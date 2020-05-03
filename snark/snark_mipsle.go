// +build linux,mipsle

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/mipsel-unknown-linux-gnu -lepoch_snark -ldl -lm
*/
import "C"
