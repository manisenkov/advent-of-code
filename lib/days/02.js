const { range } = require("../util");
const { runCode } = require("../intcode");

const runSpec = async (opcodes, noun, verb) =>
  runCode([opcodes[0], noun, verb, ...opcodes.slice(3)]);

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

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const part1 = await runSpec(opcodes, 12, 2);
  const [noun, verb] = await findNounAndVerb(opcodes, 19690720);
  const part2 = 100 * noun + verb;
  return [part1, part2];
};
