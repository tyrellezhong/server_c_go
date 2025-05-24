#include "Math.h"

// 定义extern变量
int header_extern_var = 101;
int mymath::namespace_extern_var = 102;

void GetHeaderVar(){
    printf("header_static_var=%d addr=%p \n \
        header_extern_var=%d addr=%p \n \
        header_namespace_static_var=%d addr=%p \n \
        namespace_extern_var=%d addr=%p \n",
        header_static_var, &header_static_var, // 每个文件的值不同，有自己的副本
        header_extern_var, &header_extern_var,// 每个文件的值相同
        mymath::header_namespace_static_var, &mymath::header_namespace_static_var,// 每个文件的值不同，有自己的副本
        mymath::namespace_extern_var, &mymath::namespace_extern_var  // 每个文件的值相同
    );
}