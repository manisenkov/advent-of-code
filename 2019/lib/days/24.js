const { range, sum } = require("../util");

const parseInput = input =>
  input.flatMap(s => Array.from(s).map(c => (c === "#" ? 1 : 0)));

const get = (row, col, layout) => (layout ? layout[row * 5 + col] : 0);

const set = (row, col, layout, value) => {
  layout[row * 5 + col] = value;
};

const part1 = layout => {
  let current = layout;

  const neighbors = (row, col) =>
    sum([
      col > 0 ? get(row, col - 1, current) : 0,
      col < 4 ? get(row, col + 1, current) : 0,
      row > 0 ? get(row - 1, col, current) : 0,
      row < 4 ? get(row + 1, col, current) : 0
    ]);

  const trans = () => {
    const result = Array(25).fill(0);
    for (let row = 0; row < 5; row++) {
      for (let col = 0; col < 5; col++) {
        const c = get(row, col, current);
        const n = neighbors(row, col);
        if ((c === 1 && n === 1) || (c === 0 && (n === 1 || n === 2))) {
          set(row, col, result, 1);
        }
      }
    }
    return result;
  };

  const bioDivs = new Set();
  while (true) {
    const d = sum(current.map((c, i) => c * 2 ** i));
    if (bioDivs.has(d)) {
      return d;
    }
    bioDivs.add(d);
    current = trans(current);
  }
};

const part2 = layout => {
  let layouts = { 0: layout };
  let maxLevel = 0;

  const neighbors = (row, col, level) =>
    sum([
      col > 0 ? get(row, col - 1, layouts[level]) : 0,
      col < 4 ? get(row, col + 1, layouts[level]) : 0,
      row > 0 ? get(row - 1, col, layouts[level]) : 0,
      row < 4 ? get(row + 1, col, layouts[level]) : 0,
      // Outer neighbors
      col === 0 ? get(2, 1, layouts[level - 1]) : 0,
      col === 4 ? get(2, 3, layouts[level - 1]) : 0,
      row === 0 ? get(1, 2, layouts[level - 1]) : 0,
      row === 4 ? get(3, 2, layouts[level - 1]) : 0,
      // Inner neighbors
      col === 2 && row === 1
        ? sum([
            get(0, 0, layouts[level + 1]),
            get(0, 1, layouts[level + 1]),
            get(0, 2, layouts[level + 1]),
            get(0, 3, layouts[level + 1]),
            get(0, 4, layouts[level + 1])
          ])
        : 0,
      col === 3 && row === 2
        ? sum([
            get(0, 4, layouts[level + 1]),
            get(1, 4, layouts[level + 1]),
            get(2, 4, layouts[level + 1]),
            get(3, 4, layouts[level + 1]),
            get(4, 4, layouts[level + 1])
          ])
        : 0,
      col === 2 && row === 3
        ? sum([
            get(4, 0, layouts[level + 1]),
            get(4, 1, layouts[level + 1]),
            get(4, 2, layouts[level + 1]),
            get(4, 3, layouts[level + 1]),
            get(4, 4, layouts[level + 1])
          ])
        : 0,
      col === 1 && row === 2
        ? sum([
            get(0, 0, layouts[level + 1]),
            get(1, 0, layouts[level + 1]),
            get(2, 0, layouts[level + 1]),
            get(3, 0, layouts[level + 1]),
            get(4, 0, layouts[level + 1])
          ])
        : 0
    ]);

  const trans = () => {
    const result = Object.fromEntries(
      Array.from(range(-maxLevel, maxLevel + 1)).map(i => [
        i,
        Array(25).fill(0)
      ])
    );
    for (let i = -maxLevel; i <= maxLevel; i++) {
      for (let row = 0; row < 5; row++) {
        for (let col = 0; col < 5; col++) {
          if (row === 2 && col === 2) {
            continue;
          }
          const c = get(row, col, layouts[i]);
          const n = neighbors(row, col, i);
          if ((c === 1 && n === 1) || (c === 0 && (n === 1 || n === 2))) {
            set(row, col, result[i], 1);
          }
        }
      }
    }
    layouts = result;
  };

  for (const i of range(0, 200)) {
    if (i % 2 === 0) {
      maxLevel++;
      layouts[maxLevel] = Array(25).fill(0);
      layouts[-maxLevel] = Array(25).fill(0);
    }
    trans();
  }

  return sum(Object.values(layouts).flatMap(r => r));
};

exports.run = async input => {
  const layout = parseInput(input);
  return [part1(layout), part2(layout)];
};
