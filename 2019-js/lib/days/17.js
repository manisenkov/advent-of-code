const { runCode } = require("../intcode");
const { sum } = require("../util");

// Solved manually :-)
const ROBOT_COMMANDS = [
  "A,B,A,C,A,B,C,C,A,B\n",
  "R,8,L,10,R,8\n",
  "R,12,R,8,L,8,L,12\n",
  "L,12,L,10,L,8\n",
  "n\n" // I don't need camera
];

const findIntersections = scaffold => {
  const result = [];
  const height = scaffold.length;
  const width = scaffold[0].length;
  for (let row = 0; row < height; row++) {
    for (let col = 0; col < width; col++) {
      if (scaffold[row][col] !== "#") {
        continue;
      }
      let nNeighbours = 0;
      if (row > 0 && scaffold[row - 1][col] === "#") nNeighbours++;
      if (row < height - 1 && scaffold[row + 1][col] === "#") nNeighbours++;
      if (col > 0 && scaffold[row][col - 1] === "#") nNeighbours++;
      if (col < width - 1 && scaffold[row][col + 1] === "#") nNeighbours++;
      if (nNeighbours > 2) {
        result.push([row, col]);
      }
    }
  }
  return result;
};

const wander = async opcodes => {
  let currentCommand = Array.from(ROBOT_COMMANDS[0]).map(c => c.charCodeAt(0));
  const commands = ROBOT_COMMANDS.slice(1);
  let dustCollected;
  await runCode(
    [2, ...opcodes.slice(1)],
    async () => {
      if (currentCommand.length === 0) {
        currentCommand = Array.from(commands.shift()).map(c => c.charCodeAt(0));
      }
      return currentCommand.shift();
    },
    async value => {
      dustCollected = value;
    }
  );
  return dustCollected;
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);
  const codes = [];
  await runCode(
    opcodes,
    async () => {},
    async code => {
      codes.push(code);
    }
  );
  const field = String.fromCharCode(...codes);
  const scaffold = field.split("\n");
  const dustCollected = await wander(opcodes);

  console.log(field);

  return [
    sum(findIntersections(scaffold).map(([row, col]) => row * col)),
    dustCollected
  ];
};
