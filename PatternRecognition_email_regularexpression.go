package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, err := regexp.Compile(`Nattawut`)

	if err != nil {
		fmt.Printf("Error : Don't have Nattawut")
		return
	}

	if r.MatchString("Hello Regular Expression. from Nattawut") == true {
		fmt.Printf("Match ")
	} else {
		fmt.Printf("No match ")
	}

	r2, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Printf("Error : Type not INT")
		return
	}

	// Will print 'true':
	fmt.Printf("%v", r2.MatchString("Seven times seven is 49."))
	// Will print 'false':
	fmt.Printf("%v", r2.MatchString("Seven times seven is forty-nine."))

}
