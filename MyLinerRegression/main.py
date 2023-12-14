from Gradient_Descent import LinerRegression
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt


if __name__ == "__main__":
    # 获取数据
    data = pd.read_csv("D:\\Stu\PythonStu\\LinerReression\\data\\world-happiness-report-2017.csv")
    # print(data.shape)
    
    # x_data = data[['Economy..GDP.per.Capita.']]
    # y_data = data[['Happiness.Score']]
    
    train_set = data.sample(frac=0.8)
    test_set = data.drop(train_set.index)
    
    # print(train_set)
    # print(test_set)
    
    # print(train_set.shape)
    # print(test_set.shape)
    '''
    数据集的划分
    '''
    train_x_data = train_set[['Economy..GDP.per.Capita.']].to_numpy()
    train_y_data = train_set[['Happiness.Score']].to_numpy()
    
    test_x_data = test_set[['Economy..GDP.per.Capita.']].to_numpy()
    test_y_data = test_set[['Happiness.Score']].to_numpy()
    '''
    data["column"] 返回始终为shape(,n)的Pandas系列,也就是说,它没有列,总是只有一行。
    在 data[["column"]] 返回形状为(m,n)的Pandas数据帧
    用 dataframe.to_numpy()也可以
    '''
    
    liner_regression = LinerRegression(train_x_data,train_y_data)
    
    learning_rate = 0.1
    iteration_times = 100
    
    theta , cost = liner_regression.train(learning_rate,iteration_times=iteration_times)
    