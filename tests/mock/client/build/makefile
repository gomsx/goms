APP=clienttest
TAG=v1.5.5
IMAGE=$(DHUB)/$(DHUB_NAMESPACE)/$(APP):$(TAG)

all: docker push

docker:
	cp /etc/timezone ./
	docker build --no-cache -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
	rm ./timezone
push:
	docker push $(IMAGE)
