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
var sender_name_array [3976]string
var username_array [3976]string
var domainname_array [3976]string
var recipient_email_array [3976]string
var reply_email_array [3976]string

// var datesend_array [3976]string
// var subject_array [3976]string
// var contentemail_array [3976]string

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

		s_email := regexp.MustCompile("<.*[^\t\n\f\r ]@[a-zA-Z]+.[a-zA-Z]+")
		email_sender := s_email.FindString(sender_data)
		emails_sender := s.Trim(email_sender, "<")

		if emails_sender != "" {
			sender_email_array[index] = emails_sender
			email_account := s.Split(emails_sender, "@")
			username_array[index] = email_account[0]
			domainname_array[index] = email_account[1]
		} else {
			sender_email_array[index] = "None"
			username_array[index] = "None"
			domainname_array[index] = "None"
		}

		s_name := regexp.MustCompile("From:[\t\n\f\r ]+.*[\t\n\f\r ]+")
		name_sender := s_name.FindString(sender_data)
		name_sender = s.Replace(name_sender, "From: ", "", -1)

		// remore \" text \"

		// if name_sender[0] == '"' {
		// 	name_sender = s.Trim(name_sender, "\"*\"")
		// }

		// append name_sender to array list
		if name_sender != "" {
			sender_name_array[index] = name_sender
		} else {
			sender_name_array[index] = "None"
		}

		recipient := regexp.MustCompile("\nTo:.*")
		recipient_data := recipient.FindString(email)
		recipient_data = s.Replace(recipient_data, "\nTo: ", "", -1)
		//fmt.Println("", index+1, " : ", recipient_data)
		if recipient_data != "" {
			recipient_email_array[index] = recipient_data
		} else {
			recipient_email_array[index] = "None"
		}

		reply := regexp.MustCompile("\nReply-To:.*")
		reply_data := reply.FindString(email)
		reply_data = s.Replace(reply_data, "\nReply-To: ", "", -1)
		fmt.Println("", index+1, " : ", reply_data)
		if reply_data != "" {
			reply_email_array[index] = reply_data
		} else {
			reply_email_array[index] = "None"
		}

	}

	//fmt.Println(sender_email_array)
	fmt.Println(len(sender_email_array))
	//fmt.Println(sender_name_array)
	fmt.Println(len(sender_name_array))
	//fmt.Println(recipient_email_array)
	fmt.Println(len(recipient_email_array))

}
