const { runCode } = require("../intcode");

const checkOutputs = outputs => {
  outputs.slice(0, outputs.length - 1).forEach((value, index) => {
    if (value !== 0) {
      throw new Error(`Error in output at index ${index}: ${value}`);
    }
  });
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);
  const output = [];

  await runCode(
    opcodes,
    async () => 1,
    async n => output.push(n)
  );
  checkOutputs(output);
  const part1 = output.pop();

  await runCode(
    opcodes,
    async () => 5,
    async n => output.push(n)
  );
  checkOutputs(output);
  const part2 = output.pop();

  return [part1, part2];
};
