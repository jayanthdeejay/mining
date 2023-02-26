package main

import (
	"fmt"

	"github.com/jayanthdeejay/mining/producer/binhex"
	"github.com/jayanthdeejay/mining/producer/necklace"
)

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258

// =====================================================================
// Generate de Bruijn sequences by iteratively applying a successor rule
// =====================================================================
func DB(n uint16, initial string) {
	var a [MAX]uint8
	var count = 0
	var i uint16 = 1
	if len(initial) == 0 {
		for i <= n {
			a[i] = 0 // First n bits
			i++
		}
	} else {
		a = binhex.HexToBinary(initial)
		fmt.Println(initial, a)
	}
	for {
		// fmt.Printf("%d", a[1])
		if count == 1000000 {
			fmt.Println(binhex.BinaryToHex(a[1:MAX-1]), a)
			count = 0
		}
		count++
		// fmt.Println(a[1 : MAX-1])
		new_bit := necklace.Granddaddy(a, n)
		i = 1
		for i <= n {
			a[i] = a[i+1]
			i++
		}
		a[n] = new_bit
		if necklace.Zeros(a, n) {
			break
		}
	}
}

// ===========================================
func main() {
	var n uint16
	initial := "8000000000000000000000000000000000000000000000000000000000003d09"
	// initial := ""
	// fmt.Printf("Enter n: ")
	// fmt.Scanf("%d", &n)
	n = 256
	DB(n, initial)
	fmt.Printf("\n\n")
}
