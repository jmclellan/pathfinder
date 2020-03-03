VER = 0.0.1
TAG = $(shell git rev-parse --short HEAD)
GO_EXE_NAME = pathfinder-$(VER)-$(TAG)
OUTPUT_DIR = build-dir

# MAKE_EXE = $(which make)
# hardcoded path ties to provision set up in Vagrantfile
GO_EXE = /usr/local/go/bin/go
export PATH=$PATH:$(GO_EXE)

export GOROOT = /usr/local/go
export GOPATH = $(abspath .)

dev-pathfinder:
	/usr/bin/make clean
	/usr/bin/make get-dependencies
	/usr/bin/make make build-go-server
	$(make run-go-server)
	$(make start-node-server)
	$(make start-nginx)


get-dependencies:
	$(GO_EXE) get github.com/gitchander/permutation
	$(GO_EXE) get github.com/gorilla/mux
	$(GO_EXE) get github.com/umahmood/haversine

build-go-server:
	$(GO_EXE) build -o $(OUTPUT_DIR)/$(GO_EXE_NAME) -x src/main.go

run-go-server: #add target
	./$(OUTPUT_DIR)/$(GO_EXE_NAME)

clean:
	/usr/bin/rm -rf $(OUTPUT_DIR)

start-node-server: # not fixed
	/usr/bin/npm start --prefix /vagrant_data/src/react-code/

start-nginx:
	# use supervisorctl to manage this 
	/usr/sbin/nginx -c /vagrant_data/src/nginx/nginx.conf

test:
	echo "hey this is working $(which rm) test test"