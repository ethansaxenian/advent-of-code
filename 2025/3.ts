import { run } from "./util";

function joltage(battery: number[], size: number): number {
  if (size === 0) {
    return 0;
  }

  const candidates = battery.slice(0, battery.length - size + 1);
  const next = Math.max(...candidates);
  const remaining = battery.slice(battery.indexOf(next) + 1);

  return next * 10 ** (size - 1) + joltage(remaining, --size);
}

async function part1(input: string): Promise<number> {
  let j = 0;

  for (const line of input.split("\n")) {
    const battery = line.split("").map(Number);
    j += joltage(battery, 2);
  }

  return j;
}

async function part2(input: string): Promise<number> {
  let j = 0;
  for (const line of input.split("\n")) {
    let battery = line.split("").map(Number);
    j += joltage(battery, 12);
  }
  return j;
}

run(3, part1, part2);
