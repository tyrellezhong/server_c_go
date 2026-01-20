

#include <vector>
#include <unordered_map>


// 二维向量节点
struct Vec2D {
    int x, y;
    Vec2D(int x = 0, int y = 0) : x(x), y(y) {}
    
    // 用于unordered_map的键
    bool operator==(const Vec2D& other) const {
        return x == other.x && y == other.y;
    }
};

// 哈希函数用于Vec2D
struct Vec2DHash {
    size_t operator()(const Vec2D& v) const {
        return std::hash<int>()(v.x) ^ (std::hash<int>()(v.y) << 1);
    }
};


int TopoTest();

