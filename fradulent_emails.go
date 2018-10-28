package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	s "strings"
)

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func counter(data_array [3976]string) map[string]int {
	counter := make(map[string]int)
	for _, row := range data_array {
		counter[row]++
	}
	return counter
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

type Countryz struct {
	CountryName string
	Count       int
}

type Countriess []Countryz

var sender_email_array [3976]string
var sender_name_array [3976]string
var username_array [3976]string
var domainname_array [3976]string
var domaingroup_array [3976]string
var recipient_email_array [3976]string
var reply_email_array [3976]string
var match_email_array [3976]string
var datesend_array [3976]string
var date_month_array [3976]string
var date_array [3976]string
var month_array [3976]string
var year_array [3976]string
var day_array [3976]string
var time_array [3976]string
var day_time_array [3976]string
var subject_array [3976]string
var subject_word_array []string
var contentemail_array [3976]string
var countryArray [3976]string

func main() {
	data, err := ioutil.ReadFile("fradulent_emails_dataset.txt")
	checkerror(err)
	emails_data_str := string(data)
	emails := s.Split(emails_data_str, "From r ")
	emails = emails[1:]

	d, er := ioutil.ReadFile("country.txt")
	checkerror(er)
	country_data_str := string(d)
	countries := s.Split(country_data_str, "\r\n")

	fmt.Println("Total : ", len(emails))

	for index, email := range emails {

		sender := regexp.MustCompile("\nFrom:.*")
		sender_data := sender.FindString(email)
		s_email := regexp.MustCompile(`\w+\S@(\w+|\d+|\.)+`)
		emails_sender := s_email.FindString(sender_data)

		if emails_sender != "" {

			sender_email_array[index] = emails_sender
			email_account := s.Split(emails_sender, "@")
			username_array[index] = email_account[0]
			domainname_array[index] = email_account[1]
			if s.Contains(email_account[1], ".") {
				domaingroup := s.Split(email_account[1], ".")
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
		if s.Contains(name_sender, "\"") {
			name_sender = s.Replace(name_sender, "\"", "", -1)
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
		// fmt.Println("", index+1, " : ", recipient_data)
		if recipient_data != "" {
			recipient_email_array[index] = recipient_data
		} else {
			recipient_email_array[index] = "None"
		}

		// reply email
		reply := regexp.MustCompile("\nReply-To:.*")
		reply_data := reply.FindString(email)
		reply_data = s.Replace(reply_data, "\nReply-To: ", "", -1)
		re_email := regexp.MustCompile(`\w+\S@(\w+|\d+|\.)+`)
		reply_emails_sender := re_email.FindString(reply_data)

		// if s.Contains(reply_data, "<") {
		// 	reply_data = s.Replace(reply_data, "<", "", -1)
		// 	reply_data = s.Replace(reply_data, ">", "", -1)
		// }
		// fmt.Println("", index+1, " : ", reply_emails_sender)

		if reply_data != "" {
			reply_email_array[index] = reply_emails_sender
		} else {
			reply_email_array[index] = "None"
		}

		// match email sender and reply-to
		if emails_sender == reply_data {
			match_email_array[index] = "Match"
		} else {
			match_email_array[index] = "non-Match"
		}

		// date
		date := regexp.MustCompile("Date:.*")
		date_data := date.FindString(email)
		// fmt.Println(date_data)
		date_re := regexp.MustCompile(`\d+.\w+.\d+.`)
		date_my := date_re.FindString(date_data)
		date_month_find := regexp.MustCompile(`\d+.\w+`)
		date_month := date_month_find.FindString(date_my)
		date_d := regexp.MustCompile(`\d{2}`)
		dateD := date_d.FindString(date_my)
		date_m := regexp.MustCompile(`\w{3}`)
		dateM := date_m.FindString(date_my)
		date_y := regexp.MustCompile(`\d{4}`)
		dateY := date_y.FindString(date_my)
		date_week := regexp.MustCompile(`\w{3}\,`)
		day := date_week.FindString(date_data)
		day = s.Replace(day, ",", "", -1)
		time_date := regexp.MustCompile(`\d{2}:\d{2}:\d{2}\s.\d{4}`)
		time := time_date.FindString(date_data)
		day_time_find := regexp.MustCompile(`\s\d{2}:`)
		day_time := day_time_find.FindString(date_data)
		day_time = s.Replace(day_time, " ", "", -1)
		day_time = s.Replace(day_time, ":", "", -1)
		day_time_new := day + " " + day_time
		// fmt.Println(day_time_new)

		if date_data != "" {
			datesend_array[index] = date_my
			date_month_array[index] = date_month
			date_array[index] = dateD
			month_array[index] = dateM
			year_array[index] = dateY
			day_array[index] = day
			time_array[index] = time
			day_time_array[index] = day_time_new
		} else {
			datesend_array[index] = "None"
			date_month_array[index] = "None"
			date_array[index] = "None"
			month_array[index] = "None"
			year_array[index] = "None"
			day_array[index] = "None"
			time_array[index] = "None"
			day_time_array[index] = "None"
		}

		subject := regexp.MustCompile(`\nSubject:\s*.*`)
		subject_data := subject.FindString(email)
		subject_data = s.Replace(subject_data, "\nSubject: ", "", -1)
		// fmt.Println(subject_data)
		if subject_data != "" {
			subject_array[index] = subject_data
		} else {
			subject_array[index] = "None"
		}
		subject_word := s.Split(subject_data, " ")
		for _, word := range subject_word {
			subject_word_array = append(subject_word_array, word)
		}
		// fmt.Println(subject_word)

		content := regexp.MustCompile(`\nStatus:\s+[a-zA-Z0-9]+(\n|[0-9A-Za-z_]+|.*|\s)+`)
		content_data := content.FindString(email)
		content_data = s.Replace(content_data, "\nStatus: ", "", -1)
		// fmt.Println("", index+1, " : ", content_data)
		if content_data != "" {
			contentemail_array[index] = content_data
		} else {
			contentemail_array[index] = "None"
		}
		for indx, country := range countries {
			// fmt.Println("----- Mail. ", index+1, " ------")
			countryArray[indx] = country
			if s.Contains(contentemail_array[index], countryArray[indx]) {
				var countrys = Countriess{
					Countryz{
						countryArray[indx],
						s.Count(contentemail_array[index], countryArray[indx]),
					},
				}
				// Found!
				pagesJson, err := json.Marshal(countrys)
				if err != nil {
					log.Fatal("Cannot encode to JSON ", err)
				}
				fmt.Print("----- Mail. ", index+1, " ------", " index: ")
				fmt.Fprintf(os.Stdout, "%s", pagesJson)
				//fmt.Print(" Time(s)")
				fmt.Println("")
				// fmt.Println("----- Mail. ", index+1, " ------", " Country Found: ", countryArray[indx], " Count: ", s.Count(contentemail_array[index], countryArray[indx]))

			}
		}

	}

	// sender_email_count := counter(sender_email_array)
	// sender_name_count := counter(sender_name_array)
	// username_count := counter(username_array)
	// domainname_count := counter(domainname_array)
	// domaingroup_count := counter(domaingroup_array)
	// day_count := counter(day_array)
	// date_count := counter(date_array)
	// month_count := counter(month_array)
	// year_count := counter(year_array)
	// subject_word_count := counter(subject_word_array)
	// day_time_count := counter(day_time_array)
	// date_month_count := counter(date_month_array)
	// match_email_count := counter(match_email_array)
	// datesend_count := counter(datesend_array)

	// fmt.Println("sender_email_count : ", sender_email_count)
	// fmt.Println("sender_name_count : ", sender_name_count)
	// fmt.Println("username_count : ", username_count)
	// fmt.Println("domainname_count : ", domainname_count)
	// fmt.Println("domaingroup_count : ", domaingroup_count)
	// fmt.Println("day_count : ", day_count)
	// fmt.Println("date_count : ", date_count)
	// fmt.Println("month_count : ", month_count)
	// fmt.Println("year_count : ", year_count)

	// fmt.Println("sender_email_count : ", sender_email_count)
	// fmt.Println(subject_word_array)
	// fmt.Println(len(subject_word_array))
	// fmt.Println(subject_word_count)

	// sort value in map
	// n := map[int][]string{}
	// var a []int
	// for k, v := range year_count {
	// 	n[v] = append(n[v], k)
	// }
	// for k := range n {
	// 	a = append(a, k)
	// }
	// sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// for _, k := range a {
	// 	for _, s := range n[k] {
	// 		fmt.Printf("%s, %d\n", s, k)
	// 	}
	// }

}
