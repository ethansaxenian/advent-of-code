const { fetchInput } = require("./util/util.js");

const parseInput = (input) => {
  return input.map((line) => {
    return line.split(" ").map((num) => Number(num));
  });
};

function getDifferences(nums) {
  const lists = [];
  let next = [...nums];
  while (!next.every((num) => num === 0)) {
    const differences = [];
    for (let i = 0; i < next.length - 1; i++) {
      differences.push(next[i + 1] - next[i]);
    }
    lists.push(differences);
    next = [...differences];
  }
  return lists;
}

function getNext(nums) {
  const lists = getDifferences(nums);
  for (let i = lists.length - 1; i > 0; i--) {
    const diffs = lists[i];
    const nextDiff = diffs[diffs.length - 1];
    const nextList = lists[i - 1];
    nextList.push(nextList[nextList.length - 1] + nextDiff);
  }
  return nums[nums.length - 1] + lists[0][lists[0].length - 1];
}

function getFirst(nums) {
  const lists = getDifferences(nums);
  for (let i = lists.length - 1; i > 0; i--) {
    const diffs = lists[i];
    const nextDiff = diffs[0];
    const nextList = lists[i - 1];
    nextList.unshift(nextList[0] - nextDiff);
  }
  return nums[0] - lists[0][0];
}

async function day1() {
  const input = await fetchInput(9);
  const parsedInput = parseInput(input);
  let sum = 0;
  for (const line of parsedInput) {
    sum += getNext(line);
  }
  console.log(sum);
}

async function day2() {
  const input = await fetchInput(9);
  const parsedInput = parseInput(input);
  let sum = 0;
  for (const line of parsedInput) {
    sum += getFirst(line);
  }
  console.log(sum);
}

day1();
day2();
