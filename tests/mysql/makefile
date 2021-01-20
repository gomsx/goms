REPO_NAME:=goms-mysqltest
TAG:=v2.1.1
NAME_SPACE:=$(DOCKER_USERNAME)
IMAGE:=$(NAME_SPACE)/$(REPO_NAME):$(TAG)

all: docker push

docker:
	docker build -t $(REPO_NAME) -f dockerfile .
	docker tag $(REPO_NAME) $(IMAGE)
push:
	docker push $(IMAGE)
