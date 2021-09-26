from flask import Flask, jsonify, request
from statsmodels.tsa.arima.model import ARIMA
from random import randrange

app = Flask(__name__)

data = [randrange(0,4) for x in range(1, 90)]

model = ARIMA(data, order=(1, 1, 1))
model_fit = model.fit()

@app.route('/predictions', methods=['POST'])
def get_prediction():
  response = request.get_json()
  data = response['data']
  predictions = response['n']
  for i in range(predictions):
    yhat = model_fit.predict(len(data), len(data), typ='levels')
    data.append(int(yhat))
  return {"predictions": data[-predictions:]}

if __name__ == "__main__":
    app.run()
