APP_IMAGE:=eRedis
DOCKER_IMAGE:=eredis


all:compile docker push clean

compile:
	go build -o $(APP_IMAGE) ../cmd
	ls -l ./$(APP_IMAGE) ../
	chmod +x ./$(APP_IMAGE) 
docker: 
	docker build -t $(DOCKER_IMAGE) -f dockerfile  ../    
	docker images
push: 
	docker tag $(DOCKER_IMAGE) $(DOCKER_USERNAME)/$(DOCKER_IMAGE)
	docker push $(DOCKER_USERNAME)/$(DOCKER_IMAGE)
clean:
	rm $(APP_IMAGE)






