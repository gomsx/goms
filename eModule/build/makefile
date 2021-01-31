APP=emodule
TAG=v1.1.0
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(NAME_SPACE)/$(APP):$(TAG)

all:compile docker push clean

debug:
	go build -o $(APP) ../cmd
	ls -l ./$(APP) ../
	chmod +x ./$(APP)
compile:
	go build -o $(APP) ../cmd
docker:
	docker build --no-cache -t $(APP) -f dockerfile ../
	docker tag $(APP) $(IMAGE)
push:
	docker push $(IMAGE)
clean:
	rm -f $(APP)
