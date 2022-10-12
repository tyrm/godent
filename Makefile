PROJECT_NAME=godent

build-snapshot:
	goreleaser release --snapshot --rm-dist

new-migration: export BUN_TIMESTAMP=$(shell date +%Y%m%d%H%M%S | head -c 14)
new-migration:
	touch internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go
	cat internal/db/bun/migrations/migration.go.tmpl > internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go
	sed -i "s/\"CHANGEME\"/\"${BUN_TIMESTAMP}\"/g" internal/db/bun/migrations/${BUN_TIMESTAMP}_new.go

check:
	golangci-lint run

check-fix:
	golangci-lint run --fix

docker-pull:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml pull

docker-restart: docker-stop docker-start

docker-start:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml up -d

docker-stop:
	docker-compose --project-name ${PROJECT_NAME} -f deployments/docker-compose-test.yaml down

fmt:
	@echo formatting
	@go fmt $(shell go list ./... | grep -v /vendor/)

test:
	go test -cover ./...

tidy:
	go mod tidy

vendor: tidy
	go mod vendor

.PHONY: build-snapshot check check-fix docker-pull docker-restart docker-start docker-stop fmt new-migration test tidy vendor
