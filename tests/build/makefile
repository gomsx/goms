all:docker push

docker:
	cd ../alpinex;make docker
	cd ../mysql/build;make docker
	cd ../redis/build;make docker
	cd ../client;make docker

push:
	cd ../alpinex;make push
	cd ../mysql/build;make push
	cd ../redis/build;make push
	cd ../client;make push
