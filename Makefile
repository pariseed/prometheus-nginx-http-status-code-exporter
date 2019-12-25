GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get


deps:
	$(GOGET) "github.com/daniele-parise/tail"
	$(GOGET) "github.com/prometheus/client_golang/prometheus"
	$(GOGET) "github.com/prometheus/client_golang/prometheus/promhttp"



build:
	$(GOBUILD) "nginx-httpcode-exporter.go" 



install: deps build



.PHONY: deps build install
