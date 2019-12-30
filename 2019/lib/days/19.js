const { runCode } = require("../intcode");
const { Future } = require("../util");

const SHIP_SIZE = 100;

const getSignalFactory = opcodes => async (x, y) => {
  const resultFuture = new Future();
  let isX = true;
  runCode(
    opcodes,
    async () => {
      if (isX) {
        isX = !isX;
        return x;
      }
      return y;
    },
    async value => {
      resultFuture.resolve(value);
    }
  );
  const result = await resultFuture;
  return result;
};

const part1 = async opcodes => {
  const getSignal = getSignalFactory(opcodes);
  let beamCount = 1; // include [0; 0]
  for (let x = 1; x < 50; x++) {
    for (let y = 1; y < 50; y++) {
      beamCount += await getSignal(x, y);
    }
  }
  return beamCount;
};

const part2 = async opcodes => {
  const getSignal = getSignalFactory(opcodes);
  let [x, y] = [1, SHIP_SIZE];
  while (true) {
    const leftSignal = await getSignal(x, y);
    if (leftSignal) {
      const topSignal = await getSignal(
        x + (SHIP_SIZE - 1),
        y - (SHIP_SIZE - 1)
      );
      if (topSignal) {
        break;
      }
      y++;
    } else {
      x++;
    }
  }
  return [x, y - (SHIP_SIZE - 1)];
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);
  const beamCount = await part1(opcodes);
  const [left, top] = await part2(opcodes);
  return [beamCount, left * 10000 + top];
};
