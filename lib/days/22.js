/* global BigInt */

const { compose } = require("../util");

const N_CARDS = BigInt(119315717514047);
const N_SHUFFLES = BigInt(101741582076661);

const cut = (count, deckSize) => pos =>
  count < 0
    ? cut(deckSize + count, deckSize)(pos)
    : (pos + (deckSize - count)) % deckSize;

const dealIntoNewStack = deckSize => pos => deckSize - pos - 1;

const dealWithIncrement = (count, deckSize) => pos => (pos * count) % deckSize;

const parseCommand = deckSize => s => {
  if (s.startsWith("cut")) {
    return cut(Number(s.slice("cut".length + 1)), deckSize);
  }
  if (s.startsWith("deal with increment")) {
    return dealWithIncrement(
      Number(s.slice("deal with increment".length + 1)),
      deckSize
    );
  }
  return dealIntoNewStack(deckSize);
};

function powerMod(n, exp, mod) {
  let b = n;
  let p = exp;
  let result = 1n;
  while (p > 0n) {
    if (p % 2n === 1n) {
      result = (result * b) % mod;
    }
    p /= 2n;
    b = b ** 2n % mod;
  }
  return (mod + result) % mod;
}

const extGCD = (a, b) => {
  if (a === 0n) {
    return [b, 0n, 1n];
  }
  const [g, y, x] = extGCD(b % a, a);
  return [g, x - ((b / a) | 0n) * y, y];
};

const modInv = (a, mod) => {
  const [g, x] = extGCD(a, mod);
  if (g !== 1n) {
    return undefined;
  }
  return x % mod;
};

const calcReverseShuffleParams = input => {
  let a = 1n;
  let b = 0n;
  for (const s of [...input].reverse()) {
    if (s.startsWith("cut")) {
      const arg = BigInt(s.slice("cut".length + 1));
      b += arg;
    } else if (s.startsWith("deal with increment")) {
      const arg = BigInt(s.slice("deal with increment".length + 1));
      const m = modInv(arg, N_CARDS);
      a *= m;
      b *= m;
    } else {
      a = -a;
      b = -b - 1n;
    }
    [a, b] = [(N_CARDS + a) % N_CARDS, (N_CARDS + b) % N_CARDS];
  }
  return [a, b];
};

const calcPart1 = input => {
  const commands = compose(...input.map(parseCommand(10007)));
  return commands(2019);
};

const calcPart2 = input => {
  const [a, b] = calcReverseShuffleParams(input);
  const pa = powerMod(a, N_SHUFFLES, N_CARDS);
  const pb =
    b *
    (powerMod(a, N_SHUFFLES, N_CARDS) - 1n) *
    ((N_CARDS + modInv(a - 1n, N_CARDS)) % N_CARDS);
  return (N_CARDS + (pa * 2020n + pb)) % N_CARDS;
};

exports.run = async input => {
  const part1 = calcPart1(input);
  const part2 = calcPart2(input);

  return [part1, part2];
};
