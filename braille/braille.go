package braille

const (
	brailleStart = 0x2800 // 0x2800 - 0x28FF is reserved for braille in unicode
	size         = 8
)

var (
	// Braille patterns don't appear in the way one would naturally associate
	// with the bits. There first six form a sequence of their own, and the two
	// bottom dots are added as "extension" bits.
	brailleOrder = [size]int{
		1, 4,
		2, 5,
		3, 6,
		7, 8,
	}
)

// Get a braille representation for a given [8]bool. The array is a 4x2
// representation of the braille pattern.
func Get(dots [size]bool) string {
	offset := 0
	for idx, bit := range brailleOrder {
		if dots[idx] == true {
			offset |= (1 << (bit - 1))
		}
	}
	return string(brailleStart + offset)
}
