package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	//[[cat c] [bat b] [sat s] [mat m]]
	re, err := regexp.Compile(`(.)at`)
	if err != nil {
		log.Fatal(err)
		return
	}
	res := re.FindAllStringSubmatch("The cat bat sat on the mat.", -1)
	fmt.Printf("%v", res)

	s := "Mr. Nattawut Ruangvivattanaroj"
	re1, err := regexp.Compile(`(Mr)(s)?\. (\w+) (\w+)`)
	result := re1.FindStringSubmatch(s)

	for k, v := range result {
		fmt.Printf("%d. %s\n", k, v)
	}
	// Prints
	// 1. Mr
	// 2.
	// 3. Nattawut
	// 4. Ruangvivattanaroj
}
