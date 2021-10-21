APP=./bin/app

test:
	go test  ./... -cover

build:
	go build -o $(APP) ./cmd/app/main.go

run_im:
	$(APP)

run_db:
	$(APP) -db

fmt:
	go fmt ./...
	#go imports -l ./
	go mod tidy

generate:
	go generate ./...
