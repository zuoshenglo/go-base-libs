package go_base_libs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/cihub/seelog"
)

//叮叮告警是一个接口，下面应该对应多种格式

var Dingding = &dingding{}

type dingding struct {
	PostStruct postStruct
	Url        string
}

type postStruct struct {
	Msgtype string `json:"msgtype"`
	Text    text   `json:"text"`
}

type text struct {
	Content string `json:"content"`
}

func (d *dingding) SetUrl(url string) *dingding {
	d.Url = url
	return d
}

func (d *dingding) SetContent(content string) *dingding {
	d.PostStruct.Text.Content = content
	return d
}

func (d *dingding) SetMsgtype(msgType string) *dingding {
	d.PostStruct.Msgtype = msgType
	return d
}

func (d *dingding) SendDingDingAlert(content string) string {

	d.SetContent(content)

	postJson, err := json.Marshal(d.PostStruct)

	if err != nil {
		log.Error("转换结构体为json格式失败，请检查程序", err)
		return fmt.Sprintf("转换结构体为json格式失败，请检查程序:%s", err)
	}
	//https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http 主体
	req, err := http.NewRequest("POST", d.Url, bytes.NewBuffer(postJson))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: tr}
	resp, herr := client.Do(req)
	if herr != nil {
		log.Error("给叮叮发送告警信息失败", herr)
		return fmt.Sprintf("给叮叮发送告警信息失败:%s", herr)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("叮叮的返回信息为:", string(body))
	return string(body)
}

// 使用
//dd.setMsgtype("text") # 叮叮告警结构的类型
//dd.setUrl(ServiceConf.Elasticsearch.Dingurl) # 叮叮的地址
//dd.setContent(content) # 告警的内容

type UserDingDing struct {
	MsgContent string `json:"msgContent"`
	Url        string `json:"url"`
	MsgType    string `json:"msgType"`
}

// 最基础的告警 结构体, dingding struct
type DingDingBase struct {
	Msgtype string `json:"msgtype,omitempty"`
	Text    struct {
		Content string `json:"content,omitempty"`
	} `json:"text,omitempty"`
}

type DingDingMarkdown struct {
	Msgtype string
	Markdown struct {
		Title string
		Text string
	}
	At struct {
		AtMobiles [] string
		IsAtAll bool
	}
}

func NewDingDingMarkdown() *DingDingMarkdown {
	return &DingDingMarkdown{
		Msgtype: "markdown",
		Markdown: struct {
			Title string
			Text  string
		}{Title: "", Text: ""},
		At: struct {
			AtMobiles [] string
			IsAtAll   bool
		}{AtMobiles: [] string{}, IsAtAll: false},
	}
}

func (dm * DingDingMarkdown) SetText(markdownText string) * DingDingMarkdown {
	dm.Markdown.Text = markdownText
	return dm
}

func (dm * DingDingMarkdown) SetTitle(markdownTitle string) * DingDingMarkdown {
	dm.Markdown.Title = markdownTitle
	return dm
}

func (dm * DingDingMarkdown) CreateMarkdownText(formatData map[string]interface{}) * DingDingMarkdown {
	dm.Markdown.Text = "geshihuacshi"
	return dm
}

func NewUserDingDing() *UserDingDing {
	return &UserDingDing{
		MsgType: "text",
	}
}

func (u *UserDingDing) SetMsgContent(msgContent string) *UserDingDing {
	u.MsgContent = msgContent
	return u
}

func (u *UserDingDing) SetUrl(url string) *UserDingDing {
	u.Url = url
	return u
}


// func NewBaseDingDingAlter(userReqBody []byte) ([]byte, error) {
// 	var userDingDing *UserDingDing
// 	if err := json.Unmarshal(userReqBody, userDingDing); err != nil {
// 		return []byte(""), err
// 	}

// 	//
// 	var dingDingBase = &DingDingBase{
// 		Msgtype: "text",
// 	}

// 	return []byte("111"), nil
// }
