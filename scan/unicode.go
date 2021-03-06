package scan

import "unicode"

// isIdentifierStart reports whether r is a valid unicode ID_Start
// character according to unicode.org/reports/tr31
func isIdentifierStart(r rune) bool {
	switch r {
	case '_', '$',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return true
	}

	// All ASCII identifier start code points are listed above
	if r < 0x7F {
		return false
	}

	return unicode.Is(IDStart, r)
}

// isIdentifierContinue reports whether r is a valid unicode ID_Continue
// character according to unicode.org/reports/tr31
func isIdentifierContinue(r rune) bool {
	switch r {
	case '_', '$', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		return true
	}

	// All ASCII identifier start code points are listed above
	if r < 0x7F {
		return false
	}

	// ZWNJ and ZWJ are allowed in identifiers
	if r == 0x200C || r == 0x200D {
		return true
	}

	return unicode.Is(IDContinue, r)
}

// IDStart defines the characters derived from the Unicode General_Category of uppercase letters,
// lowercase letters, titlecase letters, modifier letters, other letters, letter numbers, plus Other_ID_Start,
// minus Pattern_Syntax and Pattern_White_Space code points
var IDStart = &unicode.RangeTable{
	LatinOffset: 117,
	R16: []unicode.Range16{
		{0x41, 0x5a, 1},
		{0x61, 0x7a, 1},
		{0xaa, 0xaa, 1},
		{0xb5, 0xb5, 1},
		{0xba, 0xba, 1},
		{0xc0, 0xd6, 1},
		{0xd8, 0xf6, 1},
		{0xf8, 0x2c1, 1},
		{0x2c6, 0x2d1, 1},
		{0x2e0, 0x2e4, 1},
		{0x2ec, 0x2ec, 1},
		{0x2ee, 0x2ee, 1},
		{0x370, 0x374, 1},
		{0x376, 0x377, 1},
		{0x37a, 0x37d, 1},
		{0x37f, 0x37f, 1},
		{0x386, 0x386, 1},
		{0x388, 0x38a, 1},
		{0x38c, 0x38c, 1},
		{0x38e, 0x3a1, 1},
		{0x3a3, 0x3f5, 1},
		{0x3f7, 0x481, 1},
		{0x48a, 0x52f, 1},
		{0x531, 0x556, 1},
		{0x559, 0x559, 1},
		{0x560, 0x588, 1},
		{0x5d0, 0x5ea, 1},
		{0x5ef, 0x5f2, 1},
		{0x620, 0x64a, 1},
		{0x66e, 0x66f, 1},
		{0x671, 0x6d3, 1},
		{0x6d5, 0x6d5, 1},
		{0x6e5, 0x6e6, 1},
		{0x6ee, 0x6ef, 1},
		{0x6fa, 0x6fc, 1},
		{0x6ff, 0x6ff, 1},
		{0x710, 0x710, 1},
		{0x712, 0x72f, 1},
		{0x74d, 0x7a5, 1},
		{0x7b1, 0x7b1, 1},
		{0x7ca, 0x7ea, 1},
		{0x7f4, 0x7f5, 1},
		{0x7fa, 0x7fa, 1},
		{0x800, 0x815, 1},
		{0x81a, 0x81a, 1},
		{0x824, 0x824, 1},
		{0x828, 0x828, 1},
		{0x840, 0x858, 1},
		{0x860, 0x86a, 1},
		{0x8a0, 0x8b4, 1},
		{0x8b6, 0x8c7, 1},
		{0x904, 0x939, 1},
		{0x93d, 0x93d, 1},
		{0x950, 0x950, 1},
		{0x958, 0x961, 1},
		{0x971, 0x980, 1},
		{0x985, 0x98c, 1},
		{0x98f, 0x990, 1},
		{0x993, 0x9a8, 1},
		{0x9aa, 0x9b0, 1},
		{0x9b2, 0x9b2, 1},
		{0x9b6, 0x9b9, 1},
		{0x9bd, 0x9bd, 1},
		{0x9ce, 0x9ce, 1},
		{0x9dc, 0x9dd, 1},
		{0x9df, 0x9e1, 1},
		{0x9f0, 0x9f1, 1},
		{0x9fc, 0x9fc, 1},
		{0xa05, 0xa0a, 1},
		{0xa0f, 0xa10, 1},
		{0xa13, 0xa28, 1},
		{0xa2a, 0xa30, 1},
		{0xa32, 0xa33, 1},
		{0xa35, 0xa36, 1},
		{0xa38, 0xa39, 1},
		{0xa59, 0xa5c, 1},
		{0xa5e, 0xa5e, 1},
		{0xa72, 0xa74, 1},
		{0xa85, 0xa8d, 1},
		{0xa8f, 0xa91, 1},
		{0xa93, 0xaa8, 1},
		{0xaaa, 0xab0, 1},
		{0xab2, 0xab3, 1},
		{0xab5, 0xab9, 1},
		{0xabd, 0xabd, 1},
		{0xad0, 0xad0, 1},
		{0xae0, 0xae1, 1},
		{0xaf9, 0xaf9, 1},
		{0xb05, 0xb0c, 1},
		{0xb0f, 0xb10, 1},
		{0xb13, 0xb28, 1},
		{0xb2a, 0xb30, 1},
		{0xb32, 0xb33, 1},
		{0xb35, 0xb39, 1},
		{0xb3d, 0xb3d, 1},
		{0xb5c, 0xb5d, 1},
		{0xb5f, 0xb61, 1},
		{0xb71, 0xb71, 1},
		{0xb83, 0xb83, 1},
		{0xb85, 0xb8a, 1},
		{0xb8e, 0xb90, 1},
		{0xb92, 0xb95, 1},
		{0xb99, 0xb9a, 1},
		{0xb9c, 0xb9c, 1},
		{0xb9e, 0xb9f, 1},
		{0xba3, 0xba4, 1},
		{0xba8, 0xbaa, 1},
		{0xbae, 0xbb9, 1},
		{0xbd0, 0xbd0, 1},
		{0xc05, 0xc0c, 1},
		{0xc0e, 0xc10, 1},
		{0xc12, 0xc28, 1},
		{0xc2a, 0xc39, 1},
		{0xc3d, 0xc3d, 1},
		{0xc58, 0xc5a, 1},
		{0xc60, 0xc61, 1},
		{0xc80, 0xc80, 1},
		{0xc85, 0xc8c, 1},
		{0xc8e, 0xc90, 1},
		{0xc92, 0xca8, 1},
		{0xcaa, 0xcb3, 1},
		{0xcb5, 0xcb9, 1},
		{0xcbd, 0xcbd, 1},
		{0xcde, 0xcde, 1},
		{0xce0, 0xce1, 1},
		{0xcf1, 0xcf2, 1},
		{0xd04, 0xd0c, 1},
		{0xd0e, 0xd10, 1},
		{0xd12, 0xd3a, 1},
		{0xd3d, 0xd3d, 1},
		{0xd4e, 0xd4e, 1},
		{0xd54, 0xd56, 1},
		{0xd5f, 0xd61, 1},
		{0xd7a, 0xd7f, 1},
		{0xd85, 0xd96, 1},
		{0xd9a, 0xdb1, 1},
		{0xdb3, 0xdbb, 1},
		{0xdbd, 0xdbd, 1},
		{0xdc0, 0xdc6, 1},
		{0xe01, 0xe30, 1},
		{0xe32, 0xe33, 1},
		{0xe40, 0xe46, 1},
		{0xe81, 0xe82, 1},
		{0xe84, 0xe84, 1},
		{0xe86, 0xe8a, 1},
		{0xe8c, 0xea3, 1},
		{0xea5, 0xea5, 1},
		{0xea7, 0xeb0, 1},
		{0xeb2, 0xeb3, 1},
		{0xebd, 0xebd, 1},
		{0xec0, 0xec4, 1},
		{0xec6, 0xec6, 1},
		{0xedc, 0xedf, 1},
		{0xf00, 0xf00, 1},
		{0xf40, 0xf47, 1},
		{0xf49, 0xf6c, 1},
		{0xf88, 0xf8c, 1},
	},
	R32: []unicode.Range32{
		{0x1000, 0x102a, 1},
		{0x103f, 0x103f, 1},
		{0x1050, 0x1055, 1},
		{0x105a, 0x105d, 1},
		{0x1061, 0x1061, 1},
		{0x1065, 0x1066, 1},
		{0x106e, 0x1070, 1},
		{0x1075, 0x1081, 1},
		{0x108e, 0x108e, 1},
		{0x10a0, 0x10c5, 1},
		{0x10c7, 0x10c7, 1},
		{0x10cd, 0x10cd, 1},
		{0x10d0, 0x10fa, 1},
		{0x10fc, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x1258, 1},
		{0x125a, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c0, 1},
		{0x12c2, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x1380, 0x138f, 1},
		{0x13a0, 0x13f5, 1},
		{0x13f8, 0x13fd, 1},
		{0x1401, 0x166c, 1},
		{0x166f, 0x167f, 1},
		{0x1681, 0x169a, 1},
		{0x16a0, 0x16ea, 1},
		{0x16ee, 0x16f8, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1711, 1},
		{0x1720, 0x1731, 1},
		{0x1740, 0x1751, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1780, 0x17b3, 1},
		{0x17d7, 0x17d7, 1},
		{0x17dc, 0x17dc, 1},
		{0x1820, 0x1878, 1},
		{0x1880, 0x18a8, 1},
		{0x18aa, 0x18aa, 1},
		{0x18b0, 0x18f5, 1},
		{0x1900, 0x191e, 1},
		{0x1950, 0x196d, 1},
		{0x1970, 0x1974, 1},
		{0x1980, 0x19ab, 1},
		{0x19b0, 0x19c9, 1},
		{0x1a00, 0x1a16, 1},
		{0x1a20, 0x1a54, 1},
		{0x1aa7, 0x1aa7, 1},
		{0x1b05, 0x1b33, 1},
		{0x1b45, 0x1b4b, 1},
		{0x1b83, 0x1ba0, 1},
		{0x1bae, 0x1baf, 1},
		{0x1bba, 0x1be5, 1},
		{0x1c00, 0x1c23, 1},
		{0x1c4d, 0x1c4f, 1},
		{0x1c5a, 0x1c7d, 1},
		{0x1c80, 0x1c88, 1},
		{0x1c90, 0x1cba, 1},
		{0x1cbd, 0x1cbf, 1},
		{0x1ce9, 0x1cec, 1},
		{0x1cee, 0x1cf3, 1},
		{0x1cf5, 0x1cf6, 1},
		{0x1cfa, 0x1cfa, 1},
		{0x1d00, 0x1dbf, 1},
		{0x1e00, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f59, 1},
		{0x1f5b, 0x1f5b, 1},
		{0x1f5d, 0x1f5d, 1},
		{0x1f5f, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fbc, 1},
		{0x1fbe, 0x1fbe, 1},
		{0x1fc2, 0x1fc4, 1},
		{0x1fc6, 0x1fcc, 1},
		{0x1fd0, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fe0, 0x1fec, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffc, 1},
		{0x2071, 0x2071, 1},
		{0x207f, 0x207f, 1},
		{0x2090, 0x209c, 1},
		{0x2102, 0x2102, 1},
		{0x2107, 0x2107, 1},
		{0x210a, 0x2113, 1},
		{0x2115, 0x2115, 1},
		{0x2118, 0x211d, 1},
		{0x2124, 0x2124, 1},
		{0x2126, 0x2126, 1},
		{0x2128, 0x2128, 1},
		{0x212a, 0x2139, 1},
		{0x213c, 0x213f, 1},
		{0x2145, 0x2149, 1},
		{0x214e, 0x214e, 1},
		{0x2160, 0x2188, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c60, 0x2ce4, 1},
		{0x2ceb, 0x2cee, 1},
		{0x2cf2, 0x2cf3, 1},
		{0x2d00, 0x2d25, 1},
		{0x2d27, 0x2d27, 1},
		{0x2d2d, 0x2d2d, 1},
		{0x2d30, 0x2d67, 1},
		{0x2d6f, 0x2d6f, 1},
		{0x2d80, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x3005, 0x3007, 1},
		{0x3021, 0x3029, 1},
		{0x3031, 0x3035, 1},
		{0x3038, 0x303c, 1},
		{0x3041, 0x3096, 1},
		{0x309b, 0x309f, 1},
		{0x30a1, 0x30fa, 1},
		{0x30fc, 0x30ff, 1},
		{0x3105, 0x312f, 1},
		{0x3131, 0x318e, 1},
		{0x31a0, 0x31bf, 1},
		{0x31f0, 0x31ff, 1},
		{0x3400, 0x4dbf, 1},
		{0x4e00, 0x9ffc, 1},
		{0xa000, 0xa48c, 1},
		{0xa4d0, 0xa4fd, 1},
		{0xa500, 0xa60c, 1},
		{0xa610, 0xa61f, 1},
		{0xa62a, 0xa62b, 1},
		{0xa640, 0xa66e, 1},
		{0xa67f, 0xa69d, 1},
		{0xa6a0, 0xa6ef, 1},
		{0xa717, 0xa71f, 1},
		{0xa722, 0xa788, 1},
		{0xa78b, 0xa7bf, 1},
		{0xa7c2, 0xa7ca, 1},
		{0xa7f5, 0xa801, 1},
		{0xa803, 0xa805, 1},
		{0xa807, 0xa80a, 1},
		{0xa80c, 0xa822, 1},
		{0xa840, 0xa873, 1},
		{0xa882, 0xa8b3, 1},
		{0xa8f2, 0xa8f7, 1},
		{0xa8fb, 0xa8fb, 1},
		{0xa8fd, 0xa8fe, 1},
		{0xa90a, 0xa925, 1},
		{0xa930, 0xa946, 1},
		{0xa960, 0xa97c, 1},
		{0xa984, 0xa9b2, 1},
		{0xa9cf, 0xa9cf, 1},
		{0xa9e0, 0xa9e4, 1},
		{0xa9e6, 0xa9ef, 1},
		{0xa9fa, 0xa9fe, 1},
		{0xaa00, 0xaa28, 1},
		{0xaa40, 0xaa42, 1},
		{0xaa44, 0xaa4b, 1},
		{0xaa60, 0xaa76, 1},
		{0xaa7a, 0xaa7a, 1},
		{0xaa7e, 0xaaaf, 1},
		{0xaab1, 0xaab1, 1},
		{0xaab5, 0xaab6, 1},
		{0xaab9, 0xaabd, 1},
		{0xaac0, 0xaac0, 1},
		{0xaac2, 0xaac2, 1},
		{0xaadb, 0xaadd, 1},
		{0xaae0, 0xaaea, 1},
		{0xaaf2, 0xaaf4, 1},
		{0xab01, 0xab06, 1},
		{0xab09, 0xab0e, 1},
		{0xab11, 0xab16, 1},
		{0xab20, 0xab26, 1},
		{0xab28, 0xab2e, 1},
		{0xab30, 0xab5a, 1},
		{0xab5c, 0xab69, 1},
		{0xab70, 0xabe2, 1},
		{0xac00, 0xd7a3, 1},
		{0xd7b0, 0xd7c6, 1},
		{0xd7cb, 0xd7fb, 1},
		{0xf900, 0xfa6d, 1},
		{0xfa70, 0xfad9, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb1d, 0xfb1d, 1},
		{0xfb1f, 0xfb28, 1},
		{0xfb2a, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb3e, 1},
		{0xfb40, 0xfb41, 1},
		{0xfb43, 0xfb44, 1},
		{0xfb46, 0xfbb1, 1},
		{0xfbd3, 0xfd3d, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfb, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xff21, 0xff3a, 1},
		{0xff41, 0xff5a, 1},
		{0xff66, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
		{0x10000, 0x1000b, 1},
		{0x1000d, 0x10026, 1},
		{0x10028, 0x1003a, 1},
		{0x1003c, 0x1003d, 1},
		{0x1003f, 0x1004d, 1},
		{0x10050, 0x1005d, 1},
		{0x10080, 0x100fa, 1},
		{0x10140, 0x10174, 1},
		{0x10280, 0x1029c, 1},
		{0x102a0, 0x102d0, 1},
		{0x10300, 0x1031f, 1},
		{0x1032d, 0x1034a, 1},
		{0x10350, 0x10375, 1},
		{0x10380, 0x1039d, 1},
		{0x103a0, 0x103c3, 1},
		{0x103c8, 0x103cf, 1},
		{0x103d1, 0x103d5, 1},
		{0x10400, 0x1049d, 1},
		{0x104b0, 0x104d3, 1},
		{0x104d8, 0x104fb, 1},
		{0x10500, 0x10527, 1},
		{0x10530, 0x10563, 1},
		{0x10600, 0x10736, 1},
		{0x10740, 0x10755, 1},
		{0x10760, 0x10767, 1},
		{0x10800, 0x10805, 1},
		{0x10808, 0x10808, 1},
		{0x1080a, 0x10835, 1},
		{0x10837, 0x10838, 1},
		{0x1083c, 0x1083c, 1},
		{0x1083f, 0x10855, 1},
		{0x10860, 0x10876, 1},
		{0x10880, 0x1089e, 1},
		{0x108e0, 0x108f2, 1},
		{0x108f4, 0x108f5, 1},
		{0x10900, 0x10915, 1},
		{0x10920, 0x10939, 1},
		{0x10980, 0x109b7, 1},
		{0x109be, 0x109bf, 1},
		{0x10a00, 0x10a00, 1},
		{0x10a10, 0x10a13, 1},
		{0x10a15, 0x10a17, 1},
		{0x10a19, 0x10a35, 1},
		{0x10a60, 0x10a7c, 1},
		{0x10a80, 0x10a9c, 1},
		{0x10ac0, 0x10ac7, 1},
		{0x10ac9, 0x10ae4, 1},
		{0x10b00, 0x10b35, 1},
		{0x10b40, 0x10b55, 1},
		{0x10b60, 0x10b72, 1},
		{0x10b80, 0x10b91, 1},
		{0x10c00, 0x10c48, 1},
		{0x10c80, 0x10cb2, 1},
		{0x10cc0, 0x10cf2, 1},
		{0x10d00, 0x10d23, 1},
		{0x10e80, 0x10ea9, 1},
		{0x10eb0, 0x10eb1, 1},
		{0x10f00, 0x10f1c, 1},
		{0x10f27, 0x10f27, 1},
		{0x10f30, 0x10f45, 1},
		{0x10fb0, 0x10fc4, 1},
		{0x10fe0, 0x10ff6, 1},
		{0x11003, 0x11037, 1},
		{0x11083, 0x110af, 1},
		{0x110d0, 0x110e8, 1},
		{0x11103, 0x11126, 1},
		{0x11144, 0x11144, 1},
		{0x11147, 0x11147, 1},
		{0x11150, 0x11172, 1},
		{0x11176, 0x11176, 1},
		{0x11183, 0x111b2, 1},
		{0x111c1, 0x111c4, 1},
		{0x111da, 0x111da, 1},
		{0x111dc, 0x111dc, 1},
		{0x11200, 0x11211, 1},
		{0x11213, 0x1122b, 1},
		{0x11280, 0x11286, 1},
		{0x11288, 0x11288, 1},
		{0x1128a, 0x1128d, 1},
		{0x1128f, 0x1129d, 1},
		{0x1129f, 0x112a8, 1},
		{0x112b0, 0x112de, 1},
		{0x11305, 0x1130c, 1},
		{0x1130f, 0x11310, 1},
		{0x11313, 0x11328, 1},
		{0x1132a, 0x11330, 1},
		{0x11332, 0x11333, 1},
		{0x11335, 0x11339, 1},
		{0x1133d, 0x1133d, 1},
		{0x11350, 0x11350, 1},
		{0x1135d, 0x11361, 1},
		{0x11400, 0x11434, 1},
		{0x11447, 0x1144a, 1},
		{0x1145f, 0x11461, 1},
		{0x11480, 0x114af, 1},
		{0x114c4, 0x114c5, 1},
		{0x114c7, 0x114c7, 1},
		{0x11580, 0x115ae, 1},
		{0x115d8, 0x115db, 1},
		{0x11600, 0x1162f, 1},
		{0x11644, 0x11644, 1},
		{0x11680, 0x116aa, 1},
		{0x116b8, 0x116b8, 1},
		{0x11700, 0x1171a, 1},
		{0x11800, 0x1182b, 1},
		{0x118a0, 0x118df, 1},
		{0x118ff, 0x11906, 1},
		{0x11909, 0x11909, 1},
		{0x1190c, 0x11913, 1},
		{0x11915, 0x11916, 1},
		{0x11918, 0x1192f, 1},
		{0x1193f, 0x1193f, 1},
		{0x11941, 0x11941, 1},
		{0x119a0, 0x119a7, 1},
		{0x119aa, 0x119d0, 1},
		{0x119e1, 0x119e1, 1},
		{0x119e3, 0x119e3, 1},
		{0x11a00, 0x11a00, 1},
		{0x11a0b, 0x11a32, 1},
		{0x11a3a, 0x11a3a, 1},
		{0x11a50, 0x11a50, 1},
		{0x11a5c, 0x11a89, 1},
		{0x11a9d, 0x11a9d, 1},
		{0x11ac0, 0x11af8, 1},
		{0x11c00, 0x11c08, 1},
		{0x11c0a, 0x11c2e, 1},
		{0x11c40, 0x11c40, 1},
		{0x11c72, 0x11c8f, 1},
		{0x11d00, 0x11d06, 1},
		{0x11d08, 0x11d09, 1},
		{0x11d0b, 0x11d30, 1},
		{0x11d46, 0x11d46, 1},
		{0x11d60, 0x11d65, 1},
		{0x11d67, 0x11d68, 1},
		{0x11d6a, 0x11d89, 1},
		{0x11d98, 0x11d98, 1},
		{0x11ee0, 0x11ef2, 1},
		{0x11fb0, 0x11fb0, 1},
		{0x12000, 0x12399, 1},
		{0x12400, 0x1246e, 1},
		{0x12480, 0x12543, 1},
		{0x13000, 0x1342e, 1},
		{0x14400, 0x14646, 1},
		{0x16800, 0x16a38, 1},
		{0x16a40, 0x16a5e, 1},
		{0x16ad0, 0x16aed, 1},
		{0x16b00, 0x16b2f, 1},
		{0x16b40, 0x16b43, 1},
		{0x16b63, 0x16b77, 1},
		{0x16b7d, 0x16b8f, 1},
		{0x16e40, 0x16e7f, 1},
		{0x16f00, 0x16f4a, 1},
		{0x16f50, 0x16f50, 1},
		{0x16f93, 0x16f9f, 1},
		{0x16fe0, 0x16fe1, 1},
		{0x16fe3, 0x16fe3, 1},
		{0x17000, 0x187f7, 1},
		{0x18800, 0x18cd5, 1},
		{0x18d00, 0x18d08, 1},
		{0x1b000, 0x1b11e, 1},
		{0x1b150, 0x1b152, 1},
		{0x1b164, 0x1b167, 1},
		{0x1b170, 0x1b2fb, 1},
		{0x1bc00, 0x1bc6a, 1},
		{0x1bc70, 0x1bc7c, 1},
		{0x1bc80, 0x1bc88, 1},
		{0x1bc90, 0x1bc99, 1},
		{0x1d400, 0x1d454, 1},
		{0x1d456, 0x1d49c, 1},
		{0x1d49e, 0x1d49f, 1},
		{0x1d4a2, 0x1d4a2, 1},
		{0x1d4a5, 0x1d4a6, 1},
		{0x1d4a9, 0x1d4ac, 1},
		{0x1d4ae, 0x1d4b9, 1},
		{0x1d4bb, 0x1d4bb, 1},
		{0x1d4bd, 0x1d4c3, 1},
		{0x1d4c5, 0x1d505, 1},
		{0x1d507, 0x1d50a, 1},
		{0x1d50d, 0x1d514, 1},
		{0x1d516, 0x1d51c, 1},
		{0x1d51e, 0x1d539, 1},
		{0x1d53b, 0x1d53e, 1},
		{0x1d540, 0x1d544, 1},
		{0x1d546, 0x1d546, 1},
		{0x1d54a, 0x1d550, 1},
		{0x1d552, 0x1d6a5, 1},
		{0x1d6a8, 0x1d6c0, 1},
		{0x1d6c2, 0x1d6da, 1},
		{0x1d6dc, 0x1d6fa, 1},
		{0x1d6fc, 0x1d714, 1},
		{0x1d716, 0x1d734, 1},
		{0x1d736, 0x1d74e, 1},
		{0x1d750, 0x1d76e, 1},
		{0x1d770, 0x1d788, 1},
		{0x1d78a, 0x1d7a8, 1},
		{0x1d7aa, 0x1d7c2, 1},
		{0x1d7c4, 0x1d7cb, 1},
		{0x1e100, 0x1e12c, 1},
		{0x1e137, 0x1e13d, 1},
		{0x1e14e, 0x1e14e, 1},
		{0x1e2c0, 0x1e2eb, 1},
		{0x1e800, 0x1e8c4, 1},
		{0x1e900, 0x1e943, 1},
		{0x1e94b, 0x1e94b, 1},
		{0x1ee00, 0x1ee03, 1},
		{0x1ee05, 0x1ee1f, 1},
		{0x1ee21, 0x1ee22, 1},
		{0x1ee24, 0x1ee24, 1},
		{0x1ee27, 0x1ee27, 1},
		{0x1ee29, 0x1ee32, 1},
		{0x1ee34, 0x1ee37, 1},
		{0x1ee39, 0x1ee39, 1},
		{0x1ee3b, 0x1ee3b, 1},
		{0x1ee42, 0x1ee42, 1},
		{0x1ee47, 0x1ee47, 1},
		{0x1ee49, 0x1ee49, 1},
		{0x1ee4b, 0x1ee4b, 1},
		{0x1ee4d, 0x1ee4f, 1},
		{0x1ee51, 0x1ee52, 1},
		{0x1ee54, 0x1ee54, 1},
		{0x1ee57, 0x1ee57, 1},
		{0x1ee59, 0x1ee59, 1},
		{0x1ee5b, 0x1ee5b, 1},
		{0x1ee5d, 0x1ee5d, 1},
		{0x1ee5f, 0x1ee5f, 1},
		{0x1ee61, 0x1ee62, 1},
		{0x1ee64, 0x1ee64, 1},
		{0x1ee67, 0x1ee6a, 1},
		{0x1ee6c, 0x1ee72, 1},
		{0x1ee74, 0x1ee77, 1},
		{0x1ee79, 0x1ee7c, 1},
		{0x1ee7e, 0x1ee7e, 1},
		{0x1ee80, 0x1ee89, 1},
		{0x1ee8b, 0x1ee9b, 1},
		{0x1eea1, 0x1eea3, 1},
		{0x1eea5, 0x1eea9, 1},
		{0x1eeab, 0x1eebb, 1},
		{0x20000, 0x2a6dd, 1},
		{0x2a700, 0x2b734, 1},
		{0x2b740, 0x2b81d, 1},
		{0x2b820, 0x2cea1, 1},
		{0x2ceb0, 0x2ebe0, 1},
		{0x2f800, 0x2fa1d, 1},
		{0x30000, 0x3134a, 1},
	},
}

// IDContinue defines characters including ID_Start characters, plus characters having the
// Unicode General_Category of nonspacing marks, spacing combining marks, decimal number, connector punctuation,
// plus Other_ID_Continue , minus Pattern_Syntax and Pattern_White_Space code points.
var IDContinue = &unicode.RangeTable{
	LatinOffset: 129,
	R16: []unicode.Range16{
		{0x30, 0x39, 1},
		{0x41, 0x5a, 1},
		{0x5f, 0x5f, 1},
		{0x61, 0x7a, 1},
		{0xaa, 0xaa, 1},
		{0xb5, 0xb5, 1},
		{0xb7, 0xb7, 1},
		{0xba, 0xba, 1},
		{0xc0, 0xd6, 1},
		{0xd8, 0xf6, 1},
		{0xf8, 0x2c1, 1},
		{0x2c6, 0x2d1, 1},
		{0x2e0, 0x2e4, 1},
		{0x2ec, 0x2ec, 1},
		{0x2ee, 0x2ee, 1},
		{0x300, 0x374, 1},
		{0x376, 0x377, 1},
		{0x37a, 0x37d, 1},
		{0x37f, 0x37f, 1},
		{0x386, 0x38a, 1},
		{0x38c, 0x38c, 1},
		{0x38e, 0x3a1, 1},
		{0x3a3, 0x3f5, 1},
		{0x3f7, 0x481, 1},
		{0x483, 0x487, 1},
		{0x48a, 0x52f, 1},
		{0x531, 0x556, 1},
		{0x559, 0x559, 1},
		{0x560, 0x588, 1},
		{0x591, 0x5bd, 1},
		{0x5bf, 0x5bf, 1},
		{0x5c1, 0x5c2, 1},
		{0x5c4, 0x5c5, 1},
		{0x5c7, 0x5c7, 1},
		{0x5d0, 0x5ea, 1},
		{0x5ef, 0x5f2, 1},
		{0x610, 0x61a, 1},
		{0x620, 0x669, 1},
		{0x66e, 0x6d3, 1},
		{0x6d5, 0x6dc, 1},
		{0x6df, 0x6e8, 1},
		{0x6ea, 0x6fc, 1},
		{0x6ff, 0x6ff, 1},
		{0x710, 0x74a, 1},
		{0x74d, 0x7b1, 1},
		{0x7c0, 0x7f5, 1},
		{0x7fa, 0x7fa, 1},
		{0x7fd, 0x7fd, 1},
		{0x800, 0x82d, 1},
		{0x840, 0x85b, 1},
		{0x860, 0x86a, 1},
		{0x8a0, 0x8b4, 1},
		{0x8b6, 0x8c7, 1},
		{0x8d3, 0x8e1, 1},
		{0x8e3, 0x963, 1},
		{0x966, 0x96f, 1},
		{0x971, 0x983, 1},
		{0x985, 0x98c, 1},
		{0x98f, 0x990, 1},
		{0x993, 0x9a8, 1},
		{0x9aa, 0x9b0, 1},
		{0x9b2, 0x9b2, 1},
		{0x9b6, 0x9b9, 1},
		{0x9bc, 0x9c4, 1},
		{0x9c7, 0x9c8, 1},
		{0x9cb, 0x9ce, 1},
		{0x9d7, 0x9d7, 1},
		{0x9dc, 0x9dd, 1},
		{0x9df, 0x9e3, 1},
		{0x9e6, 0x9f1, 1},
		{0x9fc, 0x9fc, 1},
		{0x9fe, 0x9fe, 1},
		{0xa01, 0xa03, 1},
		{0xa05, 0xa0a, 1},
		{0xa0f, 0xa10, 1},
		{0xa13, 0xa28, 1},
		{0xa2a, 0xa30, 1},
		{0xa32, 0xa33, 1},
		{0xa35, 0xa36, 1},
		{0xa38, 0xa39, 1},
		{0xa3c, 0xa3c, 1},
		{0xa3e, 0xa42, 1},
		{0xa47, 0xa48, 1},
		{0xa4b, 0xa4d, 1},
		{0xa51, 0xa51, 1},
		{0xa59, 0xa5c, 1},
		{0xa5e, 0xa5e, 1},
		{0xa66, 0xa75, 1},
		{0xa81, 0xa83, 1},
		{0xa85, 0xa8d, 1},
		{0xa8f, 0xa91, 1},
		{0xa93, 0xaa8, 1},
		{0xaaa, 0xab0, 1},
		{0xab2, 0xab3, 1},
		{0xab5, 0xab9, 1},
		{0xabc, 0xac5, 1},
		{0xac7, 0xac9, 1},
		{0xacb, 0xacd, 1},
		{0xad0, 0xad0, 1},
		{0xae0, 0xae3, 1},
		{0xae6, 0xaef, 1},
		{0xaf9, 0xaff, 1},
		{0xb01, 0xb03, 1},
		{0xb05, 0xb0c, 1},
		{0xb0f, 0xb10, 1},
		{0xb13, 0xb28, 1},
		{0xb2a, 0xb30, 1},
		{0xb32, 0xb33, 1},
		{0xb35, 0xb39, 1},
		{0xb3c, 0xb44, 1},
		{0xb47, 0xb48, 1},
		{0xb4b, 0xb4d, 1},
		{0xb55, 0xb57, 1},
		{0xb5c, 0xb5d, 1},
		{0xb5f, 0xb63, 1},
		{0xb66, 0xb6f, 1},
		{0xb71, 0xb71, 1},
		{0xb82, 0xb83, 1},
		{0xb85, 0xb8a, 1},
		{0xb8e, 0xb90, 1},
		{0xb92, 0xb95, 1},
		{0xb99, 0xb9a, 1},
		{0xb9c, 0xb9c, 1},
		{0xb9e, 0xb9f, 1},
		{0xba3, 0xba4, 1},
		{0xba8, 0xbaa, 1},
		{0xbae, 0xbb9, 1},
		{0xbbe, 0xbc2, 1},
		{0xbc6, 0xbc8, 1},
		{0xbca, 0xbcd, 1},
		{0xbd0, 0xbd0, 1},
		{0xbd7, 0xbd7, 1},
		{0xbe6, 0xbef, 1},
		{0xc00, 0xc0c, 1},
		{0xc0e, 0xc10, 1},
		{0xc12, 0xc28, 1},
		{0xc2a, 0xc39, 1},
		{0xc3d, 0xc44, 1},
		{0xc46, 0xc48, 1},
		{0xc4a, 0xc4d, 1},
		{0xc55, 0xc56, 1},
		{0xc58, 0xc5a, 1},
		{0xc60, 0xc63, 1},
		{0xc66, 0xc6f, 1},
		{0xc80, 0xc83, 1},
		{0xc85, 0xc8c, 1},
		{0xc8e, 0xc90, 1},
		{0xc92, 0xca8, 1},
		{0xcaa, 0xcb3, 1},
		{0xcb5, 0xcb9, 1},
		{0xcbc, 0xcc4, 1},
		{0xcc6, 0xcc8, 1},
		{0xcca, 0xccd, 1},
		{0xcd5, 0xcd6, 1},
		{0xcde, 0xcde, 1},
		{0xce0, 0xce3, 1},
		{0xce6, 0xcef, 1},
		{0xcf1, 0xcf2, 1},
		{0xd00, 0xd0c, 1},
		{0xd0e, 0xd10, 1},
		{0xd12, 0xd44, 1},
		{0xd46, 0xd48, 1},
		{0xd4a, 0xd4e, 1},
		{0xd54, 0xd57, 1},
		{0xd5f, 0xd63, 1},
		{0xd66, 0xd6f, 1},
		{0xd7a, 0xd7f, 1},
		{0xd81, 0xd83, 1},
		{0xd85, 0xd96, 1},
		{0xd9a, 0xdb1, 1},
		{0xdb3, 0xdbb, 1},
		{0xdbd, 0xdbd, 1},
		{0xdc0, 0xdc6, 1},
		{0xdca, 0xdca, 1},
		{0xdcf, 0xdd4, 1},
		{0xdd6, 0xdd6, 1},
		{0xdd8, 0xddf, 1},
		{0xde6, 0xdef, 1},
		{0xdf2, 0xdf3, 1},
		{0xe01, 0xe3a, 1},
		{0xe40, 0xe4e, 1},
		{0xe50, 0xe59, 1},
		{0xe81, 0xe82, 1},
		{0xe84, 0xe84, 1},
		{0xe86, 0xe8a, 1},
		{0xe8c, 0xea3, 1},
		{0xea5, 0xea5, 1},
		{0xea7, 0xebd, 1},
		{0xec0, 0xec4, 1},
		{0xec6, 0xec6, 1},
		{0xec8, 0xecd, 1},
		{0xed0, 0xed9, 1},
		{0xedc, 0xedf, 1},
		{0xf00, 0xf00, 1},
		{0xf18, 0xf19, 1},
		{0xf20, 0xf29, 1},
		{0xf35, 0xf35, 1},
		{0xf37, 0xf37, 1},
		{0xf39, 0xf39, 1},
		{0xf3e, 0xf47, 1},
		{0xf49, 0xf6c, 1},
		{0xf71, 0xf84, 1},
		{0xf86, 0xf97, 1},
		{0xf99, 0xfbc, 1},
		{0xfc6, 0xfc6, 1},
	},
	R32: []unicode.Range32{
		{0x1000, 0x1049, 1},
		{0x1050, 0x109d, 1},
		{0x10a0, 0x10c5, 1},
		{0x10c7, 0x10c7, 1},
		{0x10cd, 0x10cd, 1},
		{0x10d0, 0x10fa, 1},
		{0x10fc, 0x1248, 1},
		{0x124a, 0x124d, 1},
		{0x1250, 0x1256, 1},
		{0x1258, 0x1258, 1},
		{0x125a, 0x125d, 1},
		{0x1260, 0x1288, 1},
		{0x128a, 0x128d, 1},
		{0x1290, 0x12b0, 1},
		{0x12b2, 0x12b5, 1},
		{0x12b8, 0x12be, 1},
		{0x12c0, 0x12c0, 1},
		{0x12c2, 0x12c5, 1},
		{0x12c8, 0x12d6, 1},
		{0x12d8, 0x1310, 1},
		{0x1312, 0x1315, 1},
		{0x1318, 0x135a, 1},
		{0x135d, 0x135f, 1},
		{0x1369, 0x1371, 1},
		{0x1380, 0x138f, 1},
		{0x13a0, 0x13f5, 1},
		{0x13f8, 0x13fd, 1},
		{0x1401, 0x166c, 1},
		{0x166f, 0x167f, 1},
		{0x1681, 0x169a, 1},
		{0x16a0, 0x16ea, 1},
		{0x16ee, 0x16f8, 1},
		{0x1700, 0x170c, 1},
		{0x170e, 0x1714, 1},
		{0x1720, 0x1734, 1},
		{0x1740, 0x1753, 1},
		{0x1760, 0x176c, 1},
		{0x176e, 0x1770, 1},
		{0x1772, 0x1773, 1},
		{0x1780, 0x17d3, 1},
		{0x17d7, 0x17d7, 1},
		{0x17dc, 0x17dd, 1},
		{0x17e0, 0x17e9, 1},
		{0x180b, 0x180d, 1},
		{0x1810, 0x1819, 1},
		{0x1820, 0x1878, 1},
		{0x1880, 0x18aa, 1},
		{0x18b0, 0x18f5, 1},
		{0x1900, 0x191e, 1},
		{0x1920, 0x192b, 1},
		{0x1930, 0x193b, 1},
		{0x1946, 0x196d, 1},
		{0x1970, 0x1974, 1},
		{0x1980, 0x19ab, 1},
		{0x19b0, 0x19c9, 1},
		{0x19d0, 0x19da, 1},
		{0x1a00, 0x1a1b, 1},
		{0x1a20, 0x1a5e, 1},
		{0x1a60, 0x1a7c, 1},
		{0x1a7f, 0x1a89, 1},
		{0x1a90, 0x1a99, 1},
		{0x1aa7, 0x1aa7, 1},
		{0x1ab0, 0x1abd, 1},
		{0x1abf, 0x1ac0, 1},
		{0x1b00, 0x1b4b, 1},
		{0x1b50, 0x1b59, 1},
		{0x1b6b, 0x1b73, 1},
		{0x1b80, 0x1bf3, 1},
		{0x1c00, 0x1c37, 1},
		{0x1c40, 0x1c49, 1},
		{0x1c4d, 0x1c7d, 1},
		{0x1c80, 0x1c88, 1},
		{0x1c90, 0x1cba, 1},
		{0x1cbd, 0x1cbf, 1},
		{0x1cd0, 0x1cd2, 1},
		{0x1cd4, 0x1cfa, 1},
		{0x1d00, 0x1df9, 1},
		{0x1dfb, 0x1f15, 1},
		{0x1f18, 0x1f1d, 1},
		{0x1f20, 0x1f45, 1},
		{0x1f48, 0x1f4d, 1},
		{0x1f50, 0x1f57, 1},
		{0x1f59, 0x1f59, 1},
		{0x1f5b, 0x1f5b, 1},
		{0x1f5d, 0x1f5d, 1},
		{0x1f5f, 0x1f7d, 1},
		{0x1f80, 0x1fb4, 1},
		{0x1fb6, 0x1fbc, 1},
		{0x1fbe, 0x1fbe, 1},
		{0x1fc2, 0x1fc4, 1},
		{0x1fc6, 0x1fcc, 1},
		{0x1fd0, 0x1fd3, 1},
		{0x1fd6, 0x1fdb, 1},
		{0x1fe0, 0x1fec, 1},
		{0x1ff2, 0x1ff4, 1},
		{0x1ff6, 0x1ffc, 1},
		{0x203f, 0x2040, 1},
		{0x2054, 0x2054, 1},
		{0x2071, 0x2071, 1},
		{0x207f, 0x207f, 1},
		{0x2090, 0x209c, 1},
		{0x20d0, 0x20dc, 1},
		{0x20e1, 0x20e1, 1},
		{0x20e5, 0x20f0, 1},
		{0x2102, 0x2102, 1},
		{0x2107, 0x2107, 1},
		{0x210a, 0x2113, 1},
		{0x2115, 0x2115, 1},
		{0x2118, 0x211d, 1},
		{0x2124, 0x2124, 1},
		{0x2126, 0x2126, 1},
		{0x2128, 0x2128, 1},
		{0x212a, 0x2139, 1},
		{0x213c, 0x213f, 1},
		{0x2145, 0x2149, 1},
		{0x214e, 0x214e, 1},
		{0x2160, 0x2188, 1},
		{0x2c00, 0x2c2e, 1},
		{0x2c30, 0x2c5e, 1},
		{0x2c60, 0x2ce4, 1},
		{0x2ceb, 0x2cf3, 1},
		{0x2d00, 0x2d25, 1},
		{0x2d27, 0x2d27, 1},
		{0x2d2d, 0x2d2d, 1},
		{0x2d30, 0x2d67, 1},
		{0x2d6f, 0x2d6f, 1},
		{0x2d7f, 0x2d96, 1},
		{0x2da0, 0x2da6, 1},
		{0x2da8, 0x2dae, 1},
		{0x2db0, 0x2db6, 1},
		{0x2db8, 0x2dbe, 1},
		{0x2dc0, 0x2dc6, 1},
		{0x2dc8, 0x2dce, 1},
		{0x2dd0, 0x2dd6, 1},
		{0x2dd8, 0x2dde, 1},
		{0x2de0, 0x2dff, 1},
		{0x3005, 0x3007, 1},
		{0x3021, 0x302f, 1},
		{0x3031, 0x3035, 1},
		{0x3038, 0x303c, 1},
		{0x3041, 0x3096, 1},
		{0x3099, 0x309f, 1},
		{0x30a1, 0x30fa, 1},
		{0x30fc, 0x30ff, 1},
		{0x3105, 0x312f, 1},
		{0x3131, 0x318e, 1},
		{0x31a0, 0x31bf, 1},
		{0x31f0, 0x31ff, 1},
		{0x3400, 0x4dbf, 1},
		{0x4e00, 0x9ffc, 1},
		{0xa000, 0xa48c, 1},
		{0xa4d0, 0xa4fd, 1},
		{0xa500, 0xa60c, 1},
		{0xa610, 0xa62b, 1},
		{0xa640, 0xa66f, 1},
		{0xa674, 0xa67d, 1},
		{0xa67f, 0xa6f1, 1},
		{0xa717, 0xa71f, 1},
		{0xa722, 0xa788, 1},
		{0xa78b, 0xa7bf, 1},
		{0xa7c2, 0xa7ca, 1},
		{0xa7f5, 0xa827, 1},
		{0xa82c, 0xa82c, 1},
		{0xa840, 0xa873, 1},
		{0xa880, 0xa8c5, 1},
		{0xa8d0, 0xa8d9, 1},
		{0xa8e0, 0xa8f7, 1},
		{0xa8fb, 0xa8fb, 1},
		{0xa8fd, 0xa92d, 1},
		{0xa930, 0xa953, 1},
		{0xa960, 0xa97c, 1},
		{0xa980, 0xa9c0, 1},
		{0xa9cf, 0xa9d9, 1},
		{0xa9e0, 0xa9fe, 1},
		{0xaa00, 0xaa36, 1},
		{0xaa40, 0xaa4d, 1},
		{0xaa50, 0xaa59, 1},
		{0xaa60, 0xaa76, 1},
		{0xaa7a, 0xaac2, 1},
		{0xaadb, 0xaadd, 1},
		{0xaae0, 0xaaef, 1},
		{0xaaf2, 0xaaf6, 1},
		{0xab01, 0xab06, 1},
		{0xab09, 0xab0e, 1},
		{0xab11, 0xab16, 1},
		{0xab20, 0xab26, 1},
		{0xab28, 0xab2e, 1},
		{0xab30, 0xab5a, 1},
		{0xab5c, 0xab69, 1},
		{0xab70, 0xabea, 1},
		{0xabec, 0xabed, 1},
		{0xabf0, 0xabf9, 1},
		{0xac00, 0xd7a3, 1},
		{0xd7b0, 0xd7c6, 1},
		{0xd7cb, 0xd7fb, 1},
		{0xf900, 0xfa6d, 1},
		{0xfa70, 0xfad9, 1},
		{0xfb00, 0xfb06, 1},
		{0xfb13, 0xfb17, 1},
		{0xfb1d, 0xfb28, 1},
		{0xfb2a, 0xfb36, 1},
		{0xfb38, 0xfb3c, 1},
		{0xfb3e, 0xfb3e, 1},
		{0xfb40, 0xfb41, 1},
		{0xfb43, 0xfb44, 1},
		{0xfb46, 0xfbb1, 1},
		{0xfbd3, 0xfd3d, 1},
		{0xfd50, 0xfd8f, 1},
		{0xfd92, 0xfdc7, 1},
		{0xfdf0, 0xfdfb, 1},
		{0xfe00, 0xfe0f, 1},
		{0xfe20, 0xfe2f, 1},
		{0xfe33, 0xfe34, 1},
		{0xfe4d, 0xfe4f, 1},
		{0xfe70, 0xfe74, 1},
		{0xfe76, 0xfefc, 1},
		{0xff10, 0xff19, 1},
		{0xff21, 0xff3a, 1},
		{0xff3f, 0xff3f, 1},
		{0xff41, 0xff5a, 1},
		{0xff66, 0xffbe, 1},
		{0xffc2, 0xffc7, 1},
		{0xffca, 0xffcf, 1},
		{0xffd2, 0xffd7, 1},
		{0xffda, 0xffdc, 1},
		{0x10000, 0x1000b, 1},
		{0x1000d, 0x10026, 1},
		{0x10028, 0x1003a, 1},
		{0x1003c, 0x1003d, 1},
		{0x1003f, 0x1004d, 1},
		{0x10050, 0x1005d, 1},
		{0x10080, 0x100fa, 1},
		{0x10140, 0x10174, 1},
		{0x101fd, 0x101fd, 1},
		{0x10280, 0x1029c, 1},
		{0x102a0, 0x102d0, 1},
		{0x102e0, 0x102e0, 1},
		{0x10300, 0x1031f, 1},
		{0x1032d, 0x1034a, 1},
		{0x10350, 0x1037a, 1},
		{0x10380, 0x1039d, 1},
		{0x103a0, 0x103c3, 1},
		{0x103c8, 0x103cf, 1},
		{0x103d1, 0x103d5, 1},
		{0x10400, 0x1049d, 1},
		{0x104a0, 0x104a9, 1},
		{0x104b0, 0x104d3, 1},
		{0x104d8, 0x104fb, 1},
		{0x10500, 0x10527, 1},
		{0x10530, 0x10563, 1},
		{0x10600, 0x10736, 1},
		{0x10740, 0x10755, 1},
		{0x10760, 0x10767, 1},
		{0x10800, 0x10805, 1},
		{0x10808, 0x10808, 1},
		{0x1080a, 0x10835, 1},
		{0x10837, 0x10838, 1},
		{0x1083c, 0x1083c, 1},
		{0x1083f, 0x10855, 1},
		{0x10860, 0x10876, 1},
		{0x10880, 0x1089e, 1},
		{0x108e0, 0x108f2, 1},
		{0x108f4, 0x108f5, 1},
		{0x10900, 0x10915, 1},
		{0x10920, 0x10939, 1},
		{0x10980, 0x109b7, 1},
		{0x109be, 0x109bf, 1},
		{0x10a00, 0x10a03, 1},
		{0x10a05, 0x10a06, 1},
		{0x10a0c, 0x10a13, 1},
		{0x10a15, 0x10a17, 1},
		{0x10a19, 0x10a35, 1},
		{0x10a38, 0x10a3a, 1},
		{0x10a3f, 0x10a3f, 1},
		{0x10a60, 0x10a7c, 1},
		{0x10a80, 0x10a9c, 1},
		{0x10ac0, 0x10ac7, 1},
		{0x10ac9, 0x10ae6, 1},
		{0x10b00, 0x10b35, 1},
		{0x10b40, 0x10b55, 1},
		{0x10b60, 0x10b72, 1},
		{0x10b80, 0x10b91, 1},
		{0x10c00, 0x10c48, 1},
		{0x10c80, 0x10cb2, 1},
		{0x10cc0, 0x10cf2, 1},
		{0x10d00, 0x10d27, 1},
		{0x10d30, 0x10d39, 1},
		{0x10e80, 0x10ea9, 1},
		{0x10eab, 0x10eac, 1},
		{0x10eb0, 0x10eb1, 1},
		{0x10f00, 0x10f1c, 1},
		{0x10f27, 0x10f27, 1},
		{0x10f30, 0x10f50, 1},
		{0x10fb0, 0x10fc4, 1},
		{0x10fe0, 0x10ff6, 1},
		{0x11000, 0x11046, 1},
		{0x11066, 0x1106f, 1},
		{0x1107f, 0x110ba, 1},
		{0x110d0, 0x110e8, 1},
		{0x110f0, 0x110f9, 1},
		{0x11100, 0x11134, 1},
		{0x11136, 0x1113f, 1},
		{0x11144, 0x11147, 1},
		{0x11150, 0x11173, 1},
		{0x11176, 0x11176, 1},
		{0x11180, 0x111c4, 1},
		{0x111c9, 0x111cc, 1},
		{0x111ce, 0x111da, 1},
		{0x111dc, 0x111dc, 1},
		{0x11200, 0x11211, 1},
		{0x11213, 0x11237, 1},
		{0x1123e, 0x1123e, 1},
		{0x11280, 0x11286, 1},
		{0x11288, 0x11288, 1},
		{0x1128a, 0x1128d, 1},
		{0x1128f, 0x1129d, 1},
		{0x1129f, 0x112a8, 1},
		{0x112b0, 0x112ea, 1},
		{0x112f0, 0x112f9, 1},
		{0x11300, 0x11303, 1},
		{0x11305, 0x1130c, 1},
		{0x1130f, 0x11310, 1},
		{0x11313, 0x11328, 1},
		{0x1132a, 0x11330, 1},
		{0x11332, 0x11333, 1},
		{0x11335, 0x11339, 1},
		{0x1133b, 0x11344, 1},
		{0x11347, 0x11348, 1},
		{0x1134b, 0x1134d, 1},
		{0x11350, 0x11350, 1},
		{0x11357, 0x11357, 1},
		{0x1135d, 0x11363, 1},
		{0x11366, 0x1136c, 1},
		{0x11370, 0x11374, 1},
		{0x11400, 0x1144a, 1},
		{0x11450, 0x11459, 1},
		{0x1145e, 0x11461, 1},
		{0x11480, 0x114c5, 1},
		{0x114c7, 0x114c7, 1},
		{0x114d0, 0x114d9, 1},
		{0x11580, 0x115b5, 1},
		{0x115b8, 0x115c0, 1},
		{0x115d8, 0x115dd, 1},
		{0x11600, 0x11640, 1},
		{0x11644, 0x11644, 1},
		{0x11650, 0x11659, 1},
		{0x11680, 0x116b8, 1},
		{0x116c0, 0x116c9, 1},
		{0x11700, 0x1171a, 1},
		{0x1171d, 0x1172b, 1},
		{0x11730, 0x11739, 1},
		{0x11800, 0x1183a, 1},
		{0x118a0, 0x118e9, 1},
		{0x118ff, 0x11906, 1},
		{0x11909, 0x11909, 1},
		{0x1190c, 0x11913, 1},
		{0x11915, 0x11916, 1},
		{0x11918, 0x11935, 1},
		{0x11937, 0x11938, 1},
		{0x1193b, 0x11943, 1},
		{0x11950, 0x11959, 1},
		{0x119a0, 0x119a7, 1},
		{0x119aa, 0x119d7, 1},
		{0x119da, 0x119e1, 1},
		{0x119e3, 0x119e4, 1},
		{0x11a00, 0x11a3e, 1},
		{0x11a47, 0x11a47, 1},
		{0x11a50, 0x11a99, 1},
		{0x11a9d, 0x11a9d, 1},
		{0x11ac0, 0x11af8, 1},
		{0x11c00, 0x11c08, 1},
		{0x11c0a, 0x11c36, 1},
		{0x11c38, 0x11c40, 1},
		{0x11c50, 0x11c59, 1},
		{0x11c72, 0x11c8f, 1},
		{0x11c92, 0x11ca7, 1},
		{0x11ca9, 0x11cb6, 1},
		{0x11d00, 0x11d06, 1},
		{0x11d08, 0x11d09, 1},
		{0x11d0b, 0x11d36, 1},
		{0x11d3a, 0x11d3a, 1},
		{0x11d3c, 0x11d3d, 1},
		{0x11d3f, 0x11d47, 1},
		{0x11d50, 0x11d59, 1},
		{0x11d60, 0x11d65, 1},
		{0x11d67, 0x11d68, 1},
		{0x11d6a, 0x11d8e, 1},
		{0x11d90, 0x11d91, 1},
		{0x11d93, 0x11d98, 1},
		{0x11da0, 0x11da9, 1},
		{0x11ee0, 0x11ef6, 1},
		{0x11fb0, 0x11fb0, 1},
		{0x12000, 0x12399, 1},
		{0x12400, 0x1246e, 1},
		{0x12480, 0x12543, 1},
		{0x13000, 0x1342e, 1},
		{0x14400, 0x14646, 1},
		{0x16800, 0x16a38, 1},
		{0x16a40, 0x16a5e, 1},
		{0x16a60, 0x16a69, 1},
		{0x16ad0, 0x16aed, 1},
		{0x16af0, 0x16af4, 1},
		{0x16b00, 0x16b36, 1},
		{0x16b40, 0x16b43, 1},
		{0x16b50, 0x16b59, 1},
		{0x16b63, 0x16b77, 1},
		{0x16b7d, 0x16b8f, 1},
		{0x16e40, 0x16e7f, 1},
		{0x16f00, 0x16f4a, 1},
		{0x16f4f, 0x16f87, 1},
		{0x16f8f, 0x16f9f, 1},
		{0x16fe0, 0x16fe1, 1},
		{0x16fe3, 0x16fe4, 1},
		{0x16ff0, 0x16ff1, 1},
		{0x17000, 0x187f7, 1},
		{0x18800, 0x18cd5, 1},
		{0x18d00, 0x18d08, 1},
		{0x1b000, 0x1b11e, 1},
		{0x1b150, 0x1b152, 1},
		{0x1b164, 0x1b167, 1},
		{0x1b170, 0x1b2fb, 1},
		{0x1bc00, 0x1bc6a, 1},
		{0x1bc70, 0x1bc7c, 1},
		{0x1bc80, 0x1bc88, 1},
		{0x1bc90, 0x1bc99, 1},
		{0x1bc9d, 0x1bc9e, 1},
		{0x1d165, 0x1d169, 1},
		{0x1d16d, 0x1d172, 1},
		{0x1d17b, 0x1d182, 1},
		{0x1d185, 0x1d18b, 1},
		{0x1d1aa, 0x1d1ad, 1},
		{0x1d242, 0x1d244, 1},
		{0x1d400, 0x1d454, 1},
		{0x1d456, 0x1d49c, 1},
		{0x1d49e, 0x1d49f, 1},
		{0x1d4a2, 0x1d4a2, 1},
		{0x1d4a5, 0x1d4a6, 1},
		{0x1d4a9, 0x1d4ac, 1},
		{0x1d4ae, 0x1d4b9, 1},
		{0x1d4bb, 0x1d4bb, 1},
		{0x1d4bd, 0x1d4c3, 1},
		{0x1d4c5, 0x1d505, 1},
		{0x1d507, 0x1d50a, 1},
		{0x1d50d, 0x1d514, 1},
		{0x1d516, 0x1d51c, 1},
		{0x1d51e, 0x1d539, 1},
		{0x1d53b, 0x1d53e, 1},
		{0x1d540, 0x1d544, 1},
		{0x1d546, 0x1d546, 1},
		{0x1d54a, 0x1d550, 1},
		{0x1d552, 0x1d6a5, 1},
		{0x1d6a8, 0x1d6c0, 1},
		{0x1d6c2, 0x1d6da, 1},
		{0x1d6dc, 0x1d6fa, 1},
		{0x1d6fc, 0x1d714, 1},
		{0x1d716, 0x1d734, 1},
		{0x1d736, 0x1d74e, 1},
		{0x1d750, 0x1d76e, 1},
		{0x1d770, 0x1d788, 1},
		{0x1d78a, 0x1d7a8, 1},
		{0x1d7aa, 0x1d7c2, 1},
		{0x1d7c4, 0x1d7cb, 1},
		{0x1d7ce, 0x1d7ff, 1},
		{0x1da00, 0x1da36, 1},
		{0x1da3b, 0x1da6c, 1},
		{0x1da75, 0x1da75, 1},
		{0x1da84, 0x1da84, 1},
		{0x1da9b, 0x1da9f, 1},
		{0x1daa1, 0x1daaf, 1},
		{0x1e000, 0x1e006, 1},
		{0x1e008, 0x1e018, 1},
		{0x1e01b, 0x1e021, 1},
		{0x1e023, 0x1e024, 1},
		{0x1e026, 0x1e02a, 1},
		{0x1e100, 0x1e12c, 1},
		{0x1e130, 0x1e13d, 1},
		{0x1e140, 0x1e149, 1},
		{0x1e14e, 0x1e14e, 1},
		{0x1e2c0, 0x1e2f9, 1},
		{0x1e800, 0x1e8c4, 1},
		{0x1e8d0, 0x1e8d6, 1},
		{0x1e900, 0x1e94b, 1},
		{0x1e950, 0x1e959, 1},
		{0x1ee00, 0x1ee03, 1},
		{0x1ee05, 0x1ee1f, 1},
		{0x1ee21, 0x1ee22, 1},
		{0x1ee24, 0x1ee24, 1},
		{0x1ee27, 0x1ee27, 1},
		{0x1ee29, 0x1ee32, 1},
		{0x1ee34, 0x1ee37, 1},
		{0x1ee39, 0x1ee39, 1},
		{0x1ee3b, 0x1ee3b, 1},
		{0x1ee42, 0x1ee42, 1},
		{0x1ee47, 0x1ee47, 1},
		{0x1ee49, 0x1ee49, 1},
		{0x1ee4b, 0x1ee4b, 1},
		{0x1ee4d, 0x1ee4f, 1},
		{0x1ee51, 0x1ee52, 1},
		{0x1ee54, 0x1ee54, 1},
		{0x1ee57, 0x1ee57, 1},
		{0x1ee59, 0x1ee59, 1},
		{0x1ee5b, 0x1ee5b, 1},
		{0x1ee5d, 0x1ee5d, 1},
		{0x1ee5f, 0x1ee5f, 1},
		{0x1ee61, 0x1ee62, 1},
		{0x1ee64, 0x1ee64, 1},
		{0x1ee67, 0x1ee6a, 1},
		{0x1ee6c, 0x1ee72, 1},
		{0x1ee74, 0x1ee77, 1},
		{0x1ee79, 0x1ee7c, 1},
		{0x1ee7e, 0x1ee7e, 1},
		{0x1ee80, 0x1ee89, 1},
		{0x1ee8b, 0x1ee9b, 1},
		{0x1eea1, 0x1eea3, 1},
		{0x1eea5, 0x1eea9, 1},
		{0x1eeab, 0x1eebb, 1},
		{0x1fbf0, 0x1fbf9, 1},
		{0x20000, 0x2a6dd, 1},
		{0x2a700, 0x2b734, 1},
		{0x2b740, 0x2b81d, 1},
		{0x2b820, 0x2cea1, 1},
		{0x2ceb0, 0x2ebe0, 1},
		{0x2f800, 0x2fa1d, 1},
		{0x30000, 0x3134a, 1},
		{0xe0100, 0xe01ef, 1},
	},
}
