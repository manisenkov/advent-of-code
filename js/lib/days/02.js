const { Range } = require("immutable");
const { runCode } = require("../intcode");

const runSpec = async (opcodes, noun, verb) =>
  runCode([opcodes[0], noun, verb, ...opcodes.slice(3)]);

const findNounAndVerb = async (opcodes, output) => {
  for (const noun of Range(0, 100)) {
    for (const verb of Range(0, 100)) {
      const result = await runSpec(opcodes, noun, verb);
      if (result === output) {
        return [noun, verb];
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
