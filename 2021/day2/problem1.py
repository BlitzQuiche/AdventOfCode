h = d = 0
with open('input.txt') as f:
    for line in f:
        command, value = line.split()
        if command == "forward":
            h += int(value)
        elif command == "down":
            d += int(value)
        else:
            d -= int(value)

print(f"Horizontal:{h}\nDepth: {d}\nAnswer: {d*h}")


