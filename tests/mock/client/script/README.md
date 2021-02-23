# 测试脚本使用

## http and grpc

格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "http端口" "grpc端口"

```
./test.sh

./test.sh "0.01" "v1"

./test.sh "0.01" "v1" "100" "localhost" "8080" "50051"

./test.sh "0.01" "v1" "" "" "8080" "50051"
```

## http

格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "http端口"

```
./test-http.sh

./test-http.sh "0.01" "v1"

./test-http.sh "0.01" "v1" "100" "localhost" "8080"

./test-http.sh "0.01" "v1" "" "" "8080"

```

## grpc

格式: cmd "间隔时间(单位s)" "版本" "读操作次数" "地址" "grpc端口"

```
./test-grpc.sh

./test-grpc.sh "0.01" "v1"

./test-grpc.sh "0.01" "v1" "100" "localhost" "50051"

./test-grpc.sh "0.01" "v1" "" "" "50051"
```

## 注意

测试不带版本api时，用 "" 代替 "v1" 参数,如:

```
./test.sh "0.01" "" "localhost"
```
