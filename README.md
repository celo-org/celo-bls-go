celo-bls-go
-----------

Go module for [celo-bls-snark-rs](https://github.com/celo-org/celo-bls-snark-rs/).

## Release process

* Create a new branch
* Create a PR
* The CI will now build all the libs and store them as tar.gz. artifact. Use it to update all the libs and commit them.
* After the PR is merged, tag the version using `./scripts/tag_version.sh`

If needed, you can remove an old tag using `./scripts/remove_tag.sh VERSION`.
