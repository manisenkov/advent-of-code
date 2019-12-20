const { createGraph, calcGraphDist } = require("../graphs");

const { parseInput } = require("./20-ext");

exports.run = async input => {
  const maze = parseInput(input);
  const graph = createGraph(maze.edges);
  const distFromAA = calcGraphDist(graph, maze.portals.AA.outer);
  const mazeMult = parseInput(input, 20);
  const graphMult = createGraph(mazeMult.edges);
  const distFromAAMult = calcGraphDist(
    graphMult,
    `0:${mazeMult.portals.AA.outer}`
  );
  return [
    distFromAA[maze.portals.ZZ.outer],
    distFromAAMult[`0:${mazeMult.portals.ZZ.outer}`]
  ];
};
