package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	s "strings"
)

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

var sender_email_array [3976]string

func main() {
	data, err := ioutil.ReadFile("fradulent_emails_dataset.txt")
	checkerror(err)
	emails_data_str := string(data)
	emails := s.Split(emails_data_str, "From r ")
	emails = emails[1:]

	fmt.Println("Total : ", len(emails))

	for index, email := range emails {

		// fmt.Println("----- Email No. ", index+1, " ------")
		sender := regexp.MustCompile("From:.*")
		sender_data := sender.FindString(email)

		fmt.Println(sender_data)

		s_email := regexp.MustCompile("<.*@[a-zA-Z]+.[a-zA-Z]+")
		email_sender := s_email.FindString(sender_data)
		emails_sender := s.Trim(email_sender, "<")
		fmt.Println(emails_sender)
		if emails_sender != "" {
			sender_email_array[index] = emails_sender
		} else {
			sender_email_array[index] = "None"
		}
		s_name := regexp.MustCompile("From:[\t\n\f\r ]+.*[\t\n\f\r ]+")
		name_sender := s_name.FindString(sender_data)
		name_sender = s.Replace(name_sender, "From: ", "", -1)
		// if name_sender[0] == '"' {
		// 	name_sender = s.Trim(name_sender, "\"*\"")
		// }
		fmt.Println(name_sender)

	}

	//fmt.Println(sender_email_array)
	fmt.Println(len(sender_email_array))

}
