
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


def check_if_update_correctly_ordered(update):
    for i in range(len(update)):
        if rules_dict.__contains__(update[i]) == False:
            continue 
        for j in range(i):
            print("j: " + str(j) + "i: " + str(i))
            if(update[j] in rules_dict[update[i]]):
                return False
    return True

def get_middle_num(singleUpdate):
    middleNum = int(singleUpdate[int(((len(singleUpdate) - 1) / 2))]) 
    print(middleNum)
    return middleNum

updates_list = []
total_sum_result = 0
for update in updates:
    singleUpdate = update.split(",")
    if check_if_update_correctly_ordered(singleUpdate):
       total_sum_result += get_middle_num(singleUpdate) 


print(total_sum_result)