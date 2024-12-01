from calendar import c
from itertools import count
from re import A


def getInput(filename):
    with open(filename) as inp:
        return inp.read().split('\n')[:-1:]


def main():
    inp = getInput('input')
    x = 3
    treeCount = 0
    width = len(inp[0])

    for y in range(1, len(inp)):
        if x > width - 1:
            x = abs(x - width)
        if inp[y][x] == '#':
            treeCount += 1
        x += 3

    print(treeCount)


def printArray(array):
    for i in array:
        print(i)


def repeatArray(array):
    newArray = array
    for i in range(len(array)):
        newArray[i] += array[i]
    return newArray


if __name__ == '__main__':
    main()
