# General Notes
Notes for working with GO, Docker and Kubernetes ( minikube for local development and testing )

## VM DRIVER: ( hyperkit )
curl -Lo docker-machine-driver-hyperkit https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
&& chmod +x docker-machine-driver-hyperkit \
&& sudo cp docker-machine-driver-hyperkit /usr/local/bin/ \
&& rm docker-machine-driver-hyperkit \
&& sudo chown root:wheel /usr/local/bin/docker-machine-driver-hyperkit \
&& sudo chmod u+s /usr/local/bin/docker-machine-driver-hyperkit


kubernetes-cli curl download:
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/darwin/amd64/kubectl

## Install ( specific-version: v1.9.2 ) - kubectl version -> 1.9.6
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.9.6/bin/darwin/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl


## ZSH Configuration
~./zshrc

if [ $commands[kubectl] ]; then
  source <(kubectl completion zsh)
fi

### OR: ( oh-my-zsh)
plugins=(kubectl)


minikube start --vm-driver=hyperkit
minikube start --vm-driver hyperkit


- Use the Docker registry of the Docker daemon running inside Minikube’s vm instance. To point the ‘docker’ command to your Minikube’s Docker daemon, type (unix shells):
eval $(minikube docker-env)

docker build -t hello-node:v1 .

- List docker images Minikube’s Docker daemon:
minikube ssh docker images


## Create deployment :
- A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking
- Use the kubectl run command to create a Deployment that manages a Pod

kubectl run hello-node --image=hello-node:v1 --port=8080 --image-pull-policy=Never

kubectl run k8test --image=k8test:v1 --port=8080 --image-pull-policy=Never


- Clean up:
kubectl delete service hello-node
kubectl delete deployment hello-node

- or:



###########################################################
#######   UPDATES                                   #######
###########################################################

### Kubernetes macOS Local setup:
* Hyperkit driver:
```
curl -Lo docker-machine-driver-hyperkit https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
&& chmod +x docker-machine-driver-hyperkit \
&& sudo cp docker-machine-driver-hyperkit /usr/local/bin/ \
&& rm docker-machine-driver-hyperkit \
&& sudo chown root:wheel /usr/local/bin/docker-machine-driver-hyperkit \
&& sudo chmod u+s /usr/local/bin/docker-machine-driver-hyperkit
```




## VM DRIVER: ( hyperkit )
curl -Lo docker-machine-driver-hyperkit https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
&& chmod +x docker-machine-driver-hyperkit \
&& sudo cp docker-machine-driver-hyperkit /usr/local/bin/ \
&& rm docker-machine-driver-hyperkit \
&& sudo chown root:wheel /usr/local/bin/docker-machine-driver-hyperkit \
&& sudo chmod u+s /usr/local/bin/docker-machine-driver-hyperkit


kubernetes-cli curl download:
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/darwin/amd64/kubectl

Alternatively with homebrew:
`brew install kubernetes-cli`
`brew uninstall kubernetes-cli`

## Install ( specific-version: v1.9.2 ) - kubectl version -> 1.9.6
curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.9.6/bin/darwin/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl


## ZSH Configuration
~./zshrc

if [ $commands[kubectl] ]; then
  source <(kubectl completion zsh)
fi

### OR: ( oh-my-zsh)
plugins=(kubectl)


minikube start --vm-driver=hyperkit
minikube start --vm-driver hyperkit


- Use the Docker registry of the Docker daemon running inside Minikube’s vm instance. To point the ‘docker’ command to your Minikube’s Docker daemon, type (unix shells):
`eval $(minikube docker-env)`

`docker build -t hello-node:v1` .

- List docker images Minikube’s Docker daemon:
`minikube ssh docker images`


## Create deployment :
- A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking
- Use the kubectl run command to create a Deployment that manages a Pod

`kubectl run hello-node --image=hello-node:v1 --port=8080 --image-pull-policy=Never`



```
kubectl get pod:
NAME                      READY     STATUS    RESTARTS   AGE
k8test-858f6f7d45-zllvn   1/1       Running   0          11s
```


```
kubectl get deployment
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
k8test    1         1         1            1           17s
```

```
kubectl expose deployment k8test --type=LoadBalancer
service "k8test" exposed
```

```
kubectl get service
NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
k8test       LoadBalancer   10.104.188.23   <pending>     8080:32192/TCP   6s
kubernetes   ClusterIP      10.96.0.1       <none>        443/TCP          21m
```

```
minikube service k8test
Opening kubernetes service default/k8test in default browser...
```

* service running at: ```http://192.168.64.3:32192/```

### Clean up:

* Delete Service:
```
kubectl delete service k8test
service "k8test" deleted
```
* Delete Deployment:
```
kubectl delete deployment k8test
deployment "k8test" deleted
```

* Delete images running on minikube daemon:
```
docker rmi k8test:v1 -f
```

* Stop minikube vm:
```
minikube stop
eval $(minikube docker-env -u)
```

`kubectl delete service hello-node`
`kubectl delete deployment hello-node`

- or Optionally, force removal of the Docker images created:
`docker rmi hello-node:v1 hello-node:v2 -f`

- Stop minikube vm:
`minikube stop
eval $(minikube docker-env -u)`


- Set of Pods targeted by a service is usually determined by a LabelSelector:

* ClusterIP (default) - Exposes the Service on an internal IP in the cluster. This type makes the Service only reachable from within the cluster.
* NodePort - Exposes the Service on the same port of each selected Node in the cluster using NAT. Makes a Service accessible from outside the cluster using <NodeIP>:<NodePort>. Superset of ClusterIP.
* LoadBalancer - Creates an external load balancer in the current cloud (if supported) and assigns a fixed, external IP to the Service. Superset of NodePort.
* ExternalName - Exposes the Service using an arbitrary name (specified by externalName in the spec) by returning a CNAME record with the name. No proxy is used. This type requires v1.7 or higher of kube-dns.

- Examples:
* `kubectl expose deployment/kubernetes-bootcamp --type="NodePort" --port 8080`
* `export NODE_PORT=$(kubectl get services/kubernetes-bootcamp -o go-template='{{(index .spec.ports 0).nodePort}}')
echo NODE_PORT=$NODE_PORT`
* `curl $(minikube ip):$NODE_PORT`

## Accessing service:
- `minikube service k8-test`
- `minikube service k8-test --url`
- `curl $(minikube service k8-test --url)`
