// +build !android,musl

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-unknown-linux-musl -lbls_snark_sys -ldl -lm
*/
import "C"

