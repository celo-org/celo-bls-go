// +build darwin,amd64,!ios

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-apple-darwin -lepoch_snark -ldl -lm
*/
import "C"

