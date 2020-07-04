# eDocker

容器化技术,使用 docker.

## docker 主要对象

- Docker image  -  数据层的集合  
- Docker file  -  定义各种镜像层  
- Docker engine  -  构建和管理镜像  
- Docker registry  -   存储镜像的地方  
- Container  -  运行中的镜像实例  

```
engine: build/run/rmi/push/pull/...

              inspect       commit 
inspectdata <---------+   +----------+
         :            |   |          |
         :            |   |          |
         V   build    |   V   run    |
dockerfile ---------> image -------> container
                      |   ^
                      |   |
                 push |   |  pull 
                      V   |
                     registry
```
## docker image  

* 被生产 docker build
* 被消费 docker run
* 中间环节  
       * 入库 docker push  
       * 出库 docker pull  
       * 丢弃 docker rmi  
       * ...

## docker file  

* dockerfile 用于定义 dockerimage.
* dockerimage 被 build 生产出来,被 run 消费成 container 实例.
* container 实例是一个进程.用户可以定义的部分:

进程    | dockerfile
-------|------------  
所属用户|  USER/...  
文件系统|  ADD/COPY/WORKDIR...  
运行指令|  CMD/...   

docker build -t image_name -f dockerfile_path  build_root_path
> build 时指定一个构建的根路径,此处的 build_root_path,COPY/ADD时,会在这个路径下查找.
## 其他容器技术  

[gvisor](https://github.com/google/gvisor)  
[Kata Containers](https://github.com/kata-containers/runtime)  

## 注意事项

1, 把用户加入docker组后就不用 sudo 来执行 docker 命令
```
//设置
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo systemctl restart docker

//确认
cat /etc/group | grep docker | grep ${USER}
docker version
```

2, 有些 docker registry 在国外，要设置国内镜像
```
cat /etc/docker/daemon.json
{
        "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
```

