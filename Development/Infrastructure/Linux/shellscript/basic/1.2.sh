#!/bin/bash

# display if the UID does NOT match 1000
UID_TO_TEST_FOR='1000';

if [[ "${UID}" -ne "${UID_TO_TEST_FOR}" ]];
then
    echo "Your UID does not match with: ${UID_TO_TEST_FOR}"
    exit 1;
fi

# display the username
username=`id -un` # or whoami or $(id -un)


if [[ "${?}" -ne 0 ]];
then
    echo "not excuted successfully ${?}"
    exit 1;
fi

echo "Your UID is :${UID}"
echo "Your username is :${username}";

exit 0