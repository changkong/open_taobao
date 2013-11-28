// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package eai

import (
	"github.com/yaofangou/open_taobao"
)





/* 用户接入时调用该 API 进行授权初始化的操作。 
调用该 API的目的是告知订单数据中心， 
外部用户需要的是哪个卖家的订单数据。 
同时告知我们对开放出去的该卖家的订单数据的负责用户是谁。 
"负责用户"可以是卖家本身。也可以是ISV。 */
type TmallEaiBaseGatewayRegisterRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 用户应用的回调URL,若是普通OPEN-API用户(user_type=0)这个字段可以不填。但是若是JIP用户注册，也就是(user_type=1)时,则此字段必填。否则调用不会成功。 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetCallBackUrl(value string) {
	r.SetValue("call_back_url", value)
}

/* 字符集编码GBK,UTF-8默认,GB2312 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetCharSet(value string) {
	r.SetValue("char_set", value)
}

/* 用户所在的城市 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 数据格式:XML,JSON。 
默认:JSON */
func (r *TmallEaiBaseGatewayRegisterRequest) SetContentType(value string) {
	r.SetValue("content_type", value)
}

/* 接入方式描述，为了更好的方便我们为您服务请务必详细描述下您的APP情况。 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 用户所在区域 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetDistrict(value string) {
	r.SetValue("district", value)
}

/* 用户的联系邮箱 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetEmail(value string) {
	r.SetValue("email", value)
}

/* 暂时还没有支持 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetMaxFlow(value string) {
	r.SetValue("max_flow", value)
}

/* 用户手机号码 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetMobile(value string) {
	r.SetValue("mobile", value)
}

/* 用户通知接收方式,邮件,电话或者其他。 
email:邮件; 
message:短信; 
aliwangwang:阿里旺旺弹出消息。 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetNotifyType(value string) {
	r.SetValue("notify_type", value)
}

/* 联系人姓名 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetPrincipal(value string) {
	r.SetValue("principal", value)
}

/* 用户联系电话-座机 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetTelephone(value string) {
	r.SetValue("telephone", value)
}

/* 接口/api名称: 
天猫退款消息可选值为: 
tmall.eai.order.refund.refundBill.push 
tmall.eai.order.refund.refundStatus.push */
func (r *TmallEaiBaseGatewayRegisterRequest) SetTopic(value string) {
	r.SetValue("topic", value)
}

/* Topic组.暂时还不能支持. */
func (r *TmallEaiBaseGatewayRegisterRequest) SetTopicGroup(value string) {
	r.SetValue("topic_group", value)
}

/* url协议 
HTTP默认 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetUrlProtocal(value string) {
	r.SetValue("url_protocal", value)
}

/* 0:普通TOP-OPEN-API用户; 
1:JIP用户。 */
func (r *TmallEaiBaseGatewayRegisterRequest) SetUserType(value string) {
	r.SetValue("user_type", value)
}


func (r *TmallEaiBaseGatewayRegisterRequest) GetResponse(accessToken string) (*TmallEaiBaseGatewayRegisterResponse, []byte, error) {
	var resp TmallEaiBaseGatewayRegisterResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.base.gateway.register", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiBaseGatewayRegisterResponse struct {
	Count int `json:"count"`
	IsSuccess bool `json:"is_success"`
}

type TmallEaiBaseGatewayRegisterResponseResult struct {
	Response *TmallEaiBaseGatewayRegisterResponse `json:"tmall_eai_base_gateway_register_response"`
}





/* 查询退货单 或者 退款单数量 */
type TmallEaiOrderRefundBillsumGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查找数量的单据类型 refund_bill:退款单， return_bill:退货单 */
func (r *TmallEaiOrderRefundBillsumGetRequest) SetBillType(value string) {
	r.SetValue("bill_type", value)
}

/* 批量查询结束时间 */
func (r *TmallEaiOrderRefundBillsumGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 批量查询开始时间 */
func (r *TmallEaiOrderRefundBillsumGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 退货单 退款单状态 */
func (r *TmallEaiOrderRefundBillsumGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *TmallEaiOrderRefundBillsumGetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundBillsumGetResponse, []byte, error) {
	var resp TmallEaiOrderRefundBillsumGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.billsum.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundBillsumGetResponse struct {
	BillSum int `json:"bill_sum"`
}

type TmallEaiOrderRefundBillsumGetResponseResult struct {
	Response *TmallEaiOrderRefundBillsumGetResponse `json:"tmall_eai_order_refund_billsum_get_response"`
}





/* 退款单审核结果回写，标记前端批量可退款 */
type TmallEaiOrderRefundExamineRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 审核留言 */
func (r *TmallEaiOrderRefundExamineRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 审核人姓名 */
func (r *TmallEaiOrderRefundExamineRequest) SetOperator(value string) {
	r.SetValue("operator", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundExamineRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundExamineRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 退款版本号 */
func (r *TmallEaiOrderRefundExamineRequest) SetRefundVersion(value string) {
	r.SetValue("refund_version", value)
}


func (r *TmallEaiOrderRefundExamineRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundExamineResponse, []byte, error) {
	var resp TmallEaiOrderRefundExamineResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.examine", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundExamineResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundExamineResponseResult struct {
	Response *TmallEaiOrderRefundExamineResponse `json:"tmall_eai_order_refund_examine_response"`
}





/* 退款单反审核结果回写 */
type TmallEaiOrderRefundExamineCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 反审核留言 */
func (r *TmallEaiOrderRefundExamineCancelRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 反审核人姓名 */
func (r *TmallEaiOrderRefundExamineCancelRequest) SetOperator(value string) {
	r.SetValue("operator", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundExamineCancelRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundExamineCancelRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 退款版本号 */
func (r *TmallEaiOrderRefundExamineCancelRequest) SetRefundVersion(value string) {
	r.SetValue("refund_version", value)
}


func (r *TmallEaiOrderRefundExamineCancelRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundExamineCancelResponse, []byte, error) {
	var resp TmallEaiOrderRefundExamineCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.examine.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundExamineCancelResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundExamineCancelResponseResult struct {
	Response *TmallEaiOrderRefundExamineCancelResponse `json:"tmall_eai_order_refund_examine_cancel_response"`
}





/* 查询单笔退款单 */
type TmallEaiOrderRefundGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 退款单编号 */
func (r *TmallEaiOrderRefundGetRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundGetRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}


func (r *TmallEaiOrderRefundGetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGetResponse, []byte, error) {
	var resp TmallEaiOrderRefundGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGetResponse struct {
	RefundBill *RefundBill `json:"refund_bill"`
}

type TmallEaiOrderRefundGetResponseResult struct {
	Response *TmallEaiOrderRefundGetResponse `json:"tmall_eai_order_refund_get_response"`
}





/* 卖家同意退货 */
type TmallEaiOrderRefundGoodReturnAgreeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 同意退货留言 */
func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) SetMessage(value string) {
	r.SetValue("message", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 退款版本号 */
func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) SetRefundVersion(value string) {
	r.SetValue("refund_version", value)
}

/* 卖家收货地址编号 */
func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) SetSellerLogisticsAddressId(value string) {
	r.SetValue("seller_logistics_address_id", value)
}


func (r *TmallEaiOrderRefundGoodReturnAgreeRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGoodReturnAgreeResponse, []byte, error) {
	var resp TmallEaiOrderRefundGoodReturnAgreeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.good.return.agree", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGoodReturnAgreeResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundGoodReturnAgreeResponseResult struct {
	Response *TmallEaiOrderRefundGoodReturnAgreeResponse `json:"tmall_eai_order_refund_good_return_agree_response"`
}





/* 卖家验货回写订单号 */
type TmallEaiOrderRefundGoodReturnCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流公司编号 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* 1.验货通过 
2.验货不通过 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetConfirmResult(value string) {
	r.SetValue("confirm_result", value)
}

/* 验货时间 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetConfirmTime(value string) {
	r.SetValue("confirm_time", value)
}

/* 验货人员 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetOperator(value string) {
	r.SetValue("operator", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 物流运单号 */
func (r *TmallEaiOrderRefundGoodReturnCheckRequest) SetSid(value string) {
	r.SetValue("sid", value)
}


func (r *TmallEaiOrderRefundGoodReturnCheckRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGoodReturnCheckResponse, []byte, error) {
	var resp TmallEaiOrderRefundGoodReturnCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.good.return.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGoodReturnCheckResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundGoodReturnCheckResponseResult struct {
	Response *TmallEaiOrderRefundGoodReturnCheckResponse `json:"tmall_eai_order_refund_good_return_check_response"`
}





/* 查询单笔退货单 */
type TmallEaiOrderRefundGoodReturnGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 退款单编号 */
func (r *TmallEaiOrderRefundGoodReturnGetRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundGoodReturnGetRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}


func (r *TmallEaiOrderRefundGoodReturnGetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGoodReturnGetResponse, []byte, error) {
	var resp TmallEaiOrderRefundGoodReturnGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.good.return.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGoodReturnGetResponse struct {
	ReturnBill *ReturnBill `json:"return_bill"`
}

type TmallEaiOrderRefundGoodReturnGetResponseResult struct {
	Response *TmallEaiOrderRefundGoodReturnGetResponse `json:"tmall_eai_order_refund_good_return_get_response"`
}





/* 批量退货单查询 */
type TmallEaiOrderRefundGoodReturnMgetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 批量查询结束时间 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:10;最大值:40 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 批量查询开始时间 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 退货单 退款单状态 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
func (r *TmallEaiOrderRefundGoodReturnMgetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *TmallEaiOrderRefundGoodReturnMgetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGoodReturnMgetResponse, []byte, error) {
	var resp TmallEaiOrderRefundGoodReturnMgetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.good.return.mget", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGoodReturnMgetResponse struct {
	HasNext bool `json:"has_next"`
	ReturnBillList []*ReturnBill `json:"return_bill_list"`
	TotalResults int `json:"total_results"`
}

type TmallEaiOrderRefundGoodReturnMgetResponseResult struct {
	Response *TmallEaiOrderRefundGoodReturnMgetResponse `json:"tmall_eai_order_refund_good_return_mget_response"`
}





/* 卖家拒绝退货 */
type TmallEaiOrderRefundGoodReturnRefuseRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 退款单编号 */
func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 退款版本号 */
func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) SetRefundVersion(value string) {
	r.SetValue("refund_version", value)
}

/* 拒绝退款原因留言 */
func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) SetRefuseMessage(value string) {
	r.SetValue("refuse_message", value)
}

/* 拒绝退款举证上传 */
func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) SetRefuseProof(value string) {
	r.SetValue("refuse_proof", value)
}


func (r *TmallEaiOrderRefundGoodReturnRefuseRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundGoodReturnRefuseResponse, []byte, error) {
	var resp TmallEaiOrderRefundGoodReturnRefuseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.good.return.refuse", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundGoodReturnRefuseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundGoodReturnRefuseResponseResult struct {
	Response *TmallEaiOrderRefundGoodReturnRefuseResponse `json:"tmall_eai_order_refund_good_return_refuse_response"`
}





/* 退款留言查询 */
type TmallEaiOrderRefundMessageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *TmallEaiOrderRefundMessageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:40 */
func (r *TmallEaiOrderRefundMessageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundMessageGetRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundMessageGetRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}


func (r *TmallEaiOrderRefundMessageGetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundMessageGetResponse, []byte, error) {
	var resp TmallEaiOrderRefundMessageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.message.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundMessageGetResponse struct {
	MessageList []*TmallRefundMessage `json:"message_list"`
	TotalResults int `json:"total_results"`
}

type TmallEaiOrderRefundMessageGetResponseResult struct {
	Response *TmallEaiOrderRefundMessageGetResponse `json:"tmall_eai_order_refund_message_get_response"`
}





/* 批量查询退款单 */
type TmallEaiOrderRefundMgetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 批量查询结束时间 */
func (r *TmallEaiOrderRefundMgetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *TmallEaiOrderRefundMgetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:10;最大值:40 */
func (r *TmallEaiOrderRefundMgetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 批量查询开始时间 */
func (r *TmallEaiOrderRefundMgetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 退货单 退款单状态 */
func (r *TmallEaiOrderRefundMgetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
func (r *TmallEaiOrderRefundMgetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}


func (r *TmallEaiOrderRefundMgetRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundMgetResponse, []byte, error) {
	var resp TmallEaiOrderRefundMgetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.mget", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundMgetResponse struct {
	HasNext bool `json:"has_next"`
	RefundBillList []*RefundBill `json:"refund_bill_list"`
	TotalResults int `json:"total_results"`
}

type TmallEaiOrderRefundMgetResponseResult struct {
	Response *TmallEaiOrderRefundMgetResponse `json:"tmall_eai_order_refund_mget_response"`
}





/* 订单截停回写 */
type TmallEaiOrderRefundOrderHoldRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 1. 截停成功 
2. 截停失败 */
func (r *TmallEaiOrderRefundOrderHoldRequest) SetHoldResult(value string) {
	r.SetValue("hold_result", value)
}

/* 截停环节 
1. 订单处理环节 
2. 发货环节 */
func (r *TmallEaiOrderRefundOrderHoldRequest) SetHoldStep(value string) {
	r.SetValue("hold_step", value)
}

/* 截单时间 */
func (r *TmallEaiOrderRefundOrderHoldRequest) SetHoldTime(value string) {
	r.SetValue("hold_time", value)
}

/* 退款单编号 */
func (r *TmallEaiOrderRefundOrderHoldRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundOrderHoldRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}


func (r *TmallEaiOrderRefundOrderHoldRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundOrderHoldResponse, []byte, error) {
	var resp TmallEaiOrderRefundOrderHoldResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.order.hold", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundOrderHoldResponse struct {
	IsSuccess string `json:"is_success"`
}

type TmallEaiOrderRefundOrderHoldResponseResult struct {
	Response *TmallEaiOrderRefundOrderHoldResponse `json:"tmall_eai_order_refund_order_hold_response"`
}





/* 卖家拒绝退款 */
type TmallEaiOrderRefundRefuseRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 退款单编号 */
func (r *TmallEaiOrderRefundRefuseRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 售中：onsale 
售后：aftersale */
func (r *TmallEaiOrderRefundRefuseRequest) SetRefundPhase(value string) {
	r.SetValue("refund_phase", value)
}

/* 退款版本号 */
func (r *TmallEaiOrderRefundRefuseRequest) SetRefundVersion(value string) {
	r.SetValue("refund_version", value)
}

/* 拒绝退款原因留言 */
func (r *TmallEaiOrderRefundRefuseRequest) SetRefuseMessage(value string) {
	r.SetValue("refuse_message", value)
}

/* 拒绝退款举证上传 */
func (r *TmallEaiOrderRefundRefuseRequest) SetRefuseProof(value string) {
	r.SetValue("refuse_proof", value)
}


func (r *TmallEaiOrderRefundRefuseRequest) GetResponse(accessToken string) (*TmallEaiOrderRefundRefuseResponse, []byte, error) {
	var resp TmallEaiOrderRefundRefuseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.eai.order.refund.refuse", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallEaiOrderRefundRefuseResponse struct {
	IsSuccess bool `json:"is_success"`
}

type TmallEaiOrderRefundRefuseResponseResult struct {
	Response *TmallEaiOrderRefundRefuseResponse `json:"tmall_eai_order_refund_refuse_response"`
}



