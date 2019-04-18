FROM python:3.7-slim
WORKDIR /app
ADD . /app
RUN apt-get -y update
RUN apt-get install -y python-pip python-dev build-essential
RUN pip install Flask
EXPOSE 81:81
ENTRYPOINT ["python","hello.py"]
