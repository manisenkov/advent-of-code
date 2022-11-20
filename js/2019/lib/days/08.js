const { Range } = require("immutable");

const { min } = require("../util");

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
  const lines = [];
  for (let i = 0; i < image.length; i += IMAGE_WIDTH) {
    const line = image
      .slice(i, i + IMAGE_WIDTH)
      .map(c => ({ 0: " ", 1: "█" }[c]))
      .join("");
    lines.push(line);
  }
  return lines.join("\n");
};

exports.run = async ([input]) => {
  const layers = getLayers(input);

  const digitsCount = layers.map(getDigitsCount);
  const minZeros = min(digitsCount, d => d[0]);
  const part1 = minZeros[1] * minZeros[2];

  const image = Array.from(Range(0, IMAGE_SIZE)).map((_, index) =>
    getPixel(layers, index)
  );
  const part2 = printImage(image);

  return [part1, part2];
};
