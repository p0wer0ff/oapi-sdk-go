// Package sheets code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larksheets

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *SheetsService {
	s := &SheetsService{config: config}
	s.Spreadsheet = &spreadsheet{service: s}
	s.SpreadsheetSheet = &spreadsheetSheet{service: s}
	s.SpreadsheetSheetFilter = &spreadsheetSheetFilter{service: s}
	s.SpreadsheetSheetFilterView = &spreadsheetSheetFilterView{service: s}
	s.SpreadsheetSheetFilterViewCondition = &spreadsheetSheetFilterViewCondition{service: s}
	s.SpreadsheetSheetFloatImage = &spreadsheetSheetFloatImage{service: s}
	return s
}

type SheetsService struct {
	config                              *larkcore.Config
	Spreadsheet                         *spreadsheet                         // 表格
	SpreadsheetSheet                    *spreadsheetSheet                    // 单元格
	SpreadsheetSheetFilter              *spreadsheetSheetFilter              // 筛选
	SpreadsheetSheetFilterView          *spreadsheetSheetFilterView          // 筛选视图
	SpreadsheetSheetFilterViewCondition *spreadsheetSheetFilterViewCondition // 筛选条件
	SpreadsheetSheetFloatImage          *spreadsheetSheetFloatImage          // 浮动图片
}

type spreadsheet struct {
	service *SheetsService
}
type spreadsheetSheet struct {
	service *SheetsService
}
type spreadsheetSheetFilter struct {
	service *SheetsService
}
type spreadsheetSheetFilterView struct {
	service *SheetsService
}
type spreadsheetSheetFilterViewCondition struct {
	service *SheetsService
}
type spreadsheetSheetFloatImage struct {
	service *SheetsService
}

// 创建表格
//
// - 使用该接口可以在指定的目录下创建在线表格。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/create_spreadsheet.go
func (s *spreadsheet) Create(ctx context.Context, req *CreateSpreadsheetReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查找单元格
//
// - 按照指定的条件查找子表的某个范围内的数据符合条件的单元格位置。请求体中的 range 和 find 字段为必填。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/find
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/find_spreadsheetSheet.go
func (s *spreadsheetSheet) Find(ctx context.Context, req *FindSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*FindSpreadsheetSheetResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/find"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &FindSpreadsheetSheetResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 移动行列
//
// - 该接口用于移动行列，行列被移动到目标位置后，原本在目标位置的行列会对应右移或下移。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/move_dimension
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/moveDimension_spreadsheetSheet.go
func (s *spreadsheetSheet) MoveDimension(ctx context.Context, req *MoveDimensionSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*MoveDimensionSpreadsheetSheetResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/move_dimension"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MoveDimensionSpreadsheetSheetResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 替换单元格
//
// - 按照指定的条件查找子表的某个范围内的数据符合条件的单元格并替换值，返回替换成功的单元格位置。一次请求最多允许替换5000个单元格，如果超过请将range缩小范围再操作。请求体中的 range、find、replaccement 字段必填。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet/replace
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/replace_spreadsheetSheet.go
func (s *spreadsheetSheet) Replace(ctx context.Context, req *ReplaceSpreadsheetSheetReq, options ...larkcore.RequestOptionFunc) (*ReplaceSpreadsheetSheetResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/replace"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ReplaceSpreadsheetSheetResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建筛选
//
// - 在子表内创建筛选。
//
// - 参数值可参考[筛选指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/create_spreadsheetSheetFilter.go
func (s *spreadsheetSheetFilter) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除筛选
//
// - 删除子表的筛选
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/delete_spreadsheetSheetFilter.go
func (s *spreadsheetSheetFilter) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取筛选
//
// - 获取子表的详细筛选信息
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/get_spreadsheetSheetFilter.go
func (s *spreadsheetSheetFilter) Get(ctx context.Context, req *GetSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新筛选
//
// - 更新子表筛选范围中的列筛选条件。
//
// - 参数值可参考[筛选指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/filter-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter/update
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/update_spreadsheetSheetFilter.go
func (s *spreadsheetSheetFilter) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterReq, options ...larkcore.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建筛选视图
//
// - 根据传入的参数创建一个筛选视图。Id 和 名字可选，不填的话会默认生成；range 必填。Id 长度为10，由 0-9、a-z、A-Z 组合生成。名字长度不超过100。单个子表内的筛选视图个数不超过 150。
//
// - 筛选范围的设置参考：[筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/create_spreadsheetSheetFilterView.go
func (s *spreadsheetSheetFilterView) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除筛选视图
//
// - 删除指定 id 对应的筛选视图。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/delete_spreadsheetSheetFilterView.go
func (s *spreadsheetSheetFilterView) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取筛选视图
//
// - 获取指定筛选视图 id 的名字和筛选范围。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/get_spreadsheetSheetFilterView.go
func (s *spreadsheetSheetFilterView) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新筛选视图
//
// - 更新筛选视图的名字或者筛选范围。名字长度不超过100，不能重复即子表内唯一；筛选范围不超过子表的最大范围。
//
// - 筛选范围的设置参考：[筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/patch_spreadsheetSheetFilterView.go
func (s *spreadsheetSheetFilterView) Patch(ctx context.Context, req *PatchSpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*PatchSpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFilterViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询筛选视图
//
// - 查询子表内所有的筛选视图基本信息，包括 id、name 和 range
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/query_spreadsheetSheetFilterView.go
func (s *spreadsheetSheetFilterView) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建筛选条件
//
// - 在筛选视图的筛选范围的某一列创建筛选条件。
//
// - 筛选条件参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/create_spreadsheetSheetFilterViewCondition.go
func (s *spreadsheetSheetFilterViewCondition) Create(ctx context.Context, req *CreateSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFilterViewConditionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除筛选条件
//
// - 删除筛选视图的筛选范围某一列的筛选条件。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/delete_spreadsheetSheetFilterViewCondition.go
func (s *spreadsheetSheetFilterViewCondition) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFilterViewConditionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取筛选条件
//
// - 获取筛选视图某列的筛选条件信息。
//
// - 筛选条件含义参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/get_spreadsheetSheetFilterViewCondition.go
func (s *spreadsheetSheetFilterViewCondition) Get(ctx context.Context, req *GetSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFilterViewConditionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询筛选条件
//
// - 查询一个筛选视图的所有筛选条件，返回筛选视图的筛选范围内的筛选条件。
//
// - 筛选条件含义可参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/query_spreadsheetSheetFilterViewCondition.go
func (s *spreadsheetSheetFilterViewCondition) Query(ctx context.Context, req *QuerySpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFilterViewConditionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新筛选条件
//
// - 更新筛选视图范围的某列的筛选条件，condition id 即为列的字母号。
//
// - 筛选条件参数可参考 [筛选视图的筛选条件指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/filter-view-condition-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-filter_view-condition/update
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/update_spreadsheetSheetFilterViewCondition.go
func (s *spreadsheetSheetFilterViewCondition) Update(ctx context.Context, req *UpdateSpreadsheetSheetFilterViewConditionReq, options ...larkcore.RequestOptionFunc) (*UpdateSpreadsheetSheetFilterViewConditionResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/filter_views/:filter_view_id/conditions/:condition_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateSpreadsheetSheetFilterViewConditionResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建浮动图片
//
// - 根据传入的参数创建一张浮动图片。Float_image_token （[上传图片至表格后得到](/ssl:ttdoc/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_all)）和range（只支持一个单元格） 必填。Float_image_id 可选，不填的话会默认生成，长度为10，由 0-9、a-z、A-Z 组合生成。表格内不重复的图片（浮动图片+单元格图片）总数不超过4000。width 和 height 为图片展示的宽高，可选，不填的话会使用图片的真实宽高。offset_x 和 offset_y 为图片左上角距离所在单元格左上角的偏移，可选，默认为 0。
//
// - 浮动图片的设置参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/create_spreadsheetSheetFloatImage.go
func (s *spreadsheetSheetFloatImage) Create(ctx context.Context, req *CreateSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*CreateSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateSpreadsheetSheetFloatImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除浮动图片
//
// - 删除 float_image_id 对应的浮动图片。
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/delete_spreadsheetSheetFloatImage.go
func (s *spreadsheetSheetFloatImage) Delete(ctx context.Context, req *DeleteSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*DeleteSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteSpreadsheetSheetFloatImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取浮动图片
//
// - 根据 float_image_id 获取对应浮动图片的信息。
//
// - 浮动图片参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/get_spreadsheetSheetFloatImage.go
func (s *spreadsheetSheetFloatImage) Get(ctx context.Context, req *GetSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*GetSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetSpreadsheetSheetFloatImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新浮动图片
//
// - 更新已有的浮动图片位置和宽高，包括 range、width、height、offset_x 和 offset_y，不包括 float_image_id 和 float_image_token。
//
// - 浮动图片更新参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/patch
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/patch_spreadsheetSheetFloatImage.go
func (s *spreadsheetSheetFloatImage) Patch(ctx context.Context, req *PatchSpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*PatchSpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/:float_image_id"
	apiReq.HttpMethod = http.MethodPatch
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchSpreadsheetSheetFloatImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询浮动图片
//
// - 返回子表内所有的浮动图片信息。
//
// - 浮动图片参考：[浮动图片指南](/ssl:ttdoc/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/float-image-user-guide)
//
// - 官网API文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/sheets-v3/spreadsheet-sheet-float_image/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/sheetsv3/query_spreadsheetSheetFloatImage.go
func (s *spreadsheetSheetFloatImage) Query(ctx context.Context, req *QuerySpreadsheetSheetFloatImageReq, options ...larkcore.RequestOptionFunc) (*QuerySpreadsheetSheetFloatImageResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/sheets/v3/spreadsheets/:spreadsheet_token/sheets/:sheet_id/float_images/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QuerySpreadsheetSheetFloatImageResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
