APP=econf
TAG=v1.1.0

IMAGE=$(DHUB_NAMESPACE)/$(APP):$(TAG)

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
