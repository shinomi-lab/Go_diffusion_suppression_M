import networkx as nx
import random

# ランダムグラフの作成
G = nx.erdos_renyi_graph(1000, 0.01)

# 残したいノードの個数を指定
remaining_nodes = random.sample(list(G.nodes()), k=500)

# 指定したノード以外を削除
G_sub = G.subgraph(remaining_nodes).copy()

print(f"元のグラフのノード数: {G.number_of_nodes()}")
print(f"部分グラフのノード数: {G_sub.number_of_nodes()}")
