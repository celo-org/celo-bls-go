package utils

import "reflect"

const FIELD_SIZE = 48
const FIELD_SIZE_IN_CONTRACT = 32

func ReverseAnyAndPad(s []byte) []byte {
	s = ReverseAny(s)
	padding := make([]byte, FIELD_SIZE_IN_CONTRACT- (len(s) %FIELD_SIZE_IN_CONTRACT))
	z := append(padding, s...)
	return z
}

func ReverseAny(s []byte) []byte {
	z := make([]byte, len(s))
	copy(z, s)
	n := reflect.ValueOf(z).Len()
	swap := reflect.Swapper(z)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return z
}
