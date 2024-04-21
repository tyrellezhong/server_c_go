#include <algorithm>
#include <chrono>
#include <cstdint>
#include <ctime>
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
#include "Sockets.h"
#include "src/proto_test.pb.h"
#include "mychrono.h"
#include "algorithm.h"


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
int main() {

    // socket test

    // ByteOrderHostTest();
    // ByteOrderTransform();
    // TCPServer();
    // UDPServer();
    // TCPClient();
    //  UDPClient();

    // auto t = std::make_index_sequence<10>();
    // print(t);
    // auto t2 = make_index_seq<5>();
    // print(t2);

    // // tutorial::Person Person;
    // // auto Descripter = Person.GetDescriptor();


    // hours_type h_oneday (24);                  // 24h
    // seconds_type s_oneday (60*60*24);          // 86400s
    // milliseconds_type ms_oneday(s_oneday);    // 86400000ms

    // seconds_type s_onehour (60*60);            // 3600s
    // hours_type h_onehour (std::chrono::duration_cast<hours_type>(s_onehour));
    // milliseconds_type ms_onehour (s_onehour);  // 3600000ms (ok, no type truncation)

    // std::cout << ms_onehour.count() << "ms in 1h" << std::endl;
    // return 0;
    // using namespace std::chrono;

    // system_clock::time_point tp_epoch;    // epoch value

    // time_point <system_clock,duration<int>> tp_seconds (duration<int>(1000));

    // system_clock::time_point tp (tp_seconds);

    // std::cout << "1 second since system_clock epoch = ";
    // std::cout << tp_seconds.time_since_epoch().count();
    // std::cout << " system_clock periods." << std::endl;

    // // display time_point:
    // std::time_t tt = system_clock::to_time_t(tp);
    // std::cout << "time_point tp is: " << ctime(&tt);

    // std::time_t tt2 = system_clock::to_time_t(system_clock::now());
    // std::cout << "time_point tp is: " << ctime(&tt2);

    // std::vector<int> temp{1,2,3,4,5};
    TestQsort();
    return 0;

}

