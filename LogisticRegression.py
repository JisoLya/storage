import numpy as np
from scipy.optimize import minimize
from utils.features import prepare_for_training
'''
minimize 自动求标量函数的最值以及最值点
逻辑回归实现多分类转化为三个二分类

'''
def sigmoid(matrix):

    return 1 / (1 + np.exp(-matrix))

class Logistic_Regression():
    def __init__(self,data,labels,polynomial_degree = 0,sinusoid_degree = 0,normalize_data = True):
        (data_processed , 
         features_mean  ,
         features_deviation) = prepare_for_training(data,polynomial_degree,sinusoid_degree,normalize_data)
        # 数据预处理
        
        self.data = data_processed
        self.labels = labels
        self.unique_labels  = np.unique(labels)# 分类任务，有多少类
        self.features_mean = features_mean
        self.features_deviation = features_deviation
        self.polynomial_degree = polynomial_degree
        self.sinusoid_degree = sinusoid_degree
        self.normalize_data = normalize_data
        
        nums_features = self.data.shape[1]
        nums_unique_labels = np.unique(labels).shape[0]
        #[ theta11 theta12 theta13 ]
        #[ theta21 theta22 theta23 ]
        #[ theta31 theta32 theta33 ]
        # 针对每一个类别都有一个theta参数
        self.theta = np.zeros((nums_unique_labels,nums_features))
        
        
    def train(self,iteration_times = 1000):
        # 输入迭代次数，学习率，每轮进行梯度下降
        # 返回历史损失，以及当前theta值
        cost_histories = []
        # nums_features = self.data.shape[1]
        for labels_index , unique_labels in enumerate(self.unique_labels):
            current_theta = self.theta[labels_index].transpose()
            # 当前标签对应的theta变为列向量
            current_labels = (self.labels == unique_labels).astype(float)
            # 标签列表代表选中哪一行theta
            # label_matrix = np.array(current_labels , ndmin = 2)
            'series类型支持矩阵运算,不转化为矩阵'
            current_theta ,cost_history = Logistic_Regression.gradient_descent(self.data,iteration_times,current_theta,current_labels)
            # 训练
            self.theta[labels_index] = current_theta.transpose()
            cost_histories.append(cost_history)
            # 训练结束赋值
        return self.theta, cost_histories
    
    @staticmethod
    def gradient_descent(data,iteration_times,theta,current_labels):
        # 执行iteration_times次梯度下降
        num_features = data.shape[1]
        cost_history = []
        result = minimize(
            lambda theta :Logistic_Regression.cost_func(data,current_labels,theta)
            ,x0 = theta
            ,method = 'CG'
            ,jac = lambda theta : Logistic_Regression.Step(data,theta,current_labels)
            ,callback = lambda theta:cost_history.append(Logistic_Regression.cost_func(data,current_labels,theta))
            ,options = {'maxiter': iteration_times}
        )
        if not result.success:
            raise ArithmeticError('Can not minimize cost function'+result.message)
        optimized_theta = result.x.reshape(num_features ,1)
        return optimized_theta,cost_history
    
    
    @staticmethod
    def Step(data,theta,labels):
        # 实现一步的梯度下降计算
        nums_examples = data.shape[0]
        pred = Logistic_Regression.hypothesis(data,theta)
        label_diff = pred - labels
        gradient = (1/nums_examples)*np.dot(data.transpose(),label_diff)
        # 计算梯度值, (y_i - h(x_i)) * x_j
        return gradient.transpose().flatten()
    
    @staticmethod
    def cost_func(data,labels,theta):
        # 损失函数定义,交叉熵
        nums_examples = data.shape[0]
        pred = Logistic_Regression.hypothesis(data,theta)
        y_is_set = np.dot(labels[labels==1].transpose(),np.log(pred[labels == 1]))
        y_is_not_set = np.dot(1-labels[labels == 0].transpose(),np.log(1-pred[labels == 0]))
        cur_cost = -1 / nums_examples*(y_is_not_set + y_is_set)
        return cur_cost
    '函数用到的theta已全部转换为列向量'
    @staticmethod
    def hypothesis(data,theta):
        # 概率值
        return sigmoid(np.dot(data,theta))
    
    def predict(self,data):
        '''
        对外部输入数据进行预测
        其余模块使用此模块需要对 data赋值为self.data
        '''
        nums_examples = data.shape[0]
        data_processed = prepare_for_training(data, self.polynomial_degree, self.sinusoid_degree,self.normalize_data)[0]
        # 外部传入数据需要先进行处理
        # 否则会对预测结果产生影响
        pred_probility = Logistic_Regression.hypothesis(data_processed,self.theta.T)
        max_prob_index = np.argmax(pred_probility, axis=1)
        'axis = 0表示取每一行的最大值,axis = 1表示取每一列的最大值(索引)' 
        class_prediction = np.empty(max_prob_index.shape,dtype = object)
        for index,label in enumerate(self.unique_labels):
            class_prediction[max_prob_index == index] = label
        return class_prediction.reshape((nums_examples,1))