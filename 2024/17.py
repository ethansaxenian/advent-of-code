import util


def run(a: int, b: int, c: int, prog: list[int]) -> list[int]:
    out = []
    ip = 0
    while ip < len(prog):
        opcode = prog[ip]
        literal_operand = prog[ip + 1]
        combo_operand = (
            literal_operand
            if 0 <= literal_operand <= 3
            else {4: a, 5: b, 6: c}[literal_operand]
        )
        match opcode:
            case 0:
                a //= 2**combo_operand
            case 1:
                b ^= literal_operand
            case 2:
                b = combo_operand % 8
            case 3:
                if a > 0:
                    ip = literal_operand
                    continue
            case 4:
                b ^= c
            case 5:
                out.append(combo_operand % 8)
            case 6:
                b = a // (2**combo_operand)
            case 7:
                c = a // (2**combo_operand)
            case _:
                pass

        ip += 2

    return out


def part1(input: str) -> str:
    a, b, c, _, prog = input.splitlines()
    A = int(a.split()[-1])
    B = int(b.split()[-1])
    C = int(c.split()[-1])
    prog = list(map(int, prog.split()[-1].split(",")))

    return ",".join(map(str, run(A, B, C, prog)))


def part2(input: str) -> int:
    _, b, c, _, prog = input.splitlines()
    B = int(b.split()[-1])
    C = int(c.split()[-1])
    prog = list(map(int, prog.split()[-1].split(",")))

    """
    b = a % 8
    b = b ^ 1
    c = a >> b
    a = a >> 3
    b = b ^ c
    b = b ^ 6
    out -> b % 8


    out -> ((((a%8)^1)^(a//(2**((a%8)^1)))) ^ 6) % 8
    a = a // (2**3)
    """
    A = 0
    s = -1
    while (out := run(A, B, C, prog)) != prog:
        if out[s:] == prog[s:]:
            A *= 8
            s -= 1
            continue

        A += 1

    return A


if __name__ == "__main__":
    util.run(17, part1, part2)
