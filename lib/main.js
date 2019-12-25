/* eslint import/no-dynamic-require: 0 */
/* eslint global-require: 0 */

const { createReadStream } = require("fs");

const { readInput } = require("./util");

const main = async (dayNumberStr, inputFileName) => {
  const dayModule = require(`./days/${dayNumberStr}`);
  const input = await readInput(createReadStream(inputFileName));
  const result = await dayModule.run(input);
  result.forEach((value, i) => {
    console.log(
      `Part ${i + 1}:`,
      typeof value === "string" ? `\n${value}` : value
    );
  });
};

module.exports = main;

if (require.main === module) {
  main(...process.argv.slice(2));
}
