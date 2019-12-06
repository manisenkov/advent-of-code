const {
  calcDist,
  createGraph,
  entries,
  fromEntries,
  joinGraphs,
  readInput,
  sum,
  values
} = require("./util");

const joinOrbitCounts = (orbitCounts1, orbitCounts2) => ({
  ...orbitCounts1,
  ...fromEntries(
    entries(orbitCounts2).map(([vertex, count]) => [
      vertex,
      (orbitCounts1[vertex] || 0) + count
    ])
  )
});

const calcOrbitCounts = (graph, start) =>
  [...(graph[start] || [])]
    .map(dest =>
      fromEntries([
        ...entries(calcOrbitCounts(graph, dest)).map(([vertex, orbitCount]) => [
          vertex,
          orbitCount + 1
        ]),
        [dest, 1]
      ])
    )
    .reduce(joinOrbitCounts, {});

const main = async () => {
  const input = await readInput();
  const edges = input.map(s => s.split(")"));

  // Part 1
  const graph = createGraph(edges);
  console.log("Part 1:", sum(values(calcOrbitCounts(graph, "COM"))));

  // Part 2
  // Revert graph nodes and join it with original one to create undirected version
  const undirectedGraph = joinGraphs(
    graph,
    createGraph(edges.map(([x, y]) => [y, x]))
  );
  console.log("Part 2:", calcDist(undirectedGraph, "YOU", "SAN") - 2); // Minus hops from [YOU] and to [SAN] nodes
};

main();
