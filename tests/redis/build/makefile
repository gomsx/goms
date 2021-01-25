REPO:=goms-redistest
TAG:=v1.4.2
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(CHUB)/$(NAME_SPACE)/$(REPO):$(TAG)

all: docker push

docker:
	docker build -t $(REPO) -f dockerfile ../
	docker tag $(REPO) $(IMAGE)
push:
	docker push $(IMAGE)
