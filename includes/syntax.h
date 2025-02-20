#include <iostream>
#include <ostream>
#include <utility>
#include <vector>
#include <initializer_list>
#include <iostream>


template <int... N>
struct index_seq {};

// make_index_seq<5> 
// 展开：
// 4 4
// 3 3 4 
// 2 2 3 4
// 1 1 2 3 4
// 0 0 1 2 3 4
// 0 1 2 3 4 
// index_seq<0, 1, 2, 3, 4>{}
template <int N, int... M>
struct make_index_seq : public make_index_seq<N - 1, N - 1, M...> {};

// 递归终止条件
template <int... M>
struct make_index_seq<0, M...> : public index_seq<M...> {};


template <int... N> 
void PrintN(index_seq<N...>) {
    (void)std::initializer_list<int>{((std::cout << N << " "), 0)...}; // c++11 包展开, 逗号运算符 ，多个表达式依次执行，返回最后一个表达式的值
    // ((std::cout << N << " "), ...); // c++17 折叠表达式
    std::cout << std::endl;
}

template <size_t... N> void PrintN(std::index_sequence<N...>) {
    std::vector<int> res;
    auto xx = std::initializer_list<int>{((res.push_back(N), std::cout << N << " ", 10))...};
    std::cout << "list size:" << xx.size() << std::endl;
}

// c 语言风格可变参数
extern int Sum(int n, ...);


template<typename... Args>
void print_values(Args... args) {
    // 生成，括号右结合表达式
    // static_cast<void>(std::initializer_list<int>{((std::operator<<(std::cout.operator<<(__args0), " ")) , 0) , (((std::operator<<(std::cout.operator<<(__args1), " ")) , 0) , (((std::operator<<(std::cout.operator<<(__args2), " ")) , 0) , ((std::operator<<(std::cout.operator<<(__args3), " ")) , 0)))});
    (void)std::initializer_list<int>{(((std::cout << args << " "), 0), ...)}; // c++17 折叠表达式
}
template<typename... Args>
void print_values2(Args... args) {
    // 生成
    //  static_cast<void>(std::initializer_list<int>{((std::operator<<(std::cout.operator<<(__args0), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args1), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args2), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args3), " ")) , 0)});
    (void)std::initializer_list<int>{((std::cout << args << " "), 0)...}; // c++11 包展开, 逗号运算符 ，多个表达式依次执行，返回最后一个表达式的值
}  