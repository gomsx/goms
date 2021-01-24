REPO:=goms-redistest
TAG:=v1.4.0
NAME_SPACE:=$(DOCKER_USERNAME)
IMAGE:=$(NAME_SPACE)/$(REPO):$(TAG)

all: docker push

docker:
	docker build -t $(REPO) -f dockerfile ../
	docker tag $(REPO) $(IMAGE)
push:
	docker push $(IMAGE)
