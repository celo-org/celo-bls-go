// +build !android,linux,amd64,!musl

package bls

import (
    blsRoute "github.com/celo-org/celo-bls-go-linux/bls"
)

const (
	MODULUS377        = blsRoute.MODULUS377
	MODULUSBITS       = blsRoute.MODULUSBITS
	MODULUSMASK       = blsRoute.MODULUSMASK
	PRIVATEKEYBYTES   = blsRoute.PRIVATEKEYBYTES
	PUBLICKEYBYTES    = blsRoute.PUBLICKEYBYTES
	SIGNATUREBYTES    = blsRoute.SIGNATUREBYTES
	EPOCHENTROPYBYTES = blsRoute.EPOCHENTROPYBYTES
)

var (
	GeneralError       = blsRoute.GeneralError
	NotVerifiedError   = blsRoute.NotVerifiedError
	IncorrectSizeError = blsRoute.IncorrectSizeError
	NilPointerError    = blsRoute.NilPointerError
	EmptySliceError    = blsRoute.EmptySliceError
)

type PrivateKey = blsRoute.PrivateKey
type PublicKey = blsRoute.PublicKey
type Signature = blsRoute.Signature
type SignedBlockHeader = blsRoute.SignedBlockHeader
type EpochEntropy  = blsRoute.EpochEntropy

func InitBLSCrypto() {
    blsRoute.InitBLSCrypto()
}

func GeneratePrivateKey() (*PrivateKey, error) {
    return blsRoute.GeneratePrivateKey()
}

func DeserializePrivateKey(privateKeyBytes []byte) (*PrivateKey, error) {
    return blsRoute.DeserializePrivateKey(privateKeyBytes)
}

func HashDirect(message []byte, usePoP bool) ([]byte, error) {
    return blsRoute.HashDirect(message, usePoP)
}

func HashDirectWithAttempt(message []byte, usePoP bool) ([]byte, uint, error) {
    return blsRoute.HashDirectWithAttempt(message, usePoP)
}

func HashComposite(message []byte, extraData []byte) ([]byte, error) {
    return blsRoute.HashComposite(message, extraData)
}

func HashDirectFirstStep(message []byte, hashBytes int32) ([]byte, error) {
    return blsRoute.HashDirectFirstStep(message, hashBytes)
}

func HashCRH(message []byte, hashBytes int32) ([]byte, error) {
    return blsRoute.HashCRH(message, hashBytes)
}

func HashCompositeCIP22(message []byte, extraData []byte) ([]byte, uint8, error) {
    return blsRoute.HashCompositeCIP22(message, extraData)
}

func CompressSignature(signature []byte) ([]byte, error) {
    return blsRoute.CompressSignature(signature)
}

func CompressPublickey(pubkey []byte) ([]byte, error) {
    return blsRoute.CompressPublickey(pubkey)
}

func DeserializePublicKey(publicKeyBytes []byte) (*PublicKey, error) {
    return blsRoute.DeserializePublicKey(publicKeyBytes)
}

func DeserializePublicKeyCached(publicKeyBytes []byte) (*PublicKey, error) {
    return blsRoute.DeserializePublicKeyCached(publicKeyBytes)
}

func BatchVerifyEpochs(signedHeaders []*SignedBlockHeader, shouldUseCompositeHasher, shouldUseCIP22 bool) error {
    return blsRoute.BatchVerifyEpochs(signedHeaders, shouldUseCompositeHasher, shouldUseCIP22)
}

func DeserializeSignature(signatureBytes []byte) (*Signature, error) {
    return DeserializeSignature(signatureBytes)
}

func AggregatePublicKeys(publicKeys []*PublicKey) (*PublicKey, error) {
    return blsRoute.AggregatePublicKeys(publicKeys)
}

func AggregatePublicKeysSubtract(aggregatedPublicKey *PublicKey, publicKeys []*PublicKey) (*PublicKey, error) {
    return blsRoute.AggregatePublicKeysSubtract(aggregatedPublicKey, publicKeys)
}

func AggregateSignatures(signatures []*Signature) (*Signature, error) {
    return blsRoute.AggregateSignatures(signatures)
}

func EncodeEpochToBytesCIP22(epochIndex uint16, round uint8, blockHash, parentHash EpochEntropy, maximumNonSigners, maximumValidators uint32, addedPublicKeys []*PublicKey) ([]byte, []byte, error) {
    return blsRoute.EncodeEpochToBytesCIP22(epochIndex, round, blockHash, parentHash, maximumNonSigners, maximumValidators, addedPublicKeys)
}

func EncodeEpochToBytes(epochIndex uint16, maximumNonSigners uint32, addedPublicKeys []*PublicKey) ([]byte, error) {
    return blsRoute.EncodeEpochToBytes(epochIndex, maximumNonSigners, addedPublicKeys)
}
