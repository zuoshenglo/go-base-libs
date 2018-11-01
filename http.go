package go_base_libs

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
)

func SendDingDingWebHook(sendData []byte, url string) string {

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

type HttpRequestCustom struct {
	body      []byte
	method    string
	url       string
	protocol  string
	transport *http.Transport
	user      string
	password  string
}

// protocol == http or https
func (hrc *HttpRequestCustom) SetRequestProtocol(protocol string) *HttpRequestCustom {

	switch protocol {
	case "https":
		hrc.transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	case "http":
		hrc.transport = &http.Transport{
			Dial:              dialTimeout,
			DisableKeepAlives: true,
		}
	}
	return hrc
}

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Second*10)
}

func (hrc *HttpRequestCustom) SetBasicAuth(user string, password string) *HttpRequestCustom {

	hrc.user = user
	hrc.password = password
	return hrc
}

func (hrc *HttpRequestCustom) ExecRequest() (string, error) {

	tr := hrc.transport

	//http 主体
	req, err := http.NewRequest(hrc.method, hrc.url, bytes.NewBuffer(hrc.body))
	if err != nil {
		return "", err
	}

	//
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(hrc.user, hrc.password)
	client := &http.Client{Transport: tr}
	resp, herr := client.Do(req)
	if herr != nil {
		return "", herr
	}
	defer resp.Body.Close()
	body, ierr := ioutil.ReadAll(resp.Body)
	if ierr != nil {
		return "", herr
	}
	return string(body), nil
}

// example
// res, err := baseLibs.NewHttpRequestCustom([]byte(queeryString), "POST", "http://172.16.28.120:9200/sjgc-logs-2018.10.30.01/_search").SetRequestProtocol("http").SetBasicAuth("admin", "1313GHGHG321dd").ExecRequest()
// if err != nil {
// 	fmt.Println(err)
// }
func NewHttpRequestCustom(body []byte, method string, url string) *HttpRequestCustom {
	return &HttpRequestCustom{
		body:   body,
		method: method,
		url:    url,
	}
}
