APP=mysqltest
TAG=v2.2.1
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(NAME_SPACE)/$(APP):$(TAG)

all: docker push

docker:
	docker build --no-cache -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
