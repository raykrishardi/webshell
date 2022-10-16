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

kubectl create ns webshell
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

kubectl create ns webshell
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

## Using custom root command for the webshell

### docker-compose
```
cd webshell
# Copy the custom binary to webshell-ws folder
cp <path_to_custom_binary> webshell-ws/
# Modify the build args and environment var with the new binary
vim docker-compose.yaml
`
webshell-ws:
    build:
      args:
        ROOT_CMD: "<custom_binary>"
    environment:
      ROOT_CMD: "<custom_binary>"
`
make down && make up
```

#### Example (using [iot-controller CLI](https://github.com/raykrishardi/iot-controller))
```
cd webshell
git clone https://github.com/raykrishardi/iot-controller.git
cd iot-controller
make build-iot-cli
cd ..
cp iot-controller/iot-controller-cli/iot webshell-ws/
vim docker-compose.yaml
`
webshell-ws:
    build:
      args:
        ROOT_CMD: "iot"
    environment:
      ROOT_CMD: "iot"
      GRPC_HOST: "<domain address of the IoT controller>"
      GRPC_PORT: "<port number of the IoT controller service> (eg. 50001)"
`
make down && make up
```

### k8s
```
cd webshell
# Copy the custom binary to webshell-ws folder
cp <path_to_custom_binary> webshell-ws/
# Example custom binary called iot in this case
docker build --no-cache --build-arg ROOT_CMD=iot -f webshell-ws.Dockerfile -t raylayadi/webshell-ws:iot-latest .
docker push raylayadi/webshell-ws:iot-latest

# Set k8s/webshell/ws/ws-deploy.yaml to use the new image (eg. raylayadi/webshell-ws:iot-latest)
# Set the env var value in config map k8s/webshell/ws/ws-cm.yaml
`
data:
  ROOT_CMD: "iot"
  # GRPC_HOST and GRPC_PORT are specific to iot binary
  GRPC_HOST: "<domain address of the IoT controller>"
  GRPC_PORT: "<port number of the IoT controller service> (eg. 50001)"
`
```
