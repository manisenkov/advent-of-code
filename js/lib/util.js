const readline = require("readline");

const compose = (...fns) => (...args) => {
  let result = fns[0](...args);
  for (const fn of fns.slice(1)) {
    result = fn(result);
  }
  return result;
};

const cross = (set1, set2) =>
  new Set([...set1, ...set2].filter(s => set1.has(s) && set2.has(s)));

const ident = x => x;

const floatEq = (x, y) => Math.abs(x - y) < 0.001;

const angle = (y, x) => {
  if (y >= 0 && x >= 0) return Math.atan(y / x);
  if (y >= 0 && x < 0) return Math.PI / 2 + Math.atan(-x / y);
  if (y < 0 && x < 0) return Math.PI + Math.atan(-y / -x);
  return (3 * Math.PI) / 2 + Math.atan(x / -y);
};

const dist = (y, x) => Math.sqrt(y ** 2 + x ** 2);

const manhattan = ([x, y]) => Math.abs(x) + Math.abs(y);

const gcd = (a, b) => (b === 0 ? a : gcd(b, a % b));

const lcm = (a, b) => Math.abs(a * b) / gcd(a, b);

const max = (iter, keyFn = ident) => {
  let isFirstEl = true;
  let maxKey;
  let maxEl;
  for (const el of iter) {
    if (isFirstEl) {
      isFirstEl = false;
      maxKey = keyFn(el);
      maxEl = el;
    } else {
      const key = keyFn(el);
      if (key > maxKey) {
        maxKey = key;
        maxEl = el;
      }
    }
  }
  return maxEl;
};

const min = (iter, keyFn = ident) => {
  let isFirstEl = true;
  let minKey;
  let minEl;
  for (const el of iter) {
    if (isFirstEl) {
      isFirstEl = false;
      minKey = keyFn(el);
      minEl = el;
    } else {
      const key = keyFn(el);
      if (key < minKey) {
        minKey = key;
        minEl = el;
      }
    }
  }
  return minEl;
};

function* range(start = 0, stop = Infinity, step = undefined) {
  let i = start;
  if (typeof step === "undefined") {
    step = stop >= start ? 1 : -1;
  }
  while (i < stop) {
    yield i;
    i += step;
  }
}

function* permutate(arr) {
  if (arr.length <= 1) {
    yield arr;
  } else {
    for (const i of range(0, arr.length)) {
      const el = arr[i];
      for (const rest of permutate([...arr.slice(0, i), ...arr.slice(i + 1)])) {
        yield [el, ...rest];
      }
    }
  }
}

const readInput = async inputStream => {
  const rl = readline.createInterface({
    input: inputStream
  });
  const result = [];
  for await (const line of rl) {
    result.push(line);
  }
  return result;
};

const sum = arr => arr.reduce((acc, m) => acc + m, 0);

const union = (set1, set2) => new Set([...set1, ...set2]);

const update = (arr, index, value) => {
  const result = [...arr];
  result[index] = value;
  return result;
};

const subtract = (set1, set2) => new Set([...set1].filter(s => !set2.has(s)));

class Future {
  static wait(ms = 0) {
    const f = new Future();
    setTimeout(f.resolve, ms);
    return f;
  }

  constructor() {
    this.isResolved = false;
    const p = new Promise((resolve, reject) => {
      this.resolve = result => {
        resolve(result);
        this.isResolved = true;
      };
      this.reject = reject;
    });
    this.then = p.then.bind(p);
    this.catch = p.catch.bind(p);
  }
}

module.exports = {
  Future,
  angle,
  compose,
  cross,
  dist,
  floatEq,
  gcd,
  lcm,
  manhattan,
  max,
  min,
  permutate,
  range,
  readInput,
  sum,
  subtract,
  union,
  update
};
