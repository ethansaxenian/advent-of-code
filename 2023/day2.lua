local http = require("socket.http")
local ltn12 = require("ltn12")
local json = require("lunajson")


local file = io.open("../aoc-cookie.json", "rb")
if not file then return nil end
local json_string = file:read("*a")
file:close()

local t = json.decode(json_string)

local puzzle_input = {}
http.request({
  url = "https://adventofcode.com/2023/day/2/input",
  headers = { Cookie = "session=" .. t.aoc_cookie },
  sink = ltn12.sink.table(puzzle_input)
})


local lines = {}
for line in table.concat(puzzle_input):gmatch("[^\r\n]+") do
  table.insert(lines, line)
end

local function part1()
  local maxes = {
    r = 12,
    g = 13,
    b = 14
  }

  local total = 0

  for i, line in pairs(lines) do
    local is_valid = true

    for num, color in line:gmatch("(%d+) ([rgb])") do
      if tonumber(num) > maxes[color] then
        is_valid = false
        break
      end
    end

    if is_valid then
      total = total + i
    end
  end

  print(total)
end



local function part2()
  local total = 0

  for _, line in pairs(lines) do
    local maxes = {
      r = 0,
      g = 0,
      b = 0
    }
    for num, color in line:gmatch("(%d+) ([rgb])") do
      if tonumber(num) > maxes[color] then
        maxes[color] = tonumber(num)
      end
    end
    total = total + maxes.r * maxes.g * maxes.b
  end

  print(total)
end




part1()
part2()
