IMAGE_NAME:=redistest
IMAGE_VERSION:=v3
DOCKER_IMAGE:=$(DOCKER_USERNAME)/$(IMAGE_NAME):$(IMAGE_VERSION)

all: docker push

docker: 
	docker build -t $(IMAGE_NAME) -f dockerfile ../../
	docker images | grep $(IMAGE_NAME)
push: 
	docker tag $(IMAGE_NAME) $(DOCKER_IMAGE)
	docker push $(DOCKER_IMAGE)
