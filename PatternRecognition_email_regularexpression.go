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
	fmt.Println("Email - Model of Computation")
	fmt.Println("--- Reading file ---")
	data, err := ioutil.ReadFile("emails.txt")
	checkerror(err)
	// convert data from bytes type to strings
	emails_data_str := string(data)
	//fmt.Println(email_data_str)
	emails := s.Split(emails_data_str, "From r ")
	emails = emails[1:]
	//fmt.Println("Split : ", email)
	fmt.Println(len(emails))

	for _, email := range emails {
		sender := regexp.MustCompile("From:.*")
		sender_data := sender.FindString(email)
		// fmt.Println("Email-Sender ", index, " : ", sender_data)
		// From: "MR. JAMES NGOLA." <james_ngola2002@maktoob.com>

		s_email := regexp.MustCompile("<.*@[a-zA-Z]+.[a-zA-Z]+")
		email_sender := s_email.FindString(sender_data)

		s_name := regexp.MustCompile("\".*\"")
		name_sender := s_name.FindString(sender_data)

		fmt.Println("Email-Sender : ", email_sender[1:])
		fmt.Println("Name-Sender : ", s.Trim(name_sender, "\""))

		recipient := regexp.MustCompile("To:.*")
		recipient_data := recipient.FindString(email)
		email_rep := regexp.MustCompile(" .*@.*$")
		recipient_email := email_rep.FindString(recipient_data)
		fmt.Println("Name-Recipient : ", s.Replace(recipient_email, " ", "", -1))

		date := regexp.MustCompile("Date:.*")
		date_data := date.FindString(email)
		date_re := regexp.MustCompile("[0-9]+.[a-zA-Z]+.[0-9]+.")
		date_my := date_re.FindString(date_data)
		fmt.Println("Date send : ", date_my)

		subject := regexp.MustCompile("Subject: .*")
		subject_data := subject.FindString(email)
		subject_title := s.Replace(subject_data, "Subject: ", "", -1)
		fmt.Println("Subject : ", subject_title)

		// email_rep := regexp.MustCompile(" .*@.*$")
		// recipient_email := email_rep.FindString(recipient_data)
		// fmt.Println("Name-Recipient : ", s.Replace(recipient_email, " ", "", -1))

		// data := map[string]string{
		// 	"Email_sender": email_sender[1:],
		// }
		// fmt.Println(data)
	}

}
