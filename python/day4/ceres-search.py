matrix = []
with open("input.txt") as file:
    for line in file:
        matrix.append(line.rstrip())

def checkIfLeftUpDiagonalIsImpossible(row, col):
    row = row - 3
    col = col - 3
    if(row < 0 or col < 0):
        return True

def checkLeftUpDiagonal(row, col):
    if(checkIfLeftUpDiagonalIsImpossible(row, col)):
        return False 
    if(matrix[row-1][col-1] == 'M'):
        if(matrix[row-2][col-2] == 'A'):
            if(matrix[row-3][col-3] == 'S'):
                return True
    return False



def checkIfUpIsImpossilble(row):
    row = row - 3
    if(row < 0):
        return True

def checkUp(row, col):
    if(checkIfUpIsImpossilble(row)):
        return False 
    if(matrix[row-1][col] == 'M'):
        if(matrix[row-2][col] == 'A'):
            if(matrix[row-3][col] == 'S'):
                return True
    return False



def checkIfRightUpIsImpossible(row, col):
    row = row - 3
    col = col + 3
    if(row < 0 or col > (len(matrix[0]) - 1)):
        return True

def checkRightUpDiagonal(row, col):
    if(checkIfRightUpIsImpossible(row, col)):
        return False 
    if(matrix[row-1][col+1] == 'M'):
        if(matrix[row-2][col+2] == 'A'):
            if(matrix[row-3][col+3] == 'S'):
                return True
    return False


def checkIfLeftIsImpossible(col):
    col = col - 3
    if(col < 0):
        return True

def checkLeft(row, col):
    if(checkIfLeftIsImpossible(col)):
        return False 
    if(matrix[row][col-1] == 'M'):
        if(matrix[row][col-2] == 'A'):
            if(matrix[row][col-3] == 'S'):
                return True
    return False


def checkIfRightIsImpossible(col):
    col = col + 3
    if(col > (len(matrix[0]) - 1)):
        return True

def checkRight(row, col):
    if(checkIfRightIsImpossible(col)):
        return False 
    if(matrix[row][col+1] == 'M'):
        if(matrix[row][col+2] == 'A'):
            if(matrix[row][col+3] == 'S'):
                return True
    return False

def checkIfLeftDownIsImpossible(row, col):
    col = col - 3 
    row = row + 3
    if(col < 0 or row > (len(matrix) - 1)):
        return True

def checkLeftDownDiagonal(row, col):
    if(checkIfLeftDownIsImpossible(row, col)):
        return False 
    if(matrix[row+1][col-1] == 'M'):
        if(matrix[row+2][col-2] == 'A'):
            if(matrix[row+3][col-3] == 'S'):
                return True
    return False

def checkIfDownIsImpossible(row):
    row = row + 3
    if(row > (len(matrix) - 1)):
        return True
        
def checkDown(row, col):
    if(checkIfDownIsImpossible(row)):
        return False 
    if(matrix[row+1][col] == 'M'):
        if(matrix[row+2][col] == 'A'):
            if(matrix[row+3][col] == 'S'):
                return True
    return False

def checkIfRightDownIsImpossible(row, col):
    row = row + 3 
    col = col + 3 
    if(row > (len(matrix)-1) or col > (len(matrix[0]) - 1)):
        return True

def checkRightDownDiagonal(row, col):
    if(checkIfRightDownIsImpossible(row, col)):
        return False 
    if(matrix[row+1][col+1] == 'M'):
        if(matrix[row+2][col+2] == 'A'):
            if(matrix[row+3][col+3] == 'S'):
                return True
    return False


totalXMASes = 0
for row in range(len(matrix)):
    for col in range(len(matrix[0])):
        if matrix[row][col] == 'X': 
            if checkLeftUpDiagonal(row, col):
                print(matrix[row][col] + matrix[row-1][col-1] + matrix[row-2][col-2] + matrix[row-3][col-3])
                totalXMASes += 1
            if checkUp(row, col):
                print(matrix[row][col] + matrix[row-1][col] + matrix[row-2][col] + matrix[row-3][col])
                totalXMASes += 1
            if checkRightUpDiagonal(row, col):
                print(matrix[row][col] + matrix[row-1][col+1] + matrix[row-2][col+2] + matrix[row-3][col+3])
                totalXMASes += 1
            if checkLeft(row, col):
                print(matrix[row][col] + matrix[row][col-1] + matrix[row][col-2] + matrix[row][col-3])
                totalXMASes += 1
            if checkRight(row, col):
                print(matrix[row][col] + matrix[row][col+1] + matrix[row][col+2] + matrix[row][col+3])
                totalXMASes += 1
            if checkLeftDownDiagonal(row, col):
                print(matrix[row][col] + matrix[row+1][col-1] + matrix[row+2][col-2] + matrix[row+3][col-3])
                totalXMASes += 1
            if checkDown(row, col): 
                print(matrix[row][col] + matrix[row+1][col] + matrix[row+2][col] + matrix[row+3][col])
                totalXMASes += 1
            if checkRightDownDiagonal(row, col):
                print(matrix[row][col] + matrix[row+1][col+1] + matrix[row+2][col+2] + matrix[row+3][col+3])
                totalXMASes += 1

print(totalXMASes)