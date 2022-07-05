#!/bin/bash

sudo yum update -y
sudo yum install -y docker

cd /home/ec2-user/metrogalaxy-api
sudo docker build -t quangkeudocker/metrogalaxy-api .

sudo systemctl daemon-reload
sudo systemctl start docker