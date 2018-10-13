package main

import "fmt"

//scope global variable
var globalVariable int = 5000

func main() {
	array()
	summation(1, 2, 3) //data
	fmt.Println(factorial(5))
}

func slicing() { // don't fix len of array

}
func array() {
	var x [5]int
	x[3] = 200
	fmt.Println(x)

	y := [10]float64{2.5, 5.3, 7, 3.2, 9.0}
	var totalY float64
	for _, n := range y {
		totalY += n
	}
	fmt.Println(totalY / float64(len(y)))
}

func summation(nums ...int) { // Variadic Function : unknow lenght of list
	var total int
	for _, n := range nums {
		total += n
	}
	fmt.Println(total)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func training() {
	fmt.Println("Kor Computer Science KMITL")

	var number int = 200
	number2 := 500 // automatic assign type of variable
	var boolean bool = true
	age1, age2 := 50, 80
	fmt.Println("The number in this project is : ", number)
	fmt.Println("The number which automatic assign is : ", number2)
	fmt.Println(boolean)
	fmt.Println(age1, age2)

	text1 := "tanakorn"
	text2 := "puraram"

	// concaticaiton
	textConcat := text1 + text2
	fmt.Println(textConcat)
	fmt.Println(textConcat[1])  //ASCII
	fmt.Println(textConcat[0:]) //TEXT

	isEqual := 5 < 3
	fmt.Println(isEqual)

	fmt.Println(globalVariable)

	fmt.Print("Input your ID : ")
	var iduser string
	fmt.Scanf("%s", &iduser)
	fmt.Println("ID user : ", iduser)
	if iduser == "61605045" {
		fmt.Println("You are KMITL Master Degree Student")
	}
}
