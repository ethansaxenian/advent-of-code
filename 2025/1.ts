import { fetchInput } from "./util";

async function part1() {
  const input = await fetchInput(1);
  const lines = input
    .replace(/R/g, "")
    .replace(/L/g, "-")
    .split("\n")
    .map(Number);

  let d = 50;
  let ans = 0;
  for (const rot of lines) {
    d += rot;
    d %= 100;
    if (d == 0) {
      ans++;
    }
  }

  console.log(ans);
}

async function part2() {
  const input = await fetchInput(1);
  const lines = input
    .replace(/R/g, "")
    .replace(/L/g, "-")
    .split("\n")
    .map(Number);

  let d = 50;
  let ans = 0;
  for (const rot of lines) {
    for (let i = 0; i < Math.abs(rot); i++) {
      d += rot < 0 ? -1 : 1;
      d %= 100;
      if (d == 0) {
        ans++;
      }
    }
  }

  console.log(ans);
}

part1();
part2();
