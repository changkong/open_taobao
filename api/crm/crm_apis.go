// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package crm

import (
	"github.com/yaofangou/open_taobao"
)





/* 卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息 */
type CrmGradeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *CrmGradeGetRequest) GetResponse(accessToken string) (*CrmGradeGetResponse, []byte, error) {
	var resp CrmGradeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grade.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGradeGetResponse struct {
	GradePromotions *GradePromotionListObject `json:"grade_promotions"`
}

type CrmGradeGetResponseResult struct {
	Response *CrmGradeGetResponse `json:"crm_grade_get_response"`
}





/* 设置等级信息，可以设置层级等级，也可以单独设置一个等级。出于安全原因，折扣现最低只能设置到700即7折。 */
type CrmGradeSetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br> 
amount参数的填写规范：升级到下一个级别的需要的交易额，单位为分,必须全部填写.例如10000,20000,30000，其中10000表示非会员升级到普通的所需的交易额，20000表示普通升级到高级所需的交易额，层级等级中最高等级的下一个等级默认为0。会员等级越高，所需交易额必须越高。 */
func (r *CrmGradeSetRequest) SetAmount(value string) {
	r.SetValue("amount", value)
}

/* 只对设置的层级等级有效，必须要在amount和count参数中选择一个<br> 
count参数的填写规范： 
升级到下一个级别的需要的交易量,必须全部填写. 以逗号分隔,例如100,200,300，其中100表示非会员升级到普通会员交易量。层级等级中最高等级的下一个等级的交易量默认为0。会员等级越高，交易量必须越高。 */
func (r *CrmGradeSetRequest) SetCount(value string) {
	r.SetValue("count", value)
}

/* 会员级别折扣率。会员等级越高，折扣必须越低。 
950即9.5折，888折即8.88折。出于安全原因，折扣现最低只能设置到700即7折。 */
func (r *CrmGradeSetRequest) SetDiscount(value string) {
	r.SetValue("discount", value)
}

/* 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP */
func (r *CrmGradeSetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 是否设置达到某一会员等级的交易量和交易额，必填。4个级别都需要设置，如入参为true,true,true,false时，表示设置达到高级会员、VIP会员的交易量或者交易额，不设置达到至尊会员的交易量和交易额 */
func (r *CrmGradeSetRequest) SetHierarchy(value string) {
	r.SetValue("hierarchy", value)
}


func (r *CrmGradeSetRequest) GetResponse(accessToken string) (*CrmGradeSetResponse, []byte, error) {
	var resp CrmGradeSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grade.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGradeSetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGradeSetResponseResult struct {
	Response *CrmGradeSetResponse `json:"crm_grade_set_response"`
}





/* 商家通过该接口吸纳线上店铺会员。 */
type CrmGrademktMemberAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 会员nick */
func (r *CrmGrademktMemberAddRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 系统属性，json格式 */
func (r *CrmGrademktMemberAddRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 会员属性-json format 
生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberAddRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberAddRequest) GetResponse(accessToken string) (*CrmGrademktMemberAddResponse, []byte, error) {
	var resp CrmGrademktMemberAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberAddResponse struct {
	Model bool `json:"model"`
}

type CrmGrademktMemberAddResponseResult struct {
	Response *CrmGrademktMemberAddResponse `json:"crm_grademkt_member_add_response"`
}





/* 创建商品等级营销明细 */
type CrmGrademktMemberDetailCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 扩展字段 */
func (r *CrmGrademktMemberDetailCreateRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailCreateRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberDetailCreateRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailCreateResponse, []byte, error) {
	var resp CrmGrademktMemberDetailCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailCreateResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberDetailCreateResponseResult struct {
	Response *CrmGrademktMemberDetailCreateResponse `json:"crm_grademkt_member_detail_create_response"`
}





/* 删除商品等级营销明细 */
type CrmGrademktMemberDetailDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 扩展字段 */
func (r *CrmGrademktMemberDetailDeleteRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailDeleteRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberDetailDeleteRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailDeleteResponse, []byte, error) {
	var resp CrmGrademktMemberDetailDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailDeleteResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberDetailDeleteResponseResult struct {
	Response *CrmGrademktMemberDetailDeleteResponse `json:"crm_grademkt_member_detail_delete_response"`
}





/* 商家通过该接口查询等级营销活动 */
type CrmGrademktMemberDetailQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 扩展字段 */
func (r *CrmGrademktMemberDetailQueryRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberDetailQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberDetailQueryRequest) GetResponse(accessToken string) (*CrmGrademktMemberDetailQueryResponse, []byte, error) {
	var resp CrmGrademktMemberDetailQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.detail.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberDetailQueryResponse struct {
	Model string `json:"model"`
}

type CrmGrademktMemberDetailQueryResponseResult struct {
	Response *CrmGrademktMemberDetailQueryResponse `json:"crm_grademkt_member_detail_query_response"`
}





/* 商家通过该接口设置等级活动 */
type CrmGrademktMemberGradeactivityInitRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 扩展字段 */
func (r *CrmGrademktMemberGradeactivityInitRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 活动名称，不传默认为“等级营销” */
func (r *CrmGrademktMemberGradeactivityInitRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberGradeactivityInitRequest) GetResponse(accessToken string) (*CrmGrademktMemberGradeactivityInitResponse, []byte, error) {
	var resp CrmGrademktMemberGradeactivityInitResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.gradeactivity.init", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberGradeactivityInitResponse struct {
	Module bool `json:"module"`
}

type CrmGrademktMemberGradeactivityInitResponseResult struct {
	Response *CrmGrademktMemberGradeactivityInitResponse `json:"crm_grademkt_member_gradeactivity_init_response"`
}





/* 商家通过该接口查询线上店铺会员。 */
type CrmGrademktMemberQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 会员nick */
func (r *CrmGrademktMemberQueryRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 系统属性，json格式 */
func (r *CrmGrademktMemberQueryRequest) SetFeather(value string) {
	r.SetValue("feather", value)
}

/* 会员属性-json format 
生成方法见http://open.taobao.com/doc/detail.htm?id=101281 */
func (r *CrmGrademktMemberQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *CrmGrademktMemberQueryRequest) GetResponse(accessToken string) (*CrmGrademktMemberQueryResponse, []byte, error) {
	var resp CrmGrademktMemberQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grademkt.member.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrademktMemberQueryResponse struct {
	Module string `json:"module"`
}

type CrmGrademktMemberQueryResponseResult struct {
	Response *CrmGrademktMemberQueryResponse `json:"crm_grademkt_member_query_response"`
}





/* 卖家创建一个新的分组，接口返回一个创建成功的分组的id */
type CrmGroupAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分组名称，每个卖家最多可以拥有100个分组 */
func (r *CrmGroupAddRequest) SetGroupName(value string) {
	r.SetValue("group_name", value)
}


func (r *CrmGroupAddRequest) GetResponse(accessToken string) (*CrmGroupAddResponse, []byte, error) {
	var resp CrmGroupAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupAddResponse struct {
	GroupId int `json:"group_id"`
	IsSuccess bool `json:"is_success"`
}

type CrmGroupAddResponseResult struct {
	Response *CrmGroupAddResponse `json:"crm_group_add_response"`
}





/* 将某分组下的所有会员添加到另一个分组,注：1.该操作为异步任务，建议先调用taobao.crm.grouptask.check 确保涉及分组上没有任务；2.若分组下某会员分组数超最大限额，则该会员不会被添加到新分组，同时不影响其余会员添加分组，接口调用依然返回成功。 */
type CrmGroupAppendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 添加的来源分组 */
func (r *CrmGroupAppendRequest) SetFromGroupId(value string) {
	r.SetValue("from_group_id", value)
}

/* 添加的目标分组 */
func (r *CrmGroupAppendRequest) SetToGroupId(value string) {
	r.SetValue("to_group_id", value)
}


func (r *CrmGroupAppendRequest) GetResponse(accessToken string) (*CrmGroupAppendResponse, []byte, error) {
	var resp CrmGroupAppendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.append", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupAppendResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupAppendResponseResult struct {
	Response *CrmGroupAppendResponse `json:"crm_group_append_response"`
}





/* 将该分组下的所有会员移除出该组，同时删除该分组。注：删除分组为异步任务，必须先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type CrmGroupDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要删除的分组id */
func (r *CrmGroupDeleteRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}


func (r *CrmGroupDeleteRequest) GetResponse(accessToken string) (*CrmGroupDeleteResponse, []byte, error) {
	var resp CrmGroupDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupDeleteResponseResult struct {
	Response *CrmGroupDeleteResponse `json:"crm_group_delete_response"`
}





/* 将一个分组下的所有会员移动到另一个分组，会员从原分组中删除 
注：移动属性为异步任务建议先调用taobao.crm.grouptask.check 确保涉及属性上没有任务。 */
type CrmGroupMoveRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要移动的分组 */
func (r *CrmGroupMoveRequest) SetFromGroupId(value string) {
	r.SetValue("from_group_id", value)
}

/* 目的分组 */
func (r *CrmGroupMoveRequest) SetToGroupId(value string) {
	r.SetValue("to_group_id", value)
}


func (r *CrmGroupMoveRequest) GetResponse(accessToken string) (*CrmGroupMoveResponse, []byte, error) {
	var resp CrmGroupMoveResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.move", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupMoveResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupMoveResponseResult struct {
	Response *CrmGroupMoveResponse `json:"crm_group_move_response"`
}





/* 修改一个已经存在的分组，接口返回分组的修改是否成功 */
type CrmGroupUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分组的id */
func (r *CrmGroupUpdateRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 新的分组名，分组名称不能包含|或者： */
func (r *CrmGroupUpdateRequest) SetNewGroupName(value string) {
	r.SetValue("new_group_name", value)
}


func (r *CrmGroupUpdateRequest) GetResponse(accessToken string) (*CrmGroupUpdateResponse, []byte, error) {
	var resp CrmGroupUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.group.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmGroupUpdateResponseResult struct {
	Response *CrmGroupUpdateResponse `json:"crm_group_update_response"`
}





/* 查询卖家的分组，返回查询到的分组列表，分页返回分组 */
type CrmGroupsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 显示第几页的分组，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页码为1 */
func (r *CrmGroupsGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 每页显示的记录数，其最大值不能超过100条，最小值为1，默认20条 */
func (r *CrmGroupsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *CrmGroupsGetRequest) GetResponse(accessToken string) (*CrmGroupsGetResponse, []byte, error) {
	var resp CrmGroupsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.groups.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGroupsGetResponse struct {
	Groups *GroupListObject `json:"groups"`
	TotalResult int `json:"total_result"`
}

type CrmGroupsGetResponseResult struct {
	Response *CrmGroupsGetResponse `json:"crm_groups_get_response"`
}





/* 检查一个分组上是否有异步任务,异步任务包括1.将一个分组下的所有用户添加到另外一个分组2.将一个分组下的所有用户移动到另外一个分组3.删除某个分组 
若分组上有任务则该属性不能被操作。 */
type CrmGrouptaskCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分组id */
func (r *CrmGrouptaskCheckRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}


func (r *CrmGrouptaskCheckRequest) GetResponse(accessToken string) (*CrmGrouptaskCheckResponse, []byte, error) {
	var resp CrmGrouptaskCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.grouptask.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmGrouptaskCheckResponse struct {
	IsFinished bool `json:"is_finished"`
}

type CrmGrouptaskCheckResponseResult struct {
	Response *CrmGrouptaskCheckResponse `json:"crm_grouptask_check_response"`
}





/* 设置会员等级 */
type CrmMembergradeSetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *CrmMembergradeSetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 买家会员级别有四种1：普通会员。2：高级会员。 3VIP会员。 4：至尊VIP */
func (r *CrmMembergradeSetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}


func (r *CrmMembergradeSetRequest) GetResponse(accessToken string) (*CrmMembergradeSetResponse, []byte, error) {
	var resp CrmMembergradeSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.membergrade.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembergradeSetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembergradeSetResponseResult struct {
	Response *CrmMembergradeSetResponse `json:"crm_membergrade_set_response"`
}





/* 编辑会员的基本资料，接口返回会员信息修改是否成功 */
type CrmMemberinfoUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *CrmMemberinfoUpdateRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 城市 */
func (r *CrmMemberinfoUpdateRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 交易关闭金额，单位：分 */
func (r *CrmMemberinfoUpdateRequest) SetCloseTradeAmount(value string) {
	r.SetValue("close_trade_amount", value)
}

/* 交易关闭次数 */
func (r *CrmMemberinfoUpdateRequest) SetCloseTradeCount(value string) {
	r.SetValue("close_trade_count", value)
}

/* 会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip 
 
只有正常会员才给予升级，对于status 为delete或者blacklist的会员 升级无效 */
func (r *CrmMemberinfoUpdateRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 分组的id集合字符串 */
func (r *CrmMemberinfoUpdateRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}

/* 宝贝件数 */
func (r *CrmMemberinfoUpdateRequest) SetItemNum(value string) {
	r.SetValue("item_num", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置 */
func (r *CrmMemberinfoUpdateRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 用于描述会员的状态，normal表示正常，blacklist表示黑名单，delete表示删除会员(只有潜在未交易成功的会员才能删除) */
func (r *CrmMemberinfoUpdateRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易金额，单位：分 */
func (r *CrmMemberinfoUpdateRequest) SetTradeAmount(value string) {
	r.SetValue("trade_amount", value)
}

/* 交易笔数 */
func (r *CrmMemberinfoUpdateRequest) SetTradeCount(value string) {
	r.SetValue("trade_count", value)
}


func (r *CrmMemberinfoUpdateRequest) GetResponse(accessToken string) (*CrmMemberinfoUpdateResponse, []byte, error) {
	var resp CrmMemberinfoUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.memberinfo.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMemberinfoUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMemberinfoUpdateResponseResult struct {
	Response *CrmMemberinfoUpdateResponse `json:"crm_memberinfo_update_response"`
}





/* 查询卖家的会员，进行基本的查询，返回符合条件的会员列表 */
type CrmMembersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家的昵称 */
func (r *CrmMembersGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为1000 */
func (r *CrmMembersGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员。如果不传入值则默认为全部等级。 */
func (r *CrmMembersGetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 最迟上次交易时间 */
func (r *CrmMembersGetRequest) SetMaxLastTradeTime(value string) {
	r.SetValue("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (r *CrmMembersGetRequest) SetMaxTradeAmount(value string) {
	r.SetValue("max_trade_amount", value)
}

/* 最大交易量 */
func (r *CrmMembersGetRequest) SetMaxTradeCount(value string) {
	r.SetValue("max_trade_count", value)
}

/* 最早上次交易时间 */
func (r *CrmMembersGetRequest) SetMinLastTradeTime(value string) {
	r.SetValue("min_last_trade_time", value)
}

/* 最小交易额,单位为元 */
func (r *CrmMembersGetRequest) SetMinTradeAmount(value string) {
	r.SetValue("min_trade_amount", value)
}

/* 最小交易量 */
func (r *CrmMembersGetRequest) SetMinTradeCount(value string) {
	r.SetValue("min_trade_count", value)
}

/* 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1， */
func (r *CrmMembersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *CrmMembersGetRequest) GetResponse(accessToken string) (*CrmMembersGetResponse, []byte, error) {
	var resp CrmMembersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGetResponse struct {
	Members *BasicMemberListObject `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersGetResponseResult struct {
	Response *CrmMembersGetResponse `json:"crm_members_get_response"`
}





/* 为一批会员添加分组，接口返回添加是否成功,如至少有一个会员的分组添加成功，接口就返回成功，否则返回失败，如果当前会员已经拥有当前分组，则直接跳过 */
type CrmMembersGroupBatchaddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 会员的id（一次最多传入10个） */
func (r *CrmMembersGroupBatchaddRequest) SetBuyerIds(value string) {
	r.SetValue("buyer_ids", value)
}

/* 分组id */
func (r *CrmMembersGroupBatchaddRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}


func (r *CrmMembersGroupBatchaddRequest) GetResponse(accessToken string) (*CrmMembersGroupBatchaddResponse, []byte, error) {
	var resp CrmMembersGroupBatchaddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.group.batchadd", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGroupBatchaddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembersGroupBatchaddResponseResult struct {
	Response *CrmMembersGroupBatchaddResponse `json:"crm_members_group_batchadd_response"`
}





/* 批量删除多个会员的公共分组，接口返回删除是否成功，该接口只删除多个会员的公共分组，不是公共分组的，不进行删除。如果入参只输入一个会员，则表示删除该会员的某些分组。 */
type CrmMembersGroupsBatchdeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家的Id集合 */
func (r *CrmMembersGroupsBatchdeleteRequest) SetBuyerIds(value string) {
	r.SetValue("buyer_ids", value)
}

/* 会员需要删除的分组 */
func (r *CrmMembersGroupsBatchdeleteRequest) SetGroupIds(value string) {
	r.SetValue("group_ids", value)
}


func (r *CrmMembersGroupsBatchdeleteRequest) GetResponse(accessToken string) (*CrmMembersGroupsBatchdeleteResponse, []byte, error) {
	var resp CrmMembersGroupsBatchdeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.groups.batchdelete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersGroupsBatchdeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmMembersGroupsBatchdeleteResponseResult struct {
	Response *CrmMembersGroupsBatchdeleteResponse `json:"crm_members_groups_batchdelete_response"`
}





/* 增量获取会员列表，接口返回符合查询条件的所有会员。任何状态更改都会返回,最大允许100 */
type CrmMembersIncrementGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1 */
func (r *CrmMembersIncrementGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 对应买家 最后一次 修改交易订单的时间，如果不填写此字段，默认为当前时间 */
func (r *CrmMembersIncrementGetRequest) SetEndModify(value string) {
	r.SetValue("end_modify", value)
}

/* 会员等级，0：店铺客户，1：普通会员，2：高级会员，3：VIP会员， 4：至尊VIP会员 */
func (r *CrmMembersIncrementGetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 每页显示的会员数，page_size的值不能超过100，最小值要大于1 */
func (r *CrmMembersIncrementGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 对应买家 最后一次 修改交易订单的时间 */
func (r *CrmMembersIncrementGetRequest) SetStartModify(value string) {
	r.SetValue("start_modify", value)
}


func (r *CrmMembersIncrementGetRequest) GetResponse(accessToken string) (*CrmMembersIncrementGetResponse, []byte, error) {
	var resp CrmMembersIncrementGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.increment.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersIncrementGetResponse struct {
	Members *BasicMemberListObject `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersIncrementGetResponseResult struct {
	Response *CrmMembersIncrementGetResponse `json:"crm_members_increment_get_response"`
}





/* 会员列表的高级查询，接口返回符合条件的会员列表.<br> 
注：建议获取09年以后的数据，09年之前的数据不是很完整 */
type CrmMembersSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *CrmMembersSearchRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 城市 */
func (r *CrmMembersSearchRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1.最大1000页 */
func (r *CrmMembersSearchRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 会员等级，0：店铺客户，1：普通客户，2：高级会员，3：VIP会员, 4：至尊VIP会员 */
func (r *CrmMembersSearchRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 分组id */
func (r *CrmMembersSearchRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 最大平均客单价，单位为元 */
func (r *CrmMembersSearchRequest) SetMaxAvgPrice(value string) {
	r.SetValue("max_avg_price", value)
}

/* 最大交易关闭笔数 */
func (r *CrmMembersSearchRequest) SetMaxCloseTradeNum(value string) {
	r.SetValue("max_close_trade_num", value)
}

/* 最大交易宝贝件数 */
func (r *CrmMembersSearchRequest) SetMaxItemNum(value string) {
	r.SetValue("max_item_num", value)
}

/* 最迟上次交易时间 */
func (r *CrmMembersSearchRequest) SetMaxLastTradeTime(value string) {
	r.SetValue("max_last_trade_time", value)
}

/* 最大交易额，单位为元 */
func (r *CrmMembersSearchRequest) SetMaxTradeAmount(value string) {
	r.SetValue("max_trade_amount", value)
}

/* 最大交易量 */
func (r *CrmMembersSearchRequest) SetMaxTradeCount(value string) {
	r.SetValue("max_trade_count", value)
}

/* 最少平均客单价，单位为元 */
func (r *CrmMembersSearchRequest) SetMinAvgPrice(value string) {
	r.SetValue("min_avg_price", value)
}

/* 最小交易关闭的笔数 */
func (r *CrmMembersSearchRequest) SetMinCloseTradeNum(value string) {
	r.SetValue("min_close_trade_num", value)
}

/* 最小交易宝贝件数 */
func (r *CrmMembersSearchRequest) SetMinItemNum(value string) {
	r.SetValue("min_item_num", value)
}

/* 最早上次交易时间 */
func (r *CrmMembersSearchRequest) SetMinLastTradeTime(value string) {
	r.SetValue("min_last_trade_time", value)
}

/* 最小交易额，单位为元 */
func (r *CrmMembersSearchRequest) SetMinTradeAmount(value string) {
	r.SetValue("min_trade_amount", value)
}

/* 最小交易量 */
func (r *CrmMembersSearchRequest) SetMinTradeCount(value string) {
	r.SetValue("min_trade_count", value)
}

/* 每页显示的会员数量，page_size的最大值不能超过100，最小值不能小于1 */
func (r *CrmMembersSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35 */
func (r *CrmMembersSearchRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 关系来源，1交易成功，2未成交，3卖家手动吸纳 */
func (r *CrmMembersSearchRequest) SetRelationSource(value string) {
	r.SetValue("relation_source", value)
}


func (r *CrmMembersSearchRequest) GetResponse(accessToken string) (*CrmMembersSearchResponse, []byte, error) {
	var resp CrmMembersSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.members.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmMembersSearchResponse struct {
	Members *CrmMemberListObject `json:"members"`
	TotalResult int `json:"total_result"`
}

type CrmMembersSearchResponseResult struct {
	Response *CrmMembersSearchResponse `json:"crm_members_search_response"`
}





/* 此接口用于取消VIP优惠 */
type CrmShopvipCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *CrmShopvipCancelRequest) GetResponse(accessToken string) (*CrmShopvipCancelResponse, []byte, error) {
	var resp CrmShopvipCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.shopvip.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmShopvipCancelResponse struct {
	IsSuccess bool `json:"is_success"`
}

type CrmShopvipCancelResponseResult struct {
	Response *CrmShopvipCancelResponse `json:"crm_shopvip_cancel_response"`
}





/* 查询是否是指定类型的无线未下单会员 0：所有无线未购买 1：主客登录未购买 */
type CrmWirelessMemberGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家nick */
func (r *CrmWirelessMemberGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 0：所有无线未下单 1：主客登录未下单 */
func (r *CrmWirelessMemberGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *CrmWirelessMemberGetRequest) GetResponse(accessToken string) (*CrmWirelessMemberGetResponse, []byte, error) {
	var resp CrmWirelessMemberGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.crm.wireless.member.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CrmWirelessMemberGetResponse struct {
	IsMatch bool `json:"is_match"`
}

type CrmWirelessMemberGetResponseResult struct {
	Response *CrmWirelessMemberGetResponse `json:"crm_wireless_member_get_response"`
}





/* 根据条件查询action列表 */
type HanoiActionGetListRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分页查询当前查询的页数； */
func (r *HanoiActionGetListRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* action的状态（0：审核中 1：审核通过，正常） */
func (r *HanoiActionGetListRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *HanoiActionGetListRequest) GetResponse(accessToken string) (*HanoiActionGetListResponse, []byte, error) {
	var resp HanoiActionGetListResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.action.get.list", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiActionGetListResponse struct {
	Hanoiactions *ActionInfoListObject `json:"hanoiactions"`
	Total int `json:"total"`
}

type HanoiActionGetListResponseResult struct {
	Response *HanoiActionGetListResponse `json:"hanoi_action_get_list_response"`
}





/* 查询单个action */
type HanoiActionGetSingleRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* action的code */
func (r *HanoiActionGetSingleRequest) SetActionCode(value string) {
	r.SetValue("action_code", value)
}

/* action的id */
func (r *HanoiActionGetSingleRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* action的name */
func (r *HanoiActionGetSingleRequest) SetName(value string) {
	r.SetValue("name", value)
}


func (r *HanoiActionGetSingleRequest) GetResponse(accessToken string) (*HanoiActionGetSingleResponse, []byte, error) {
	var resp HanoiActionGetSingleResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.action.get.single", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiActionGetSingleResponse struct {
	Action string `json:"action"`
}

type HanoiActionGetSingleResponseResult struct {
	Response *HanoiActionGetSingleResponse `json:"hanoi_action_get_single_response"`
}





/* * 1、查询属性时，<b>只能查询Attribute的简单信息</b>，不会级联任何表结构从而得到其他对象，如果需要用到其他对象，请从其他接口获得 
	 * 2、属性的必填参数分两种，一种是随模板一起传入，这种必填参数创建模板时必须赋值Value */
type HanoiAttributesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性编码 */
func (r *HanoiAttributesGetRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 分页时需要用。默认第一页。 */
func (r *HanoiAttributesGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 属性的唯一标识 */
func (r *HanoiAttributesGetRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 是否支持模糊查询 */
func (r *HanoiAttributesGetRequest) SetIndistinctQuery(value string) {
	r.SetValue("indistinct_query", value)
}

/* 分页时 每页显示的条数。最小1 最大30 默认10页 */
func (r *HanoiAttributesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 属性名称 */
func (r *HanoiAttributesGetRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* Top 的访问权限。按位与。1：可读，2：可写，4：可规则计算。如可读且可写的权限值为3。 */
func (r *HanoiAttributesGetRequest) SetTopAccess(value string) {
	r.SetValue("top_access", value)
}

/* 类型 唯一标识 */
func (r *HanoiAttributesGetRequest) SetTypeId(value string) {
	r.SetValue("type_id", value)
}

/* 类型名称 */
func (r *HanoiAttributesGetRequest) SetTypeName(value string) {
	r.SetValue("type_name", value)
}


func (r *HanoiAttributesGetRequest) GetResponse(accessToken string) (*HanoiAttributesGetResponse, []byte, error) {
	var resp HanoiAttributesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.attributes.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiAttributesGetResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values *AttributeVOListObject `json:"values"`
}

type HanoiAttributesGetResponseResult struct {
	Response *HanoiAttributesGetResponse `json:"hanoi_attributes_get_response"`
}





/* 店铺会员数据读取。一次可读取一个用户的多个属性值，但所有属性值的主键必须一致。主键不一致的属性值，请分多次读取。 */
type HanoiDataserviceReadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiDataserviceReadRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要查询的属性的id(Long型)。以json的数组的形式。所有属性的querykey必须相同，如有不同的需要分多次查询。属性的taobao.hanoi.ranges.get不为空时，值为key的拼接 */
func (r *HanoiDataserviceReadRequest) SetAttrs(value string) {
	r.SetValue("attrs", value)
}

/* 查询的主键。如查询关系数据，主键是sellerId+buyerId，但sellerId是系统自动传过来的，所以只需要传"BUYERNICK"或"BUYERID"二选一。再例如查询类目数据，则需要传"CATEGORYID"。 */
func (r *HanoiDataserviceReadRequest) SetPk(value string) {
	r.SetValue("pk", value)
}


func (r *HanoiDataserviceReadRequest) GetResponse(accessToken string) (*HanoiDataserviceReadResponse, []byte, error) {
	var resp HanoiDataserviceReadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.dataservice.read", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDataserviceReadResponse struct {
	Values string `json:"values"`
}

type HanoiDataserviceReadResponseResult struct {
	Response *HanoiDataserviceReadResponse `json:"hanoi_dataservice_read_response"`
}





/* 数据实时回流 */
type HanoiDataserviceWriteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiDataserviceWriteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* json对象，key为属性ID，值为需要回写的内容,如果属性值域不为空（通过taobao.hanoi.ranges.get获取）则需要把值转成key回写，并以逗号","隔开拼接成字符串 */
func (r *HanoiDataserviceWriteRequest) SetData(value string) {
	r.SetValue("data", value)
}

/* 回流数据的主键或上下文环境，如卖家id，类目id等。 */
func (r *HanoiDataserviceWriteRequest) SetParams(value string) {
	r.SetValue("params", value)
}


func (r *HanoiDataserviceWriteRequest) GetResponse(accessToken string) (*HanoiDataserviceWriteResponse, []byte, error) {
	var resp HanoiDataserviceWriteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.dataservice.write", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDataserviceWriteResponse struct {
	Model bool `json:"model"`
}

type HanoiDataserviceWriteResponseResult struct {
	Response *HanoiDataserviceWriteResponse `json:"hanoi_dataservice_write_response"`
}





/* 增加一个数据模板.只增加基本信息，不会增加Detail详情 必填字段：name,owner,opened 
增加数据模板后，数据模板状态:未审核 <br/> 
	   
名词解释：<br/> 
<ul> 
<li> 
数据模板：将一系列的指标（AttributeVO）组合在一起，当做指标集合使用。数据模板包含多个数据模板详情，数据模板详情和指标一一对应。 
 数据模板主要应用在，标签匹配以及数据服务时使用。标明此次调用时，我需要哪些指标。</li> 
<li>数据模板详情：1个数据模板包含多个详情，一个详情和指标一一对应，用户对模板添加详情即指明此模板需要哪些指标。指标（AttributeVO）中如果包含参数（ParamKeyVO）时，用户可以选择性的添加参数的Value属性（AttributeVO通过其他服务接口获得，如果ParamKeyVO的value在创建DT时不赋值，后续调用其他服务接口时，也可以动态添加）</li> 
</ul> 
使用方法：<br/> 
{"name":"TESTS__","opened":0,"owner":"TEST"} */
type HanoiDatatemplateAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* appName */
func (r *HanoiDatatemplateAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* name:String类型，数据模板的名称 
opened:int类型，标识此数据模板是否对其他人可见 */
func (r *HanoiDatatemplateAddRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *HanoiDatatemplateAddRequest) GetResponse(accessToken string) (*HanoiDatatemplateAddResponse, []byte, error) {
	var resp HanoiDatatemplateAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateAddResponse struct {
	Value int `json:"value"`
}

type HanoiDatatemplateAddResponseResult struct {
	Response *HanoiDatatemplateAddResponse `json:"hanoi_datatemplate_add_response"`
}





/* ID必填， Opened必填。<br/> 
现阶段只根据ID删除,只能删除未审核的数据模板，数据模板审核后,不能删除。 
用法：<br/>  
{"id":888888,"opened":0} */
type HanoiDatatemplateDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* APPName */
func (r *HanoiDatatemplateDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* id：通过query服务接口得到ID */
func (r *HanoiDatatemplateDeleteRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *HanoiDatatemplateDeleteRequest) GetResponse(accessToken string) (*HanoiDatatemplateDeleteResponse, []byte, error) {
	var resp HanoiDatatemplateDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateDeleteResponse struct {
	Value bool `json:"value"`
}

type HanoiDatatemplateDeleteResponseResult struct {
	Response *HanoiDatatemplateDeleteResponse `json:"hanoi_datatemplate_delete_response"`
}





/* 向某个DT批量添加指标关系 注意 不会校验指标关系是否重复 返回的是添加数量 
必填字段： <br /> 
DataTemplateVO: <br /> 
id <br/> 
opened<br/> 
DataTemplateDetailVO:<br /> 
name--用户可以根据这个name来创建模板<br/> 
Attr--需要把查询出来的AttributeVO对象放入到Detail中,如果Attr对应有ParamKey，paramKey的value可填入也可不填入，填入后，查询出来的Detail中ParamMap就会包含此信息<br /> 
 
用法<br/> 
DataTemplateVO  {"id":1,"opened":1} <br/> 
DataTemplateDetailVO  [{"attr":{"code":"SEX","description":"按男、女、男+女、无四个维度","documentId":1219,"id":35666,"params":[{"documentId":null,,"id":7757,"key":"buyerID","name":"buyerID","required":0,"sort":0,"type":3,"value":123}],"title":"性别","typeId":23157,"unit":"","valueType":1},"name":"d0c24612-b0f5-4814-9cf8-8eccb20f2017"}] */
type HanoiDatatemplateDetailAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* appName */
func (r *HanoiDatatemplateDetailAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* attr: 将AttributeVO转换成JSON格式 
name: 详情的名称 */
func (r *HanoiDatatemplateDetailAddRequest) SetDataTemplateDetails(value string) {
	r.SetValue("data_template_details", value)
}

/* id:数据模板唯一标识 */
func (r *HanoiDatatemplateDetailAddRequest) SetDataTemplateVo(value string) {
	r.SetValue("data_template_vo", value)
}


func (r *HanoiDatatemplateDetailAddRequest) GetResponse(accessToken string) (*HanoiDatatemplateDetailAddResponse, []byte, error) {
	var resp HanoiDatatemplateDetailAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.detail.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateDetailAddResponse struct {
	Value int `json:"value"`
}

type HanoiDatatemplateDetailAddResponseResult struct {
	Response *HanoiDatatemplateDetailAddResponse `json:"hanoi_datatemplate_detail_add_response"`
}





/* 批量删除某个DT下面的关联指标关系。 Opened必填，dataTemplateId必填。 */
type HanoiDatatemplateDetailDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* appName */
func (r *HanoiDatatemplateDetailDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* Long类型数组的JSON格式 */
func (r *HanoiDatatemplateDetailDeleteRequest) SetDataTemplateDetailIds(value string) {
	r.SetValue("data_template_detail_ids", value)
}

/* id:数据模板唯一标识 */
func (r *HanoiDatatemplateDetailDeleteRequest) SetDataTemplateVo(value string) {
	r.SetValue("data_template_vo", value)
}


func (r *HanoiDatatemplateDetailDeleteRequest) GetResponse(accessToken string) (*HanoiDatatemplateDetailDeleteResponse, []byte, error) {
	var resp HanoiDatatemplateDetailDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.detail.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateDetailDeleteResponse struct {
	Value int `json:"value"`
}

type HanoiDatatemplateDetailDeleteResponseResult struct {
	Response *HanoiDatatemplateDetailDeleteResponse `json:"hanoi_datatemplate_detail_delete_response"`
}





/* 查询Detail详情, 根据DataTemplateDetailCriteriaVO信息查询<br/> 
justQueryParamNotInput 为true时， 则查询出来的Detail详情,只是参数没有填入的列表(（Detail中的Attribute中的ParamKey是调用其他匹配接口必填的值，按照其他接口的规范填入即可）)<br/> 
默认分页查询，每页10条记录。如果查询总条数，needRetPage = true<br/> 
DataTemplateDetailCriteriaVO 
<ul> 
<li>attrId(Long):AttributeVO的唯一标识</li> 
<li>templateId(Long 必填参数):数据模板的唯一标识</li> 
<li>name(String):数据模板详情的名称</li> 
<li>id(Long):根据模板唯一标识去查询</li> 
<li>pageSize:分页大小（最大值30）</li> 
<li>currentPage:当前页码</li> 
<li>needRetPage(Boolean 默认False):是否返回总条数</li> 
<li>justQueryParamNotInput（Boolean 默认False）:是否只查询每天如PK的详情列表</li> 
</ul> */
type HanoiDatatemplateDetailQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* appName */
func (r *HanoiDatatemplateDetailQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* attrId(Long):AttributeVO的唯一标识<br/> 
templateId(Long):数据模板的唯一标识<br/> 
name(String):数据模板详情的名称<br/> 
id(Long):根据模板唯一标识去查询<br/> 
pageSize:分页大小（最大值30）<br/> 
currentPage:当前页码<br/> 
needRetPage(Boolean 默认False):是否返回总条数<br/> 
justQueryParamNotInput（Boolean 默认False）:是否只查询每天如PK的详情列表<br/> */
func (r *HanoiDatatemplateDetailQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *HanoiDatatemplateDetailQueryRequest) GetResponse(accessToken string) (*HanoiDatatemplateDetailQueryResponse, []byte, error) {
	var resp HanoiDatatemplateDetailQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.detail.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateDetailQueryResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values string `json:"values"`
}

type HanoiDatatemplateDetailQueryResponseResult struct {
	Response *HanoiDatatemplateDetailQueryResponse `json:"hanoi_datatemplate_detail_query_response"`
}





/* 查询指标模板 默认不返回详情（detail）列表。 如果想要返回detail，只查询正常的detail <br /> 
必填参数：opened参数必填（opened表示开放策略：0：只根据sellerId开放，1：全网开放 2：只根据ISV开放） <br /> 
选填参数：isNeedDetail(如果needDetail为true时会返回attr列表，效率低)<br /> 
此接口分页处理。默认分页 每页10条，返回总条数 请设置快关：needRetPage=true,参数为true是critera对象中totalAmount为符合条件的总条数,否则，不返回总条数。<br /> 
 
{"currentPage":1,"id":123,"isNeedDetail":false,"needRetPage":false,"opened":1,"pageSize":10,"templateName":"xxx"} */
type HanoiDatatemplateQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* AppName */
func (r *HanoiDatatemplateQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* templateName:String 根据模板的名字查找 
isNeedDetail：Boolean 是否返回模板详情 
opened:int 开放策略 
owner:创建者，填入appkey */
func (r *HanoiDatatemplateQueryRequest) SetParameter(value string) {
	r.SetValue("parameter", value)
}


func (r *HanoiDatatemplateQueryRequest) GetResponse(accessToken string) (*HanoiDatatemplateQueryResponse, []byte, error) {
	var resp HanoiDatatemplateQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.datatemplate.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDatatemplateQueryResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values string `json:"values"`
}

type HanoiDatatemplateQueryResponseResult struct {
	Response *HanoiDatatemplateQueryResponse `json:"hanoi_datatemplate_query_response"`
}





/* 这里只是简单的查询档案的信息，<b>不会返回档案下面的值域信息</b> <br> */
type HanoiDocumentsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 档案资源对象的编码信息 */
func (r *HanoiDocumentsGetRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 分页时需要用。默认第一页。 */
func (r *HanoiDocumentsGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* id:唯一标示 */
func (r *HanoiDocumentsGetRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 档案的名称 */
func (r *HanoiDocumentsGetRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 分页时 每页显示的条数。最小1 最大30 默认10页 */
func (r *HanoiDocumentsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *HanoiDocumentsGetRequest) GetResponse(accessToken string) (*HanoiDocumentsGetResponse, []byte, error) {
	var resp HanoiDocumentsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.documents.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiDocumentsGetResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values *DocumentVOListObject `json:"values"`
}

type HanoiDocumentsGetResponseResult struct {
	Response *HanoiDocumentsGetResponse `json:"hanoi_documents_get_response"`
}





/* 添加function */
type HanoiFunctionAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 函数配置名称 */
func (r *HanoiFunctionAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 函数规则类型 0：JSON 1：条件表达式 */
func (r *HanoiFunctionAddRequest) SetParseType(value string) {
	r.SetValue("parse_type", value)
}

/* 函数规则定义。支持JSON格式，条件表达式等等。groupId：要匹配人群的标签Id。actionId：需要执行动作的动作Id。filterType：匹配类型。0代表动态标签 1代表标签组 doAction:true代表执行动作 getData:true 或者false true代表要匹配结果 */
func (r *HanoiFunctionAddRequest) SetRule(value string) {
	r.SetValue("rule", value)
}

/* 规则开放策略 0：user_id私有 1：所有user_id可以使用 2：同一创建者下的user_id拥有 */
func (r *HanoiFunctionAddRequest) SetStrategy(value string) {
	r.SetValue("strategy", value)
}


func (r *HanoiFunctionAddRequest) GetResponse(accessToken string) (*HanoiFunctionAddResponse, []byte, error) {
	var resp HanoiFunctionAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionAddResponse struct {
	Id int `json:"id"`
}

type HanoiFunctionAddResponseResult struct {
	Response *HanoiFunctionAddResponse `json:"hanoi_function_add_response"`
}





/* 删除一个function */
type HanoiFunctionDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* id:函数配置Id strategy必须输入，需要判断权限 */
func (r *HanoiFunctionDeleteRequest) SetSdata(value string) {
	r.SetValue("sdata", value)
}


func (r *HanoiFunctionDeleteRequest) GetResponse(accessToken string) (*HanoiFunctionDeleteResponse, []byte, error) {
	var resp HanoiFunctionDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionDeleteResponse struct {
	Model bool `json:"model"`
}

type HanoiFunctionDeleteResponseResult struct {
	Response *HanoiFunctionDeleteResponse `json:"hanoi_function_delete_response"`
}





/* 调用function */
type HanoiFunctionExecuteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionExecuteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* Param的json格式，map中通常包括functionId,strategy 等所需业务参数 */
func (r *HanoiFunctionExecuteRequest) SetPara(value string) {
	r.SetValue("para", value)
}


func (r *HanoiFunctionExecuteRequest) GetResponse(accessToken string) (*HanoiFunctionExecuteResponse, []byte, error) {
	var resp HanoiFunctionExecuteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.execute", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionExecuteResponse struct {
	Model string `json:"model"`
}

type HanoiFunctionExecuteResponseResult struct {
	Response *HanoiFunctionExecuteResponse `json:"hanoi_function_execute_response"`
}





/* 根据function的id查询function */
type HanoiFunctionGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionGetRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* id:函数配置Id strategy必须输入，需要判断权限 */
func (r *HanoiFunctionGetRequest) SetSdata(value string) {
	r.SetValue("sdata", value)
}


func (r *HanoiFunctionGetRequest) GetResponse(accessToken string) (*HanoiFunctionGetResponse, []byte, error) {
	var resp HanoiFunctionGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionGetResponse struct {
	Function *Function `json:"function"`
}

type HanoiFunctionGetResponseResult struct {
	Response *HanoiFunctionGetResponse `json:"hanoi_function_get_response"`
}





/* 根据条件检索function。current_page:最大100页 page_size:10-30 */
type HanoiFunctionSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionSearchRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* FunctionQuery的json格式 */
func (r *HanoiFunctionSearchRequest) SetSdata(value string) {
	r.SetValue("sdata", value)
}


func (r *HanoiFunctionSearchRequest) GetResponse(accessToken string) (*HanoiFunctionSearchResponse, []byte, error) {
	var resp HanoiFunctionSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionSearchResponse struct {
	Functions *FunctionListObject `json:"functions"`
	Total int `json:"total"`
}

type HanoiFunctionSearchResponseResult struct {
	Response *HanoiFunctionSearchResponse `json:"hanoi_function_search_response"`
}





/* 更新一个function */
type HanoiFunctionUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分配给调用方的名称信息，内部统计使用 */
func (r *HanoiFunctionUpdateRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 函数配置ID */
func (r *HanoiFunctionUpdateRequest) SetFunctionId(value string) {
	r.SetValue("function_id", value)
}

/* 函数配置名称 */
func (r *HanoiFunctionUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 函数规则类型 0：JSON 1：条件表达式 */
func (r *HanoiFunctionUpdateRequest) SetParseType(value string) {
	r.SetValue("parse_type", value)
}

/* 函数规则定义。支持JSON格式，条件表达式等等。groupId：要匹配人群的标签Id。actionId：需要执行动作的动作Id。filterType：匹配类型。0代表动态标签 1代表标签组 .doAction:true代表执行动作 getData:true 或者false true代表要匹配结果 */
func (r *HanoiFunctionUpdateRequest) SetRule(value string) {
	r.SetValue("rule", value)
}

/* 规则开放策略 0：user_id私有 1：所有user_id可以使用 2：同一创建者下的user_id拥有 */
func (r *HanoiFunctionUpdateRequest) SetStrategy(value string) {
	r.SetValue("strategy", value)
}


func (r *HanoiFunctionUpdateRequest) GetResponse(accessToken string) (*HanoiFunctionUpdateResponse, []byte, error) {
	var resp HanoiFunctionUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.function.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiFunctionUpdateResponse struct {
	Model bool `json:"model"`
}

type HanoiFunctionUpdateResponseResult struct {
	Response *HanoiFunctionUpdateResponse `json:"hanoi_function_update_response"`
}





/* 添加一个分组 */
type HanoiGroupAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的简单描述 */
func (r *HanoiGroupAddRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 分组的编码 */
func (r *HanoiGroupAddRequest) SetGroupCode(value string) {
	r.SetValue("group_code", value)
}

/* 分组的名称 */
func (r *HanoiGroupAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 true开放，默认为false */
func (r *HanoiGroupAddRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 使用场景 */
func (r *HanoiGroupAddRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 分组的类型，0 互斥，1 共存，默认为0 */
func (r *HanoiGroupAddRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *HanoiGroupAddRequest) GetResponse(accessToken string) (*HanoiGroupAddResponse, []byte, error) {
	var resp HanoiGroupAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupAddResponse struct {
	Value int `json:"value"`
}

type HanoiGroupAddResponseResult struct {
	Response *HanoiGroupAddResponse `json:"hanoi_group_add_response"`
}





/* 创建一个标签分组，自动审核 */
type HanoiGroupAddPapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupAddPapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的简单描述 */
func (r *HanoiGroupAddPapRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 分组的编码 */
func (r *HanoiGroupAddPapRequest) SetGroupCode(value string) {
	r.SetValue("group_code", value)
}

/* 分组的名称 */
func (r *HanoiGroupAddPapRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 true开放，默认为false */
func (r *HanoiGroupAddPapRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 使用场景 */
func (r *HanoiGroupAddPapRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 分组的类型，0 互斥，1 共存，默认为0 */
func (r *HanoiGroupAddPapRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *HanoiGroupAddPapRequest) GetResponse(accessToken string) (*HanoiGroupAddPapResponse, []byte, error) {
	var resp HanoiGroupAddPapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.add.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupAddPapResponse struct {
	Value int `json:"value"`
}

type HanoiGroupAddPapResponseResult struct {
	Response *HanoiGroupAddPapResponse `json:"hanoi_group_add_pap_response"`
}





/* 删除一个标签分组 */
type HanoiGroupDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要删除的分组id */
func (r *HanoiGroupDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *HanoiGroupDeleteRequest) GetResponse(accessToken string) (*HanoiGroupDeleteResponse, []byte, error) {
	var resp HanoiGroupDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupDeleteResponse struct {
	Value int `json:"value"`
}

type HanoiGroupDeleteResponseResult struct {
	Response *HanoiGroupDeleteResponse `json:"hanoi_group_delete_response"`
}





/* 分组匹配 */
type HanoiGroupFilterRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupFilterRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 买家的nick */
func (r *HanoiGroupFilterRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 分组的id */
func (r *HanoiGroupFilterRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 卖家nick */
func (r *HanoiGroupFilterRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}


func (r *HanoiGroupFilterRequest) GetResponse(accessToken string) (*HanoiGroupFilterResponse, []byte, error) {
	var resp HanoiGroupFilterResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.filter", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupFilterResponse struct {
	LabelList []int `json:"label_list"`
}

type HanoiGroupFilterResponseResult struct {
	Response *HanoiGroupFilterResponse `json:"hanoi_group_filter_response"`
}





/* 给一个分组添加标签 */
type HanoiGroupLabelAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupLabelAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的id */
func (r *HanoiGroupLabelAddRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 标签的id */
func (r *HanoiGroupLabelAddRequest) SetLabelId(value string) {
	r.SetValue("label_id", value)
}

/* 标签的优先级。对于互斥分组必须填 */
func (r *HanoiGroupLabelAddRequest) SetLevel(value string) {
	r.SetValue("level", value)
}


func (r *HanoiGroupLabelAddRequest) GetResponse(accessToken string) (*HanoiGroupLabelAddResponse, []byte, error) {
	var resp HanoiGroupLabelAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.label.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupLabelAddResponse struct {
	Value int `json:"value"`
}

type HanoiGroupLabelAddResponseResult struct {
	Response *HanoiGroupLabelAddResponse `json:"hanoi_group_label_add_response"`
}





/* 给标签分组添加一个标签，自动审核 */
type HanoiGroupLabelAddPapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupLabelAddPapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的id */
func (r *HanoiGroupLabelAddPapRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 标签的id */
func (r *HanoiGroupLabelAddPapRequest) SetLabelId(value string) {
	r.SetValue("label_id", value)
}

/* 标签的优先级。对于互斥分组必须填 */
func (r *HanoiGroupLabelAddPapRequest) SetLevel(value string) {
	r.SetValue("level", value)
}


func (r *HanoiGroupLabelAddPapRequest) GetResponse(accessToken string) (*HanoiGroupLabelAddPapResponse, []byte, error) {
	var resp HanoiGroupLabelAddPapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.label.add.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupLabelAddPapResponse struct {
	Value int `json:"value"`
}

type HanoiGroupLabelAddPapResponseResult struct {
	Response *HanoiGroupLabelAddPapResponse `json:"hanoi_group_label_add_pap_response"`
}





/* 删除标签组中的某一个标签 */
type HanoiGroupLabelDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupLabelDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的id */
func (r *HanoiGroupLabelDeleteRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 标签的id */
func (r *HanoiGroupLabelDeleteRequest) SetLabelId(value string) {
	r.SetValue("label_id", value)
}


func (r *HanoiGroupLabelDeleteRequest) GetResponse(accessToken string) (*HanoiGroupLabelDeleteResponse, []byte, error) {
	var resp HanoiGroupLabelDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.label.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupLabelDeleteResponse struct {
	Value int `json:"value"`
}

type HanoiGroupLabelDeleteResponseResult struct {
	Response *HanoiGroupLabelDeleteResponse `json:"hanoi_group_label_delete_response"`
}





/* 删除分组中的一个标签，自动审核 */
type HanoiGroupLabelDeletePapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupLabelDeletePapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的id */
func (r *HanoiGroupLabelDeletePapRequest) SetGroupId(value string) {
	r.SetValue("group_id", value)
}

/* 标签的id */
func (r *HanoiGroupLabelDeletePapRequest) SetLabelId(value string) {
	r.SetValue("label_id", value)
}


func (r *HanoiGroupLabelDeletePapRequest) GetResponse(accessToken string) (*HanoiGroupLabelDeletePapResponse, []byte, error) {
	var resp HanoiGroupLabelDeletePapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.label.delete.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupLabelDeletePapResponse struct {
	Value int `json:"value"`
}

type HanoiGroupLabelDeletePapResponseResult struct {
	Response *HanoiGroupLabelDeletePapResponse `json:"hanoi_group_label_delete_pap_response"`
}





/* 查询分组使用的标签及优先级 */
type HanoiGroupLabelQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupLabelQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要查询的分组id */
func (r *HanoiGroupLabelQueryRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *HanoiGroupLabelQueryRequest) GetResponse(accessToken string) (*HanoiGroupLabelQueryResponse, []byte, error) {
	var resp HanoiGroupLabelQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.label.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupLabelQueryResponse struct {
	Values *InnerLabelListObject `json:"values"`
}

type HanoiGroupLabelQueryResponseResult struct {
	Response *HanoiGroupLabelQueryResponse `json:"hanoi_group_label_query_response"`
}





/* 根据条件检索分组 */
type HanoiGroupListQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupListQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 模板业务状态 1 审核 0 创建 */
func (r *HanoiGroupListQueryRequest) SetBizStatus(value string) {
	r.SetValue("biz_status", value)
}

/* 当前页、默认为1 */
func (r *HanoiGroupListQueryRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 根据时间查询 创建时间结束点 */
func (r *HanoiGroupListQueryRequest) SetGmtCreateEnd(value string) {
	r.SetValue("gmt_create_end", value)
}

/* 根据时间查询 创建时间截止点 */
func (r *HanoiGroupListQueryRequest) SetGmtCreateStart(value string) {
	r.SetValue("gmt_create_start", value)
}

/* 根据时间查询，最近修改时间截止 */
func (r *HanoiGroupListQueryRequest) SetGmtModifiedEnd(value string) {
	r.SetValue("gmt_modified_end", value)
}

/* 根据时间查询，最近修改时间起点 */
func (r *HanoiGroupListQueryRequest) SetGmtModifiedStart(value string) {
	r.SetValue("gmt_modified_start", value)
}

/* 分组的编码 */
func (r *HanoiGroupListQueryRequest) SetGroupCode(value string) {
	r.SetValue("group_code", value)
}

/* 分组id */
func (r *HanoiGroupListQueryRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 是否查询总数，默认false */
func (r *HanoiGroupListQueryRequest) SetIsTotal(value string) {
	r.SetValue("is_total", value)
}

/* 分组名称 */
func (r *HanoiGroupListQueryRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略(true 开放) */
func (r *HanoiGroupListQueryRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 分页查询，默认20 */
func (r *HanoiGroupListQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 使用场景 */
func (r *HanoiGroupListQueryRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 分组类型 0互斥 1 共存 */
func (r *HanoiGroupListQueryRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *HanoiGroupListQueryRequest) GetResponse(accessToken string) (*HanoiGroupListQueryResponse, []byte, error) {
	var resp HanoiGroupListQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.list.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupListQueryResponse struct {
	Total int `json:"total"`
	Values *HanoiLabelGroupListObject `json:"values"`
}

type HanoiGroupListQueryResponseResult struct {
	Response *HanoiGroupListQueryResponse `json:"hanoi_group_list_query_response"`
}





/* 更新一个标签的分组 */
type HanoiGroupUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupUpdateRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的描述 */
func (r *HanoiGroupUpdateRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 分组上次更新时间 */
func (r *HanoiGroupUpdateRequest) SetGmtModified(value string) {
	r.SetValue("gmt_modified", value)
}

/* 分组的编码 */
func (r *HanoiGroupUpdateRequest) SetGroupCode(value string) {
	r.SetValue("group_code", value)
}

/* 需要更新的分组id */
func (r *HanoiGroupUpdateRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 分组的名称 */
func (r *HanoiGroupUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略（true 开放，默认为false） */
func (r *HanoiGroupUpdateRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 使用场景 */
func (r *HanoiGroupUpdateRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 分组的类型，0 互斥，1 共存，默认为0 */
func (r *HanoiGroupUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *HanoiGroupUpdateRequest) GetResponse(accessToken string) (*HanoiGroupUpdateResponse, []byte, error) {
	var resp HanoiGroupUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupUpdateResponse struct {
	Value int `json:"value"`
}

type HanoiGroupUpdateResponseResult struct {
	Response *HanoiGroupUpdateResponse `json:"hanoi_group_update_response"`
}





/* 更新一个标签分组，自动审核 */
type HanoiGroupUpdatePapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiGroupUpdatePapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 分组的描述 */
func (r *HanoiGroupUpdatePapRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 分组上次更新时间 */
func (r *HanoiGroupUpdatePapRequest) SetGmtModified(value string) {
	r.SetValue("gmt_modified", value)
}

/* 分组的编码 */
func (r *HanoiGroupUpdatePapRequest) SetGroupCode(value string) {
	r.SetValue("group_code", value)
}

/* 需要更新的分组id */
func (r *HanoiGroupUpdatePapRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 分组的名称 */
func (r *HanoiGroupUpdatePapRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略（true 开放，默认为false） */
func (r *HanoiGroupUpdatePapRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 使用场景 */
func (r *HanoiGroupUpdatePapRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 分组的类型，0 互斥，1 共存，默认为0 */
func (r *HanoiGroupUpdatePapRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *HanoiGroupUpdatePapRequest) GetResponse(accessToken string) (*HanoiGroupUpdatePapResponse, []byte, error) {
	var resp HanoiGroupUpdatePapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.group.update.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiGroupUpdatePapResponse struct {
	Value int `json:"value"`
}

type HanoiGroupUpdatePapResponseResult struct {
	Response *HanoiGroupUpdatePapResponse `json:"hanoi_group_update_pap_response"`
}





/* 用于创建一个新的标签 */
type HanoiLabelAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 针对标签的一个简单描述 */
func (r *HanoiLabelAddRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 标签的编码，用于检索 */
func (r *HanoiLabelAddRequest) SetLabelCode(value string) {
	r.SetValue("label_code", value)
}

/* 标签的名称 */
func (r *HanoiLabelAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 标签是否开放 */
func (r *HanoiLabelAddRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 针对模板表达式中需要设置的参数的一个实例化，List<ParameterVO>对象的json格式 */
func (r *HanoiLabelAddRequest) SetParas(value string) {
	r.SetValue("paras", value)
}

/* 标签的使用场景 */
func (r *HanoiLabelAddRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签对应的模板id */
func (r *HanoiLabelAddRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelAddRequest) GetResponse(accessToken string) (*HanoiLabelAddResponse, []byte, error) {
	var resp HanoiLabelAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelAddResponse struct {
	Value int `json:"value"`
}

type HanoiLabelAddResponseResult struct {
	Response *HanoiLabelAddResponse `json:"hanoi_label_add_response"`
}





/* 和创建标签类似，区别在于自动审核，可立即使用 */
type HanoiLabelAddPapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelAddPapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 针对标签的一个简单描述 */
func (r *HanoiLabelAddPapRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 标签的编码，用于检索 */
func (r *HanoiLabelAddPapRequest) SetLabelCode(value string) {
	r.SetValue("label_code", value)
}

/* 标签的名称 */
func (r *HanoiLabelAddPapRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 标签是否开放 */
func (r *HanoiLabelAddPapRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 针对模板表达式中需要设置的参数的一个实例化，List<ParameterVO>对象的json格式 */
func (r *HanoiLabelAddPapRequest) SetParas(value string) {
	r.SetValue("paras", value)
}

/* 标签的使用场景 */
func (r *HanoiLabelAddPapRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签对应的模板id */
func (r *HanoiLabelAddPapRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelAddPapRequest) GetResponse(accessToken string) (*HanoiLabelAddPapResponse, []byte, error) {
	var resp HanoiLabelAddPapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.add.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelAddPapResponse struct {
	Value int `json:"value"`
}

type HanoiLabelAddPapResponseResult struct {
	Response *HanoiLabelAddPapResponse `json:"hanoi_label_add_pap_response"`
}





/* 删除某个标签，前提是该标签没有被标签组引用 */
type HanoiLabelDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要删除的标签id */
func (r *HanoiLabelDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *HanoiLabelDeleteRequest) GetResponse(accessToken string) (*HanoiLabelDeleteResponse, []byte, error) {
	var resp HanoiLabelDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelDeleteResponse struct {
	Value int `json:"value"`
}

type HanoiLabelDeleteResponseResult struct {
	Response *HanoiLabelDeleteResponse `json:"hanoi_label_delete_response"`
}





/* 根据标签id，卖家和买家Nick，匹配标签 */
type HanoiLabelFilterRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelFilterRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 买家的nick */
func (r *HanoiLabelFilterRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 标签的id */
func (r *HanoiLabelFilterRequest) SetLabelId(value string) {
	r.SetValue("label_id", value)
}

/* 卖家nick */
func (r *HanoiLabelFilterRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}


func (r *HanoiLabelFilterRequest) GetResponse(accessToken string) (*HanoiLabelFilterResponse, []byte, error) {
	var resp HanoiLabelFilterResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.filter", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelFilterResponse struct {
	Filter bool `json:"filter"`
}

type HanoiLabelFilterResponseResult struct {
	Response *HanoiLabelFilterResponse `json:"hanoi_label_filter_response"`
}





/* 根据条件检索标签 */
type HanoiLabelListQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelListQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 模板业务状态 1 审核 0 创建 */
func (r *HanoiLabelListQueryRequest) SetBizStatus(value string) {
	r.SetValue("biz_status", value)
}

/* 模板的编码 */
func (r *HanoiLabelListQueryRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 当前页、默认为1 */
func (r *HanoiLabelListQueryRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 根据时间查询 创建时间结束点 */
func (r *HanoiLabelListQueryRequest) SetGmtCreateEnd(value string) {
	r.SetValue("gmt_create_end", value)
}

/* 根据时间查询 创建时间截止点 */
func (r *HanoiLabelListQueryRequest) SetGmtCreateStart(value string) {
	r.SetValue("gmt_create_start", value)
}

/* 根据时间查询，最近修改时间截止 */
func (r *HanoiLabelListQueryRequest) SetGmtModifiedEnd(value string) {
	r.SetValue("gmt_modified_end", value)
}

/* 根据时间查询，最近修改时间起点 */
func (r *HanoiLabelListQueryRequest) SetGmtModifiedStart(value string) {
	r.SetValue("gmt_modified_start", value)
}

/* 模板id */
func (r *HanoiLabelListQueryRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 是否查询总数，默认false */
func (r *HanoiLabelListQueryRequest) SetIsTotal(value string) {
	r.SetValue("is_total", value)
}

/* 模板名称 */
func (r *HanoiLabelListQueryRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 true 开放 */
func (r *HanoiLabelListQueryRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 分页查询，默认20 */
func (r *HanoiLabelListQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 添加使用场景字段 */
func (r *HanoiLabelListQueryRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签的源模板id */
func (r *HanoiLabelListQueryRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelListQueryRequest) GetResponse(accessToken string) (*HanoiLabelListQueryResponse, []byte, error) {
	var resp HanoiLabelListQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.list.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelListQueryResponse struct {
	Total int `json:"total"`
	Values *LabelListObject `json:"values"`
}

type HanoiLabelListQueryResponseResult struct {
	Response *HanoiLabelListQueryResponse `json:"hanoi_label_list_query_response"`
}





/* 查询指定标签的参数信息 */
type HanoiLabelParaQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelParaQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要删除的标签id */
func (r *HanoiLabelParaQueryRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *HanoiLabelParaQueryRequest) GetResponse(accessToken string) (*HanoiLabelParaQueryResponse, []byte, error) {
	var resp HanoiLabelParaQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.para.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelParaQueryResponse struct {
	Values *ParameterVOListObject `json:"values"`
}

type HanoiLabelParaQueryResponseResult struct {
	Response *HanoiLabelParaQueryResponse `json:"hanoi_label_para_query_response"`
}





/* 根据条件，检索标签数目 */
type HanoiLabelQueryCountRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelQueryCountRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 模板业务状态 1 审核 0 创建 */
func (r *HanoiLabelQueryCountRequest) SetBizStatus(value string) {
	r.SetValue("biz_status", value)
}

/* 模板的编码 */
func (r *HanoiLabelQueryCountRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 根据时间查询 创建时间结束点 */
func (r *HanoiLabelQueryCountRequest) SetGmtCreateEnd(value string) {
	r.SetValue("gmt_create_end", value)
}

/* 根据时间查询 创建时间截止点 */
func (r *HanoiLabelQueryCountRequest) SetGmtCreateStart(value string) {
	r.SetValue("gmt_create_start", value)
}

/* 根据时间查询，最近修改时间截止 */
func (r *HanoiLabelQueryCountRequest) SetGmtModifiedEnd(value string) {
	r.SetValue("gmt_modified_end", value)
}

/* 根据时间查询，最近修改时间起点 */
func (r *HanoiLabelQueryCountRequest) SetGmtModifiedStart(value string) {
	r.SetValue("gmt_modified_start", value)
}

/* 模板id */
func (r *HanoiLabelQueryCountRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 模板名称 */
func (r *HanoiLabelQueryCountRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 */
func (r *HanoiLabelQueryCountRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 使用场景 */
func (r *HanoiLabelQueryCountRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签的源模板id */
func (r *HanoiLabelQueryCountRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelQueryCountRequest) GetResponse(accessToken string) (*HanoiLabelQueryCountResponse, []byte, error) {
	var resp HanoiLabelQueryCountResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.query.count", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelQueryCountResponse struct {
	TemplateResult *TemplateResult `json:"template_result"`
	Value int `json:"value"`
}

type HanoiLabelQueryCountResponseResult struct {
	Response *HanoiLabelQueryCountResponse `json:"hanoi_label_query_count_response"`
}





/* 更新一个标签，前提是该标签未审核。 */
type HanoiLabelUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelUpdateRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 针对标签的一个简单描述 */
func (r *HanoiLabelUpdateRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 标签最近一次修改时间 */
func (r *HanoiLabelUpdateRequest) SetGmtModified(value string) {
	r.SetValue("gmt_modified", value)
}

/* 要修改的标签的id */
func (r *HanoiLabelUpdateRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 标签的编码，用于检索 */
func (r *HanoiLabelUpdateRequest) SetLabelCode(value string) {
	r.SetValue("label_code", value)
}

/* 标签的名称 */
func (r *HanoiLabelUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 true 开放 */
func (r *HanoiLabelUpdateRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 针对模板表达式中需要设置的参数的一个实例化，List<ParameterVO>对象的json格式 */
func (r *HanoiLabelUpdateRequest) SetParas(value string) {
	r.SetValue("paras", value)
}

/* 场景字段 */
func (r *HanoiLabelUpdateRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签对应的模板id。修改了模板，必须同时修改标签的参数paras */
func (r *HanoiLabelUpdateRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelUpdateRequest) GetResponse(accessToken string) (*HanoiLabelUpdateResponse, []byte, error) {
	var resp HanoiLabelUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelUpdateResponse struct {
	Value int `json:"value"`
}

type HanoiLabelUpdateResponseResult struct {
	Response *HanoiLabelUpdateResponse `json:"hanoi_label_update_response"`
}





/* 更新标签，标签自动审核，立即可用 */
type HanoiLabelUpdatePapRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiLabelUpdatePapRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 针对标签的一个简单描述 */
func (r *HanoiLabelUpdatePapRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 标签最近一次修改时间 */
func (r *HanoiLabelUpdatePapRequest) SetGmtModified(value string) {
	r.SetValue("gmt_modified", value)
}

/* 要修改的标签的id */
func (r *HanoiLabelUpdatePapRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 标签的编码，用于检索 */
func (r *HanoiLabelUpdatePapRequest) SetLabelCode(value string) {
	r.SetValue("label_code", value)
}

/* 标签的名称 */
func (r *HanoiLabelUpdatePapRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 开放策略 true 开放 */
func (r *HanoiLabelUpdatePapRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 针对模板表达式中需要设置的参数的一个实例化，List<ParameterVO>对象的json格式 */
func (r *HanoiLabelUpdatePapRequest) SetParas(value string) {
	r.SetValue("paras", value)
}

/* 场景字段 */
func (r *HanoiLabelUpdatePapRequest) SetScene(value string) {
	r.SetValue("scene", value)
}

/* 标签对应的模板id。修改了模板，必须同时修改标签的参数paras */
func (r *HanoiLabelUpdatePapRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *HanoiLabelUpdatePapRequest) GetResponse(accessToken string) (*HanoiLabelUpdatePapResponse, []byte, error) {
	var resp HanoiLabelUpdatePapResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.label.update.pap", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiLabelUpdatePapResponse struct {
	Value int `json:"value"`
}

type HanoiLabelUpdatePapResponseResult struct {
	Response *HanoiLabelUpdatePapResponse `json:"hanoi_label_update_pap_response"`
}





/* 通过代理方法获取属性数据，主要是结构化评价和dsr。 */
type HanoiProxyValueGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* app的密匙 */
func (r *HanoiProxyValueGetRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 结构化数据组装参数。 */
func (r *HanoiProxyValueGetRequest) SetData(value string) {
	r.SetValue("data", value)
}


func (r *HanoiProxyValueGetRequest) GetResponse(accessToken string) (*HanoiProxyValueGetResponse, []byte, error) {
	var resp HanoiProxyValueGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.proxy.value.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiProxyValueGetResponse struct {
	Model string `json:"model"`
}

type HanoiProxyValueGetResponseResult struct {
	Response *HanoiProxyValueGetResponse `json:"hanoi_proxy_value_get_response"`
}





/* 档案ID必填字段 
1、查询值域信息，根据值域ID查询，最多返回一条记录，不会返回分页信息。<br> 
	 * 2、根据档案ID查询，默认返回10条记录<br> 
	 * 3、值域Key只能配合档案ID来唯一确定一条记录，不会返回分页信息。<br> */
type HanoiRangesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分页时需要用。默认第一页。 */
func (r *HanoiRangesGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 根据档案ID查询下面的值域信息 */
func (r *HanoiRangesGetRequest) SetDocumentId(value string) {
	r.SetValue("document_id", value)
}

/* 属性的唯一标识 */
func (r *HanoiRangesGetRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 取值范围的Key。可以和Document_ID搭配使用。 */
func (r *HanoiRangesGetRequest) SetKey(value string) {
	r.SetValue("key", value)
}

/* 分页时 每页显示的条数。最小1 最大30 默认10页 */
func (r *HanoiRangesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *HanoiRangesGetRequest) GetResponse(accessToken string) (*HanoiRangesGetResponse, []byte, error) {
	var resp HanoiRangesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.ranges.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiRangesGetResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values *RangeVOListObject `json:"values"`
}

type HanoiRangesGetResponseResult struct {
	Response *HanoiRangesGetResponse `json:"hanoi_ranges_get_response"`
}





/* 用于创建模板 */
type HanoiTemplateAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiTemplateAddRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 所使用的数据模板 */
func (r *HanoiTemplateAddRequest) SetDataTemplateId(value string) {
	r.SetValue("data_template_id", value)
}

/* 模板的描述 */
func (r *HanoiTemplateAddRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 模板的表达式 */
func (r *HanoiTemplateAddRequest) SetExpression(value string) {
	r.SetValue("expression", value)
}

/* 模板的名称 */
func (r *HanoiTemplateAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 模板的开放策略，默认false不开放 */
func (r *HanoiTemplateAddRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 设置表达式中参数的类型。ParaVO对象的json格式 */
func (r *HanoiTemplateAddRequest) SetParaList(value string) {
	r.SetValue("para_list", value)
}

/* 模板的源模板id，用于模板分配，不填或者为0时，则视为创建新模板 */
func (r *HanoiTemplateAddRequest) SetSourceTemplateId(value string) {
	r.SetValue("source_template_id", value)
}

/* 模板的编码 */
func (r *HanoiTemplateAddRequest) SetTemplateCode(value string) {
	r.SetValue("template_code", value)
}


func (r *HanoiTemplateAddRequest) GetResponse(accessToken string) (*HanoiTemplateAddResponse, []byte, error) {
	var resp HanoiTemplateAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.template.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTemplateAddResponse struct {
	Value int `json:"value"`
}

type HanoiTemplateAddResponseResult struct {
	Response *HanoiTemplateAddResponse `json:"hanoi_template_add_response"`
}





/* 删除一个模板，只能删除未审核的 */
type HanoiTemplateDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiTemplateDeleteRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 要删除的模板id */
func (r *HanoiTemplateDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *HanoiTemplateDeleteRequest) GetResponse(accessToken string) (*HanoiTemplateDeleteResponse, []byte, error) {
	var resp HanoiTemplateDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.template.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTemplateDeleteResponse struct {
	Value int `json:"value"`
}

type HanoiTemplateDeleteResponseResult struct {
	Response *HanoiTemplateDeleteResponse `json:"hanoi_template_delete_response"`
}





/* 检索模板 */
type HanoiTemplateListQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiTemplateListQueryRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 模板业务状态 1 审核 0 创建 */
func (r *HanoiTemplateListQueryRequest) SetBizStatus(value string) {
	r.SetValue("biz_status", value)
}

/* 创建者id */
func (r *HanoiTemplateListQueryRequest) SetCreater(value string) {
	r.SetValue("creater", value)
}

/* 当前页、默认为1 */
func (r *HanoiTemplateListQueryRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 所使用的数据模板 */
func (r *HanoiTemplateListQueryRequest) SetDataTemplateId(value string) {
	r.SetValue("data_template_id", value)
}

/* 根据时间查询 创建时间结束点 */
func (r *HanoiTemplateListQueryRequest) SetGmtCreateEnd(value string) {
	r.SetValue("gmt_create_end", value)
}

/* 根据时间查询 创建时间截止点 */
func (r *HanoiTemplateListQueryRequest) SetGmtCreateStart(value string) {
	r.SetValue("gmt_create_start", value)
}

/* 根据时间查询，最近修改时间截止 */
func (r *HanoiTemplateListQueryRequest) SetGmtModifiedEnd(value string) {
	r.SetValue("gmt_modified_end", value)
}

/* 根据时间查询，最近修改时间起点 */
func (r *HanoiTemplateListQueryRequest) SetGmtModifiedStart(value string) {
	r.SetValue("gmt_modified_start", value)
}

/* 模板id */
func (r *HanoiTemplateListQueryRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 是否查询总数，默认false */
func (r *HanoiTemplateListQueryRequest) SetIsTotal(value string) {
	r.SetValue("is_total", value)
}

/* 模板名称 */
func (r *HanoiTemplateListQueryRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 模板的开放策略 */
func (r *HanoiTemplateListQueryRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 分页查询，默认20 */
func (r *HanoiTemplateListQueryRequest) SetPagesize(value string) {
	r.SetValue("pagesize", value)
}

/* 源模板id */
func (r *HanoiTemplateListQueryRequest) SetSourceTemplateId(value string) {
	r.SetValue("source_template_id", value)
}

/* 模板编码 */
func (r *HanoiTemplateListQueryRequest) SetTemplateCode(value string) {
	r.SetValue("template_code", value)
}


func (r *HanoiTemplateListQueryRequest) GetResponse(accessToken string) (*HanoiTemplateListQueryResponse, []byte, error) {
	var resp HanoiTemplateListQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.template.list.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTemplateListQueryResponse struct {
	Total int `json:"total"`
	Values *TemplateListObject `json:"values"`
}

type HanoiTemplateListQueryResponseResult struct {
	Response *HanoiTemplateListQueryResponse `json:"hanoi_template_list_query_response"`
}





/* 根据条件检索模板数目 */
type HanoiTemplateQueryCountRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiTemplateQueryCountRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 模板业务状态 1 审核 0 创建 */
func (r *HanoiTemplateQueryCountRequest) SetBizStatus(value string) {
	r.SetValue("biz_status", value)
}

/* 创建者id */
func (r *HanoiTemplateQueryCountRequest) SetCreater(value string) {
	r.SetValue("creater", value)
}

/* 所使用的数据模板 */
func (r *HanoiTemplateQueryCountRequest) SetDataTemplateId(value string) {
	r.SetValue("data_template_id", value)
}

/* 根据时间查询 创建时间结束点 */
func (r *HanoiTemplateQueryCountRequest) SetGmtCreateEnd(value string) {
	r.SetValue("gmt_create_end", value)
}

/* 根据时间查询 创建时间截止点 */
func (r *HanoiTemplateQueryCountRequest) SetGmtCreateStart(value string) {
	r.SetValue("gmt_create_start", value)
}

/* 根据时间查询，最近修改时间截止 */
func (r *HanoiTemplateQueryCountRequest) SetGmtModifiedEnd(value string) {
	r.SetValue("gmt_modified_end", value)
}

/* 根据时间查询，最近修改时间起点 */
func (r *HanoiTemplateQueryCountRequest) SetGmtModifiedStart(value string) {
	r.SetValue("gmt_modified_start", value)
}

/* 模板id */
func (r *HanoiTemplateQueryCountRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 模板名称 */
func (r *HanoiTemplateQueryCountRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 模板的开放策略 */
func (r *HanoiTemplateQueryCountRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 源模板id */
func (r *HanoiTemplateQueryCountRequest) SetSourceTemplateId(value string) {
	r.SetValue("source_template_id", value)
}

/* 模板编码 */
func (r *HanoiTemplateQueryCountRequest) SetTemplateCode(value string) {
	r.SetValue("template_code", value)
}


func (r *HanoiTemplateQueryCountRequest) GetResponse(accessToken string) (*HanoiTemplateQueryCountResponse, []byte, error) {
	var resp HanoiTemplateQueryCountResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.template.query.count", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTemplateQueryCountResponse struct {
	Value int `json:"value"`
}

type HanoiTemplateQueryCountResponseResult struct {
	Response *HanoiTemplateQueryCountResponse `json:"hanoi_template_query_count_response"`
}





/* 更新一个模板。包括模板的名称、编码、使用的属性和参数 */
type HanoiTemplateUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 认证信息 */
func (r *HanoiTemplateUpdateRequest) SetAppName(value string) {
	r.SetValue("app_name", value)
}

/* 所使用的数据模板 */
func (r *HanoiTemplateUpdateRequest) SetDataTemplateId(value string) {
	r.SetValue("data_template_id", value)
}

/* 模板的描述 */
func (r *HanoiTemplateUpdateRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 模板的表达式 */
func (r *HanoiTemplateUpdateRequest) SetExpression(value string) {
	r.SetValue("expression", value)
}

/* 模板上次更新时间 */
func (r *HanoiTemplateUpdateRequest) SetGmtModified(value string) {
	r.SetValue("gmt_modified", value)
}

/* 需要更新的模板id */
func (r *HanoiTemplateUpdateRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 模板的名称 */
func (r *HanoiTemplateUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 模板的开放策略，默认为false，不开放 */
func (r *HanoiTemplateUpdateRequest) SetOpen(value string) {
	r.SetValue("open", value)
}

/* 设置表达式中参数的类型。ParaVO对象的json格式 */
func (r *HanoiTemplateUpdateRequest) SetParaList(value string) {
	r.SetValue("para_list", value)
}

/* 模板的源模板id，0表示没有源模板 */
func (r *HanoiTemplateUpdateRequest) SetSourceTemplateId(value string) {
	r.SetValue("source_template_id", value)
}

/* 模板的编码 */
func (r *HanoiTemplateUpdateRequest) SetTemplateCode(value string) {
	r.SetValue("template_code", value)
}


func (r *HanoiTemplateUpdateRequest) GetResponse(accessToken string) (*HanoiTemplateUpdateResponse, []byte, error) {
	var resp HanoiTemplateUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.template.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTemplateUpdateResponse struct {
	Value int `json:"value"`
}

type HanoiTemplateUpdateResponseResult struct {
	Response *HanoiTemplateUpdateResponse `json:"hanoi_template_update_response"`
}





/* 只能根据ID或者名称查询类型 
 <b>不会返回类型下面关联的属性列表</b>，如果需要得到类型下面的属性列表，需要调用queryAttribute方法<br> */
type HanoiTypesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分页时需要用。默认第一页。 */
func (r *HanoiTypesGetRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 类型的唯一标识 */
func (r *HanoiTypesGetRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 类型的名称 */
func (r *HanoiTypesGetRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 分页时 每页显示的条数。最小1 最大30 默认10页 */
func (r *HanoiTypesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *HanoiTypesGetRequest) GetResponse(accessToken string) (*HanoiTypesGetResponse, []byte, error) {
	var resp HanoiTypesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.hanoi.types.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type HanoiTypesGetResponse struct {
	PageResult *PageResult `json:"page_result"`
	Values *TypeVOListObject `json:"values"`
}

type HanoiTypesGetResponseResult struct {
	Response *HanoiTypesGetResponse `json:"hanoi_types_get_response"`
}





/* 获取天猫卖家设置的等级权益 */
type TmallCrmEquityGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *TmallCrmEquityGetRequest) GetResponse(accessToken string) (*TmallCrmEquityGetResponse, []byte, error) {
	var resp TmallCrmEquityGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.crm.equity.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallCrmEquityGetResponse struct {
	GradeEquitys *GradeEquityListObject `json:"grade_equitys"`
}

type TmallCrmEquityGetResponseResult struct {
	Response *TmallCrmEquityGetResponse `json:"tmall_crm_equity_get_response"`
}





/* 天猫卖家设置权益 */
type TmallCrmEquitySetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 不免邮区域，只在包邮条件设置的时候生效。要和等级一一对应。包邮条件为false的时候不起作用。 */
func (r *TmallCrmEquitySetRequest) SetExcludeArea(value string) {
	r.SetValue("exclude_area", value)
}

/* 会员等级，用逗号分隔。买家会员级别0：店铺客户 1：普通会员 2 ：高级会员 3：VIP会员 4：至尊VIP */
func (r *TmallCrmEquitySetRequest) SetGrade(value string) {
	r.SetValue("grade", value)
}

/* 返几倍天猫积分，可以不设置。如果设置则要和等级一一对应。不设置代表清空。 */
func (r *TmallCrmEquitySetRequest) SetPoint(value string) {
	r.SetValue("point", value)
}

/* 是否包邮，可以不设置，如果设置则要和等级一一对应。不设置代表清空 */
func (r *TmallCrmEquitySetRequest) SetPostage(value string) {
	r.SetValue("postage", value)
}


func (r *TmallCrmEquitySetRequest) GetResponse(accessToken string) (*TmallCrmEquitySetResponse, []byte, error) {
	var resp TmallCrmEquitySetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.crm.equity.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallCrmEquitySetResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallCrmEquitySetResponseResult struct {
	Response *TmallCrmEquitySetResponse `json:"tmall_crm_equity_set_response"`
}



