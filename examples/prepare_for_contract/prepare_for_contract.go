package main

import (
	"encoding/hex"
	"fmt"
	"github.com/celo-org/celo-bls-go/bls"
	"github.com/celo-org/celo-bls-go/examples/utils"
	"os"
)

/*
examples:
./prepare_for_contract pk aab506de1ef9b0df75f202b0813904e08d99ba0dbbf2084c3a983d9190c41f5f773489a6ee530da67d517d3151805101860dfdb8d7d72d768643af1b07a468f93ba1e08edb2a7f22c85bad3c2c02545f036647f11ce63eed3bd44e2cc080c480
./prepare_for_contract sig 063c39b0d4f7fa61cef3b97fe6705f02ffc209a4f3c91442588f15a884e25e68fb63cc60215ac6912f67b9fee9276501
*/

func main() {
	bls.InitBLSCrypto()
	kind := os.Args[1]
	arg := os.Args[2]
	switch kind {
	case "pk":
		dec, _ := hex.DecodeString(arg)
		key, _ := bls.DeserializePublicKey(dec)
		enc, _ := key.SerializeUncompressed()
		enc1 := enc[0:utils.FIELD_SIZE]
		enc1 = utils.ReverseAnyAndPad(enc1)
		enc2 := enc[utils.FIELD_SIZE :2*utils.FIELD_SIZE]
		enc2 = utils.ReverseAnyAndPad(enc2)
		enc3 := enc[2*utils.FIELD_SIZE :3*utils.FIELD_SIZE]
		enc3 = utils.ReverseAnyAndPad(enc3)
		enc4 := enc[3*utils.FIELD_SIZE :4*utils.FIELD_SIZE]
		enc4 = utils.ReverseAnyAndPad(enc4)
		fmt.Printf("Encoded public key: %x%x%x%x\n", enc1, enc2, enc3, enc4)

	case "sig":
		dec, _ := hex.DecodeString(arg)
		key, _ := bls.DeserializeSignature(dec)
		enc, _ := key.SerializeUncompressed()
		enc1 := enc[0:utils.FIELD_SIZE]
		enc1 = utils.ReverseAnyAndPad(enc1)
		enc2 := enc[utils.FIELD_SIZE :2*utils.FIELD_SIZE]
		enc2 = utils.ReverseAnyAndPad(enc2)
		fmt.Printf("Encoded signature: %x%x\n", enc1, enc2)
	}
}
