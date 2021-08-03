celo-bls-go
-----------

Go module for [celo-bls-snark-rs](https://github.com/celo-org/celo-bls-snark-rs/).

## Release process

* Create a new branch
* Delete the old libs
* Create a PR
* The CI will now build all the libs and store them as tar.gz. artifact.
* Download the tar file and extract it in the root of the repository. This will create the libs directory.
* Run `go run cmd/distribute/distribute.go . platforms/platforms.json`. This will create all the repositories for the different packages.
* Merge the PR, so master will be up to date.
* Run `./scripts/push_and_tag.sh TAG` with a chosen `TAG`. This will create a tag in each of the repos, update the go.mod and tag the main repository.
