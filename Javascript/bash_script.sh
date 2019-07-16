#!/bin/bash

# until [ "`/usr/local/bin/docker inspect -f {{.State.Running}} golang_app`"=="true" ]; do
#     sleep 0.1;
# done;

IP=$(curl app:8080/posts)

data={'"data"':$IP}

echo "$data" > patients.json

