const { readInput } = require("./util");
const { maximizeThrust } = require("./thruster");

const main = async () => {
  const [input] = await readInput();

  const opcodes = input.split(",").map(Number);

  // Part 1
  const phases = [0, 1, 2, 3, 4];
  const [maxThrust] = await maximizeThrust(opcodes, phases);
  console.log("Part 1:", maxThrust);

  // Part 2
  const phasesWithFeedback = [5, 6, 7, 8, 9];
  const [maxThrustWithFeedback] = await maximizeThrust(
    opcodes,
    phasesWithFeedback
  );
  console.log("Part 2:", maxThrustWithFeedback);
};

main();
