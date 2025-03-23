# Dev

## Wgo
> https://github.com/bokwoon95/wgo
Golang web server hot reload

```
go install github.com/bokwoon95/wgo@latest
wgo run cmd/main.go
```

## Templ
Compile templates into .go files
```
wget https://github.com/a-h/templ/releases/download/v0.3.833/templ_Linux_x86_64.tar.gz
tar -xzf templ_Linux_x86_64.tar.gz 

go get github.com/a-h/templ

./templ generate --watch
```
*Also, get templ-vscode extension*

## Tailwind
```
apt install watchman
wget https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.14/tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
chmod +x tailwindcss
./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch
```

# Build & Run locally
> https://hub.docker.com/repository/docker/golle/ipcalc/general

```
docker build . -t golle/ipcalc --target production
docker run -p 8000:8000 golle/ipcalc
```
*You may need to run **docker login** first*

# Deploy to Kubernetes
We create a deployment/service file and apply it using kubectl.

## golang-ipcalc.yaml:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: golang-ipcalc
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      ipcalc: web
  template:
    metadata:
      labels:
        ipcalc: web
    spec:
      containers:
        - name: golang-ipcalc
          image: golle/ipcalc
          imagePullPolicy: Always
---
apiVersion: v1
kind: Service
metadata:
  name: golang-ipcalc
  namespace: default
spec:
  type: NodePort
  selector:
    ipcalc: web
  ports:
    - port: 8000
      targetPort: 8000
      nodePort: 30001
```

Apply the file on your kubernetes cluster:
```
root@k3s-1:~# kubectl apply -f golang-ipcalc.yaml 
deployment.apps/golang-ipcalc configured
service/ipcalc-entrypoint unchanged

root@k3s-1:~# kubectl get deployments
NAME            READY   UP-TO-DATE   AVAILABLE   AGE
golang-ipcalc   1/1     1            1           49m

root@k3s-1:~# kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
golang-ipcalc-c4d9bdb58-n2wkv   1/1     Running   0          46m

root@k3s-1:~# kubectl get services
NAME                TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
ipcalc-entrypoint   NodePort    10.43.99.25   <none>        8000:30001/TCP   50m
kubernetes          ClusterIP   10.43.0.1     <none>        443/TCP          18h
```
*You can now reach the service on the kubernetes <IP-address>:30001*
