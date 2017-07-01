package main

// {% lorem %} will output the common 'lorem ipsum' paragraph.
// {% lorem 3 p %} will output the common lorem 'ipsum paragraph' and
//     two random paragraphs each wrapped in HTML <p> tags.
// {% lorem 2 w random %} will output two random Latin words.

import (
	"fmt"

	"oldcode.org/web/lorem"
)

func main() {
	fmt.Printf("Main ...\n")
	//fmt.Println(lorem.Sentence())
	//fmt.Println(lorem.Paragraph())
	ps := lorem.ParagraphsC(5)
	for _, p := range ps {
		fmt.Println(p)
		fmt.Println()
	}

}
