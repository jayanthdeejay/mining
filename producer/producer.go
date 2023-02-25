package main

import (
	"encoding/hex"
	"fmt"
	"time"
)

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258

// =====================
// Test if a[1..n] = 0^n
// =====================
func Zeros(a [MAX]uint16, n uint16) uint16 {
	var i uint16 = 1
	for i <= n {
		if a[i] == 1 {
			return 0
		}
		i++
	}
	return 1
}

// =============================
// Test if b[1..n] is a necklace
// =============================
func IsNecklace(b [MAX]uint16, n uint16) uint16 {
	var p uint16 = 1
	var i uint16 = 2
	for i <= n {
		if b[i-p] > b[i] {
			return 0
		}
		if b[i-p] < b[i] {
			p = i
		}
		i++
	}
	if n%p != 0 {
		return 0
	}
	return 1
}

// ===========================================
// Necklace Successor Rules
// ===========================================
func Granddaddy(a [MAX]uint16, n uint16) uint16 {
	var b [MAX]uint16
	var j uint16 = 2
	for j <= n && a[j] == 1 {
		j++
	}
	var i uint16 = j
	for i <= n {
		b[i-j+1] = a[i]
		i++
	}
	b[n-j+2] = 0
	i = 2
	for i < j {
		b[n-j+i+1] = a[i]
		i++
	}
	if IsNecklace(b, n) != 0 {
		return 1 - a[1]
	}
	return a[1]
}

// =====================================================================
// Generate de Bruijn sequences by iteratively applying a successor rule
// =====================================================================
func DB(n uint16) {
	var a [MAX]uint16
	var i uint16 = 1
	for i <= n {
		a[i] = 0 // First n bits
		i++
	}
	for {
		// fmt.Printf("%d", a[1])
		time.Sleep(time.Millisecond * 300)
		fmt.Println(binaryToHex(a[1 : MAX-1]))
		// fmt.Println(a[1 : MAX-1])
		new_bit := Granddaddy(a, n)
		i = 1
		for i <= n {
			a[i] = a[i+1]
			i++
		}
		a[n] = new_bit
		if Zeros(a, n) != 0 {
			break
		}
	}
}

func binaryToHex(a []uint16) string {
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

// ===========================================
func main() {
	var n uint16
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)
	DB(n)
	fmt.Printf("\n\n")
}
