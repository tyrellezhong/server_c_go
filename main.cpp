
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
#include "class_construction.h"

using namespace std;


int main(int argc, char **argv)
{
    
    LogInfo("begin new program : %s", argv[0]);

    printf("test result begin: \n");
    // std::cout << bool_test(true, false, true);
    printf("\ntest result end \n");
    TestClassFunctions();
   
}
