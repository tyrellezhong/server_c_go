#!/usr/bin/env python3
#coding=utf-8

def _get_join_msg(*args):
  return " ".join(map(str, args))

def _get_log_pattern(msg, format="DEBUG"):
  return "[%s]:%s" % (format, msg)

def _log_red(*args):
  errMsg = _get_log_pattern(_get_join_msg(*args), "ERROR")
  print("\033[31m%s\033[0m" % errMsg)

def _log_yellow(*args):
  errMsg = _get_log_pattern(_get_join_msg(*args), "Warn")
  print("\033[33m%s\033[0m" % errMsg)

def _log_green(*args):
  errMsg = _get_log_pattern(_get_join_msg(*args), "INFO")
  print("\033[32m%s\033[0m" % errMsg)

log_info = _log_green
log_warn = _log_yellow
log_error = _log_red

def print_head(head : str):
    log_info("------------------", head, "-----------------")

def string_test():
    print_head("string_test")
    strt = 'Hello, World!'
    str1 = 'Hello, World!'
    strc = "Hello, 世界!"
    print(str1)

    # 使用双引号创建字符串
    str2 = "Hello, World!"
    print(str2)

    # 使用三引号创建多行字符串
    str3 = """Hello,
    World!"""
    print(str3)

    # 使用 + 操作符连接字符串
    result = str1 + ", " + str2 + "!"
    print(result)

    # 使用 join 方法连接字符串
    parts = ["Hello", "World"]
    result = ", ".join(parts)
    print(result)

    # 获取子字符串
    sub_str = str1[7:12]
    print(sub_str)  # 输出: World

    # 查找子字符串
    index = strt.find("World")
    print(index)  # 输出: 7

    # 检查子字符串是否存在
    contains = "World" in strt
    print(contains)  # 输出: True

    # 替换子字符串
    new_str = strt.replace("World", "Python")
    print(new_str)  # 输出: Hello, Python!

    strx = "a,b,c,d,e"

    # 分割字符串
    parts = strx.split(",")
    print(parts)  # 输出: ['a', 'b', 'c', 'd', 'e']

    # 将字符串转换为列表
    str_list = list(strc)

    # 修改字符
    str_list[7] = 'P'
    str_list[8] = 'y'

    # 将列表转换回字符串
    new_str = ''.join(str_list)
    print(new_str)  # 输出: Hello, Pyorld!

    # 使用 for 循环遍历字符串
    for char in strc:
        print(char)

    # 获取字符串长度
    length = len(strc)
    print(length)  # 输出: 10

    # 整数转换为字符串
    int_val = 123
    str_val = str(int_val)
    print(str_val)  # 输出: 123

    # 字符串转换为整数
    str_val = "456"
    int_val = int(str_val)
    print(int_val)  # 输出: 456

    name = "Alice"
    age = 30

    # 使用 % 操作符格式化字符串
    formatted_str = "Name: %s, Age: %d" % (name, age)
    print(formatted_str)  # 输出: Name: Alice, Age: 30

    # 使用 str.format 方法格式化字符串
    formatted_str = "Name: {}, Age: {}".format(name, age)
    print(formatted_str)  # 输出: Name: Alice, Age: 30

    # 使用 f-string 格式化字符串 (Python 3.6+)
    formatted_str = f"Name: {name}, Age: {age}"
    print(formatted_str)  # 输出: Name: Alice, Age: 30

def map_test():
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

    print(map)
    print(map.keys())
    print(map.values())
    print(map.items())
    map.copy()

def list_test():
    li = [3, 2, 1, 7, 9, 8, 4, 5, 6]

    li.sort()

    li.pop()

    li.append(10)
    li.pop()
    li[0] = 111


    li.insert(0, [12, 13])

    li.pop()
    for i in li:
      print(i)
    for i in range(0, 10):
       print(i)

def set_test():
    s = {1, 2, 3, 8, 5, 7, 0}
    s.copy()
    s.add("s")
    s.remove(1)
    if "s" in s:
        print("set find")

if __name__ == "__main__":
   string_test()
   list_test()
   map_test()
   set_test()