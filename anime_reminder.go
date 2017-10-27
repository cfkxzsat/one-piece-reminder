package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/cfkxzsat/one-piece-reminder/submail"
)

const onePieceURL string = "http://www.one-piece.cn"
const listURL string = "/comic"

var nextIssueRegex = regexp.MustCompile(`<p\s+class\s*=\s*"next"\s*>\s*第(\d+)话\s+预计(\d+)月(\d+)日.*</p>`)

type IssueInfo struct {
	IssueNo  string
	NextDate MyDate
}

type MyDate struct {
	Month time.Month
	Day   int
}

var ii IssueInfo

func getHTMLStr() string {

	req, _ := http.NewRequest("GET", onePieceURL+listURL, nil)
	res, _ := http.DefaultClient.Do(req)

	for res.StatusCode != 200 {
		req, _ = http.NewRequest("GET", onePieceURL+listURL, nil)
		res, _ = http.DefaultClient.Do(req)
	}

	rb := res.Body
	defer rb.Close()

	b, _ := ioutil.ReadAll(rb)
	return string(b)

}

func UpdateInfo() {

	htmlStr := getHTMLStr()

	ptag := nextIssueRegex.FindString(htmlStr)

	cg := nextIssueRegex.FindStringSubmatch(ptag) //stands for capture group [0] whole string [1] issue num [2] month [3] year

	m, _ := strconv.Atoi(cg[2])
	d, _ := strconv.Atoi(cg[3])

	ii = IssueInfo{
		IssueNo: cg[1],
		NextDate: MyDate{
			Month: time.Month(m),
			Day:   d,
		},
	}
}

func main() {

	UpdateInfo()

	for {
		now := time.Now()
		location, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Println(err)
		}
		next := time.Date(now.Year(), ii.NextDate.Month, ii.NextDate.Day, 0, 0, 0, 0, location)
		//When around the last few days of the year, it is likely that we get a wrong date if we use the year of the current time.Should set the new year for next

		if now.After(next) {
			next = time.Date(now.Year()+1, ii.NextDate.Month, ii.NextDate.Day, 0, 0, 0, 0, location)
		}
		fmt.Println("I'm about to sleep")
		time.Sleep(next.Sub(now))
		fmt.Println("I awake now")

		ticker := time.NewTicker(time.Minute * 30)
		for {
			<-ticker.C
			if title, link, have := haveNewIssue(); have {
				submail.Notify(ii.IssueNo, title, link)
				break
			}
			fmt.Println("still not update...")
		}

		UpdateInfo()
	}

}

func haveNewIssue() (title, link string, have bool) {
	s := `<a\s+href\s*=\s*"(/post/\d+/)"[^>]*>第` + ii.IssueNo + `话\s+([^<]+)</a>`
	var newIssueLinkRegex = regexp.MustCompile(s)
	htmlStr := getHTMLStr()
	atag := newIssueLinkRegex.FindString(htmlStr)
	if atag == "" {
		return "", "", false
	}
	cg := nextIssueRegex.FindStringSubmatch(atag)
	link = cg[1]
	title = cg[2]
	return title, onePieceURL + link, true
}
