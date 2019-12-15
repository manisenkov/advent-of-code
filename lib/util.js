const readline = require("readline");
const { Range } = require("immutable");

const { entries, fromEntries, keys, values } = Object;

const compose = (...fns) => (...args) => {
  let result = fns[0](...args);
  for (const fn of fns.slice(1)) {
    result = fn(result);
  }
  return result;
};

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

const joinGraphs = (graph1, graph2) => ({
  ...graph1,
  ...fromEntries(
    entries(graph2).map(([from, toSet]) => [
      from,
      new Set([...(graph1[from] || []), ...toSet])
    ])
  )
});

const createGraph = edges =>
  edges
    .map(([from, to]) => ({
      [from]: new Set([to])
    }))
    .reduce(joinGraphs, {});

const calcGraphDist = (graph, start) => {
  // Dijkstra algorithm
  const rest = keys(graph);
  const dists = fromEntries(rest.map(v => [v, v === start ? 0 : Infinity]));
  while (rest.length > 0) {
    rest.sort((v1, v2) => {
      if (dists[v1] > dists[v2]) return 1;
      if (dists[v1] === dists[v2]) return 0;
      return -1;
    });
    const current = rest.shift();
    for (const v of graph[current]) {
      dists[v] = Math.min(dists[v], dists[current] + 1);
    }
  }
  return dists;
};

function* permutate(arr) {
  if (arr.length <= 1) {
    yield arr;
  } else {
    for (const i of Range(0, arr.length)) {
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

const update = (arr, index, value) => {
  const result = [...arr];
  result[index] = value;
  return result;
};

class Future {
  constructor() {
    const p = new Promise((resolve, reject) => {
      this.resolve = resolve;
      this.reject = reject;
    });
    this.then = p.then.bind(p);
    this.catch = p.catch.bind(p);
  }
}

module.exports = {
  Future,
  angle,
  calcGraphDist,
  compose,
  createGraph,
  dist,
  entries,
  floatEq,
  fromEntries,
  gcd,
  lcm,
  joinGraphs,
  keys,
  manhattan,
  max,
  min,
  permutate,
  readInput,
  sum,
  update,
  values
};
