all:docker push

docker:
	cd ../mysql/build;make docker
	cd ../redis/build;make docker
	cd ../client/build;make docker
push:
	cd ../mysql/build;make push
	cd ../redis/build;make push
	cd ../client/build;make push
