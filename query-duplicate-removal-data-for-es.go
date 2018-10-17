//从es的一个index中查询去重复的数据列表
package go_base_libs

import (
	"time"
	//"gopkg.in/olivere/elastic.v6"
	"gopkg.in/olivere/elastic.v6"
	"context"
	"encoding/json"
)

var QueryDuplicateRemovalData = & queryDuplicateRemovalData{
	Type: "log",
	Size: 1000,
	Aggregation: "groupby",
	BucketsKeySlice: make([] string, 0),
}

type queryDuplicateRemovalData struct {
	IndexName string //必须
	FieldName string //必须
	Type string
	Aggregation string
	Size int
	BeginTime time.Duration
	EndTime time.Duration
	Res * elastic.SearchResult
	BucketsKeySlice [] string
}

func (q * queryDuplicateRemovalData) SetIndexName(indexName string) * queryDuplicateRemovalData{
	q.IndexName = indexName
	return q
}

func (q * queryDuplicateRemovalData) SetFieldName(filedName string) * queryDuplicateRemovalData {
	q.FieldName = filedName
	return q
}

func (q * queryDuplicateRemovalData) SetType(docType string) * queryDuplicateRemovalData{
	q.Type = docType
	return q
}

func (q * queryDuplicateRemovalData) SetAggregation(aggregation string)  * queryDuplicateRemovalData {
	q.Aggregation = aggregation
	return q
}

func (q * queryDuplicateRemovalData) SetSize(size int) * queryDuplicateRemovalData {
	q.Size = size
	return q
}

func (q * queryDuplicateRemovalData) QueryMain() error {

	OperateEs.User = "admin"
	OperateEs.Password = "1313GHGHG321dd"
	OperateEs.Address = "http://172.16.28.120:9200"
	OperateEs.initClient()
	esInitClient := elastic.NewTermsAggregation().Field(q.FieldName)
	res, err :=OperateEs.client.Search(q.IndexName).Type(q.Type).Size(q.Size).Aggregation(q.Aggregation, esInitClient).Do(context.Background())
	if err != nil {
		return err
		//fmt.Println("11212121", err)
	} else {
		q.Res = res
	}

	//解析数据
	var p AggregationsGroupby
	errJson := json.Unmarshal(* res.Aggregations["groupby"], & p)
	if errJson !=nil {
		return errJson
	}

	// 有元素的时候，才开始轮训,把key存于一个动态的数组slice中
	if len(p.Buckets) > 0 {
		for _, v := range p.Buckets {
			q.BucketsKeySlice = append(q.BucketsKeySlice, v.Key)
		}
	}

	return nil
}

type AggregationsGroupby struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount int `json:"sum_other_doc_count"`
	Buckets [] GroupbyBuckets
}

type GroupbyBuckets struct {
	Key string `json:"key"`
	DocCount int `json:"doc_count"`
}


