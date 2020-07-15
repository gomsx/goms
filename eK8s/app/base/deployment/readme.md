## 部署

deploy-mysql
```
kubectl apply -f mysql-deploy.yaml 
kubectl get rs,pod,deploy,svc,ep 

kubectl describe pod mysql-deploy 
mysql -h 192.168.43.204 -P 31001 -u root -p  
```

deploy-redis  
```
kubectl apply -f redis-sts.yaml 
kubectl get rs,pod,deploy,sts,svc,ep 

kubectl describe pod redis-deploy 
redis-cli -h 192.168.43.204 -p 31002  
```

deploy-user  
```
kubectl apply -f user-deploy.yaml 
kubectl get rs,pod,deploy,svc,ep 

kubectl describe pod user-deploy 
curl 192.168.43.204:31003/ping  
```

