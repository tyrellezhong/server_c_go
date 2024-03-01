#include <algorithm>
#include <cstdint>
#include <list>
#include <string>
#include <unordered_set>
#include <vector>
#include "includes/Gun.h"
#include "includes/Log.h"
#include "includes/Soldier.h"
#include "includes/Math.h"
#include <stdio.h>
#include <arpa/inet.h>



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
  (void)std::initializer_list<int>{((std::cout << N << " "), 0)...};
}

// int main() {
//   auto r = make_index_seq<2>();
//   print(r);
//   return 0;
// }


template <size_t... N> void print(std::index_sequence<N...>) {
  std::vector<int> res;
  (void)std::initializer_list<int>{
      ((res.push_back(N), std::cout << N << " "), 0)...};
    std::cout << std::endl;
  std::for_each(res.begin(), res.end(), [](int x) {std::cout << x << " ";});
}

void SoliderShoot() {
    Solider sanduo("xusanduo");
    sanduo.AddGun(new Gun("AK47"));
    sanduo.AddBulletToGun(20);
    sanduo.fire();
}
// int main() {
//     LogInfo("start run cmake learn.");
//     SoliderShoot();
//     int a = 10;
//     int b = 5;
//     std::cout << "a > b : " << Greater(a, b) << std::endl;
//     auto t = std::make_index_sequence<10>();
//     print(t);
//     std::vector<int32_t> Numbers{0,1,2,3,4,5};
//     std::string string_test("helloworld");
//     // static_assert(const_nums<101>(100) == 100 * 100);
//     return 0;

// }

int main() {

    // 创建一个ip字符串,点分十进制的IP地址字符串
    char buf[] = "192.168.1.4";
    unsigned int num = 0;

    // 将点分十进制的IP字符串转换成网络字节序的整数
    inet_pton(AF_INET, buf, &num);
    unsigned char * p = (unsigned char *)&num;
    printf("%d %d %d %d\n", *p, *(p+1), *(p+2), *(p+3));


    // 将网络字节序的IP整数转换成点分十进制的IP字符串
    char ip[16] = "";
    const char * str =  inet_ntop(AF_INET, &num, ip, 16);
    printf("str : %s\n", str);
    printf("ip : %s\n", str);
    printf("%d\n", ip == str);

    return 0;
}