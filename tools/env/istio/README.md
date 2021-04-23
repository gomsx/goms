
## Istio

```
bash install-istio.sh
```

## 插件

file
```
$ ls /usr/local/istio-1.9.1/samples/addons/
extras  grafana.yaml  jaeger.yaml  kiali.yaml  prometheus.yaml  README.md
```

deploy
```
bash deploy.sh
```

patch
```
bash patch-xxx.sh
```

## 访问

http://120.79.1.69:31131/kiali  
http://120.79.1.69:31133/grafana  
http://120.79.1.69:31130/graph  
