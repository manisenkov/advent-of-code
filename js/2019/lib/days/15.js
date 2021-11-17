const { createGraph, calcGraphDist } = require("../graphs");
const { Future, max } = require("../util");
const { runCode } = require("../intcode");

const NORTH = 1;
const SOUTH = 2;
const WEST = 3;
const EAST = 4;

const WALL = 0;
const EMPTY = 1;
const OXYGEN = 2;

const DIR_ORDER = [SOUTH, WEST, NORTH, EAST];

const re = /Pos\(x=(-?\d+), ?y=(-?\d+)\)/i;

class Pos {
  static parse(s) {
    const m = s.match(re);
    return new Pos(Number(m[1]), Number(m[2]));
  }

  constructor(x, y) {
    this.x = x;
    this.y = y;
  }

  equals(other) {
    return this.x === other.x && this.y === other.y;
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

const ZERO_POS = new Pos(0, 0);

const uturn = dir =>
  ({
    [NORTH]: SOUTH,
    [EAST]: WEST,
    [SOUTH]: NORTH,
    [WEST]: EAST
  }[dir]);

const discover = async opcodes => {
  let inputFuture = new Future();
  let outputFuture = new Future();
  const cells = new Map([[ZERO_POS.toString(), EMPTY]]);
  const tracks = [{ pos: ZERO_POS }];

  const where = () => {
    const currentPos = tracks[tracks.length - 1].pos;
    for (const dir of DIR_ORDER) {
      const destPos = currentPos.move(dir);
      const destKey = destPos.toString();
      const destCell = cells.get(destKey);
      if (typeof destCell !== "undefined") {
        continue;
      }
      return { destPos, dir };
    }
    return null;
  };

  const move = async dir => {
    inputFuture.resolve(dir);
    const result = await outputFuture;
    outputFuture = new Future();
    return result;
  };

  // Run IntCode computer in parallel (without async)
  runCode(
    opcodes,
    async () => {
      const dir = await inputFuture;
      inputFuture = new Future();
      return dir;
    },
    async res => {
      outputFuture.resolve(res);
    }
  );

  // Main loop
  while (tracks.length > 0) {
    const whereTo = where();
    if (whereTo == null) {
      const { dir } = tracks.pop();
      if (dir != null) {
        await move(uturn(dir));
      }
      continue;
    }
    const { destPos, dir } = whereTo;
    const destKey = destPos.toString();
    const moveResult = await move(dir);
    cells.set(destKey, moveResult);
    if (moveResult === EMPTY || moveResult === OXYGEN) {
      tracks.push({ pos: destPos, dir });
    }
  }

  return cells;
};

const drawMap = cells => {
  let minX = +Infinity;
  let minY = +Infinity;
  let maxX = -Infinity;
  let maxY = -Infinity;
  for (const [posStr] of cells) {
    const pos = Pos.parse(posStr);
    if (pos.x < minX) minX = pos.x;
    if (pos.y < minY) minY = pos.y;
    if (pos.x > maxX) maxX = pos.x;
    if (pos.y > maxY) maxY = pos.y;
  }
  const image = Array(maxY - minY + 1)
    .fill(null)
    .map(() => Array(maxX - minX + 1).fill(" "));
  for (const [posStr, type] of cells) {
    const pos = Pos.parse(posStr);
    let symbol;
    if (pos.x === 0 && pos.y === 0) {
      symbol = "D";
    } else {
      switch (type) {
        case WALL:
          symbol = "█";
          break;
        case EMPTY:
          symbol = ".";
          break;
        case OXYGEN:
          symbol = "#";
          break;
      }
    }
    image[pos.y - minY][pos.x - minX] = symbol;
  }
  return image.map(row => row.join("")).join("\n");
};

const createCellGraph = cells => {
  const edges = [];
  for (const [posKey, type] of cells) {
    if (type === WALL) {
      continue;
    }
    const pos = Pos.parse(posKey);
    for (const dir of DIR_ORDER) {
      const destKey = pos.move(dir).toString();
      const destCell = cells.get(destKey);
      if (posKey === destKey) {
        continue;
      }
      if (destCell === EMPTY || destCell === OXYGEN) {
        edges.push([posKey, destKey]);
      }
    }
  }
  return createGraph(edges);
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const cells = await discover(opcodes, ZERO_POS);
  const oxygenCell = Array.from(cells).find(r => r[1] === OXYGEN)[0];
  const cellGraph = createCellGraph(cells);
  const distToOxygen = calcGraphDist(cellGraph, ZERO_POS.toString())[
    oxygenCell
  ];
  const fillTime = max(
    Object.entries(calcGraphDist(cellGraph, oxygenCell)),
    r => r[1]
  )[1];

  console.log(drawMap(cells));

  return [distToOxygen, fillTime];
};
