## 部署

service-mysql
```
kubectl apply -f mysql-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms-ek8s  

kubectl describe pod mysql-deploy -n goms-ek8s  
mysql -h 192.168.43.204 -P 31001 -u root -p  
```

service-redis  
```
kubectl apply -f redis-svc.yaml --record  
kubectl get rs,pod,deploy,sts,svc,ep -n goms-ek8s  

kubectl describe pod redis-svc -n goms-ek8s  
redis-cli -h 192.168.43.204 -p 31002  
```

service-user  
```
kubectl apply -f user-svc.yaml --record  
kubectl get rs,pod,deploy,svc,ep -n goms-ek8s  

kubectl describe pod user-svc -n goms-ek8s  
curl 192.168.43.204:31003/ping  
```

## 调试 

log
```
journalctl -f -u kubelet
kubectl logs -n goms-ek8s service/redis-svc  
```

login
```
kubectl exec -it pod/user-deploy-7fc88fdcbf-gn7kl -n goms-ek8s -- /bin/sh  
kubectl exec -it deployment.extensions/mysql-deploy -n goms-ek8s -- /bin/sh  
kubectl exec -it service/user-svc -n goms-ek8s -- /bin/sh  
```

other
```
kubectl describe node
kubectl get ns
```