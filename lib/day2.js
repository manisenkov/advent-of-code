const { readInput, range } = require("./util");

const run = (opcodes, noun, verb) => {
  const mem = [...opcodes];
  let pointer = 0;

  mem[1] = noun;
  mem[2] = verb;

  while (mem[pointer] !== 99) {
    if ([1, 2].includes(mem[pointer])) {
      const [pointer1, pointer2, pointerResult] = [
        mem[pointer + 1],
        mem[pointer + 2],
        mem[pointer + 3]
      ];
      if (mem[pointer] === 1) {
        mem[pointerResult] = mem[pointer1] + mem[pointer2];
      } else {
        mem[pointerResult] = mem[pointer1] * mem[pointer2];
      }
      pointer += 4;
    } else {
      return -1; // Wrong instruction
    }
  }
  return mem[0];
};

const findNounAndVerb = (opcodes, output) => {
  for (const i of range(0, 100)) {
    for (const j of range(0, 100)) {
      if (run(opcodes, i, j) === output) {
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
  console.log("Part 1:", run(opcodes, 12, 2));

  // Part 2
  const [noun, verb] = findNounAndVerb(opcodes, 19690720);
  console.log("Part 2:", 100 * noun + verb);
};

main();
