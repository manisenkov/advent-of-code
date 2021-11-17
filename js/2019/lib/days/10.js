const { List, Set, Range } = require("immutable");

const { angle, dist, floatEq, gcd } = require("../util");

class Asteroid {
  constructor(y, x) {
    this.y = y;
    this.x = x;
  }

  hashCode() {
    return (this.x << 5) & this.y;
  }

  equals(other) {
    return this.x === other.x && this.y === other.y;
  }

  toString() {
    return `Asteroid(${this.y}, ${this.x})`;
  }
}

const getAsteroids = spacemap =>
  Set(
    spacemap.flatMap((line, x) =>
      line
        .map((c, y) => (c !== "." ? new Asteroid(y, x) : null))
        .filter(v => v != null)
    )
  );

const calcDirections = (w, h) => {
  const directions = [];
  for (const dy of Range(-w, w)) {
    for (const dx of Range(-h, h)) {
      if ((dy === 0 && dx === 0) || gcd(Math.abs(dy), Math.abs(dx)) > 1) {
        continue;
      }
      directions.push([dy, dx]);
    }
  }
  return directions;
};

const calcVisibility = (current, asteroids, [w, h], directions) => {
  let visibility = 0;
  for (const [dy, dx] of directions) {
    for (const i of Range(1)) {
      const pos = new Asteroid(current.y + i * dy, current.x + i * dx);
      if (pos.y < 0 || pos.x < 0 || pos.y >= h || pos.x >= w) {
        break;
      }
      if (asteroids.has(pos)) {
        visibility++;
        break;
      }
    }
  }
  return visibility;
};

const getEvaporizeOrder = (station, asteroids) => {
  const laserAngle = asteroid =>
    angle(asteroid.y - station.y, -asteroid.x + station.x);
  const laserDist = asteroid =>
    dist(asteroid.y - station.y, -asteroid.x + station.x);

  const sorted = asteroids
    .filter(a => !a.equals(station))
    .toList()
    .sort((asteroid1, asteroid2) => {
      const angle1 = laserAngle(asteroid1);
      const angle2 = laserAngle(asteroid2);
      if (floatEq(angle1 - angle2)) {
        return laserDist(asteroid1) - dist(asteroid2);
      }
      return angle1 - angle2;
    });
  let result = List([sorted.first()]);
  let rest = List(sorted.rest());
  let nextTurn = List();
  let current;
  while (rest.size + nextTurn.size > 0) {
    if (rest.size === 0) {
      rest = nextTurn;
      nextTurn = List();
    }
    [current, rest] = [rest.first(), rest.rest()];
    if (floatEq(laserAngle(current), laserAngle(result.get(-1)))) {
      nextTurn = nextTurn.push(current);
    } else {
      result = result.push(current);
    }
  }
  return result;
};

exports.run = async input => {
  const mapSize = [input[0].length, input.length];
  const spacemap = List(input.map(List));
  const asteroids = getAsteroids(spacemap);
  const directions = calcDirections(...mapSize);
  const visibility = asteroids
    .map(current => ({
      asteroid: current,
      visibility: calcVisibility(current, asteroids, mapSize, directions)
    }))
    .toList();
  const { asteroid: maxAsteroid, visibility: maxVisibility } = visibility.max(
    ({ visibility: a }, { visibility: b }) => a - b
  );
  const { y: evY, x: evX } = getEvaporizeOrder(maxAsteroid, asteroids).get(199);

  return [maxVisibility, evY * 100 + evX];
};
