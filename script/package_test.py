#!/usr/bin/env python3
#coding=utf-8

import getopt
import os
import sys

# 闭包测试
def make_accumulator():
    total = 0  # 这是外部函数的局部变量

    def accumulator(value):
        nonlocal total  # 声明使用外部函数的局部变量
        total += value
        return total

    return accumulator


class TestClass:
    name = "testclass"

if __name__ == "__main__":

    namespace = ""
    processname = ""
    yaml_name = ""
    local_port = ""
    remote_port = ""
    destination_ip = ""
    # 解析参数
    # "hi:o:": 短格式分析串, h 后面没有冒号, 表示后面不带参数; i 和 o 后面带有冒号, 表示后面带参数
    # ["help", "input_file=", "output_file="]: 长格式分析串列表, help后面没有等号, 表示后面不带参数; input_file和output_file后面带冒号, 表示后面带参数
    # 返回值包括 `opts` 和 `args`, opts 是以元组为元素的列表, 每个元组的形式为: (选项, 附加参数)，如: ('-i', 'test.png');
    # args是个列表，其中的元素是那些不含'-'或'--'的参数
    opts, args = getopt.getopt(sys.argv[1:], "hn:p:f:l:r:i:",
                               ["help", "namespace=", "processname=", "yamlname=", "localport=", "remoteport=", "desip="])
    for opt, arg in opts:
        if opt in ("h", "--help"):
            help()
            sys.exit()
        elif opt in ("-n", "--namespace"):
            namespace = arg
        elif opt in ("-p", "--processname"):
            processname = arg
        elif opt in ("-f", "--yamlname"):
            yaml_name = arg
        elif opt in ("-l", "--localport"):
            local_port = arg
        elif opt in ("r", "--remoteport"):
            remote_port = arg
        elif opt in ("i", "--desip"):
            destination_ip = arg

    test_class = TestClass()
    print(test_class)
    print(test_class.__dict__)
    print(__file__)
    print(os.getcwd())
    print(os.chdir(os.path.dirname(__file__)))
    print(os.path.basename(__file__))
    print(os.getcwd())
    print(os.path.exists(os.getcwd() + "/file"))
    print(os.path.isfile(os.getcwd() + "/file"))
    print(os.listdir(os.getcwd()))
    li = [1, 2, 3, 4, 5]
    print([ item * 2 for item in li if item % 2 == 0])

    # 使用累加器
    # 创建一个累加器实例
    acc = make_accumulator()
    print(acc(10))  # 输出: 10
    print(acc(5))   # 输出: 15
    print(acc(3))   # 输出: 18






