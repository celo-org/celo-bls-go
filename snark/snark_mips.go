// +build linux,mips

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mips-unknown-linux-gnu -lbls_snark_sys -ldl -lm
*/
import "C"
