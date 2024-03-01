# cmake 最小版本要求
cmake_minimum_required(VERSION 2.8)

# 项目名称
project(SOLDIERFIRE)

#clangd 配置
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# 在CMAKE_CXX_FLAGS编译选项后追加-std=c++11等
set( CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++14 -g -Wall")

set(CMAKE_CXX_STANDARD 14)


# 设定编译类型为debug，调试时需要选择debug
set(CMAKE_BUILD_TYPE Debug)

# # 设定编译类型为release，发布时需要选择release
# set(CMAKE_BUILD_TYPE Release)

# 向工程添加多个特定的头文件搜索路径
include_directories(${CMAKE_SOURCE_DIR}/includes)

# 编译SRC变量所代表的源代码文件，生成main可执行文件
add_executable(myexe main.cpp src/Gun.cpp src/soldier.cpp src/Log.cpp src/Sockets.cpp)