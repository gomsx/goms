APP=redistest
TAG=v1.4.2
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(NAME_SPACE)/$(APP):$(TAG)

all: docker push

docker:
	docker build -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
