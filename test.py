import numpy as np
import pandas as pd
from scipy.optimize import minimize
from LogisticRegression import Logistic_Regression
import matplotlib.pyplot as plt

'''
do some test
2023 / 12 / 7
逻辑回归实现多分类任务
变为3个二分类
'''


data =  pd.read_csv("D:\Stu\PythonStu\LinerReression\data\iris.csv")
test_data = data
# print(test_data.shape)
labels = test_data['class']
# print(labels)
test_data.drop('class',axis = 1,inplace = True)
test_data = test_data.to_numpy()

logis = Logistic_Regression(test_data,labels)
iteration_times = 100
theta, cost = logis.train(iteration_times = iteration_times)
'''
可视化损失展示
# print(theta)
# print(cost)

# y = cost[0]
# # print(range(len(y)))
# x = range(len(y))
# print(y)
# plt.scatter(x,y)
# plt.show()
'''


'''
模型验证
'''

ans = logis.predict(data)
print(ans)

# labels = ['SETOSA','VERSICOLOR','VIRGINICA']

# for index , labels_tab in enumerate(data['class'].unique()):
#     # print(index,labels_tab)
#     current_label = (data['class'] == labels_tab).astype(float)
#     print(current_label.shape)
# 'series类型支持 np.dot()'



