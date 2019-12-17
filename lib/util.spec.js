const { Future, permutate, update } = require("./util");

describe("update", () => {
  it("should change value by index provided", () => {
    const arr = [1, 2, 3, 4, 5];
    const nextArr = update(arr, 2, 42);
    expect(nextArr).toEqual([1, 2, 42, 4, 5]);
    expect(arr).toEqual([1, 2, 3, 4, 5]);
  });
});

describe("permutate", () => {
  it("should generate all permutations", () => {
    const permutations = Array.from(permutate(["a", "b", "c"]));
    expect(permutations).toEqual([
      ["a", "b", "c"],
      ["a", "c", "b"],
      ["b", "a", "c"],
      ["b", "c", "a"],
      ["c", "a", "b"],
      ["c", "b", "a"]
    ]);
  });
});

describe("Future", () => {
  it("should eventually resolves with value", async () => {
    const f = new Future();
    setTimeout(() => f.resolve(42), 0);
    await expect(f).resolves.toEqual(42);
  });

  it("should immediately resolves with value", async () => {
    const f = new Future();
    f.resolve(42);
    await expect(f).resolves.toEqual(42);
  });
});
