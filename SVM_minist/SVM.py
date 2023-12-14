import pandas as pd
import numpy as np
from sklearn.preprocessing import StandardScaler
from sklearn import svm
from PIL import Image
'''
支持向量机的代码实现
'''
# data = pd.read_csv("train-images-idx3-ubyte.gz",compression='gzip')

# a = [1,2,3,4,5]
# print(a[0:3])
data = open('train-images.idx3-ubyte','rb')
labels = open('train-labels.idx1-ubyte','rb')
f = data.read()
f_l = labels.read()

# magic_number = int.from_bytes(f[0:4],'big')
# number_images = int.from_bytes(f[4:8],'big')
# num_rows = int.from_bytes(f[8:12],'big')
# num_columbs = int.from_bytes(f[12:16],'big')
label = []
for i in range(1,60001):
    label.append(int.from_bytes(f_l[7+i:8+i],'big'))
# print(label)
# print(len(label))

input_x = []
for i in range(1,60001):
    image = [item for item in f[16 + 28 * 28 *(i-1):16 + 28 * 28 * i]]
    input_x.append(image)
    '第16字节之后以28*28(字节)为单位60000张图片'
    # image_np = np.array(image,dtype=np.uint8).reshape(28,28)
    # print(image_np)
processed_data = StandardScaler().fit_transform(input_x)
# for i in range(2,60001):
#     image = [item for item in f[16 + 28 * 28*(i-1):16 + 28 * 28 * i]]
#     image_nparray = np.array(image,dtype = np.uint8)
#     data = np.concatenate((image_np,image))

clf = svm.SVC(kernel='sigmoid',decision_function_shape='ovr')
clf.fit(processed_data,label)

# image_np = np.array(image,dtype=np.uint8).reshape(28,28)
# im = Image.fromarray(image_np)
# processed_data = StandardScaler().fit_transform(image_np).flatten()
# sklearn标准化库,对数据进行处理

pre = Image.open('predict.png')
grayscale_image = pre.convert("L")
pixel_values = list(grayscale_image.getdata())


