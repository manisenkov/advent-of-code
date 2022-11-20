const { runCode } = require("../intcode");
const { Future } = require("../util");

// The answer is "astrolabe, ornament, sand, shell"

const part1 = async opcodes => {
  let inputFuture = new Future();
  let outputFuture = new Future();
  let inputBuffer = [];
  let outputStr = "";

  const inputFn = async () => {
    if (inputBuffer.length === 0) {
      await inputFuture;
      inputFuture = new Future();
    }
    return inputBuffer.shift();
  };

  const outputFn = async code => {
    if (code === 10) {
      outputFuture.resolve(outputStr);
      outputStr = "";
      return;
    }
    outputStr += String.fromCharCode(code);
  };

  const showOutput = async () => {
    while (true) {
      const output = await outputFuture;
      outputFuture = new Future();
      console.log(output);
    }
  };

  process.stdin.on("data", chunk => {
    inputBuffer = [
      ...inputBuffer,
      ...Array.from(chunk.toString()).map(s => s.charCodeAt(0))
    ];
    if (inputBuffer[inputBuffer.length - 1] === 10) {
      inputFuture.resolve();
    }
  });

  runCode(opcodes, inputFn, outputFn);
  showOutput();
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  await part1(opcodes);

  return [0, 0];
};
