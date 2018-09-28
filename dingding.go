package go_base_libs

import (
	"encoding/json"
	"net/http"
	"crypto/tls"
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
)

//叮叮告警是一个接口，下面应该对应多种格式

var Dingding = & dingding{}

type dingding struct {
	PostStruct postStruct
	Url string

}

type postStruct struct {
	Msgtype string `json:"msgtype"`
	Text text `json:"text"`
}

type text struct {
	Content string `json:"content"`
}

func (d *dingding) SetUrl(url string) * dingding{
	d.Url = url
	return d
}

func (d * dingding) SetContent(content string ) * dingding{
	d.PostStruct.Text.Content = content
	return d
}

func (d * dingding)  SetMsgtype(msgType string) * dingding{
	d.PostStruct.Msgtype = msgType
	return d
}

func (d * dingding) SendDingDingAlert(content string) string {

	d.SetContent(content)

	postJson,err := json.Marshal(d.PostStruct)

	if err != nil{
		log.Error("转换结构体为json格式失败，请检查程序", err)
		return
	}
	//https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http 主体
	req, err := http.NewRequest("POST", d.Url, bytes.NewBuffer(postJson))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport:tr}
	resp, herr := client.Do(req)
	if herr != nil {
		log.Error("给叮叮发送告警信息失败", herr)
		return
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