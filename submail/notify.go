package submail

import (
	"fmt"
)

func Notify(issueNo, title, link string) {
	messageconfig := make(map[string]string)
	messageconfig["appid"] = "17028"
	messageconfig["appkey"] = "d2388e0c7deccf527fed3ab68b56e38b"
	messageconfig["signtype"] = "md5"

	messagexsend := CreateMessageXSend()
	MessageXSendAddTo(messagexsend, "13501147622")
	MessageXSendSetProject(messagexsend, "WMlRw3")
	MessageXSendAddVar(messagexsend, "issueNo", issueNo)
	MessageXSendAddVar(messagexsend, "title", title)
	MessageXSendAddVar(messagexsend, "findout", link)
	fmt.Println("MessageXSend ", MessageXSendRun(MessageXSendBuildRequest(messagexsend), messageconfig))
}
