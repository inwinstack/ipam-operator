[![Build Status](https://travis-ci.org/inwinstack/ipam.svg?branch=master)](https://travis-ci.org/inwinstack/ipam) [![codecov](https://codecov.io/gh/inwinstack/ipam/branch/master/graph/badge.svg)](https://codecov.io/gh/inwinstack/ipam) [![Docker Pulls](https://img.shields.io/docker/pulls/inwinstack/ipam.svg)](https://hub.docker.com/r/inwinstack/ipam/) ![Hex.pm](https://img.shields.io/hexpm/l/plug.svg)
# IPAM 
The IPAM provides `Pool` and `IP` custom resource to take care of assigning and unassigning individual addresses from pools because Kubernetes cannot create IP addresses out of thin air, so we need to give it CRDs that it can use. 

![](images/architecture.png)

## Building from Source
Clone repo into your go path under `$GOPATH/src`:
```sh
$ git clone https://github.com/inwinstack/ipam.git $GOPATH/src/github.com/inwinstack/ipam
$ cd $GOPATH/src/github.com/inwinstack/ipam
$ make
```

## Debug out of the cluster
Run the following command to debug:
```sh
$ go run cmd/main.go \
    --kubeconfig $HOME/.kube/config \
    --logtostderr \
    -v=2
```

## Deploy in the cluster
Run the following command to deploy the controller:
```sh
$ kubectl apply -f deploy/
$ kubectl -n kube-system get po -l ipam
```
