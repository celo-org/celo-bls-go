// +build linux,arm64

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/aarch64-unknown-linux-gnu -lepoch_snark -ldl -lm
*/
import "C"

