// +build !android,musl

package snark

/*
#cgo LDFLAGS: ${SRCDIR}/../libs/x86_64-unknown-linux-musl/libbls_snark_sys.so -ldl -lm
*/
import "C"

