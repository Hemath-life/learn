#!/bin/bash

# this script creates and account on the local system.
# you will be promped for the account name and pasword

# Ask fo the user name.
read -p 'Type user name: ' username # -p for prompt
# Ask for the real name.
read -p 'Type real name: ' realname # -p for prompt
# Ask for the password.
read -p 'Type password: ' password

echo "
username: ${username} 
realname: ${realname}
password: ${password}
";
# Create the user.
sudo useradd 

# Set the password for the user.

# force password change on first login.
