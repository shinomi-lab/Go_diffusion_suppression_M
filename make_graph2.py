import pandas as pd
import numpy as np

df = pd.read_csv("Twitter_node_folower_supp_allkurikasi10.csv",header=None)
df = df.rename(columns={0:'node_num', 1:'follower_num',2:"true_supp",3:"false_supp"})

import csv

filename="Nodes_Twitter_node_folower_supp_allkurikasi10.csv"

with open(filename, encoding='utf8', newline='') as f:
    csvreader = csv.reader(f)
    content = [row for row in csvreader]


df["nodes"] = content

import matplotlib.pyplot as plt
df_node_num = df.loc[df.groupby('node_num')['true_supp'].idxmax()]


df2 = df.copy()

df2['follower_num'] = np.round(df2['follower_num']*3,-2)/3

df_follower = df2.loc[df2.groupby('follower_num')['true_supp'].idxmax()]

indexes = np.setdiff1d(df.index,df_node_num.index)

not_max_df = df.loc[indexes]

x = not_max_df['node_num'].values
y = not_max_df['follower_num'].values

x_max = df_node_num['node_num'].values
y_max = df_node_num['follower_num'].values

x3 = df_follower['node_num'].values
y3 = df_follower['follower_num'].values

print(type(x[0]), type(y[0]))
plt.scatter(x, y)
plt.scatter(x_max, y_max, color = 'red', alpha = 1)
plt.scatter(x3, y3, color = 'yellow', alpha = 0.8)
