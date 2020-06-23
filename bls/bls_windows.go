// +build windows,386

package bls

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-pc-windows-msvc -lbls_snark_sys -lm -lws2_32 -luserenv -lunwind
*/
import "C"

