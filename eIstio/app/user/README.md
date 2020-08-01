# istio

## label
```
kubectl label namespace test-goms istio-injection=enabled
kubectl get namespace -L istio-injection
```

## apply

```
kubectl apply -n test-goms -k app/user
``` 

## 注意
- 要给对象 label istio-injection=enabled,才能使用 istio
- label 后 pod 要重发布才会注入 pause

