#!/bin/bash -e

pushd celo-bls-snark-rs
export VERSION_FROM_CARGO=v`cargo pkgid -p epoch-snark | cut -d'#' -f2 | cut -d: -f2`
export VERSION=${1:-${VERSION_FROM_CARGO}}
popd
git tag ${VERSION}
git push --tags
