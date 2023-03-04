#! /bin/bash

# variable
name1=zyz
name2='zyz'
name3="zyz"

echo $name1
echo ${name2}
echo ${name3}bash

# readonly variable
name=zyz
readonly name
declare -r name
echo $name

# delete variable
var=variable
echo $var
unset var
echo $var

# string
echo 'hello, $name \"hh\"'
echo "hello, $name \"hh\""
echo hello, $name \"hh\"

# len
echo ${#name}

# substr
str="Hello World"
echo ${str:0:5}

