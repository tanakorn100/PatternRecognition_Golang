package main

import (
	"fmt"
	"io/ioutil"
)

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Email - Model of Computation")

	fmt.Println("--- Reading file ---")
	data, err := ioutil.ReadFile("emails.txt")
	checkerror(err)
	// convert data from bytes type to strings
	str := string(data)
	fmt.Println(str)
}
