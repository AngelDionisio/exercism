package proverb

import "fmt"

// Proverb returns a proverbial rhyme based on the input list
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	var result []string

	for i, word := range rhyme {
		// when index == last position in list, add closing statement
		// covers case when there's only one item in list
		if i == len(rhyme)-1 {
			result = append(result, closingStatement(rhyme[0]))
		} else {
			row := rhymeGenerator(word, rhyme[i+1])
			result = append(result, row)
		}
	}
	return result
}

func rhymeGenerator(verb1 string, verb2 string) string {
	return fmt.Sprintf("For want of a %s the %s was lost.", verb1, verb2)
}

func closingStatement(s string) string {
	return fmt.Sprintf("And all for the want of a %s.", s)
}
