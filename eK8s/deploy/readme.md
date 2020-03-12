## 部署服务

### 创建命名空间
```
kubectl create namespace goms  
```
### 部署服务
```
./apllay.sh
```
### 删除服务
```
./delete.sh
```

## 逐个部署
### mysql
```
kubectl apllay -f mysql-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl apllay -f mysql-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl describe pod mysql-deploy -n goms  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```
### redis  
```
kubectl apllay -f redis-sts.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n goms  

kubectl apllay -f redis-svc.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n goms  

kubectl describe pod redis-deploy -n goms  
redis-cli -h 192.168.43.204 -p 31002  
```
### user  
```
kubectl apllay -f user-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl apllay -f user-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl describe pod user-deploy -n goms  
curl 192.168.43.204:31003/ping  
```

## 其他

### log
```
kubectl logs -n goms service/redis-svc  
```
### login
```
kubectl exec -it pod/user-deploy-7fc88fdcbf-gn7kl  -n goms  -- /bin/sh  
kubectl exec -it deployment.extensions/mysql-deploy  -n goms  -- /bin/sh  
kubectl exec -it service/user-svc  -n goms  -- /bin/sh  
```

