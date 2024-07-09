#!/bin/bash

# 打印所有传入的参数
echo "参数1:$1"
echo "参数2:$2"
echo "所有参数: $@"
for arg in "$@"; do
    echo "参数: $arg"
done