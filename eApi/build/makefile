APP:=eApi
REPO_NAME:=goms-eapi
TAG:=v1.2.0
NAME_SPACE:=$(DOCKER_USERNAME)
IMAGE:=$(NAME_SPACE)/$(REPO_NAME):$(TAG)

all:test compile docker push clean

test:
	go test ../... -v -gcflags=-l -count=1
compile:
	go build -o $(APP) ../cmd
	ls -l ./$(APP) ../
	chmod +x ./$(APP)
docker:
	docker build -t $(REPO_NAME) -f dockerfile  ../
	docker tag $(REPO_NAME) $(IMAGE)
push:
	docker push $(IMAGE)
clean:
	rm -f $(APP)
