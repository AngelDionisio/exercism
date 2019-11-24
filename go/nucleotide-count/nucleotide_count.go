package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	histogram := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	// iterate over each rune in string
	for _, r := range d {
		// check if rune exists in histogram
		_, isFound := histogram[r]
		if isFound {
			histogram[r]++
		} else {
			return nil, errors.New("Invalid Nucleotide")
		}
	}
	return histogram, nil
}
