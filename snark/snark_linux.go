// +build linux,arm64 !android,linux,amd64,!musl arm7 !android,linux,386,!musl !android,musl

package snark

import (
    snarkRoute "github.com/celo-org/celo-bls-go-linux/snark"
)

var VerificationError = snarkRoute.VerificationError

type Proof = snarkRoute.Proof

type VerifyingKey = snarkRoute.VerifyingKey

type EpochBlock = snarkRoute.EpochBlock

const PUBLIC_KEY_BYTES = snarkRoute.PUBLIC_KEY_BYTES

func VerifyEpochs(
	verifyingKey VerifyingKey,
	proof Proof,
	firstEpoch EpochBlock,
	lastEpoch EpochBlock,
) error {
    return snarkRoute.VerifyEpochs(verifyingKey , proof, firstEpoch, lastEpoch)
}
