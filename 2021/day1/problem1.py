with open('problem1input.txt') as f:
    prev_depth = 1000000
    current_depth = 0
    num_increases = 0
    for line in f:
        current_depth = int(line)
        if current_depth > prev_depth:
            num_increases += 1
        prev_depth = current_depth
print(num_increases)