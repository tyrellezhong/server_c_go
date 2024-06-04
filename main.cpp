#include <algorithm>
#include <chrono>
#include <cstdint>
#include <ctime>
#include <istream>
#include <sstream>
#include <string>
#include <vector>
#include "includes/Gun.h"
#include "includes/Log.h"
#include "includes/Soldier.h"
#include "includes/Math.h"
#include <stdio.h>
#include <arpa/inet.h>
#include <stdio.h>
#include <arpa/inet.h>
#include <unistd.h>
#include "Sockets.h"
#include "src/proto_test.pb.h"
#include "mychrono.h"
#include "algorithm.h"
#include "sstream"
#include "syntax.h"

template <int... N>
struct index_seq {};

template <int N, int... M>
struct make_index_seq : public make_index_seq<N - 1, N - 1, M...> {};
template <int... M>
struct make_index_seq<0, M...> : public index_seq<M...> {};


template <int... N> void print(index_seq<N...>) {
    // (void)std::initializer_list<int>{((std::cout << N << " "), 0)...};
    ((std::cout << N << " "), ...);
    std::cout << std::endl;
}

template <size_t... N> void print(std::index_sequence<N...>) {
    std::vector<int> res;
    auto xx = std::initializer_list<int>{((res.push_back(N), std::cout << N << " ", 10))...};
    std::cout << "list size:" << xx.size() << std::endl;
    // for (auto iter = xx.begin(); iter != xx.end(); ++iter)
    // {
    //     std::cout << *iter << " ";
    // }
    // std::cout << std::endl;
    // std::cout << "For each : ";
    // std::for_each(res.begin(), res.end(), [](int x) {std::cout << x << " ";});
    // std::cout << std::endl;
}

void SoliderShoot() {
    Solider sanduo("xusanduo");
    sanduo.AddGun(new Gun("AK47"));
    sanduo.AddBulletToGun(20);
    sanduo.fire();
}
using namespace std;
class CStudent
{
public:
    char szName[20];
    int age;
};
int main()
{
    LogInfo("new program begin !!!")
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
    std::cout << "sum:" << sum(5, 1,1,1,1,1) << std::endl;
}
