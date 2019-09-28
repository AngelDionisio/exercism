package hamming

import "errors"

/**
In Go a string is nothing more than a slice of bytes.
A rune is the accurate representation of a string, as each character
 is stored as an int32 value.
A slice of bytes of a string may not be equal to a slice of rune of a string
Words like ñ require more characters as a slice of byte, e.g.
The character ñ
[]rune: [241]
[]byte: [195 177]

It is important to note that a string holds arbitrary bytes. It is not required
 to hold Unicode text, UTF-8 text or any other predefined format.
 A string is exactly equivalent to a slice of bytes. This allows for printing
 it's contents into different formats, a plain string, ASCII, hexadecimal etc.

When we store a character value in a string, we store
 const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
 prints out: ��=� ⌘

 To find out what the string really holds, we need to take it apart and examine its pieces
The most obvious one is to loop over its contents.
However indexing a string, YIELDS ITS BYTES not it's characters.

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
for i := 0; i < len(sample); i++ {
	fmt.Printf("%x ", sample[i]) // bd b2 3d bc 20 e2 8c 98
}

** Runes
	a rune / code point is the atomic unit of information.
	Text is a sequence of code points. Each code point is an int32
	which is given meaning by the Unicode standard

** Character
	is an overloaded term that can mean many things.

https://play.golang.org/p/w7HQAx_rsgx
*/

// Distance calculates hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	// a slice of runes is like a slice of bytes, but instead of bytes
	// it holds a character at each index.
	ra := []rune(a)
	rb := []rune(b)

	if len(ra) != len(rb) {
		return 0, errors.New("strand resprestantions must be of equal length")
	}

	var distance int
	for i := range ra {
		if ra[i] != rb[i] {
			distance++
		}
	}

	return distance, nil
}
