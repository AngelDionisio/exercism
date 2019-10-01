package strand

// ToRNA transforms a DNA strand to RNA
func ToRNA(dna string) string {
	rnaMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	resultStrand := []rune{}
	// for-ranging over a string, each loop starts
	// at the start of the next rune
	for _, v := range dna {
		resultStrand = append(resultStrand, rnaMap[v])
	}

	return string(resultStrand)

}
