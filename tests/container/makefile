all:docker push 

docker: 
	cd ./mysqltest;make docker 
	cd ./redistest;make docker 
	cd ./clienttest;make docker 

push: 
	cd ./mysqltest;make push
	cd ./redistest;make push
	cd ./clienttest;make push

