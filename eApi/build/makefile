APP_NAME:=eApi
IMAGE_NAME:=goms-eapi
IMAGE_VERSION:=v1.1.1
DOCKER_IMAGE:=$(DOCKER_USERNAME)/$(IMAGE_NAME):$(IMAGE_VERSION)

all:test compile docker push clean

test:
	go test ../... -v -gcflags=-l -count=1
compile:
	go build -o $(APP_NAME) ../cmd
	ls -l ./$(APP_NAME) ../
	chmod +x ./$(APP_NAME) 
docker: 
	docker build -t $(IMAGE_NAME) -f dockerfile  ../    
	docker images
push: 
	docker tag $(IMAGE_NAME) $(DOCKER_IMAGE)
	docker push $(DOCKER_IMAGE)
clean:
	rm -f $(APP_NAME)

