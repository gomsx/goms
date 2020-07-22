# test

# apply kustomize
```
kubectl apply -k ./
kubectl get -k ./
kubectl describe -k ./
```

# rolling update user-deploy v2
```
kubectl apply -f /other/user-deploy-v2-rollingupdate.yaml
```

# deploy user-deploy v2
```
kubectl apply -f /other/user-deploy-v2.yaml
```