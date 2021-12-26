with open('problem1input.txt') as f:
    depths = f.readlines()

num_increases = 0
for i in range(len(depths)-3):
    if int(depths[i]) < int(depths[i+3]):
        num_increases +=1

print(num_increases)