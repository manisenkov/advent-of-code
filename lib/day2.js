const { readInput, range } = require("./util");
const { run } = require("./intcode");

const runSpec = async (opcodes, noun, verb) =>
  run([opcodes[0], noun, verb, ...opcodes.slice(3)]);

const findNounAndVerb = async (opcodes, output) => {
  for (const i of range(0, 100)) {
    for (const j of range(0, 100)) {
      const result = await runSpec(opcodes, i, j);
      if (result === output) {
        return [i, j];
      }
    }
  }
  return [-1, -1];
};

const main = async () => {
  const [input] = await readInput();

  const opcodes = input.split(",").map(Number);

  // Part 1
  console.log("Part 1:", await runSpec(opcodes, 12, 2));

  // Part 2
  const [noun, verb] = await findNounAndVerb(opcodes, 19690720);
  console.log("Part 2:", 100 * noun + verb);
};

main();
