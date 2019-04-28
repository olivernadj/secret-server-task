# Secret Server Coding Task

## Introduction
Your task is to implement a secret server. The secret server can be used to store and share secrets
using the random generated URL. But the secret can be read only a limited number of times after that
it will expire and won’t be available. The secret may have TTL. After the expiration time the secret
won’t be available anymore. You can find the detailed API documentation in the swagger.yaml file.
We recommend to use [Swagger](https://editor.swagger.io/) or any other OpenAPI implementation to
read the documentation. 

## Task
**Implementation**: You have to implement the whole Secret Server API. If it is not specified you can choose the technology
you want to use (database, programming language, etc). However it would be wise to store the data using encryption now this is not part of the task. You can use plain text to store your secrets.
The response can be XML or JSON too. Use a configuration file to switch between the two response type. 

**Hosting**: You also have to deploy and host the service. There are hundreds of free solutions to do this. So this shouldn't
be an issue. If this API was used in production, then HTTPS would be a must but this is not the case now, use HTTP instead of HTTPS.

**Share the code**: Upload the code to your GitHub account and share with us.

## Questions
It is totaly OK to ask if something is not clear. 

## Production endpoints
Make sure your hosts config resolve these hosts.
```
178.128.143.215 api.your-secret-server.com grafana.your-secret-server.com prometheus.your-secret-server.com
```


#### Development Endpoints of services
- Secret API
  - http://api.your-secret-server.com/metrics Prometheus exporter
  - http://api.your-secret-server.com/v1/ui/ Swagger Documentation
  - http://api.your-secret-server.com/ Hello world
- Grafana Access (admin/5ecret):
  - http://grafana.your-secret-server.com/
- Prometheus: scraps metrics and store it in time series
  - http://prometheus.your-secret-server.com/graph for debugging purpose


## Build and maintain the containers

### Development on docker compose
Git clone:
```
~$git clone https://github.com/olivernadj/secret-server-task
~$ cd secret-server-task
```
To get the containers up and running, run this command:
```
$ docker-compose up -d --build
```
To check running containers, run this command:
```
$ watch "sudo docker ps --format='table{{.Image}}\t{{.Names}}\t{{.Status}}\t{{.Ports}}'"
```
For stop services, run this command:
```
$ docker-compose stop
```

#### Development Endpoints of services
- Secret API
  - http://localhost:8080/metrics Prometheus exporter
  - http://localhost:8080/v1/ui/ Swagger Documentation
  - http://localhost:8080/ Hello world
- Prometheus: scraps metrics and store it in time series
  - http://localhost:9090/graph
- Grafana Access (admin/5ecret):
  - http://localhost:3000/
  
## Build docker images
```
./docker-build.sh
```


## Maintain Kubernetes

### Create deployments and services
```
cd kube
kube$ kubectl create -f goapi-deployment.yaml,grafana-deployment.yaml,nginx-deployment.yaml,prometheus-deployment.yaml,\
redis-deployment.yaml,redis-exporter-service.yaml,goapi-service.yaml,grafana-service.yaml,nginx-service.yaml,\
prometheus-service.yaml,redis-exporter-deployment.yaml,redis-service.yaml

```

### Inspect pod
```
kubectl describe pod <pod-id> 
```


### Update image
```
cd kube
kubectl set image deployment/prometheus prometheus=olivernadj/secret-api-prometheus:1.0.2

```

### List services

```
kube$ kubectl get svc
NAME         CLUSTER-IP       EXTERNAL-IP      PORT(S)        AGE
goapi        10.245.164.247   <none>           8080/TCP       28s
grafana      10.245.84.205    <none>           3000/TCP       26s
kubernetes   10.245.0.1       <none>           443/TCP        30m
nginx        10.245.180.76    178.128.143.215  80:30906/TCP   26s
prometheus   10.245.120.173   <none>           9090/TCP       27s

```

### Clean up everything

```
kubectl delete --all pods && \
  kubectl delete --all deployments && \
  kubectl delete --all services
```

## Fake load test

```
$ ./fakeload.sh http://api.your-secret-server.com/v1

```
