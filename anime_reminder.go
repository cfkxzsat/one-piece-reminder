package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

const onePieceURL string = "http://www.onepiece.cc/comic"

var nextIssueRegex = regexp.MustCompile(`<p\s+class\s*=\s*"next"\s*>\s*第(\d+)话\s+预计(\d+)月(\d+)日.*</p>`)

type issueInfo struct {
	issueNo  string
	nextDate myDate
}

type myDate struct {
	month int
	day   int
}

func main() {

	updateInfo()

}

func updateInfo() {
	req, _ := http.NewRequest("GET", onePieceURL, nil)
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == 200 {
		rb := res.Body
		defer rb.Close()

		b, _ := ioutil.ReadAll(rb)
		htmlStr := string(b)

		ptag := nextIssueRegex.FindString(htmlStr)

		//	s := `<p class="next"> 第883话 预计10月27日晚上更新</p>`
		//	b := nextIssueRegex.MatchString(s) //Q:tab also match \s...
		cg := nextIssueRegex.FindStringSubmatch(ptag) //stands for capture group [0] whole string [1] issue num [2] month [3] year

		m, _ := strconv.Atoi(cg[2])
		d, _ := strconv.Atoi(cg[3])

		ii := issueInfo{
			issueNo: cg[1],
			nextDate: myDate{
				month: m,
				day:   d,
			},
		}

		fmt.Println(ii)
	}
}
