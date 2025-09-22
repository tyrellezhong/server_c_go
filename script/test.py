import json
data = json.load(open("tmp.json"))
print(json.dumps(data, indent=4, sort_keys=True))