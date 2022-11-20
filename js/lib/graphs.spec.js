const { createGraph, joinGraphs } = require("./graphs");

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
      a: new Map([
        ["b", 1],
        ["c", 1]
      ]),
      b: new Map([["d", 1]]),
      c: new Map(),
      d: new Map(),
      e: new Map([
        ["a", 1],
        ["c", 1],
        ["f", 1]
      ]),
      f: new Map()
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
      a: new Map([
        ["b", 1],
        ["c", 1],
        ["d", 1]
      ]),
      b: new Map([
        ["a", 1],
        ["d", 1]
      ]),
      c: new Map([["f", 1]]),
      d: new Map(),
      e: new Map([
        ["a", 1],
        ["c", 1],
        ["f", 1]
      ]),
      f: new Map([
        ["d", 1],
        ["e", 1]
      ])
    });
  });
});
