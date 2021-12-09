
inputFile = open("./days/day9.txt", "r")

heights = inputFile.read().split("\n")

# add a border to the matrix
h = len(heights)+2
w = len(heights[0])+2
heightmap = [[9 for x in range(w)] for y in range(h)] 


for i in range(len(heights)):
    for j in range(len(heights[0])):
        heightmap[i+1][j+1] = heights[i][j]

heightsum = 0
risksum = 0

for i in range(1,len(heightmap)-1):
    for j in range(1,len(heightmap[i])-1):
        
        point = int(heightmap[i][j])

        top, left, right, down = 0,0,0,0
        top = int(heightmap[i-1][j])
        left = int(heightmap[i][j-1])
        right = int(heightmap[i][j+1])
        down = int(heightmap[i+1][j])

        if(point < top and point < left and point < right and point < down):
            heightsum += point
            risksum += 1

print(heightsum)
print(risksum)