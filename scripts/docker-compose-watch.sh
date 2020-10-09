#!/usr/bin/env bash
# Use this like docker-compose
DOTENV="$(source scripts/dotenv-path.sh)" docker-compose -f infra/local/docker-compose.yml -f infra/local/docker-compose.watch.yml --project-directory="$(pwd)" $@
