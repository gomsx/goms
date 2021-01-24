REPO:=goms-mysqltest
TAG:=v2.2.1
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(NAME_SPACE)/$(REPO):$(TAG)
IMAGEX:=$(NAME_SPACE)/$(REPO)

all: docker push

docker:
	docker build --no-cache -t $(REPO) -f dockerfile ../
	docker tag $(REPO) $(IMAGE)
	docker tag $(REPO) $(IMAGEX)
push:
	docker push $(IMAGE)
	docker push $(IMAGEX)
