MAIN_PATH=cmd/cxp-api-starter/main.go   # Contains main.go file
BINARY_FILE=binary 		      			# Binary file name

# Builds the project
#
# Removes binary
clean:
	go clean
	rm -rf bin/${BINARY_FILE}

# Builds main.go
build: clean
	go build -o bin/${BINARY_FILE} ${MAIN_PATH}

# First removes the binary file than rebuilds main.go and runs the binary
run: build
	./bin/${BINARY_FILE}

# Development
up-dev-down:
	docker compose -f docker-compose.dev.yml down
up-dev: up-dev-down
	docker compose -f docker-compose.dev.yml --env-file .env.dev up --build