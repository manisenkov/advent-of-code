const { runCode } = require("./intcode");
const { Future, permutate } = require("./util");

const calcThrust = async (opcodes, phaseSettings) => {
  const numPhases = phaseSettings.length;
  const inputFutures = Array(numPhases)
    .fill(0)
    .map(() => new Future());
  const outputs = [Array(numPhases).fill(0)];

  inputFutures[0].resolve(0);

  await Promise.all(
    phaseSettings.map(
      (phase, i) =>
        new Promise(resolve => {
          let isPhaseSent = false;
          const input = async () => {
            if (!isPhaseSent) {
              isPhaseSent = true;
              return phase;
            }
            const result = await inputFutures[i];
            inputFutures[i] = new Future();
            return result;
          };
          const output = async value => {
            outputs[(i + 1) % numPhases] = value;
            inputFutures[(i + 1) % numPhases].resolve(value);
          };

          runCode(opcodes, input, output).then(resolve);
        })
    )
  );

  return outputs[0];
};

const maximizeThrust = async (opcodes, possiblePhases) => {
  let maxThrust = -Infinity;
  let maxPhases = null;
  for (const phaseSettings of permutate(possiblePhases)) {
    const thrust = await calcThrust(opcodes, phaseSettings);
    if (thrust > maxThrust) {
      maxThrust = thrust;
      maxPhases = phaseSettings;
    }
  }
  return [maxThrust, maxPhases];
};

module.exports = {
  calcThrust,
  maximizeThrust
};
