from numpy.random import *
import numpy as np
from matplotlib import pyplot as plt
import seaborn as sns

# 平均
mu = np.array([[0.7], [0.7]])

# 分散共分散行列
cov = np.array([[1.0, 0.9], [0.9, 1.0]])

L = np.linalg.cholesky(cov)
dim = len(cov)

random_list = []

for i in range(1000):
    z = np.random.randn(dim, 1)
    random_list.append((np.dot(L, z)+mu).reshape(1, -1).tolist()[0])

# print(random_list)
# print(random_list[0, :])

# sns.jointplot([x[0] for x in random_list], [x[1] for x in random_list])
# plt.show()

mu = [0, 0]
sigma = [[1.0, 0.9], [0.9, 1.0]]

# 2次元正規乱数を1万個生成
values = multivariate_normal(mu, sigma, 1000)

# 散布図
sns.jointplot(values[:, 0].map(), values[:, 1])
plt.show()
