## 安装
```
git clone https://github.com/wg/wrk.git wrk
cd wrk
make
sudo install ./wrk /usr/local/bin
```

## 使用
```
wrk -t4 -c100 -d5s http://localhost:8080/ping
```

