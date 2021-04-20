# dev 

# apply kustomize

```
kubectl apply -k ./
kubectl get -k ./
kubectl describe -k ./
```

# rolling update user-deploy v2

```
kubectl apply -f user/user-deploy-v2-rollingupdate.yaml
```

# deploy user-deploy v2

```
kubectl apply -f user/user-deploy-v2.yaml
```

# patch user-deploy v1

```
kubectl patch deployment.apps/user-deploy --patch "$(cat patch/user-deploy-v1.yaml)"
```

# patch client-deploy

```
# replicas
kubectl patch deploy client-deploy -p '{"spec": {"replicas": 2}}'

# command
kubectl patch deploy client-deploy -p '{"spec": {"template": {"spec": {"containers": [{"name": "client","command": ["sh", "/bash/test_http.sh","100000","v1","user-svc"]}]}}}}'
```
