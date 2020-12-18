# 测试脚本使用

## http and grpc

```
./test.sh

./test.sh "0.01" "v1"

./test.sh "0.01" "v1" "localhost"

./test.sh "0.01" "v1" "localhost" "8080"

./test.sh "0.01" "v1" "localhost" "8080" "50051"

./test.sh "0.01" "v1" "localhost" "" "50051"

./test.sh "0.01" "v1" "" "" "50051"
```

## http
```
./testhttp.sh

./testhttp.sh "0.01" "v1"

./testhttp.sh "0.01" "v1" "localhost"

./testhttp.sh "0.01" "v1" "localhost" "8080"

./testhttp.sh "0.01" "v1" "localhost" ""

./testhttp.sh "0.01" "v1" "" ""
```

## grpc

```
./testgrpc.sh

./testgrpc.sh "0.01" "v1"

./testgrpc.sh "0.01" "v1" "localhost"

./testgrpc.sh "0.01" "v1" "localhost" "50051"

./testgrpc.sh "0.01" "v1" "localhost" ""

./testgrpc.sh "0.01" "v1" "" ""
```

## 注意

测试无版本API时，用 "" 代替 "v1" 参数,如:
```
./test.sh "0.01" "" "localhost"
```

