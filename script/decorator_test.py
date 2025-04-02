#!/usr/bin/env python3
#coding=utf-8

# 定义一个简单的装饰器
def simple_decorator(func):
    def wrapper():
        print("装饰器：函数调用之前")
        func()
        print("装饰器：函数调用之后")
    return wrapper

# 使用装饰器修饰函数
@simple_decorator
def say_hello():
    print("Hello, World!")

# 调用修饰后的函数
say_hello()