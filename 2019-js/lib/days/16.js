const { sum } = require("../util");

const fft = (signal, pat) => {
  let result = 0;
  // +1
  for (let i = pat; i < signal.length; i += (pat + 1) * 4) {
    for (let j = i; j < i + pat + 1 && j < signal.length; j++) {
      result += signal[j];
    }
  }
  // -1
  for (let i = pat + (pat + 1) * 2; i < signal.length; i += (pat + 1) * 4) {
    for (let j = i; j < i + pat + 1 && j < signal.length; j++) {
      result -= signal[j];
    }
  }
  return Math.abs(result) % 10;
};

const phase100 = signal => {
  let result = signal;
  for (let i = 0; i < 100; i++) {
    const next = [];
    for (let j = 0; j < result.length; j++) {
      next.push(fft(result, j));
    }
    result = next;
  }
  return result;
};

const phase100Part2 = signal => {
  // bottom right part of pattern matrix is just triangle of "1"s
  // so we can just summarize elements there
  let result = signal;
  for (let i = 0; i < 100; i++) {
    const next = [];
    let current = sum(result);
    for (let j = 0; j < result.length; j++) {
      next.push(Math.abs(current) % 10);
      current -= result[j];
    }
    result = next;
  }
  return result;
};

exports.run = async ([input]) => {
  const signal = Array.from(input).map(Number);
  const realSignal = Array(signal.length * 10000);
  for (let i = 0; i < signal.length * 10000; i++) {
    realSignal[i] = signal[i % signal.length];
  }
  const offset = Number(signal.slice(0, 7).join(""));

  return [
    phase100(signal)
      .slice(0, 8)
      .join(""),
    phase100Part2(realSignal.slice(offset))
      .slice(0, 8)
      .join("")
  ];
};
