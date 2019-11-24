package accumulate

// Accumulate applies the 'operation' function to each item in the slice list
func Accumulate(list []string, operation func(string) string) []string {
	if list == nil || len(list) == 0 {
		return nil
	}

	var listResult []string
	for _, str := range list {
		listResult = append(listResult, operation(str))
	}

	return listResult
}
