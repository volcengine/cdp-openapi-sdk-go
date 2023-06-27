/*
 * Copyright 2022 ByteDance and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * CDP开放接口
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Linger please
var (
	_ context.Context
	_ os.File
)

type InsightApiService service

/*
InsightApiService 获取洞察跳转key
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body dsl结构
@return InlineResponse200
*/
func (a *InsightApiService) GetDSLInsightKey(ctx context.Context, body interface{}) (InlineResponse200, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse200
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action", "QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version", "2022-12-16")
	localVarQueryParams.Add("ApiAction", "GetDSLInsightKey")
	localVarQueryParams.Add("ApiVersion", "2023-02-10")

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if len(localVarHttpHeaderAccepts) > 0 {
		respType := localVarHttpResponse.Header.Values("Content-Type")
		for _, respType := range respType {
			for _, accept := range localVarHttpHeaderAccepts {
				if respType == accept {
					goto RESP_TYPE_CHECK_END
				}
			}
		}
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("Content-Type %v not accept, body: \"%v\"", respType, string(localVarBody))
	}
RESP_TYPE_CHECK_END:

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header)
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse200
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header)
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsightApiService 洞察报告详情
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenant
 * @param reportId
 * @param optional nil or *InsightApiGetInsightReportByIdOpts - Optional Parameters:
     * @param "XEnv" (optional.String) -
     * @param "SingleValue" (optional.Bool) -
     * @param "InsightSortObj" (optional.String) -
     * @param "SortType" (optional.String) -
     * @param "PDate" (optional.String) -
@return InlineResponse2002
*/

type InsightApiGetInsightReportByIdOpts struct {
	XEnv           optional.String
	SingleValue    optional.Bool
	InsightSortObj optional.String
	SortType       optional.String
	PDate          optional.String
}

func (a *InsightApiService) GetInsightReportById(ctx context.Context, xTenant int64, reportId int64, localVarOptionals *InsightApiGetInsightReportByIdOpts) (InlineResponse2002, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse2002
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("reportId", fmt.Sprintf("%v", reportId))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action", "QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version", "2022-12-16")
	localVarQueryParams.Add("ApiAction", "GetInsightReportById")
	localVarQueryParams.Add("ApiVersion", "2023-02-10")

	if localVarOptionals != nil && localVarOptionals.SingleValue.IsSet() {
		localVarQueryParams.Add("singleValue", parameterToString(localVarOptionals.SingleValue.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.InsightSortObj.IsSet() {
		localVarQueryParams.Add("insightSortObj", parameterToString(localVarOptionals.InsightSortObj.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortType.IsSet() {
		localVarQueryParams.Add("sortType", parameterToString(localVarOptionals.SortType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PDate.IsSet() {
		localVarQueryParams.Add("pDate", parameterToString(localVarOptionals.PDate.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(xTenant, "")
	if localVarOptionals != nil && localVarOptionals.XEnv.IsSet() {
		localVarHeaderParams["X-Env"] = parameterToString(localVarOptionals.XEnv.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if len(localVarHttpHeaderAccepts) > 0 {
		respType := localVarHttpResponse.Header.Values("Content-Type")
		for _, respType := range respType {
			for _, accept := range localVarHttpHeaderAccepts {
				if respType == accept {
					goto RESP_TYPE_CHECK_END
				}
			}
		}
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("Content-Type %v not accept, body: \"%v\"", respType, string(localVarBody))
	}
RESP_TYPE_CHECK_END:

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header)
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2002
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header)
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsightApiService 洞察报告列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenant
 * @param optional nil or *InsightApiGetInsightReportListOpts - Optional Parameters:
     * @param "XEnv" (optional.String) -
@return InlineResponse2001
*/

type InsightApiGetInsightReportListOpts struct {
	XEnv optional.String
}

func (a *InsightApiService) GetInsightReportList(ctx context.Context, xTenant int64, localVarOptionals *InsightApiGetInsightReportListOpts) (InlineResponse2001, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse2001
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action", "QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version", "2022-12-16")
	localVarQueryParams.Add("ApiAction", "GetInsightReportList")
	localVarQueryParams.Add("ApiVersion", "2023-02-10")

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(xTenant, "")
	if localVarOptionals != nil && localVarOptionals.XEnv.IsSet() {
		localVarHeaderParams["X-Env"] = parameterToString(localVarOptionals.XEnv.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if len(localVarHttpHeaderAccepts) > 0 {
		respType := localVarHttpResponse.Header.Values("Content-Type")
		for _, respType := range respType {
			for _, accept := range localVarHttpHeaderAccepts {
				if respType == accept {
					goto RESP_TYPE_CHECK_END
				}
			}
		}
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("Content-Type %v not accept, body: \"%v\"", respType, string(localVarBody))
	}
RESP_TYPE_CHECK_END:

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header)
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2001
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header)
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
InsightApiService 统计洞察报告查看次数
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param xTenant
 * @param optional nil or *InsightApiGetInsightReportUsageOpts - Optional Parameters:
     * @param "XEnv" (optional.String) -
     * @param "OpenapiOnly" (optional.Bool) -
@return InlineResponse2003
*/

type InsightApiGetInsightReportUsageOpts struct {
	XEnv        optional.String
	OpenapiOnly optional.Bool
}

func (a *InsightApiService) GetInsightReportUsage(ctx context.Context, xTenant int64, localVarOptionals *InsightApiGetInsightReportUsageOpts) (InlineResponse2003, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue InlineResponse2003
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action", "QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version", "2022-12-16")
	localVarQueryParams.Add("ApiAction", "GetInsightReportUsage")
	localVarQueryParams.Add("ApiVersion", "2023-02-10")

	if localVarOptionals != nil && localVarOptionals.OpenapiOnly.IsSet() {
		localVarQueryParams.Add("openapiOnly", parameterToString(localVarOptionals.OpenapiOnly.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["X-Tenant"] = parameterToString(xTenant, "")
	if localVarOptionals != nil && localVarOptionals.XEnv.IsSet() {
		localVarHeaderParams["X-Env"] = parameterToString(localVarOptionals.XEnv.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if len(localVarHttpHeaderAccepts) > 0 {
		respType := localVarHttpResponse.Header.Values("Content-Type")
		for _, respType := range respType {
			for _, accept := range localVarHttpHeaderAccepts {
				if respType == accept {
					goto RESP_TYPE_CHECK_END
				}
			}
		}
		return localVarReturnValue, localVarHttpResponse, fmt.Errorf("Content-Type %v not accept, body: \"%v\"", respType, string(localVarBody))
	}
RESP_TYPE_CHECK_END:

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header)
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InlineResponse2003
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header)
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
