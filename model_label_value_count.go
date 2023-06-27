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

type LabelValueCount struct {
	// 当前标签枚举值覆盖人数
	Count int32 `json:"count,omitempty"`
	// 当前标签枚举值原始值
	OriginValue *interface{} `json:"originValue,omitempty"`
	// 当前标签枚举值映射值，如果未做映射则等于原始值
	MappingValue string `json:"mappingValue,omitempty"`
	// 当前标签枚举值映射值覆盖人数占当前标签总覆盖人数的比例
	TagCoverage string `json:"tagCoverage,omitempty"`
	// 当前标签枚举值映射值覆盖人数占当前项目总覆盖人数的比例
	TotalCoverage string `json:"totalCoverage,omitempty"`
}