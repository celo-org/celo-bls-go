package snark

/*
#include "snark.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

var VerificationError = errors.New("SNARK proof verification failed")

/// Serialized Groth16 Proof
type Proof []byte

/// Serialized Verifying Key
type VerifyingKey []byte

/// The EpochBlock to be used for Verification in the SNARK
type EpochBlock struct {
	/// Index of the epoch
	Index uint16
	/// Max non signers per block
	MaxNonSigners uint32
	/// Max validators per block
	MaxValidators uint
	/// Epoch entropy value
	EpochEntropy []byte
	/// Parent entropy value
	ParentEntropy []byte
	/// Serialized public keys of the validators in this epoch (each `PUBLIC_KEY_BYTES` long)
	PublicKeys [][]byte
}

const PUBLIC_KEY_BYTES = 96

func sliceToPtr(slice []byte) (*C.uchar, C.int) {
	if len(slice) == 0 {
		return nil, 0
	} else {
		return (*C.uchar)(unsafe.Pointer(&slice[0])), C.int(len(slice))
	}
}

func vecToPtr(slice [][]byte) (*C.uchar, C.int) {
	if len(slice) == 0 {
		return nil, 0
	} else {
		return (*C.uchar)(unsafe.Pointer(&slice[0][0])), C.int(len(slice) * PUBLIC_KEY_BYTES)
	}
}

func VerifyEpochs(
	verifyingKey VerifyingKey,
	proof Proof,
	firstEpoch EpochBlock,
	lastEpoch EpochBlock,
) error {
	vkPtr, vkLen := sliceToPtr(verifyingKey)
	proofPtr, proofLen := sliceToPtr(proof)

	firstPublicKeysPtr, _ := vecToPtr(firstEpoch.PublicKeys)
	firstEpochEntropyPtr, _ := sliceToPtr(firstEpoch.EpochEntropy)
	firstParentEntropyPtr, _ := sliceToPtr(firstEpoch.ParentEntropy)
	firstEpochRaw := C.EpochBlockFFI{
		index:               C.ushort(firstEpoch.Index),
		maximum_non_signers: C.uint(firstEpoch.MaxNonSigners),
		maximum_validators:  C.ulong(firstEpoch.MaxValidators),
		epoch_entropy:		 firstEpochEntropyPtr,
		parent_entropy:		 firstParentEntropyPtr,
		pubkeys_num:         C.ulong(len(firstEpoch.PublicKeys)),
		pubkeys:             firstPublicKeysPtr,
	}

	lastPublicKeysPtr, _ := vecToPtr(lastEpoch.PublicKeys)
	lastEpochEntropyPtr, _ := sliceToPtr(lastEpoch.EpochEntropy)
	lastParentEntropyPtr, _ := sliceToPtr(lastEpoch.ParentEntropy)
	lastEpochRaw := C.EpochBlockFFI{
		index:               C.ushort(lastEpoch.Index),
		maximum_non_signers: C.uint(lastEpoch.MaxNonSigners),
		maximum_validators:  C.ulong(lastEpoch.MaxValidators),
		epoch_entropy:		 lastEpochEntropyPtr,
		parent_entropy:		 lastParentEntropyPtr,
		pubkeys_num:         C.ulong(len(lastEpoch.PublicKeys)),
		pubkeys:             lastPublicKeysPtr,
	}

	success := C.verify(
		vkPtr,
		C.uint(vkLen),
		proofPtr,
		C.uint(proofLen),
		firstEpochRaw,
		lastEpochRaw,
	)

	if !success {
		return VerificationError
	}

	return nil
}
