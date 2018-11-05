package go_base_libs

import (
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
