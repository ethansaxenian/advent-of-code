import { parseArgs } from "util";
import { execSync } from "child_process";
import { readFile, writeFile } from "fs/promises";

const AOC_EMAIL = process.env.AOC_EMAIL!;
const AOC_COOKIE = process.env.AOC_COOKIE!;
const AOC_YEAR = process.env.AOC_YEAR!;

export async function fetchInput(day: number): Promise<string> {
  const inputFile = `${__dirname}/input/day${day}.txt`;

  try {
    const text = await readFile(inputFile, "utf8");
    return text.trim();
  } catch (err: any) {}

  const res = await fetch(
    `https://adventofcode.com/${AOC_YEAR}/day/${day}/input`,
    {
      headers: new Headers({
        Cookie: `session=${AOC_COOKIE}`,
        "User-Agent": `github.com/ethansaxenian/advent-of-code/tree/main/${AOC_YEAR} by ${AOC_EMAIL}`,
      }),
    },
  );

  const text = await res.text();

  await writeFile(inputFile, text);

  return text.trim();
}

function parseArguments() {
  const { values } = parseArgs({
    options: {
      part: {
        type: "string",
        short: "p",
        default: "1",
      },
      test: {
        type: "boolean",
        short: "t",
        default: false,
      },
    },
  });

  if (values.part !== "1" && values.part !== "2") {
    throw new Error("Part must be 1 or 2");
  }

  return values;
}

export async function run(
  day: number,
  part1: (input: string) => Promise<string | number>,
  part2: (input: string) => Promise<string | number>,
): Promise<void> {
  const args = parseArguments();

  const input = await (args.test
    ? readFile("/dev/stdin", "utf-8")
    : fetchInput(day));

  const func = args.part === "1" ? part1 : part2;

  const start = process.hrtime.bigint();
  const answer = await func(input.trim());
  const duration = process.hrtime.bigint() - start;

  if (!args.test) {
    execSync("pbcopy", { input: answer.toString() });
  }

  console.log(answer);
  console.log(`Ran in: ${Number(duration) / 1e6}ms`);
}
