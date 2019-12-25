const joinGraphs = (graph1, graph2) => {
  const result = {};
  for (const from of Object.keys(graph1)) {
    result[from] = new Map(graph1[from]);
  }
  for (const from of Object.keys(graph2)) {
    for (const [val, dist] of graph2[from].entries()) {
      if (from in result) {
        result[from].set(val, dist);
      } else {
        result[from] = new Map([[val, dist]]);
      }
    }
  }
  return result;
};

const createGraph = edges => {
  const result = {};
  for (const [from, to, dist] of edges) {
    result[from] = (result[from] || new Map()).set(
      to,
      typeof dist === "undefined" ? 1 : dist
    );
    result[to] = result[to] || new Map();
  }
  return result;
};

const calcGraphDist = (graph, start) => {
  // Dijkstra algorithm
  const rest = Object.keys(graph);
  const dists = Object.fromEntries(
    rest.map(v => [v, v === start ? 0 : Infinity])
  );
  while (rest.length > 0) {
    rest.sort((v1, v2) => {
      if (dists[v1] > dists[v2]) return 1;
      if (dists[v1] === dists[v2]) return 0;
      return -1;
    });
    const current = rest.shift();
    for (const [v, dist] of graph[current]) {
      dists[v] = Math.min(dists[v], dists[current] + dist);
    }
  }
  return dists;
};

module.exports = {
  joinGraphs,
  createGraph,
  calcGraphDist
};
