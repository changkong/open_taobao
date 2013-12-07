// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package subuser

const VersionNo = "20131207"


/* 子账号角色 */
type Role struct {
	CreateTime string `json:"create_time"`
	Description string `json:"description"`
	ModifiedTime string `json:"modified_time"`
	Permissions *PermissionListObject `json:"permissions"`
	RoleId int `json:"role_id"`
	RoleName string `json:"role_name"`
	SellerId int `json:"seller_id"`

}

/*  */
type PermissionListObject struct {
	Permission []*Permission `json:"permission"`

}

/* 权限信息 */
type Permission struct {
	IsAuthorize int `json:"is_authorize"`
	ParentCode string `json:"parent_code"`
	PermissionCode string `json:"permission_code"`
	PermissionName string `json:"permission_name"`

}

/*  */
type SubUserInfoListObject struct {
	SubUserInfo []*SubUserInfo `json:"sub_user_info"`

}

/* 子账号基本信息 */
type SubUserInfo struct {
	FullName string `json:"full_name"`
	IsOnline int `json:"is_online"`
	Nick string `json:"nick"`
	SellerId int `json:"seller_id"`
	SellerNick string `json:"seller_nick"`
	Status int `json:"status"`
	SubId int `json:"sub_id"`

}

/*  */
type RoleListObject struct {
	Role []*Role `json:"role"`

}

/* 子账号所拥有的权限对象(直接赋予的权限和通过角色赋予的权限的总和对象) */
type SubUserPermission struct {
	Permissions *PermissionListObject `json:"permissions"`
	Roles *RoleListObject `json:"roles"`

}

/*  */
type DepartmentListObject struct {
	Department []*Department `json:"department"`

}

/* 部门信息 */
type Department struct {
	DepartmentId int `json:"department_id"`
	DepartmentName string `json:"department_name"`
	ParentId int `json:"parent_id"`

}

/*  */
type DutyListObject struct {
	Duty []*Duty `json:"duty"`

}

/* 子账号职务信息 */
type Duty struct {
	DutyId int `json:"duty_id"`
	DutyLevel int `json:"duty_level"`
	DutyName string `json:"duty_name"`

}

/* 子账号详细信息，其中包括账号基本信息、员工信息和部门职务信息 */
type SubUserFullInfo struct {
	DepartmentId int `json:"department_id"`
	DepartmentName string `json:"department_name"`
	DutyId int `json:"duty_id"`
	DutyLevel int `json:"duty_level"`
	DutyName string `json:"duty_name"`
	EmployeeId int `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	EmployeeNickname string `json:"employee_nickname"`
	EmployeeNum string `json:"employee_num"`
	EntryDate string `json:"entry_date"`
	LeaderId int `json:"leader_id"`
	OfficePhone string `json:"office_phone"`
	ParentDepartment int `json:"parent_department"`
	Sex int `json:"sex"`
	SubDispatchStatus bool `json:"sub_dispatch_status"`
	SubId int `json:"sub_id"`
	SubNick string `json:"sub_nick"`
	SubOwedStatus bool `json:"sub_owed_status"`
	SubStatus int `json:"sub_status"`
	SubuserEmail string `json:"subuser_email"`
	UserEmail string `json:"user_email"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`
	WorkLocation string `json:"work_location"`

}

/*  */
type SubAccountInfoListObject struct {
	SubAccountInfo []*SubAccountInfo `json:"sub_account_info"`

}

/* 子账号基本信息 */
type SubAccountInfo struct {
	SubDispatchStatus bool `json:"sub_dispatch_status"`
	SubId int `json:"sub_id"`
	SubNick string `json:"sub_nick"`
	SubOwedStatus bool `json:"sub_owed_status"`
	SubStatus int `json:"sub_status"`
	UserId int `json:"user_id"`
	UserNick string `json:"user_nick"`

}

