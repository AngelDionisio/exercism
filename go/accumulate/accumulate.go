package accumulate

// Accumulate applies the 'operation' function to each item in the slice list
func Accumulate(list []string, operation func(string) string) []string {
	var result = make([]string, len(list))

	for i, v := range list {
		result[i] = operation(v)
	}

	return result
}

// Accumulate2 is a less performant operation, if we know the size of the result slice
// append will create a new array, doubling the size and copying all the contents when
// reaching the capacity of the bucket slice. We can avoid this by instantiating an array
// like in the function above
func Accumulate2(list []string, operation func(string) string) []string {
	if list == nil || len(list) == 0 {
		return nil
	}

	var listResult []string
	for _, str := range list {
		listResult = append(listResult, operation(str))
	}

	return listResult
}
