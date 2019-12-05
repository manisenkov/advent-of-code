const readline = require("readline");

module.exports.compose = (...fns) => (...args) => {
  let result = fns[0](...args);
  for (const fn of fns.slice(1)) {
    result = fn(result);
  }
  return result;
};

module.exports.forEach = (iter, fn) => {
  for (const el of iter) {
    fn(el);
  }
};

function* range(start, stop) {
  let current = start;
  if (stop > start) {
    while (current < stop) {
      yield current;
      current++;
    }
  } else {
    while (current > stop) {
      yield current;
      current--;
    }
  }
}
module.exports.range = range;

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

module.exports.update = (arr, index, value) => {
  const result = [...arr];
  result[index] = value;
  return result;
};
