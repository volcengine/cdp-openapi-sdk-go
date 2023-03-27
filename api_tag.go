
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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"os"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
	_ os.File
)

type TagApiService service
/*
TagApiService 添加或者修改某个用户身上的人工标签
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param tenantCode 租户Code
@return ByteDanceNormalResponseString
*/
func (a *TagApiService) AddOrModifyManualTags(ctx context.Context, body ManualPersonTagRequest, tenantCode string) (ByteDanceNormalResponseString, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ByteDanceNormalResponseString
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("tenantCode",fmt.Sprintf("%v", tenantCode))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action","QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version","2022-12-16")
	localVarQueryParams.Add("ApiAction","AddOrModifyManualTags")
	localVarQueryParams.Add("ApiVersion","2023-02-10")

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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header);
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ByteDanceNormalResponseString
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
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
TagApiService 删除某个用户身上的人工标签
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenantCode 租户Code
 * @param baseId 基准ID
 * @param tagId 人工标签ID
@return ByteDanceNormalResponseString
*/
func (a *TagApiService) DeleteManualTagsInUser(ctx context.Context, tenantCode string, baseId int64, tagId int32) (ByteDanceNormalResponseString, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ByteDanceNormalResponseString
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("tenantCode",fmt.Sprintf("%v", tenantCode))
	localVarQueryParams.Add("baseId",fmt.Sprintf("%v", baseId))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action","QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version","2022-12-16")
	localVarQueryParams.Add("ApiAction","DeleteManualTagsInUser")
	localVarQueryParams.Add("ApiVersion","2023-02-10")

	localVarQueryParams.Add("tagId", parameterToString(tagId, ""))
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header);
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ByteDanceNormalResponseString
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
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
TagApiService 查询某个用户身上的人工标签列表
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenantCode 租户Code
@return ByteDanceSeqResponseOpenApiManualTagInfo
*/
func (a *TagApiService) GetManualTagsList(ctx context.Context, tenantCode string) (ByteDanceSeqResponseOpenApiManualTagInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ByteDanceSeqResponseOpenApiManualTagInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("tenantCode",fmt.Sprintf("%v", tenantCode))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action","QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version","2022-12-16")
	localVarQueryParams.Add("ApiAction","GetManualTagsList")
	localVarQueryParams.Add("ApiVersion","2023-02-10")

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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header);
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ByteDanceSeqResponseOpenApiManualTagInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
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
TagApiService 获取某个标签的取值
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenantCode 租户Code
 * @param tagId 标签id
@return ByteDanceSeqResponseTagValueResp
*/
func (a *TagApiService) GetTagValues(ctx context.Context, tenantCode string, tagId int32) (ByteDanceSeqResponseTagValueResp, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ByteDanceSeqResponseTagValueResp
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("tenantCode",fmt.Sprintf("%v", tenantCode))
	localVarQueryParams.Add("tagId",fmt.Sprintf("%v", tagId))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action","QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version","2022-12-16")
	localVarQueryParams.Add("ApiAction","GetTagValues")
	localVarQueryParams.Add("ApiVersion","2023-02-10")

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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header);
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ByteDanceSeqResponseTagValueResp
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
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
TagApiService 获取所有标签列表，包含各个标签的取值
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param tenantCode 租户Code
 * @param optional nil or *TagApiGetTagsListOpts - Optional Parameters:
     * @param "IdType" (optional.Int32) -  标签的idType，传每个主题下的base_id类型的ID
     * @param "DiscardValue" (optional.Bool) -  是否移除标签取值势力。默认是false
@return ByteDanceSeqResponseSeqDomainGroupedTags
*/

type TagApiGetTagsListOpts struct {
    IdType optional.Int32
    DiscardValue optional.Bool
}

func (a *TagApiService) GetTagsList(ctx context.Context, tenantCode string, localVarOptionals *TagApiGetTagsListOpts) (ByteDanceSeqResponseSeqDomainGroupedTags, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ByteDanceSeqResponseSeqDomainGroupedTags
	)

	// create path and map variables
	localVarPath := a.client.cfg.Host + a.client.cfg.BasePath
	localVarQueryParams := url.Values{}
	localVarQueryParams.Add("tenantCode",fmt.Sprintf("%v", tenantCode))

	localVarHeaderParams := make(map[string]string)

	localVarFormParams := url.Values{}
	localVarQueryParams.Add("Action","QueryOpenPlatformOpenApi")
	localVarQueryParams.Add("Version","2022-12-16")
	localVarQueryParams.Add("ApiAction","GetTagsList")
	localVarQueryParams.Add("ApiVersion","2023-02-10")

	if localVarOptionals != nil && localVarOptionals.IdType.IsSet() {
		localVarQueryParams.Add("idType", parameterToString(localVarOptionals.IdType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DiscardValue.IsSet() {
		localVarQueryParams.Add("discardValue", parameterToString(localVarOptionals.DiscardValue.Value(), ""))
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
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header);
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ByteDanceSeqResponseSeqDomainGroupedTags
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header);
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