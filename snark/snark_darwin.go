// +build darwin,386,!ios

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-apple-darwin -lbls_snark_sys -ldl -lm
*/
import "C"
