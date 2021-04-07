APP=mysqltest
TAG=v2.2.4
IMAGE=$(DHUB)/$(DHUB_NAMESPACE)/$(APP):$(TAG)

all: docker push

docker:
	docker build --no-cache -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
