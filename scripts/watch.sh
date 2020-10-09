#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd /go/src/github.com/maxheckel/covid_county/cmd/covid_county
chmod +x ../../infra/local/watcher/watcher-0.2.4-linux-amd64
../../infra/local/watcher/watcher-0.2.4-linux-amd64 -watch github.com/maxheckel/covid_county
