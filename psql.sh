#!/bin/sh

CONTAINER=$1

sudo docker exec -it $CONTAINER psql -U root
