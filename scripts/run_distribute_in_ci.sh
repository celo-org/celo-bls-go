#!/bin/bash -e

if [[ ! -z $CIRCLE_SHA1 && $(git log --format=oneline -n 1 $CIRCLE_SHA1) != *"BUNDLE"* ]]; then
  echo "Not local and commit message doesn't contain BUNDLE, distributing only linux."
  go run cmd/distribute/distribute.go . ./platforms/platforms_linux.json
  exit 0
fi

go run cmd/distribute/distribute.go . ./platforms/platforms.json
