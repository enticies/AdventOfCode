import re

def getInput(filename):
    with open(filename) as inp:
        return inp.read().split('\n')
    
def main():
    inp = getInput('input')

    for row in range(len(inp)):
        for column in range(len(inp[row])):
            print(inp[row][column])
            

if __name__ == '__main__':
    main()