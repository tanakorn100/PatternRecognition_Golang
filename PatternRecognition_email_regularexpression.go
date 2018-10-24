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

var sender_email_array []string
var sender_name_array []string
var username_array []string
var domainname_array []string
var recipient_email_array []string
var datesend_array []string
var subject_array []string
var contentemail_array []string

func main() {
	data, err := ioutil.ReadFile("emails.txt")
	checkerror(err)
	emails_data_str := string(data)
	emails := s.Split(emails_data_str, "From r ")
	emails = emails[1:]

	fmt.Println("Total : ", len(emails))

	for _, email := range emails {

		// fmt.Println("----- Email No. ", index+1, " ------")
		sender := regexp.MustCompile("From:.*")
		sender_data := sender.FindString(email)

		s_email := regexp.MustCompile("<.*@[a-zA-Z]+.[a-zA-Z]+")
		email_sender := s_email.FindString(sender_data)
		emails_sender := s.Trim(email_sender, "<")

		s_name := regexp.MustCompile("\".*\"")
		name_sender := s_name.FindString(sender_data)
		names_sender := s.Trim(name_sender, "\"")

		// fmt.Println("Sender Email : ", emails_sender)
		// fmt.Println("Sender Name : ", names_sender)

		sender_email_array := append(sender_email_array, emails_sender)
		sender_name_array := append(sender_name_array, names_sender)

		email_account := s.Split(emails_sender, "@")
		//fmt.Println("Username : ", email_account[0], "\tDomain : ", email_account[1])
		username_array := append(username_array, email_account[0])
		domainname_array := append(domainname_array, email_account[1])

		recipient := regexp.MustCompile("To:.*")
		recipient_data := recipient.FindString(email)
		email_rep := regexp.MustCompile(" .*@.*$")
		recipient_email := email_rep.FindString(recipient_data)

		recipient_email_array := append(recipient_email_array, recipient_email)

		//fmt.Println("Name-Recipient : ", s.Replace(recipient_email, " ", "", -1))
		//fmt.Println("Recipient Email :", recipient_email)

		date := regexp.MustCompile("Date:.*")
		date_data := date.FindString(email)
		date_re := regexp.MustCompile("[0-9]+.[a-zA-Z]+.[0-9]+.")
		date_my := date_re.FindString(date_data)

		datesend_array := append(datesend_array, date_my)
		//fmt.Println("Date send : ", date_my)

		subject := regexp.MustCompile("Subject:[\t\n\f\r ]*.*")
		subject_data := subject.FindString(email)

		subject_array := append(subject_array, subject_data)

		//subject_title := s.Replace(subject_data, "Subject: ", "", -1)
		//fmt.Println(subject_data)

		content := regexp.MustCompile("Status:[\t\n\f\r ]+[a-zA-Z0-9]+(\n|[0-9A-Za-z_]+|.*|[\t\n\f\r ])+")
		content_data := content.FindString(email)

		contentemail_array := append(contentemail_array, content_data)
		//fmt.Println("Email-Content : ", content_data)
	}

	for index, _ := range sender_email_array {
		email := map[string]string{
			"sender_email":    sender_email_array[index],
			"sender_name":     sender_name_array[index],
			"username":        username_array[index],
			"domainname":      domainname_array[index],
			"recipient_email": recipient_email_array[index],
			"datesent":        datesend_array[index],
			"subject":         subject_array[index],
			"contentemail":    contentemail_array[index],
		}

		fmt.Println(email)

		// email_datas[index] = email
	}
}
