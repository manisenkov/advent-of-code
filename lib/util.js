const readline = require("readline");

const { entries, fromEntries, keys, values } = Object;

const compose = (...fns) => (...args) => {
  let result = fns[0](...args);
  for (const fn of fns.slice(1)) {
    result = fn(result);
  }
  return result;
};

const forEach = (iter, fn) => {
  for (const el of iter) {
    fn(el);
  }
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

const calcDist = (graph, start, dest) => {
  // Dijkstra algorithm
  const rest = keys(graph);
  const dist = fromEntries(rest.map(v => [v, v === start ? 0 : Infinity]));
  while (rest.length > 0) {
    rest.sort((v1, v2) => {
      if (dist[v1] > dist[v2]) return 1;
      if (dist[v1] === dist[v2]) return 0;
      return -1;
    });
    const current = rest.shift();
    for (const v of graph[current]) {
      dist[v] = Math.min(dist[v], dist[current] + 1);
    }
  }
  return dist[dest];
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

const readInput = async () => {
  const rl = readline.createInterface({
    input: process.stdin
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
  calcDist,
  compose,
  createGraph,
  entries,
  forEach,
  fromEntries,
  joinGraphs,
  keys,
  permutate,
  range,
  readInput,
  sum,
  update,
  values
};
