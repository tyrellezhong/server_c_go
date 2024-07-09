
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
#include <vector>
#include "std_example.h"

using namespace std;




int main(int argc, char **argv)
{
    LogInfo("begin new program : %s", argv[0]);

    StdContainers containertest;
    containertest.AllTest();

    TimeTest time_test;
    time_test.ChronoTimeTest();
    time_test.CTimeTest();

    std::cout << "make_index_seq : ";
    PrintN(make_index_seq<5>());
}
