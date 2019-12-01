package acronym

import (
	"strings"
)

// Abbreviate returns the first letter of each word in caps.
// first use replacer to clean text of special characters
// use Title to uppercase the first letter of each word encountered after one or more spaces
// use Fields to separate string into a []string around one more consecutive white spaces
func Abbreviate(s string) string {
	r := strings.NewReplacer("-", " ", "_", "")
	cleandedText := r.Replace(s)

	titled := strings.Title(cleandedText)
	sliceOfStrings := strings.Fields(titled)

	// strings builder is an efficient way to build strings
	// use it instead of concatenating
	var result strings.Builder
	for _, str := range sliceOfStrings {
		result.WriteString(string(str[0]))
	}
	return result.String()
}
