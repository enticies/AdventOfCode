measurements = []

with open("input", "r") as file:
    for measurement in file:
        measurements.append(measurement.replace("\n", ""))

measurements = list(map(int, measurements))

increased = 0

for i in range(1, len(measurements)):
    if measurements[i] > measurements[i-1]:
        increased += 1

start = 0
stop = 3

sum_increased = 0

# measurements = [607, 618, 618, 617, 647, 716, 769, 792]

while stop < len(measurements):
    current_sum = sum(measurements[start:stop:])
    start += 1
    stop += 1

    if sum(measurements[start:stop:]) > current_sum:
        sum_increased += 1

print(sum_increased)


