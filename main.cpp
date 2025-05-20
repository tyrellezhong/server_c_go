
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
    test.ChronoTimeTest();
    std::unordered_map<int, int> map;
    auto erase_num = map.erase(1);
    LogInfo("format, ... %zu", erase_num);
    TestProtoOption();
    for (;;){
        // 睡眠10秒
        sleep(10);
    }

}
