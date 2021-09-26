# Using lightweight alpine image
#FROM python:3.6-alpine
FROM python:3.6-slim-buster

#RUN apk add g++ 

# Installing packages
RUN pip install --no-cache-dir pipenv 

# Defining working directory and adding source code
WORKDIR /usr/src/app
COPY bootstrap.sh ./
COPY inference.py ./
COPY requirements.txt ./

RUN pip install -r requirements.txt

# Install API dependencies
#RUN pipenv install

# Start app
EXPOSE 5000
ENTRYPOINT ["/usr/src/app/bootstrap.sh"]
