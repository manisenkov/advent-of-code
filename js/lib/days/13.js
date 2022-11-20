const { Set } = require("immutable");

const { runCode } = require("../intcode");

const EMPTY = 0;
const WALL = 1;
const BLOCK = 2;
const PADDLE = 3;
const BALL = 4;

class Tile {
  constructor({ row, col, type }) {
    this.row = row;
    this.col = col;
    this.type = type;
  }

  equals(other) {
    return (
      this.row === other.row &&
      this.col === other.col &&
      this.type === other.type
    );
  }

  hashCode() {
    return (((this.row * 31) & this.col) * 31) & this.type;
  }

  toString() {
    return `Tile(row=${this.row}, col=${this.col}, type=${this.type})`;
  }
}

const calcDisplaySize = tileSet => {
  let maxRow = -Infinity;
  let maxCol = -Infinity;
  for (const { row, col } of tileSet) {
    if (row > maxRow) maxRow = row;
    if (col > maxCol) maxCol = col;
  }
  return {
    width: maxCol + 1,
    height: maxRow + 1
  };
};

const display = tileSet => {
  const { width, height } = calcDisplaySize(tileSet);
  const image = Array(height)
    .fill(null)
    .map(() => Array(width).fill(" "));
  for (const { row, col, type } of tileSet) {
    let symbol;
    switch (type) {
      case EMPTY:
        symbol = " ";
        break;
      case WALL:
        symbol = "█";
        break;
      case BLOCK:
        symbol = "■";
        break;
      case PADDLE:
        symbol = "▄";
        break;
      case BALL:
        symbol = "o";
        break;
    }
    image[row][col] = symbol;
  }
  return image.map(row => row.join("")).join("\n");
};

const steer = (ballTile, paddleTile) => {
  if (ballTile.col < paddleTile.col) {
    return -1;
  }
  if (ballTile.col > paddleTile.col) {
    return 1;
  }
  return 0;
};

const calcField = async opcodes => {
  let buffer = [];
  let tileSet = Set();
  await runCode(
    opcodes,
    async () => {},
    async value => {
      buffer.push(value);
      if (buffer.length === 3) {
        const [x, y, val] = buffer;
        const tile = new Tile({ row: y, col: x, type: val });
        tileSet = tileSet.add(tile);
        buffer = [];
      }
    }
  );
  return tileSet;
};

const playGame = async (opcodes, width, height) => {
  let buffer = [];
  let score = 0;
  let ballTile = null;
  let paddleTile = null;
  const field = Array(height)
    .fill(null)
    .map((_, row) =>
      Array(width)
        .fill(null)
        .map((__, col) => new Tile(row, col, EMPTY))
    );

  await runCode(
    [2, ...opcodes.slice(1)],
    async () => steer(ballTile, paddleTile),
    async value => {
      buffer.push(value);
      if (buffer.length === 3) {
        const [x, y, val] = buffer;
        if (x === -1 && y === 0) {
          score = val;
        } else {
          const tile = new Tile({ row: y, col: x, type: val });
          if (val === BALL) {
            ballTile = tile;
          }
          if (val === PADDLE) {
            paddleTile = tile;
          }
          field[y][x] = tile;
        }
        buffer = [];
      }
    }
  );

  return score;
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  // Determine blocks count
  const tileSet = await calcField(opcodes);
  const { width, height } = calcDisplaySize(tileSet);
  const blocksCount = tileSet.filter(tile => tile.type === BLOCK).size;

  console.log(display(tileSet));

  // Play the game
  const score = await playGame(opcodes, width, height);

  return [blocksCount, score];
};
