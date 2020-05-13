## 部署

namespace
```
kubectl create namespace goms-ek8s
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