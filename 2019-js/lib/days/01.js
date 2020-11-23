const { sum } = require("../util");

const fuelReq = mass => Math.floor(mass / 3) - 2;

const totalFuelReq = mass => {
  const m = fuelReq(mass);
  return m <= 0 ? 0 : m + totalFuelReq(m);
};

exports.run = async input => {
  const moduleMasses = input.map(Number);
  return [sum(moduleMasses.map(fuelReq)), sum(moduleMasses.map(totalFuelReq))];
};
