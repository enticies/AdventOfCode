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

def first_pass(line):
    h = {2: 1, 4: 4, 3: 7, 7: 8}
    unique = dict()
    notunique = [] 
    for value in line:
        if len(value) in h:
            unique[h[len(value)]] = value
        else:
            notunique.append(value)
    
    return unique, notunique

def count(line):
    ret = dict()
    letters = ["a", "b", "c", "d", "e", "f", "g"]

    for letter in letters:
        for word in line:
            for a in word:
                if letter == a and a not in ret:
                    ret[a] = 1
                elif letter == a:
                    ret[a] += 1

    return ret

def find_nine(nonunique, counted):
    a = None
    for i in counted:
        if counted[i] == 4:
            a = i
    for val in nonunique:
        if len(val) == 6 and a not in val:
            return val

def find_zero(nonunique):
    for val in nonunique:
        if len(val) == 6:
            return val

def find_two(nonunique, counted):
    a = None
    for i in counted:
        if counted[i] == 4:
            a = i

    for l in nonunique:
        if a in l:
            return l


def find_five_three(nonunique, c):
    five = None
    three = None
    for a in c:
        if c[a] == 8:
            for l in nonunique:
                if a not in l:
                    five = l
                else:
                    three = l
    return five, three
                    

def find_six(unique, nonunique, counted):
    a = None
    for letter in unique[4]:
        if  counted[letter] == 7:
            a = letter
    for val in nonunique:
        if len(val) == 6 and a in val:
            return val


def com(unique, nonunique, c):
    nine = find_nine(nonunique, c)
    nonunique.remove(nine)
    six = find_six(unique, nonunique, c)
    nonunique.remove(six)
    zero = find_zero(nonunique)
    nonunique.remove(zero)
    two = find_two(nonunique, c)
    nonunique.remove(two)
    five, three = find_five_three(nonunique, c)

    ret = unique
    ret[5] = five
    ret[6] = six
    ret[0] = zero
    ret[2] = two
    ret[5] = five
    ret[3] = three
    ret[9] = nine

    ret = {tuple(sorted(v)): k for k, v in ret.items()}

    return ret

def calculate(parsed):
    summ = []
    for line in parsed:
        unique, nonunique = first_pass(line[0].split())
        c = count(line[0].split())
        ret = ""
        final = com(unique, nonunique, c)
        print(final)
        for a in line[1].split():
            if tuple(sorted(a)) in final:
                ret += str(final[tuple(sorted(a))])

        summ.append(int(ret))

    print(summ)






def main(parsed):
    calculate(parsed)


if __name__=="__main__":
    parsed = get_input()
    main(parsed)
