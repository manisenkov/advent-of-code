const { update } = require("./util");

const CMD_SUM = 1;
const CMD_MUL = 2;
const CMD_INPUT = 3;
const CMD_OUTPUT = 4;
const CMD_JUMP_IF_TRUE = 5;
const CMD_JUMP_IF_FALSE = 6;
const CMD_IF_LESS = 7;
const CMD_IF_EQUAL = 8;
const CMD_TERMINATE = 99;

const getArgs = (mem, pointer, count) =>
  mem.slice(pointer + 1, pointer + count + 1);

const getCommandCode = opcode => opcode % 100;

const getModes = opcode => [
  ((opcode / 100) | 0) % 10,
  ((opcode / 1000) | 0) % 10,
  ((opcode / 10000) | 0) % 10
];

const commands = {
  [CMD_SUM]: (mem, pointer, [modeArg1, modeArg2]) => {
    const [arg1, arg2, pResult] = getArgs(mem, pointer, 3);
    const val1 = modeArg1 ? arg1 : mem[arg1];
    const val2 = modeArg2 ? arg2 : mem[arg2];
    return [update(mem, pResult, val1 + val2), pointer + 4];
  },
  [CMD_MUL]: (mem, pointer, [modeArg1, modeArg2]) => {
    const [arg1, arg2, pResult] = getArgs(mem, pointer, 3);
    const val1 = modeArg1 ? arg1 : mem[arg1];
    const val2 = modeArg2 ? arg2 : mem[arg2];
    return [update(mem, pResult, val1 * val2), pointer + 4];
  },
  [CMD_INPUT]: (mem, pointer, _, input) => {
    const [pResult] = getArgs(mem, pointer, 1);
    return [update(mem, pResult, input()), pointer + 2];
  },
  [CMD_OUTPUT]: (mem, pointer, [modeArg], _, output) => {
    const [arg] = getArgs(mem, pointer, 1);
    output(modeArg ? arg : mem[arg]);
    return [mem, pointer + 2];
  },
  [CMD_JUMP_IF_TRUE]: (mem, pointer, [modeArg, modePointer]) => {
    const [arg, pNext] = getArgs(mem, pointer, 2);
    const checkVal = modeArg ? arg : mem[arg];
    const nextPointer = modePointer ? pNext : mem[pNext];
    return [mem, checkVal ? nextPointer : pointer + 3];
  },
  [CMD_JUMP_IF_FALSE]: (mem, pointer, [modeArg, modePointer]) => {
    const [arg, pNext] = getArgs(mem, pointer, 2);
    const checkVal = modeArg ? arg : mem[arg];
    const nextPointer = modePointer ? pNext : mem[pNext];
    return [mem, !checkVal ? nextPointer : pointer + 3];
  },
  [CMD_IF_LESS]: (mem, pointer, [modeArg1, modeArg2]) => {
    const [arg1, arg2, pResult] = getArgs(mem, pointer, 3);
    const val1 = modeArg1 ? arg1 : mem[arg1];
    const val2 = modeArg2 ? arg2 : mem[arg2];
    return [update(mem, pResult, val1 < val2 ? 1 : 0), pointer + 4];
  },
  [CMD_IF_EQUAL]: (mem, pointer, [modeArg1, modeArg2]) => {
    const [arg1, arg2, pResult] = getArgs(mem, pointer, 3);
    const val1 = modeArg1 ? arg1 : mem[arg1];
    const val2 = modeArg2 ? arg2 : mem[arg2];
    return [update(mem, pResult, val1 === val2 ? 1 : 0), pointer + 4];
  }
};

module.exports.run = (opcodes, input, output) => {
  let mem = [...opcodes];
  let pointer = 0;

  while (mem[pointer] !== CMD_TERMINATE) {
    const opcode = mem[pointer];
    const commandCode = getCommandCode(opcode);
    const modes = getModes(opcode);
    const cmdFn = commands[commandCode];
    if (!cmdFn) {
      return -1;
    }
    [mem, pointer] = cmdFn(mem, pointer, modes, input, output);
  }
  return mem[0];
};
