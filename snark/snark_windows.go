// +build windows,386 windows,amd64

package snark

import (
    snarkRoute "github.com/celo-org/celo-bls-go-windows/snark"
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
