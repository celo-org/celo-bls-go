// +build !android,!musl

package bls

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-unknown-linux-gnu -L${SRCDIR}/../libs/x86_64-unknown-linux-gnu -lbls_snark_sys -ldl -lm
*/
import "C"

