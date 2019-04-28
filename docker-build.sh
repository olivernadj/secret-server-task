#!/usr/bin/env bash


# validate the arguments
if [[ -z "$1" ]]
then
    echo "Version must be provided"
    exit 1
fi


cd redis &&\
  docker build -t secret-api-redis:$1 . &&\
  docker tag secret-api-redis:$1 olivernadj/secret-api-redis:$1 &&\
  docker push olivernadj/secret-api-redis:$1 &&\
  docker tag secret-api-redis:$1 olivernadj/secret-api-redis &&\
  docker push olivernadj/secret-api-redis &&\
  cd ..  &&\
cd redis-exporter &&\
  docker build -t secret-api-redis-exporter:$1 . &&\
  docker tag secret-api-redis-exporter:$1 olivernadj/secret-api-redis-exporter:$1 &&\
  docker push olivernadj/secret-api-redis-exporter:$1 &&\
  docker tag secret-api-redis-exporter:$1 olivernadj/secret-api-redis-exporter &&\
  docker push olivernadj/secret-api-redis-exporter &&\
  cd ..  &&\
cd goapi &&\
  docker build -t secret-api-goapi:$1 . &&\
  docker tag secret-api-goapi:$1 olivernadj/secret-api-goapi:$1 &&\
  docker push olivernadj/secret-api-goapi:$1 &&\
  docker tag secret-api-goapi:$1 olivernadj/secret-api-goapi &&\
  docker push olivernadj/secret-api-goapi &&\
  cd ..  &&\
cd prometheus &&\
  docker build -t secret-api-prometheus:$1 . &&\
  docker tag secret-api-prometheus:$1 olivernadj/secret-api-prometheus:$1 &&\
  docker push olivernadj/secret-api-prometheus:$1 &&\
  docker tag secret-api-prometheus:$1 olivernadj/secret-api-prometheus &&\
  docker push olivernadj/secret-api-prometheus &&\
  cd ..  &&\
cd grafana &&\
  docker build -t secret-api-grafana:$1 . &&\
  docker tag secret-api-grafana:$1 olivernadj/secret-api-grafana:$1 &&\
  docker push olivernadj/secret-api-grafana:$1 &&\
  docker tag secret-api-grafana:$1 olivernadj/secret-api-grafana &&\
  docker push olivernadj/secret-api-grafana &&\
  cd ..  &&\
cd nginx &&\
  docker build -t secret-api-nginx:$1 . &&\
  docker tag secret-api-nginx:$1 olivernadj/secret-api-nginx:$1 &&\
  docker push olivernadj/secret-api-nginx:$1 &&\
  docker tag secret-api-nginx:$1 olivernadj/secret-api-nginx &&\
  docker push olivernadj/secret-api-nginx &&\
  cd ..