
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
#include <cstdint>
#include <iostream>
#include <vector>
#include "std_example.h"
#include "proto_test.h"
#include "class_construction.h"
#include "template_test.h"

using namespace std;


int main(int argc, char **argv)
{
    
    LogInfo("begin new program : %s", argv[0]);

    printf("test result begin: \n");
    float x = 1.15f;
    int y = static_cast<int>(x);
    printf("float to int: %d \n", y);
    printf("\ntest result end \n");

    ForwardTest();

   
}

