#!/bin/sh

mkdir -p ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw)
echo "create .dropstack.json"
sed s!'GROUPBOX_STACKLET_NAME'!${GROUPBOX_STACKLET_NAME}! < template.dropstack.json > ../dropstack/.dropstack.json

# go executable
echo "create go executable"
export GOOS=linux
export GOARCH=amd64
export VERSION_NUMBER=`git describe --always --tags --dirty="*"`
go build -ldflags "-X main.VersionNumber=$VERSION_NUMBER" -o ../dropstack/groupbox ../src

# docker image
echo "create Dockerfile"
cp template.Dockerfile Dockerfile
sed -i .backup "s|MONGODB_URL|${MONGODB_URL}|g" Dockerfile
sed -i .backup "s|GROUPBOX_ROOT_URI|${GROUPBOX_ROOT_URI}|g" Dockerfile
sed -i .backup "s|SMTP_USERNAME|${SMTP_USERNAME}|g" Dockerfile
sed -i .backup "s|SMTP_PASSWORD|${SMTP_PASSWORD}|g" Dockerfile
sed -i .backup "s|SMTP_NO_REPLY_EMAIL|${SMTP_NO_REPLY_EMAIL}|g" Dockerfile
sed -i .backup "s|SMTP_SERVER_ADDRESS|${SMTP_SERVER_ADDRESS}|g" Dockerfile

cp deploy.sh ../dropstack

rm Dockerfile.backup
mv Dockerfile ../dropstack/Dockerfile

echo "Jetzt cd ../dropstack und deployen mit ./deploy.sh"

