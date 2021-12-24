def get_input():
    with open("input", "r") as file:
        file = list(map(int, file.read().split(",")))

    return file

def find_min(positions):
    min_fuel = 1000000000

    for i in positions:
        ret = 0
        for a in positions:
            ret += abs(a - i)

        min_fuel = min(ret, min_fuel)


    return min_fuel

if __name__=="__main__":
    positions = get_input()
    min_fuel = find_min(positions)

    print(min_fuel)

