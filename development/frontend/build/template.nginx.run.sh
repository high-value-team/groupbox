#!/bin/sh

echo
echo "Replacing environment variables"
sh /replace.sh /usr/share/nginx/html/static/**/*
if [[ $? != 0 ]]
then
    echo "FAILED"
    exit 0;
fi
echo

echo "starting NGINX"
nginx -g "daemon off;"
