# -*- coding: utf-8 -*-
"""
Created on Fri Sep 30 10:39:37 2022

@author: finkt
"""
import sys
import os
import pandas as pd
import networkx as nx
from pyvis.network import Network

from viral_centrality import viral_centrality
import json
import numpy as np
from matplotlib import pyplot as plt

tol = 0.001

f = open('congress_network_data.json')
data = json.load(f)

inList = data[0]['inList']
inWeight = data[0]['inWeight']
outList = data[0]['outList']
# outWeight = data[0]['outWeight']
usernameList = data[0]['usernameList']

# print(inList)
# print()
# print()
# print()
# print(outList)


list1 = [[0 for i in range(475)] for j in range(475)]
list2 = []
# list1[0,0] = 4
# list1[0][4] = 1
print("nodes",len(usernameList))

for i,l in enumerate(inList):
    k=0
    for j in l:
        # print(i,j)
        list1[i][j] = 1
        k = k+1
    list2.append(k)

max_index = np.argmax(list2)
adj = pd.DataFrame(list1)
G = nx.from_pandas_adjacency(adj, create_using=nx.Graph)
#
print("edges:",len(G.edges))
print("radius",nx.radius(G))
print("diameter",nx.diameter(G))
print("argmax:",max_index)
print("max:",np.max(list2))
print("ave:",np.mean(list2))
print("min:",np.min(list2))
print("all:",list2)

# degree_sequence = [d for n, d in G.degree()]
#
# # 次数分布のプロット（ヒストグラム）
# plt.hist(degree_sequence, bins=30, density=True)
# plt.title("Degree Distribution")
# plt.xlabel("Degree")
# plt.ylabel("Frequency")
# plt.show()

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
plt.show()

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

# adj = to_np_adjmat(G)
#Golangへ無理やり持っていくように書いたけど使わない
# adj = adj.astype(np.int32)
# print(adj[0])
# print(adj.shape)
# df = DataFrame(adj)
# df.to_csv("graph100node.csv")
# adj.to_json("adj_jsonTwitterInteractionUCongress.txt")

# net = Network(notebook = True)
#
# print(type(G))
# net.from_nx(G)
# net.show("abc.html")
# num_activated = viral_centrality(inList, inWeight, outList, Niter = -1, tol = tol)
#
# plt.scatter(np.array(range(len(num_activated))),num_activated,color='red',label='Viral Centrality')
# plt.xlabel('Node ID',fontsize=15)
# plt.ylabel('Avg Number Activaated',fontsize=15)
