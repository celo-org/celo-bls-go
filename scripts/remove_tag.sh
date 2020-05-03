#!/bin/bash -e

export VERSION=v$1
git push origin --delete ${VERSION}
git tag -d ${VERSION}
