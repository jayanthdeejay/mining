package necklace

// For string of length n, you need MAX = n + 2
// For 256 bit strings, set MAX = 258
const MAX = 258

// =====================
// Test if a[1..n] = 0^n
// =====================
func Zeros(a [MAX]uint8, n uint16) bool {
	var i uint16 = 1
	for i <= n {
		if a[i] == 1 {
			return false
		}
		i++
	}
	return true
}

// =============================
// Test if b[1..n] is a necklace
// =============================
func IsNecklace(b [MAX]uint8, n uint16) bool {
	var p uint16 = 1
	var i uint16 = 2
	for i <= n {
		if b[i-p] > b[i] {
			return false
		}
		if b[i-p] < b[i] {
			p = i
		}
		i++
	}
	return n%p == 0
}

// ===========================================
// Necklace Successor Rules
// ===========================================
func Granddaddy(a [MAX]uint8, n uint16) uint8 {
	var b [MAX]uint8
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
	if IsNecklace(b, n) {
		return 1 - a[1]
	}
	return a[1]
}
