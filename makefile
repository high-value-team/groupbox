.PHONY: all test build docker deploy run deploy_local deploy_docker deploy_dropstack

all: test

test:
	echo "execute tests"
	cd src/backend && go test -v `(go list ./... | grep -v vendor)`

build:
	echo "build go executable"
	go build -ldflags "-X main.VersionNumber=`git describe --always --tags --dirty="*"`" -o groupbox ./src/

docker:
	echo "build in build directory"
	cd build && ./build.sh

deploy:
	echo "deploy to dropstack"
	cd dropstack && ./deploy.sh

run:
	echo "run docker container on localhost:9090"
	cd dropstack && docker build -t groupbox .
	cd dropstack && docker run --rm -ti -p 9090:80 groupbox

#
# test, build and run on specified plattform (local, docker on localhost, docker on dropstack)
#

deploy_local: test build
	./groupbox

deploy_docker: test docker run
	echo "run tests build docker container and run docker container"

deploy_dropstack: test docker deploy
	echo "build docker container and run docker container"

