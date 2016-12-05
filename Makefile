CWD=$(shell pwd)
PATH := $(PATH):/usr/local/go/bin/
prefix := $(DESTDIR)/usr
vardir := $(DESTDIR)/var
confdir := $(DESTDIR)/etc
NAME:=snipcart-webhook-laposte


export GOPATH:=${CWD}
export GO15VENDOREXPERIMENT=1

all: build

build:
	go install $(NAME)

check:
	go test $(NAME)/...

clean:
	go clean
	rm bin/*

install: build
	install -d ${prefix}/bin
	install -m 755 ./bin/$(NAME) ${prefix}/bin/$(NAME)

.PHONY: build
