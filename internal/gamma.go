package internal

// Gamma8 converts 8-bit luminance y according to CIE 1976.
func Gamma8(y byte) byte {
	return gamma[y]
}

// Gamma table according to CIE 1976.
//
//	const max = 255
//
//	for i := 0; i <= max; i++ {
//		var x float64
//		y := 100 * float64(i) / max
//		if y > 8 {
//			x = math.Pow((y+16)/116, 3)
//		} else {
//			x = y / 903.3
//		}
//		fmt.Printf("\\x%02x", int(max*x+0.5))
//	}
const gamma = "" +
	"\x00\x00\x00\x00\x00\x01\x01\x01\x01\x01\x01\x01\x01\x01\x02\x02" +
	"\x02\x02\x02\x02\x02\x02\x02\x03\x03\x03\x03\x03\x03\x03\x03\x04" +
	"\x04\x04\x04\x04\x04\x05\x05\x05\x05\x05\x06\x06\x06\x06\x06\x07" +
	"\x07\x07\x07\x08\x08\x08\x08\x09\x09\x09\x0a\x0a\x0a\x0a\x0b\x0b" +
	"\x0b\x0c\x0c\x0c\x0d\x0d\x0d\x0e\x0e\x0f\x0f\x0f\x10\x10\x11\x11" +
	"\x11\x12\x12\x13\x13\x14\x14\x15\x15\x16\x16\x17\x17\x18\x18\x19" +
	"\x19\x1a\x1a\x1b\x1c\x1c\x1d\x1d\x1e\x1f\x1f\x20\x20\x21\x22\x22" +
	"\x23\x24\x25\x25\x26\x27\x27\x28\x29\x2a\x2b\x2b\x2c\x2d\x2e\x2f" +
	"\x2f\x30\x31\x32\x33\x34\x35\x36\x36\x37\x38\x39\x3a\x3b\x3c\x3d" +
	"\x3e\x3f\x40\x41\x42\x43\x44\x46\x47\x48\x49\x4a\x4b\x4c\x4d\x4f" +
	"\x50\x51\x52\x53\x55\x56\x57\x58\x5a\x5b\x5c\x5e\x5f\x60\x62\x63" +
	"\x64\x66\x67\x69\x6a\x6c\x6d\x6e\x70\x71\x73\x74\x76\x78\x79\x7b" +
	"\x7c\x7e\x80\x81\x83\x84\x86\x88\x8a\x8b\x8d\x8f\x91\x92\x94\x96" +
	"\x98\x9a\x9b\x9d\x9f\xa1\xa3\xa5\xa7\xa9\xab\xad\xaf\xb1\xb3\xb5" +
	"\xb7\xb9\xbb\xbd\xbf\xc1\xc4\xc6\xc8\xca\xcc\xcf\xd1\xd3\xd6\xd8" +
	"\xda\xdc\xdf\xe1\xe4\xe6\xe8\xeb\xed\xf0\xf2\xf5\xf7\xfa\xfc\xff"
