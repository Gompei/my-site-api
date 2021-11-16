FROM python:3.9-buster

WORKDIR /work/my-site-api

RUN apt-get update && \
    apt-get install -y less vim curl unzip sudo

RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    sudo ./aws/install \

COPY init.sh init.sh
