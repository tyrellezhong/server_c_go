#!/usr/bin/env python3
#coding=utf-8

import datetime
import getopt
import os
import sys
import time

# 闭包测试
def make_accumulator():
    total = 0  # 这是外部函数的局部变量

    def accumulator(value):
        nonlocal total  # 声明使用外部函数的局部变量
        total += value
        return total

    return accumulator

def time_test():
    timestamp = time.time()
    print("timestamp", timestamp)
    # 获取当前时间并格式化
    utc_time = time.strftime("%Y-%m-%d %H:%M:%S", time.gmtime(timestamp))
    print("UTC 时间:", utc_time)
    formatted_time = time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(timestamp))
    print("格式化后的当前时间:", formatted_time)
    c_time = time.ctime(timestamp)
    print("c_time:", c_time, "\n asctime:", time.asctime())

        # 获取当前日期和时间
    now = datetime.datetime.now()
    print("当前日期和时间:", now)

def os_test():
    print(os.getcwd())
    print(os.chdir(os.path.dirname(__file__)))
    print(os.path.basename(__file__))
    print(os.getcwd())
    print(os.path.exists(os.getcwd() + "/file"))
    print(os.path.isfile(os.getcwd() + "/file"))
    print(os.listdir(os.getcwd()))

def getopt_test():
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


class TestClass:
    name = "testclass"

if __name__ == "__main__":

    # 列表推导
    li = [1, 2, 3, 4, 5]
    print("列表推导：", [ item * 2 for item in li if item % 2 == 0])

    # 使用累加器
    # 创建一个累加器实例
    acc = make_accumulator()
    print("闭包累加器：", acc(10), acc(5), acc(3))  # 输出: 10 15 18
    time_test()








