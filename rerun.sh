#!/bin/bash
docker-compose down --rmi all -v
docker-compose build --no-cache
docker-compose up -d