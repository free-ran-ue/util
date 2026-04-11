package util_test

import (
	"testing"

	"github.com/free-ran-ue/util"
	"github.com/go-playground/assert/v2"
)

var testSupiCases = []struct {
	name           string
	mccLength      int
	mncLength      int
	supi           string
	expectedLength int
	expectedBytes  []byte
}{
	{
		name:           "imsi-208930000007487",
		mccLength:      3,
		mncLength:      2,
		supi:           "208930000007487",
		expectedLength: 13,
		expectedBytes:  []byte{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78},
	},
	{
		name:           "imsi-208930000000001",
		mccLength:      3,
		mncLength:      2,
		supi:           "208930000000001",
		expectedLength: 13,
		expectedBytes:  []byte{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10},
	},
	{
		name:           "imsi-001001000000001",
		mccLength:      3,
		mncLength:      3,
		supi:           "001001000000001",
		expectedLength: 13,
		expectedBytes:  []byte{0x01, 0x00, 0x11, 0x00, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10},
	},
}

func TestSupiToBytes(t *testing.T) {
	for _, testCase := range testSupiCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := util.SupiToBytes(testCase.mccLength, testCase.mncLength, testCase.supi)
			assert.Equal(t, testCase.expectedLength, len(result))
			assert.Equal(t, testCase.expectedBytes, result)
		})
	}
}
