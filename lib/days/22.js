const { compose } = require("../util");

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

exports.run = async input => {
  const commands = compose(...input.map(parseCommand(10007)));
  const part1 = commands(2019);

  return [part1, 0];
};
