def getInput(filename):
    with open(filename) as inp:
        return inp.read().split('\n')

def main():
    inp = getInput('input')
    count = 0
    for line in inp:
        if line != '':
            splitLine = line.split(' ')
            minAmount = int(splitLine[0].split('-')[0])
            maxAmount = int(splitLine[0].split('-')[1])
            letter = splitLine[1][0]
            password = splitLine[2]
            letterCount = password.count(letter)
            if letterCount >= minAmount and letterCount <= maxAmount:
                count += 1
    print(count) 
    
if __name__ == '__main__':
    main() 