from functools import cache

z_divs = (1, 1, 1, 26, 1, 1, 1, 26, 26, 1, 26, 26, 26, 26)
check = (14, 12, 11, -4, 10, 10, 15, -9, -9, 12, -15, -7, -10, 0)
offsets = (7, 4, 8, 1, 5, 14, 12, 10, 5, 7, 6, 8, 4, 6)

"""
push n[0] + 7
push n[1] + 4
push n[2] + 8
pop n[3] == n[2] + 8 - 4
push n[4] + 5
push n[5] + 14
push n[6] + 12
pop n[7] == n[6] + 12 - 9
pop n[8] == n[5] + 14 - 9
push n[9] + 7
pop n[10] == n[9] + 7 - 15
pop n[11] == n[4] + 5 - 7
pop n[12] == n[1] + 4 - 10
pop n[13] == n[0] + 7

n[3] == n[2] + 4
n[7] == n[6] + 3
n[8] == n[5] + 5
n[10] == n[9] - 8
n[11] == n[4] - 2
n[12] == n[1] - 6
n[13] == n[0] + 7

n[0] = 1
n[1] = 7
n[2] = 1
n[3] = 5
n[4] = 3
n[5] = 1
n[6] = 1
n[7] = 4
n[8] = 6
n[9] = 9
n[10] = 1
n[11] = 1
n[12] = 1
n[13] = 8

n = 17153114691118
"""


@cache
def process_digit(d, prev_z, i):
	z = prev_z
	x = (z % 26 + check[i]) == d
	if check[i] <= 9:
		z //= z_divs[i]  # will always be 26 here
	if not x:
		z = z * 26 + d + offsets[i]

	return z


def process_number(n):
	z = 0

	for i, d in enumerate(n):
		z = process_digit(int(d), z, i)

	return z


print(process_number("17153114691118"))
