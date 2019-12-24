const { sum } = require("../util");

const parseInput = input =>
  input.flatMap(s => Array.from(s).map(c => (c === "#" ? 1 : 0)));

const get = (row, col, layout) => layout[row * 5 + col];

const set = (row, col, layout, value) => {
  layout[row * 5 + col] = value;
};

const part1 = layout => {
  let current = layout;

  const neighbors = (col, row) =>
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
        const n = neighbors(col, row);
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

exports.run = async input => {
  const layout = parseInput(input);
  return [part1(layout), 0];
};
