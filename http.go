package go_base_libs

import (
	"net/http"
	"crypto/tls"
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
	"fmt"
)

func SendDingDingWebHook(sendData [] byte, url string) string {

	//postJson,err := json.Marshal(sendDataStruct)
	//
	//if err != nil{
	//	log.Error("转换结构体为json格式失败，请检查程序", err)
	//	return
	//}
	//https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	//http 主体
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(sendData))
	if err != nil {
		log.Error("新建请求数据格式错误！")
		return fmt.Sprintf("新建请求数据格式错误！->%s", err)
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport:tr}
	resp, herr := client.Do(req)
	if herr != nil {
		log.Error("给叮叮发送告警信息失败", herr)
		return fmt.Sprintf("给叮叮发送告警信息失败:%s",herr)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("叮叮的返回信息为:", string(body))
	return string(body)
}