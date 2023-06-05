# dev stage
FROM golang:1.17-alpine as dev

WORKDIR /app

# Needed for VS Code Remote Development
RUN apk add bash openssh musl libgcc libstdc++ git curl wget bash ca-certificates gcompat openssh-client

# Copy mod and sum files
COPY go.mod go.sum ./

# Install swaggo and wire
RUN go get github.com/swaggo/swag/cmd/swag
RUN go get github.com/google/wire/cmd/wire

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Generate Swagger document
RUN go run github.com/swaggo/swag/cmd/swag init -g ./cmd/api/main.go -o ./docs

# Generate dependencies by wire
RUN go run github.com/google/wire/cmd/wire ./internal/wired/mongo.go

# Build the Go api
RUN go build -o ./todoapi ./cmd/api

EXPOSE 8080

# Run the executable
CMD ["todoapi"]

# prod stage
FROM dev as prod

COPY --from=dev /app/todoapi /go/bin/todoapi

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["todoapi"]
