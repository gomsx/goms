APP=redistest
TAG=v1.4.4
IMAGE=$(DHUB)/$(DHUB_NAMESPACE)/$(APP):$(TAG)

all: docker push

docker:
	docker build -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
