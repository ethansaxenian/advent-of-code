import { run } from "./util";

async function part1(input: string): Promise<number> {
  const grid = input.split("\n").map((line) => line.split(""));

  let splits = 0;
  for (let r = 1; r < grid.length; r++) {
    const row = grid[r];
    for (let c = 0; c < row.length; c++) {
      if (grid[r - 1][c] === "S" || grid[r - 1][c] === "|") {
        switch (grid[r][c]) {
          case ".":
            grid[r][c] = "|";
            break;
          case "^":
            if (grid[r][c - 1] == ".") {
              grid[r][c - 1] = "|";
            }
            if (grid[r][c + 1] == ".") {
              grid[r][c + 1] = "|";
            }
            splits += 1;
            break;
        }
      }
    }
  }
  return splits;
}

async function part2(input: string): Promise<number> {
  const grid = input.split("\n").map((line) => line.split(""));

  const memo = new Map<string, number>();

  function dfs(r: number, c: number): number {
    if (r === grid.length) {
      return 1;
    }

    const key = `${r},${c}`;
    let val = memo.get(key);
    if (val !== undefined) return val;

    let res;
    const curr = grid[r][c];
    switch (curr) {
      case ".":
        res = dfs(r + 1, c);
        break;
      case "^":
        res = dfs(r + 1, c - 1) + dfs(r + 1, c + 1);
        break;
      default:
        res = 0;
    }

    memo.set(key, res);
    return res;
  }

  return dfs(1, grid[0].indexOf("S"));
}

run(7, part1, part2);
