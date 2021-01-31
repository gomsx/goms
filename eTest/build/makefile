APP=etest
TAG=v1.1.3
NAME_SPACE=$(GOMS_NAME_SPACE)
IMAGE=$(NAME_SPACE)/$(APP):$(TAG)

all:test compile docker push clean

test:
	go test -v -gcflags=-l -count=1 ../...
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
