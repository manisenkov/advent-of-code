const { readInput, range } = require("./util");
const { run } = require("./intcode");

const runSpec = (opcodes, noun, verb) =>
  run([opcodes[0], noun, verb, ...opcodes.slice(3)]);

const findNounAndVerb = (opcodes, output) => {
  for (const i of range(0, 100)) {
    for (const j of range(0, 100)) {
      if (runSpec(opcodes, i, j) === output) {
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
  console.log("Part 1:", runSpec(opcodes, 12, 2));

  // Part 2
  const [noun, verb] = findNounAndVerb(opcodes, 19690720);
  console.log("Part 2:", 100 * noun + verb);
};

main();
