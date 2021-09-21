// +build linux,arm64 !android,linux,amd64,!musl linux,arm,!arm7 arm7 !android,linux,386,!musl !android,musl linux,mips linux,mips64 linux,mips64le linux,mipsle

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
