const { fetchInput } = require("./util/util.js");

function parseLine(line) {
  [left, right] = line.split("|");
  const leftNums = left
    .split(" ")
    .splice(2)
    .filter((x) => x !== "");
  const rightNums = right.split(" ").filter((x) => x !== "");

  let numOverlap = 0;
  for (const num of rightNums) {
    if (leftNums.includes(num)) {
      numOverlap++;
    }
  }

  return numOverlap;
}

async function part1() {
  let sum = 0;
  const puzzleInput = await fetchInput(4);
  for (const line of puzzleInput) {
    const numOverlap = parseLine(line);

    sum += Math.floor(2 ** numOverlap / 2);
  }

  console.log(sum);
}

async function part2() {
  const puzzleInput = await fetchInput(4);

  const cards = [];
  for (const line of puzzleInput) {
    const numOverlap = parseLine(line);
    cards.push({ numOverlap, count: 1 });
  }

  for (let i = 0; i < cards.length; i++) {
    const card = cards[i];
    for (let j = 1; j <= card.numOverlap; j++) {
      cards[i + j].count += card.count;
    }
  }

  const totalNumCards = cards.reduce((acc, curr) => acc + curr.count, 0);
  console.log(totalNumCards);
}

part1();
part2();
