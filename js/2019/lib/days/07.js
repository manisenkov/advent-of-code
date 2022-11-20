const { maximizeThrust } = require("../thruster");

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const phases = [0, 1, 2, 3, 4];
  const [part1] = await maximizeThrust(opcodes, phases);

  const phasesWithFeedback = [5, 6, 7, 8, 9];
  const [part2] = await maximizeThrust(opcodes, phasesWithFeedback);

  return [part1, part2];
};
