#!/bin/sh
mkdir ../dropstack

# GROUPBOX_STACKLETNAME in ~/.bashrc definiert (ralfw)
#sed -i "s/$stackletname/$GROUPBOX_STACKLETNAME" template.dropstack.json
sed s!'$stackletname'!$GROUPBOX_STACKLETNAME! < template.dropstack.json > ../dropstack/.dropstack.json

cp dropstackdemo.html ../dropstack/index.html