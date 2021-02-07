APP=clienttest
TAG=v1.5.1
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(CHUB)/$(NAME_SPACE)/$(APP):$(TAG)

all: docker push

docker:
	cp /etc/timezone ./
	docker build --no-cache -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
	rm ./timezone
push:
	docker push $(IMAGE)
