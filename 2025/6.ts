import { run } from "@util";

async function part1(input: string): Promise<number> {
  const lines = input.split("\n");
  const ops = new Map<number, string>();
  const totals = new Map<number, number>();

  const opsLine = lines[lines.length - 1].split(/\s+/);
  for (let i = 0; i < opsLine.length; i++) {
    ops.set(i, opsLine[i]);
  }

  for (let i = lines.length - 2; i >= 0; i--) {
    const nums = lines[i]
      .split(/\s+/)
      .filter((v) => v !== "")
      .map(Number);

    for (let j = 0; j < nums.length; j++) {
      const fn = ops.get(j);
      let curr;
      switch (fn) {
        case "+":
          curr = totals.get(j) || 0;
          totals.set(j, curr + nums[j]);
          break;
        case "*":
          curr = totals.get(j) || 1;
          totals.set(j, curr * nums[j]);
          break;
      }
    }
  }

  return Array.from(totals.values()).reduce((acc, val) => acc + val, 0);
}

async function part2(input: string): Promise<number> {
  const lines = input.split("\n");

  const ops = new Map<number, string>();
  const lastLine = lines[lines.length - 1].split("");
  for (let i = 0; i < lastLine.length; i++) {
    const char = lastLine[i];
    if (char === "*" || char === "+") {
      ops.set(i, char);
    }
  }

  const rawNums: Array<Array<string | undefined>> = [];
  for (let r = 0; r < lines[0].length; r++) {
    const row = [];
    for (let c = 0; c < lines.length - 1; c++) {
      row.push(undefined);
    }
    rawNums.push(row);
  }

  for (let c = lines.length - 2; c >= 0; c--) {
    const line = lines[c];
    for (let r = 0; r < line.length; r++) {
      rawNums[r][c] = line[r];
    }
  }
  const nums: Array<number | undefined> = rawNums
    .map((n) => n.filter(Boolean).join(""))
    .map((n) => (n.match(/^\s+$/) ? undefined : Number(n)));

  let total = 0;
  let currOp = ops.get(0);
  let currTotal = nums[0];

  if (currTotal === undefined) {
    throw new Error("impossible");
  }

  for (let i = 1; i < nums.length; i++) {
    const num = nums[i];
    currOp = ops.get(i) || currOp;

    if (num === undefined) {
      total += currTotal;
      currTotal = ops.get(i + 1) === "*" ? 1 : 0;
      continue;
    }

    switch (currOp) {
      case "+":
        currTotal += num;
        break;
      case "*":
        currTotal *= num;
        break;
    }
  }

  return total + currTotal;
}

run(6, 2025, part1, part2);
