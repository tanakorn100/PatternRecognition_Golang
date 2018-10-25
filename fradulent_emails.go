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
var datesend_array [3976]string
var date_array [3976]string
var month_array [3976]string
var year_array [3976]string
var day_array [3976]string
var time_array [3976]string
var subject_array [3976]string

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

		// recipient_data email
		recipient := regexp.MustCompile("\nTo:.*")
		recipient_data := recipient.FindString(email)
		recipient_data = s.Replace(recipient_data, "\nTo: ", "", -1)
		//fmt.Println("", index+1, " : ", recipient_data)
		if recipient_data != "" {
			recipient_email_array[index] = recipient_data
		} else {
			recipient_email_array[index] = "None"
		}

		// reply email
		reply := regexp.MustCompile("\nReply-To:.*")
		reply_data := reply.FindString(email)
		reply_data = s.Replace(reply_data, "\nReply-To: ", "", -1)
		// fmt.Println("", index+1, " : ", reply_data)
		if reply_data != "" {
			reply_email_array[index] = reply_data
		} else {
			reply_email_array[index] = "None"
		}

		// date
		date := regexp.MustCompile("Date:.*")
		date_data := date.FindString(email)
		//fmt.Println(date_data)
		date_re := regexp.MustCompile("[0-9]+.[a-zA-Z]+.[0-9]+.")
		date_my := date_re.FindString(date_data)
		// fmt.Println("", index+1, " : ", date_my)
		date_d := regexp.MustCompile("[0-9]{2}")
		dateD := date_d.FindString(date_my)
		// fmt.Println(dateD)
		date_m := regexp.MustCompile("[a-zA-Z]{3}")
		dateM := date_m.FindString(date_my)
		// fmt.Println(dateM)
		date_y := regexp.MustCompile("[0-9]{4}")
		dateY := date_y.FindString(date_my)
		// fmt.Println(dateY)
		date_week := regexp.MustCompile("[\t\n\f\r ][a-zA-Z]{3}")
		day := date_week.FindString(date_data)
		time_date := regexp.MustCompile("[0-9]{2}:[0-9]{2}:[0-9]{2}[\t\n\f\r ].*")
		time := time_date.FindString(date_data)
		//fmt.Println(day, "\t", time)
		if date_data != "" {
			datesend_array[index] = date_my
			date_array[index] = dateD
			month_array[index] = dateM
			year_array[index] = dateY
			day_array[index] = day
			time_array[index] = time
		} else {
			datesend_array[index] = "None"
			date_array[index] = "None"
			month_array[index] = "None"
			year_array[index] = "None"
			day_array[index] = "None"
			time_array[index] = "None"
		}

		subject := regexp.MustCompile("\nSubject:[\t\n\f\r ]*.*")
		subject_data := subject.FindString(email)
		subject_data = s.Replace(subject_data, "\nSubject: ", "", -1)
		// fmt.Println("", index+1, " : ", subject_data)
		if subject_data != "" {
			subject_array[index] = subject_data
		} else {
			subject_array[index] = "None"
		}

		content := regexp.MustCompile("Status:[\t\n\f\r ]+[a-zA-Z0-9]+(\n|[0-9A-Za-z_]+|.*|[\t\n\f\r ])+")
		content_data := content.FindString(email)
		content_data = s.Replace(content_data, "\nStatus: ", "", -1)
		fmt.Println("", index+1, " : ", content_data)
	}

	fmt.Println("sender_email_array : ", len(sender_email_array))
	fmt.Println("sender_name_array : ", len(sender_name_array))
	fmt.Println("recipient_email_array : ", len(recipient_email_array))
	fmt.Println("reply_email_array : ", len(reply_email_array))
	fmt.Println("datesend_array : ", len(datesend_array))
	fmt.Println("date_array : ", len(date_array))
	fmt.Println("month_array : ", len(month_array))
	fmt.Println("year_array : ", len(year_array))
	fmt.Println("day_array : ", len(day_array))
	fmt.Println("time_array : ", len(time_array))
	fmt.Println("subject_array : ", len(subject_array))

}
