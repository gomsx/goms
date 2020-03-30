APP_NAME:=eTest
IMAGE_NAME:=goms-etest
IMAGE_VERSION:=v1
DOCKER_IMAGE:=$(DOCKER_USERNAME)/$(IMAGE_NAME):$(IMAGE_VERSION)

all:test compile docker push clean

test:
	chmod +x ../internal/dao/testdata/docker.sh
	../internal/dao/testdata/docker.sh
	go test -v ../... 
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







