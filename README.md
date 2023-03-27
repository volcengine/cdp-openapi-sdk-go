# CDP openAPI for GO

欢迎使用CDP openAPI SDK for GO, 本文档为您介绍如何获取及调用SDK。
此SDK 适用于CDP私部场景。
## 前置准备
###服务开通
如果调用高速openAPI,请确保已经开通在线服务。

###获取安全凭证
Access Key（访问密钥）是访问CDP openAPI服务的安全凭证，包含Access Key ID（简称为AK）和Secret Access Key（简称为SK）两部分。
您可以登录CDP, 点击“项目中心”->"资产输出"->"渠道管理"->"自定义渠道"， 点击“添加渠道应用”，创建并管理您的Access Key。

###环境检查
Go 版本需要不低于1.14

##获取与安装

go get -u github.com/volcengine/cdp-openapi-sdk-golang

##样例
```
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"github.com/volcengine/cdp-openapi-sdk-go"
)

func main() {
	testAk := "your-ak"
	testSk := "your-sk"
	basePath := "https://XXX/open_platform/openapi"
	
	httpCLient := http.Client{Timeout: 1 * time.Second}
	Config := Configuration{AccessKeyId: testAk,
	    AccessKeySecret: testSk,
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

```




