#include <iostream>
#include <ostream>
#include <queue>
#include <unordered_set>
#include <unordered_map>
#include <vector>
#include <set>
#include <map>

class StdContainers {

public:

    void AllTest();

    void VectorTest();

    void StringTest();

    void DequeTest();

    void ListTest();

    void SetTest();

    void MapTest();

    void UnorderedSetTest();

    void UnorderedMapTest();

    void MinPqTest();




private:

    std::set<int> set_;

    std::map<int, int> map_;

    std::unordered_set<int> unordered_set_;

    std::unordered_map<int, int> unordered_map_;

    std::priority_queue<int, std::vector<int>, std::greater<int>> min_pq_;
};

class AllocatorTest {

public:

    void StdAllocatorTest();

    void SelfAllocatorTset();
};

// 自定义分配器
template <typename T>
struct MyAllocator {
    using value_type = T;

    MyAllocator() = default;

    template <typename U>
    MyAllocator(const MyAllocator<U>&) {}

    T* allocate(std::size_t n) {
        std::cout << "Allocating " << n << " elements" << std::endl;
        return static_cast<T*>(::operator new(n * sizeof(T)));
    }

    void deallocate(T* p, std::size_t n) {
        std::cout << "Deallocating " << n << " elements" << std::endl;
        ::operator delete(p);
    }

    template <typename U, typename... Args>
    void construct(U* p, Args&&... args) {
        new (p) U(std::forward<Args>(args)...);
    }

    template <typename U>
    void destroy(U* p) {
        p->~U();
    }
};

template <typename T, typename U>
bool operator==(const MyAllocator<T>&, const MyAllocator<U>&) { return true; }

template <typename T, typename U>
bool operator!=(const MyAllocator<T>&, const MyAllocator<U>&) { return false; }