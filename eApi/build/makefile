APP=eapi
TAG=v1.4.0
IMAGE=$(DHUB)/$(DHUB_NAMESPACE)/$(APP):$(TAG)

RIMAGE=$(RHUB)/$(RHUB_NAMESPACE)/$(APP):$(TAG)

all: test compile docker push clean

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
