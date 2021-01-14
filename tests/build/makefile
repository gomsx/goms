all:docker push

docker:
	cd ../mysql;make docker
	cd ../redis;make docker
	cd ../client;make docker

push:
	cd ../mysql;make push
	cd ../redis;make push
	cd ../client;make push
