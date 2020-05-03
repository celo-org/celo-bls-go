celo-bls-go
-----------

Go module for [bls-zexe](https://github.com/celo-org/bls-zexe/).

## Release process

* Create a new branch
* Update all the libs using `./scripts/release.sh`
* Create a PR
* After the PR is merged, tag the version using `./scripts/tag_version.sh`

If needed, you can remove an old tag using `./scripts/remove_tag.sh VERSION`.
