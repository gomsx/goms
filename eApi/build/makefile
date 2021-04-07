APP=eapi
TAG=v1.3.0
NAME_SPACE=$(DHUB_NAMESPACE)
IMAGE=$(DHUB)/$(NAME_SPACE)/$(APP):$(TAG)

RIMAGE=$(RHUB)/$(RSPACE)/$(APP):$(TAG)

all:test compile docker push clean

test:
	go test -gcflags=-l -count=1 ../...
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
release:
	docker tag $(APP) $(RIMAGE)
	docker push $(RIMAGE)
clean:
	rm -f $(APP)
