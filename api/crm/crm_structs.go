// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crm

const VersionNo = "20131207"


/*  */
type GradePromotionListObject struct {
	GradePromotion []*GradePromotion `json:"grade_promotion"`

}

/* 卖家设置的等级优惠信息 */
type GradePromotion struct {
	CurGrade string `json:"cur_grade"`
	CurGradeName string `json:"cur_grade_name"`
	Discount int `json:"discount"`
	NextGrade string `json:"next_grade"`
	NextGradeName string `json:"next_grade_name"`
	NextUpgradeAmount int `json:"next_upgrade_amount"`
	NextUpgradeCount int `json:"next_upgrade_count"`

}

/*  */
type GroupListObject struct {
	Group []*Group `json:"group"`

}

/* 描述分组的数据结构 */
type Group struct {
	GroupCreate string `json:"group_create"`
	GroupId int `json:"group_id"`
	GroupModify string `json:"group_modify"`
	GroupName string `json:"group_name"`
	MemberCount int `json:"member_count"`
	Status string `json:"status"`

}

/*  */
type BasicMemberListObject struct {
	BasicMember []*BasicMember `json:"basic_member"`

}

/* 表示会员关系的基本信息字段，用于会员列表的基本查询 */
type BasicMember struct {
	BizOrderId int `json:"biz_order_id"`
	BuyerId int `json:"buyer_id"`
	BuyerNick string `json:"buyer_nick"`
	CloseTradeAmount float64 `json:"close_trade_amount"`
	CloseTradeCount int `json:"close_trade_count"`
	Grade int `json:"grade"`
	GroupIds string `json:"group_ids"`
	ItemNum int `json:"item_num"`
	LastTradeTime string `json:"last_trade_time"`
	RelationSource int `json:"relation_source"`
	Status string `json:"status"`
	TradeAmount float64 `json:"trade_amount"`
	TradeCount int `json:"trade_count"`

}

/*  */
type CrmMemberListObject struct {
	CrmMember []*CrmMember `json:"crm_member"`

}

/* 会员信息对象 */
type CrmMember struct {
	AvgPrice float64 `json:"avg_price"`
	BizOrderId int `json:"biz_order_id"`
	BuyerId int `json:"buyer_id"`
	BuyerNick string `json:"buyer_nick"`
	City string `json:"city"`
	CloseTradeAmount float64 `json:"close_trade_amount"`
	CloseTradeCount int `json:"close_trade_count"`
	Grade int `json:"grade"`
	GroupIds string `json:"group_ids"`
	ItemCloseCount int `json:"item_close_count"`
	ItemNum int `json:"item_num"`
	LastTradeTime string `json:"last_trade_time"`
	Province int `json:"province"`
	RelationSource int `json:"relation_source"`
	Status string `json:"status"`
	TradeAmount float64 `json:"trade_amount"`
	TradeCount int `json:"trade_count"`

}

/*  */
type ActionInfoListObject struct {
	ActionInfo []*ActionInfo `json:"action_info"`

}

/* 动作action */
type ActionInfo struct {
	ActionCode string `json:"action_code"`
	Applier string `json:"applier"`
	Id int `json:"id"`
	Name string `json:"name"`
	SellerId int `json:"seller_id"`
	Status int `json:"status"`
	Verifier string `json:"verifier"`

}

/* 分页结果 */
type PageResult struct {
	CurrentPage int `json:"current_page"`
	PageSize int `json:"page_size"`
	TotalAmount int `json:"total_amount"`
	TotalPage int `json:"total_page"`

}

/*  */
type AttributeVOListObject struct {
	AttributeVO []*AttributeVO `json:"attribute_v_o"`

}

/* 属性（指标）信息 */
type AttributeVO struct {
	ClazzType int `json:"clazz_type"`
	Code string `json:"code"`
	Description string `json:"description"`
	DocumentId int `json:"document_id"`
	Id int `json:"id"`
	ParamKeys *ParamKeyVOListObject `json:"param_keys"`
	Title string `json:"title"`
	TopAccess int `json:"top_access"`
	TypeId int `json:"type_id"`
	Unit string `json:"unit"`

}

/*  */
type ParamKeyVOListObject struct {
	ParamKeyVO []*ParamKeyVO `json:"param_key_v_o"`

}

/* 属性必填参数的对象 */
type ParamKeyVO struct {
	ClazzType int `json:"clazz_type"`
	Code string `json:"code"`
	DocumentId int `json:"document_id"`
	Id int `json:"id"`
	Name string `json:"name"`
	Sort int `json:"sort"`
	Type int `json:"type"`
	Value string `json:"value"`

}

/*  */
type DocumentVOListObject struct {
	DocumentVO []*DocumentVO `json:"document_v_o"`

}

/* 档案对象 */
type DocumentVO struct {
	Code string `json:"code"`
	Dscription string `json:"dscription"`
	Id int `json:"id"`
	Name string `json:"name"`

}

/* 函数配置 */
type Function struct {
	Creator string `json:"creator"`
	Id int `json:"id"`
	Name string `json:"name"`
	ParseType int `json:"parse_type"`
	Rule string `json:"rule"`
	Status int `json:"status"`
	Strategy int `json:"strategy"`
	UserId int `json:"user_id"`

}

/*  */
type FunctionListObject struct {
	Function []*Function `json:"function"`

}

/*  */
type InnerLabelListObject struct {
	InnerLabel []*InnerLabel `json:"inner_label"`

}

/* 用于表示分组中的标签 */
type InnerLabel struct {
	LabelId int `json:"label_id"`
	Level int `json:"level"`

}

/*  */
type HanoiLabelGroupListObject struct {
	HanoiLabelGroup []*HanoiLabelGroup `json:"hanoi_label_group"`

}

/* 标签组 */
type HanoiLabelGroup struct {
	BizStatus int `json:"biz_status"`
	Code string `json:"code"`
	Description string `json:"description"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Id int `json:"id"`
	Name string `json:"name"`
	Open bool `json:"open"`
	Scene int `json:"scene"`
	Type int `json:"type"`

}

/*  */
type LabelListObject struct {
	Label []*Label `json:"label"`

}

/* 标签 */
type Label struct {
	BizStatus int `json:"biz_status"`
	Code string `json:"code"`
	Description string `json:"description"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Id int `json:"id"`
	Name string `json:"name"`
	Open bool `json:"open"`
	Scene int `json:"scene"`
	TemplateId int `json:"template_id"`

}

/*  */
type ParameterVOListObject struct {
	ParameterVO []*ParameterVO `json:"parameter_v_o"`

}

/* 标签的参数实例化 */
type ParameterVO struct {
	Name string `json:"name"`
	StringValue string `json:"string_value"`

}

/* 模板、标签、分组的操作结果 */
type TemplateResult struct {
	ErrorMessage string `json:"error_message"`
	IsSuccess bool `json:"is_success"`
	Total int `json:"total"`

}

/*  */
type RangeVOListObject struct {
	RangeVO []*RangeVO `json:"range_v_o"`

}

/* 值域对象 */
type RangeVO struct {
	Description string `json:"description"`
	DocumentId int `json:"document_id"`
	Id int `json:"id"`
	Key string `json:"key"`
	Value string `json:"value"`

}

/*  */
type TemplateListObject struct {
	Template []*Template `json:"template"`

}

/* 模板 */
type Template struct {
	BizStatus int `json:"biz_status"`
	Creater int `json:"creater"`
	Description string `json:"description"`
	Expression string `json:"expression"`
	GmtCreate string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Id int `json:"id"`
	Name string `json:"name"`
	TemplateCode string `json:"template_code"`

}

/*  */
type TypeVOListObject struct {
	TypeVO []*TypeVO `json:"type_v_o"`

}

/* 类型对象 */
type TypeVO struct {
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"`

}

/*  */
type GradeEquityListObject struct {
	GradeEquity []*GradeEquity `json:"grade_equity"`

}

/* tmall权益 */
type GradeEquity struct {
	CurGrade string `json:"cur_grade"`
	CurGradeName string `json:"cur_grade_name"`
	ExcludeArea string `json:"exclude_area"`
	Point int `json:"point"`
	Postage bool `json:"postage"`

}

