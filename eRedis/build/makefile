APP_NAME:=eRedis
REPO_NAME:=goms-eredis
TAG:=v1.1.0
NAME_SPACE:=$(DOCKER_USERNAME)
IMAGE:=$(NAME_SPACE)/$(REPO_NAME):$(TAG)

all:compile docker push clean

compile:
	go build -o $(APP_NAME) ../cmd
	ls -l ./$(APP_NAME) ../
	chmod +x ./$(APP_NAME)
docker:
	docker build -t $(REPO_NAME) -f dockerfile  ../
	docker tag $(REPO_NAME) $(IMAGE)	
push:
	docker push $(IMAGE)
clean:
	rm -f $(APP_NAME)

