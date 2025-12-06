import { run } from "./util";

async function part1(input: string): Promise<number> {
  const parts = input.split("\n\n");
  const ranges = parts[0]
    .split("\n")
    .map((line) => line.split("-").map(Number));
  const ids = parts[1].split("\n").map(Number);
  let n = 0;
  for (const id of ids) {
    for (const range of ranges) {
      if (id >= range[0] && id <= range[1]) {
        n++;
        break;
      }
    }
  }
  return n;
}

async function part2(input: string): Promise<number> {
  const ranges = input
    .split("\n\n")[0]
    .split("\n")
    .map((line) => line.split("-").map(Number))
    .sort((a, b) => a[0] - b[0]);

  let n = 0;
  let curr_max = 0;

  for (let [l, r] of ranges) {
    if (r >= curr_max) {
      n += r - Math.max(l, curr_max) + 1;
      curr_max = r + 1;
    }
  }
  return n;
}

run(5, part1, part2);
