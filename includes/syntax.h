#include <iostream>
#include <ostream>
#include <utility>
#include <vector>
#include <initializer_list>
#include <iostream>

/*
* 可变参数
*/

// c 语言风格可变参数
extern int sum_c_macro(int n, ...);

// C++11 变参模版
// 基本案例：当只有一个参数时，直接返回该参数
template<typename T>
T sum_cpp11_template(T t) {
    return t;
}

// 递归展开：接受一个参数和一个参数包，处理一个参数后，将剩余的参数包传递给下一个递归调用
template<typename T, typename... Args>
T sum_cpp11_template(T first, Args... args) {
    return first + sum_cpp11_template(args...);
}

// C++17折叠表达式
/*
左折叠 ((... op args))：如果参数包为 {1, 2, 3}，结果为 ((1 + 2) + 3)。左折叠的应用场景如逻辑运算 AND 或 OR 操作，可以确保从左到右的短路评估。
右折叠 ((args op ...))：如果参数包为 {1, 2, 3}，结果为 (1 + (2 + 3))。右折叠的应用场景如函数组合，从右至左组合函数更自然，因为这符合数学中的复合函数（g(f(x))）顺序。
语法：
一元：(... op pack) 或 (pack op ...)
二元（带初值）：(init op ... op pack) 或 (pack op ... op init)
*/
/*
变参的 sum 函数使用折叠表达式实现
展开示例：
template<>
int sum<int, int, int, int, int>(int __args0, int __args1, int __args2, int __args3, int __args4)
{
  return __args0 + (__args1 + (__args2 + (__args3 + __args4)));
}
*/
template<typename... Args>
auto sum_cpp17_folder(Args... args) -> decltype((args + ...)) {
    return (args + ...);
}

// 变参的打印日志函数
template<typename... Args>
void print_cpp17_folder(Args&&... args) {
    (std::cout << ... << args) << '\n';  // C++17 折叠表达式
}


// 模版变长参数展开
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


template<typename... Args>
void print_values(Args... args) {
    // 生成，括号右结合表达式，打了括号，从右到左执行
    // static_cast<void>(std::initializer_list<int>{((std::operator<<(std::cout.operator<<(__args0), " ")) , 0) , (((std::operator<<(std::cout.operator<<(__args1), " ")) , 0) , (((std::operator<<(std::cout.operator<<(__args2), " ")) , 0) , ((std::operator<<(std::cout.operator<<(__args3), " ")) , 0)))});
    (void)std::initializer_list<int>{(((std::cout << args << " "), 0), ...)}; // c++17 折叠表达式
}
template<typename... Args>
void print_values2(Args... args) {
    // 生成，从左到右依次执行，c++11的包展开，要么直接函数式展开包参数，要么利用逗号，初始化列表展开，没有其他操作符
    //  static_cast<void>(std::initializer_list<int>{((std::operator<<(std::cout.operator<<(__args0), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args1), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args2), " ")) , 0), ((std::operator<<(std::cout.operator<<(__args3), " ")) , 0)});
    (void)std::initializer_list<int>{((std::cout << args << " "), 0)...}; // c++11 包展开, 逗号运算符 ，多个表达式依次执行，返回最后一个表达式的值
} 

template <typename... Args>
bool bool_test(Args... args) {
    bool ret = true;
    return ((ret = ret && args), ...);
}
// bool_test(true, false, true);
// 代码生成示例
#ifdef INSIGHTS_USE_TEMPLATE
template<>
bool bool_test<bool, bool, bool>(bool __args0, bool __args1, bool __args2)
{
  bool ret = true;
  return (ret = (ret && __args0)) , ((ret = (ret && __args1)) , (ret = (ret && __args2)));
}
#endif

