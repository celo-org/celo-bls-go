#!/bin/bash -e

pushd bls-zexe
export VERSION=`cargo pkgid -p epoch-snark | cut -d'#' -f2 | cut -d: -f2`
popd
git checkout -b v${VERSION}
./scripts/release.sh
git add libs
git commit -m"adds libs"
git push -u origin v${VERSION}
