package swagger

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/antihax/optional"
)

var accessKeyId = "XXX"
var accessKeySecret = "XXX"
var basePath = "https://XXX/open_platform/openapi"

func TestGetSegment(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	responseBody, httpRespose, err := client.SegmentationApi.LegacyGetSegment(context.Background(), 1, 1000302)
	if err != nil {
		fmt.Printf("getSegment query fail, err:%v, statusCod:%v", err, httpRespose.StatusCode)
	} else {
		fmt.Printf("getSegment success, body is:%v, data is:%v", responseBody, *responseBody.Data)
	}
}

func TestGetSegmentList(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}
	opts := SegmentationApiLegacyGetSegmentListOpts{Current: optional.NewInt32(1), PageSize: optional.NewInt32(20)}
	responseBody, httpRespose, err := client.SegmentationApi.LegacyGetSegmentList(context.Background(), 1, &opts)
	if err != nil {
		fmt.Printf("getSegmentList query fail, err:%v, statusCod:%v", err, httpRespose.StatusCode)
	} else {
		fmt.Printf("getSegmentList success, body is:%v, data is:%v", responseBody, *responseBody.Data)
	}

	responseBody5, _, err5 := client.SegmentationApi.DownloadSegFile(context.Background(), 1, 1000560, "CSV", nil)
	if err5 != nil {
		fmt.Println("DownloadSegFile   query fail", err5)

	} else {
		data, err := ioutil.ReadAll(responseBody5)
		responseBody5.Close()
		if err != nil {
			fmt.Println("read error:", err)
		} else {
			fmt.Println("size of data is: ", string(data))
		}
	}
}

func TestDownLoadSegmention(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	opts := SegmentationApiDownloadSegFileOpts{optional.NewBool(false)}
	responseBody, httpRespose, err := client.SegmentationApi.DownloadSegFile(context.Background(), 1, 1000560, "CSV", &opts)
	if err != nil {
		fmt.Printf("DownloadSegFile query fail, err:%v, statusCod:%v", err, httpRespose.StatusCode)
	} else {
		//读取到内存，也可以读取到文件
		fmt.Println(responseBody.Name())
		responseBody.Seek(0, 0)
		data, err := ioutil.ReadAll(responseBody)
		responseBody.Close()
		if err != nil {
			fmt.Println("read error:", err)
		} else {
			fmt.Println("data is: ", string(data))
		}
	}
}

func TestUploadSegment(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	file, err := os.Open("./test_11675.csv")
	if err != nil {
		fmt.Errorf("open file err:%v", err)
		return
	}
	defer file.Close()
	responseBody, httpRespose, err := client.SegmentationApi.UploadSegFile(context.Background(), file, 1)
	if err != nil {
		fmt.Printf("UploadSegFile qfail, err:%v, statusCod:%v", err, httpRespose.StatusCode)

	} else {
		fmt.Printf("UploadSegFile succes:%+v", *responseBody.Data)
	}
}

func TestCreateUploadSegment(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	requeset := LegacyUploadedSegCreateRequest{Name: "testUpload", IdType: "baseid", Description: "test openapi", AbaseEnabled: false, SourcePlatform: "gateway", Detail: &LegacyUploadedSegmentDetail{
		123,
	}}

	responseBody, _, err := client.SegmentationApi.LegacyCreateUploadSegment(context.Background(), requeset, 1)
	if err != nil {
		fmt.Println("UploadSegFile query fail", err)

	} else {
		fmt.Printf("UploadSegFile succes:%+v", *responseBody.Data)
	}
}

func TestGetTagList(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}
	opts := TagApiGetTagsListOpts{IdType: optional.NewInt32(1)}
	responseBody, _, err := client.TagApi.GetTagsList(context.Background(), "2", &opts)
	if err != nil {
		fmt.Println("GetTagsList query fail", err)

	} else {
		fmt.Printf("GetTagsList succes:%+v", responseBody)
	}
}

func TestGetTagValues(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	responseBody, _, err := client.TagApi.GetTagValues(context.Background(), "2", 311)
	if err != nil {
		fmt.Println("GetTagsList query fail", err)

	} else {
		fmt.Printf("GetTagsList succes:%+v", responseBody)
	}
}

func TestGetManualTagsList(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	responseBody, _, err := client.TagApi.GetManualTagsList(context.Background(), "1")
	if err != nil {
		fmt.Println("GetManualTagsList query fail", err)

	} else {
		fmt.Printf("GetManualTagsList succes:%+v", responseBody)
		for _, v := range responseBody.Data {
			if v.TagValues != nil {
				fmt.Println("tag value is", *v.TagValues)
			}
		}
	}
}

func TestAddOrModifyManualTags(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}
	newTagValues := make([]TagResult, 0, 3)
	var intValue interface{} = 10
	tag1 := TagResult{TagId: 59, Value: &intValue}
	newTagValues = append(newTagValues, tag1)

	//var stringValue interface{} = "fdwe"
	//tag2 := TagResult{TagId: 60,Value: &stringValue}
	//newTagValues = append(newTagValues, tag2)
	//
	//var sliceValue interface{} = []string{"男","女"}
	//tag3 :=  TagResult{TagId: 61, Value: &sliceValue}
	//newTagValues = append(newTagValues, tag3)
	//
	//var doubleValue interface{} = 10.4
	//tag4 := TagResult{TagId: 62,Value: &doubleValue}
	//newTagValues = append(newTagValues, tag4)

	request := ManualPersonTagRequest{&IdFilter{Id: "1223", Type_: "baseid"}, newTagValues}
	responseBody, _, err := client.TagApi.AddOrModifyManualTags(context.Background(), request, "1")
	if err != nil {
		fmt.Println("AddOrModifyManualTags query fail", err)

	} else {
		fmt.Printf("AddOrModifyManualTags succes:%+v", responseBody)
	}
}

func TestDeleteManualTagsInUser(t *testing.T) {
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: accessKeyId, AccessKeySecret: accessKeySecret,
		BasePath: basePath, HTTPClient: &httpCLient}
	client, err := NewAPIClient(&Config)
	if err != nil {
		fmt.Println("NewAPIClient err", err)
		return
	}

	responseBody, _, err := client.TagApi.DeleteManualTagsInUser(context.Background(), "1", 12233, 59)
	if err != nil {
		fmt.Println("DeleteManualTagsInUser query fail", err)
	} else {
		fmt.Printf("DeleteManualTagsInUser succes:%+v", responseBody)
	}
}
