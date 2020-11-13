package notice

import (
	"github.com/gpmgo/gopm/modules/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

///server酱的微信推送

func SendMsg(title, content string) {
	client := http.Client{}

	data := url.Values{
		"test": {title},
		"desp": {content},
	}
	req, _ := http.NewRequest("post", "https://sc.ftqq.com/SCU59910Ta14c7c99d7946deda54f1f1d70ca845b5d7705b0ded8a.send", strings.NewReader(data.Encode()))
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Info(string(body))
}
