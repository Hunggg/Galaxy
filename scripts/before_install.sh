#!/bin/bash

DIR="/home/ec2-user/metrogalaxy-api/"
if [ -d "$DIR" ]; then
  echo "${DIR} exists"
else
  echo "Creating ${DIR} directory"
  mkdir -p ${DIR}
fi

rm -f /home/ec2-user/metrogalaxy-api/*
