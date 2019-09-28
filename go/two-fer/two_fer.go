// Package twofer provides a simple function that prints one of two
// strings depending on whether a name param was passed to the
// ShareWith function.
package twofer

import "fmt"

// ShareWith returns the sentence "one for 'name'", one for me
// if no name is provided, it returns 'you' instead of 'name'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
