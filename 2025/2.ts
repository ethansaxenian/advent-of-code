import { fetchInput } from "./util";

async function part1() {
  const input = await fetchInput(2);

  let s = 0;

  for (const line of input.split(",")) {
    const [a, b] = line.split("-").map(Number);

    for (let i = a; i <= b; i++) {
      const id = i.toString();
      if (
        id.slice(0, Math.floor(id.length / 2)) ===
        id.slice(Math.floor(id.length / 2), id.length)
      ) {
        s += i;
      }
    }
  }

  console.log(s);
}

function is_invalid(id: string): boolean {
  for (let i = 1; i < Math.floor(id.length / 2) + 1; i++) {
    if (id.length % i !== 0) {
      continue;
    }

    let valid = false;
    let part = id.slice(0, i);
    for (let j = i; j < id.length; j += i) {
      if (id.slice(j, j + i) !== part) {
        valid = true;
        break;
      }
    }
    if (!valid) {
      return true;
    }
  }

  return false;
}

async function part2() {
  const input = await fetchInput(2);

  let s = 0;

  for (const line of input.split(",")) {
    const [a, b] = line.split("-").map(Number);

    for (let i = a; i <= b; i++) {
      if (is_invalid(i.toString())) {
        s += i;
      }
    }
  }

  console.log(s);
}

part1();
part2();
