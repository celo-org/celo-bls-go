// +build linux,mips

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mips-unknown-linux-gnu -lepoch_snark -ldl -lm
*/
import "C"
