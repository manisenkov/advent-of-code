const { runCode } = require("../intcode");
const { Future } = require("../util");

const runNetwork = async (opcodes, count) => {
  let currentNATState;
  let stopNAT = false;
  let lastPushedY;
  const firstNATPacketY = new Future();
  const twicePushedY = new Future();
  const terminateFutures = Array(count)
    .fill()
    .map(() => new Future());
  const inputQueues = Array(count)
    .fill()
    .map((_, i) => [i]);
  const outputQueues = Array(count)
    .fill()
    .map(() => []);
  const idleStatus = Array(count).fill(false);

  const natWorker = async () => {
    while (!stopNAT) {
      const f = Future.wait();
      await f;
      if (currentNATState && idleStatus.every(s => s)) {
        inputQueues[0].push(...currentNATState);
        if (lastPushedY === currentNATState[1]) {
          twicePushedY.resolve(lastPushedY);
        }
        [, lastPushedY] = currentNATState;
      }
    }
  };

  const inputFn = i => async () => {
    const f = Future.wait();
    await f;

    if (inputQueues[i].length === 0) {
      idleStatus[i] = true;
      return -1;
    }
    idleStatus[i] = false;
    return inputQueues[i].shift();
  };

  const outputFn = i => async value => {
    const f = Future.wait();
    await f;

    idleStatus[i] = false;

    outputQueues[i].push(value);
    if (outputQueues[i].length === 3) {
      const [address, x, y] = [
        outputQueues[i].shift(),
        outputQueues[i].shift(),
        outputQueues[i].shift()
      ];
      if (address === 255) {
        // Part 1
        if (!firstNATPacketY.isResolved) {
          firstNATPacketY.resolve(y);
        }

        // Part 2
        currentNATState = [x, y];
      } else {
        inputQueues[address].push(x);
        inputQueues[address].push(y);
      }
    }
  };

  Array(count)
    .fill()
    .forEach((_, i) => {
      runCode(opcodes, inputFn(i), outputFn(i), terminateFutures[i]);
    });
  natWorker();

  const firstY = await firstNATPacketY;
  const twiceY = await twicePushedY;

  for (const f of terminateFutures) {
    f.resolve();
  }
  stopNAT = true;

  return [firstY, twiceY];
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const [firstY, twiceY] = await runNetwork(opcodes, 50);

  return [firstY, twiceY];
};
