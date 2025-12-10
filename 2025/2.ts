import { run } from "@util";

async function part1(input: string): Promise<number> {
  const pattern = /^(.+)\1$/;

  let s = 0;

  for (const line of input.split(",")) {
    const [a, b] = line.split("-").map(Number);

    for (let id = a; id <= b; id++) {
      if (id.toString().match(pattern)) {
        s += id;
      }
    }
  }

  return s;
}

async function part2(input: string): Promise<number> {
  const pattern = /^(.+)\1+$/;

  let s = 0;

  for (const line of input.split(",")) {
    const [a, b] = line.split("-").map(Number);

    for (let id = a; id <= b; id++) {
      if (id.toString().match(pattern)) {
        s += id;
      }
    }
  }

  return s;
}

run(2, 2025, part1, part2);
