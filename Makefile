
provision:
	@echo "Provisioning WOOT Cluster"	
	bash scripts/provision.sh

woot-build:
	@echo "Building WOOT Docker Image"	
	docker build -t woot -f Dockerfile .

woot-run:
	@echo "Running Single WOOT Docker Container"
	docker run -p 8080:8080 -d woot

info:
	echo "WOOT Cluster Nodes"
	docker ps | grep 'woot'
	docker network ls | grep woot_network

clean:
	@echo "Cleaning WOOT Cluster"
	docker ps -a | awk '$$2 ~ /woot/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm woot_network

build:
	@echo "Building WOOT Server"	
	go build -o bin/woot main.go

fmt:
	@echo "go fmt WOOT Server"	
	go fmt ./...

test:
	@echo "Testing WOOT"	
	go test -v --cover ./...