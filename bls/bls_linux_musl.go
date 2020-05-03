// +build !android,musl

package bls

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-unknown-linux-musl -lepoch_snark -ldl -lm
*/
import "C"

