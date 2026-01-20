from collections import deque

def compute_max_dependency_length(deps):
    # deps: list of (u, v) 表示 u 依赖 v（u -> v）
    nodes = set()
    for u, v in deps:
        nodes.add(u)
        nodes.add(v)
    n = max(nodes) + 1  # 假设节点编号从 0 到 max_node
    
    adj = [[] for _ in range(n)]
    indeg = [0] * n
    
    for u, v in deps:
        adj[u].append(v)
        indeg[v] += 1
    
    # 拓扑排序
    q = deque()
    for i in range(n):
        if indeg[i] == 0:
            q.append(i)
    
    topo = []
    while q:
        u = q.popleft()
        topo.append(u)
        for v in adj[u]:
            indeg[v] -= 1
            if indeg[v] == 0:
                q.append(v)
    
    # 逆序计算 dp
    dp = [0] * n
    for u in reversed(topo):
        for v in adj[u]:
            dp[u] = max(dp[u], dp[v] + 1)
    
    return dp

# 测试例子
deps = [(0, 1), (0, 2), (1, 3), (2, 4), (4, 5)]
result = compute_max_dependency_length(deps)
print("每个节点的最大依赖长度（边数）：")
for i, val in enumerate(result):
    print(f"节点 {i}: {val}")