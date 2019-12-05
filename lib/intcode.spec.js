const { run } = require("./intcode");

describe("IntCode test program", () => {
  const testProgram = [3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9];

  it("should output 0", () => {
    const output = [];
    run(
      testProgram,
      () => 0,
      val => output.push(val)
    );
    expect(output).toEqual([0]);
  });

  it("should output 1", () => {
    const output = [];
    run(
      testProgram,
      () => 0,
      val => output.push(val)
    );
    expect(output).toEqual([0]);
  });
});
