## 部署

# patch
```
patch-svc.sh
```
# login

mysql
```
kubectl get rs,pod,deploy,svc,ep

kubectl describe pod mysql-deploy
mysql -h 192.168.43.204 -P 31001 -u root -p
```

redis
```
kubectl get rs,pod,deploy,sts,svc,ep

kubectl describe pod redis-deploy
redis-cli -h 192.168.43.204 -p 31002
```

user
```
kubectl get rs,pod,deploy,svc,ep

kubectl describe pod user-deploy
curl 192.168.43.204:31003/ping
```

