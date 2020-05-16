// +build linux,amd64

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/x86_64-unknown-linux-gnu -lbls_snark_sys -ldl -lm
*/
import "C"
