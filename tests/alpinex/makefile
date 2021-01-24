REPO:=goms-alpinex
TAG:=v0.1.0
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(NAME_SPACE)/$(REPO):$(TAG)

all: docker push

docker:
	docker build -t $(REPO) -f dockerfile .
	docker tag $(REPO) $(IMAGE)
push:
	docker push $(IMAGE)
