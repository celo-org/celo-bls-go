// +build linux,arm,!arm7

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/arm-unknown-linux-gnueabi -lepoch_snark -ldl -lm
*/
import "C"
