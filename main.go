package main

// {% lorem %} will output the common 'lorem ipsum' paragraph.
// {% lorem 3 p %} will output the common lorem 'ipsum paragraph' and
//     two random paragraphs each wrapped in HTML <p> tags.
// {% lorem 2 w random %} will output two random Latin words.

import (
	"fmt"

	"oldcode.org/lorem/lib"
)

func main() {
	fmt.Printf("Main ...\n")

	//fmt.Println(lib.Sentence())
	//fmt.Println(lib.Paragraph())

	//ps := lib.ParagraphsC(5)
	//for _, p := range ps {
	//	fmt.Println(p)
	//	fmt.Println()
	//}

	fmt.Println(lib.Words(4, false))
	fmt.Println()
	fmt.Println(lib.Words(5, true))
}
