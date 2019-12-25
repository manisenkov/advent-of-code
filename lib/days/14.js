const { List } = require("immutable");

const MAX_ORE = 1000000000000;

const parseMaterial = input => {
  const [amountStr, chem] = input.split(" ").map(s => s.trim());
  return { amount: Number(amountStr), chem };
};

const parseRecipe = input => {
  const [compStr, resultMaterialStr] = input.split("=>").map(s => s.trim());
  const components = Object.fromEntries(
    compStr
      .split(",")
      .map(s => s.trim())
      .map(parseMaterial)
      .map(({ chem, amount }) => [chem, amount])
  );
  const result = parseMaterial(resultMaterialStr);
  return {
    chem: result.chem,
    amount: result.amount,
    components
  };
};

const calcDistsFromFuel = recipes => {
  const dists = {};
  const queue = [{ chem: "FUEL", dist: 0 }];
  while (queue.length > 0) {
    const { chem, dist } = queue.shift();
    dists[chem] = dist;
    if (chem === "ORE") {
      continue;
    }
    for (const compChem of Object.keys(recipes[chem].components)) {
      queue.push({ chem: compChem, dist: dist + 1 });
    }
  }
  return dists;
};

const calcOreAmount = (recipes, dists, fuelAmount) => {
  let state = List([{ chem: "FUEL", amount: fuelAmount }]);
  let totalOre = 0;

  while (state.size > 0) {
    const [{ chem, amount }, stateRest] = [state.first(), state.rest()];

    if (chem === "ORE") {
      totalOre += amount;
      state = stateRest;
      continue;
    }

    const { components, amount: batchAmount } = recipes[chem];
    const batchCount = Math.ceil(amount / batchAmount);
    const stateObj = Object.fromEntries(stateRest.map(s => [s.chem, s.amount]));
    const toAdd = Object.fromEntries(
      Object.entries(components).map(([compChem, compAmount]) => [
        compChem,
        (stateObj[compChem] || 0) + batchCount * compAmount
      ])
    );

    state = List(
      Object.entries({ ...stateObj, ...toAdd }).map(
        ([compChem, compAmount]) => ({
          chem: compChem,
          amount: compAmount
        })
      )
    ).sort((x, y) => dists[x.chem] - dists[y.chem]);
  }

  return totalOre;
};

const calcMaxFuel = (recipes, dists) => {
  let [min, max] = [1, 10000000000];
  while (max - min > 1) {
    const fuel = ((min + max) / 2) | 0;
    const ore = calcOreAmount(recipes, dists, fuel);
    if (ore < MAX_ORE) {
      min = fuel;
    } else if (ore >= MAX_ORE) {
      max = fuel;
    }
  }
  return min;
};

exports.run = async input => {
  const recipes = Object.fromEntries(
    input.map(parseRecipe).map(r => [r.chem, r])
  );
  const dists = calcDistsFromFuel(recipes);
  return [calcOreAmount(recipes, dists, 1), calcMaxFuel(recipes, dists)];
};
