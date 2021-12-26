h = d = a = 0
with open('input.txt') as f:
    for line in f:
        command, value = line.split()
        if command == "forward":
            h += int(value)
            d += a*int(value)
        elif command == "down":
            a += int(value)
        else:
            a -= int(value)

print(f"Horizontal:{h}\nDepth: {d}\nAnswer: {d*h}")


