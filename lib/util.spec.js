const { createGraph, joinGraphs, update } = require("./util");

describe("update", () => {
  it("should change value by index provided", () => {
    const arr = [1, 2, 3, 4, 5];
    const nextArr = update(arr, 2, 42);
    expect(nextArr).toEqual([1, 2, 42, 4, 5]);
    expect(arr).toEqual([1, 2, 3, 4, 5]);
  });
});

describe("createGraph", () => {
  it("should create directed graph", () => {
    const graph = createGraph([
      ["a", "b"],
      ["a", "c"],
      ["b", "d"],
      ["e", "a"],
      ["e", "c"],
      ["e", "f"]
    ]);
    expect(graph).toEqual({
      a: new Set(["b", "c"]),
      b: new Set(["d"]),
      e: new Set(["a", "c", "f"])
    });
  });
});

describe("joinGraphs", () => {
  it("should join two directed graphs", () => {
    const graph1 = createGraph([
      ["a", "b"],
      ["a", "c"],
      ["b", "d"],
      ["e", "a"],
      ["e", "c"],
      ["e", "f"]
    ]);
    const graph2 = createGraph([
      ["a", "d"],
      ["b", "a"],
      ["c", "f"],
      ["e", "c"],
      ["f", "d"],
      ["f", "e"]
    ]);
    expect(joinGraphs(graph1, graph2)).toEqual({
      a: new Set(["b", "c", "d"]),
      b: new Set(["a", "d"]),
      c: new Set(["f"]),
      e: new Set(["a", "c", "f"]),
      f: new Set(["d", "e"])
    });
  });
});