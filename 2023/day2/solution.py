def getInput(filename):
    with open(filename) as inp:
        return inp.read().split('\n')
    
def main():
    inp = getInput('input')
    m = {
        'red': 12,
        'green': 13,
        'blue': 14
    }
    correctIds = [] 
    for line in inp:
        correct = True
        for cubes in line.split(':')[1].split(';'):
            for cube in cubes.split(','):
                splitCube = cube.split(' ')
                
                if m[splitCube[2]] < int(splitCube[1]):
                    correct = False
        if correct:
            correctIds.append(int(line.split(':')[0].split(' ')[1]))

    

        
           
         
    print(sum(correctIds)) 
            
    
        

if __name__ == '__main__':
    main()
