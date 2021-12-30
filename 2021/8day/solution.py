def get_input():
    ret = []
    with open("input", "r") as file:
        line = file.readline()
        while line:
            t = line.strip().split("|")
            if len(t) == 2:
                left, right = t[0][:-1:], t[1][1:]
                ret.append([left, right])
            line = file.readline()
    return ret

def count(inp):
    h = {2: 1, 4: 4, 3: 7, 7: 8}
    ret = 0
    for line in inp:
        for value in line[1].split():
            if len(value) in h:
                ret += 1

    return ret




if __name__=="__main__":
    parsed = get_input()

    count = count(parsed)

    print(count)


