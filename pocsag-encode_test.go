package pocsagencode

import (
	"testing"
)

func TestEncodeShort(t *testing.T) {
	capcode := 1
	message := "ABC"
	data := []byte{
		0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA,
		0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xD8, 0x15, 0xD2, 0x7C, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0xA5, 0x1D, 0x00, 0x00,
		0xB8, 0x87, 0x43, 0xC1, 0xD1, 0x05, 0x00, 0xD4, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A,
		0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A,
	}

	tcx := EncodeTransmission(capcode, message)

	if len(data) != len(tcx) {
		t.Errorf("Expected %d uint32's got %d\nExpected: %X\nGot     : %X\n", len(data), len(tcx), data, tcx)
		return
	}

	for i, v := range data {
		if tcx[i] != v {
			t.Errorf("byte at index %d: expected %X got %X\n", i, v, tcx[i])
		}
	}

	t.Logf("DATA: % X", data)
	t.Logf("TCX:  % X", tcx)
}

func TestEncode(t *testing.T) {
	capcode := 12345678
	message := "The quick brown fox jumped the lay cow."
	data := []byte{
		0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA,
		0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xD8, 0x15, 0xD2, 0x7C, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A,
		0x97, 0xC1, 0x89, 0x7A, 0x10, 0x38, 0x85, 0xF1, 0xB5, 0x4F, 0x17, 0x95, 0xBC, 0xAA, 0x47, 0xC1, 0x08, 0xEC, 0xF1, 0xF2, 0x24, 0xD0, 0x48, 0xB0, 0x1A, 0xEE, 0xBE, 0xBF, 0x94, 0x3B, 0x13, 0xEC, 0xD5, 0x13, 0x3C, 0xF6, 0x67, 0x6C, 0xAF, 0xAB, 0x23, 0x24, 0xD3, 0xC3,
		0xD1, 0x89, 0x8B, 0xE0, 0x45, 0x8B, 0x60, 0xBA, 0xD8, 0x15, 0xD2, 0x7C, 0xCC, 0xE5, 0x39, 0xDC, 0x1D, 0xBA, 0x1F, 0x8B, 0xD4, 0x43, 0xE9, 0xEE, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A,
		0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A, 0x97, 0xC1, 0x89, 0x7A,
	}

	tcx := EncodeTransmission(capcode, message)

	if len(data) != len(tcx) {
		t.Errorf("Expected %d uint32's got %d\nExpected: %X\nGot     : %X\n", len(data), len(tcx), data, tcx)
		return
	}

	for i, v := range data {
		if tcx[i] != v {
			t.Errorf("byte at index %d: expected %X got %X\n", i, v, tcx[i])
		}
	}

	t.Logf("%X", data)
	t.Logf("%X", tcx)
}