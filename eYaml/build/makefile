APP=eyaml
TAG=v1.1.0
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(CHUB)/$(NAME_SPACE)/$(APP):$(TAG)

all:compile docker push clean

debug:
	go build -o $(APP) ../cmd
	ls -l ./$(APP) ../
	chmod +x ./$(APP)
compile:
	go build -o $(APP) ../cmd
docker:
	docker build -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
clean:
	rm -f $(APP)
