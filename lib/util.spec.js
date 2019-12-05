const { update } = require("./util");

describe("update", () => {
  it("should change value by index provided", () => {
    const arr = [1, 2, 3, 4, 5];
    const nextArr = update(arr, 2, 42);
    expect(nextArr).toEqual([1, 2, 42, 4, 5]);
    expect(arr).toEqual([1, 2, 3, 4, 5]);
  });
});
