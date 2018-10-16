//从es的一个index中查询去重复的数据列表
package go_base_libs

import (
	"time"
	//"gopkg.in/olivere/elastic.v6"
)

type QueryDuplicateRemovalData struct {
	IndexName string
	FieldName string
	BeginTime time.Duration
	EndTime time.Duration
}

