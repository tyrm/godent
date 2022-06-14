PROJECT_NAME=godent

docker-pull:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml pull

docker-restart: docker-stop docker-start

docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

.PHONY: docker-pull docker-restart docker-start docker-stop tidy vendor
