unit-test:
	go test -cover ./internal/... -covermode=atomic

.PHONY: unit_test

integration-test:
	@docker-compose -f ./docker-compose.test.yml up -d
	@sleep 1 && \
		MONGO_URL="mongodb://127.0.0.1:27017/" \
		MONGO_TODO_DB=TodoDbTest \
        MONGO_CONNECTION_TIMEOUT=20 \
        MONGO_MAX_POOL_SIZE=10 \
		JAEGER_DISABLED=true \
		go test -coverpkg ./... ./test/...
	@docker-compose -f ./docker-compose.test.yml down

.PHONY: integration_test

lint:
	golangci-lint run --fix -v

.PHONY: lint

swag:
	swag init -g ./cmd/api/main.go -o ./docs

wire-mongo:
	wire ./internal/wired/mongo.go

wire-redis:
	wire ./internal/wired/redis.go

build:
	go build -o ./todoapi ./cmd/api

run:
	make build
	./todoapi

clean:
	rm -rf ./todoapi

docker-mongo-start:
	BACKEND_URL=http://localhost:3080 \
	FRONTEND_URL=http://localhost:8080 \
	FRONTEND_BUILD_TARGET=prod \
	BACKEND_BUILD_TARGET=prod \
	docker-compose up --build -d

docker-mongo-stop:
	docker-compose down

docker-redis-start:
	docker-compose -f ./docker-compose.redis.yml up --build -d

docker-redis-stop:
	docker-compose -f ./docker-compose.redis.yml down

format:
	go fmt ./internal/...
