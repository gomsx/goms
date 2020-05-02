## 部署服务

### 初始化
```
./init.sh
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

### namespace
```
kubectl create namespace ek8sv2
```
### volume
```
sudo mkdir -p /var/lib/mysqlx/vol-2
sudo chmod 777 /var/lib/mysqlx/vol-2
ls -l -a /var/lib/mysqlx/vol-2
```
### service-mysql
```
kubectl apply -f mysql-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv2  

kubectl apply -f mysql-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv2  

kubectl describe pod mysql-deploy -n ek8sv2  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```
### service-redis  
```
kubectl apply -f redis-sts.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n ek8sv2  

kubectl apply -f redis-svc.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n ek8sv2  

kubectl describe pod redis-deploy -n ek8sv2  
redis-cli -h 192.168.43.204 -p 31002  
```
### service-user  
```
kubectl apply -f user-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv2  

kubectl apply -f user-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n ek8sv2  

kubectl describe pod user-deploy -n ek8sv2  
curl 192.168.43.204:31003/ping  
```

## 调试

### log
```
journalctl -f -u kubelet
kubectl describe node
kubectl logs -n ek8sv2 service/redis-svc  
```
### login
```
kubectl exec -it pod/user-deploy-7fc88fdcbf-gn7kl -n ek8sv2 -- /bin/sh  
kubectl exec -it deployment.extensions/mysql-deploy -n ek8sv2 -- /bin/sh  
kubectl exec -it service/user-svc -n ek8sv2 -- /bin/sh  
```