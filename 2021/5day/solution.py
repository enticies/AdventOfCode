def parse_file():
    lines = []
    with open("input", "r") as file:
        line = file.readline()
        while line:
            lines.append(line.strip().replace(" ", "").split("->"))
            line = file.readline()

    return list(map(lambda x: list(map(lambda a: a.split(","), x)), lines))
        

if __name__=="__main__":
    lines = parse_file()
    print(lines)
