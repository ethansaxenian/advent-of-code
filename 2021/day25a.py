import time

rights = set()
downs = set()

lines = [l.strip() for l in open("day25_input.txt", "r").readlines()]
for y, row in enumerate(lines):
	for x in range(len(row)):
		if row[x] == ">":
			rights.add((x, y))
		if row[x] == "v":
			downs.add((x, y))


height = len(lines)
width = len(lines[0])


def print_grid(rights, downs):
	for y in range(height):
		for x in range(width):
			if (x, y) in rights:
				print(">", end="")
			elif (x, y) in downs:
				print("v", end="")
			else:
				print(".", end="")
		print()


i = 0
while True:
	no_movement = True
	new_rights = set()
	new_downs = set()

	occupied_slots = downs | rights
	for x, y in rights:
		if ((x+1) % width, y) in occupied_slots:
			new_rights.add((x, y))
		else:
			new_rights.add(((x+1) % width, y))
			no_movement = False

	occupied_slots = new_rights | downs
	for x, y in downs:
		if (x, (y+1) % height) in occupied_slots:
			new_downs.add((x, y))
		else:
			new_downs.add((x, (y+1) % height))
			no_movement = False

	rights = new_rights
	downs = new_downs
	i += 1
	if no_movement:
		print(i)
		break
