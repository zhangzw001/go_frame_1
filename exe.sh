#!/bin/sh
# you should exec compile.sh first

if [ $# != 4 ]
then
	echo "Usage: $0 aunt_file order_file result_path timefile"
	exit 1
fi

BIN_PATH=./
cd ${BIN_PATH}

echo "status:running" >> coder_info

./a.out $1 $2 $3 $4 >./exe_result 2>&1
