// +build linux,arm64

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-unknown-linux-gnu -lbls_snark_sys -ldl -lm
*/
import "C"

