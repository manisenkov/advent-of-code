const { lcm, sum } = require("../util");

const inputRegex = /<x=(-?\d+), ?y=(-?\d+), ?z=(-?\d+)>/i;

const NUM_STEPS = 1000;

const parseCoords = input => {
  const m = input.match(inputRegex);
  return { x: Number(m[1]), y: Number(m[2]), z: Number(m[3]) };
};

const applyGravity = (moon1, moon2, velocity1, velocity2) => {
  const newVelocity1 = { ...velocity1 };
  const newVelocity2 = { ...velocity2 };
  if (moon1.x < moon2.x) {
    newVelocity1.vx++;
    newVelocity2.vx--;
  }
  if (moon1.x > moon2.x) {
    newVelocity1.vx--;
    newVelocity2.vx++;
  }
  if (moon1.y < moon2.y) {
    newVelocity1.vy++;
    newVelocity2.vy--;
  }
  if (moon1.y > moon2.y) {
    newVelocity1.vy--;
    newVelocity2.vy++;
  }
  if (moon1.z < moon2.z) {
    newVelocity1.vz++;
    newVelocity2.vz--;
  }
  if (moon1.z > moon2.z) {
    newVelocity1.vz--;
    newVelocity2.vz++;
  }
  return [newVelocity1, newVelocity2];
};

const applyVelocity = (moon, velocity) => ({
  x: moon.x + velocity.vx,
  y: moon.y + velocity.vy,
  z: moon.z + velocity.vz
});

const calcPot = moon => Math.abs(moon.x) + Math.abs(moon.y) + Math.abs(moon.z);
const calcKin = velocity =>
  Math.abs(velocity.vx) + Math.abs(velocity.vy) + Math.abs(velocity.vz);

exports.run = async input => {
  const moons = input.map(parseCoords);
  const startMoons = moons.map(m => ({ ...m }));
  const velocities = moons.map(() => ({ vx: 0, vy: 0, vz: 0 }));
  let totalEnergy;

  const stepToRepeat = { x: 0, y: 0, z: 0 };
  let step = 0;

  while (
    step <= NUM_STEPS ||
    !stepToRepeat.x ||
    !stepToRepeat.y ||
    !stepToRepeat.z
  ) {
    if (step === NUM_STEPS) {
      totalEnergy = sum(
        moons
          .map((moon, i) => [moon, velocities[i]])
          .map(([moon, velocity]) => calcPot(moon) * calcKin(velocity))
      );
    }
    for (let i = 0; i < moons.length - 1; i++) {
      for (let j = i + 1; j < moons.length; j++) {
        const moon1 = moons[i];
        const moon2 = moons[j];
        const velocity1 = velocities[i];
        const velocity2 = velocities[j];
        [velocities[i], velocities[j]] = applyGravity(
          moon1,
          moon2,
          velocity1,
          velocity2
        );
      }
    }
    for (let i = 0; i < moons.length; i++) {
      moons[i] = applyVelocity(moons[i], velocities[i]);
    }
    step++;

    // Check repeating cycles
    if (
      !stepToRepeat.x &&
      velocities.every(v => v.vx === 0) &&
      moons.every((m, i) => m.x === startMoons[i].x)
    ) {
      stepToRepeat.x = step;
    }
    if (
      !stepToRepeat.y &&
      velocities.every(v => v.vy === 0) &&
      moons.every((m, i) => m.y === startMoons[i].y)
    ) {
      stepToRepeat.y = step;
    }
    if (
      !stepToRepeat.z &&
      velocities.every(v => v.vz === 0) &&
      moons.every((m, i) => m.z === startMoons[i].z)
    ) {
      stepToRepeat.z = step;
    }
  }

  const totalStepsToRepeat = lcm(
    lcm(stepToRepeat.x, stepToRepeat.y),
    stepToRepeat.z
  );

  return [totalEnergy, totalStepsToRepeat];
};
