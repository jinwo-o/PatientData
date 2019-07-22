#!/bin/bash

sleep 15

IP=$(curl app:8080/get)
echo "$IP" > patients.json
