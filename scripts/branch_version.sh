#!/bin/bash -e

pushd bls-zexe
export VERSION_FROM_CARGO=v`cargo pkgid -p epoch-snark | cut -d'#' -f2 | cut -d: -f2`
export VERSION=${1:-${VERSION_FROM_CARGO}}
git checkout ${VERSION}
popd
./scripts/release.sh
git checkout -b ${VERSION}
git add libs
git commit -m"adds libs"
git push -u origin ${VERSION}
