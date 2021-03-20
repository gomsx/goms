
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
k apply -f mysql-sts-svc.yaml
k get all,pv,pvc
k delete deploy,pod,sts,svc,pv,pvc --all &

k apply -f mysql-deploy-svc.yaml
k get all,pv,pvc
k delete deploy,pod,sts,svc,pv,pvc --all &
```