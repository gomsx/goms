
## debug

1, deploy app

```
kubectl apply -k .
```

2, deploy pv,cm

```
kubectl apply -k ./debug
```

3, get

```
kubectl get all,pv,cm
```

## step

```
kubectl apply -f debug/mysql-pv/pv-local.yaml
kubectl apply -f mysql/mysql-sts-svc.yaml


kubectl apply -f debug/redis-pv/pv-local.yaml
kubectl apply -f redis/redis-sts-svc.yaml

kubectl apply -k user/user-cm
kubectl apply -f user/user-deploy-svc.yaml

kubectl get all,pv,pvc,cm

kubectl delete deploy,pod,sts,svc,pv,pvc --all
```