celo-bls-go
-----------

Go module for [celo-bls-snark-rs](https://github.com/celo-org/celo-bls-snark-rs/).

## Release process

[CircleCI] is configured to build release libraries for all supported architectures and operating
systems. Building and bundling of new releases can be achieved through the following process.

* Create a new branch.
* Delete the old `./libs` folder, if it exists.
* Create a PR.
* On each commit, CI will run:
    * By default, it will build and test the code on amd64 Linux.
    * Upon [approval] by a contributor with push access, CI will build the libs for all targets and store them as a `libs.tar.gz` artifact on the `build-and-bundle-libs-all-targets` job.

* Download the `libs.tar.gz` file and extract it in the root of the repository. This will create the `./libs` directory.
* Run `go run cmd/distribute/distribute.go . platforms/platforms.json`. This will create all the repositories for the different packages.
* Run `./scripts/push_and_tag.sh TAG` with a chosen `TAG`. This will create a tag in each of the repos and update the go.mod.
    * Tag should be formated as a semver such as `v0.1.2`.
* Merge the PR, so master will be up to date.
* Run `./scripts/push_and_tag_master.sh TAG`, to push the tag to the main repository.

[CircleCI]: https://app.circleci.com/pipelines/github/celo-org/celo-bls-go
[approval]: https://circleci.com/docs/2.0/workflows/#holding-a-workflow-for-a-manual-approval
