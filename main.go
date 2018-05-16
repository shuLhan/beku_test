package main

import (
	"fmt"

	"github.com/shuLhan/share/lib/text"
)

func main() {
	fmt.Println(text.StringJSONEscape("v0.1.0"))
	fmt.Println("Add feature A.")
	fmt.Println("Add feature B.")
}
