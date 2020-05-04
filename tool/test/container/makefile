
# 查看工程目录
PRODIR:=$(shell cd ./;pwd)# 工程目录
$(warning $(PRODIR))   

# 查看环境变量
$(warning $(HOME))
$(warning $(PATH))
$(warning $(shell which go))


all:docker push 

docker: 
	cd $(PRODIR)/mysqltest;make docker 

push: 
	cd $(PRODIR)/mysqltest;make push