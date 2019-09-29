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

** FOR LOOPS and Strings **
When for ranging on a utf-8 value (string), the index will start at each rune,
 measured in bytes.

 e.g.
 const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

 output:
	U+65E5 '日' starts at byte position 0
	U+672C '本' starts at byte position 3
	U+8A9E '語' starts at byte position 6

 	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("%#U starts at byte position %d\n", nihongo[i], i)
	}

 output:
	U+0097 starts at byte position 1
	U+00A5 '¥' starts at byte position 2
	U+00E6 'æ' starts at byte position 3
	U+009C starts at byte position 4
	U+00AC '¬' starts at byte position 5
	U+00E8 'è' starts at byte position 6
	U+00AA 'ª' starts at byte position 7
	U+009E starts at byte position 8

Conclusion:
To answer the question posed at the beginning:
Strings are built from bytes so indexing them yields bytes, not characters.
A string might not even hold characters. In fact, the definition of "character" is ambiguous
and it would be a mistake to try to resolve the ambiguity by defining that strings are made of characters.
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
