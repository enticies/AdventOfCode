def getInput(filename):
    with open(filename) as inp:
        return inp.read().split()


def main():
    inp = getInput('./input')
    
    sum = 0
    
    m = {
        'one': 1,
        'two': 2,
        'three': 3,
        'four': 4,
        'five': 5,
        'six': 6,
        'seven': 7,
        'eight': 8,
        'nine': 9
    }
    
    for line in inp:
        num = ""

        convertedLine = line 
        for key, value in m.items():
            convertedLine = convertedLine.replace(str(value), key)
       
        left = [99999]
        right = [99999]
        for key, value in m.items():
            l = convertedLine.find(key)
            r = convertedLine[::-1].find(key[::-1])
            if l != -1 and l < left[0]:
                left = [l, str(value)]
            if r != -1 and r < right[0]:
                right = [r, str(value)]
       
     
        sum += int(left[1] + right[1])
    
        print(sum) 
    

if __name__ == '__main__':
    main()
