#!/usr/bin/env bash

# Use this like docker-compose
DOTENV="$(source scripts/dotenv-path.sh)" DEPLOY_DATETIME="$(date -u +%Y-%m-%dT%H:%M:%SZ)" docker-compose -f infra/local/docker-compose.yml --project-directory="$(pwd)" $@
