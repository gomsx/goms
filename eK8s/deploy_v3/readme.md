## 部署服务

### 创建命名空间
```
kubectl create namespace ek8sv3
```
### 创建 cofigmap
```
kubectl create configmap ek8sv3 --from-file=./configs
kubectl describe configmaps ek8sv3
kubectl get configmaps ek8sv3 -o yaml
```
### 部署服务
```
./apply.sh
```
### 删除服务
```
./delete.sh
```

## 逐个部署

### volume
```
sudo mkdir -p /var/lib/mysqlx/vol-3
sudo chmod 777 /var/lib/mysqlx/vol-3
ls -l -a /var/lib/mysqlx/vol-3
```
### mysql
```
kubectl apply -f mysql-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv3  

kubectl apply -f mysql-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv3  

kubectl describe pod mysql-deploy -n ek8sv3  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```
### redis  
```
kubectl apply -f redis-sts.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n ek8sv3  

kubectl apply -f redis-svc.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n ek8sv3  

kubectl describe pod redis-deploy -n ek8sv3  
redis-cli -h 192.168.43.204 -p 31002  
```
### user  
```
kubectl apply -f user-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv3  

kubectl apply -f user-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv3  

kubectl describe pod user-deploy -n ek8sv3  
curl 192.168.43.204:31003/ping  
```

## 其他

### log
```
kubectl logs -n ek8sv3 service/redis-svc  
```
### login
```
kubectl exec -it pod/user-deploy-7fc88fdcbf-gn7kl -n ek8sv3 -- /bin/sh  
kubectl exec -it deployment.extensions/mysql-deploy -n ek8sv3 -- /bin/sh  
kubectl exec -it service/user-svc -n ek8sv3 -- /bin/sh  
```