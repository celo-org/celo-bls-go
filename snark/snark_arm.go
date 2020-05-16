// +build linux,arm,!arm7

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/arm-unknown-linux-gnueabi -lbls_snark_sys -ldl -lm
*/
import "C"
