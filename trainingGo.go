package main

import (
	"fmt"
	s "strings"
)

//scope global variable
var globalVariable int = 5000
var p = fmt.Println

func main() {
	stringF()
	x := 15
	fmt.Println("value x is ", x)
	fmt.Println("Address x is ", &x)
	var p *int
	p = &x //pointer
	fmt.Println("Pointer p is %x", p)
	*p = 20 //เปลี่ยนค่าให้เป็น 20 ด้วย pointer
	fmt.Println("value x is ", x)

	structure()
	closure()
	mp()
	slicing()
	array()
	summation(1, 2, 3) //data
	fmt.Println(factorial(5))
}

func stringF() {
	p("Contains : ", s.Contains("test", "es"))      //ตรวจสอบว่ามี es ภายใน test หรือไม่
	p("Count : ", s.Count("test", "t"))             //นับว่ามีตัว t ภายใน test ทั้งหมดกี่ตัว
	p("HasPrefix : ", s.HasPrefix("test", "te"))    //ตรวจสอบว่า test ขึ้นต้นด้วย te หรือไม่
	p("HasSuffix : ", s.HasSuffix("test", "st"))    //ตรวจสอบว่า test ลงท้ายด้วย st หรือไม่
	p("Index : ", s.Index("test", "e"))             //ดึงค่า index ของตัวอักษร e ในคำว่า test
	p("Join : ", s.Join([]string{"a", "b"}, "-"))   // นำสตริง a และ b มาต่อกัน ซึ่งมี - มาคั่นกลาง
	p("Repeat : ", s.Repeat("a", 5))                //กำหนดให้ตัวอักษร a ซ้ำกันจำนวน 5 ตัวอักษร จะได้ aaaaa
	p("Replace : ", s.Replace("foo", "o", "0", -2)) //กำหนด 1 จะเปลี่ยนแค่ 1 ตัวเริ่มจากซ้ายไปขวา แต่ถ้ากำหนด -n จะเปลี่ยนทุกตัว
	p("Split : ", s.Split("t-e-s-t", "-"))          //แยกตัวอักษร หรือ คำ ต้องสิ่งที่กำหนด ซึ่งอาจจะเป็นตัวอักษร ช่องว่าง หรือ สัญลักษณ์
	p("ToUpper : ", s.ToUpper("test"))
	p("ToLower : ", s.ToLower("TEST"))
	p("Len : ", len("TEST"))
}

type Books struct {
	title    string
	author   string
	subtitle string
	prices   float64
}

func structure() {
	//like database use to store data
	var Book1 Books
	Book1.title = "Go Programming"
	Book1.author = "Kor CS KMITL"
	Book1.subtitle = "Model of Computation"
	Book1.prices = 600

	fmt.Println(Book1.title)

	Golang := Books{title: "Go Programming", author: "Kor CS KMITL", prices: 599.99}
	fmt.Println(Golang)
}
func closure() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
}

func slicing() { // don't fix len of array
	// x := make([]int, 5)
	// fmt.Println(x)

	//second for
	slice := []float64{1, 2, 3}
	fmt.Println(slice)
	slice2 := append(slice, 4, 5)
	fmt.Println(slice2)

	//copy
	sliceA := []int{1, 2, 3}
	sliceB := make([]int, 2)
	copy(sliceB, sliceA)
	fmt.Println(sliceB)
}

func mp() {
	// like a dictionary of python's type : key , value
	mp := make(map[string]string)
	mp["TH"] = "THAILAND"
	mp["SG"] = "SINGAPORE"
	mp["JP"] = "JAPAN"

	fmt.Println(mp["JP"])

	smp := map[string]string{
		"TH": "THAILAND",
		"SG": "SINGAPORE",
		"JP": "JAPAN",
	}
	fmt.Println(smp)
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
