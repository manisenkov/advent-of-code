const { runCode } = require("./intcode");

describe("IntCode basic features", () => {
  it("should output input value", async () => {
    const testProgram = JSON.parse(
      "[3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9]"
    );
    const output = [];
    await runCode(
      testProgram,
      async () => 0,
      async val => output.push(val)
    );
    await runCode(
      testProgram,
      async () => 1,
      async val => output.push(val)
    );
    expect(output).toEqual([0, 1]);
  });

  it("should output input value (in relative mode)", async () => {
    const testProgram = JSON.parse(
      "[109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99]"
    );
    const output = [];
    await runCode(
      testProgram,
      async () => 0,
      async val => output.push(val)
    );
    expect(output).toEqual(testProgram);
  });

  it("should output 16-digit number", async () => {
    const testProgram = JSON.parse("[1102,34915192,34915192,7,4,7,99,0]");
    const output = [];
    await runCode(
      testProgram,
      async () => 0,
      async val => output.push(val)
    );
    expect(output).toEqual([1219070632396864]);
  });

  it("should output 1125899906842624", async () => {
    const testProgram = JSON.parse("[104,1125899906842624,99]");
    const output = [];
    await runCode(
      testProgram,
      async () => 0,
      async val => output.push(val)
    );
    expect(output).toEqual([1125899906842624]);
  });
});
