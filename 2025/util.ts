import { readFile, writeFile } from "fs/promises";

export async function fetchInput(day: number): Promise<string> {
  const inputFile = `${process.cwd()}/input/day${day}.txt`;

  try {
    const text = await readFile(inputFile, "utf8");
    return text.trim();
  } catch (err: any) {}

  const res = await fetch(
    `https://adventofcode.com/${process.env.AOC_YEAR}/day/${day}/input`,
    {
      headers: new Headers({
        Cookie: `session=${process.env.AOC_COOKIE}`,
      }),
    },
  );

  const text = await res.text();

  await writeFile(inputFile, text);

  return text.trim();
}
