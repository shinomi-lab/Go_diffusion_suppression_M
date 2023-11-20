import matplotlib.pyplot as mplot
import numpy as np

n=10

targetData=np.loadtxt('Twitter_node_folower_supp_allkurikasi'+str(n)+'.csv',skiprows=1,delimiter=',')
print(targetData.shape)
x=[]
y=[]
z=[]
w=[]

for xyz in targetData:
    x.append(xyz[0])
    y.append(xyz[1])
    z.append(xyz[2])
    w.append(xyz[3])

# print(x)
# print(y)
# print(z)
fig=mplot.figure()
ax=fig.add_subplot(projection='3d')
ax.scatter(x,y,z,color="blue")
ax.scatter(x,y,w,color="red")

ax.set_xlabel('node_num')
ax.set_ylabel('folower_num')
ax.set_zlabel('suppression')

mplot.show()
