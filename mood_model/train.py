import numpy as np 
import pandas as pd 
import matplotlib.pyplot as plt
import seaborn as sns
from keras.models import Sequential
from keras.layers import Dense
from keras.wrappers.scikit_learn import KerasClassifier
from keras.utils import np_utils
import tensorflow as tf
from sklearn.model_selection import cross_val_score, KFold, train_test_split
from sklearn.preprocessing import LabelEncoder,MinMaxScaler
from sklearn.pipeline import Pipeline
from sklearn.metrics import confusion_matrix, accuracy_score
from keras.models import model_from_json

tf.compat.v1.disable_eager_execution()
tf.compat.v1.disable_v2_behavior()

df = pd.read_csv("data_moods.csv")

col_features = df.columns[6:-3]
X= MinMaxScaler().fit_transform(df[col_features])
X2 = np.array(df[col_features])
Y = df['mood']

#Encodethe categories
encoder = LabelEncoder()
encoder.fit(Y)
encoded_y = encoder.transform(Y)

X_train,X_test,Y_train,Y_test = train_test_split(X,encoded_y,test_size=0.2,random_state=15)

target = pd.DataFrame({'mood':df['mood'].tolist(),'encode':encoded_y}).drop_duplicates().sort_values(['encode'],ascending=True)
print(target)

model = Sequential()
#Add 1 layer with 8 nodes,input of 4 dim with relu function
model.add(Dense(8,input_dim=10,activation='relu'))
#Add 1 layer with output 3 and softmax function
model.add(Dense(4,activation='softmax'))
#Compile the model using sigmoid loss function and adam optim
model.compile(loss='sparse_categorical_crossentropy',optimizer='adam',
                 metrics=['accuracy'])

model.fit(X_train,Y_train, epochs=300, batch_size=200, verbose=0)

# serialize model to JSON
model_json = model.to_json()
with open("model.json", "w") as json_file:
    json_file.write(model_json)
# serialize weights to HDF5
model.save_weights("model.h5")
print("Saved model to disk")
