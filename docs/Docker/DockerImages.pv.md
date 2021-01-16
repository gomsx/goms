# iamges 管理

## 退出容器

exit
Ctrl+P+Q
## 所有有用的 images 都在 run 时，可以删除无用的不在 run 的 images

```
docker images -q | xargs docker rmi
```

## 长时间运行，硬盘爆满时，删除容量暴涨的 docker

```
du -d 1 -hm /var/lib/docker/containers | sort -t $'\t' -k 1rn,1
```

## 搜索特定 docker 并管理

```
docker ps | grep mysql | awk '{ print $1 }' | xargs docker stop
docker ps | grep mysql | awk '{ print $1 }' | xargs docker rm -f
docker ps | grep mysql | awk '{ print $1 }' | xargs docker logs
```

## Docker build 提高 apt-get 速度

把 /etc/apt/sources.list 文件中的 archive.ubuntu.com 替换为 mirrors.aliyun.com. 提高速度.
在 Dockerfile 添加两行语句：
```
RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list 
RUN apt-get clear
```
