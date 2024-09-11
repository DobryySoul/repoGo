package main

import (
	"errors"
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	tests := []struct {
		i []byte
		o int
		e error
	}{
		{
			i: []byte("Hello"),
			o: 5,
			e: nil,
		},
		{
			i: []byte("Привет"),
			o: 6,
			e: nil,
		},
		{
			i: []byte("こんにちは"),
			o: 5,
			e: nil,
		},
		{
			i: []byte("Invalid\x80UTF8"),
			o: 0,
			e: ErrInvalidUTF8,
		},
	}
	for _, test := range tests {
		res, err := GetUTFLength(test.i)
		if res != test.o || !errors.Is(err, test.e) {
			t.Errorf("GetUTFLength(%q) = %d, %v; expected %d, %v", test.i, res, err, test.o, test.e)
		}
	}
}
