depth = 0
x = 0
aim = 0

with open("input", "r") as file:
    for line in file.read().split("\n"):
        if line:
            line = line.split()
            if line[0] == "forward":
                x += int(line[1])
                depth += aim * int(line[1])
            elif line[0] == "up":
                aim -= int(line[1])
            elif line[0] == "down":
                aim += int(line[1])


print(x * depth)
