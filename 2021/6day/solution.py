def parse_file():
    with open("input", "r") as file:
        ints = list(map(int, file.read()[:-1:].split(",")))

    return ints

def fun(ints, days):
    fish = dict()

    for i in range(9):
        if i in ints:
            fish[i] = ints.count(i)
        else:
            fish[i] = 0

    for i in range(days):
        cpy = dict(fish)
        day0 = fish[0]
        fish[0] = 0
        for a in range(8, 0, -1):
            cpy[a - 1] = fish[a]
        cpy[8] = day0
        cpy[6] += day0

        fish = cpy


    return fish.values()


if __name__=="__main__":
    ints = parse_file()
    result = fun(ints, 256)


    print(result)
    print(sum(result))
