
## deploy

1, deploy app

```
kubectl apply -k .
```

2, get

```
kubectl get all
```

## step

```
kubectl apply -f mysql/mysql-sts-svc.yaml
kubectl apply -f redis/redis-sts-svc.yaml
kubectl apply -f user/user-deploy-svc.yaml

kubectl get all
kubectl delete svc,deploy,pod,sts --all
```
