#!/bin/bash -e

pushd celo-bls-snark-rs
export VERSION=`cargo pkgid -p epoch-snark | cut -d'#' -f2 | cut -d: -f2`
popd
echo ${VERSION}
