## 部署

deploy-mysql
```
kubectl apply -f mysql-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms-ek8s  

kubectl describe pod mysql-deploy -n goms-ek8s  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```

deploy-redis  
```
kubectl apply -f redis-sts.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n goms-ek8s  

kubectl describe pod redis-deploy -n goms-ek8s  
redis-cli -h 192.168.43.204 -p 31002  
```

deploy-user  
```
kubectl apply -f user-deploy.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms-ek8s  

kubectl describe pod user-deploy -n goms-ek8s  
curl 192.168.43.204:31003/ping  
```

