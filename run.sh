#!/bin/bash

sudo docker build -t server .

sudo docker run -e TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN} server
