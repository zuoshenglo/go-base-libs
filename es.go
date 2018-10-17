package go_base_libs

import (
	"gopkg.in/olivere/elastic.v6"
	"fmt"
)

var OperateEs = & operateElasticsearch{}

// 定义使用es基础库的结构
type operateElasticsearch struct {
	User string
	Password string
	Address string
	client * elastic.Client
}

// es库初始化
func (e *operateElasticsearch) initClient() {
	client,err := elastic.NewClient(
		elastic.SetURL(e.Address),
		elastic.SetBasicAuth(e.User,e.Password),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false))

	if err != nil {
		fmt.Println("初始化失败",err)
		return
	}
	e.client = client
}

