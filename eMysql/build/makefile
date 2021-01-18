APP:=eMysql
REPO_NAME:=goms-emysql
TAG:=v1.1.0
NAME_SPACE:=$(DOCKER_USERNAME)
IMAGE:=$(NAME_SPACE)/$(REPO_NAME):$(TAG)

all:compile docker push clean

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
