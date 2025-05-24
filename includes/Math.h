#pragma once

#include <cstdio>
template <typename T = int >
bool Greater(const T &v1, const T &v2) {
    return v1 > v2;
}
// 当你在 namespace 中定义一个 static 变量时，这个变量具有内部链接。
// 这意味着该变量的作用域仅限于定义它的翻译单元（通常是一个源文件）。
// 其他翻译单元无法直接访问这个变量，即使它们包含相同的头文件。
// 同理，在头文件中直接定义static变量，其他文件也无法访问
// 通常不会在头文件定义static变量
namespace mymath {
    // static变量定义
    static int header_namespace_static_var = 20;
    // 全局变量定义    
    extern int namespace_extern_var;   
}

// 全局变量定义
extern int header_extern_var;

// 这个变量在头文件中声明为static，所有包含这个文件的文件单元，都会生成一个同名的变量副本，供自己访问
static int header_static_var = 10;

// 获取头文件中定义的变量
extern void GetHeaderVar();