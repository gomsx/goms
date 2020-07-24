# istio

## label
```
kubectl label namespace test-goms istio-injection=enabled
kubectl get namespace -L istio-injection
```

## apply

```
kubectl apply -n test-goms -f user-vs.yaml
kubectl apply -n test-goms -f user-dr.yaml
``` 
![](https://raw.githubusercontent.com/aivuca/goms/master/docs/pics/eistio-user-vs-dr.png?token=AQC6BPTQSI4MNNIPIS4ZAA27DJRA2)
## 注意
- 要给对象 label istio-injection=enabled,才能使用 istio
- label 后 pod 要重启才会注入 pause
