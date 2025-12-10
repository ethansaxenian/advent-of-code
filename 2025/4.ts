import { run } from "@util";

const directions = [
  [-1, 0],
  [-1, 1],
  [0, 1],
  [1, 1],
  [1, 0],
  [1, -1],
  [0, -1],
  [-1, -1],
];

function neighbors(roll: string, g: Map<string, string>): number {
  const [r, c] = roll.split(",").map(Number);
  let n = 0;
  directions.forEach(([dr, dc]) => {
    if (g.get([r + dr, c + dc].join(",")) === "@") {
      n++;
    }
  });
  return n;
}

async function part1(input: string): Promise<number> {
  const g = new Map<string, string>();
  const rolls = new Set<string>();
  input.split("\n").forEach((row, r) => {
    row.split("").forEach((char, c) => {
      if (char === "@") {
        g.set([r, c].join(","), char);
        rolls.add([r, c].join(","));
      }
    });
  });

  let n = 0;
  for (const roll of rolls.values()) {
    if (neighbors(roll, g) < 4) {
      n++;
    }
  }
  return n;
}

async function part2(input: string): Promise<number> {
  let g = new Map<string, string>();
  const rolls = new Set<string>();
  input.split("\n").forEach((row, r) => {
    row.split("").forEach((char, c) => {
      if (char === "@") {
        g.set([r, c].join(","), char);
        rolls.add([r, c].join(","));
      }
    });
  });

  let total = 0;
  while (true) {
    const gNext = new Map<string, string>();
    let n = 0;
    for (const roll of rolls.values()) {
      if (neighbors(roll, g) < 4) {
        n++;
        gNext.set(roll, ".");
        rolls.delete(roll);
      } else {
        gNext.set(roll, "@");
      }
    }
    g = gNext;
    if (n === 0) {
      break;
    }
    total += n;
  }
  return total;
}

run(4, 2025, part1, part2);
