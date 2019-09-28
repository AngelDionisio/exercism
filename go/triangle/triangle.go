package triangle

import (
	"math"
)

// Kind denotes the type of a triagle
type Kind string

// Def of triangle types
const (
	// NaT: not a triagle
	NaT = "Not a triangle"
	// Equ: equilateral triangle, all three sideas have same length
	Equ = "Equilateral"
	// Iso: has at least two sides with same length
	Iso = "Isosceles"
	// Sca: all sides have different length
	Sca = "Scalene"
	// Dge: the sume of two sides equal to that of the third
	Dge = "Degenerate"
)

// KindFromSides returns whether inputs can define a triangle,
// if so, it returns it's type
func KindFromSides(a, b, c float64) Kind {

	if !isTriangle(a, b, c) {
		return NaT
	} else if isEquilateral(a, b, c) {
		return Equ
	} else if isIsosceles(a, b, c) {
		return Iso
	} else if isScalene(a, b, c) {
		return Sca
	}

	return NaT
}

// for a shape to be a triangle at all, all sides have to be of length > 0,
// and the sum of the lengths of any two sides must be greater than
// or equal to the length of the third side
func isTriangle(a, b, c float64) bool {
	return isValidSide(a) && isValidSide(b) && isValidSide(c) &&
		(a <= b+c) && (b <= a+c) && (c <= b+a)
}

func isValidSide(side float64) bool {
	return side > 0 && side != math.Inf(1)
}

// An equilateral triangle has all three sides the same length
func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

// An isosceles triangle has at least two sides the same length
func isIsosceles(a, b, c float64) bool {
	return (a == b || a == c || b == c) && !isEquilateral(a, b, c)
}

// IsScalene triangle has all sides of different lengths.
func isScalene(a, b, c float64) bool {
	return a != b && a != c && b != c
}
