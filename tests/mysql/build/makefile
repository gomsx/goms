REPO:=goms-mysqltest
TAG:=v2.2.1
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(NAME_SPACE)/$(REPO):$(TAG)

all: docker push

docker:
	docker build --no-cache -t $(REPO) -f dockerfile ../
	docker tag $(REPO) $(IMAGE)
push:
	docker push $(IMAGE)
