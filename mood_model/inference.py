from keras.models import model_from_json
import numpy as np
import sys

target = ['Calm', 'Energetic', 'Happy', 'Sad']

# load json and create model
json_file = open('model.json', 'r')
loaded_model_json = json_file.read()
json_file.close()
loaded_model = model_from_json(loaded_model_json)

# load weights into new model
loaded_model.load_weights("model.h5")
print("Loaded model from disk")

# evaluate loaded model on test data
loaded_model.compile(loss='sparse_categorical_crossentropy', optimizer='adam', metrics=['accuracy'])

col_features = [float(i) for i in sys.argv[1].split(',')]

x = np.array(col_features).reshape(-1,1).T
test = np.array(loaded_model.predict(x))

index = np.argmin(test)
mood = target[index]

print(mood)
