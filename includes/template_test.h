#include <iostream>
#include <string>

class Wrapper {
public:
    // 构造函数接受左值引用（拷贝）
    Wrapper(const std::string& s) : data(s) {
        std::cout << "Copied: " << s << std::endl;
    }
    // 构造函数接受右值引用（移动）
    Wrapper(std::string&& s) : data(std::move(s)) {
        std::cout << "Moved: " << data << std::endl;
    }
private:
    std::string data;
};

// 工厂函数模板 - 使用通用引用 (T&&)
template<typename T>
Wrapper makeWrapper(T&& arg) {
    // 问题：在函数内部，arg 是一个有名字的变量，所以它永远是左值！
    // 即使外部传入的是一个右值，到了这里也会变成左值。
    return Wrapper(arg); // 错误！总是调用左值版本的构造函数
}

// 正确的工厂函数模板
template<typename T>
Wrapper makeWrapper2(T&& arg) { // T&& 是通用引用，可以绑定到左值或右值
    // 使用 std::forward<T> 来有条件地转换回原始的值类别
    return Wrapper(std::forward<T>(arg));
}

extern int ForwardTest();

