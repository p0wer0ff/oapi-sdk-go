// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf           *config.Config
	AdminDeptStats *AdminDeptStatService
	AdminUserStats *AdminUserStatService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.AdminDeptStats = newAdminDeptStatService(s)
	s.AdminUserStats = newAdminUserStatService(s)
	return s
}

type AdminDeptStatService struct {
	service *Service
}

func newAdminDeptStatService(service *Service) *AdminDeptStatService {
	return &AdminDeptStatService{
		service: service,
	}
}

type AdminUserStatService struct {
	service *Service
}

func newAdminUserStatService(service *Service) *AdminUserStatService {
	return &AdminUserStatService{
		service: service,
	}
}

type AdminDeptStatListReqCall struct {
	ctx            *core.Context
	adminDeptStats *AdminDeptStatService
	queryParams    map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AdminDeptStatListReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}
func (rc *AdminDeptStatListReqCall) SetStartDate(startDate string) {
	rc.queryParams["start_date"] = startDate
}
func (rc *AdminDeptStatListReqCall) SetEndDate(endDate string) {
	rc.queryParams["end_date"] = endDate
}
func (rc *AdminDeptStatListReqCall) SetDepartmentId(departmentId string) {
	rc.queryParams["department_id"] = departmentId
}
func (rc *AdminDeptStatListReqCall) SetContainsChildDept(containsChildDept bool) {
	rc.queryParams["contains_child_dept"] = containsChildDept
}
func (rc *AdminDeptStatListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *AdminDeptStatListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}

func (rc *AdminDeptStatListReqCall) Do() (*AdminDeptStatListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AdminDeptStatListResult{}
	req := request.NewRequest("admin/v1/admin_dept_stats", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.adminDeptStats.service.conf, req)
	return result, err
}

func (adminDeptStats *AdminDeptStatService) List(ctx *core.Context, optFns ...request.OptFn) *AdminDeptStatListReqCall {
	return &AdminDeptStatListReqCall{
		ctx:            ctx,
		adminDeptStats: adminDeptStats,
		queryParams:    map[string]interface{}{},
		optFns:         optFns,
	}
}

type AdminUserStatListReqCall struct {
	ctx            *core.Context
	adminUserStats *AdminUserStatService
	queryParams    map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AdminUserStatListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *AdminUserStatListReqCall) SetDepartmentIdType(departmentIdType string) {
	rc.queryParams["department_id_type"] = departmentIdType
}
func (rc *AdminUserStatListReqCall) SetStartDate(startDate string) {
	rc.queryParams["start_date"] = startDate
}
func (rc *AdminUserStatListReqCall) SetEndDate(endDate string) {
	rc.queryParams["end_date"] = endDate
}
func (rc *AdminUserStatListReqCall) SetDepartmentId(departmentId string) {
	rc.queryParams["department_id"] = departmentId
}
func (rc *AdminUserStatListReqCall) SetUserId(userId string) {
	rc.queryParams["user_id"] = userId
}
func (rc *AdminUserStatListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *AdminUserStatListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}

func (rc *AdminUserStatListReqCall) Do() (*AdminUserStatListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AdminUserStatListResult{}
	req := request.NewRequest("admin/v1/admin_user_stats", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.adminUserStats.service.conf, req)
	return result, err
}

func (adminUserStats *AdminUserStatService) List(ctx *core.Context, optFns ...request.OptFn) *AdminUserStatListReqCall {
	return &AdminUserStatListReqCall{
		ctx:            ctx,
		adminUserStats: adminUserStats,
		queryParams:    map[string]interface{}{},
		optFns:         optFns,
	}
}