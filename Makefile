docker.routing.build:
	docker build -t mbenz-routing -f ./mbenz_planning/Dockerfile ./mbenz_planning

docker.poc.build:
	docker build -t mbenz-routing -f ./mbenz_poc/Dockerfile ./mbenz_poc

run:
	docker-compose up -d

stop:
	docker-compose stop

all: docker.routing.build docker.poc.build
	docker-compose up -d

