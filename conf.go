package go_base_libs

import (
	"fmt"
	"io/ioutil"

	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
)

var Conf = &conf{}

type conf struct {
	ServiceStruct interface{} //需要解析数据的结构体
	ConfigFile    string      //yml配置文件
}

func (c *conf) GetConf() *conf {
	log.Info("读取service的配置信息。")

	ymalFile, ierr := ioutil.ReadFile(c.ConfigFile)
	if ierr != nil {
		log.Error(fmt.Sprintf("读取项目的配置文件->%s->失败", ymalFile), ierr)
		panic(ierr)
	}

	uerr := yaml.Unmarshal(ymalFile, c.ServiceStruct)
	if uerr != nil {
		log.Error("反序列化servcer的配置文件失败", uerr)
		panic(uerr)
	}
	return c
}

//test
//package main
//
//import (
//baseLibs "github.com/zuoshenglo/go-base-libs"
//"fmt"
//)
//
//type esInfo struct {
//	Elasticsearch struct{
//		User string
//		Password string
//	}
//}
//
//func main() {
//	baseLibs.Conf.ServiceStruct = & esInfo{}
//	baseLibs.Conf.ConfigFile = "/Users/zuoshenglo/goland-workspace/src/github.com/zuoshenglo/conf/service.yml"
//	baseLibs.Conf.GetConf()
//	fmt.Println(baseLibs.Conf.ServiceStruct)
//}
