def get_input():
    s = None
    ret = []

    with open("input", "r") as file:
        s = list(map(list, file.read().split()))

    for line in s:
        r = list(map(int, line))
        r.insert(0, 1000)
        r.append(1000)
        ret.append(r)
    ones = [1000] * len(ret[0]) 
    ret.insert(0, ones)
    ret.insert(len(ret), ones)

    return ret

def find_lows(inp):
    coors = []
    for height in range(1, len(inp) - 1):
        for width in range(1, len(inp[height]) - 1):
            curr = inp[height][width]
            if (inp[height-1][width] > curr and inp[height][width-1] > curr
                and inp[height+1][width] > curr and inp[height][width+1] > curr):
                coors.append([height, width])

    return coors

def find_basins(inp, lows):
    sizes = []
    for low in lows:
        y, x = low
        sizes.append(len(traverse(inp, y, x, [[x, y]])))

    return sizes

def traverse(inp, y, x, visited):
    if ([x + 1, y] not in visited and x + 1 < len(inp[0]) and inp[y][x + 1] not in (9, 1000) and
        inp[y][x + 1] > inp[y][x]):
        visited.append([x + 1, y])
        traverse(inp, y, x + 1, visited)

    if ([x, y + 1] not in visited and y + 1 < len(inp) and inp[y + 1][x] not in (9, 1000) and
        inp[y + 1][x] > inp[y][x]):
        visited.append([ x, y + 1 ])
        traverse(inp, y + 1, x, visited)

    if ([x - 1, y] not in visited and x - 1 > 0 and inp[y][x - 1] not in (9, 1000) and
        inp[y][x - 1] > inp[y][x]):
        visited.append([ x - 1, y ])
        traverse(inp, y, x - 1, visited)

    if ([x, y - 1] not in visited and y - 1 > 0 and inp[y - 1][x] not in (9, 1000) and
        inp[y - 1][x] > inp[y][x]):
        visited.append([ x, y - 1 ])
        traverse(inp, y - 1, x, visited)

    return visited

def calc(basins):
    size = 1
    for i in sorted(basins)[len(basins) - 3:]:
        size *= i

    return size


if __name__=="__main__":
    inp = get_input()
    lows = find_lows(inp)
    basins = find_basins(inp, lows)

    res = calc(basins)

    print(res)