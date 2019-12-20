const { entries } = require("../util");

const re = /Pos\(x=(-?\d+), ?y=(-?\d+)\)/i;

const CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

const NORTH = 1;
const SOUTH = 2;
const WEST = 3;
const EAST = 4;

const DIR_ORDER = [SOUTH, WEST, NORTH, EAST];

class Pos {
  static fromString(key) {
    const m = key.match(re);
    return new Pos(Number(m[1]), Number(m[2]));
  }

  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  move(dir) {
    switch (dir) {
      case NORTH:
        return new Pos(this.x, this.y - 1);
      case EAST:
        return new Pos(this.x + 1, this.y);
      case SOUTH:
        return new Pos(this.x, this.y + 1);
      case WEST:
        return new Pos(this.x - 1, this.y);
    }
    return null;
  }

  toString() {
    return `Pos(x=${this.x}, y=${this.y})`;
  }
}

const parseInput = (input, numLevels = 1) => {
  const height = input.length;
  const width = Math.max(...input.map(s => s.length));
  const maze = {
    passages: new Set(),
    edges: [],
    portals: {}
  };

  // Parse maze
  for (let y = 2; y < height - 2; y++) {
    for (let x = 2; x < width - 2; x++) {
      const cell = input[y][x];
      if (cell === "#") {
        continue;
      }

      if (cell !== ".") {
        continue;
      }
      const pos = new Pos(x, y);
      maze.passages.add(pos.toString());
      for (const dir of DIR_ORDER) {
        const c = pos.move(dir);
        const ci = input[c.y][c.x];
        if (ci === ".") {
          maze.edges.push([pos.toString(), c.toString()]);
        } else if (CHARS.includes(ci)) {
          let portal;
          const isOuter =
            c.y === 1 || c.y === height - 2 || c.x === 1 || c.x === width - 2;
          switch (dir) {
            case NORTH:
              portal = input[c.y - 1][c.x] + ci;
              break;
            case EAST:
              portal = ci + input[c.y][c.x + 1];
              break;
            case SOUTH:
              portal = ci + input[c.y + 1][c.x];
              break;
            case WEST:
              portal = input[c.y][c.x - 1] + ci;
              break;
          }
          maze.portals[portal] = {
            ...(maze.portals[portal] || {}),
            ...(isOuter ? { outer: pos.toString() } : { inner: pos.toString() })
          };
        }
      }
    }
  }

  // Add portal edges
  if (numLevels === 1) {
    for (const [, entrances] of entries(maze.portals)) {
      if (!entrances.inner) {
        continue;
      }
      maze.edges.push([entrances.inner, entrances.outer]);
      maze.edges.push([entrances.outer, entrances.inner]);
    }
  } else {
    // Repeat maze [numLevels] times
    maze.edges = Array(numLevels)
      .fill()
      .map((_, i) => i)
      .flatMap(i => maze.edges.map(([e1, e2]) => [`${i}:${e1}`, `${i}:${e2}`]));

    // Connect portals
    for (let i = 0; i < numLevels; i++) {
      for (const [, entrances] of entries(maze.portals)) {
        if (!entrances.inner) {
          continue;
        }
        if (i > 0) {
          maze.edges.push([
            `${i}:${entrances.outer}`,
            `${i - 1}:${entrances.inner}`
          ]);
        }
        if (i < numLevels - 1) {
          maze.edges.push([
            `${i}:${entrances.inner}`,
            `${i + 1}:${entrances.outer}`
          ]);
        }
      }
    }
  }

  return maze;
};

module.exports = {
  parseInput
};
