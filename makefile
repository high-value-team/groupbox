.PHONY: all build deploy local

all: local

build:
	cd build && ./build.sh

deploy:
	cd dropstack && dropstack deploy


local:
	go build -ldflags "-X main.VersionNumber=1.2.3" -o groupbox ./src/

run:
	cd dropstack && docker build -t groupbox .
	cd dropstack && docker run --rm -ti -p 8080:80 groupbox