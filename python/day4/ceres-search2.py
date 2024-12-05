import re

letters = 0
matrix = []
with open("input.txt") as file:
    for line in file:
        letters += len(line.rstrip())
        matrix.append(line.rstrip())

def checkIfImpossible(row, col):
    if(row+1 > (len(matrix)-1) or row-1 < 0 or col+1 > (len(matrix[0])-1) or col-1 < 0):
        return True 
    return False

def checkForMAS(row, col):
    if(checkIfImpossible(row, col)):
        return False
    if((((matrix[row-1][col-1] == 'M' and matrix[row+1][col+1] == 'S') 
         or (matrix[row-1][col-1] == 'S' and matrix[row+1][col+1] == 'M')) 
        and ((matrix[row-1][col+1] == 'M' and matrix[row+1][col-1] == 'S' ) 
        or (matrix[row-1][col+1] == 'S' and matrix[row+1][col-1] == 'M')))):
        return True


totalXMASes = 0
for row in range(len(matrix)):
    for col in range(len(matrix[0])):
        if matrix[row][col] == 'A': 
            if checkForMAS(row, col):
                totalXMASes += 1
            
print(totalXMASes)
print(letters)

