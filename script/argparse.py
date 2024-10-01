#!/usr/bin/env python3.11
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