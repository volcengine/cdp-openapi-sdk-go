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

type IdTypeDetail struct {
	Id int32 `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	DataType string `json:"dataType,omitempty"`
	EntityType string `json:"entityType,omitempty"`
	NeedEncrypt bool `json:"needEncrypt,omitempty"`
	SubjectId int32 `json:"subjectId,omitempty"`
	SubjectName string `json:"subjectName,omitempty"`
	OrgId int32 `json:"orgId,omitempty"`
	SubjectType string `json:"subjectType,omitempty"`
	AvailableMappingIds []AvailableMappingIdsDetail `json:"availableMappingIds,omitempty"`
}