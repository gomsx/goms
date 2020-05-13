

https://blog.csdn.net/networken/article/details/85607593
https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/
https://github.com/kubernetes/dashboard
https://www.jianshu.com/p/be2a12a8bc0b


## 部署
```
kubectl apply -f 2_0_0/recommended-my.yaml
```

## 调试
```
kubectl get ns
kubectl get rs,pod,deploy,svc,ep -n kubernetes-dashboard
```

## 生成 token， 用于 dashboard ui 登录
```
./get_token.sh
```