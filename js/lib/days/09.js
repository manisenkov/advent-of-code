const { runCode } = require("../intcode");

const checkOutput = output => {
  const errorCodes = output.slice(0, output.length - 2);
  if (errorCodes.length > 1) {
    throw new Error(`Error in commands ${errorCodes}`);
  }
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);
  const output = [];

  await runCode(
    opcodes,
    async () => 1,
    async value => output.push(value)
  );
  checkOutput(output);
  const part1 = output.pop();

  await runCode(
    opcodes,
    async () => 2,
    async value => output.push(value)
  );
  const part2 = output.pop();

  return [part1, part2];
};
