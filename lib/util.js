const readline = require("readline");

module.exports.readInput = async () => {
  const rl = readline.createInterface({
    input: process.stdin
  });
  const result = [];
  for await (const line of rl) {
    result.push(line);
  }
  return result;
};

module.exports.sum = arr => arr.reduce((acc, m) => acc + m, 0);

function* range(start, stop) {
  let current = start;
  while (current < stop) {
    yield current;
    current++;
  }
}
module.exports.range = range;
