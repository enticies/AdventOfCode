def getInput(filename):
    with open(filename) as inp:
        return inp.read().split('\n')
    
def main():
    inp = getInput('input')
    red, blue, green = 0, 0, 0
    total = 0
    for line in inp:
        correct = True
        for cubes in line.split(':')[1].split(';'):
            for cube in cubes.split(','):
                splitCube = cube.split(' ')
                print(splitCube) 
                    
                if splitCube[2] == 'red':
                    red = max(red, int(splitCube[1]))
                elif splitCube[2] == 'green':
                    green = max(green, int(splitCube[1]))
                elif splitCube[2] == 'blue':
                    blue = max(blue, int(splitCube[1]))
        total += red * blue * green
        red, blue, green = 0, 0, 0
           
    print(total) 
    
        

if __name__ == '__main__':
    main()
