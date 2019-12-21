const { runCode } = require("../intcode");

const WALK_PROGRAM = [
  "NOT A J\n",
  "NOT B T\n",
  "OR T J\n",
  "NOT C T\n",
  "OR T J\n",
  "AND D J\n",
  "WALK\n"
];

const RUN_PROGRAM = [
  "NOT A J\n",
  "NOT B T\n",
  "OR T J\n",
  "NOT C T\n",
  "OR T J\n",
  "AND D J\n",
  "NOT D T\n",
  "OR H T\n",
  "OR E T\n",
  "AND T J\n",
  "RUN\n"
];

const walk = async opcodes => {
  let currentCommand = Array.from(WALK_PROGRAM[0]).map(c => c.charCodeAt(0));
  const commands = WALK_PROGRAM.slice(1);
  let hullDamage = -1;
  let output = "";
  await runCode(
    opcodes,
    async () => {
      if (currentCommand.length === 0) {
        currentCommand = Array.from(commands.shift()).map(c => c.charCodeAt(0));
      }
      return currentCommand.shift();
    },
    async value => {
      if (value > 255) {
        hullDamage = value;
      } else {
        output += String.fromCharCode(value);
      }
    }
  );
  if (output !== "") {
    console.log(output);
  }
  return hullDamage;
};

const run = async opcodes => {
  let currentCommand = Array.from(RUN_PROGRAM[0]).map(c => c.charCodeAt(0));
  const commands = RUN_PROGRAM.slice(1);
  let hullDamage = -1;
  let output = "";
  await runCode(
    opcodes,
    async () => {
      if (currentCommand.length === 0) {
        currentCommand = Array.from(commands.shift()).map(c => c.charCodeAt(0));
      }
      return currentCommand.shift();
    },
    async value => {
      if (value > 255) {
        hullDamage = value;
      } else {
        output += String.fromCharCode(value);
      }
    }
  );
  if (output !== "") {
    console.log(output);
  }
  return hullDamage;
};

exports.run = async ([input]) => {
  const opcodes = input.split(",").map(Number);

  const hullDamage = await walk(opcodes);
  const hullDamageRun = await run(opcodes);

  return [hullDamage, hullDamageRun];
};
