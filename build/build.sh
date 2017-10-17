#!/bin/sh
mkdir ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw), für den Namen
sed s!'$stackletname'!$GROUPBOX_STACKLETNAME! < template.dropstack.json > ../dropstack/.dropstack.json

cp dropstackdemo.html ../dropstack/index.html

cd ../dropstack

echo "Jetzt das Deployment starten mit: dropstack deploy"