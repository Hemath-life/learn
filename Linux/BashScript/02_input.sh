#! /usr/bin/bash

# #getting sing e value
# echo "Enter you Name: "
# read name 
# echo "Entered name $name"

# #getting more then one  value
# echo "Enter you Name: "
# read $name1 name2 name3
# echo " Name:  $name1 $name2 $name3"

# # getting input in same line
# read  -p "username :" user_val
# read  -sp "password :" user_pass
# echo "username : $user_val"
# echo "username : $user_pass"

#  getting using defualt variable
echo "username :"
read
echo "username : $REPLY"