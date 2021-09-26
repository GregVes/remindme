#!/bin/sh

TELEGRAM_TOKEN=${TELEGRAM_BOT_TOKEN}
CLOUD_FUNCTION_URL="https://radio4000-dev-api.space/bot"

curl --data "url=$CLOUD_FUNCTION_URL" https://api.telegram.org/bot$TELEGRAM_TOKEN/SetWebhook
