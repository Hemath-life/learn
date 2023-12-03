#!/bin/bash

# Display 'Hello'
echo 'Hello'

# Assing value to variable
name='Hemath' # no space after the equal sign
echo "Hi $name"
echo "How are you ${name}"

# String interpulation
echo "${name}'s god"

# Combine the two varialbe
greet='bye'
echo "${name} ${greet}." # echo "$name $greet"


# Reassing the variable
greet='Thank you'
echo "${greet}."
