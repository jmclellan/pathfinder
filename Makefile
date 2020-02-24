VER = 0.0.1
TAG = $(shell git rev-parse --short HEAD)
GO_EXE_NAME = pathfinder-$(VER)-$(TAG)
OUTPUT_DIR = build-dir

export GOPATH = $(abspath .)

dev-pathfinder:
	make clean
	make get-dependencies
	make build-go-server
	$(make run-go-server)
	$(make start-node-server)
	$(make start-nginx)


get-dependencies:
	go get github.com/gitchander/permutation
	go get github.com/gorilla/mux
	go get github.com/umahmood/haversine

build-go-server:
	go build -o $(OUTPUT_DIR)/$(GO_EXE_NAME) -x src/main.go

run-go-server:
	./$(OUTPUT_DIR)/$(GO_EXE_NAME)

clean:
	rm -rf $(OUTPUT_DIR)

start-node-server:
	npm start --prefix src/react-code/

start-nginx:
	sudo nginx -c /vagrant_data/src/nginx/nginx.conf

