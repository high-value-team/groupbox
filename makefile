.PHONY: all build deploy local

all: local

build:
	cd build && ./build.sh

deploy:
	cd dropstack && dropstack deploy

local:
	go build -ldflags "-X main.VersionNumber=0.0.0" -o groupbox ./src/
