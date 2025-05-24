
#include "Math.h"
#include "includes/Log.h"
#include <stdio.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <arpa/inet.h>
#include <unistd.h>
#include "Sockets.h"
#include "time_test.h"
#include "algorithm.h"
#include "syntax.h"
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <iostream>
#include <vector>
#include "std_example.h"
#include "proto_test.h"

using namespace std;


int main(int argc, char **argv)
{
    
    LogInfo("begin new program : %s", argv[0]);

    TimeTest test;
    // test.ChronoTimeTest();
    std::unordered_map<int, int> map;
    map.insert({1, 1});
    map.insert({2, 2});
    map.insert({3, 3});
    map.insert({4, 4});
    map.erase(1);

    // TestProtoOption();
    header_static_var = 100;
    header_extern_var = 1000;
    mymath::header_namespace_static_var = 200;
    mymath::namespace_extern_var = 2000;

    printf("header_static_var=%d addr=%p \n \
        header_extern_var=%d addr=%p \n \
        header_namespace_static_var=%d addr=%p \n \
        namespace_extern_var=%d addr=%p \n",
        header_static_var, &header_static_var, // 每个文件的值不同，有自己的副本
        header_extern_var, &header_extern_var,// 每个文件的值相同
        mymath::header_namespace_static_var, &mymath::header_namespace_static_var,// 每个文件的值不同，有自己的副本
        mymath::namespace_extern_var, &mymath::namespace_extern_var  // 每个文件的值相同
    );

    GetHeaderVar();
   
}
