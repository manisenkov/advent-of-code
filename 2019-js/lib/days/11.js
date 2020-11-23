const { Set } = require("immutable");

const { runCode } = require("../intcode");

const UP = { dRow: -1, dCol: 0 };
const RIGHT = { dRow: 0, dCol: 1 };
const DOWN = { dRow: 1, dCol: 0 };
const LEFT = { dRow: 0, dCol: -1 };

const turn = (dir, input) => {
  if (input === 0) {
    // Turn left
    switch (dir) {
      case UP:
        return LEFT;
      case LEFT:
        return DOWN;
      case DOWN:
        return RIGHT;
      case RIGHT:
        return UP;
    }
  } else {
    // Turn right
    switch (dir) {
      case UP:
        return RIGHT;
      case RIGHT:
        return DOWN;
      case DOWN:
        return LEFT;
      case LEFT:
        return UP;
    }
  }
  return null;
};

class Tile {
  constructor({ row, col }) {
    this.row = row;
    this.col = col;
  }

  hashCode() {
    return (this.row << 5) & this.col;
  }

  equals(other) {
    return this.row === other.row && this.col === other.col;
  }

  toString() {
    return `Tile(row=${this.row}, col=${this.col})`;
  }
}

const paint = async (opcodes, startWhiteTiles) => {
  let pos = { row: 0, col: 0 };
  let whiteTiles = startWhiteTiles;
  let changedTiles = Set();
  let outputMode = 0;
  let dir = UP;
  await runCode(
    opcodes,
    async () => (whiteTiles.has(new Tile(pos)) ? 1 : 0),
    async output => {
      if (outputMode === 0) {
        // Change color
        const tile = new Tile(pos);
        if (output === 0) {
          whiteTiles = whiteTiles.remove(tile);
        } else {
          whiteTiles = whiteTiles.add(tile);
        }
        changedTiles = changedTiles.add(tile);
        outputMode = 1;
      } else {
        // Change direction
        dir = turn(dir, output);
        pos = { row: pos.row + dir.dRow, col: pos.col + dir.dCol };
        outputMode = 0;
      }
    }
  );
  return [whiteTiles, changedTiles];
};

const drawTiles = whiteTiles => {
  let minRow = +Infinity;
  let minCol = +Infinity;
  let maxRow = -Infinity;
  let maxCol = -Infinity;
  for (const { row, col } of whiteTiles) {
    if (row < minRow) minRow = row;
    if (col < minCol) minCol = col;
    if (row > maxRow) maxRow = row;
    if (col > maxCol) maxCol = col;
  }
  const image = Array(maxRow - minRow + 1)
    .fill(null)
    .map(() => Array(maxCol - minCol + 1).fill(" "));
  for (const { row, col } of whiteTiles) {
    image[row - minRow][col - minCol] = "█";
  }
  return image.map(row => row.join("")).join("\n");
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const changedTiles = (await paint(opcodes, Set()))[1];
  const whiteTiles = (
    await paint(opcodes, Set([new Tile({ row: 0, col: 0 })]))
  )[0];
  const image = drawTiles(whiteTiles);

  return [changedTiles.size, image];
};
