const { readInput, sum } = require("./util");

const fuelReq = mass => Math.floor(mass / 3) - 2;

const totalFuelReq = mass => {
  const m = fuelReq(mass);
  return m <= 0 ? 0 : m + totalFuelReq(m);
};

const main = async () => {
  const moduleMasses = (await readInput()).map(Number);

  // Part 1
  console.log("Part 1:", sum(moduleMasses.map(fuelReq)));

  // Part 2
  console.log("Part 2:", sum(moduleMasses.map(totalFuelReq)));
};

main();
