# Server CLI

This repo is for learning and practice the following topics:

* How to debug on a remote Kubernetes cluster with `skaffold` and Cloud Code  pluing in VSCode.

* Using Gin framework to create web services with TDD.

* Using Cobra framework to create a CLI to interact with a rEST API server.

## Requirements

* minikube (v1.9.2+)
* skaffold (v1.9.1+)
* VSCode (v1.45.1+) + Cloud Code plugin

## Run local environment

In order to create the local environment for trying the remote debug in Kubernetes 
the first thing that we need is to have a Kubernetes cluster where deploy our app.
For this example we are going ot use Minikube. Once you have installed Minikube run 
the following command.

```bash
$ minikube start
```

This will start a default context of Minikube with a single host Kubernetes cluster in your local machine, depending of the driver you are using will create a VM or a Docker container to create the cluster. At this point you should have a `~/.kube/config` file that allow you access to your local cluster with `kubectl`. So we are ready to run our app.

Be aware to update the entry of the `server-ingress.yaml` with the ip address of your minikube instance. In order to the this ip run the following command.

```bash
$ minikube ip
172.17.0.3
```

In my case the content of the ingress object is like this

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: server
spec:
  rules:
    - host: www.172.17.0.2.nip.io
      http:
        paths:
          - path: /
            backend:
              serviceName: server
              servicePort: 8080
```

We are using `nip.io` to avoid the editing of `/etc/hosts` file, but if you feel more confortable with that option you can use my college tool [hostctl](https://github.com/guumaster/hostctl) to automate this proccess.

The next step is to clone this repo, open it with VSCode and run from the Debugger `Run/Debug on Kubernetes`.

If everything works fine you should open the URL [http://www.172.17.0.3.nip.io](http://www.172.17.0.3.nip.io) and you should see the main page of our blog server.

## Build server

```bash
go build -o server ./cmd/server
```

## Links

[Semaphore CI article part 1](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin)

[Semaphore CI article part 2](https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin)

[Cloud Code in VSCode](https://www.youtube.com/watch?v=62GLbBDLiPE)
