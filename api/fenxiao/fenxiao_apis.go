// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package fenxiao

import (
	"github.com/changkong/open_taobao"
)





/* 合作授权审批 */
type FenxiaoCooperationAuditRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 1:审批通过，审批通过后要加入授权产品线列表； 
2:审批拒绝 */
func (r *FenxiaoCooperationAuditRequest) SetAuditResult(value string) {
	r.SetValue("audit_result", value)
}

/* 当审批通过时需要指定授权产品线id列表(代销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
func (r *FenxiaoCooperationAuditRequest) SetProductLineListAgent(value string) {
	r.SetValue("product_line_list_agent", value)
}

/* 当审批通过时需要指定授权产品线id列表(经销)，如果没传则表示不开通，且经销和代销的授权产品线列表至少传入一种，同时传入则表示都开通。 */
func (r *FenxiaoCooperationAuditRequest) SetProductLineListDealer(value string) {
	r.SetValue("product_line_list_dealer", value)
}

/* 备注 */
func (r *FenxiaoCooperationAuditRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* 合作申请Id */
func (r *FenxiaoCooperationAuditRequest) SetRequisitionId(value string) {
	r.SetValue("requisition_id", value)
}


func (r *FenxiaoCooperationAuditRequest) GetResponse(accessToken string) (*FenxiaoCooperationAuditResponse, []byte, error) {
	var resp FenxiaoCooperationAuditResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.cooperation.audit", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoCooperationAuditResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoCooperationAuditResponseResult struct {
	Response *FenxiaoCooperationAuditResponse `json:"fenxiao_cooperation_audit_response"`
}





/* 获取供应商的合作关系信息 */
type FenxiaoCooperationGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 合作结束时间yyyy-MM-dd HH:mm:ss */
func (r *FenxiaoCooperationGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 页码（大于0的整数，默认1） */
func (r *FenxiaoCooperationGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (r *FenxiaoCooperationGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 合作开始时间yyyy-MM-dd HH:mm:ss */
func (r *FenxiaoCooperationGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}

/* 合作状态： NORMAL(合作中)、 ENDING(终止中) 、END (终止) */
func (r *FenxiaoCooperationGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 分销方式：AGENT(代销) 、DEALER（经销） */
func (r *FenxiaoCooperationGetRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoCooperationGetRequest) GetResponse(accessToken string) (*FenxiaoCooperationGetResponse, []byte, error) {
	var resp FenxiaoCooperationGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.cooperation.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoCooperationGetResponse struct {
	Cooperations []*Cooperation `json:"cooperations"`
	TotalResults int `json:"total_results"`
}

type FenxiaoCooperationGetResponseResult struct {
	Response *FenxiaoCooperationGetResponse `json:"fenxiao_cooperation_get_response"`
}





/* 追加授权产品线 */
type FenxiaoCooperationProductcatAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 合作关系id */
func (r *FenxiaoCooperationProductcatAddRequest) SetCooperateId(value string) {
	r.SetValue("cooperate_id", value)
}

/* 等级ID（为空则不修改） */
func (r *FenxiaoCooperationProductcatAddRequest) SetGradeId(value string) {
	r.SetValue("grade_id", value)
}

/* 产品线id列表，若有多个，以逗号分隔 */
func (r *FenxiaoCooperationProductcatAddRequest) SetProductLineList(value string) {
	r.SetValue("product_line_list", value)
}


func (r *FenxiaoCooperationProductcatAddRequest) GetResponse(accessToken string) (*FenxiaoCooperationProductcatAddResponse, []byte, error) {
	var resp FenxiaoCooperationProductcatAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.cooperation.productcat.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoCooperationProductcatAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoCooperationProductcatAddResponseResult struct {
	Response *FenxiaoCooperationProductcatAddResponse `json:"fenxiao_cooperation_productcat_add_response"`
}





/* 终止与分销商的合作关系 */
type FenxiaoCooperationTerminateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 合作编号 */
func (r *FenxiaoCooperationTerminateRequest) SetCooperateId(value string) {
	r.SetValue("cooperate_id", value)
}

/* 终止天数，可选1,2,3,5,7,15，在多少天数内终止 */
func (r *FenxiaoCooperationTerminateRequest) SetEndRemainDays(value string) {
	r.SetValue("end_remain_days", value)
}

/* 终止说明（5-2000字） */
func (r *FenxiaoCooperationTerminateRequest) SetEndRemark(value string) {
	r.SetValue("end_remark", value)
}


func (r *FenxiaoCooperationTerminateRequest) GetResponse(accessToken string) (*FenxiaoCooperationTerminateResponse, []byte, error) {
	var resp FenxiaoCooperationTerminateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.cooperation.terminate", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoCooperationTerminateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoCooperationTerminateResponseResult struct {
	Response *FenxiaoCooperationTerminateResponse `json:"fenxiao_cooperation_terminate_response"`
}





/* 供应商更新合作的分销商等级 */
type FenxiaoCooperationUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销商ID */
func (r *FenxiaoCooperationUpdateRequest) SetDistributorId(value string) {
	r.SetValue("distributor_id", value)
}

/* 等级ID(0代表取消) */
func (r *FenxiaoCooperationUpdateRequest) SetGradeId(value string) {
	r.SetValue("grade_id", value)
}

/* 分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销) */
func (r *FenxiaoCooperationUpdateRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoCooperationUpdateRequest) GetResponse(accessToken string) (*FenxiaoCooperationUpdateResponse, []byte, error) {
	var resp FenxiaoCooperationUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.cooperation.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoCooperationUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoCooperationUpdateResponseResult struct {
	Response *FenxiaoCooperationUpdateResponse `json:"fenxiao_cooperation_update_response"`
}





/* 供应商或分销商通过采购申请/经销采购单审核 */
type FenxiaoDealerRequisitionorderAgreeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 采购申请/经销采购单编号 */
func (r *FenxiaoDealerRequisitionorderAgreeRequest) SetDealerOrderId(value string) {
	r.SetValue("dealer_order_id", value)
}


func (r *FenxiaoDealerRequisitionorderAgreeRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderAgreeResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderAgreeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.agree", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderAgreeResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDealerRequisitionorderAgreeResponseResult struct {
	Response *FenxiaoDealerRequisitionorderAgreeResponse `json:"fenxiao_dealer_requisitionorder_agree_response"`
}





/* 供应商或分销商关闭采购申请/经销采购单 */
type FenxiaoDealerRequisitionorderCloseRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 经销采购单编号 */
func (r *FenxiaoDealerRequisitionorderCloseRequest) SetDealerOrderId(value string) {
	r.SetValue("dealer_order_id", value)
}

/* 关闭原因： 
1：长时间无法联系到分销商，取消交易。 
2：分销商错误提交申请，取消交易。 
3：缺货，近段时间都无法发货。 
4：分销商恶意提交申请单。 
5：其他原因。 */
func (r *FenxiaoDealerRequisitionorderCloseRequest) SetReason(value string) {
	r.SetValue("reason", value)
}

/* 关闭详细原因，字数5-200字 */
func (r *FenxiaoDealerRequisitionorderCloseRequest) SetReasonDetail(value string) {
	r.SetValue("reason_detail", value)
}


func (r *FenxiaoDealerRequisitionorderCloseRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderCloseResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderCloseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.close", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderCloseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDealerRequisitionorderCloseResponseResult struct {
	Response *FenxiaoDealerRequisitionorderCloseResponse `json:"fenxiao_dealer_requisitionorder_close_response"`
}





/* 批量查询采购申请/经销采购单，目前支持供应商和分销商查询 */
type FenxiaoDealerRequisitionorderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 采购申请/经销采购单最迟修改时间。与start_date字段的最大时间间隔不能超过30天 */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有采购申请/经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 查询者自己在所要查询的采购申请/经销采购单中的身份。 
1：供应商。 
2：分销商。 
注：填写其他值当做错误处理。 */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetIdentity(value string) {
	r.SetValue("identity", value)
}

/* 采购申请/经销采购单状态。 
0：全部状态。 
1：分销商提交申请，待供应商审核。 
2：供应商驳回申请，待分销商确认。 
3：供应商修改后，待分销商确认。 
4：分销商拒绝修改，待供应商再审核。 
5：审核通过下单成功，待分销商付款。 
7：付款成功，待供应商发货。 
8：供应商发货，待分销商收货。 
9：分销商收货，交易成功。 
10：采购申请/经销采购单关闭。 
 
注：无值按默认值0计，超出状态范围返回错误信息。 */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetOrderStatus(value string) {
	r.SetValue("order_status", value)
}

/* 页码（大于0的整数。无值或小于1的值按默认值1计） */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数（大于0但小于等于50的整数。无值或大于50或小于1的值按默认值50计） */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 采购申请/经销采购单最早修改时间 */
func (r *FenxiaoDealerRequisitionorderGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *FenxiaoDealerRequisitionorderGetRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderGetResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderGetResponse struct {
	DealerOrders []*DealerOrder `json:"dealer_orders"`
	TotalResults int `json:"total_results"`
}

type FenxiaoDealerRequisitionorderGetResponseResult struct {
	Response *FenxiaoDealerRequisitionorderGetResponse `json:"fenxiao_dealer_requisitionorder_get_response"`
}





/* 供应商或分销商修改采购申请/经销采购单明细，如果商品总数量被修改为0（或删除所有商品明细）则关闭采购申请/经销采购单，否则状态变为供应商或分销商审核通过。 */
type FenxiaoDealerRequisitionorderModifyRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 经销采购单编号 */
func (r *FenxiaoDealerRequisitionorderModifyRequest) SetDealerOrderId(value string) {
	r.SetValue("dealer_order_id", value)
}

/* 要删除的商品明细id列表，多个id使用英文符号的逗号隔开 */
func (r *FenxiaoDealerRequisitionorderModifyRequest) SetDeleteDetailIds(value string) {
	r.SetValue("delete_detail_ids", value)
}

/* 经销采购单的商品明细的新的采购价格。格式为商品明细id:价格修改值,商品明细id:价格修改值 */
func (r *FenxiaoDealerRequisitionorderModifyRequest) SetDetailIdPrices(value string) {
	r.SetValue("detail_id_prices", value)
}

/* 修改经销采购单的商品明细的新的库存。格式为商品明细id:库存修改值,商品明细id:库存修改值 */
func (r *FenxiaoDealerRequisitionorderModifyRequest) SetDetailIdQuantities(value string) {
	r.SetValue("detail_id_quantities", value)
}

/* 新邮费（单位：分，示例值1005表示10.05元）。必须大于等于0。自提方式不可修改邮费。不提交该参数表示不修改邮费。 */
func (r *FenxiaoDealerRequisitionorderModifyRequest) SetNewPostFee(value string) {
	r.SetValue("new_post_fee", value)
}


func (r *FenxiaoDealerRequisitionorderModifyRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderModifyResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderModifyResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.modify", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderModifyResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDealerRequisitionorderModifyResponseResult struct {
	Response *FenxiaoDealerRequisitionorderModifyResponse `json:"fenxiao_dealer_requisitionorder_modify_response"`
}





/* 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。 */
type FenxiaoDealerRequisitionorderQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 经销采购单编号。 
多个编号用英文符号的逗号隔开。最多支持50个经销采购单编号的查询。 */
func (r *FenxiaoDealerRequisitionorderQueryRequest) SetDealerOrderIds(value string) {
	r.SetValue("dealer_order_ids", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有经销采购单对象(dealer_orders)字段。 如果不为空：返回指定采购单对象(dealer_orders)字段。 例1： dealer_order_details.product_id 表示只返回product_id 例2： dealer_order_details 表示只返回明细列表 */
func (r *FenxiaoDealerRequisitionorderQueryRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *FenxiaoDealerRequisitionorderQueryRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderQueryResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderQueryResponse struct {
	DealerOrders []*DealerOrder `json:"dealer_orders"`
}

type FenxiaoDealerRequisitionorderQueryResponseResult struct {
	Response *FenxiaoDealerRequisitionorderQueryResponse `json:"fenxiao_dealer_requisitionorder_query_response"`
}





/* 供应商或分销商驳回采购申请/经销采购单 */
type FenxiaoDealerRequisitionorderRefuseRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 采购申请/经销采购单编号 */
func (r *FenxiaoDealerRequisitionorderRefuseRequest) SetDealerOrderId(value string) {
	r.SetValue("dealer_order_id", value)
}

/* 驳回原因（1：价格不合理；2：采购数量不合理；3：其他原因） */
func (r *FenxiaoDealerRequisitionorderRefuseRequest) SetReason(value string) {
	r.SetValue("reason", value)
}

/* 驳回详细原因，字数范围5-200字 */
func (r *FenxiaoDealerRequisitionorderRefuseRequest) SetReasonDetail(value string) {
	r.SetValue("reason_detail", value)
}


func (r *FenxiaoDealerRequisitionorderRefuseRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderRefuseResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderRefuseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.refuse", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderRefuseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDealerRequisitionorderRefuseResponseResult struct {
	Response *FenxiaoDealerRequisitionorderRefuseResponse `json:"fenxiao_dealer_requisitionorder_refuse_response"`
}





/* 供应商修改经销采购单备注 */
type FenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 经销采购单ID */
func (r *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetDealerOrderId(value string) {
	r.SetValue("dealer_order_id", value)
}

/* 备注留言，可为空 */
func (r *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemo(value string) {
	r.SetValue("supplier_memo", value)
}

/* 旗子的标记，必选。 
1-5之间的数字。 
非1-5之间，都采用1作为默认。 
1:红色 
2:黄色 
3:绿色 
4:蓝色 
5:粉红色 */
func (r *FenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemoFlag(value string) {
	r.SetValue("supplier_memo_flag", value)
}


func (r *FenxiaoDealerRequisitionorderRemarkUpdateRequest) GetResponse(accessToken string) (*FenxiaoDealerRequisitionorderRemarkUpdateResponse, []byte, error) {
	var resp FenxiaoDealerRequisitionorderRemarkUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.dealer.requisitionorder.remark.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDealerRequisitionorderRemarkUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDealerRequisitionorderRemarkUpdateResponseResult struct {
	Response *FenxiaoDealerRequisitionorderRemarkUpdateResponse `json:"fenxiao_dealer_requisitionorder_remark_update_response"`
}





/* 新增等级折扣 */
type FenxiaoDiscountAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 折扣名称,长度不能超过25字节 */
func (r *FenxiaoDiscountAddRequest) SetDiscountName(value string) {
	r.SetValue("discount_name", value)
}

/* PERCENT（按折扣优惠）、PRICE（按减价优惠），例如"PERCENT,PRICE,PERCENT" */
func (r *FenxiaoDiscountAddRequest) SetDiscountTypes(value string) {
	r.SetValue("discount_types", value)
}

/* 优惠比率或者优惠价格，例如：”8000,-2300,7000”,大小为-100000000到100000000之间（单位：分） */
func (r *FenxiaoDiscountAddRequest) SetDiscountValues(value string) {
	r.SetValue("discount_values", value)
}

/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (r *FenxiaoDiscountAddRequest) SetTargetIds(value string) {
	r.SetValue("target_ids", value)
}

/* GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠），例如"GRADE,DISTRIBUTOR" */
func (r *FenxiaoDiscountAddRequest) SetTargetTypes(value string) {
	r.SetValue("target_types", value)
}


func (r *FenxiaoDiscountAddRequest) GetResponse(accessToken string) (*FenxiaoDiscountAddResponse, []byte, error) {
	var resp FenxiaoDiscountAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.discount.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDiscountAddResponse struct {
	DiscountId int `json:"discount_id"`
}

type FenxiaoDiscountAddResponseResult struct {
	Response *FenxiaoDiscountAddResponse `json:"fenxiao_discount_add_response"`
}





/* 修改等级折扣 */
type FenxiaoDiscountUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 详情ID，例如：”0,1002,1003” */
func (r *FenxiaoDiscountUpdateRequest) SetDetailIds(value string) {
	r.SetValue("detail_ids", value)
}

/* ADD(新增)、UPDATE（更新）、DEL（删除，对应的target_type等信息填NULL），例如：”UPDATE,DEL,DEL” */
func (r *FenxiaoDiscountUpdateRequest) SetDetailStatuss(value string) {
	r.SetValue("detail_statuss", value)
}

/* 折扣ID */
func (r *FenxiaoDiscountUpdateRequest) SetDiscountId(value string) {
	r.SetValue("discount_id", value)
}

/* 折扣名称，长度不能超过25字节 */
func (r *FenxiaoDiscountUpdateRequest) SetDiscountName(value string) {
	r.SetValue("discount_name", value)
}

/* 状态DEL（删除）UPDATE(更新) */
func (r *FenxiaoDiscountUpdateRequest) SetDiscountStatus(value string) {
	r.SetValue("discount_status", value)
}

/* PERCENT（按折扣优惠）、PRICE（按减价优惠），例如"PERCENT,PRICE,PERCENT" */
func (r *FenxiaoDiscountUpdateRequest) SetDiscountTypes(value string) {
	r.SetValue("discount_types", value)
}

/* 优惠比率或者优惠价格，例如：”8000,-2300,7000”,大小为-100000000到100000000之间（单位：分） */
func (r *FenxiaoDiscountUpdateRequest) SetDiscountValues(value string) {
	r.SetValue("discount_values", value)
}

/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (r *FenxiaoDiscountUpdateRequest) SetTargetIds(value string) {
	r.SetValue("target_ids", value)
}

/* GRADE（按会员等级优惠）、DISTRIBUTOR（按分销商优惠），例如"GRADE,DISTRIBUTOR" */
func (r *FenxiaoDiscountUpdateRequest) SetTargetTypes(value string) {
	r.SetValue("target_types", value)
}


func (r *FenxiaoDiscountUpdateRequest) GetResponse(accessToken string) (*FenxiaoDiscountUpdateResponse, []byte, error) {
	var resp FenxiaoDiscountUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.discount.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDiscountUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDiscountUpdateResponseResult struct {
	Response *FenxiaoDiscountUpdateResponse `json:"fenxiao_discount_update_response"`
}





/* 查询折扣信息 */
type FenxiaoDiscountsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 折扣ID */
func (r *FenxiaoDiscountsGetRequest) SetDiscountId(value string) {
	r.SetValue("discount_id", value)
}

/* 指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用） */
func (r *FenxiaoDiscountsGetRequest) SetExtFields(value string) {
	r.SetValue("ext_fields", value)
}


func (r *FenxiaoDiscountsGetRequest) GetResponse(accessToken string) (*FenxiaoDiscountsGetResponse, []byte, error) {
	var resp FenxiaoDiscountsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.discounts.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDiscountsGetResponse struct {
	Discounts []*Discount `json:"discounts"`
	TotalResults int `json:"total_results"`
}

type FenxiaoDiscountsGetResponseResult struct {
	Response *FenxiaoDiscountsGetResponse `json:"fenxiao_discounts_get_response"`
}





/* 获取当前供应商授权分销商的产品的下载率，计算逻辑同后台页面。downLoadRatio 
获取当前供应商授权分销商的产品的上架率，计算逻辑同后台页面。获取upSelfRatio 
获取当前供应商在分销商店铺中的上架商品占比。upShopRatio 
获取当前供应商在分销商店铺中的成交（已付款）订单笔数占比。 orderShopRatio 
获取当前供应商在分销商店铺中铺货商品UV占店铺商品总UV的比。uvShopRatio */
type FenxiaoDistributorArchivesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销商淘宝店主nick */
func (r *FenxiaoDistributorArchivesGetRequest) SetDistributorUserNick(value string) {
	r.SetValue("distributor_user_nick", value)
}


func (r *FenxiaoDistributorArchivesGetRequest) GetResponse(accessToken string) (*FenxiaoDistributorArchivesGetResponse, []byte, error) {
	var resp FenxiaoDistributorArchivesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.distributor.archives.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDistributorArchivesGetResponse struct {
	DistributorArchive *DistributorArchive `json:"distributor_archive"`
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDistributorArchivesGetResponseResult struct {
	Response *FenxiaoDistributorArchivesGetResponse `json:"fenxiao_distributor_archives_get_response"`
}





/* 供应商查询分销商商品下载记录。 */
type FenxiaoDistributorItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销商ID 。 */
func (r *FenxiaoDistributorItemsGetRequest) SetDistributorId(value string) {
	r.SetValue("distributor_id", value)
}

/* 设置结束时间,空为不设置。 */
func (r *FenxiaoDistributorItemsGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 页码（大于0的整数，默认1） */
func (r *FenxiaoDistributorItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (r *FenxiaoDistributorItemsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 产品ID */
func (r *FenxiaoDistributorItemsGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 设置开始时间。空为不设置。 */
func (r *FenxiaoDistributorItemsGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}


func (r *FenxiaoDistributorItemsGetRequest) GetResponse(accessToken string) (*FenxiaoDistributorItemsGetResponse, []byte, error) {
	var resp FenxiaoDistributorItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.distributor.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDistributorItemsGetResponse struct {
	Records []*FenxiaoItemRecord `json:"records"`
	TotalResults int `json:"total_results"`
}

type FenxiaoDistributorItemsGetResponseResult struct {
	Response *FenxiaoDistributorItemsGetResponse `json:"fenxiao_distributor_items_get_response"`
}





/* 获取分销商品流量，以天为单位统计分销商品的PV，UV */
type FenxiaoDistributorProcuctStaticGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销商淘宝店主nick */
func (r *FenxiaoDistributorProcuctStaticGetRequest) SetDistributorUserNick(value string) {
	r.SetValue("distributor_user_nick", value)
}

/* 供应商商品id，一次可以传多个，每次最多40个。 
以,(英文)作为分隔符。 */
func (r *FenxiaoDistributorProcuctStaticGetRequest) SetProductIdArray(value string) {
	r.SetValue("product_id_array", value)
}


func (r *FenxiaoDistributorProcuctStaticGetRequest) GetResponse(accessToken string) (*FenxiaoDistributorProcuctStaticGetResponse, []byte, error) {
	var resp FenxiaoDistributorProcuctStaticGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.distributor.procuct.static.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDistributorProcuctStaticGetResponse struct {
	DistributorItemFlows []*DistributorItemFlow `json:"distributor_item_flows"`
	IsSuccess bool `json:"is_success"`
}

type FenxiaoDistributorProcuctStaticGetResponseResult struct {
	Response *FenxiaoDistributorProcuctStaticGetResponse `json:"fenxiao_distributor_procuct_static_get_response"`
}





/* 分销商查询供应商产品信息 */
type FenxiaoDistributorProductsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 下载状态，默认未下载。UNDOWNLOAD：未下载，DOWNLOADED：已下载。 */
func (r *FenxiaoDistributorProductsGetRequest) SetDownloadStatus(value string) {
	r.SetValue("download_status", value)
}

/* 结束时间 */
func (r *FenxiaoDistributorProductsGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
func (r *FenxiaoDistributorProductsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (r *FenxiaoDistributorProductsGetRequest) SetItemIds(value string) {
	r.SetValue("item_ids", value)
}

/* 排序。QUANTITY_DESC：库存降序，CREATE_TIME_DESC，创建时间降序。 */
func (r *FenxiaoDistributorProductsGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 页码（大于0的整数，默认1） */
func (r *FenxiaoDistributorProductsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (r *FenxiaoDistributorProductsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 产品ID列表，优先级最高，传了忽略其他查询条件。用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (r *FenxiaoDistributorProductsGetRequest) SetPids(value string) {
	r.SetValue("pids", value)
}

/* 产品线ID */
func (r *FenxiaoDistributorProductsGetRequest) SetProductcatId(value string) {
	r.SetValue("productcat_id", value)
}

/* 开始时间 */
func (r *FenxiaoDistributorProductsGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 供应商nick，分页查询时，必传 */
func (r *FenxiaoDistributorProductsGetRequest) SetSupplierNick(value string) {
	r.SetValue("supplier_nick", value)
}

/* 查询时间类型，默认更新时间。MODIFIED:更新时间，CREATE:创建时间 */
func (r *FenxiaoDistributorProductsGetRequest) SetTimeType(value string) {
	r.SetValue("time_type", value)
}

/* 分销方式，分页查询时，必传。AGENT：代销，DEALER：经销 */
func (r *FenxiaoDistributorProductsGetRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoDistributorProductsGetRequest) GetResponse(accessToken string) (*FenxiaoDistributorProductsGetResponse, []byte, error) {
	var resp FenxiaoDistributorProductsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.distributor.products.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDistributorProductsGetResponse struct {
	HasNext bool `json:"has_next"`
	Products []*FenxiaoProduct `json:"products"`
}

type FenxiaoDistributorProductsGetResponseResult struct {
	Response *FenxiaoDistributorProductsGetResponse `json:"fenxiao_distributor_products_get_response"`
}





/* 查询和当前登录供应商有合作关系的分销商的信息 */
type FenxiaoDistributorsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。 */
func (r *FenxiaoDistributorsGetRequest) SetNicks(value string) {
	r.SetValue("nicks", value)
}


func (r *FenxiaoDistributorsGetRequest) GetResponse(accessToken string) (*FenxiaoDistributorsGetResponse, []byte, error) {
	var resp FenxiaoDistributorsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.distributors.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoDistributorsGetResponse struct {
	Distributors []*Distributor `json:"distributors"`
}

type FenxiaoDistributorsGetResponseResult struct {
	Response *FenxiaoDistributorsGetResponse `json:"fenxiao_distributors_get_response"`
}





/* 新建等级 */
type FenxiaoGradeAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 等级名称，等级名称不可重复 */
func (r *FenxiaoGradeAddRequest) SetName(value string) {
	r.SetValue("name", value)
}


func (r *FenxiaoGradeAddRequest) GetResponse(accessToken string) (*FenxiaoGradeAddResponse, []byte, error) {
	var resp FenxiaoGradeAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.grade.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoGradeAddResponse struct {
	GradeId int `json:"grade_id"`
	IsSuccess bool `json:"is_success"`
}

type FenxiaoGradeAddResponseResult struct {
	Response *FenxiaoGradeAddResponse `json:"fenxiao_grade_add_response"`
}





/* 删除等级 */
type FenxiaoGradeDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 等级ID */
func (r *FenxiaoGradeDeleteRequest) SetGradeId(value string) {
	r.SetValue("grade_id", value)
}


func (r *FenxiaoGradeDeleteRequest) GetResponse(accessToken string) (*FenxiaoGradeDeleteResponse, []byte, error) {
	var resp FenxiaoGradeDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.grade.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoGradeDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoGradeDeleteResponseResult struct {
	Response *FenxiaoGradeDeleteResponse `json:"fenxiao_grade_delete_response"`
}





/* 修改等级 */
type FenxiaoGradeUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 等级ID */
func (r *FenxiaoGradeUpdateRequest) SetGradeId(value string) {
	r.SetValue("grade_id", value)
}

/* 等级名称，等级名称不可重复 */
func (r *FenxiaoGradeUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}


func (r *FenxiaoGradeUpdateRequest) GetResponse(accessToken string) (*FenxiaoGradeUpdateResponse, []byte, error) {
	var resp FenxiaoGradeUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.grade.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoGradeUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoGradeUpdateResponseResult struct {
	Response *FenxiaoGradeUpdateResponse `json:"fenxiao_grade_update_response"`
}





/* 根据供应商ID，查询他的分销商等级信息 */
type FenxiaoGradesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *FenxiaoGradesGetRequest) GetResponse(accessToken string) (*FenxiaoGradesGetResponse, []byte, error) {
	var resp FenxiaoGradesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.grades.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoGradesGetResponse struct {
	FenxiaoGrades []*FenxiaoGrade `json:"fenxiao_grades"`
	TotalResults int `json:"total_results"`
}

type FenxiaoGradesGetResponseResult struct {
	Response *FenxiaoGradesGetResponse `json:"fenxiao_grades_get_response"`
}





/* 获取用户登录信息 */
type FenxiaoLoginUserGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *FenxiaoLoginUserGetRequest) GetResponse(accessToken string) (*FenxiaoLoginUserGetResponse, []byte, error) {
	var resp FenxiaoLoginUserGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.login.user.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoLoginUserGetResponse struct {
	LoginUser *LoginUser `json:"login_user"`
}

type FenxiaoLoginUserGetResponseResult struct {
	Response *FenxiaoLoginUserGetResponse `json:"fenxiao_login_user_get_response"`
}





/* 供应商关闭未付款采购单.可传入一个主单号或是多个子单号，多个子单号之间以‘，’隔开 */
type FenxiaoOrderCloseRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 关闭理由,特殊字符会被转义，会改变长度，有特殊字符是请注意 */
func (r *FenxiaoOrderCloseRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 采购单编号 */
func (r *FenxiaoOrderCloseRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}

/* 子采购单ID，可传多笔子单ID，逗号分隔 */
func (r *FenxiaoOrderCloseRequest) SetSubOrderIds(value string) {
	r.SetValue("sub_order_ids", value)
}


func (r *FenxiaoOrderCloseRequest) GetResponse(accessToken string) (*FenxiaoOrderCloseResponse, []byte, error) {
	var resp FenxiaoOrderCloseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.close", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderCloseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoOrderCloseResponseResult struct {
	Response *FenxiaoOrderCloseResponse `json:"fenxiao_order_close_response"`
}





/* 供应商确认收款（非支付宝交易）。 */
type FenxiaoOrderConfirmPaidRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 确认支付信息（字数小于100） */
func (r *FenxiaoOrderConfirmPaidRequest) SetConfirmRemark(value string) {
	r.SetValue("confirm_remark", value)
}

/* 采购单编号。 */
func (r *FenxiaoOrderConfirmPaidRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}


func (r *FenxiaoOrderConfirmPaidRequest) GetResponse(accessToken string) (*FenxiaoOrderConfirmPaidResponse, []byte, error) {
	var resp FenxiaoOrderConfirmPaidResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.confirm.paid", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderConfirmPaidResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoOrderConfirmPaidResponseResult struct {
	Response *FenxiaoOrderConfirmPaidResponse `json:"fenxiao_order_confirm_paid_response"`
}





/* 分销商创建经销采购单. */
type FenxiaoOrderCreateDealerRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 街道 */
func (r *FenxiaoOrderCreateDealerRequest) SetAddr(value string) {
	r.SetValue("addr", value)
}

/* 买家姓名 */
func (r *FenxiaoOrderCreateDealerRequest) SetBuyerName(value string) {
	r.SetValue("buyer_name", value)
}

/* 市 */
func (r *FenxiaoOrderCreateDealerRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 区 */
func (r *FenxiaoOrderCreateDealerRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 运费，单位为分 */
func (r *FenxiaoOrderCreateDealerRequest) SetLogisticFee(value string) {
	r.SetValue("logistic_fee", value)
}

/* 运输方式，快递,平邮等 */
func (r *FenxiaoOrderCreateDealerRequest) SetLogisticType(value string) {
	r.SetValue("logistic_type", value)
}

/* 留言 */
func (r *FenxiaoOrderCreateDealerRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 买家手机号码和电话号码两者中必须有一个 */
func (r *FenxiaoOrderCreateDealerRequest) SetMobilePhone(value string) {
	r.SetValue("mobile_phone", value)
}

/* erp主订单号，用于去重。当传入号已存在将返回原来的采购单 */
func (r *FenxiaoOrderCreateDealerRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 支付类型,需要供应支持该支付类型 */
func (r *FenxiaoOrderCreateDealerRequest) SetPayType(value string) {
	r.SetValue("pay_type", value)
}

/* 买家电话号码 */
func (r *FenxiaoOrderCreateDealerRequest) SetPhone(value string) {
	r.SetValue("phone", value)
}

/* 省 */
func (r *FenxiaoOrderCreateDealerRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 子单信息,子单内部以‘,’隔开，多个子单以‘;’隔开. 
例(分销产品id,skuid,购买数量,单价;分销产品id:,skuid,购买数量,单价) 
单价的单位位分 */
func (r *FenxiaoOrderCreateDealerRequest) SetSubOrderDetail(value string) {
	r.SetValue("sub_order_detail", value)
}

/* 邮编 */
func (r *FenxiaoOrderCreateDealerRequest) SetZipCode(value string) {
	r.SetValue("zip_code", value)
}


func (r *FenxiaoOrderCreateDealerRequest) GetResponse(accessToken string) (*FenxiaoOrderCreateDealerResponse, []byte, error) {
	var resp FenxiaoOrderCreateDealerResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.create.dealer", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderCreateDealerResponse struct {
	GetModule int `json:"get_module"`
}

type FenxiaoOrderCreateDealerResponseResult struct {
	Response *FenxiaoOrderCreateDealerResponse `json:"fenxiao_order_create_dealer_response"`
}





/* 采购单自定义字段 */
type FenxiaoOrderCustomfieldUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 自定义key */
func (r *FenxiaoOrderCustomfieldUpdateRequest) SetIsvCustomKey(value string) {
	r.SetValue("isv_custom_key", value)
}

/* 自定义的值 */
func (r *FenxiaoOrderCustomfieldUpdateRequest) SetIsvCustomValue(value string) {
	r.SetValue("isv_custom_value", value)
}

/* 采购单id */
func (r *FenxiaoOrderCustomfieldUpdateRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}


func (r *FenxiaoOrderCustomfieldUpdateRequest) GetResponse(accessToken string) (*FenxiaoOrderCustomfieldUpdateResponse, []byte, error) {
	var resp FenxiaoOrderCustomfieldUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.customfield.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderCustomfieldUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoOrderCustomfieldUpdateResponseResult struct {
	Response *FenxiaoOrderCustomfieldUpdateResponse `json:"fenxiao_order_customfield_update_response"`
}





/* 添加采购单留言，最多20条（供应商分销商都可添加） */
type FenxiaoOrderMessageAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 留言内容 */
func (r *FenxiaoOrderMessageAddRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 采购单id */
func (r *FenxiaoOrderMessageAddRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}


func (r *FenxiaoOrderMessageAddRequest) GetResponse(accessToken string) (*FenxiaoOrderMessageAddResponse, []byte, error) {
	var resp FenxiaoOrderMessageAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.message.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderMessageAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoOrderMessageAddResponseResult struct {
	Response *FenxiaoOrderMessageAddResponse `json:"fenxiao_order_message_add_response"`
}





/* 供应商修改采购单备注 */
type FenxiaoOrderRemarkUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 采购单编号 */
func (r *FenxiaoOrderRemarkUpdateRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}

/* 备注旗子(供应商操作) */
func (r *FenxiaoOrderRemarkUpdateRequest) SetSupplierMemo(value string) {
	r.SetValue("supplier_memo", value)
}

/* 旗子的标记，1-5之间的数字。非1-5之间，都采用1作为默认。 
1:红色 
2:黄色 
3:绿色 
4:蓝色 
5:粉红色 */
func (r *FenxiaoOrderRemarkUpdateRequest) SetSupplierMemoFlag(value string) {
	r.SetValue("supplier_memo_flag", value)
}


func (r *FenxiaoOrderRemarkUpdateRequest) GetResponse(accessToken string) (*FenxiaoOrderRemarkUpdateResponse, []byte, error) {
	var resp FenxiaoOrderRemarkUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.order.remark.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrderRemarkUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoOrderRemarkUpdateResponseResult struct {
	Response *FenxiaoOrderRemarkUpdateResponse `json:"fenxiao_order_remark_update_response"`
}





/* 分销商或供应商均可用此接口查询采购单信息. (发货处理请调用物流API中的发货接口) */
type FenxiaoOrdersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
func (r *FenxiaoOrdersGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 多个字段用","分隔。 
 
fields 
如果为空：返回所有采购单对象(purchase_orders)字段。 
如果不为空：返回指定采购单对象(purchase_orders)字段。 
 
例1： 
sub_purchase_orders.tc_order_id 表示只返回tc_order_id  
例2： 
sub_purchase_orders表示只返回子采购单列表 */
func (r *FenxiaoOrdersGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。（大于0的整数。默认为1） */
func (r *FenxiaoOrdersGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。（每页条数不超过50条） */
func (r *FenxiaoOrdersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 采购单编号或分销流水号，若其它参数没传，则此参数必传。 */
func (r *FenxiaoOrdersGetRequest) SetPurchaseOrderId(value string) {
	r.SetValue("purchase_order_id", value)
}

/* 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。若purchase_order_id没传，则此参数必传。 */
func (r *FenxiaoOrdersGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}

/* 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值:<br> *供应商：<br> WAIT_SELLER_SEND_GOODS(等待发货)<br> WAIT_SELLER_CONFIRM_PAY(待确认收款)<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(已发货)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> *分销商：<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(待收货确认)<br> TRADE_FOR_PAY(已付款)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> */
func (r *FenxiaoOrdersGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 采购单下游买家订单id */
func (r *FenxiaoOrdersGetRequest) SetTcOrderId(value string) {
	r.SetValue("tc_order_id", value)
}

/* 可选值：trade_time_type(采购单按照成交时间范围查询),update_time_type(采购单按照更新时间范围查询) */
func (r *FenxiaoOrdersGetRequest) SetTimeType(value string) {
	r.SetValue("time_type", value)
}


func (r *FenxiaoOrdersGetRequest) GetResponse(accessToken string) (*FenxiaoOrdersGetResponse, []byte, error) {
	var resp FenxiaoOrdersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.orders.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoOrdersGetResponse struct {
	PurchaseOrders []*PurchaseOrder `json:"purchase_orders"`
	TotalResults int `json:"total_results"`
}

type FenxiaoOrdersGetResponseResult struct {
	Response *FenxiaoOrdersGetResponse `json:"fenxiao_orders_get_response"`
}





/* 添加分销平台产品数据。业务逻辑与分销系统前台页面一致。 
 
    * 产品图片默认为空 
    * 产品发布后默认为下架状态 */
type FenxiaoProductAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 警戒库存必须是0到29999。 */
func (r *FenxiaoProductAddRequest) SetAlarmNumber(value string) {
	r.SetValue("alarm_number", value)
}

/* 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
func (r *FenxiaoProductAddRequest) SetCategoryId(value string) {
	r.SetValue("category_id", value)
}

/* 所在地：市，例：“杭州” */
func (r *FenxiaoProductAddRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetCostPrice(value string) {
	r.SetValue("cost_price", value)
}

/* 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetDealerCostPrice(value string) {
	r.SetValue("dealer_cost_price", value)
}

/* 产品描述，长度为5到25000字符。 */
func (r *FenxiaoProductAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 折扣ID */
func (r *FenxiaoProductAddRequest) SetDiscountId(value string) {
	r.SetValue("discount_id", value)
}

/* 是否有保修，可选值：false（否）、true（是），默认false。 */
func (r *FenxiaoProductAddRequest) SetHaveGuarantee(value string) {
	r.SetValue("have_guarantee", value)
}

/* 是否有发票，可选值：false（否）、true（是），默认false。 */
func (r *FenxiaoProductAddRequest) SetHaveInvoice(value string) {
	r.SetValue("have_invoice", value)
}

/* 产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片 */
func (r *FenxiaoProductAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 自定义属性。格式为pid:value;pid:value */
func (r *FenxiaoProductAddRequest) SetInputProperties(value string) {
	r.SetValue("input_properties", value)
}

/* 添加产品时，添加入参isAuthz:yes|no  
yes:需要授权  
no:不需要授权  
默认是需要授权 */
func (r *FenxiaoProductAddRequest) SetIsAuthz(value string) {
	r.SetValue("is_authz", value)
}

/* 导入的商品ID */
func (r *FenxiaoProductAddRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 产品名称，长度不超过60个字节。 */
func (r *FenxiaoProductAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商家编码，长度不能超过60个字节。 */
func (r *FenxiaoProductAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (r *FenxiaoProductAddRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。 */
func (r *FenxiaoProductAddRequest) SetPostageEms(value string) {
	r.SetValue("postage_ems", value)
}

/* 快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
func (r *FenxiaoProductAddRequest) SetPostageFast(value string) {
	r.SetValue("postage_fast", value)
}

/* 运费模板ID，参考taobao.postages.get。 */
func (r *FenxiaoProductAddRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。 */
func (r *FenxiaoProductAddRequest) SetPostageOrdinary(value string) {
	r.SetValue("postage_ordinary", value)
}

/* 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。 */
func (r *FenxiaoProductAddRequest) SetPostageType(value string) {
	r.SetValue("postage_type", value)
}

/* 产品线ID */
func (r *FenxiaoProductAddRequest) SetProductcatId(value string) {
	r.SetValue("productcat_id", value)
}

/* 产品属性，格式为pid:vid;pid:vid */
func (r *FenxiaoProductAddRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* 属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名） */
func (r *FenxiaoProductAddRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 所在地：省，例：“浙江” */
func (r *FenxiaoProductAddRequest) SetProv(value string) {
	r.SetValue("prov", value)
}

/* 产品库存必须是1到999999。 */
func (r *FenxiaoProductAddRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
func (r *FenxiaoProductAddRequest) SetRetailPriceHigh(value string) {
	r.SetValue("retail_price_high", value)
}

/* 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetRetailPriceLow(value string) {
	r.SetValue("retail_price_low", value)
}

/* sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (r *FenxiaoProductAddRequest) SetSkuCostPrices(value string) {
	r.SetValue("sku_cost_prices", value)
}

/* sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetSkuDealerCostPrices(value string) {
	r.SetValue("sku_dealer_cost_prices", value)
}

/* sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (r *FenxiaoProductAddRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (r *FenxiaoProductAddRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (r *FenxiaoProductAddRequest) SetSkuQuantitys(value string) {
	r.SetValue("sku_quantitys", value)
}

/* sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序 */
func (r *FenxiaoProductAddRequest) SetSkuStandardPrices(value string) {
	r.SetValue("sku_standard_prices", value)
}

/* 采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetStandardPrice(value string) {
	r.SetValue("standard_price", value)
}

/* 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductAddRequest) SetStandardRetailPrice(value string) {
	r.SetValue("standard_retail_price", value)
}

/* 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做） */
func (r *FenxiaoProductAddRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoProductAddRequest) GetResponse(accessToken string) (*FenxiaoProductAddResponse, []byte, error) {
	var resp FenxiaoProductAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductAddResponse struct {
	Created string `json:"created"`
	Pid int `json:"pid"`
}

type FenxiaoProductAddResponseResult struct {
	Response *FenxiaoProductAddResponse `json:"fenxiao_product_add_response"`
}





/* 等级折扣查询 */
type FenxiaoProductGradepriceGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品id */
func (r *FenxiaoProductGradepriceGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* skuId */
func (r *FenxiaoProductGradepriceGetRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}

/* 经、代销模式（1：代销、2：经销） */
func (r *FenxiaoProductGradepriceGetRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoProductGradepriceGetRequest) GetResponse(accessToken string) (*FenxiaoProductGradepriceGetResponse, []byte, error) {
	var resp FenxiaoProductGradepriceGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.gradeprice.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductGradepriceGetResponse struct {
	GradeDiscounts []*GradeDiscount `json:"grade_discounts"`
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductGradepriceGetResponseResult struct {
	Response *FenxiaoProductGradepriceGetResponse `json:"fenxiao_product_gradeprice_get_response"`
}





/* 供应商可以针对产品不同的sku，指定对应交易类型（代销or经销）方式下，设定折扣方式（按等级or指定分销商）以及对应优惠后的采购价格 */
type FenxiaoProductGradepriceUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 会员等级的id或者分销商id，例如：”1001,2001,1002” */
func (r *FenxiaoProductGradepriceUpdateRequest) SetIds(value string) {
	r.SetValue("ids", value)
}

/* 优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25 */
func (r *FenxiaoProductGradepriceUpdateRequest) SetPrices(value string) {
	r.SetValue("prices", value)
}

/* 产品Id */
func (r *FenxiaoProductGradepriceUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选 */
func (r *FenxiaoProductGradepriceUpdateRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}

/* 选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR" */
func (r *FenxiaoProductGradepriceUpdateRequest) SetTargetType(value string) {
	r.SetValue("target_type", value)
}

/* 交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销) */
func (r *FenxiaoProductGradepriceUpdateRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *FenxiaoProductGradepriceUpdateRequest) GetResponse(accessToken string) (*FenxiaoProductGradepriceUpdateResponse, []byte, error) {
	var resp FenxiaoProductGradepriceUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.gradeprice.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductGradepriceUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductGradepriceUpdateResponseResult struct {
	Response *FenxiaoProductGradepriceUpdateResponse `json:"fenxiao_product_gradeprice_update_response"`
}





/* 产品图片删除，只删除图片信息，不真正删除图片 */
type FenxiaoProductImageDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片位置 */
func (r *FenxiaoProductImageDeleteRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 产品ID */
func (r *FenxiaoProductImageDeleteRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
func (r *FenxiaoProductImageDeleteRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *FenxiaoProductImageDeleteRequest) GetResponse(accessToken string) (*FenxiaoProductImageDeleteResponse, []byte, error) {
	var resp FenxiaoProductImageDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.image.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductImageDeleteResponse struct {
	Created string `json:"created"`
	Result bool `json:"result"`
}

type FenxiaoProductImageDeleteResponseResult struct {
	Response *FenxiaoProductImageDeleteResponse `json:"fenxiao_product_image_delete_response"`
}





/* 产品主图图片空间相对路径或绝对路径添加或更新，或者是图片上传。如果指定位置的图片已存在，则覆盖原有信息。如果位置为1,自动设为主图；如果位置为0，表示属性图片 */
type FenxiaoProductImageUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品图片 */
func (r *FenxiaoProductImageUploadRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (r *FenxiaoProductImageUploadRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图 */
func (r *FenxiaoProductImageUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 产品ID */
func (r *FenxiaoProductImageUploadRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项 */
func (r *FenxiaoProductImageUploadRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *FenxiaoProductImageUploadRequest) GetResponse(accessToken string) (*FenxiaoProductImageUploadResponse, []byte, error) {
	var resp FenxiaoProductImageUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.image.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductImageUploadResponse struct {
	Created string `json:"created"`
	Result bool `json:"result"`
}

type FenxiaoProductImageUploadResponseResult struct {
	Response *FenxiaoProductImageUploadResponse `json:"fenxiao_product_image_upload_response"`
}





/* 创建分销和供应链商品映射关系。 */
type FenxiaoProductMapAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否需要校验商家编码，true不校验，false校验。 */
func (r *FenxiaoProductMapAddRequest) SetNotCheckOuterCode(value string) {
	r.SetValue("not_check_outer_code", value)
}

/* 分销产品id。 */
func (r *FenxiaoProductMapAddRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 后端商品id（如果当前分销产品没有sku和后端商品时需要指定）。 */
func (r *FenxiaoProductMapAddRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 在有sku的情况下，与各个sku对应的后端商品id列表。逗号分隔，顺序需要保证与sku_ids一致。 */
func (r *FenxiaoProductMapAddRequest) SetScItemIds(value string) {
	r.SetValue("sc_item_ids", value)
}

/* 分销产品的sku id。逗号分隔，顺序需要保证与sc_item_ids一致（没有sku就不传）。 */
func (r *FenxiaoProductMapAddRequest) SetSkuIds(value string) {
	r.SetValue("sku_ids", value)
}


func (r *FenxiaoProductMapAddRequest) GetResponse(accessToken string) (*FenxiaoProductMapAddResponse, []byte, error) {
	var resp FenxiaoProductMapAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.map.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductMapAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductMapAddResponseResult struct {
	Response *FenxiaoProductMapAddResponse `json:"fenxiao_product_map_add_response"`
}





/* 删除分销和供应链商品映射关系。 */
type FenxiaoProductMapDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分销产品id。 */
func (r *FenxiaoProductMapDeleteRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 分销产品的sku id列表，逗号分隔，在有sku时需要指定。 */
func (r *FenxiaoProductMapDeleteRequest) SetSkuIds(value string) {
	r.SetValue("sku_ids", value)
}


func (r *FenxiaoProductMapDeleteRequest) GetResponse(accessToken string) (*FenxiaoProductMapDeleteResponse, []byte, error) {
	var resp FenxiaoProductMapDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.map.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductMapDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductMapDeleteResponseResult struct {
	Response *FenxiaoProductMapDeleteResponse `json:"fenxiao_product_map_delete_response"`
}





/* 添加产品SKU信息 */
type FenxiaoProductSkuAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销采购价 */
func (r *FenxiaoProductSkuAddRequest) SetAgentCostPrice(value string) {
	r.SetValue("agent_cost_price", value)
}

/* 经销采购价 */
func (r *FenxiaoProductSkuAddRequest) SetDealerCostPrice(value string) {
	r.SetValue("dealer_cost_price", value)
}

/* 产品ID */
func (r *FenxiaoProductSkuAddRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* sku属性 */
func (r *FenxiaoProductSkuAddRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* sku产品库存，在0到1000000之间，如果不传，则库存为0 */
func (r *FenxiaoProductSkuAddRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 商家编码 */
func (r *FenxiaoProductSkuAddRequest) SetSkuNumber(value string) {
	r.SetValue("sku_number", value)
}

/* 采购基准价，最大值1000000000 */
func (r *FenxiaoProductSkuAddRequest) SetStandardPrice(value string) {
	r.SetValue("standard_price", value)
}


func (r *FenxiaoProductSkuAddRequest) GetResponse(accessToken string) (*FenxiaoProductSkuAddResponse, []byte, error) {
	var resp FenxiaoProductSkuAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.sku.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductSkuAddResponse struct {
	Created string `json:"created"`
	Result bool `json:"result"`
}

type FenxiaoProductSkuAddResponseResult struct {
	Response *FenxiaoProductSkuAddResponse `json:"fenxiao_product_sku_add_response"`
}





/* 根据sku properties删除sku数据 */
type FenxiaoProductSkuDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品id */
func (r *FenxiaoProductSkuDeleteRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* sku属性 */
func (r *FenxiaoProductSkuDeleteRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *FenxiaoProductSkuDeleteRequest) GetResponse(accessToken string) (*FenxiaoProductSkuDeleteResponse, []byte, error) {
	var resp FenxiaoProductSkuDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.sku.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductSkuDeleteResponse struct {
	Created string `json:"created"`
	Result bool `json:"result"`
}

type FenxiaoProductSkuDeleteResponseResult struct {
	Response *FenxiaoProductSkuDeleteResponse `json:"fenxiao_product_sku_delete_response"`
}





/* 产品SKU信息更新 */
type FenxiaoProductSkuUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销采购价 */
func (r *FenxiaoProductSkuUpdateRequest) SetAgentCostPrice(value string) {
	r.SetValue("agent_cost_price", value)
}

/* 经销采购价 */
func (r *FenxiaoProductSkuUpdateRequest) SetDealerCostPrice(value string) {
	r.SetValue("dealer_cost_price", value)
}

/* 产品ID */
func (r *FenxiaoProductSkuUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* sku属性 */
func (r *FenxiaoProductSkuUpdateRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* 产品SKU库存 */
func (r *FenxiaoProductSkuUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 商家编码 */
func (r *FenxiaoProductSkuUpdateRequest) SetSkuNumber(value string) {
	r.SetValue("sku_number", value)
}

/* 采购基准价 */
func (r *FenxiaoProductSkuUpdateRequest) SetStandardPrice(value string) {
	r.SetValue("standard_price", value)
}


func (r *FenxiaoProductSkuUpdateRequest) GetResponse(accessToken string) (*FenxiaoProductSkuUpdateResponse, []byte, error) {
	var resp FenxiaoProductSkuUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.sku.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductSkuUpdateResponse struct {
	Created string `json:"created"`
	Result bool `json:"result"`
}

type FenxiaoProductSkuUpdateResponseResult struct {
	Response *FenxiaoProductSkuUpdateResponse `json:"fenxiao_product_sku_update_response"`
}





/* 产品sku查询 */
type FenxiaoProductSkusGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品ID */
func (r *FenxiaoProductSkusGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}


func (r *FenxiaoProductSkusGetRequest) GetResponse(accessToken string) (*FenxiaoProductSkusGetResponse, []byte, error) {
	var resp FenxiaoProductSkusGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.skus.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductSkusGetResponse struct {
	Skus []*FenxiaoSku `json:"skus"`
	TotalResults int `json:"total_results"`
}

type FenxiaoProductSkusGetResponseResult struct {
	Response *FenxiaoProductSkusGetResponse `json:"fenxiao_product_skus_get_response"`
}





/* 更新分销平台产品数据，不传更新数据返回失败<br> 
1. 对sku进行增、删操作时，原有的sku_ids字段会被忽略，请使用sku_properties和sku_properties_del。<br> */
type FenxiaoProductUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 警戒库存必须是0到29999。 */
func (r *FenxiaoProductUpdateRequest) SetAlarmNumber(value string) {
	r.SetValue("alarm_number", value)
}

/* 所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。 */
func (r *FenxiaoProductUpdateRequest) SetCategoryId(value string) {
	r.SetValue("category_id", value)
}

/* 所在地：市，例：“杭州” */
func (r *FenxiaoProductUpdateRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetCostPrice(value string) {
	r.SetValue("cost_price", value)
}

/* 经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetDealerCostPrice(value string) {
	r.SetValue("dealer_cost_price", value)
}

/* 产品描述，长度为5到25000字符。 */
func (r *FenxiaoProductUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 折扣ID */
func (r *FenxiaoProductUpdateRequest) SetDiscountId(value string) {
	r.SetValue("discount_id", value)
}

/* 是否有保修，可选值：false（否）、true（是），默认false。 */
func (r *FenxiaoProductUpdateRequest) SetHaveGuarantee(value string) {
	r.SetValue("have_guarantee", value)
}

/* 是否有发票，可选值：false（否）、true（是），默认false。 */
func (r *FenxiaoProductUpdateRequest) SetHaveInvoice(value string) {
	r.SetValue("have_invoice", value)
}

/* 主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数 */
func (r *FenxiaoProductUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 自定义属性。格式为pid:value;pid:value */
func (r *FenxiaoProductUpdateRequest) SetInputProperties(value string) {
	r.SetValue("input_properties", value)
}

/* 产品是否需要授权isAuthz:yes|no  
yes:需要授权  
no:不需要授权 */
func (r *FenxiaoProductUpdateRequest) SetIsAuthz(value string) {
	r.SetValue("is_authz", value)
}

/* 产品名称，长度不超过60个字节。 */
func (r *FenxiaoProductUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商家编码，长度不能超过60个字节。 */
func (r *FenxiaoProductUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 产品主图图片空间相对路径或绝对路径 */
func (r *FenxiaoProductUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 产品ID */
func (r *FenxiaoProductUpdateRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (r *FenxiaoProductUpdateRequest) SetPostageEms(value string) {
	r.SetValue("postage_ems", value)
}

/* 快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (r *FenxiaoProductUpdateRequest) SetPostageFast(value string) {
	r.SetValue("postage_fast", value)
}

/* 运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。 */
func (r *FenxiaoProductUpdateRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。 */
func (r *FenxiaoProductUpdateRequest) SetPostageOrdinary(value string) {
	r.SetValue("postage_ordinary", value)
}

/* 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。 */
func (r *FenxiaoProductUpdateRequest) SetPostageType(value string) {
	r.SetValue("postage_type", value)
}

/* 产品属性 */
func (r *FenxiaoProductUpdateRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* 属性别名 */
func (r *FenxiaoProductUpdateRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 所在地：省，例：“浙江” */
func (r *FenxiaoProductUpdateRequest) SetProv(value string) {
	r.SetValue("prov", value)
}

/* 产品库存必须是1到999999。 */
func (r *FenxiaoProductUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。 */
func (r *FenxiaoProductUpdateRequest) SetRetailPriceHigh(value string) {
	r.SetValue("retail_price_high", value)
}

/* 最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetRetailPriceLow(value string) {
	r.SetValue("retail_price_low", value)
}

/* sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
func (r *FenxiaoProductUpdateRequest) SetSkuCostPrices(value string) {
	r.SetValue("sku_cost_prices", value)
}

/* sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetSkuDealerCostPrices(value string) {
	r.SetValue("sku_dealer_cost_prices", value)
}

/* sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。 */
func (r *FenxiaoProductUpdateRequest) SetSkuIds(value string) {
	r.SetValue("sku_ids", value)
}

/* sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",," */
func (r *FenxiaoProductUpdateRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid) 
通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。 */
func (r *FenxiaoProductUpdateRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* 根据sku属性删除sku信息。需要按组删除属性。 */
func (r *FenxiaoProductUpdateRequest) SetSkuPropertiesDel(value string) {
	r.SetValue("sku_properties_del", value)
}

/* sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。 */
func (r *FenxiaoProductUpdateRequest) SetSkuQuantitys(value string) {
	r.SetValue("sku_quantitys", value)
}

/* sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。 */
func (r *FenxiaoProductUpdateRequest) SetSkuStandardPrices(value string) {
	r.SetValue("sku_standard_prices", value)
}

/* 采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetStandardPrice(value string) {
	r.SetValue("standard_price", value)
}

/* 零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。 */
func (r *FenxiaoProductUpdateRequest) SetStandardRetailPrice(value string) {
	r.SetValue("standard_retail_price", value)
}

/* 发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。 */
func (r *FenxiaoProductUpdateRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *FenxiaoProductUpdateRequest) GetResponse(accessToken string) (*FenxiaoProductUpdateResponse, []byte, error) {
	var resp FenxiaoProductUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.product.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductUpdateResponse struct {
	Modified string `json:"modified"`
	Pid int `json:"pid"`
}

type FenxiaoProductUpdateResponseResult struct {
	Response *FenxiaoProductUpdateResponse `json:"fenxiao_product_update_response"`
}





/* 新增产品线 */
type FenxiaoProductcatAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销默认采购价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatAddRequest) SetAgentCostPercent(value string) {
	r.SetValue("agent_cost_percent", value)
}

/* 经销默认采购价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatAddRequest) SetDealerCostPercent(value string) {
	r.SetValue("dealer_cost_percent", value)
}

/* 产品线名称 */
func (r *FenxiaoProductcatAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 最高零售价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatAddRequest) SetRetailHighPercent(value string) {
	r.SetValue("retail_high_percent", value)
}

/* 最低零售价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatAddRequest) SetRetailLowPercent(value string) {
	r.SetValue("retail_low_percent", value)
}


func (r *FenxiaoProductcatAddRequest) GetResponse(accessToken string) (*FenxiaoProductcatAddResponse, []byte, error) {
	var resp FenxiaoProductcatAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.productcat.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductcatAddResponse struct {
	IsSuccess bool `json:"is_success"`
	ProductLineId int `json:"product_line_id"`
}

type FenxiaoProductcatAddResponseResult struct {
	Response *FenxiaoProductcatAddResponse `json:"fenxiao_productcat_add_response"`
}





/* 删除产品线 */
type FenxiaoProductcatDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品线ID */
func (r *FenxiaoProductcatDeleteRequest) SetProductLineId(value string) {
	r.SetValue("product_line_id", value)
}


func (r *FenxiaoProductcatDeleteRequest) GetResponse(accessToken string) (*FenxiaoProductcatDeleteResponse, []byte, error) {
	var resp FenxiaoProductcatDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.productcat.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductcatDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductcatDeleteResponseResult struct {
	Response *FenxiaoProductcatDeleteResponse `json:"fenxiao_productcat_delete_response"`
}





/* 修改产品线 */
type FenxiaoProductcatUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销默认采购价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatUpdateRequest) SetAgentCostPercent(value string) {
	r.SetValue("agent_cost_percent", value)
}

/* 经销默认采购价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatUpdateRequest) SetDealerCostPercent(value string) {
	r.SetValue("dealer_cost_percent", value)
}

/* 产品线名称 */
func (r *FenxiaoProductcatUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 产品线ID */
func (r *FenxiaoProductcatUpdateRequest) SetProductLineId(value string) {
	r.SetValue("product_line_id", value)
}

/* 最高零售价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatUpdateRequest) SetRetailHighPercent(value string) {
	r.SetValue("retail_high_percent", value)
}

/* 最低零售价比例，注意：100.00%，则输入为10000 */
func (r *FenxiaoProductcatUpdateRequest) SetRetailLowPercent(value string) {
	r.SetValue("retail_low_percent", value)
}


func (r *FenxiaoProductcatUpdateRequest) GetResponse(accessToken string) (*FenxiaoProductcatUpdateResponse, []byte, error) {
	var resp FenxiaoProductcatUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.productcat.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductcatUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoProductcatUpdateResponseResult struct {
	Response *FenxiaoProductcatUpdateResponse `json:"fenxiao_productcat_update_response"`
}





/* 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参 */
type FenxiaoProductcatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 返回字段列表 */
func (r *FenxiaoProductcatsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *FenxiaoProductcatsGetRequest) GetResponse(accessToken string) (*FenxiaoProductcatsGetResponse, []byte, error) {
	var resp FenxiaoProductcatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.productcats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductcatsGetResponse struct {
	Productcats []*ProductCat `json:"productcats"`
	TotalResults int `json:"total_results"`
}

type FenxiaoProductcatsGetResponseResult struct {
	Response *FenxiaoProductcatsGetResponse `json:"fenxiao_productcats_get_response"`
}





/* 查询供应商的产品数据。 
 
    * 入参传入pids将优先查询，即只按这个条件查询。 
    *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条) 
    * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。 
    * 入参fields传入images将查询多图数据，不传只返回主图数据。 
    * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据） 
    * 查询结果按照产品发布时间倒序，即时间近的数据在前。 */
type FenxiaoProductsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束修改时间 */
func (r *FenxiaoProductsGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。 */
func (r *FenxiaoProductsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 查询产品列表时，查询入参“是否需要授权” 
yes:需要授权  
no:不需要授权 */
func (r *FenxiaoProductsGetRequest) SetIsAuthz(value string) {
	r.SetValue("is_authz", value)
}

/* 可根据导入的商品ID列表查询，优先级次于产品ID、sku_numbers，高于其他分页查询条件。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (r *FenxiaoProductsGetRequest) SetItemIds(value string) {
	r.SetValue("item_ids", value)
}

/* 商家编码 */
func (r *FenxiaoProductsGetRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 页码（大于0的整数，默认1） */
func (r *FenxiaoProductsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (r *FenxiaoProductsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005” */
func (r *FenxiaoProductsGetRequest) SetPids(value string) {
	r.SetValue("pids", value)
}

/* 产品线ID */
func (r *FenxiaoProductsGetRequest) SetProductcatId(value string) {
	r.SetValue("productcat_id", value)
}

/* sku商家编码 */
func (r *FenxiaoProductsGetRequest) SetSkuNumber(value string) {
	r.SetValue("sku_number", value)
}

/* 开始修改时间 */
func (r *FenxiaoProductsGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 产品状态，可选值：up（上架）、down（下架），不传默认查询所有 */
func (r *FenxiaoProductsGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *FenxiaoProductsGetRequest) GetResponse(accessToken string) (*FenxiaoProductsGetResponse, []byte, error) {
	var resp FenxiaoProductsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.products.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoProductsGetResponse struct {
	Products []*FenxiaoProduct `json:"products"`
	TotalResults int `json:"total_results"`
}

type FenxiaoProductsGetResponseResult struct {
	Response *FenxiaoProductsGetResponse `json:"fenxiao_products_get_response"`
}





/* 用于对分销的子采购单创建退款 */
type FenxiaoRefundCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否退货，只有子采购单发货后，才能申请退货 */
func (r *FenxiaoRefundCreateRequest) SetIsReturnGoods(value string) {
	r.SetValue("is_return_goods", value)
}

/* 此条子采购单是否退邮费，邮费要记在最后一笔申请退款的子单上面 */
func (r *FenxiaoRefundCreateRequest) SetIsReturnPostFee(value string) {
	r.SetValue("is_return_post_fee", value)
}

/* 退款说明,2-200字 */
func (r *FenxiaoRefundCreateRequest) SetRefundDesc(value string) {
	r.SetValue("refund_desc", value)
}

/* 发货前： 
1、缺货，2、拍错商品，3、商品缺少所需样式，4、协商一致退款，5、未及时发货，0、其它 
发货后： 
经销： 
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 
代销： 
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、买家一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 */
func (r *FenxiaoRefundCreateRequest) SetRefundReasonId(value string) {
	r.SetValue("refund_reason_id", value)
}

/* 需要退款的金额，单位为分 */
func (r *FenxiaoRefundCreateRequest) SetReturnFee(value string) {
	r.SetValue("return_fee", value)
}

/* 需要创建退款的子采购单id */
func (r *FenxiaoRefundCreateRequest) SetSubOrderId(value string) {
	r.SetValue("sub_order_id", value)
}


func (r *FenxiaoRefundCreateRequest) GetResponse(accessToken string) (*FenxiaoRefundCreateResponse, []byte, error) {
	var resp FenxiaoRefundCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.refund.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRefundCreateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoRefundCreateResponseResult struct {
	Response *FenxiaoRefundCreateResponse `json:"fenxiao_refund_create_response"`
}





/* 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息 */
type FenxiaoRefundGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否查询下游买家的退款信息 */
func (r *FenxiaoRefundGetRequest) SetQuerySellerRefund(value string) {
	r.SetValue("query_seller_refund", value)
}

/* 要查询的退款子单的id */
func (r *FenxiaoRefundGetRequest) SetSubOrderId(value string) {
	r.SetValue("sub_order_id", value)
}


func (r *FenxiaoRefundGetRequest) GetResponse(accessToken string) (*FenxiaoRefundGetResponse, []byte, error) {
	var resp FenxiaoRefundGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.refund.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRefundGetResponse struct {
	RefundDetail *RefundDetail `json:"refund_detail"`
}

type FenxiaoRefundGetResponseResult struct {
	Response *FenxiaoRefundGetResponse `json:"fenxiao_refund_get_response"`
}





/* 添加退款留言 */
type FenxiaoRefundMessageAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 留言凭证 */
func (r *FenxiaoRefundMessageAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 留言内容 */
func (r *FenxiaoRefundMessageAddRequest) SetMessageContent(value string) {
	r.SetValue("message_content", value)
}

/* 发生退款的子采购单id */
func (r *FenxiaoRefundMessageAddRequest) SetSubOrderId(value string) {
	r.SetValue("sub_order_id", value)
}


func (r *FenxiaoRefundMessageAddRequest) GetResponse(accessToken string) (*FenxiaoRefundMessageAddResponse, []byte, error) {
	var resp FenxiaoRefundMessageAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.refund.message.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRefundMessageAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoRefundMessageAddResponseResult struct {
	Response *FenxiaoRefundMessageAddResponse `json:"fenxiao_refund_message_add_response"`
}





/* 根据子采购单id，查询对应子单的退款留言信息 */
type FenxiaoRefundMessageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 页码。（大于0的整数。默认为1） */
func (r *FenxiaoRefundMessageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。（默认10） */
func (r *FenxiaoRefundMessageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 子采购单id */
func (r *FenxiaoRefundMessageGetRequest) SetSubOrderId(value string) {
	r.SetValue("sub_order_id", value)
}


func (r *FenxiaoRefundMessageGetRequest) GetResponse(accessToken string) (*FenxiaoRefundMessageGetResponse, []byte, error) {
	var resp FenxiaoRefundMessageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.refund.message.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRefundMessageGetResponse struct {
	OrderMessages []*OrderMessage `json:"order_messages"`
	TotalResults int `json:"total_results"`
}

type FenxiaoRefundMessageGetResponseResult struct {
	Response *FenxiaoRefundMessageGetResponse `json:"fenxiao_refund_message_get_response"`
}





/* 修改退款协议 */
type FenxiaoRefundUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否退货，只有子采购单发货后，才能申请退货 */
func (r *FenxiaoRefundUpdateRequest) SetIsReturnGoods(value string) {
	r.SetValue("is_return_goods", value)
}

/* 退款说明 */
func (r *FenxiaoRefundUpdateRequest) SetRefundDesc(value string) {
	r.SetValue("refund_desc", value)
}

/* 发货前： 
1、缺货，2、拍错商品，3、商品缺少所需样式，4、协商一致退款，5、未及时发货，0、其它 
发货后： 
经销： 
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 
代销： 
1、商品质量问题，2、收到的商品不符，3、协商一致退款，4、买家一直未收到货，5、退还多付邮费，6、折扣、赠品、发票等问题，0、其它 */
func (r *FenxiaoRefundUpdateRequest) SetRefundReasonId(value string) {
	r.SetValue("refund_reason_id", value)
}

/* 需要退款的金额 */
func (r *FenxiaoRefundUpdateRequest) SetReturnFee(value string) {
	r.SetValue("return_fee", value)
}

/* 需要修改退款的子采购单id */
func (r *FenxiaoRefundUpdateRequest) SetSubOrderId(value string) {
	r.SetValue("sub_order_id", value)
}


func (r *FenxiaoRefundUpdateRequest) GetResponse(accessToken string) (*FenxiaoRefundUpdateResponse, []byte, error) {
	var resp FenxiaoRefundUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.refund.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRefundUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type FenxiaoRefundUpdateResponseResult struct {
	Response *FenxiaoRefundUpdateResponse `json:"fenxiao_refund_update_response"`
}





/* 合作申请查询 */
type FenxiaoRequisitionsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 申请结束时间yyyy-MM-dd */
func (r *FenxiaoRequisitionsGetRequest) SetApplyEnd(value string) {
	r.SetValue("apply_end", value)
}

/* 申请开始时间yyyy-MM-dd */
func (r *FenxiaoRequisitionsGetRequest) SetApplyStart(value string) {
	r.SetValue("apply_start", value)
}

/* 页码（大于0的整数，默认1） */
func (r *FenxiaoRequisitionsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页记录数（默认20，最大50） */
func (r *FenxiaoRequisitionsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 申请状态（1-申请中、2-成功、3-被退回、4-已撤消、5-过期） */
func (r *FenxiaoRequisitionsGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *FenxiaoRequisitionsGetRequest) GetResponse(accessToken string) (*FenxiaoRequisitionsGetResponse, []byte, error) {
	var resp FenxiaoRequisitionsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.requisitions.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoRequisitionsGetResponse struct {
	IsSuccess bool `json:"is_success"`
	Requisitions []*Requisition `json:"requisitions"`
	TotalResults int `json:"total_results"`
}

type FenxiaoRequisitionsGetResponseResult struct {
	Response *FenxiaoRequisitionsGetResponse `json:"fenxiao_requisitions_get_response"`
}





/* 仅限供应商调用此接口查询经销商品监控信息 */
type FenxiaoTrademonitorGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 经销商的淘宝账号 */
func (r *FenxiaoTrademonitorGetRequest) SetDistributorNick(value string) {
	r.SetValue("distributor_nick", value)
}

/* 结束时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
func (r *FenxiaoTrademonitorGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 多个字段用","分隔。 fields 如果为空：返回所有采购单对象(purchase_orders)字段。 如果不为空：返回指定采购单对象(purchase_orders)字段。例如：trade_monitors.item_title表示只返回item_title */
func (r *FenxiaoTrademonitorGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。（大于0的整数。小于1按1计） */
func (r *FenxiaoTrademonitorGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。（每页条数不超过50条，超过50或小于0均按50计） */
func (r *FenxiaoTrademonitorGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 产品id */
func (r *FenxiaoTrademonitorGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 起始时间 格式 yyyy-MM-dd HH:mm:ss.支持到秒的查询。若不传时分秒，默认为0时0分0秒。 */
func (r *FenxiaoTrademonitorGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}


func (r *FenxiaoTrademonitorGetRequest) GetResponse(accessToken string) (*FenxiaoTrademonitorGetResponse, []byte, error) {
	var resp FenxiaoTrademonitorGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.fenxiao.trademonitor.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type FenxiaoTrademonitorGetResponse struct {
	TotalResults int `json:"total_results"`
	TradeMonitors []*TradeMonitor `json:"trade_monitors"`
}

type FenxiaoTrademonitorGetResponseResult struct {
	Response *FenxiaoTrademonitorGetResponse `json:"fenxiao_trademonitor_get_response"`
}





/* 商家非交易调整库存，调拨出库、盘点等时调用 */
type InventoryAdjustExternalRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 外部订单类型：ALLOCATE:调拨、 RETURN:退货、PURCHACE：采购、BALANCE:盘点、NON_TAOBAO_TRADE：非淘宝交易、OTHERS：其他 */
func (r *InventoryAdjustExternalRequest) SetBizType(value string) {
	r.SetValue("biz_type", value)
}

/* 商家外部定单号 */
func (r *InventoryAdjustExternalRequest) SetBizUniqueCode(value string) {
	r.SetValue("biz_unique_code", value)
}

/* 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,”direction”: 1: 盘盈 -1: 盘亏,参数可选,"quantity":"数量"}] */
func (r *InventoryAdjustExternalRequest) SetItems(value string) {
	r.SetValue("items", value)
}

/* 库存占用返回的操作码.operate_type 为OUTBOUND时，如果是确认事先进行过的库存占用，需要传入当时返回的操作码，并且明细必须与申请时保持一致 */
func (r *InventoryAdjustExternalRequest) SetOccupyOperateCode(value string) {
	r.SetValue("occupy_operate_code", value)
}

/* 业务操作时间 */
func (r *InventoryAdjustExternalRequest) SetOperateTime(value string) {
	r.SetValue("operate_time", value)
}

/* 库存操作类别：INBOUND：入库 OUTBOUND: 出库BALANCE: 盘点 */
func (r *InventoryAdjustExternalRequest) SetOperateType(value string) {
	r.SetValue("operate_type", value)
}

/* 出库时库存扣减类型，operate_type为OUTBOUND时有效。  
AUTO_CALCULATE:自动计算可供扣减，如果库存不够返回失败 ClIENT_FORCE：强制要求最大化扣减，有多少扣多少 */
func (r *InventoryAdjustExternalRequest) SetReduceType(value string) {
	r.SetValue("reduce_type", value)
}

/* 商家仓库编码 */
func (r *InventoryAdjustExternalRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *InventoryAdjustExternalRequest) GetResponse(accessToken string) (*InventoryAdjustExternalResponse, []byte, error) {
	var resp InventoryAdjustExternalResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.adjust.external", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAdjustExternalResponse struct {
	OperateCode string `json:"operate_code"`
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAdjustExternalResponseResult struct {
	Response *InventoryAdjustExternalResponse `json:"inventory_adjust_external_response"`
}





/* 商家交易调整库存，淘宝交易、B2B经销等 */
type InventoryAdjustTradeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商家外部定单号 */
func (r *InventoryAdjustTradeRequest) SetBizUniqueCode(value string) {
	r.SetValue("biz_unique_code", value)
}

/* 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}] */
func (r *InventoryAdjustTradeRequest) SetItems(value string) {
	r.SetValue("items", value)
}

/* 业务操作时间 */
func (r *InventoryAdjustTradeRequest) SetOperateTime(value string) {
	r.SetValue("operate_time", value)
}

/* 订单类型：B2C、B2B */
func (r *InventoryAdjustTradeRequest) SetTbOrderType(value string) {
	r.SetValue("tb_order_type", value)
}


func (r *InventoryAdjustTradeRequest) GetResponse(accessToken string) (*InventoryAdjustTradeResponse, []byte, error) {
	var resp InventoryAdjustTradeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.adjust.trade", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAdjustTradeResponse struct {
	OperateCode string `json:"operate_code"`
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAdjustTradeResponseResult struct {
	Response *InventoryAdjustTradeResponse `json:"inventory_adjust_trade_response"`
}





/* 根据授权结果码获取授权状况 */
type InventoryAuthorizeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 库存分配授权结果码 */
func (r *InventoryAuthorizeGetRequest) SetAuthorizeCode(value string) {
	r.SetValue("authorize_code", value)
}

/* sc_item_id商品后端id */
func (r *InventoryAuthorizeGetRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 分配用户列表，多个用户使用“,“分割开 */
func (r *InventoryAuthorizeGetRequest) SetUserNickList(value string) {
	r.SetValue("user_nick_list", value)
}


func (r *InventoryAuthorizeGetRequest) GetResponse(accessToken string) (*InventoryAuthorizeGetResponse, []byte, error) {
	var resp InventoryAuthorizeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeGetResponse struct {
	AuthorizeList []*InventoryAuthorizeInfo `json:"authorize_list"`
}

type InventoryAuthorizeGetResponseResult struct {
	Response *InventoryAuthorizeGetResponse `json:"inventory_authorize_get_response"`
}





/* 货主根据多个商品列表获取每个商品的授权明细 */
type InventoryAuthorizeGetallRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品编码列表，使用”,”分割多个号码，最大50个 */
func (r *InventoryAuthorizeGetallRequest) SetScItemIdList(value string) {
	r.SetValue("sc_item_id_list", value)
}

/* 指定的商家仓库编码，使用”,”分割多个仓库 */
func (r *InventoryAuthorizeGetallRequest) SetStoreCodeList(value string) {
	r.SetValue("store_code_list", value)
}


func (r *InventoryAuthorizeGetallRequest) GetResponse(accessToken string) (*InventoryAuthorizeGetallResponse, []byte, error) {
	var resp InventoryAuthorizeGetallResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.getall", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeGetallResponse struct {
	AuthorizeList []*InventoryAuthorizeInfo `json:"authorize_list"`
}

type InventoryAuthorizeGetallResponseResult struct {
	Response *InventoryAuthorizeGetallResponse `json:"inventory_authorize_getall_response"`
}





/* 根据库存授权结果码移除该授权下的用户库存授权 */
type InventoryAuthorizeRemoveRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 库存授权结果码 */
func (r *InventoryAuthorizeRemoveRequest) SetAuthorizeCode(value string) {
	r.SetValue("authorize_code", value)
}

/* 后端商品id */
func (r *InventoryAuthorizeRemoveRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 移除授权的目标用户昵称列表，用”,”隔开 */
func (r *InventoryAuthorizeRemoveRequest) SetUserNickList(value string) {
	r.SetValue("user_nick_list", value)
}


func (r *InventoryAuthorizeRemoveRequest) GetResponse(accessToken string) (*InventoryAuthorizeRemoveResponse, []byte, error) {
	var resp InventoryAuthorizeRemoveResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.remove", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeRemoveResponse struct {
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAuthorizeRemoveResponseResult struct {
	Response *InventoryAuthorizeRemoveResponse `json:"inventory_authorize_remove_response"`
}





/* 移除该授权下的用户库存授权 */
type InventoryAuthorizeRemoveallRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后端商品id */
func (r *InventoryAuthorizeRemoveallRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 移除授权的目标用户昵称列表，用”,”隔开 */
func (r *InventoryAuthorizeRemoveallRequest) SetUserNickList(value string) {
	r.SetValue("user_nick_list", value)
}


func (r *InventoryAuthorizeRemoveallRequest) GetResponse(accessToken string) (*InventoryAuthorizeRemoveallResponse, []byte, error) {
	var resp InventoryAuthorizeRemoveallResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.removeall", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeRemoveallResponse struct {
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAuthorizeRemoveallResponseResult struct {
	Response *InventoryAuthorizeRemoveallResponse `json:"inventory_authorize_removeall_response"`
}





/* 商家向其他用户设置配额库存的共享或者独享，支持到渠道，商家自定义库存类别 */
type InventoryAuthorizeSetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 配额授权方式  
PUBLIC: 全共享 
PRIVATE:独享 */
func (r *InventoryAuthorizeSetRequest) SetAuthorizeType(value string) {
	r.SetValue("authorize_type", value)
}

/* 授权明细 
[{“index”:0,“scItemId”:232323,”scItemCode”:”A232”,”storeCode”:”Kj11”,”inventoryType”:1,”channelFlag”:0,”quotaQuantity”:1000,”nickNameList”:”s108,TY000”，“nickName":"ca11"}] 
每次请求的列表数据量是1条，如果authorize_type是PUBLIC,使用nickNameList，否则请用nickName */
func (r *InventoryAuthorizeSetRequest) SetItems(value string) {
	r.SetValue("items", value)
}


func (r *InventoryAuthorizeSetRequest) GetResponse(accessToken string) (*InventoryAuthorizeSetResponse, []byte, error) {
	var resp InventoryAuthorizeSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeSetResponse struct {
	AuthorizeResults []*InventoryAuthorizeInfo `json:"authorize_results"`
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAuthorizeSetResponseResult struct {
	Response *InventoryAuthorizeSetResponse `json:"inventory_authorize_set_response"`
}





/* 对已经授权的配额账户进行配额数量的更改； */
type InventoryAuthorizeUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 库存授权配额编码 */
func (r *InventoryAuthorizeUpdateRequest) SetAuthorizeCode(value string) {
	r.SetValue("authorize_code", value)
}

/* 期望更新后的配额账户库存数(全量覆盖) */
func (r *InventoryAuthorizeUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 后端商品id */
func (r *InventoryAuthorizeUpdateRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}


func (r *InventoryAuthorizeUpdateRequest) GetResponse(accessToken string) (*InventoryAuthorizeUpdateResponse, []byte, error) {
	var resp InventoryAuthorizeUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.authorize.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryAuthorizeUpdateResponse struct {
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryAuthorizeUpdateResponseResult struct {
	Response *InventoryAuthorizeUpdateResponse `json:"inventory_authorize_update_response"`
}





/* 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账 */
type InventoryInitialRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品初始库存信息： [{"scItemId":"商品后端ID，如果有传scItemCode,参数可以为0","scItemCode":"商品商家编码","inventoryType":"库存类型  1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义","quantity":"数量"}] */
func (r *InventoryInitialRequest) SetItems(value string) {
	r.SetValue("items", value)
}

/* 商家仓库编码 */
func (r *InventoryInitialRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *InventoryInitialRequest) GetResponse(accessToken string) (*InventoryInitialResponse, []byte, error) {
	var resp InventoryInitialResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.initial", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryInitialResponse struct {
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryInitialResponseResult struct {
	Response *InventoryInitialResponse `json:"inventory_initial_response"`
}





/* 商家仓商品初始化在各个仓中库存 */
type InventoryInitialItemRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后端商品id */
func (r *InventoryInitialItemRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 商品初始库存信息： [{"storeCode":"必选,商家仓库编号","inventoryType":"可选，库存类型 1：正常,2：损坏,3：冻结,10：质押,11-20:用户自定义,默认为1","quantity":"必选,数量"}] */
func (r *InventoryInitialItemRequest) SetStoreInventorys(value string) {
	r.SetValue("store_inventorys", value)
}


func (r *InventoryInitialItemRequest) GetResponse(accessToken string) (*InventoryInitialItemResponse, []byte, error) {
	var resp InventoryInitialItemResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.initial.item", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryInitialItemResponse struct {
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryInitialItemResponseResult struct {
	Response *InventoryInitialItemResponse `json:"inventory_initial_item_response"`
}





/* 库存占用调整 */
type InventoryOccupyAdjustRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商家外部订单号 */
func (r *InventoryOccupyAdjustRequest) SetBizUniqueCode(value string) {
	r.SetValue("biz_unique_code", value)
}

/* 商品初始库存信息： [{ "TBOrderCode":"淘宝交易号","TBSubOrderCode ":"淘宝子交易单号","originalStoreCode":"交易下单的仓库编码","storeCode":"要切换占用到的商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码","inventoryType":"仓储类型","quantity":"新仓库的占用数量，如果不传，则取用原先的占用数"}] */
func (r *InventoryOccupyAdjustRequest) SetItems(value string) {
	r.SetValue("items", value)
}

/* 业务操作时间 */
func (r *InventoryOccupyAdjustRequest) SetOperateTime(value string) {
	r.SetValue("operate_time", value)
}

/* 订单类型：B2C、B2B */
func (r *InventoryOccupyAdjustRequest) SetTbOrderType(value string) {
	r.SetValue("tb_order_type", value)
}


func (r *InventoryOccupyAdjustRequest) GetResponse(accessToken string) (*InventoryOccupyAdjustResponse, []byte, error) {
	var resp InventoryOccupyAdjustResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.occupy.adjust", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryOccupyAdjustResponse struct {
	OperateCode string `json:"operate_code"`
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryOccupyAdjustResponseResult struct {
	Response *InventoryOccupyAdjustResponse `json:"inventory_occupy_adjust_response"`
}





/* 商家查询商品总体库存信息 */
type InventoryQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后端商品的商家编码列表，控制到50个 */
func (r *InventoryQueryRequest) SetScItemCodes(value string) {
	r.SetValue("sc_item_codes", value)
}

/* 后端商品ID 列表，控制到50个 */
func (r *InventoryQueryRequest) SetScItemIds(value string) {
	r.SetValue("sc_item_ids", value)
}

/* 卖家昵称 */
func (r *InventoryQueryRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}

/* 仓库列表:GLY001^GLY002 */
func (r *InventoryQueryRequest) SetStoreCodes(value string) {
	r.SetValue("store_codes", value)
}


func (r *InventoryQueryRequest) GetResponse(accessToken string) (*InventoryQueryResponse, []byte, error) {
	var resp InventoryQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryQueryResponse struct {
	ItemInventorys []*InventorySum `json:"item_inventorys"`
	TipInfos []*TipInfo `json:"tip_infos"`
}

type InventoryQueryResponseResult struct {
	Response *InventoryQueryResponse `json:"inventory_query_response"`
}





/* 创建商家仓或者更新商家仓信息 */
type InventoryStoreManageRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 仓库的物理地址，可更新 */
func (r *InventoryStoreManageRequest) SetAddress(value string) {
	r.SetValue("address", value)
}

/* 仓库区域名，可更新 */
func (r *InventoryStoreManageRequest) SetAddressAreaName(value string) {
	r.SetValue("address_area_name", value)
}

/* 仓库简称，可更新 */
func (r *InventoryStoreManageRequest) SetAliasName(value string) {
	r.SetValue("alias_name", value)
}

/* 联系人，可更新 */
func (r *InventoryStoreManageRequest) SetContact(value string) {
	r.SetValue("contact", value)
}

/* 参数定义，ADD：新建; UPDATE：更新 */
func (r *InventoryStoreManageRequest) SetOperateType(value string) {
	r.SetValue("operate_type", value)
}

/* 联系电话，可更新 */
func (r *InventoryStoreManageRequest) SetPhone(value string) {
	r.SetValue("phone", value)
}

/* 邮编，可更新 */
func (r *InventoryStoreManageRequest) SetPostcode(value string) {
	r.SetValue("postcode", value)
}

/* 商家的仓库编码，不允许重复，不允许更新 */
func (r *InventoryStoreManageRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}

/* 商家的仓库名称，可更新 */
func (r *InventoryStoreManageRequest) SetStoreName(value string) {
	r.SetValue("store_name", value)
}

/* 仓库类型，可更新。目前只支持自有仓，TYPE_OWN：自有物理仓 */
func (r *InventoryStoreManageRequest) SetStoreType(value string) {
	r.SetValue("store_type", value)
}


func (r *InventoryStoreManageRequest) GetResponse(accessToken string) (*InventoryStoreManageResponse, []byte, error) {
	var resp InventoryStoreManageResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.store.manage", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryStoreManageResponse struct {
	StoreList []*Store `json:"store_list"`
}

type InventoryStoreManageResponseResult struct {
	Response *InventoryStoreManageResponse `json:"inventory_store_manage_response"`
}





/* 查询商家仓信息 */
type InventoryStoreQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商家的仓库编码 */
func (r *InventoryStoreQueryRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *InventoryStoreQueryRequest) GetResponse(accessToken string) (*InventoryStoreQueryResponse, []byte, error) {
	var resp InventoryStoreQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.store.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryStoreQueryResponse struct {
	StoreList []*Store `json:"store_list"`
}

type InventoryStoreQueryResponseResult struct {
	Response *InventoryStoreQueryResponse `json:"inventory_store_query_response"`
}





/* 发布后端商品 */
type ScitemAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 条形码 */
func (r *ScitemAddRequest) SetBarCode(value string) {
	r.SetValue("bar_code", value)
}

/* 品牌id */
func (r *ScitemAddRequest) SetBrandId(value string) {
	r.SetValue("brand_id", value)
}

/* brand_Name */
func (r *ScitemAddRequest) SetBrandName(value string) {
	r.SetValue("brand_name", value)
}

/* 高 单位：mm */
func (r *ScitemAddRequest) SetHeight(value string) {
	r.SetValue("height", value)
}

/* 1表示区域销售，0或是空是非区域销售 */
func (r *ScitemAddRequest) SetIsAreaSale(value string) {
	r.SetValue("is_area_sale", value)
}

/* 是否是贵重品 0:不是 1：是 */
func (r *ScitemAddRequest) SetIsCostly(value string) {
	r.SetValue("is_costly", value)
}

/* 是否危险 0：不是  1：是 */
func (r *ScitemAddRequest) SetIsDangerous(value string) {
	r.SetValue("is_dangerous", value)
}

/* 是否易碎 0：不是  1：是 */
func (r *ScitemAddRequest) SetIsFriable(value string) {
	r.SetValue("is_friable", value)
}

/* 是否保质期：0:不是 1：是 */
func (r *ScitemAddRequest) SetIsWarranty(value string) {
	r.SetValue("is_warranty", value)
}

/* 商品名称 */
func (r *ScitemAddRequest) SetItemName(value string) {
	r.SetValue("item_name", value)
}

/* 0.普通供应链商品 1.供应链组合主商品 */
func (r *ScitemAddRequest) SetItemType(value string) {
	r.SetValue("item_type", value)
}

/* 长度 单位：mm */
func (r *ScitemAddRequest) SetLength(value string) {
	r.SetValue("length", value)
}

/* 0:液体，1：粉体，2：固体 */
func (r *ScitemAddRequest) SetMatterStatus(value string) {
	r.SetValue("matter_status", value)
}

/* 商家编码 */
func (r *ScitemAddRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 价格 单位：分 */
func (r *ScitemAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品属性格式是  p1:v1,p2:v2,p3:v3 */
func (r *ScitemAddRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* remark */
func (r *ScitemAddRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* spuId或是cspuid */
func (r *ScitemAddRequest) SetSpuId(value string) {
	r.SetValue("spu_id", value)
}

/* 体积：立方厘米 */
func (r *ScitemAddRequest) SetVolume(value string) {
	r.SetValue("volume", value)
}

/* 重量 单位：g */
func (r *ScitemAddRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}

/* 宽 单位：mm */
func (r *ScitemAddRequest) SetWidth(value string) {
	r.SetValue("width", value)
}

/* 仓储商编码 */
func (r *ScitemAddRequest) SetWmsCode(value string) {
	r.SetValue("wms_code", value)
}


func (r *ScitemAddRequest) GetResponse(accessToken string) (*ScitemAddResponse, []byte, error) {
	var resp ScitemAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemAddResponse struct {
	ScItem *ScItem `json:"sc_item"`
}

type ScitemAddResponseResult struct {
	Response *ScitemAddResponse `json:"scitem_add_response"`
}





/* 根据id查询商品 */
type ScitemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品id */
func (r *ScitemGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *ScitemGetRequest) GetResponse(accessToken string) (*ScitemGetResponse, []byte, error) {
	var resp ScitemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemGetResponse struct {
	ScItem *ScItem `json:"sc_item"`
}

type ScitemGetResponseResult struct {
	Response *ScitemGetResponse `json:"scitem_get_response"`
}





/* 创建IC商品或分销商品与后端商品的映射关系 */
type ScitemMapAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 前台ic商品id */
func (r *ScitemMapAddRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 默认值为false 
true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联 
false:不进行高级校验 */
func (r *ScitemMapAddRequest) SetNeedCheck(value string) {
	r.SetValue("need_check", value)
}

/* sc_item_id和outer_code 其中一个不能为空 */
func (r *ScitemMapAddRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* sc_item_id和outer_code 其中一个不能为空 */
func (r *ScitemMapAddRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 前台ic商品skuid */
func (r *ScitemMapAddRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}


func (r *ScitemMapAddRequest) GetResponse(accessToken string) (*ScitemMapAddResponse, []byte, error) {
	var resp ScitemMapAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.map.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemMapAddResponse struct {
	OuterCode string `json:"outer_code"`
}

type ScitemMapAddResponseResult struct {
	Response *ScitemMapAddResponse `json:"scitem_map_add_response"`
}





/* 批量分页查询后端商品对应的前端商品 */
type ScitemMapBatchQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后端商品的商家编码 */
func (r *ScitemMapBatchQueryRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 当前页码数 */
func (r *ScitemMapBatchQueryRequest) SetPageIndex(value string) {
	r.SetValue("page_index", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *ScitemMapBatchQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 后端商品id */
func (r *ScitemMapBatchQueryRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}


func (r *ScitemMapBatchQueryRequest) GetResponse(accessToken string) (*ScitemMapBatchQueryResponse, []byte, error) {
	var resp ScitemMapBatchQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.map.batch.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemMapBatchQueryResponse struct {
	ScItemMaps []*ScItemMap `json:"sc_item_maps"`
}

type ScitemMapBatchQueryResponseResult struct {
	Response *ScitemMapBatchQueryResponse `json:"scitem_map_batch_query_response"`
}





/* 根据后端商品Id，失效指定用户的商品与后端商品的映射关系 */
type ScitemMapDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 后台商品ID */
func (r *ScitemMapDeleteRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系 */
func (r *ScitemMapDeleteRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *ScitemMapDeleteRequest) GetResponse(accessToken string) (*ScitemMapDeleteResponse, []byte, error) {
	var resp ScitemMapDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.map.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemMapDeleteResponse struct {
	Module int `json:"module"`
}

type ScitemMapDeleteResponseResult struct {
	Response *ScitemMapDeleteResponse `json:"scitem_map_delete_response"`
}





/* 查找IC商品或分销商品与后端商品的关联信息。skuId如果不传就查找该itemId下所有的sku */
type ScitemMapQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* map_type为1：前台ic商品id 
map_type为2：分销productid */
func (r *ScitemMapQueryRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* map_type为1：前台ic商品skuid  
map_type为2：分销商品的skuid */
func (r *ScitemMapQueryRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}


func (r *ScitemMapQueryRequest) GetResponse(accessToken string) (*ScitemMapQueryResponse, []byte, error) {
	var resp ScitemMapQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.map.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemMapQueryResponse struct {
	ScItemMaps []*ScItemMap `json:"sc_item_maps"`
}

type ScitemMapQueryResponseResult struct {
	Response *ScitemMapQueryResponse `json:"scitem_map_query_response"`
}





/* 根据outerCode查询商品 */
type ScitemOutercodeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品编码 */
func (r *ScitemOutercodeGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}


func (r *ScitemOutercodeGetRequest) GetResponse(accessToken string) (*ScitemOutercodeGetResponse, []byte, error) {
	var resp ScitemOutercodeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.outercode.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemOutercodeGetResponse struct {
	ScItem *ScItem `json:"sc_item"`
}

type ScitemOutercodeGetResponseResult struct {
	Response *ScitemOutercodeGetResponse `json:"scitem_outercode_get_response"`
}





/* 查询后端商品 */
type ScitemQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 条形码 */
func (r *ScitemQueryRequest) SetBarCode(value string) {
	r.SetValue("bar_code", value)
}

/* 商品名称 */
func (r *ScitemQueryRequest) SetItemName(value string) {
	r.SetValue("item_name", value)
}

/* ITEM类型(只允许输入以下英文或空) NORMAL 0:普通商品; COMBINE 1:是否是组合商品 DISTRIBUTION */
func (r *ScitemQueryRequest) SetItemType(value string) {
	r.SetValue("item_type", value)
}

/* 商家给商品的一个编码 */
func (r *ScitemQueryRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 当前页码数 */
func (r *ScitemQueryRequest) SetPageIndex(value string) {
	r.SetValue("page_index", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *ScitemQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓库编码 */
func (r *ScitemQueryRequest) SetWmsCode(value string) {
	r.SetValue("wms_code", value)
}


func (r *ScitemQueryRequest) GetResponse(accessToken string) (*ScitemQueryResponse, []byte, error) {
	var resp ScitemQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemQueryResponse struct {
	PageIndex int `json:"page_index"`
	PageSize int `json:"page_size"`
	QueryPagination *QueryPagination `json:"query_pagination"`
	ScItemList []*ScItem `json:"sc_item_list"`
	TotalPage int `json:"total_page"`
}

type ScitemQueryResponseResult struct {
	Response *ScitemQueryResponse `json:"scitem_query_response"`
}





/* 根据商品ID或商家编码修改后端商品 */
type ScitemUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 条形码 */
func (r *ScitemUpdateRequest) SetBarCode(value string) {
	r.SetValue("bar_code", value)
}

/* 品牌id */
func (r *ScitemUpdateRequest) SetBrandId(value string) {
	r.SetValue("brand_id", value)
}

/* brand_Name */
func (r *ScitemUpdateRequest) SetBrandName(value string) {
	r.SetValue("brand_name", value)
}

/* 高 单位：mm */
func (r *ScitemUpdateRequest) SetHeight(value string) {
	r.SetValue("height", value)
}

/* 1表示区域销售，0或是空是非区域销售 */
func (r *ScitemUpdateRequest) SetIsAreaSale(value string) {
	r.SetValue("is_area_sale", value)
}

/* 是否是贵重品 0:不是 1：是 */
func (r *ScitemUpdateRequest) SetIsCostly(value string) {
	r.SetValue("is_costly", value)
}

/* 是否危险 0：不是  0：是 */
func (r *ScitemUpdateRequest) SetIsDangerous(value string) {
	r.SetValue("is_dangerous", value)
}

/* 是否易碎 0：不是  1：是 */
func (r *ScitemUpdateRequest) SetIsFriable(value string) {
	r.SetValue("is_friable", value)
}

/* 是否保质期：0:不是 1：是 */
func (r *ScitemUpdateRequest) SetIsWarranty(value string) {
	r.SetValue("is_warranty", value)
}

/* 后端商品ID，跟outer_code必须指定一个 */
func (r *ScitemUpdateRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 商品名称 */
func (r *ScitemUpdateRequest) SetItemName(value string) {
	r.SetValue("item_name", value)
}

/* 0.普通供应链商品 1.供应链组合主商品 */
func (r *ScitemUpdateRequest) SetItemType(value string) {
	r.SetValue("item_type", value)
}

/* 长度 单位：mm */
func (r *ScitemUpdateRequest) SetLength(value string) {
	r.SetValue("length", value)
}

/* 0:液体，1：粉体，2：固体 */
func (r *ScitemUpdateRequest) SetMatterStatus(value string) {
	r.SetValue("matter_status", value)
}

/* 商家编码，跟item_id必须指定一个 */
func (r *ScitemUpdateRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* price */
func (r *ScitemUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* remark */
func (r *ScitemUpdateRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* 移除商品属性P列表,P由系统分配：p1；p2 */
func (r *ScitemUpdateRequest) SetRemoveProperties(value string) {
	r.SetValue("remove_properties", value)
}

/* 淘宝SKU产品级编码CSPU ID */
func (r *ScitemUpdateRequest) SetSpuId(value string) {
	r.SetValue("spu_id", value)
}

/* 需要更新的商品属性格式是  p1:v1,p2:v2,p3:v3 */
func (r *ScitemUpdateRequest) SetUpdateProperties(value string) {
	r.SetValue("update_properties", value)
}

/* 体积：立方厘米 */
func (r *ScitemUpdateRequest) SetVolume(value string) {
	r.SetValue("volume", value)
}

/* weight */
func (r *ScitemUpdateRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}

/* 宽 单位：mm */
func (r *ScitemUpdateRequest) SetWidth(value string) {
	r.SetValue("width", value)
}

/* 仓储商编码 */
func (r *ScitemUpdateRequest) SetWmsCode(value string) {
	r.SetValue("wms_code", value)
}


func (r *ScitemUpdateRequest) GetResponse(accessToken string) (*ScitemUpdateResponse, []byte, error) {
	var resp ScitemUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.scitem.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ScitemUpdateResponse struct {
	UpdateRows int `json:"update_rows"`
}

type ScitemUpdateResponseResult struct {
	Response *ScitemUpdateResponse `json:"scitem_update_response"`
}





/* 异步获取历史分销订单, 查询从昨天到3个月内的分销订单数据。所有数据根据gmt_create 降序排序 */
type TopatsFenxiaoOrdersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束时间 格式 yyyyMMdd表示yyyy-MM-dd 00:00:00.开始与结束时间不能超过90天。 */
func (r *TopatsFenxiaoOrdersGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 多个字段用","分隔。 
 
fields 
如果为空：返回所有采购单对象(purchase_orders)字段。 
如果不为空：返回指定采购单对象(purchase_orders)字段。 
 
例1： 
sub_purchase_orders.tc_order_id 表示只返回tc_order_id  
例2： 
sub_purchase_orders表示只返回子采购单列表 */
func (r *TopatsFenxiaoOrdersGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 起始时间 格式 yyyyMMdd表示yyyy-MM-dd 00:00:00.开始与结束时间不能超过90天且开始时间不能为90天前 */
func (r *TopatsFenxiaoOrdersGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}

/* 交易状态，不传默认查询所有采购单根据身份选择自身状态可选值:<br> *供应商：<br> WAIT_SELLER_SEND_GOODS(等待发货)<br> WAIT_SELLER_CONFIRM_PAY(待确认收款)<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(已发货)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> *分销商：<br> WAIT_BUYER_PAY(等待付款)<br> WAIT_BUYER_CONFIRM_GOODS(待收货确认)<br> TRADE_FOR_PAY(已付款)<br> TRADE_REFUNDING(退款中)<br> TRADE_FINISHED(采购成功)<br> TRADE_CLOSED(已关闭)<br> */
func (r *TopatsFenxiaoOrdersGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *TopatsFenxiaoOrdersGetRequest) GetResponse(accessToken string) (*TopatsFenxiaoOrdersGetResponse, []byte, error) {
	var resp TopatsFenxiaoOrdersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.fenxiao.orders.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsFenxiaoOrdersGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsFenxiaoOrdersGetResponseResult struct {
	Response *TopatsFenxiaoOrdersGetResponse `json:"topats_fenxiao_orders_get_response"`
}





/* 查询库存明细 */
type InventoryIpcInventorydetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主订单号，可选 */
func (r *InventoryIpcInventorydetailGetRequest) SetBizOrderId(value string) {
	r.SetValue("biz_order_id", value)
}

/* 子订单号，可选 */
func (r *InventoryIpcInventorydetailGetRequest) SetBizSubOrderId(value string) {
	r.SetValue("biz_sub_order_id", value)
}

/* 当前页数 */
func (r *InventoryIpcInventorydetailGetRequest) SetPageIndex(value string) {
	r.SetValue("page_index", value)
}

/* 一页显示多少条 */
func (r *InventoryIpcInventorydetailGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓储商品id */
func (r *InventoryIpcInventorydetailGetRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 1:查询预扣  4：查询占用 */
func (r *InventoryIpcInventorydetailGetRequest) SetStatusQuery(value string) {
	r.SetValue("status_query", value)
}


func (r *InventoryIpcInventorydetailGetRequest) GetResponse(accessToken string) (*InventoryIpcInventorydetailGetResponse, []byte, error) {
	var resp InventoryIpcInventorydetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.ipc.inventorydetail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryIpcInventorydetailGetResponse struct {
	InventoryDetails []*IpcInventoryDetailDo `json:"inventory_details"`
}

type InventoryIpcInventorydetailGetResponseResult struct {
	Response *InventoryIpcInventorydetailGetResponse `json:"inventory_ipc_inventorydetail_get_response"`
}



