APP_NAME:=eFlag
IMAGE_NAME:=goms-eflag
IMAGE_VERSION:=v1
DOCKER_IMAGE:=$(DOCKER_USERNAME)/$(IMAGE_NAME):$(IMAGE_VERSION)

all:compile docker push clean

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






