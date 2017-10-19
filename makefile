.PHONY: all build deploy local

all: local

build:
	cd build && ./build.sh

deploy:
	cd dropstack && dropstack deploy

local:
	go build -ldflags "-X main.VersionNumber=`git describe --always --tags --dirty="*"`" -o groupbox ./src/
