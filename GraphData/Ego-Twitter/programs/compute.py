import networkx as nx

# エッジリストを読み込む関数
def load_twitter_ego_graph(edges_file):
    G = nx.read_edgelist(edges_file)  # NetworkXのread_edgelistを使用
    return G

# エッジリストファイルのパスを指定
edges_path = "twitter_combined/twitter_combined.txt"  # 実際のファイルパスに変更

# グラフを読み込み
G = load_twitter_ego_graph(edges_path)

# グラフの基本情報を表示
print(G)

# グラフの可視化（matplotlibが必要）
import matplotlib.pyplot as plt
#nx.draw(G, with_labels=False, node_size=10)
#plt.show()

#print("edges:",len(G.edges))
#print("radius",nx.radius(G))
#print("diameter",nx.diameter(G))
#print("argmax:",max_index)
#print("max:",np.max(list2))
#print("ave:",np.mean(list2))
#print("min:",np.min(list2))
#print("all:",list2)

import numpy as np

# 次数分布のプロット（対数スケール）
degree_count = nx.degree_histogram(G)  # 各次数ごとのノード数
degrees = range(len(degree_count))  # 次数
frequency = np.array(degree_count) / sum(degree_count)  # 頻度

# 非ゼロのデータを抽出（対数計算が可能な部分）
nonzero_degrees = np.array([d for d in degrees if d > 0 and degree_count[d] > 0])
nonzero_frequency = np.array([frequency[d] for d in nonzero_degrees])

plt.loglog(nonzero_degrees, nonzero_frequency, marker="o", linestyle="none")
plt.title("Log-Log Degree Distribution")
plt.xlabel("Degree (log scale)")
plt.ylabel("Frequency (log scale)")
#plt.show()

from scipy.optimize import curve_fit

# フィッティング関数（べき乗則）
def power_law(x, a, b):
    return a * x ** (-b)

# フィッティングを行う
params, _ = curve_fit(power_law, nonzero_degrees, nonzero_frequency)

# フィッティング結果のプロット
plt.loglog(nonzero_degrees, nonzero_frequency, marker="o", linestyle="none", label="Data")
plt.loglog(nonzero_degrees, power_law(nonzero_degrees, *params), label=f"Fit: a={params[0]:.2f}, b={params[1]:.2f}")
plt.title("Power-Law Fit")
plt.xlabel("Degree (log scale)")
plt.ylabel("Frequency (log scale)")
plt.legend()
plt.show()
