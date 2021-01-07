#!/bin/bash
sudo docker kill $(docker ps -aq)
sudo docker rmi -f $(docker images -q)
sudo docker system prune -f
sudo docker volume rm $(docker volume ls -qf dangling=true)
sudo docker network rm $(docker network ls -q)
sudo docker-compose up -d
