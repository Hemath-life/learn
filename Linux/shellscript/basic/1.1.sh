#!/bin/bash


# Display the UID  and username of the user excuting this script
# Display if the user is the root user or not


# dislay the UID
echo "UID is ${UID}"

# diplsy the username
username=`id -un` # or whoami or $(id -un)
echo "Hi ${username}"

# display if the user is the root user
if [[ "${UID}" -eq 0 ]]; # or [ ]
then
    echo 'You are root'
else
    echo 'You are not root';
fi;



