#include <algorithm>
#include <cstdint>
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



// template <size_t ...N>
// static constexpr auto square_nums(size_t index, std::index_sequence<N...>) {
//     constexpr auto nums = std::array{(N * N)...};
//     return nums[index];
// }

// template <size_t N>
// constexpr static auto const_nums(size_t index) {
//     return square_nums(index, std::make_index_sequence<N>{});
// }
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
  for (auto iter = xx.begin(); iter != xx.end(); ++iter)
  {
    std::cout << *iter << " ";
  }
  std::cout << std::endl;
  std::cout << "For each : ";
  std::for_each(res.begin(), res.end(), [](int x) {std::cout << x << " ";});
  std::cout << std::endl;
}

void SoliderShoot() {
    Solider sanduo("xusanduo");
    sanduo.AddGun(new Gun("AK47"));
    sanduo.AddBulletToGun(20);
    sanduo.fire();
}
int main() {
    LogInfo("start tcp client.");
    // int a = 10;
    // int b = 5;
    // std::cout << "a > b : " << Greater(a, b) << std::endl;
    auto t = std::make_index_sequence<10>();
    // print(t);
    auto t2 = make_index_seq<5>();
    print(t2);
;
    // std::vector<int32_t> Numbers{0,1,2,3,4,5};
    // std::string string_test("helloworld");
    // // static_assert(const_nums<101>(100) == 100 * 100);

    // ByteOrderHostTest();
    // ByteOrderTransform();
    // TCPServer();
    // UDPServer();
    // TCPClient();
    //  UDPClient();
    tutorial::Person Person;
    auto Descripter = Person.GetDescriptor();
    return 0;

}

