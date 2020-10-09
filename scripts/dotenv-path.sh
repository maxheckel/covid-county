#!/usr/bin/env bash

if ! test -f "$DOTENV"; then
  DOTENV="./infra/local/.env"
fi

echo "$DOTENV"

