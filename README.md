# Webshell
Interactive terminal website built using xterm.js and Go websocket. Try it out at [playground.rlay.cc](https://playground.rlay.cc) 🎉

## Getting Started

## Local Deployment

### Run on local machine (docker-compose)
```
git clone https://github.com/raykrishardi/webshell.git
cd webshell
make build-websh-front
make build-websh-ws
make up

# Need to edit /etc/hosts file for front-end to connect to back-end
vim /etc/hosts
`
127.0.0.1       localhost webshell-ws
`

# Open browser and navigate to localhost:8080
```

### Run on k8s (http)
```
git clone https://github.com/raykrishardi/webshell.git
cd webshell/k8s
bash ingress/install.sh
kubectl apply -f ingress
kubectl apply -f webshell/ws
kubectl apply -f webshell/front

# Need to edit /etc/hosts file for resolving the ingress LB
vim /etc/hosts
`
<ingress LB IP>       playground.rlay.cc playground-ws.rlay.cc
` 

# Open browser and navigate to http://playground.rlay.cc
```

### Run on k8s (https)
```
git clone https://github.com/raykrishardi/webshell.git
cd webshell/k8s
bash ingress/install.sh
bash cert-manager/install.sh
kubectl apply -f cert-manager
kubectl apply -f ingress
kubectl apply -f webshell/ws
kubectl apply -f webshell/front

# Need to edit /etc/hosts file for resolving the ingress LB
vim /etc/hosts
`
<ingress LB IP>       playground.rlay.cc playground-ws.rlay.cc
` 

# Open browser and navigate to https://playground.rlay.cc
```
