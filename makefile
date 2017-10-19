.PHONY: all build deploy local run

all: local

build:
	cd build && ./build.sh

deploy:
	cd dropstack && dropstack deploy

local:
	go build -ldflags "-X main.VersionNumber=`git describe --always --tags --dirty="*"`" -o groupbox ./src/

run:
	cd dropstack && docker build -t groupbox .
	cd dropstack && docker run --rm -ti -p 8080:80 groupbox

