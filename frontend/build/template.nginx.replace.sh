#!/bin/sh

# Description:
# Read "REACT_APP..." Environment Variables and replace
# with "FRONTEND_ENVIRONMENT_VARIABLE_PLACEHOLDER_REACT_APP..." in given files.
#
# Example:
# export REACT_APP_HELLO="hello world"
# echo "FRONTEND_ENVIRONMENT_VARIABLE_PLACEHOLDER_REACT_APP" > test.txt
# sh replace.sh test.txt
# cat test.txt
# => hello world
#
# usage: replace.sh static/**/*


# check command line parameter
if [ -z "$1" ]
then
	echo "please provide directory/file name"
	exit 1
fi

# extract "REACT_APP..." environment variables keys into file
env | awk -e 'BEGIN { FS="=" }
{
	if ( $1 ~ /^REACT_APP*/ )
	{
		print $1;
	}
}
' > environment_keys.txt


echo "The following ENV will be replaced:"
cat environment_keys.txt

# replace "FRONTEND_ENVIRONMENT_VARIABLE_PLACEHOLDER_..." variables
for file in "$@"
do
	if ! [ -f "$file" ]
	then
		echo "Directory/file $file does not exist!"
		exit 1
	fi

	echo "replacing in file: $file"

    # replace in file
	while read -r line
	do
		key="$line"
		prefix_key="FRONTEND_ENVIRONMENT_VARIABLE_PLACEHOLDER_$key"
		value=$(printenv $key)

		sed -i 's|\"'"$prefix_key"'\"|"'"$value"'\"|g' $file
		sed -i 's|\"'"$prefix_key"'\\\"|"'"$value"'\\\"|g' $file
	done < environment_keys.txt
done

# cleanup
rm environment_keys.txt

#
# Testing 'sed'
#

# export A=PREFIX_ROM
# export B=PREFIX
# export C=caesar
#
# echo 'hello world "PREFIX_ROM"' | sed 's|\"'"$A"'\"|"'"$C"'\"|g'
# => hello world "caesar"
#
# echo 'hello world "PREFIX_ROM"' | sed 's|\"'"$B"'\"|"'"$C"'\"|g'
# => hello world "PREFIX_ROM"
#
# echo 'hello world "PREFIX_ROM\"' | sed 's|\"'"$A"'\\\"|"'"$C"'\\\"|g'
# => hello world "caesar\"
#
# echo 'hello world "PREFIX_ROM\"' | sed 's|\"'"$B"'\\\"|"'"$C"'\\\"|g'
# => hello world "PREFIX_ROM\"

#	sed -i 's|\"'"$key"'\"|"'"$value"'\"|g' flo.txt
#	sed -i 's|\"'"$key"'\\\"|"'"$value"'\\\"|g' flo.txt