package go_base_libs

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"sort"
	"time"
)

var Tool = &tool{}

type tool struct {
}

// 检查tcp 端口是否存在
// hostAddress 0.0.0.0 或者 0.0.0.0:80
func (t *tool) CheckPort(address string, port string) (bool, error) {
	conn, err := net.DialTimeout("tcp", address+":"+port, 2*time.Second)
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	defer conn.Close()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *tool) GetCwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (t *tool) GetEnv(envKey string) string {
	os.Getenv(envKey)
	return ""
}

// 删除文件
func (t *tool) DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println(err)
	}
}

func (t *tool) JsonStringToStruct() {

}

func (t *tool) StringToJson(formatString string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(formatString), &data); err != nil {
		return data, err
	}
	return data, nil
}

// 对字符串做json格式的二层解析
func (t *tool) StringToJsonJson(formatString string) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(formatString), &data); err != nil {
		return "", err
	}

	// 一层已经解析出来了
	var tmpString = ""
	for k, v := range data {
		if tmpStr, ok := v.(string); !ok {
			if tmpMap, ok := v.(map[string]interface{}); ok {
				for x, y := range tmpMap {
					tmpString = tmpString + "##### " + k + "." + x + " : " + fmt.Sprintf("%s", y) + "\n"
				}
			}
		} else {
			tmpString = tmpString + "##### " + k + " : " + tmpStr + "\n"
		}
	}
	return tmpString, nil
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

// 对map 按string key 来进行排序
func (t *tool) SortKeyForMap(sortMap map[string]interface{}) map[string]interface{} {
	s1 := make([]string, 0, len(sortMap))
	for k, _ := range sortMap {
		s1 = append(s1, k)
	}

	s2 := make(map[string]interface{})

	sort.Strings(s1)

	for _, v := range s1 {
		s2[v] = sortMap[v]
	}
	return s2
}

// 字符串 转为 一层json， 排序后， 返回为 string
func (t *tool) SortKeyForStringJson(formatString string) (string, error) {
	data, err := Tool.StringToJson(formatString)
	if err != nil {
		return "", err
	}
	data = Tool.SortKeyForMap(data)

	if jres, jerr := json.Marshal(data); jerr != nil {
		return "", jerr
	} else {
		return string(jres[:]), nil
	}
}

func (t *tool) GetNowTime() string {
	//return fmt.Sprintf("%s", time.Now())[:19]
	return time.Now().Format("2006-01-02 15:04:05")
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
