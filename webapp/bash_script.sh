#!/bin/bash

sleep 10

IP=$(curl app:8080/posts)
echo "$IP" > patients.json
