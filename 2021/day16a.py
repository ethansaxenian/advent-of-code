import io


class Stream:
    def __init__(self, s):
        self.s = s
        self.i = 0

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

hex = open("day16_input.txt", "r").readline().strip()

b = bin(int(hex, 16))[2:].zfill(len(hex) * 4)

def btoi(b):
    return int(b, 2)


s = Stream(b)

total = 0

while True:
    try:
        version = btoi(s.read(3))
        total += version
        type_id = btoi(s.read(3))
        if type_id == 4:
            while int(s.read()) != 0:
                s.read(4)
            s.read(4)

        else:
            length_type_id = int(s.read())
            if length_type_id == 0:
                l_sub = btoi(s.read(15))
            elif length_type_id == 1:
                n_sub = btoi(s.read(11))
    except ValueError:
        break

print(total)
