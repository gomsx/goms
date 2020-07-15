## 部署

service-mysql
```
kubectl apply -f mysql-svc.yaml 
kubectl get rs,pod,deploy,svc,ep 

kubectl describe pod mysql-svc 
mysql -h 192.168.43.204 -P 31001 -u root -p  
```

service-redis  
```
kubectl apply -f redis-svc.yaml 
kubectl get rs,pod,deploy,sts,svc,ep 

kubectl describe pod redis-svc 
redis-cli -h 192.168.43.204 -p 31002  
```

service-user  
```
kubectl apply -f user-svc.yaml 
kubectl get rs,pod,deploy,svc,ep 

kubectl describe pod user-svc 
curl 192.168.43.204:31003/ping  
```

