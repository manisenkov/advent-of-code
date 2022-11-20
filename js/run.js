const main = require("./lib/main");

const dayNumber = Number(process.argv[2]);
if (!dayNumber || dayNumber < 1 || dayNumber > 25) {
  throw new Error("Please specify day number (1-25)");
}
const dayNumberStr = (dayNumber < 10 ? "0" : "") + String(dayNumber);

main(dayNumberStr, `./inputs/day${dayNumberStr}`);
