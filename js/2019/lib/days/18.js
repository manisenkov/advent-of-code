const { min } = require("../util");

const posKey = ({ x, y }) => x * 10000 + y;

const findKeys = (maze, start, givenKeys) => {
  const nextCells = [start];
  const dists = { [posKey(start)]: 0 };
  const foundKeys = {};
  while (nextCells.length > 0) {
    const pos = nextCells.shift();
    const neighbors = [
      { x: pos.x - 1, y: pos.y },
      { x: pos.x + 1, y: pos.y },
      { x: pos.x, y: pos.y - 1 },
      { x: pos.x, y: pos.y + 1 }
    ];
    const cellDist = dists[posKey(pos)] + 1;
    for (const neighbor of neighbors) {
      const nCell = (maze[neighbor.y] || [])[neighbor.x];
      const cellKey = posKey(neighbor);
      if (
        typeof nCell === "undefined" ||
        nCell === "#" ||
        cellKey in dists ||
        (nCell >= "A" &&
          nCell <= "Z" &&
          givenKeys.indexOf(nCell.toLowerCase()) === -1)
      ) {
        continue;
      }
      dists[cellKey] = cellDist;
      if (nCell >= "a" && nCell <= "z" && givenKeys.indexOf(nCell) === -1) {
        foundKeys[nCell] = {
          dist: cellDist,
          pos: neighbor,
          startKey: posKey(start)
        };
      } else {
        nextCells.push(neighbor);
      }
    }
  }
  return foundKeys;
};

const findKeysMultiple = (maze, starts, givenKeys) =>
  Object.fromEntries(
    starts
      .map(start => Object.entries(findKeys(maze, start, givenKeys)))
      .flatMap(arr => arr)
  );

const minDist = (maze, starts, givenKeys, cache) => {
  givenKeys.sort();
  const cacheKey = `${starts.map(posKey).join(",")}:${givenKeys.join(",")}`;
  if (typeof cache[cacheKey] !== "undefined") {
    return cache[cacheKey];
  }
  const keys = findKeysMultiple(maze, starts, givenKeys);
  if (Object.entries(keys).length === 0) {
    cache[cacheKey] = 0;
    return 0;
  }
  const results = Object.entries(keys).map(([key, { dist, pos, startKey }]) => {
    const nextStarts = starts.map(s => (posKey(s) === startKey ? pos : s));
    return dist + minDist(maze, nextStarts, [...givenKeys, key], cache);
  });
  const result = min(results);
  cache[cacheKey] = result;
  return result;
};

const findStarts = maze => {
  const starts = [];
  for (let y = 0; y < maze.length; y++) {
    for (let x = 0; x < maze[y].length; x++) {
      if (maze[y][x] === "@") {
        starts.push({ x, y });
      }
    }
  }
  return starts;
};

exports.run = async input => {
  const maze = input.map(row => Array.from(row));
  const start = findStarts(maze)[0];
  const part1 = minDist(maze, [start], [], {});

  // Update maze for part 2
  maze[start.y - 1][start.x - 1] = "@";
  maze[start.y - 1][start.x] = "#";
  maze[start.y - 1][start.x + 1] = "@";
  maze[start.y][start.x - 1] = "#";
  maze[start.y][start.x] = "#";
  maze[start.y][start.x + 1] = "#";
  maze[start.y + 1][start.x - 1] = "@";
  maze[start.y + 1][start.x] = "#";
  maze[start.y + 1][start.x + 1] = "@";

  const part2 = minDist(
    maze,
    [
      { x: start.x - 1, y: start.y - 1 },
      { x: start.x - 1, y: start.y + 1 },
      { x: start.x + 1, y: start.y - 1 },
      { x: start.x + 1, y: start.y + 1 }
    ],
    [],
    {}
  );

  return [part1, part2];
};
