// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package subuser

import (
	"github.com/yaofangou/open_taobao"
)





/* 给指定的卖家创建新的子账号角色<br/> 
如果需要授权的权限点有下级权限点或上级权限点，把该权限点的父权限点和该权限点的所有子权限都一并做赋权操作，并递归处理<br/>例如：权限点列表如下<br/> 
code=sell 宝贝管理<br/> 
---------|code=sm 店铺管理<br/> 
---------|---------|code=sm-design 如店铺装修<br/> 
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/> 
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/> 
---------|---------|code=phone 手机淘宝店铺<br/> 
调用改接口给code=sm-design店铺装修赋权时，同时会将下列权限点都赋予默认角色<br/> 
code=sell 宝贝管理<br/> 
---------|code=sm 店铺管理<br/> 
---------|---------|code=sm-design 如店铺装修<br/> 
---------|---------|---------|code=sm-tbd-visit内店装修入口<br/> 
---------|---------|---------|code=sm-tbd-publish内店装修发布<br/> */
type SellercenterRoleAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 角色描述 */
func (r *SellercenterRoleAddRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 角色名 */
func (r *SellercenterRoleAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 表示卖家昵称 */
func (r *SellercenterRoleAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 需要授权的权限点permission_code列表,以","分割.其code值可以通过调用taobao.sellercenter.user.permissions.get返回，其中permission.is_authorize=1的权限点可以通过本接口授权给对应角色。 */
func (r *SellercenterRoleAddRequest) SetPermissionCodes(value string) {
	r.SetValue("permission_codes", value)
}


func (r *SellercenterRoleAddRequest) GetResponse(accessToken string) (*SellercenterRoleAddResponse, []byte, error) {
	var resp SellercenterRoleAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.role.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterRoleAddResponse struct {
	Role *Role `json:"role"`
}

type SellercenterRoleAddResponseResult struct {
	Response *SellercenterRoleAddResponse `json:"sellercenter_role_add_response"`
}





/* 获取指定角色的信息。只能查询属于自己的角色信息 (主账号或者某个主账号的子账号登陆，只能查询属于该主账号的角色信息) */
type SellercenterRoleInfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 角色id */
func (r *SellercenterRoleInfoGetRequest) SetRoleId(value string) {
	r.SetValue("role_id", value)
}


func (r *SellercenterRoleInfoGetRequest) GetResponse(accessToken string) (*SellercenterRoleInfoGetResponse, []byte, error) {
	var resp SellercenterRoleInfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.role.info.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterRoleInfoGetResponse struct {
	Role *Role `json:"role"`
}

type SellercenterRoleInfoGetResponseResult struct {
	Response *SellercenterRoleInfoGetResponse `json:"sellercenter_role_info_get_response"`
}





/* 获取指定卖家的角色下属员工列表，只能获取属于登陆者自己的信息。 */
type SellercenterRolemembersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 角色id */
func (r *SellercenterRolemembersGetRequest) SetRoleId(value string) {
	r.SetValue("role_id", value)
}


func (r *SellercenterRolemembersGetRequest) GetResponse(accessToken string) (*SellercenterRolemembersGetResponse, []byte, error) {
	var resp SellercenterRolemembersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.rolemembers.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterRolemembersGetResponse struct {
	Subusers *SubUserInfoListObject `json:"subusers"`
}

type SellercenterRolemembersGetResponseResult struct {
	Response *SellercenterRolemembersGetResponse `json:"sellercenter_rolemembers_get_response"`
}





/* 获取指定卖家的角色列表，只能获取属于登陆者自己的信息。 */
type SellercenterRolesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家昵称(只允许查询自己的信息：当前登陆者) */
func (r *SellercenterRolesGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SellercenterRolesGetRequest) GetResponse(accessToken string) (*SellercenterRolesGetResponse, []byte, error) {
	var resp SellercenterRolesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.roles.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterRolesGetResponse struct {
	Roles *RoleListObject `json:"roles"`
}

type SellercenterRolesGetResponseResult struct {
	Response *SellercenterRolesGetResponse `json:"sellercenter_roles_get_response"`
}





/* 查询指定的子账号的被直接赋予的权限信息和角色信息。<br/>返回对象中包括直接赋予子账号的权限点信息、被赋予的角色以及角色的对应权限点信息。 */
type SellercenterSubuserPermissionsRolesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 子账号昵称(子账号标识) */
func (r *SellercenterSubuserPermissionsRolesGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SellercenterSubuserPermissionsRolesGetRequest) GetResponse(accessToken string) (*SellercenterSubuserPermissionsRolesGetResponse, []byte, error) {
	var resp SellercenterSubuserPermissionsRolesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.subuser.permissions.roles.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterSubuserPermissionsRolesGetResponse struct {
	SubuserPermission *SubUserPermission `json:"subuser_permission"`
}

type SellercenterSubuserPermissionsRolesGetResponseResult struct {
	Response *SellercenterSubuserPermissionsRolesGetResponse `json:"sellercenter_subuser_permissions_roles_get_response"`
}





/* 根据主账号nick查询该账号下所有的子账号列表，只能查询属于自己的账号信息 (主账号以及所属子账号) */
type SellercenterSubusersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 表示卖家昵称 */
func (r *SellercenterSubusersGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SellercenterSubusersGetRequest) GetResponse(accessToken string) (*SellercenterSubusersGetResponse, []byte, error) {
	var resp SellercenterSubusersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.subusers.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterSubusersGetResponse struct {
	Subusers *SubUserInfoListObject `json:"subusers"`
}

type SellercenterSubusersGetResponseResult struct {
	Response *SellercenterSubusersGetResponse `json:"sellercenter_subusers_get_response"`
}





/* 获取指定用户的权限集合，并不组装成树。如果是主账号，返回所有的权限列表；如果是子账号，返回所有已授权的权限。只能查询属于自己的账号信息 (如果是主账号，则是主账号以及所属子账号，如果是子账号则是对应主账号以及所属子账号) */
type SellercenterUserPermissionsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 用户标识，次入参必须为子账号比如zhangsan:cool。如果只输入主账号zhangsan，将报错。 */
func (r *SellercenterUserPermissionsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SellercenterUserPermissionsGetRequest) GetResponse(accessToken string) (*SellercenterUserPermissionsGetResponse, []byte, error) {
	var resp SellercenterUserPermissionsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.sellercenter.user.permissions.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SellercenterUserPermissionsGetResponse struct {
	Permissions *PermissionListObject `json:"permissions"`
}

type SellercenterUserPermissionsGetResponseResult struct {
	Response *SellercenterUserPermissionsGetResponse `json:"sellercenter_user_permissions_get_response"`
}





/* 根据主账号ID以及部门名称、父部门ID创建新的部门信息（通过主账号登陆只能新建属于该主账号的部门信息） */
type SubuserDepartmentAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 部门名称 */
func (r *SubuserDepartmentAddRequest) SetDepartmentName(value string) {
	r.SetValue("department_name", value)
}

/* 父部门ID 如果是最高部门则传入0 */
func (r *SubuserDepartmentAddRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}

/* 主账号用户名 */
func (r *SubuserDepartmentAddRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDepartmentAddRequest) GetResponse(accessToken string) (*SubuserDepartmentAddResponse, []byte, error) {
	var resp SubuserDepartmentAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.department.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDepartmentAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDepartmentAddResponseResult struct {
	Response *SubuserDepartmentAddResponse `json:"subuser_department_add_response"`
}





/* 根据主账号Nick和部门ID删除该主账号下的该部门信息，如果待删除部门下有子账号则无法删除该部门（通过主账号登陆只能删除属于该主账号下的部门信息） */
type SubuserDepartmentDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 部门ID */
func (r *SubuserDepartmentDeleteRequest) SetDepartmentId(value string) {
	r.SetValue("department_id", value)
}

/* 主账号用户名 */
func (r *SubuserDepartmentDeleteRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDepartmentDeleteRequest) GetResponse(accessToken string) (*SubuserDepartmentDeleteResponse, []byte, error) {
	var resp SubuserDepartmentDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.department.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDepartmentDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDepartmentDeleteResponseResult struct {
	Response *SubuserDepartmentDeleteResponse `json:"subuser_department_delete_response"`
}





/* 通过指定主账号ID和部门ID来修改该部门的名称或该部门的父部门ID。（通过主账号登陆只能修改属于该主账号的部门信息） */
type SubuserDepartmentUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 部门ID */
func (r *SubuserDepartmentUpdateRequest) SetDepartmentId(value string) {
	r.SetValue("department_id", value)
}

/* 部门名称 */
func (r *SubuserDepartmentUpdateRequest) SetDepartmentName(value string) {
	r.SetValue("department_name", value)
}

/* 父部门ID 如果是最高部门则传入0 */
func (r *SubuserDepartmentUpdateRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}

/* 主账号用户名 */
func (r *SubuserDepartmentUpdateRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDepartmentUpdateRequest) GetResponse(accessToken string) (*SubuserDepartmentUpdateResponse, []byte, error) {
	var resp SubuserDepartmentUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.department.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDepartmentUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDepartmentUpdateResponseResult struct {
	Response *SubuserDepartmentUpdateResponse `json:"subuser_department_update_response"`
}





/* 获取指定账户的所有部门列表，其实包括有每个部门的ID、父部门ID、部门名称（通过主账号登陆只能查询属于该主账号下的所有部门信息）。 */
type SubuserDepartmentsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主账号用户名 */
func (r *SubuserDepartmentsGetRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDepartmentsGetRequest) GetResponse(accessToken string) (*SubuserDepartmentsGetResponse, []byte, error) {
	var resp SubuserDepartmentsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.departments.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDepartmentsGetResponse struct {
	Departments *DepartmentListObject `json:"departments"`
}

type SubuserDepartmentsGetResponseResult struct {
	Response *SubuserDepartmentsGetResponse `json:"subuser_departments_get_response"`
}





/* 通过主账号ID来增加指定账户的职务信息，职务信息中包括职务名称以及职务级别（通过主账号登陆只能新建属于该主账号的职务信息） */
type SubuserDutyAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 职务级别 */
func (r *SubuserDutyAddRequest) SetDutyLevel(value string) {
	r.SetValue("duty_level", value)
}

/* 职务名称 */
func (r *SubuserDutyAddRequest) SetDutyName(value string) {
	r.SetValue("duty_name", value)
}

/* 主账号用户名 */
func (r *SubuserDutyAddRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDutyAddRequest) GetResponse(accessToken string) (*SubuserDutyAddResponse, []byte, error) {
	var resp SubuserDutyAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.duty.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDutyAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDutyAddResponseResult struct {
	Response *SubuserDutyAddResponse `json:"subuser_duty_add_response"`
}





/* 通过主账号ID和职务ID来删除主账号下的该职务信息。使用此职务的员工中职务一栏将变为空（通过主账号登陆只能删除属于该主账号下的职务信息） */
type SubuserDutyDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 职务ID */
func (r *SubuserDutyDeleteRequest) SetDutyId(value string) {
	r.SetValue("duty_id", value)
}

/* 主账号用户名 */
func (r *SubuserDutyDeleteRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDutyDeleteRequest) GetResponse(accessToken string) (*SubuserDutyDeleteResponse, []byte, error) {
	var resp SubuserDutyDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.duty.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDutyDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDutyDeleteResponseResult struct {
	Response *SubuserDutyDeleteResponse `json:"subuser_duty_delete_response"`
}





/* 通过主账号ID和职务ID来修改该职务信息中的职务名称及职务级别信息（通过主账号登陆只能修改属于该主账号下的职务信息） */
type SubuserDutyUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 职务ID */
func (r *SubuserDutyUpdateRequest) SetDutyId(value string) {
	r.SetValue("duty_id", value)
}

/* 职务级别 */
func (r *SubuserDutyUpdateRequest) SetDutyLevel(value string) {
	r.SetValue("duty_level", value)
}

/* 职务名称 */
func (r *SubuserDutyUpdateRequest) SetDutyName(value string) {
	r.SetValue("duty_name", value)
}

/* 主账号用户名 */
func (r *SubuserDutyUpdateRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDutyUpdateRequest) GetResponse(accessToken string) (*SubuserDutyUpdateResponse, []byte, error) {
	var resp SubuserDutyUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.duty.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDutyUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserDutyUpdateResponseResult struct {
	Response *SubuserDutyUpdateResponse `json:"subuser_duty_update_response"`
}





/* 通过主账号Nick获取该账户下的所有职务信息，职务信息中包括职务ID、职务名称以及职务等级（通过主账号登陆只能获取属于该主账号下的职务信息） */
type SubuserDutysGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主账号用户名 */
func (r *SubuserDutysGetRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubuserDutysGetRequest) GetResponse(accessToken string) (*SubuserDutysGetResponse, []byte, error) {
	var resp SubuserDutysGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.dutys.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserDutysGetResponse struct {
	Dutys *DutyListObject `json:"dutys"`
}

type SubuserDutysGetResponseResult struct {
	Response *SubuserDutysGetResponse `json:"subuser_dutys_get_response"`
}





/* 给指定子账号新增一个员工信息（通过主账号登陆只能新建属于该主账号的员工信息） */
type SubuserEmployeeAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 当前员工所属部门ID */
func (r *SubuserEmployeeAddRequest) SetDepartmentId(value string) {
	r.SetValue("department_id", value)
}

/* 当前员工担任职务ID */
func (r *SubuserEmployeeAddRequest) SetDutyId(value string) {
	r.SetValue("duty_id", value)
}

/* 员工姓名 */
func (r *SubuserEmployeeAddRequest) SetEmployeeName(value string) {
	r.SetValue("employee_name", value)
}

/* 员工花名 */
func (r *SubuserEmployeeAddRequest) SetEmployeeNickname(value string) {
	r.SetValue("employee_nickname", value)
}

/* 员工工号(可以用大小写英文字母和数字) */
func (r *SubuserEmployeeAddRequest) SetEmployeeNum(value string) {
	r.SetValue("employee_num", value)
}

/* 员工入职时间 */
func (r *SubuserEmployeeAddRequest) SetEntryDate(value string) {
	r.SetValue("entry_date", value)
}

/* 员工身份证号码 */
func (r *SubuserEmployeeAddRequest) SetIdCardNum(value string) {
	r.SetValue("id_card_num", value)
}

/* 直接上级的员工ID */
func (r *SubuserEmployeeAddRequest) SetLeaderId(value string) {
	r.SetValue("leader_id", value)
}

/* 办公电话 */
func (r *SubuserEmployeeAddRequest) SetOfficePhone(value string) {
	r.SetValue("office_phone", value)
}

/* 员工私人邮箱 */
func (r *SubuserEmployeeAddRequest) SetPersonalEmail(value string) {
	r.SetValue("personal_email", value)
}

/* 员工手机号码 */
func (r *SubuserEmployeeAddRequest) SetPersonalMobile(value string) {
	r.SetValue("personal_mobile", value)
}

/* 员工性别 1：男; 2:女 */
func (r *SubuserEmployeeAddRequest) SetSex(value string) {
	r.SetValue("sex", value)
}

/* 子账号ID */
func (r *SubuserEmployeeAddRequest) SetSubId(value string) {
	r.SetValue("sub_id", value)
}

/* 工作地点 */
func (r *SubuserEmployeeAddRequest) SetWorkLocation(value string) {
	r.SetValue("work_location", value)
}


func (r *SubuserEmployeeAddRequest) GetResponse(accessToken string) (*SubuserEmployeeAddResponse, []byte, error) {
	var resp SubuserEmployeeAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.employee.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserEmployeeAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserEmployeeAddResponseResult struct {
	Response *SubuserEmployeeAddResponse `json:"subuser_employee_add_response"`
}





/* 修改指定账户子账号的员工信息（通过主账号登陆只能修改属于该主账号下的员工信息） */
type SubuserEmployeeUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 当前员工所属部门ID */
func (r *SubuserEmployeeUpdateRequest) SetDepartmentId(value string) {
	r.SetValue("department_id", value)
}

/* 当前员工担任职务ID(若需要将该字段的值置为空，请传入-1） */
func (r *SubuserEmployeeUpdateRequest) SetDutyId(value string) {
	r.SetValue("duty_id", value)
}

/* 员工姓名 */
func (r *SubuserEmployeeUpdateRequest) SetEmployeeName(value string) {
	r.SetValue("employee_name", value)
}

/* 员工花名(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetEmployeeNickname(value string) {
	r.SetValue("employee_nickname", value)
}

/* 员工工号(卖家自定义)(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetEmployeeNum(value string) {
	r.SetValue("employee_num", value)
}

/* 登记员工离职  true:登记员工离职 */
func (r *SubuserEmployeeUpdateRequest) SetEmployeeTurnover(value string) {
	r.SetValue("employee_turnover", value)
}

/* 员工入职时间(若需要将该字段的值置为空，请传入1900-01-01 00:00:00） */
func (r *SubuserEmployeeUpdateRequest) SetEntryDate(value string) {
	r.SetValue("entry_date", value)
}

/* 员工身份证号码(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetIdCardNum(value string) {
	r.SetValue("id_card_num", value)
}

/* 直接上级的员工ID(若需要将该字段的值置为空，请传入-1） */
func (r *SubuserEmployeeUpdateRequest) SetLeaderId(value string) {
	r.SetValue("leader_id", value)
}

/* 办公电话(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetOfficePhone(value string) {
	r.SetValue("office_phone", value)
}

/* 员工私人邮箱(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetPersonalEmail(value string) {
	r.SetValue("personal_email", value)
}

/* 员工手机号码(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetPersonalMobile(value string) {
	r.SetValue("personal_mobile", value)
}

/* 员工性别  1：男;  2:女 */
func (r *SubuserEmployeeUpdateRequest) SetSex(value string) {
	r.SetValue("sex", value)
}

/* 子账号ID */
func (r *SubuserEmployeeUpdateRequest) SetSubId(value string) {
	r.SetValue("sub_id", value)
}

/* 杭州大厦(若需要将该字段的值置为空，请传入“-1”） */
func (r *SubuserEmployeeUpdateRequest) SetWorkLocation(value string) {
	r.SetValue("work_location", value)
}


func (r *SubuserEmployeeUpdateRequest) GetResponse(accessToken string) (*SubuserEmployeeUpdateResponse, []byte, error) {
	var resp SubuserEmployeeUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.employee.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserEmployeeUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserEmployeeUpdateResponseResult struct {
	Response *SubuserEmployeeUpdateResponse `json:"subuser_employee_update_response"`
}





/* 获取指定账户子账号的详细信息，其中包括子账号的账号信息以及员工、部门、职务信息（只能通过主账号登陆并查询属于该主账号下的某个子账号详细信息） */
type SubuserFullinfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email） */
func (r *SubuserFullinfoGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
func (r *SubuserFullinfoGetRequest) SetSubId(value string) {
	r.SetValue("sub_id", value)
}

/* 子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号） */
func (r *SubuserFullinfoGetRequest) SetSubNick(value string) {
	r.SetValue("sub_nick", value)
}


func (r *SubuserFullinfoGetRequest) GetResponse(accessToken string) (*SubuserFullinfoGetResponse, []byte, error) {
	var resp SubuserFullinfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.fullinfo.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserFullinfoGetResponse struct {
	SubFullinfo *SubUserFullInfo `json:"sub_fullinfo"`
}

type SubuserFullinfoGetResponseResult struct {
	Response *SubuserFullinfoGetResponse `json:"subuser_fullinfo_get_response"`
}





/* 修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息） */
type SubuserInfoUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否停用子账号 true:表示停用该子账号false:表示开启该子账号 */
func (r *SubuserInfoUpdateRequest) SetIsDisableSubaccount(value string) {
	r.SetValue("is_disable_subaccount", value)
}

/* 子账号是否参与分流 true:参与分流 false:不参与分流 */
func (r *SubuserInfoUpdateRequest) SetIsDispatch(value string) {
	r.SetValue("is_dispatch", value)
}

/* 子账号ID */
func (r *SubuserInfoUpdateRequest) SetSubId(value string) {
	r.SetValue("sub_id", value)
}


func (r *SubuserInfoUpdateRequest) GetResponse(accessToken string) (*SubuserInfoUpdateResponse, []byte, error) {
	var resp SubuserInfoUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subuser.info.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubuserInfoUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type SubuserInfoUpdateResponseResult struct {
	Response *SubuserInfoUpdateResponse `json:"subuser_info_update_response"`
}





/* 获取主账号下的子账号简易账号信息集合。（只能通过主账号登陆并且查询该属于主账号的子账号信息） */
type SubusersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主账号用户名 */
func (r *SubusersGetRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *SubusersGetRequest) GetResponse(accessToken string) (*SubusersGetResponse, []byte, error) {
	var resp SubusersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.subusers.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SubusersGetResponse struct {
	Subaccounts *SubAccountInfoListObject `json:"subaccounts"`
}

type SubusersGetResponseResult struct {
	Response *SubusersGetResponse `json:"subusers_get_response"`
}



