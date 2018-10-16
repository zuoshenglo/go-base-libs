package go_base_libs

import (
	"gopkg.in/olivere/elastic.v6"
)

var OperateEs = & operateElasticsearch{}

type operateElasticsearch struct {
	esUser string
	esPassword string
	esAddress string
	client * elastic.Client
}


