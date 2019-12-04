const { readInput, range } = require("./util");

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

const main = async () => {
  const [input] = await readInput();
  const [low, high] = input.split("-").map(Number);

  // Part 1
  const passwords = getValidPasswords(low, high);
  console.log("Part 1:", passwords.length);

  // Part 2
  const stricterPasswords = passwords.filter(isStricterValid);
  console.log("Part 2:", stricterPasswords.length);
};

main();
