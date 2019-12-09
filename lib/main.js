/* eslint import/no-dynamic-require: 0 */
/* eslint global-require: 0 */

const { readInput } = require("./util");

const main = async (dayNumberStr, inputStream) => {
  const dayModule = require(`./days/${dayNumberStr}`);
  const input = await readInput(inputStream);
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
  main(process.argv[2], process.stdin);
}
