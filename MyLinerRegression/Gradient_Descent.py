import numpy as np
import pandas as pd
from utils.features import prepare_for_training
import matplotlib.pyplot as plt


class LinerRegression:
    def __init__(self,data,labels,polynomial_degree = 0,sinusoid_degree = 0,normalize_data = True):
        (data_processed,
         features_mean,
         features_deviation) = prepare_for_training(data,polynomial_degree = 0,sinusoid_degree = 0,normalize_data = True)

        self.data = data_processed
        self.labels = labels
        self.features_mean =features_mean
        self.features_deviation = features_deviation
        self.polynomial_degree = polynomial_degree
        self.sinusoid_degree = sinusoid_degree
        self.normalize_data = normalize_data
        
        num_features = self.data.shape[1]
        self.theta = np.zeros((num_features,1))
        
    def train(self,alpha,iteration_times = 100):
            # 训练函数
            # 返回梯度下降的theta结果，以及损失的列表
            cost_history = self.Gradient_descent(iteration_times,alpha)
            # 实际theta更新在Gradient_descent中
            return self.theta , cost_history
        
    
    def Gradient_descent(self,iteration_times,alpha):
        # 更新theta并返回损失列表
        cost_history = []
        for _ in range(iteration_times):
            self.Step(alpha)
            cost_history.append(self.cost())
        
        return cost_history
    
    def Step(self,alpha):
        # 更新theta值
        '''optional: 记录这一步的损失'''
        theta = self.theta
        nums_examples = self.data.shape[0]
        pred = LinerRegression.Pred(self.data,self.theta)
        delta = pred - self.labels
        # theta 更新为行向量
        theta = theta -(alpha*np.dot(delta.transpose(),self.data)/nums_examples).transpose()
        self.theta = theta
    
    @staticmethod
    def Pred(data,theta):
        # 预测利用当前的theta 与数据输出y
        cur_pred = np.dot(data,theta)
        # nums_examples * 1列向量
        return cur_pred
    
    def cost(self):
        #计算当前损失，返回出去
        nums_examples = self.data.shape[0]
        pred = LinerRegression.Pred(self.data,self.theta)
        delta = pred - self.labels
        cur_cost = 1/2*np.dot(delta.transpose(),delta)/nums_examples
        return cur_cost
    
    
    
    
    # if __name__ == "__main__":
    #     # vector1 = np.array([1,2,3])
    #     # vector2 = np.array([[1],[2],[3]])
        
    #     # print(np.dot(vector1,vector2)/3)
    #     data = pd.read_csv("D:\\Stu\PythonStu\\LinerReression\\data\\world-happiness-report-2017.csv")
    #     # print(data.info)
    #     x = data[['Economy..GDP.per.Capita.']].to_numpy()
    #     y = data[['Happiness.Score']].to_numpy()
    #     print(data.shape)
    #     '''
    #     data["column"] 返回始终为shape(,n)的Pandas系列,也就是说,它没有列,总是只有一行。
    #     在 data[["column"]] 返回形状为(m,n)的Pandas数据帧
    #     用 dataframe.to_numpy()也可以
    #     '''
    #     # print('x = ',x)
    #     print('y = ',y)
    #     '可视化'
    #     plt.scatter(x=x,y=y)
    #     plt.show()