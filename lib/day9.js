const { readInput } = require("./util");
const { run } = require("./intcode");

const checkOutput = output => {
  const errorCodes = output.slice(0, output.length - 2);
  if (errorCodes.length > 1) {
    throw new Error(`Error in commands ${errorCodes}`);
  }
};

const main = async () => {
  const [input] = await readInput();

  const opcodes = input.split(",").map(Number);

  // Part 1
  const testOutput = [];
  await run(
    opcodes,
    async () => 1,
    async value => testOutput.push(value)
  );
  checkOutput(testOutput);
  console.log("Part 1:", testOutput[0]);

  // Part 2
  const boostOutput = [];
  await run(
    opcodes,
    async () => 2,
    async value => boostOutput.push(value)
  );
  console.log("Part 2:", boostOutput[0]);
};

main();
