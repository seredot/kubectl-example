# Kubernetes resource sample plugin

`kubectl` plugin for retrieving resource sample YAMLs.

All samples originate from kubernetes.io.

### Usage

#### Build and test uging Go

`make build`

`./kubectl-sample deployment`

Output:
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80
```

#### Or download precompiled binaries

Github releases [link](https://github.com/seredot/kubectl-sample/releases).

#### Installing as a `kubectl` plugin

Simply move the compiled binary into a directory in `PATH`

`mv ./kubectl-sample /usr/local/bin`

and then can be used with `kubectl` as

`kubectl sample pod`


#### Build and test using Docker

`make build-docker`

`docker run --rm -i -t kubectl-sample deploy`
