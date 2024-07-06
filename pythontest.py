
li = [3, 2, 1, 7, 9, 8, 4, 5, 6]

li.sort()

li.pop()

li.append(10)
li.pop()
li[0] = 111


li.insert(0, [12, 13])

li.pop()


map = dict()
map["1"] = 1
map["2"] = 2
map["4"] = 2
map["5"] = 5
map.setdefault("6", 6)
map.pop("1")
if "5" in map:
    print(map["5"])

map.popitem()

print(li)
print(map)
print(map.keys())
print(map.values())
print(map.items())
map.copy()

s = {1, 2, 3, 8, 5, 7, 0}
s.copy()
s.add("s")
s.remove(1)
if "s" in s:
    print("set find")

