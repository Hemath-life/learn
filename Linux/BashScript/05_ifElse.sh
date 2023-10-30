#! /usr/bin/bash


# integer
# > greter then
# < less then
# <=
# >=
# ne not equal

# # sting
# = equal
# == equal
# != not equal
# < is less then ASCII aphabetic order
# > is greter then ASCII aphabetic order
# -z sting null or that as zero lenth


count=10
if  (( $count  > 9  ))
then
          echo "condtion is true"
fi

count=10
if  (($count  < 5  ))
then
          echo "condtion is true"
else
          echo "condtion is false"
fi