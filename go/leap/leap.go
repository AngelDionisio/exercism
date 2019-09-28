package leap

// IsLeapYear reports if a given year is a leap year
// on every year that is divisible by 4,
// except every year that is evenly divisible by 100, unless the year is also evenly divisible by 400
func IsLeapYear(yr int) bool {
	if yr%400 == 0 || (yr%4 == 0 && yr%100 != 0) {
		return true
	}

	return false
}
