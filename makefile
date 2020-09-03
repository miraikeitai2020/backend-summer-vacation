GOCMD=go
GORUN=$(GOCMD) run
DOCKER=docker
DOCKERBUILD=$(DOCKER) build
DOCKERRUN=$(DOCKER) run

all:
	$(GORUN) pkg/server/server.go
docker-build:
	$(DOCKERBUILD) ./ -t miraikeitai2020/summer:1.0.0
docker-run:
	$(DOCKERRUN) -d -p 8080:8080 --name summer miraikeitai2020/summer:1.0.0
docker-clean:
	$(DOCKER) stop summer
	$(DOCKER) rm summer
	$(DOCKER) rmi miraikeitai2020/summer:1.0.0
service-build:
	docker-compose build
	docker-compose run -d
service-clean:
	docker-compose stop