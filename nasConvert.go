package util

import (
	"strconv"
)

func encodePlmn(mcc, mnc string) []byte {
	// 3GPP PLMN encoding:
	// Octet 1 = MCC digit2 | MCC digit1
	// Octet 2 = MNC digit3 (or 0xF for 2-digit MNC) | MCC digit3
	// Octet 3 = MNC digit2 | MNC digit1
	mccD1, _ := strconv.Atoi(string(mcc[0]))
	mccD2, _ := strconv.Atoi(string(mcc[1]))
	mccD3, _ := strconv.Atoi(string(mcc[2]))
	mncD1, _ := strconv.Atoi(string(mnc[0]))
	mncD2, _ := strconv.Atoi(string(mnc[1]))

	mncD3 := 0xf
	if len(mnc) == 3 {
		mncD3, _ = strconv.Atoi(string(mnc[2]))
	}

	octet1 := uint8((mccD2 << 4) | mccD1)
	octet2 := uint8((mncD3 << 4) | mccD3)
	octet3 := uint8((mncD2 << 4) | mncD1)

	return []byte{octet1, octet2, octet3}
}

func encodeMsin(msin string) []byte {
	result := make([]byte, (len(msin)+1)/2)

	for i := 0; i < len(msin); i += 2 {
		var d1, d2 uint8

		tmpD1, err := strconv.Atoi(string(msin[i]))
		if err != nil {
			panic(err)
		}
		d1 = uint8(tmpD1)

		if i+1 < len(msin) {
			tmpD2, err := strconv.Atoi(string(msin[i+1]))
			if err != nil {
				panic(err)
			}
			d2 = uint8(tmpD2)
		} else {
			d2 = 0x0f
		}

		result[i/2] = (d2 << 4) | d1
	}

	return result
}

func SupiToBytes(mccLength, mncLength int, supi string) []byte {
	var mcc, mnc, msin string

	mcc = supi[0:mccLength]
	mnc = supi[mccLength : mccLength+mncLength]
	msin = supi[mccLength+mncLength:]

	plmnBytes := encodePlmn(mcc, mnc)
	msinBytes := encodeMsin(msin)

	buffer := make([]byte, 0)

	// Byte 0: SUPI Type (0x01 = IMSI)
	buffer = append(buffer, 0x01)

	// Bytes 1-3: MCC + MNC code
	buffer = append(buffer, plmnBytes...)

	// Bytes 4-5: Routing Indicator (0000)
	buffer = append(buffer, 0x00, 0x00)

	// Byte 6: Protection Scheme ID (0x00 = null scheme)
	buffer = append(buffer, 0x00)

	// Byte 7: Home Network Public Key Identifier (0x00 for null scheme)
	buffer = append(buffer, 0x00)

	// Bytes 8+: Scheme output (MSIN for null scheme)
	buffer = append(buffer, msinBytes...)

	return buffer
}
