def get_input():
    with open("input", "r") as file:
        file = list(map(int, file.read().split(",")))

    return file

def find_min(positions):
    min_fuel = 1000000000
    for i in range(max(positions)):
        ret = 0
        for a in positions:
            ret += add(i, a) 

        min_fuel = min(ret, min_fuel)


    return min_fuel


memo = {}
def add(a, b):
    if abs(a - b) in memo:
        return memo[abs(a - b)]

    ret = 0
    i = 0
    cost = 1
    while i != abs(a - b):
        ret += cost
        cost += 1
        i += 1
    memo[abs(a - b)] = ret
    return ret



if __name__=="__main__":
    positions = get_input()
    min_fuel = find_min(positions)
    print(min_fuel)

