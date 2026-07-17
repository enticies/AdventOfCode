def getInput(filename):
    with open(filename) as inp:
        return inp.read().split()

def main():
    inp = list(map(int, getInput('input')))
    
    for x in range(len(inp)):
        for y in range(len(inp)):
            if y != x and inp[x] + inp[y] == 2020:
                print(inp[y] * inp[x])
                return

if __name__ == '__main__':
    main() 