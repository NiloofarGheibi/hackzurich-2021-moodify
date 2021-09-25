FROM python:3

ADD model.json /
ADD model.h5 /
ADD inference.py /

RUN pip install keras numpy

CMD [ "python3", "./inference.py" ]
