package main

import (
	"encoding/hex"
	"fmt"
	"github.com/celo-org/celo-bls-go/examples/utils"
	"github.com/celo-org/celo-bls-go/platforms"
	"math/big"
	"os"
)

/*
examples:
./prepare_for_contract pk aab506de1ef9b0df75f202b0813904e08d99ba0dbbf2084c3a983d9190c41f5f773489a6ee530da67d517d3151805101860dfdb8d7d72d768643af1b07a468f93ba1e08edb2a7f22c85bad3c2c02545f036647f11ce63eed3bd44e2cc080c480
./prepare_for_contract sig 063c39b0d4f7fa61cef3b97fe6705f02ffc209a4f3c91442588f15a884e25e68fb63cc60215ac6912f67b9fee9276501
*/

func main() {
	platforms.InitBLSCrypto()
	kind := os.Args[1]
	arg := os.Args[2]
	switch kind {
	case "pk":
		dec, _ := hex.DecodeString(arg)
		key, _ := platforms.DeserializePublicKey(dec)
		enc, _ := key.SerializeUncompressed()
		enc1 := enc[0:utils.FIELD_SIZE]
		enc1 = utils.ReverseAnyAndPad(enc1)
		enc2 := enc[utils.FIELD_SIZE:2*utils.FIELD_SIZE]
		enc2 = utils.ReverseAnyAndPad(enc2)
		enc3 := enc[2*utils.FIELD_SIZE:3*utils.FIELD_SIZE]
		enc3 = utils.ReverseAnyAndPad(enc3)
		enc4 := enc[3*utils.FIELD_SIZE:4*utils.FIELD_SIZE]
		enc4 = utils.ReverseAnyAndPad(enc4)
		fmt.Printf("Encoded public key: %x%x%x%x\n", enc1, enc2, enc3, enc4)

	case "sig":
		dec, _ := hex.DecodeString(arg)
		key, _ := platforms.DeserializeSignature(dec)
		enc, _ := key.SerializeUncompressed()
		enc1 := enc[0:utils.FIELD_SIZE]
		enc1 = utils.ReverseAnyAndPad(enc1)
		enc2 := enc[utils.FIELD_SIZE:2*utils.FIELD_SIZE]
		enc2 = utils.ReverseAnyAndPad(enc2)
		fmt.Printf("Encoded signature: %x%x\n", enc1, enc2)

	case "prefix":
		msg, _ := hex.DecodeString(arg)
		hash, prefix, _ := platforms.HashDirectWithAttempt(msg, false)
		fmt.Printf("Prefix: %02x, Hash: %x\n", prefix, hash)

	case "hints":
		msg, _ := hex.DecodeString(arg)
		_, prefix, _ := platforms.HashDirectWithAttempt(msg, false)
		hash, _ := platforms.HashDirectFirstStep(append([]byte{byte(prefix)}, msg...), 64)
		hash = hash[0:48]
		hash = utils.ReverseAny(hash)
		hash[0] &= 1
		x := big.NewInt(0).SetBytes(hash)
		n, _ := big.NewInt(0).SetString("258664426012969094010652733694893533536393512754914660539884262666720468348340822774968888139573360124440321458177", 10)
		x = x.Exp(x, big.NewInt(3), n)
		x = x.Add(x, big.NewInt(1))
		y := big.NewInt(0).ModSqrt(x, n)
		yNeg := big.NewInt(0).Sub(n, y)
		yBytes := utils.ReverseAnyAndPad(utils.ReverseAny(y.Bytes()))
		yNegBytes := utils.ReverseAnyAndPad(utils.ReverseAny(yNeg.Bytes()))
		fmt.Printf("Encoded hints: %x%x\n", yBytes, yNegBytes)
	}
}
