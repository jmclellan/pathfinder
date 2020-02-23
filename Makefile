VER = 0.0.1
TAG = $(shell git rev-parse --short HEAD)
GO_EXE_NAME = pathfinder-$(VER)-$(TAG)
OUTPUT_DIR = build-dir

export GOPATH = $(abspath .)

get-dependencies:
	go get github.com/gitchander/permutation
	go get github.com/gorilla/mux
	go get github.com/umahmood/haversine

build-go-server:
	go build -o $(OUTPUT_DIR)/$(GO_EXE_NAME) -x src/main.go

clean:
	rm -rf $(OUTPUT_DIR)