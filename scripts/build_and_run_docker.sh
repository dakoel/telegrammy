#!/bin/bash
dir=$(dirname $(realpath $0))

docker build $dir/../src -t telegrammy
docker run \
  -e TELEGRAMMY_CONFIG_PATH=$TELEGRAMMY_CONFIG_PATH \
  -e TELEGRAM_CHAT_ID=$TELEGRAM_CHAT_ID \
  -e TELEGRAM_BOT_TOKEN=$TELEGRAM_BOT_TOKEN \
  -e CHAT_GPT_TOKEN=$CHAT_GPT_TOKEN \
  -e CHAT_GPT_CONVERSATION_PATH=$CHAT_GPT_CONVERSATION_PATH \
  -v $dir/../example_config:/etc/telegrammy:z \
  telegrammy
