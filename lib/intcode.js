const { List } = require("immutable");

const CMD_SUM = 1;
const CMD_MUL = 2;
const CMD_INPUT = 3;
const CMD_OUTPUT = 4;
const CMD_JUMP_IF_TRUE = 5;
const CMD_JUMP_IF_FALSE = 6;
const CMD_IF_LESS = 7;
const CMD_IF_EQUAL = 8;
const CMD_CHANGE_BASE = 9;
const CMD_TERMINATE = 99;

const MODE_POINTER = 0;
const MODE_IMMEDIATE = 1;
const MODE_RELATIVE = 2;

const getArgs = ({ mem, pointer, count, base, modes }) =>
  mem
    .slice(pointer + 1, pointer + count + 1)
    .zip(modes)
    .map(([pArg, mode]) => pArg + (mode & MODE_RELATIVE ? base : 0))
    .zip(modes)
    .map(([pArg, mode]) =>
      ({
        [MODE_POINTER]: () => mem.get(pArg) || 0,
        [MODE_IMMEDIATE]: () => pArg || 0
      }[mode & MODE_IMMEDIATE]())
    );

const getCommandCode = opcode => opcode % 100;

const getModes = opcode => [
  ((opcode / 100) | 0) % 10,
  ((opcode / 1000) | 0) % 10,
  ((opcode / 10000) | 0) % 10
];

const commands = {
  [CMD_SUM]: async ({ mem, pointer, base, modes }) => {
    const [val1, val2, pResult] = getArgs({
      mem,
      pointer,
      count: 3,
      base,
      modes: [modes[0], modes[1], modes[2] | MODE_IMMEDIATE]
    });
    return [mem.set(pResult, val1 + val2), pointer + 4, base];
  },
  [CMD_MUL]: async ({ mem, pointer, base, modes }) => {
    const [val1, val2, pResult] = getArgs({
      mem,
      pointer,
      count: 3,
      base,
      modes: [modes[0], modes[1], modes[2] | MODE_IMMEDIATE]
    });
    return [mem.set(pResult, val1 * val2), pointer + 4, base];
  },
  [CMD_INPUT]: async ({ mem, pointer, base, modes, input }) => {
    const [pResult] = getArgs({
      mem,
      pointer,
      count: 1,
      base,
      modes: [modes[0] | MODE_IMMEDIATE]
    });
    const inputValue = await input();
    return [mem.set(pResult, inputValue), pointer + 2, base];
  },
  [CMD_OUTPUT]: async ({ mem, pointer, base, modes, output }) => {
    const [val] = getArgs({ mem, pointer, count: 1, base, modes });
    await output(val);
    return [mem, pointer + 2, base];
  },
  [CMD_JUMP_IF_TRUE]: async ({ mem, pointer, base, modes }) => {
    const [checkVal, nextPointer] = getArgs({
      mem,
      pointer,
      count: 2,
      base,
      modes
    });
    return [mem, checkVal ? nextPointer : pointer + 3, base];
  },
  [CMD_JUMP_IF_FALSE]: async ({ mem, pointer, base, modes }) => {
    const [checkVal, nextPointer] = getArgs({
      mem,
      pointer,
      count: 2,
      base,
      modes
    });
    return [mem, !checkVal ? nextPointer : pointer + 3, base];
  },
  [CMD_IF_LESS]: async ({ mem, pointer, base, modes }) => {
    const [val1, val2, pResult] = getArgs({
      mem,
      pointer,
      count: 3,
      base,
      modes: [modes[0], modes[1], modes[2] | MODE_IMMEDIATE]
    });
    return [mem.set(pResult, val1 < val2 ? 1 : 0), pointer + 4, base];
  },
  [CMD_IF_EQUAL]: async ({ mem, pointer, base, modes }) => {
    const [val1, val2, pResult] = getArgs({
      mem,
      pointer,
      count: 3,
      base,
      modes: [modes[0], modes[1], modes[2] | MODE_IMMEDIATE]
    });
    return [mem.set(pResult, val1 === val2 ? 1 : 0), pointer + 4, base];
  },
  [CMD_CHANGE_BASE]: async ({ mem, pointer, base, modes }) => {
    const [val] = getArgs({ mem, pointer, count: 1, base, modes });
    return [mem, pointer + 2, base + val];
  }
};

exports.runCode = async (opcodes, input, output, terminateFuture = null) => {
  let mem = List(opcodes);
  let pointer = 0;
  let base = 0;
  let isTerminated = false;

  if (terminateFuture) {
    (async () => {
      await terminateFuture;
      isTerminated = true;
    })();
  }

  while (!isTerminated && mem.get(pointer) !== CMD_TERMINATE) {
    const opcode = mem.get(pointer);
    const commandCode = getCommandCode(opcode);
    const modes = getModes(opcode);
    const cmdFn = commands[commandCode];
    if (!cmdFn) {
      return -1;
    }
    [mem, pointer, base] = await cmdFn({
      mem,
      pointer,
      base,
      modes,
      input,
      output
    });
  }

  if (!isTerminated && terminateFuture) {
    terminateFuture.resolve();
  }

  return mem.get(0);
};
