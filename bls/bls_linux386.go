// +build linux,386,!musl

package bls

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-unknown-linux-gnu -lepoch_snark -ldl -lm
*/
import "C"
