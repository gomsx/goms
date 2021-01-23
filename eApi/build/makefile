APP:=eApi
REPO:=goms-eapi
TAG:=v1.2.0
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(NAME_SPACE)/$(REPO):$(TAG)

all:test compile docker push clean

test:
	go test ../... -v -gcflags=-l -count=1
debug:
	go build -o $(APP) ../cmd
	ls -l ./$(APP) ../
	chmod +x ./$(APP)
compile:
	go build -o $(APP) ../cmd
docker:
	docker build -t $(REPO) -f dockerfile ../
	docker tag $(REPO) $(IMAGE)
push:
	docker push $(IMAGE)
clean:
	rm -f $(APP)
