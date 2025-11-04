// C++类的测试用例，测试类的拷贝构造函数，拷贝赋值运算符，移动构造函数，移动赋值运算符，析构函数
#include <unistd.h>
#include <iostream>
#include <string>

class TestClass {
public:
    std::string name;
    int* data{nullptr};

    // 构造函数
    TestClass(const std::string& n, int value) : name(n), data(new int(value)) {
        std::cout << "调用构造函数 for " << name << std::endl;
    }

    // 析构函数
    ~TestClass() {
        std::cout << "调用析构函数 for " << name << std::endl;
        // delete操作在data为nullptr时是安全的，无需判空
        delete data;
    }

    // 拷贝构造函数
    TestClass(const TestClass& other) : name(other.name + "_copy"), data(new int(*other.data)) {
        std::cout << "调用拷贝构造函数 for " << name << std::endl;
    }

    // 拷贝赋值运算符
    TestClass& operator=(const TestClass& other) {
        if (this != &other) {
            name = other.name + "_copy_assigned";
            delete data;
            data = new int(*other.data);
            std::cout << "调用拷贝赋值运算符 for " << name << std::endl;
        }
        return *this;
    }

    // 移动构造函数
    TestClass(TestClass&& other) noexcept : name(std::move(other.name)), data(other.data) {
        other.data = nullptr;
        std::cout << "调用移动构造函数 for " << name << std::endl;
    }

    // 移动赋值运算符
    TestClass& operator=(TestClass&& other) noexcept {
        if (this != &other) {
            name = std::move(other.name) + "_move_assigned";
            delete data;
            data = other.data;
            other.data = nullptr;
            std::cout << "调用移动赋值运算符 for " << name << std::endl;
        }
        return *this;
    }
};

// 测试函数
void TestClassFunctions() {
    TestClass obj1("Original", 42);
    TestClass obj2 = obj1; // 调用拷贝构造函数
    TestClass obj3("Temporary", 100);
    obj3 = obj1; // 调用拷贝赋值运算符
    TestClass obj4 = std::move(obj1); // 调用移动构造函数
    TestClass obj5("Another", 200);
    obj5 = std::move(obj4); // 调用移动赋值运算符
    int* test = nullptr;
    sleep(6456464);
    delete test; // 测试delete nullptr的安全性
}
