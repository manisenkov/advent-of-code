const { List, Range } = require("immutable");

const { compose } = require("../util");

const coordToStr = (x, y) => `${x},${y}`;

const manhattan = ([x, y]) => Math.abs(x) + Math.abs(y);

const traceWire = (wirePath, fn) => {
  let [x, y] = [0, 0];
  const apply = () => fn(x, y);
  const moves = {
    L: compose(() => x--, apply),
    U: compose(() => y--, apply),
    R: compose(() => x++, apply),
    D: compose(() => y++, apply)
  };
  wirePath.forEach(path => {
    const direction = path[0];
    const distance = Number(path.slice(1));
    Range(0, distance).forEach(moves[direction]);
  });
};

const getWireGrid = wirePath => {
  const grid = new Map();
  let steps = 0;
  traceWire(wirePath, (x, y) => {
    steps++;
    const coordStr = coordToStr(x, y);
    if (!grid.has(coordStr)) {
      grid.set(coordStr, steps);
    }
  });
  return grid;
};

const getIntersections = (wire1Grid, wire2Path) => {
  const result = [];
  let steps = 0;
  traceWire(wire2Path, (x, y) => {
    steps++;
    const coordStr = coordToStr(x, y);
    if (wire1Grid.has(coordStr)) {
      result.push({
        pos: [x, y],
        steps: steps + wire1Grid.get(coordStr)
      });
    }
  });
  return result;
};

exports.run = async ([input1, input2]) => {
  const wire1Path = List(input1.split(","));
  const wire2Path = List(input2.split(","));
  const wire1Grid = getWireGrid(wire1Path);
  const intersections = getIntersections(wire1Grid, wire2Path);

  intersections.sort(
    ({ pos: pos1 }, { pos: pos2 }) => manhattan(pos1) - manhattan(pos2)
  );
  const part1 = manhattan(intersections[0].pos);

  intersections.sort(({ steps: steps1 }, { steps: steps2 }) => steps1 - steps2);
  const part2 = intersections[0].steps;

  return [part1, part2];
};
