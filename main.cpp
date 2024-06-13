
#include "includes/Log.h"
#include <stdio.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <arpa/inet.h>
#include <unistd.h>
#include "Sockets.h"
#include "mychrono.h"
#include "algorithm.h"
#include "syntax.h"
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

using namespace std;


int main(int argc, char **argv)
{
    LogInfo("begin new program : %s", argv[0]);
    // TCPServer();
    // UDPServer();
    // TCPClient();
    //  UDPClient();
    // char cwd[100];
    // getcwd(cwd, sizeof(cwd));
    // CStudent s;
    // ofstream outFile;
    // outFile.open("students.dat", ios::out | ios::binary | iostream::trunc);
    // if (!outFile) {
    //     std::cout << "error occur" << std::endl;
    // }
    // while (cin >> s.szName >> s.age)
    //     outFile.write((char*)&s, sizeof(s));
    // outFile.close();
    // std::cout << "write finish" << cwd << std::endl ;

    // std::vector arr{1, 2, 3, 4};
    // for (int& val : arr) {
    //     val = val + 10;
    // }
    // printf("%d\n", sizeof(int));
    // return 0;

    // std::cout << "sum:" << Sum(5, 1,1,1,1,1) << std::endl;
    PrintN(make_index_seq<5>());
}
