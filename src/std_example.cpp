#include "std_example.h"
#include <algorithm>
#include <array>
#include <cstddef>
#include <deque>
#include <forward_list>
#include <functional>
#include <initializer_list>
#include <queue>
#include <stack>
#include <unordered_map>
#include <utility>
#include <vector>
#include "Log.h"
#include <list>


#define  kBeginFlag  "---------------------------------%s-------------------------------"

void StdContainers::AllTest() {

    VectorTest();
    StringTest();
    DequeTest();
    ListTest();
    UnorderedSetTest();
    UnorderedMapTest();

    SetTest();
    MapTest();
    MinPqTest();

}

void StdContainers::VectorTest() {
    LogInfo(kBeginFlag, "VectorTest");
    std::vector<int> vector(5);
    vector.insert(vector.begin(), 5, 5);
    vector.assign(10, 0);
    std::initializer_list<int> xx{1, 2, 3, 4, 3, 11, 8, 0, 10};
    vector = xx;
    vector.erase(vector.begin());
    vector.pop_back();
    vector.push_back(9);
    std::sort(vector.begin(), vector.end(), std::greater<int>());

    PrintElements(vector);
}

void StdContainers::StringTest() {
    LogInfo(kBeginFlag, "StringTest");
    std::string str(5, 'a');
    std::string tmp("abcdefgh");
    const char* tmp2 = "abcdefgh";
    std::string str2(tmp2, 5);
    std::string str3(tmp, 5);
    auto sub = tmp.substr(0, 5);
    str.insert(0, "---");
    str.append(str2);
    str.insert(0, str2, 0, 5);
    str.assign(tmp2);
    auto ret1 = str.find('c', 5);
    auto ret2 = str.find("abc");
    auto ret3 = str.find_first_of("dcb");

    std::string ch = "中文测试";
    std::cout << "ch :  " << ch << "len" << ch.size() << std::endl;

    std::string x(5, 'a');
    std::string str1 = "123";
    std::string str2x = "456abc";
    std::string str3x = "0x1A"; // 16 进制
    std::string str4 = "075";  // 8 进制
    std::string str5 = "invalid";
    std::sort(str5.begin(), str5.end());


    try {
        int num1 = std::stoi(str1);
        std::cout << "num1: " << num1 << std::endl;

        std::size_t pos;
        int num2 = std::stoi(str2x, &pos);
        std::cout << "num2: " << num2 << ", first non-converted character at position: " << pos << std::endl;

        int num3 = std::stoi(str3x, nullptr, 16);
        std::cout << "num3: " << num3 << std::endl;

        int num4 = std::stoi(str4, nullptr, 8);
        std::cout << "num4: " << num4 << std::endl;

        int num5 = std::stoi(str5);
        std::cout << "num5: " << num5 << std::endl; // This line will not be executed
    } catch (const std::invalid_argument& e) {
        std::cerr << "Invalid argument: " << e.what() << std::endl;
    } catch (const std::out_of_range& e) {
        std::cerr << "Out of range: " << e.what() << std::endl;
    }


}

void StdContainers::DequeTest() {
    LogInfo(kBeginFlag, "DequeTest");
    std::deque<int> deque;
    deque.push_front(1);
    deque.push_back(2);

    std::stack<int> stack;

    std::queue<int> queue;

}

bool customCompare(int a, int b) {
    return a > b; // 按降序合并
}

void StdContainers::ListTest() {
    LogInfo(kBeginFlag, "ListTest");

    std::list<int> list1 = {7, 5, 3, 1};
    std::list<int> list2 = {8, 6, 4, 2};

    list1.merge(list2, customCompare);

    for (const auto& value : list1) {
        std::cout << value << " ";
    }
    std::cout << std::endl;

    // list2 现在为空
    std::cout << "Size of list2: " << list2.size() << std::endl;

    std::forward_list<int> for_list;
    for_list.assign({1, 2, 3, 0, 9, 8, 7});
    for_list.sort();
    PrintElements(for_list);

}

void StdContainers::UnorderedSetTest() {
    LogInfo(kBeginFlag, "UnorderedSetTest");
    auto ret1 = unordered_set_.insert(1);
    LogInfo("set insert %d", *ret1.first);
    auto ret2 = unordered_set_.emplace(2);
    unordered_set_.insert({3, 4, 5});
    std::vector<int> li{6, 7};
    unordered_set_.insert(li.begin(), li.end());
    auto ret3 = unordered_set_.erase(1);
    auto node = unordered_set_.extract(100);
    if (node) {
        std::cout << "node value : " << node.value() << std::endl;
    }
    if (unordered_set_.find(2) != unordered_set_.begin()) {
        std::cout << "find key success " << std::endl;
    }
    auto ret4 = unordered_set_.insert(std::move(node));
    std::cout << "cur_bc:" << unordered_set_.bucket_count() << " maxbc:" << unordered_set_.max_bucket_count() << std::endl;
    std::cout << "local factor:" << unordered_set_.load_factor() << std::endl;
    PrintElements(unordered_set_);
}

// 自定义哈希函数
template <typename T, std::size_t N>
struct ArrayHash {
    std::size_t operator()(const std::array<T, N>& arr) const {
        std::size_t hash = 0;
        for (const auto& elem : arr) {
            hash ^= std::hash<T>()(elem) + 0x9e3779b9 + (hash << 6) + (hash >> 2);
        }
        return hash;
    }
};

// 自定义相等性比较函数
template <typename T, std::size_t N>
struct ArrayEqual {
    bool operator()(const std::array<T, N>& lhs, const std::array<T, N>& rhs) const {
        return lhs == rhs;
    }
};

void StdContainers::UnorderedMapTest() {
    LogInfo(kBeginFlag, "UnorderedMapTest");
    auto ret1 = unordered_map_.insert(std::make_pair(1, 1));
    LogInfo("map insert %d:%d", (*ret1.first).first, (*ret1.first).second);
    auto ret2 = unordered_map_.emplace(2, 2);
    auto ret3 = unordered_map_.insert({3, 3});
    LogInfo("map insert %d:%d", (*ret2.first).first, (*ret2.first).second);
    unordered_map_[4] = 4;
    auto ret4 = unordered_map_.erase(5);
    if (unordered_map_.find(4) != unordered_map_.end()) {
        LogInfo("find key %d success", 4);
    }
    // PrintElements(unordered_map_);
    for (auto [key, value] : unordered_map_) {
        LogDebug("key:%d value:%d", key, value);
    }

    auto hash_key = [](const std::array<char, 26>& arr) {
        std::size_t hash = 0;
        for (const char& item : arr) {
            hash ^= std::hash<char>()(item) + 0x9e3779b9 + (hash << 6) + (hash >> 2);
        }
        return hash;
    };
    auto hash_equal = [](const std::array<char, 26>& la, const std::array<char, 26>& ra) {
        return la == ra;
    };

    std::unordered_map<std::array<char, 26>, std::vector<std::string>, decltype(hash_key), decltype(hash_equal)> dic(10, hash_key, hash_equal);

    // 使用 std::array 作为键值的 unordered_map
    std::unordered_map<std::array<char, 26>, std::string, ArrayHash<int, 3>, ArrayEqual<int, 3>> myMap;

    // // 插入元素
    // myMap.emplace(std::array<int, 3>{1, 2, 3}, "First");
    // myMap.emplace(std::array<int, 3>{4, 5, 6}, "Second");

    // // 查找元素
    // std::array<int, 3> keyToFind = {1, 2, 3};
    // auto it = myMap.find(keyToFind);
    // if (it != myMap.end()) {
    //     std::cout << "Found: " << it->second << std::endl;
    // } else {
    //     std::cout << "Not found" << std::endl;
    // }

}

void StdContainers::MinPqTest() {
    LogInfo(kBeginFlag, "MinPqTest");
    // 向优先队列中插入元素
    min_pq_.push(10);
    min_pq_.push(30);
    min_pq_.push(20);
    min_pq_.push(5);

    // 输出并移除优先队列中的元素
    while (!min_pq_.empty()) {
        std::cout << min_pq_.top() << " "; // 输出队列顶部元素
        min_pq_.pop(); // 移除队列顶部元素
    }
    std::cout << std::endl;
}

void StdContainers::MapTest() {
    LogInfo(kBeginFlag, "MapTest");
    map_.insert(std::make_pair(1, 1));
    map_.insert({2, 2});
    map_.insert({{3, 3}, {4, 4}});
    map_.emplace(5, 5);
    map_[6] = 6;

    for (auto& [key, value] : map_) {
        std::cout << key << " : " << value << std::endl;
    }
}

void StdContainers::SetTest() {

}

void AllocatorTest::StdAllocatorTest() {
    // 创建一个 std::allocator 对象，用于分配 int 类型的内存
    std::allocator<int> allocator;

    // 分配内存，足够存储 5 个 int 对象
    int* p = allocator.allocate(5);

    // 在分配的内存上构造对象
    for (int i = 0; i < 5; ++i) {
        allocator.construct(p + i, i * 10);  // 构造值为 i * 10 的 int 对象
    }

    // 使用分配的内存
    for (int i = 0; i < 5; ++i) {
        std::cout << p[i] << " ";  // 输出：0 10 20 30 40
    }
    std::cout << std::endl;

    // 销毁对象
    for (int i = 0; i < 5; ++i) {
        allocator.destroy(p + i);
    }

    // 释放内存
    allocator.deallocate(p, 5);
}

void AllocatorTest::SelfAllocatorTset() {
    // 使用自定义分配器的 std::vector
    std::vector<int, MyAllocator<int>> myVector;

    // 插入元素
    myVector.push_back(10);
    myVector.push_back(20);
    myVector.push_back(30);

    // 输出元素
    for (const auto& elem : myVector) {
        std::cout << elem << " ";  // 输出：10 20 30
    }
    std::cout << std::endl;

}