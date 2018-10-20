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

func main() {
	data, err := ioutil.ReadFile("emails.txt")
	checkerror(err)
	// convert data from bytes type to strings
	emails_data_str := string(data)
	// split email dataset to array list
	emails := s.Split(emails_data_str, "From r ")
	// remove index 0 is blank
	emails = emails[1:]
	fmt.Println("Total : ", len(emails))

	for index, email := range emails {

		fmt.Println("----- Email No. ", index+1, " ------")
		sender := regexp.MustCompile("From:.*")
		sender_data := sender.FindString(email)

		s_email := regexp.MustCompile("<.*@[a-zA-Z]+.[a-zA-Z]+")
		email_sender := s_email.FindString(sender_data)

		s_name := regexp.MustCompile("\".*\"")
		name_sender := s_name.FindString(sender_data)

		fmt.Println("Email-Sender : ", email_sender[1:])
		fmt.Println("Name-Sender : ", s.Trim(name_sender, "\""))
		email_account := s.Split(email_sender[1:], "@")
		fmt.Println("Account : ", email_account[0], "\tDomain : ", email_account[1])

		recipient := regexp.MustCompile("To:.*")
		recipient_data := recipient.FindString(email)
		email_rep := regexp.MustCompile(" .*@.*$")
		recipient_email := email_rep.FindString(recipient_data)
		//fmt.Println("Name-Recipient : ", s.Replace(recipient_email, " ", "", -1))
		fmt.Println("Name-Recipient :", recipient_email)

		date := regexp.MustCompile("Date:.*")
		date_data := date.FindString(email)
		date_re := regexp.MustCompile("[0-9]+.[a-zA-Z]+.[0-9]+.")
		date_my := date_re.FindString(date_data)
		fmt.Println("Date send : ", date_my)

		subject := regexp.MustCompile("Subject:[\t\n\f\r ]*.*")
		subject_data := subject.FindString(email)
		//subject_title := s.Replace(subject_data, "Subject: ", "", -1)
		fmt.Println(subject_data)

		content := regexp.MustCompile("Status:[\t\n\f\r ]+[a-zA-Z]+(\n|[0-9A-Za-z_]+|.*|[\t\n\f\r ])+")
		content_data := content.FindString(email)
		fmt.Println("Body : ", content_data)

		// data := map[string]string{
		// 	"Email_sender": email_sender[1:],
		// }
		// fmt.Println(data)
	}

}
