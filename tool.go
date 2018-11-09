package go_base_libs

import (
	"net/http"
	"os"
)

var Tool = &tool{}

type tool struct {
}

func (t *tool) GetCwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (t *tool) JsonStringToStruct() {

}

//去重复
func (t *tool) RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//http服务设置跨域， 添加头部
func (t *tool) SetCrossDomain(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	w.Header().Set("Access-Control-Max-Age", "1000")
	w.Header().Set("Access-Control-Allow-Headers", "AccessKey,Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With,X-Request-Id,X-Server-Addr,AppToken,PicAuth,kbn-version")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type,Origin,User-Agent,X-Requested-With,X-Request-Id,X-Server-Addr")
	return w
}

//nginx 设置
// #add_header 'Access-Control-Allow-Origin' '*';
// add_header 'Access-Control-Allow-Credentials' 'true';
// add_header 'Access-Control-Allow-Methods' 'GET, POST, PUT, DELETE, OPTIONS, HEAD';
// add_header 'Access-Control-Allow-Headers' 'content-type,AccessKey,Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Mx-ReqToken,X-Requested-With,X-Request-Id,X-Server-Addr,AppToken,PicAuth,kbn-version';
// add_header 'Access-Control-Expose-Headers' 'Content-Type,Origin,User-Agent,X-Requested-With,X-Request-Id,X-Server-Addr';
// int To string  string := strconv.Itoa(int)
// int64 To string string := strconv.FormatInt(int64,10)
