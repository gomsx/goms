## 部署

service-mysql
```
kubectl apply -f mysql-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl describe pod mysql-svc -n goms  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```

service-redis  
```
kubectl apply -f redis-svc.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n goms  

kubectl describe pod redis-svc -n goms  
redis-cli -h 192.168.43.204 -p 31002  
```

service-user  
```
kubectl apply -f user-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms  

kubectl describe pod user-svc -n goms  
curl 192.168.43.204:31003/ping  
```

