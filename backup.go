package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// read file (python : fh = open(r"test_emails.txt", "r").read() )
	data, err := ioutil.ReadFile("emails.txt")
	checkerror(err)
	fmt.Println(data) // data type is  'bytes'

	// convert data from bytes type to strings
	str := string(data)
	// fmt.Println(str)

	// python import re
	// for line in re.findall("From:.*", fh):
	// print(line)

	from, err := regexp.Compile(`(From:)`)
	//r := re.FindAllString("{From}", -1)
	r := from.FindAllString(str)
	fmt.Println(r)

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

	s2 := "  Is golang good for real world web apps?  "
	// <Is golang good for real world web apps?>
	fmt.Printf("<%v>", strings.TrimSpace(s2))

	// Prints
	// 1. Mr
	// 2.
	// 3. Nattawut
	// 4. Ruangvivattanaroj

	// fmt.Println("sender_email_array : ", len(sender_email_array))
	// fmt.Println("sender_name_array : ", len(sender_name_array))
	// fmt.Println("username_array : ", len(username_array))
	// fmt.Println("domainname_array : ", len(domainname_array))
	// fmt.Println("domaingroup_array : ", len(domaingroup_array))
	// fmt.Println("recipient_email_array : ", len(recipient_email_array))
	// fmt.Println("reply_email_array : ", len(reply_email_array))
	// fmt.Println("datesend_array : ", len(datesend_array))
	// fmt.Println("date_array : ", len(date_array))
	// fmt.Println("month_array : ", len(month_array))
	// fmt.Println("year_array : ", len(year_array))
	// fmt.Println("day_array : ", len(day_array))
	// fmt.Println("time_array : ", len(time_array))
	// fmt.Println("subject_array : ", len(subject_array))
	// fmt.Println("content_data : ", len(contentemail_array))

}
