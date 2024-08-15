#!/usr/bin/env python3.11
#coding=utf-8

import argparse

# 创建一个 ArgumentParser 对象，并提供描述信息
parser = argparse.ArgumentParser(description="arg parse")

# 添加命令行参数
parser.add_argument('--foo', help='foo help')
parser.add_argument('bar', help='bar help')

# 解析命令行参数
args = parser.parse_args()

# 打印解析后的参数
print("foo:", args.foo)
print("bar:", args.bar)