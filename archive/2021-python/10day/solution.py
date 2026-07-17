def get_input():
    with open("input", "r") as file:
        return file.read().split()

def find_corrupted(inp):
    ret = []
    m = {"(": ")", "{": "}", "[": "]", "<": ">"}

    for line in inp:
        stack = []
        for char in line:
            if char in m:
                stack.append(m[char])
            else:
                if char == stack[-1]:
                    stack.pop()
                else:
                    ret.append(char)
                    break
    return ret

def calc(s):
    scores = {")": 3, "]": 57, "}": 1197, ">": 25137}
    ret = 0

    for i in s:
        ret += scores[i]

    return ret
            
if __name__=="__main__":
    inp = get_input()
    r = find_corrupted(inp)
    c = calc(r)

    print(c)