package go_base_libs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
	"io/ioutil"
)

type AliOss struct {
	Endpoint string
	AccessKeyId string
	AccessKeySecret string
	BucketName string
	client *oss.Client
	bucket *oss.Bucket
}

func NewOssClient(endpoint string, accessKeyId string, accessKeySecret string, bucketName string) * AliOss {
	return & AliOss{
		Endpoint: endpoint,
		AccessKeyId: accessKeyId,
		AccessKeySecret: accessKeySecret,
		BucketName: bucketName,
	}
}

// 初始化chuang
func (aloss * AliOss) InitOssClient() (* AliOss, error)  {
	client, err := oss.New(aloss.Endpoint, aloss.AccessKeyId, aloss.AccessKeySecret)

	if err != nil {
		return aloss, err
	} else {
		aloss.client = client
	}
	return aloss, nil
}

// 获得存储空间
func (aloss * AliOss) GetSaveSpace() (* AliOss, error) {
	bucket, err := aloss.client.Bucket(aloss.BucketName)
	if err != nil {
		return aloss, err
	} else {
		aloss.bucket = bucket
	}
	return aloss, nil
}

// 获得文件列表

func (aloss * AliOss) ListFiles(marKer string) [] string {

	returnStringSlice := make([] string, 0)
	marker := marKer
	for {
		lsRes, err := aloss.bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			//fmt.Println("Bucket: ", object.Key)
			returnStringSlice = append(returnStringSlice, object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}

	return returnStringSlice
}

// 流式 下载, 最后以字符串的形式，打印到 屏幕界面。
func (aloss * AliOss) DownloadFile(objectName string) (string, error) {

	body, err := aloss.bucket.GetObject( objectName )
	if err != nil {
		return "", err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string( data ), nil
	//fmt.Println("data:", string(data))
	//fmt.Println(strings.Split(string(data), "\r"))

	//for _,v := range strings.Split(string(data), "\n") {
	//	fmt.Println("test:", v)
	//}
	//return  body, nil
}

