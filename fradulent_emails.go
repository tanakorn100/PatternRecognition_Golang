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
var domaingroup_array [3976]string
var recipient_email_array [3976]string
var reply_email_array [3976]string
var datesend_array [3976]string
var date_array [3976]string
var month_array [3976]string
var year_array [3976]string
var day_array [3976]string
var time_array [3976]string
var subject_array [3976]string
var contentemail_array [3976]string

func main() {
	data, err := ioutil.ReadFile("fradulent_emails_dataset.txt")
	checkerror(err)
	emails_data_str := string(data)
	emails := s.Split(emails_data_str, "From r ")
	emails = emails[1:]

	fmt.Println("Total : ", len(emails))

	for index, email := range emails {

		// fmt.Println("----- Email No. ", index+1, " ------")
		sender := regexp.MustCompile("\nFrom:.*")
		sender_data := sender.FindString(email)
		// fmt.Println(sender_data)

		s_email := regexp.MustCompile(`\w+\S@(\w+|\d+|\.)+`)
		emails_sender := s_email.FindString(sender_data)
		// fmt.Println(emails_sender)
		// emails_sender := s.Trim(email_sender, "<")

		if emails_sender != "" {
			// fmt.Println(emails_sender)
			sender_email_array[index] = emails_sender
			email_account := s.Split(emails_sender, "@")
			// fmt.Println(email_account)
			username_array[index] = email_account[0]
			domainname_array[index] = email_account[1]
			// fmt.Println(email_account[1])
			if s.Contains(email_account[1], ".") {
				domaingroup := s.Split(email_account[1], ".")
				// fmt.Println(domaingroup)
				domaingroup_array[index] = domaingroup[1]
			}
		} else {
			sender_email_array[index] = "None"
			username_array[index] = "None"
			domainname_array[index] = "None"
			domaingroup_array[index] = "None"
		}

		s_name := regexp.MustCompile(`From:\s+.*\s+`)
		name_sender := s_name.FindString(sender_data)
		name_sender = s.Replace(name_sender, "From: ", "", -1)
		// fmt.Println(name_sender)
		if s.Contains(name_sender, "\"") {
			name_sender = s.Replace(name_sender, "\"", "", -1)
			// fmt.Println(name_sender)
		}
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
		date_re := regexp.MustCompile(`\d+.\w+.\d+.`)
		date_my := date_re.FindString(date_data)
		// fmt.Println(index+1, " : ", date_my)
		date_d := regexp.MustCompile(`\d{2}`)
		dateD := date_d.FindString(date_my)
		// fmt.Println(dateD)
		date_m := regexp.MustCompile(`\w{3}`)
		dateM := date_m.FindString(date_my)
		// fmt.Println(dateM)
		date_y := regexp.MustCompile(`\d{4}`)
		dateY := date_y.FindString(date_my)
		// fmt.Println(dateY)
		date_week := regexp.MustCompile(`\s\w{3}`)
		day := date_week.FindString(date_data)
		time_date := regexp.MustCompile(`\d{2}:\d{2}:\d{2}\s.\d{4}`)
		time := time_date.FindString(date_data)
		// fmt.Println(day, "\t", time)
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

		subject := regexp.MustCompile(`\nSubject:\s*.*`)
		subject_data := subject.FindString(email)
		subject_data = s.Replace(subject_data, "\nSubject: ", "", -1)
		// fmt.Println("", index+1, " : ", subject_data)
		if subject_data != "" {
			subject_array[index] = subject_data
		} else {
			subject_array[index] = "None"
		}

		content := regexp.MustCompile(`\nStatus:\s+[a-zA-Z0-9]+(\n|[0-9A-Za-z_]+|.*|\s)+`)
		content_data := content.FindString(email)
		content_data = s.Replace(content_data, "\nStatus: ", "", -1)
		// fmt.Println("", index+1, " : ", content_data)
		if content_data != "" {
			contentemail_array[index] = content_data
		} else {
			contentemail_array[index] = "None"
		}
	}

	fmt.Println("sender_email_array : ", len(sender_email_array))
	fmt.Println("sender_name_array : ", len(sender_name_array))
	fmt.Println("username_array : ", len(username_array))
	fmt.Println("domainname_array : ", len(domainname_array))
	fmt.Println("domaingroup_array : ", len(domaingroup_array))
	fmt.Println("recipient_email_array : ", len(recipient_email_array))
	fmt.Println("reply_email_array : ", len(reply_email_array))
	fmt.Println("datesend_array : ", len(datesend_array))
	fmt.Println("date_array : ", len(date_array))
	fmt.Println("month_array : ", len(month_array))
	fmt.Println("year_array : ", len(year_array))
	fmt.Println("day_array : ", len(day_array))
	fmt.Println("time_array : ", len(time_array))
	fmt.Println("subject_array : ", len(subject_array))
	fmt.Println("content_data : ", len(contentemail_array))

}
