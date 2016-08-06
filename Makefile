.PHONY: build vendor_clean vendor_get vendor_update rice

GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

default: build

rice:
	$(PWD)/_vendor/bin/rice embed-go

build: vendor_update \
	rice
	go build -v -o ./bin/potemkin

vendor_clean:
	rm -rf ./_vendor/src

vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get \
	github.com/boltdb/bolt \
	github.com/GeertJohan/go.rice \
	github.com/GeertJohan/go.rice/rice \
	github.com/satori/go.uuid

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
	&& rm -rf `find ./_vendor/src -type d -name .svn`
