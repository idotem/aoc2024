import re

totalRes = 0
with open("input.txt") as f:
    txtInput = f.read().split("do()")
    validInput = ""
    for ti in txtInput:
       validInput += ti.split("don't()")[0]
    print(validInput)
    for m in re.findall(r'mul\([0-9]{1,3},[0-9]{1,3}\)', validInput):
       nums = re.findall(r'[0-9]{1,3}', m)
       totalRes += int(nums[0]) * int(nums[1])

print(totalRes)
