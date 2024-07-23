package main

import (
	"fmt"

	str "github.com/barnabassol/stringmod/quotes"
	qts "github.com/barnabassol/stringmod/strings"
)

func main() {
	o, e := str.CountOddEven("12345")
	fmt.Println(o, e)
	fmt.Println(qts.GetEmoji("turtle"))
}
