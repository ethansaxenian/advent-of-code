#!/bin/bash

set -euo pipefail

source .env

puzzle_input=$(curl \
  "https://adventofcode.com/2023/day/1/input" \
  --cookie "session=$SESSION_COOKIE" \
  --silent
)

part1() {
  echo "$puzzle_input" \
    | sed "s/[a-z]//g" \
    | sed -E "s/^(.).*(.)$/\1\2/" \
    | sed -E "s/^.$/&&/" \
    | tr "\n" "+" \
    | sed "s/+$//" \
    | bc
}

get_digit_at_index() {
  line="$1"
  index="$2"
  case "${line:index}" in
    [0-9]*) echo "${line:index:1}" && return ;;
    one*) echo 1 && return ;;
    two*) echo 2 && return ;;
    three*) echo 3 && return ;;
    four*) echo 4 && return ;;
    five*) echo 5 && return ;;
    six*) echo 6 && return ;;
    seven*) echo 7 && return ;;
    eight*) echo 8 && return ;;
    nine*) echo 9 && return ;;
  esac
}

part2() {
  calibration_values=$(
    while read -r line; do
      for ((index = 0; index < ${#line}; ++index)); do
        digit="$(get_digit_at_index "$line" "$index")"
        if [[ -n "$digit" ]]; then
          printf "%s" "$digit"
          break
        fi
      done
      for ((index = ${#line} - 1; index >= 0; index--)); do
        digit="$(get_digit_at_index "$line" "$index")"
        if [[ -n "$digit" ]]; then
          echo "$digit"
          break
        fi
      done
    done <<< "$puzzle_input"
  )
  echo "$calibration_values" | tr "\n" "+" | sed "s/+$//" | bc
}

part1
part2
