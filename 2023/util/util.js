const { aoc_cookie } = require("../../aoc-cookie.json");

async function fetchInput(day) {
  const headers = new Headers({
    Cookie: `session=${aoc_cookie}`,
  });
  const res = await fetch(`https://adventofcode.com/2023/day/${day}/input`, {
    headers,
  });

  const text = await res.text();

  return text.split("\n").filter((x) => x !== "");
}

module.exports = {
  fetchInput,
};
