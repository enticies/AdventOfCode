# hacky

gamma = ""
epsilon = ""

oxygen = []
co2 = []

with open("input", "r") as file:
    readings = file.read().split()
    for i in range(len(readings[0])):
        current = 0
        for a in readings:
            if a[i] == "1":
                current += 1

        most_common = None
        least_common = None

        if current >= len(readings) // 2:
            most_common = "0"
        else:
            most_common = "1"

        temp = []
    
        for c in readings:
            if c[i] == most_common:
                temp.append(c)

        readings = temp
    
        if len(readings) == 1:
            break

print(readings)




        
