def parse_file():
    with open("input", "r") as file:
        ints = list(map(int, file.read()[:-1:].split(",")))

    return ints

def fun(ints):
    for i in range(80):
        t = []
        for j in range(len(ints)):
            if ints[j] == 0:
                t.extend([6, 8])
            else:
                t.append(ints[j] - 1)
        ints = t

    return ints



if __name__=="__main__":
    ints = parse_file()
    result = fun(ints)

    print(len(result))
