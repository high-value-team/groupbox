.PHONY: all dropstack local

all: dropstack

dropstack:
	cd build && ./build.sh
	cd dropstack && dropstack deploy

local:
	go build -ldflags "-X main.VersionNumber=0.0.0" -o groupbox ./src/
