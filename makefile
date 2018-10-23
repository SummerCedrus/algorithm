GOPATH := $(shell pwd)
all:
	GOPATH=$(GOPATH) go install skiplist_main
