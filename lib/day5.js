const { readInput } = require("./util");
const { run } = require("./intcode");

const checkOutputs = outputs => {
  outputs.slice(0, outputs.length - 1).forEach((value, index) => {
    if (value !== 0) {
      throw new Error(`Error in output at index ${index}: ${value}`);
    }
  });
};

const main = async () => {
  const [input] = await readInput();

  const opcodes = input.split(",").map(Number);

  // Part 1
  const outputs1 = [];
  await run(
    opcodes,
    async () => 1,
    async n => outputs1.push(n)
  );
  checkOutputs(outputs1);
  console.log("Part 1:", outputs1.pop());

  // Part 2
  const outputs2 = [];
  await run(
    opcodes,
    async () => 5,
    async n => outputs2.push(n)
  );
  checkOutputs(outputs2);
  console.log("Part 2:", outputs2.pop());
};

main();
