#!/usr/bin/env bash

# This script initializes dependencies based on whether it detects this is running in a CI (CircleCI) environment or not

set -ex

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# local development tools (these are already on the circleci docker image)
if [ "${CI}" == "" ]; then
  # Install goreleaser
  HOMEBREW_NO_AUTO_UPDATE=1 brew install goreleaser/tap/goreleaser yq || echo "dependencies already present"

  # linter
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.27.0

  # `go get` will modify go.[mod|sum], so we need to execute outside the gomod context
  cd $GOPATH

  # tests
  go get -u github.com/rakyll/gotest

  # gox compiler
  go get -u github.com/mitchellh/gox

  # swagger cli
  go get -u github.com/go-swagger/go-swagger/cmd/swagger@v0.22.0

  # version bumper
  go get -u github.com/giantswarm/semver-bump

  # hotload watcher
  go get -u github.com/canthefason/go-watcher/cmd/watcher

  # chamber (ps management)
  go get -u github.com/segmentio/chamber

  # go back to repo root
  cd $DIR
fi

go get ./...
go build ./...
go version
