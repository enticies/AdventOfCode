def get_input():
    ret = []
    with open("example_input") as file:
        for line in file.read().split():
            a = list(map(int, list(line)))
            a.insert(0, -1)
            a.append(-1)
            ret.append(a)

    b = [-1] * len(ret[0])
    ret.insert(0, b)
    ret.append(b)

    return ret

def fill(inp, height, width):
    inp[height + 1][width + 1] += 1
    inp[height + 1][width] += 1
    inp[height][width + 1] += 1
    inp[height - 1][width - 1] += 1
    inp[height - 1][width + 1] += 1
    inp[height + 1][width - 1] += 1

if __name__=="__main__":
    inp = get_input()
    fill(inp, 1, 1)
    for i in inp:
        print(i)


