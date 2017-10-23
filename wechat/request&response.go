package wechat

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgID        int
}

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateURL(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	log.Println("Wechat Service: validateUrl Ok!")

	if r.Method == "POST" {
		textRequestBody := parseTextRequestBody(r)
		if textRequestBody != nil {
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!",
				textRequestBody.Content,
				textRequestBody.FromUserName)
		}

		responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
			textRequestBody.FromUserName,
			"Hello, "+textRequestBody.FromUserName)
		if err != nil {
			log.Println("Wechat Service: makeTextResponseBody error: ", err)
			return
		}
		fmt.Fprintf(w, string(responseTextBody))
	}

}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{}
	textResponseBody.FromUserName = fromUserName
	textResponseBody.ToUserName = toUserName
	textResponseBody.MsgType = "text"
	textResponseBody.Content = content
	textResponseBody.CreateTime = time.Duration(time.Now().Unix())
	return xml.MarshalIndent(textResponseBody, " ", "  ")
}

//recvtextmsg_unencrypt.go
func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}
