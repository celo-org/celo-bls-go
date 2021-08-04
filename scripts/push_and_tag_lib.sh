#!/bin/bash -ex

function push_tag() {
  TAG=$1

  COMMIT=$(git rev-parse HEAD)
  git tag -d $TAG || true
  git tag $TAG $COMMIT
  git push --delete origin $TAG || true
  OUTPUT=$(git push origin $TAG)
  if [[ ! $? = 0 ]]; then
    echo $OUTPUT
    echo $OUTPUT | grep "No new"
  fi
}
