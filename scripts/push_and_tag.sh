#!/bin/bash -ex

TAG=$1

declare -a platforms=("linux" "macos" "android" "ios" "other" )

echo -e "module github.com/celo-org/celo-bls-go\n\ngo 1.12\n\nrequire (" > go.mod

function push_tag() {
  TAG=$1
  COMMIT=$2

  git push -f origin master
  COMMIT=$(git rev-parse origin/master)
  git tag -d $TAG || true
  git tag $TAG $COMMIT
  git push --delete origin $TAG || true
  OUTPUT=$(git push origin $TAG)
  if [[ ! $? = 0 ]]; then
    echo $OUTPUT
    echo $OUTPUT | grep "No new"
  fi
}

for platform in ${platforms[@]}; do
  pushd platforms/repos/celo-bls-go-$platform
  rm -rf .git
  git init
  git add .
  git commit -m"sync master"
  git remote add origin https://github.com/celo-org/celo-bls-go-$platform
  git push -f origin master
  push_tag $TAG $COMMIT
  rm -rf .git
  popd

  echo -e "\tgithub.com/celo-org/celo-bls-go-$platform\t$TAG" >> go.mod
done

echo -e ")" >> go.mod

push_tag $TAG $COMMIT
