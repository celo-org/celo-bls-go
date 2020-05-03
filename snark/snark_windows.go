// +build windows,386

package snark

/*
#cgo LDFLAGS: -L${SRCDIR}/../libs/i686-pc-windows-gnu -lepoch_snark -lm -lws2_32 -luserenv -lunwind
*/
import "C"

