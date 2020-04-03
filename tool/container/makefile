
# 查看工程目录
PRODIR:=$(shell cd ./;pwd)# 工程目录
$(warning $(PRODIR))   

# 查看环境变量
$(warning $(HOME))
$(warning $(PATH))
$(warning $(shell which go))


all:docker push 

docker: 
	cd $(PRODIR)/mysql;make docker 
	cd $(PRODIR)/mysqltest_v1;make docker 
	cd $(PRODIR)/mysqltest_v2;make docker 

push: 
	cd $(PRODIR)/mysql;make push
	cd $(PRODIR)/mysqltest_v1;make push
	cd $(PRODIR)/mysqltest_v2;make push