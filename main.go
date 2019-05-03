package main

// lorem -- will output the common 'lorem ipsum' paragraph.
// lorem 3 p -- will output the common lorem 'ipsum paragraph' and
//     two random paragraphs
// lorem 2 w random -- will output two random Latin words.

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/krsanky/lorem/lib"
)

func usage() {
	fmt.Println("lorem [num] [p|w] [random]")
}

func main() {
	//args: 0 2 or 3
	//1st is number
	//2nd is p or w
	//3rd is random

	args := os.Args[1:]
	var (
		num    int64
		porw   string
		random bool
		err    error
	)

	switch l := len(args); {
	case l == 0:
		fmt.Println(lib.COMMON_P)
		os.Exit(0)
	case l == 2, l == 3:
		num, err = strconv.ParseInt(args[0], 10, 0)
		if err != nil {
			usage()
			os.Exit(1)
		}

		porw = args[1]

		if l == 3 {
			switch args[2] {
			case "random":
				random = true
			default:
				usage()
				os.Exit(1)
			}
		}
		//fmt.Printf("num:%d porw:%s random:%t\n", num, porw, random)
	default:
		usage()
		os.Exit(1)
	}

	switch porw {
	case "p":
		ps := lib.Paragraphs(int(num), !random)
		fmt.Println(strings.Join(ps, "\n\n"))
	case "w":
		ws := lib.Words(int(num), !random)
		fmt.Println(strings.Join(ws, " "))
	default:
		usage()
		os.Exit(1)
	}

}
