package raindrops

import (
	"strconv"
)

func allFactors(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// Convert returns "Pling", "Plang", "Plong" if
// the input is divisible by 3, 5 or 7 respectively
func Convert(n int) string {
	factors := allFactors(n)
	str := ""

	if !contains(factors, 3) && !contains(factors, 5) && !contains(factors, 7) {
		return strconv.Itoa(n)
	}

	for _, v := range factors {
		if v == 3 {
			str = str + "Pling"
		}
		if v == 5 {
			str = str + "Plang"
		}
		if v == 7 {
			str = str + "Plong"
		}
	}

	return str
}
