GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BASEPATH := $(shell pwd)
BUILDDIR=$(BASEPATH)/dist
GOGINDATA=go-bindata
GOARCH=amd64

KM_SERVER_NAME=kmpp-middleware
KM_CONFIG_DIR=etc/middleware
KM_BIN_DIR=usr/local/bin
KM_DATA_DIR=usr/local/lib/middleware


GOPROXY="https://goproxy.cn,direct"


build_server_linux:
	GOOS=linux GOARCH=$(GOARCH)  $(GOBUILD) -o $(BUILDDIR)/$(KM_BIN_DIR)/$(KM_SERVER_NAME) main.go
	mkdir -p $(BUILDDIR)/$(KM_CONFIG_DIR) && cp -r  $(BASEPATH)/conf/app.yaml $(BUILDDIR)/$(KM_CONFIG_DIR)
	mkdir -p $(BUILDDIR)/$(KM_DATA_DIR)
	cp -r  $(BASEPATH)/migration $(BUILDDIR)/$(KM_DATA_DIR)


docker_server:
	docker build -t harbor.dev.rdev.tech/kmpp/kmpp-middleware:master --build-arg GOPROXY=$(GOPROXY) --build-arg GOARCH=$(GOARCH) .

clean:
	rm -rf $(BUILDDIR)
	$(GOARCH) $(GOBUILD) clean
