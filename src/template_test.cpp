#include "template_test.h"

int ForwardTest() {
    std::string name = "Alice";

    std::cout << "--- Passing lvalue ---" << std::endl;
    Wrapper w1 = makeWrapper2(name); // 我们希望拷贝

    std::cout << "--- Passing rvalue ---" << std::endl;
    Wrapper w2 = makeWrapper2("Bob"); // 我们希望移动，但实际还是拷贝！
    // "Bob" 是右值，但进入 makeWrapper 后，arg 变成了左值，
    // 所以会调用 const std::string& 版本，进行拷贝构造。
    return 0;
}