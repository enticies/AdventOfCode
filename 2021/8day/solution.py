def get_input():
    ret = []
    with open("input", "r") as file:
        line = file.readline()
        while line:
            t = line.strip().split("|")
            if len(t) == 2:
                left, right = t[0][:-1:], t[1][1:]
                ret.append([left, right])
            line = file.readline()
    return ret

def get_unique(line):
    h = {2: 1, 4: 4, 3: 7, 7: 8}
    unique = dict()
    for value in line[0].split():
        if len(value) in h:
            unique[h[len(value)]] = value
    
    return unique

def count_letters(line):
    ret = dict()

    for value in line.split():
        for letter in value:
            if letter in ret:
                ret[letter] += 1
            else:
                ret[letter] = 1
    return ret

def find_six(line, unique):
    for value in line.split():
        if len(value) == 6:
            for letter in unique[7]:
                if letter not in value:
                    return value

def find_zero(line, unique, six):
    for value in line.split():
        if len(value) == 6:
            for letter in unique[4]:
                if letter not in value and value != six:
                    return value

def find_nine(line, unique, six, zero):
    for value in line.split():
        if len(value) == 6 and value not in (six, zero):
            return value



def find_three(line, unique):
    for value in line.split():
        p = 0
        if len(value) == 5:
            for letter in unique[1]:
                if letter in value:
                    p += 1
        if p == 2:
            return value

def find_two_five(line, unique, three):
    five = None
    two = None
    for value in line.split():
        if len(value) == 5 and value != three:
            p = 0
            for letter in unique[4]:
                if letter not in value:
                    p += 1
            if p == 1:
                five = value
            elif p == 2:
                two = value

    return two, five


def calc(inp):
    total = 0
    for line in inp:
        unique = get_unique(line)
        six = find_six(line[0], unique)
        zero = find_zero(line[0], unique, six)
        nine = find_nine(line[0], unique, six, zero)
        three = find_three(line[0], unique)
        two, five = find_two_five(line[0], unique, three) 


        table = dict()

        for i in unique:
            table[tuple(sorted(unique[i]))] = i

        table[tuple(sorted(six))] = 6
        table[tuple(sorted(zero))] = 0
        table[tuple(sorted(nine))] = 9
        table[tuple(sorted(five))] = 5
        table[tuple(sorted(three))] = 3
        table[tuple(sorted(two))] = 2
        f = ""
        for pot in line[1].split():
            if tuple(sorted(pot)) in table:
                f += str(table[tuple(sorted(pot))])
        total += int(f)

    return total

if __name__=="__main__":
    inp = get_input()
    result = calc(inp)
    print(result)

