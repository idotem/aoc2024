rules = []
updates = []
flag = False
with open("input.txt") as file:
    for line in file:
        if(line == '\n'):
            flag = True
            continue
        if(flag):
            updates.append(line.rstrip())
            continue
        rules.append(line.rstrip())

rules_dict = {}
for rule in rules:
    r = rule.split("|")
    rule_set = rules_dict.get(r[0])
    if(rule_set == None):
        rules_dict[r[0]] = []
    rules_dict[r[0]].append(r[1])

class Index:
  def __init__(self, ruleIndex, updateIndex):
    self.ruleIndex = ruleIndex
    self.updateIndex =updateIndex 

def check_if_update_correctly_ordered(update):
    for i in range(len(update)):
        if rules_dict.__contains__(update[i]) == False:
            continue 
        for j in range(i):
            if(update[j] in rules_dict[update[i]]):
                return Index(i, j)
    return Index(-1, -1) 

def fix_update_ordering(update, index):
    check = check_if_update_correctly_ordered(update)
    if(check.updateIndex == -1):
        return update 
    tmp = update[index.updateIndex]
    update[index.updateIndex] = update[index.ruleIndex]
    update[index.ruleIndex] = tmp
    return fix_update_ordering(update, check)
       
def get_middle_num(singleUpdate):
    middleNum = int(singleUpdate[(len(singleUpdate) // 2)]) 
    return middleNum

updates_list = []
total_sum_result = 0
for update in updates:
    singleUpdate = update.split(",")
    checkIndex = check_if_update_correctly_ordered(singleUpdate)
    if  checkIndex.updateIndex != -1:
       total_sum_result += get_middle_num(fix_update_ordering(singleUpdate, checkIndex)) 


print(total_sum_result)