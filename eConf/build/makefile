APP:=eConf
REPO:=goms-econf
TAG:=v1.1.0
NAME_SPACE:=$(GOMS_NAME_SPACE)
IMAGE:=$(CHUB)/$(NAME_SPACE)/$(REPO):$(TAG)

all:compile docker push clean

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
