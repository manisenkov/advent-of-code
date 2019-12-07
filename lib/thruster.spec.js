const { calcThrust, maximizeThrust } = require("./thruster");
const { range } = require("./util");

const programs = JSON.parse(
  `[
  [3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0],
  [3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0],
  [3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0]
]`
);

describe("calcThrust", () => {
  it("should return correct thruster signal", async () => {
    const phaseInputs = [
      [4, 3, 2, 1, 0],
      [0, 1, 2, 3, 4],
      [1, 0, 4, 3, 2]
    ];
    const expectedOutputs = [43210, 54321, 65210];
    for (const i of range(0, 3)) {
      const opcodes = programs[i];
      const phaseSettings = phaseInputs[i];
      const expectedOutput = expectedOutputs[i];
      expect(await calcThrust(opcodes, phaseSettings)).toEqual(expectedOutput);
    }
  });
});

describe("maximizeThrust", () => {
  it("should return max thruster signal and phase settings", async () => {
    const expectedOutputs = [
      [43210, [4, 3, 2, 1, 0]],
      [54321, [0, 1, 2, 3, 4]],
      [65210, [1, 0, 4, 3, 2]]
    ];
    for (const i of range(0, 3)) {
      expect(await maximizeThrust(programs[i], [0, 1, 2, 3, 4])).toEqual(
        expectedOutputs[i]
      );
    }
  });
});
