const { createGraph, calcGraphDist } = require("../graphs");

const TOTAL_LEVELS = 500;

const posKey = ({ x, y }) => String(x * 10000 + y);

const portalKey = ({ name, type }) => `${type}:${name}`;

const parseInput = input => {
  const height = input.length;
  const width = Math.max(...input.map(s => s.length));
  const edges = [];
  const portals = [];

  // Parse maze
  for (let y = 2; y < height - 2; y++) {
    for (let x = 2; x < width - 2; x++) {
      const cell = input[y][x];
      if (cell === "#") {
        continue;
      }

      if (cell !== ".") {
        continue;
      }
      const pos = { x, y };
      const neighbors = [
        { x: x - 1, y },
        { x, y: y - 1 },
        { x: x + 1, y },
        { x, y: y + 1 }
      ];
      for (const neighbor of neighbors) {
        const c = input[neighbor.y][neighbor.x];
        if (c === ".") {
          edges.push([posKey(pos), posKey(neighbor)]);
        } else if (c >= "A" && c <= "Z") {
          let portal;
          if (neighbor.x === x - 1) {
            // west
            portal = input[neighbor.y][neighbor.x - 1] + c;
          } else if (neighbor.y === y - 1) {
            // north
            portal = input[neighbor.y - 1][neighbor.x] + c;
          } else if (neighbor.x === x + 1) {
            // east
            portal = c + input[neighbor.y][neighbor.x + 1];
          } else {
            // south
            portal = c + input[neighbor.y + 1][neighbor.x];
          }
          const isOuter =
            neighbor.y === 1 ||
            neighbor.y === height - 2 ||
            neighbor.x === 1 ||
            neighbor.x === width - 2;
          portals.push({
            name: portal,
            type: isOuter ? "outer" : "inner",
            pos
          });
        }
      }
    }
  }

  // Calculate distances between portals
  const graph = createGraph(edges);
  const portalDists = {};
  for (let i = 0; i < portals.length - 1; i++) {
    const portalFrom = portals[i];
    const portalFromKey = portalKey(portalFrom);
    const portalFromPosKey = posKey(portalFrom.pos);
    const dists = calcGraphDist(graph, portalFromPosKey);
    if (!portalDists[portalFromKey]) {
      portalDists[portalFromKey] = {};
    }
    for (let j = i + 1; j < portals.length; j++) {
      const portalTo = portals[j];
      const portalToKey = portalKey(portalTo);
      const portalToPosKey = posKey(portalTo.pos);
      const d = dists[portalToPosKey];
      if (d === Infinity) {
        continue;
      }
      if (!portalDists[portalToKey]) {
        portalDists[portalToKey] = {};
      }
      portalDists[portalFromKey][portalToKey] = d;
      portalDists[portalToKey][portalFromKey] = d;
    }
    console.log(
      ` ... Calculating distances between portals: ${i + 1} of ${
        portals.length
      }`
    );
  }

  return {
    portals: Array.from(new Set(portals.map(p => p.name))),
    distances: portalDists
  };
};

const part1 = ({ portals, distances }) => {
  const edges = [];

  // Adding portal routes
  for (const [from, dists] of Object.entries(distances)) {
    for (const [to, d] of Object.entries(dists)) {
      edges.push([from, to, d]);
    }
  }

  // Sewing portals
  for (const p of portals) {
    if (p === "AA" || p === "ZZ") {
      continue;
    }
    edges.push([`outer:${p}`, `inner:${p}`, 1]);
    edges.push([`inner:${p}`, `outer:${p}`, 1]);
  }

  const graph = createGraph(edges);
  const dists = calcGraphDist(graph, "outer:AA");

  return dists["outer:ZZ"];
};

const part2 = ({ portals, distances }) => {
  const edges = [];

  for (let level = 0; level < TOTAL_LEVELS; level++) {
    // Adding portal routes
    for (const [from, dists] of Object.entries(distances)) {
      for (const [to, d] of Object.entries(dists)) {
        edges.push([`${level}:${from}`, `${level}:${to}`, d]);
      }
    }

    // Sewing portals
    for (const p of portals) {
      if (p === "AA" || p === "ZZ") {
        continue;
      }
      if (level > 0) {
        edges.push([`${level}:outer:${p}`, `${level - 1}:inner:${p}`, 1]);
      }
      if (level < TOTAL_LEVELS - 1) {
        edges.push([`${level}:inner:${p}`, `${level + 1}:outer:${p}`, 1]);
      }
    }
  }

  const graph = createGraph(edges);
  const dists = calcGraphDist(graph, "0:outer:AA");

  return dists["0:outer:ZZ"];
};

exports.run = async input => {
  const maze = parseInput(input);
  return [part1(maze), part2(maze)];
};
