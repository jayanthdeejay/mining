package binhex

import "encoding/hex"

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258

func BinaryToHex(a []uint8) string {
	// Convert the binary digits into a byte array
	b := make([]byte, (len(a)+7)/8)
	for i, v := range a {
		if v == 1 {
			b[i/8] |= 1 << (7 - uint(i)%8)
		}
	}

	// Convert the byte array into a hexadecimal string
	return hex.EncodeToString(b)
}

func HexToBinary(s string) [MAX]uint8 {
	// Convert the hexadecimal string into a byte array
	b, err := hex.DecodeString(s)
	if err != nil {
		// Handle decoding error
		panic(err)
	}

	// Convert the byte array into a slice of binary digits
	var a [MAX - 2]uint8
	for i, v := range b {
		for j := 0; j < 8; j++ {
			k := 8*i + j
			if v&(1<<(7-j)) != 0 {
				a[k] = 1
			}
		}
	}
	var initial [MAX]uint8
	initial[0] = 0
	for i := 1; i <= len(a); i++ {
		initial[i] = a[i-1]
	}
	initial[len(initial)-1] = 0
	return initial
}
