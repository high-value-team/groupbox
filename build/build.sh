#!/bin/sh

mkdir -p ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw)
echo "create .dropstack.json"
sed s!'GROUPBOX_STACKLET_NAME'!${GROUPBOX_STACKLET_NAME}! < template.dropstack.json > ../dropstack/.dropstack.json

# go executable
echo "go build"
export GOOS=linux
export GOARCH=amd64
go build -ldflags "-X main.VersionNumber=$VERSION_NUMBER" -o ../dropstack/groupbox ../src

# docker image
echo "create Dockerfile"
cp template.Dockerfile Dockerfile
sed -i .backup "s|MONGODB_URL|${MONGODB_URL}|g" Dockerfile
rm Dockerfile.backup
mv Dockerfile ../dropstack/Dockerfile


echo "Jetzt nach ../dropstack wechseln und deployen mit dropstack deploy"