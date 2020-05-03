// +build linux,amd64

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-unknown-linux-gnu -lepoch_snark -ldl -lm
*/
import "C"
