#!/bin/sh

LOG=

if $LOG == "error"; then
	tail +1f /var/log/nginx/bot/error_log
else
	tail +1f /var/log/nginx/bot/access_log
fi


