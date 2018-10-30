package go_base_libs

import (
	"fmt"

	"gopkg.in/olivere/elastic.v6"
)

var OperateEs = &operateElasticsearch{}

// 新建一个es查询的初始化函数
func NewQueryStringQuery(user string, password string, address string) *operateElasticsearch {
	opes := &operateElasticsearch{
		User:     user,
		Password: password,
		Address:  address,
	}

	opes.initClient()

	return opes
}

// 定义使用es基础库的结构
type operateElasticsearch struct {
	User     string
	Password string
	Address  string
	Client   *elastic.Client
}

// es库初始化
func (e *operateElasticsearch) initClient() {
	client, err := elastic.NewClient(
		elastic.SetURL(e.Address),
		elastic.SetBasicAuth(e.User, e.Password),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false))

	if err != nil {
		fmt.Println("初始化失败", err)
		return
	}
	e.Client = client
}
