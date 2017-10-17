#!/bin/sh
mkdir -p ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw)
sed s!'$stackletname'!$GROUPBOX_STACKLETNAME! < template.dropstack.json > ../dropstack/.dropstack.json


# go executable
GOOS=linux GOARCH=amd64 go build -o ../dropstack/groupbox ../src

# docker image
cp Dockerfile ../dropstack/Dockerfile



