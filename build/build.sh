#!/bin/sh

mkdir -p ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw)
echo "create .dropstack.json"
sed s!'$stackletname'!$GROUPBOX_STACKLETNAME! < template.dropstack.json > ../dropstack/.dropstack.json

# go executable
echo "go build"
export GOOS=linux
export GOARCH=amd64
export VERSION_NUMBER=0.0.3
go build -ldflags "-X main.VersionNumber=$VERSION_NUMBER" -o ../dropstack/groupbox ../src

# docker image
echo "create Dockerfile"
cp Dockerfile ../dropstack/Dockerfile

echo "Jetzt nach ../dropstack wechseln und deployen mit dropstack deploy"