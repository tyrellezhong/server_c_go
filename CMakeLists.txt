# cmake 最小版本要求
cmake_minimum_required(VERSION 3.2)

# 项目名称
project(SOLDIERFIRE)

set(app myexe)

#clangd 配置
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

# 在CMAKE_CXX_FLAGS编译选项后追加-std=c++11等
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17 -g -Wall")

set(CMAKE_CXX_STANDARD 17)

# 显示编译器
message(STATUS "C Compiler: ${CMAKE_C_COMPILER}")
message(STATUS "C++ Compiler: ${CMAKE_CXX_COMPILER}")

# 设定编译类型为debug，调试时需要选择debug
set(CMAKE_BUILD_TYPE Debug)

# # 设定编译类型为release，发布时需要选择release
# set(CMAKE_BUILD_TYPE Release)

find_package(Protobuf REQUIRED) # proto 编译不过时可以注释掉
message(STATUS "protobuf package: ${Protobuf_LIBRARIES}")

# 向工程添加多个特定的头文件搜索路径
include_directories(${CMAKE_SOURCE_DIR}/includes)
include_directories(${CMAKE_SOURCE_DIR}/proto_out)
include_directories(${CMAKE_SOURCE_DIR}/protobuf-21.9/include) # proto 编译不过时可以注释掉

file(GLOB_RECURSE srcs
        ./src/*.c
        ./src/*.cpp
        ./src/*.h
        ./proto_out/*.cc # proto 编译不过时可以注释掉
        ./proto_out/*.cpp # proto 编译不过时可以注释掉
        )

# 编译SRC变量所代表的源代码文件，生成main可执行文件
add_executable(${app} main.cpp ${srcs})
target_link_libraries(${app} ${Protobuf_LIBRARIES}) # proto 编译不过时可以注释掉

add_subdirectory(gomod)