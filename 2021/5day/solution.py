from pprint import pprint

def parse_file():
    lines = []
    with open("input", "r") as file:
        line = file.readline()
        while line:
            lines.append(line.strip().replace(" ", "").split("->"))
            line = file.readline()

    return list(map(lambda x: list(map(lambda a: list(map(int, a.split(","))), x)), lines))

def create_table(lines):
    max_x = 0
    max_y = 0

    for line in lines:
        max_x = max(max_x, line[0][0], line[1][0])
        max_y = max(max_y, line[0][1], line[1][1])

    table = []

    # max_x = 5
    # max_y = 5


    for i in range(max_y + 1):
        table.append([0] * (max_x + 1))

    return table

def mark_line(table, coors):
    x1, x2, y1, y2 = coors[0][0], coors[1][0], coors[0][1], coors[1][1]
   

    if abs(x1 - x2) == abs(y1 - y2):
        for i in range(abs(y1 - y2) + 1):
            dist_y = i if y1 < y2 else -i
            dist_x = i if x1 < x2 else -i

            table[y1 + dist_y][x1 + dist_x] += 1
    elif x1 == x2:
        for i in range(abs(y1 - y2) + 1):
            dist = i if y1 < y2 else -i
            table[y1 + dist][x1] += 1 
    elif y1 == y2:
        for i in range(abs(x1 - x2) + 1):
            dist = i if x1 < x2 else -i
            table[y1][x1 + dist] += 1
    

    return table

def mark_table(table, lines):
    for line in lines:
        table = mark_line(table, line)
    return table

def calculate_table(table):
    total = 0

    for row in table:
        for i in row:
            if i >= 2:
                total += 1
    return total

if __name__=="__main__":
    lines = parse_file()
    table = create_table(lines)
    # marked = mark_table(table, [[[2, 4], [4, 2]]])
    marked = mark_table(table, lines)
    # pprint(marked)
    total = calculate_table(marked)
    print(total)

