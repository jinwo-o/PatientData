#!/bin/bash

sleep 2

IP=$(curl app:8080/posts)
echo "$IP" > patients.json
