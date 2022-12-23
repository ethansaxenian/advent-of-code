from dataclasses import dataclass
from math import prod


@dataclass
class Stream:
    s: str
    i: int = 0

    def read(self, n=1):
        try:
            x = self.s[self.i: self.i + n]
        except IndexError:
            return None
        self.i += n
        return x

    def peek(self, n=1):
        try:
            return self.s[self.i: self.i + n]
        except IndexError:
            return None

    def empty(self):
        return self.i >= len(self.s)


@dataclass
class Packet:
    version: int
    type_id: int
    subpackets: list
    data: int = None


hex = open("day16_input.txt", "r").readline().strip()
# hex = "9C0141080250320F1802104A08"

b = bin(int(hex, 16))[2:].zfill(len(hex) * 4)


def btoi(b):
    return int(b, 2)


def get_packet(s: Stream):
    version = btoi(s.read(3))
    type_id = btoi(s.read(3))
    subpackets = []
    data = None
    if type_id == 4:
        d = ""
        while int(s.read()) != 0:
            d += s.read(4)
        d += s.read(4)
        data = btoi(d)

    else:
        length_type_id = int(s.read())
        if length_type_id == 0:
            l_sub = btoi(s.read(15))
            new_s = Stream(s.read(l_sub))
            while not new_s.empty():
                subpackets.append(get_packet(new_s))
        elif length_type_id == 1:
            n_sub = btoi(s.read(11))
            for _ in range(n_sub):
                subpackets.append(get_packet(s))

    return Packet(version, type_id, subpackets, data)


def count_versions(p):
    if not p.subpackets:
        return p.version

    total = p.version
    for a in p.subpackets:
        total += count_versions(a)
    return total


def evaluate(p: Packet):
    if p.type_id == 4:
        assert p.data is not None
        return p.data

    assert p.data is None
    assert len(p.subpackets) > 0

    if p.type_id == 0:
        return sum(evaluate(a) for a in p.subpackets)
    elif p.type_id == 1:
        return prod(evaluate(a) for a in p.subpackets)
    elif p.type_id == 2:
        return min(evaluate(a) for a in p.subpackets)
    elif p.type_id == 3:
        return max(evaluate(a) for a in p.subpackets)

    assert len(p.subpackets) == 2
    if p.type_id == 5:
        return evaluate(p.subpackets[0]) > evaluate(p.subpackets[1])
    if p.type_id == 6:
        return evaluate(p.subpackets[0]) < evaluate(p.subpackets[1])
    if p.type_id == 7:
        return evaluate(p.subpackets[0]) == evaluate(p.subpackets[1])


tree = get_packet(Stream(b))
res = evaluate(tree)
print(res)
