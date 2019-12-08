const { min, range, readInput } = require("./util");

const IMAGE_WIDTH = 25;
const IMAGE_HEIGHT = 6;
const IMAGE_SIZE = IMAGE_WIDTH * IMAGE_HEIGHT;
const TRANSPARENT = 2;

const getLayers = input => {
  const layers = [];
  for (let i = 0; i < input.length; i += IMAGE_SIZE) {
    layers.push(Array.from(input.slice(i, i + IMAGE_SIZE)).map(Number));
  }
  return layers;
};

const getDigitsCount = layer =>
  layer.reduce(
    (result, digit) => ({ ...result, [digit]: (result[digit] || 0) + 1 }),
    {}
  );

const getPixel = (layers, index) => {
  for (const layer of layers) {
    const pixel = layer[index];
    if (pixel !== TRANSPARENT) {
      return pixel;
    }
  }
  return TRANSPARENT;
};

const printImage = image => {
  for (let i = 0; i < image.length; i += IMAGE_WIDTH) {
    const line = image
      .slice(i, i + IMAGE_WIDTH)
      .map(c => ({ 0: " ", 1: "█" }[c]))
      .join("");
    console.log(line);
  }
};

const main = async () => {
  const [input] = await readInput();
  const layers = getLayers(input);

  // Part 1
  const digitsCount = layers.map(getDigitsCount);
  const minZeros = min(digitsCount, d => d[0]);
  console.log("Part 1:", minZeros[1] * minZeros[2]);

  // Part 2
  const image = Array.from(range(0, IMAGE_SIZE)).map((_, index) =>
    getPixel(layers, index)
  );
  console.log("Part 2:");
  printImage(image);
};

main();

// 11110 10010 00110 10010 10000
// 10000 10010 00010 10010 10000
// 11100 11110 00010 10010 10000
// 10000 10010 00010 10010 10000
// 10000 10010 10010 10010 10000
// 10000 10010 01100 01100 11110
