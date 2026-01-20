

#include <queue>
#include <algorithm>
#include <iostream>

#include "topo_sort.h"



// 计算最大依赖长度，节点是Vec2D
static std::unordered_map<Vec2D, int, Vec2DHash> ComputeMaxDependencyLengthVec2D(
    const std::vector<std::pair<Vec2D, Vec2D>>& dependencies) {
    
    // 第一步：收集所有节点并分配ID
    std::unordered_map<Vec2D, int, Vec2DHash> nodeToId;
    std::vector<Vec2D> idToNode;
    
    for (const auto& dep : dependencies) {
        const Vec2D& u = dep.first;
        const Vec2D& v = dep.second;
        
        if (nodeToId.find(u) == nodeToId.end()) {
            nodeToId[u] = idToNode.size();
            idToNode.push_back(u);
        }
        if (nodeToId.find(v) == nodeToId.end()) {
            nodeToId[v] = idToNode.size();
            idToNode.push_back(v);
        }
    }
    
    int n = idToNode.size();
    
    // 第二步：构建图
    std::vector<std::vector<int>> graph(n);
    std::vector<int> indegree(n, 0);
    
    for (const auto& dep : dependencies) {
        int u = nodeToId[dep.first];
        int v = nodeToId[dep.second];
        graph[u].push_back(v);
        indegree[v]++;
    }
    
    // 第三步：拓扑排序
    std::queue<int> q;
    for (int i = 0; i < n; i++) {
        if (indegree[i] == 0) {
            q.push(i);
        }
    }
    
    std::vector<int> topoOrder;
    while (!q.empty()) {
        int u = q.front();
        q.pop();
        topoOrder.push_back(u);
        
        for (int v : graph[u]) {
            indegree[v]--;
            if (indegree[v] == 0) {
                q.push(v);
            }
        }
    }
    if (topoOrder.size() != n) {
        // 存在环，无法进行拓扑排序
        std::cout << "存在环，无法进行拓扑排序" << std::endl;
        return {};
    }
    
    // 第四步：逆序计算DP
    std::vector<int> dp(n, 0);
    for (int i = topoOrder.size() - 1; i >= 0; i--) {
        int u = topoOrder[i];
        for (int v : graph[u]) {
            dp[u] = std::max(dp[u], dp[v] + 1);
        }
    }
    
    // 第五步：转换回Vec2D到长度的映射
    std::unordered_map<Vec2D, int, Vec2DHash> result;
    for (int i = 0; i < n; i++) {
        result[idToNode[i]] = dp[i];
    }
    
    return result;
}

int TopoTest() {
    // 示例依赖关系
    // (0,1) -> (0,2) -> (1,0) -> (2,0)->(0,2)
    std::vector<std::pair<Vec2D, Vec2D>> dependencies = {
        {Vec2D(0, 1), Vec2D(0, 2)},
        {Vec2D(0, 2), Vec2D(1, 0)},
        {Vec2D(1, 0), Vec2D(2, 0)},
        {Vec2D(2, 0), Vec2D(0, 2)},
    };
    
    std::cout << "依赖关系：" << std::endl;
    for (const auto& dep : dependencies) {
        std::cout << "(" << dep.first.x << "," << dep.first.y << ") -> " 
             << "(" << dep.second.x << "," << dep.second.y << ")" << std::endl;
    }
    
    auto result = ComputeMaxDependencyLengthVec2D(dependencies);
    
    std::cout << "\n每个节点的最大依赖长度：" << std::endl;
    for (const auto& entry : result) {
        std::cout << "节点 (" << entry.first.x << "," << entry.first.y 
             << "): " << entry.second << std::endl;
    }
    
    return 0;
}

