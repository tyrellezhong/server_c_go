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