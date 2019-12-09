const { range } = require("../util");

const isValid = password => {
  let hasDouble = false;
  let current = password[0];
  for (const c of password.slice(1)) {
    if (Number(c) < Number(current)) {
      return false;
    }
    if (c === current) {
      hasDouble = true;
    }
    current = c;
  }
  return hasDouble;
};

const isStricterValid = password => {
  const charCounts = Array.from(password).reduce(
    (counts, c) => ({ ...counts, [c]: (counts[c] || 0) + 1 }),
    {}
  );
  return Object.values(charCounts).filter(c => c === 2).length !== 0;
};

const getValidPasswords = (low, high) => {
  return Array.from(range(low, high + 1))
    .map(String)
    .filter(isValid);
};

exports.run = async ([input]) => {
  const [low, high] = input.split("-").map(Number);

  const passwords = getValidPasswords(low, high);
  const part1 = passwords.length;
  const stricterPasswords = passwords.filter(isStricterValid);
  const part2 = stricterPasswords.length;

  return [part1, part2];
};
