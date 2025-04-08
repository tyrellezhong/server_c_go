#!/usr/bin/env python3
#coding=utf-8

# 定义一个简单的装饰器，不带参数 等同于 say_hello = simple_decorator(say_hello)
def simple_decorator(func):
    def wrapper():
        print("装饰器：函数调用之前")
        func()
        print("装饰器：函数调用之后")
    return wrapper

# 使用装饰器修饰函数
@simple_decorator
def say_hello():
    print("Hello, World! 1---")

# 定义一个带参数的装饰器
def arg_decorator(args):
    def decorator(func):
        def wrapper(*args_, **kwargs):
            print("装饰器：函数调用之前")
            result = func(*args_, **kwargs)  # Handle function arguments and return value
            print("args:", args)
            print("装饰器：函数调用之后")
            return result
        return wrapper
    return decorator

# 使用装饰器修饰函数，等价于say_hello2 = arg_decorator(args="test")(say_hello)
@arg_decorator(
    args="test"
)
def say_hello2():
    print("Hello, World! 2----")

# 调用修饰后的函数
say_hello()

say_hello2()

# 装饰器的工作机制：
# 当 Python 解释器遇到 @decorator 语法时
# 会 立即执行 decorator 函数（在模块导入/代码加载时就会执行）
# 将被装饰的函数作为参数传递给装饰器
# 用装饰器返回的结果 替换 原函数
def decorator(func):
    print("装饰器执行了！")  # 会在 @ 时立即打印
    return func

@decorator  # 这里会立即输出"装饰器执行了！"
def test(): pass
