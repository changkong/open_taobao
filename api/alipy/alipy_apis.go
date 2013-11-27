// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package alipy

import (
	"github.com/changkong/open_taobao"
)





/* 创建生活账单 */
type AlipayEbppBillAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 如果传入调用api有淘宝授权的session可以不传这个字段 */
func (r *AlipayEbppBillAddRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 账单的账期，例如201203表示2012年3月的账单。 */
func (r *AlipayEbppBillAddRequest) SetBillDate(value string) {
	r.SetValue("bill_date", value)
}

/* 账单单据号，例如水费单号，手机号，电费号，信用卡卡号。没有唯一性要求。 */
func (r *AlipayEbppBillAddRequest) SetBillKey(value string) {
	r.SetValue("bill_key", value)
}

/* 支付宝给每个出账机构指定了一个对应的英文短名称来唯一表示该收费单位。 */
func (r *AlipayEbppBillAddRequest) SetChargeInst(value string) {
	r.SetValue("charge_inst", value)
}

/* 输出机构的业务流水号，需要保证唯一性。 */
func (r *AlipayEbppBillAddRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 用户的手机号 */
func (r *AlipayEbppBillAddRequest) SetMobile(value string) {
	r.SetValue("mobile", value)
}

/* 支付宝订单类型。公共事业缴纳JF,信用卡还款HK */
func (r *AlipayEbppBillAddRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}

/* 拥有该账单的用户姓名 */
func (r *AlipayEbppBillAddRequest) SetOwnerName(value string) {
	r.SetValue("owner_name", value)
}

/* 缴费金额。用户支付的总金额。单位为：RMB Yuan。取值范围为[0.01，100000000.00]，精确到小数点后两位。 */
func (r *AlipayEbppBillAddRequest) SetPayAmount(value string) {
	r.SetValue("pay_amount", value)
}

/* 账单的服务费。 */
func (r *AlipayEbppBillAddRequest) SetServiceAmount(value string) {
	r.SetValue("service_amount", value)
}

/* 子业务类型是业务类型的下一级概念，例如：WATER表示JF下面的水费，ELECTRIC表示JF下面的电费，GAS表示JF下面的燃气费。 */
func (r *AlipayEbppBillAddRequest) SetSubOrderType(value string) {
	r.SetValue("sub_order_type", value)
}

/* 交通违章地点，sub_order_type=TRAFFIC时填写。 */
func (r *AlipayEbppBillAddRequest) SetTrafficLocation(value string) {
	r.SetValue("traffic_location", value)
}

/* 违章行为，sub_order_type=TRAFFIC时填写。 */
func (r *AlipayEbppBillAddRequest) SetTrafficRegulations(value string) {
	r.SetValue("traffic_regulations", value)
}


func (r *AlipayEbppBillAddRequest) GetResponse(accessToken string) (*AlipayEbppBillAddResponse, []byte, error) {
	var resp AlipayEbppBillAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.ebpp.bill.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayEbppBillAddResponse struct {
	AlipayOrderNo string `json:"alipay_order_no"`
	BillDate string `json:"bill_date"`
	BillKey string `json:"bill_key"`
	ChargeInst string `json:"charge_inst"`
	ChargeInstName string `json:"charge_inst_name"`
	MerchantOrderNo string `json:"merchant_order_no"`
	OrderType string `json:"order_type"`
	OwnerName string `json:"owner_name"`
	PayAmount string `json:"pay_amount"`
	ServiceAmount string `json:"service_amount"`
	SubOrderType string `json:"sub_order_type"`
}

type AlipayEbppBillAddResponseResult struct {
	Response *AlipayEbppBillAddResponse `json:"alipay_ebpp_bill_add_response"`
}





/* 查询生活账单信息 */
type AlipayEbppBillGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝授权凭证，如果有淘宝的session可以不传 */
func (r *AlipayEbppBillGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 输出机构的业务流水号，需要保证唯一性。 */
func (r *AlipayEbppBillGetRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 支付宝订单类型。公共事业缴纳JF,信用卡还款HK */
func (r *AlipayEbppBillGetRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}


func (r *AlipayEbppBillGetRequest) GetResponse(accessToken string) (*AlipayEbppBillGetResponse, []byte, error) {
	var resp AlipayEbppBillGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.ebpp.bill.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayEbppBillGetResponse struct {
	AlipayOrderNo string `json:"alipay_order_no"`
	BillDate string `json:"bill_date"`
	BillKey string `json:"bill_key"`
	ChargeInst string `json:"charge_inst"`
	ChargeInstName string `json:"charge_inst_name"`
	MerchantOrderNo string `json:"merchant_order_no"`
	OrderStatus string `json:"order_status"`
	OrderType string `json:"order_type"`
	OwnerName string `json:"owner_name"`
	PayAmount float64 `json:"pay_amount"`
	ServiceAmount float64 `json:"service_amount"`
	SubOrderType string `json:"sub_order_type"`
	TrafficLocation string `json:"traffic_location"`
	TrafficRegulations string `json:"traffic_regulations"`
}

type AlipayEbppBillGetResponseResult struct {
	Response *AlipayEbppBillGetResponse `json:"alipay_ebpp_bill_get_response"`
}





/* 对生活账单进行支付接口 */
type AlipayEbppBillPayRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝的业务订单号，具有唯一性。 */
func (r *AlipayEbppBillPayRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 如有有淘宝授权的session可以不传这个字段 */
func (r *AlipayEbppBillPayRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 输出机构的业务流水号，需要保证唯一性。 */
func (r *AlipayEbppBillPayRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 支付宝订单类型。公共事业缴纳JF,信用卡还款HK */
func (r *AlipayEbppBillPayRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}


func (r *AlipayEbppBillPayRequest) GetResponse(accessToken string) (*AlipayEbppBillPayResponse, []byte, error) {
	var resp AlipayEbppBillPayResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.ebpp.bill.pay", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayEbppBillPayResponse struct {
	AlipayOrderNo string `json:"alipay_order_no"`
	MerchantOrderNo string `json:"merchant_order_no"`
	OrderType string `json:"order_type"`
}

type AlipayEbppBillPayResponseResult struct {
	Response *AlipayEbppBillPayResponse `json:"alipay_ebpp_bill_pay_response"`
}





/* 创建账单之后，调用此API返回一个用户自己去收银台付款的URL，用户去收银台页面完成付款操作。 */
type AlipayEbppBillPayurlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝的业务订单号，具有唯一性。 */
func (r *AlipayEbppBillPayurlGetRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 如有有淘宝授权的session可以不传这个字段 */
func (r *AlipayEbppBillPayurlGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 输出机构的业务流水号，需要保证唯一性。 */
func (r *AlipayEbppBillPayurlGetRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 支付宝订单类型。公共事业缴纳JF,信用卡还款HK。 */
func (r *AlipayEbppBillPayurlGetRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}


func (r *AlipayEbppBillPayurlGetRequest) GetResponse(accessToken string) (*AlipayEbppBillPayurlGetResponse, []byte, error) {
	var resp AlipayEbppBillPayurlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.ebpp.bill.payurl.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayEbppBillPayurlGetResponse struct {
	PayUrl string `json:"pay_url"`
}

type AlipayEbppBillPayurlGetResponseResult struct {
	Response *AlipayEbppBillPayurlGetResponse `json:"alipay_ebpp_bill_payurl_get_response"`
}





/* 创建并支付冻结单之后，使用此接口获取有密支付的url 
这个接口调用的前提是创建冻结金的参数信息pay_confirm=PAY_PASSWORD */
type AlipayMicropayOrderConfirmpayurlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝订单号，冻结流水号.这个是创建冻结订单支付宝返回的 */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 支付金额,区间必须在[0.01,30]，只能保留小数点后两位 */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetAmount(value string) {
	r.SetValue("amount", value)
}

/* 支付宝用户给应用的授权. */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 支付备注 */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 收款方的支付宝ID */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetReceiveUserId(value string) {
	r.SetValue("receive_user_id", value)
}

/* 本次转账的外部单据号（只能由字母和数字组成,maxlength=32） */
func (r *AlipayMicropayOrderConfirmpayurlGetRequest) SetTransferOutOrderNo(value string) {
	r.SetValue("transfer_out_order_no", value)
}


func (r *AlipayMicropayOrderConfirmpayurlGetRequest) GetResponse(accessToken string) (*AlipayMicropayOrderConfirmpayurlGetResponse, []byte, error) {
	var resp AlipayMicropayOrderConfirmpayurlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.confirmpayurl.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderConfirmpayurlGetResponse struct {
	SinglePayDetail *SinglePayDetail `json:"single_pay_detail"`
}

type AlipayMicropayOrderConfirmpayurlGetResponseResult struct {
	Response *AlipayMicropayOrderConfirmpayurlGetResponse `json:"alipay_micropay_order_confirmpayurl_get_response"`
}





/* 1.	创建并支付一笔冻结金后，调用此接口 
2.	这个接口调用的前提是创建冻结金的参数信息pay_confirm=NO_CONFIRM */
type AlipayMicropayOrderDirectPayRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝订单号，冻结流水号.这个是创建冻结订单支付宝返回的 */
func (r *AlipayMicropayOrderDirectPayRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 支付金额,区间必须在[0.01,30]，只能保留小数点后两位 */
func (r *AlipayMicropayOrderDirectPayRequest) SetAmount(value string) {
	r.SetValue("amount", value)
}

/* 支付宝给应用的授权 */
func (r *AlipayMicropayOrderDirectPayRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 支付备注 */
func (r *AlipayMicropayOrderDirectPayRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 收款方的支付宝ID */
func (r *AlipayMicropayOrderDirectPayRequest) SetReceiveUserId(value string) {
	r.SetValue("receive_user_id", value)
}

/* 本次转账的外部单据号（只能由字母和数字组成,maxlength=32 */
func (r *AlipayMicropayOrderDirectPayRequest) SetTransferOutOrderNo(value string) {
	r.SetValue("transfer_out_order_no", value)
}


func (r *AlipayMicropayOrderDirectPayRequest) GetResponse(accessToken string) (*AlipayMicropayOrderDirectPayResponse, []byte, error) {
	var resp AlipayMicropayOrderDirectPayResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.direct.pay", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderDirectPayResponse struct {
	SinglePayDetail *SinglePayDetail `json:"single_pay_detail"`
}

type AlipayMicropayOrderDirectPayResponseResult struct {
	Response *AlipayMicropayOrderDirectPayResponse `json:"alipay_micropay_order_direct_pay_response"`
}





/* 该接口是支付宝提成冻结支付宝资金的操作 */
type AlipayMicropayOrderFreezeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要冻结金额，[0.01,2000]，必须是正数，最多只能保留小数点两位,单位是元 */
func (r *AlipayMicropayOrderFreezeRequest) SetAmount(value string) {
	r.SetValue("amount", value)
}

/* 支付宝用户给应用的授权. */
func (r *AlipayMicropayOrderFreezeRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 冻结资金的到期时间，超过此日期，冻结金会自动解冻,时间要求是:[当前时间+24h,订购时间-8h] . */
func (r *AlipayMicropayOrderFreezeRequest) SetExpireTime(value string) {
	r.SetValue("expire_time", value)
}

/* 冻结备注,maxLength=40 */
func (r *AlipayMicropayOrderFreezeRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 商户订单号,只能由字母和数字组成，最大长度32.此外部订单号与支付宝的冻结订单号对应,注意 应用id和订购者id和外部订单号必须保证唯一性。 */
func (r *AlipayMicropayOrderFreezeRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 在解冻转账的时候的支付方式: 
NO_CONFIRM：不需要付款确认，调用接口直接扣款 
PAY_PASSWORD: 
在转账需要付款方用支付密码确认，才可以转账成功 */
func (r *AlipayMicropayOrderFreezeRequest) SetPayConfirm(value string) {
	r.SetValue("pay_confirm", value)
}


func (r *AlipayMicropayOrderFreezeRequest) GetResponse(accessToken string) (*AlipayMicropayOrderFreezeResponse, []byte, error) {
	var resp AlipayMicropayOrderFreezeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.freeze", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderFreezeResponse struct {
	MicroPayOrderDetail *MicroPayOrderDetail `json:"micro_pay_order_detail"`
}

type AlipayMicropayOrderFreezeResponseResult struct {
	Response *AlipayMicropayOrderFreezeResponse `json:"alipay_micropay_order_freeze_response"`
}





/* 用户创建一笔支付订单之后,使用此接口获取支付冻结金的url来支付冻结金 */
type AlipayMicropayOrderFreezepayurlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 冻结订单号,创建冻结订单时支付宝返回的 */
func (r *AlipayMicropayOrderFreezepayurlGetRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 支付宝用户给应用的授权。 */
func (r *AlipayMicropayOrderFreezepayurlGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}


func (r *AlipayMicropayOrderFreezepayurlGetRequest) GetResponse(accessToken string) (*AlipayMicropayOrderFreezepayurlGetResponse, []byte, error) {
	var resp AlipayMicropayOrderFreezepayurlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.freezepayurl.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderFreezepayurlGetResponse struct {
	PayFreezeUrl string `json:"pay_freeze_url"`
}

type AlipayMicropayOrderFreezepayurlGetResponseResult struct {
	Response *AlipayMicropayOrderFreezepayurlGetResponse `json:"alipay_micropay_order_freezepayurl_get_response"`
}





/* 用户对已经冻结的订单可以调用此接口查询明细 */
type AlipayMicropayOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝订单号，冻结流水号(创建冻结订单返回) */
func (r *AlipayMicropayOrderGetRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 支付宝用户给应用的授权. */
func (r *AlipayMicropayOrderGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}


func (r *AlipayMicropayOrderGetRequest) GetResponse(accessToken string) (*AlipayMicropayOrderGetResponse, []byte, error) {
	var resp AlipayMicropayOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderGetResponse struct {
	MicroPayOrderDetail *MicroPayOrderDetail `json:"micro_pay_order_detail"`
}

type AlipayMicropayOrderGetResponseResult struct {
	Response *AlipayMicropayOrderGetResponse `json:"alipay_micropay_order_get_response"`
}





/* 解冻冻结单API */
type AlipayMicropayOrderUnfreezeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 冻结资金流水号,在创建资金订单时支付宝返回的流水号 */
func (r *AlipayMicropayOrderUnfreezeRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 支付宝用户给应用的授权. */
func (r *AlipayMicropayOrderUnfreezeRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 冻结备注 */
func (r *AlipayMicropayOrderUnfreezeRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}


func (r *AlipayMicropayOrderUnfreezeRequest) GetResponse(accessToken string) (*AlipayMicropayOrderUnfreezeResponse, []byte, error) {
	var resp AlipayMicropayOrderUnfreezeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.micropay.order.unfreeze", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayMicropayOrderUnfreezeResponse struct {
	UnfreezeOrderDetail *UnfreezeOrderDetail `json:"unfreeze_order_detail"`
}

type AlipayMicropayOrderUnfreezeResponseResult struct {
	Response *AlipayMicropayOrderUnfreezeResponse `json:"alipay_micropay_order_unfreeze_response"`
}





/* 查询已采购的集分宝余额，操作流程如下： 
1、申请支付宝增值包。 
2、申请支付宝应用上线时选择集分宝API。 
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。 
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID 
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。 
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。 */
type AlipayPointBudgetGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝给用户的授权。如果没有top的授权，这个字段是必填项 */
func (r *AlipayPointBudgetGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}


func (r *AlipayPointBudgetGetRequest) GetResponse(accessToken string) (*AlipayPointBudgetGetResponse, []byte, error) {
	var resp AlipayPointBudgetGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.point.budget.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayPointBudgetGetResponse struct {
	BudgetAmount int `json:"budget_amount"`
}

type AlipayPointBudgetGetResponseResult struct {
	Response *AlipayPointBudgetGetResponse `json:"alipay_point_budget_get_response"`
}





/* 向用户发送集分宝，发放集分宝之前，操作流程如下： 
1、申请支付宝增值包。 
2、申请支付宝应用上线时选择集分宝API。 
3、引导商家签约支付宝集分宝服务，地址为https://openapi.alipay.com/subscribe.htm?id=应用ID。 
4、引导商家对授予支付宝集分宝发放权限，地址为https://openauth.alipay.com/oauth2/authorize.htm?scope=p&client_id=应用ID 
5、引导授权的商家采购集分宝，地址为https://jf.alipay.com/aop/purchase.htm，进行集分宝采购。 
6、商家采购之后可以通过集分宝余额API（alipay.point.budget.get）查询商家的集分宝数量。 
7、调用发放API（alipay.point.order.add）发放集分宝。 */
type AlipayPointOrderAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝用户给应用发放集分宝的授权。 */
func (r *AlipayPointOrderAddRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 向用户展示集分宝发放备注 */
func (r *AlipayPointOrderAddRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* isv提供的发放号订单号，由数字和组成，最大长度为32为，需要保证每笔发放的唯一性，支付宝会对该参数做唯一性控制。如果使用同样的订单号，支付宝将返回订单号已经存在的错误 */
func (r *AlipayPointOrderAddRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 发放集分宝时间 */
func (r *AlipayPointOrderAddRequest) SetOrderTime(value string) {
	r.SetValue("order_time", value)
}

/* 发放集分宝的数量 */
func (r *AlipayPointOrderAddRequest) SetPointCount(value string) {
	r.SetValue("point_count", value)
}

/* 用户标识符，用于指定集分宝发放的用户，和user_symbol_type一起使用，确定一个唯一的支付宝用户 */
func (r *AlipayPointOrderAddRequest) SetUserSymbol(value string) {
	r.SetValue("user_symbol", value)
}

/* 用户标识符类型，现在支持ALIPAY_USER_ID:表示支付宝用户ID,ALIPAY_LOGON_ID:表示支付宝登陆号 */
func (r *AlipayPointOrderAddRequest) SetUserSymbolType(value string) {
	r.SetValue("user_symbol_type", value)
}


func (r *AlipayPointOrderAddRequest) GetResponse(accessToken string) (*AlipayPointOrderAddResponse, []byte, error) {
	var resp AlipayPointOrderAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.point.order.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayPointOrderAddResponse struct {
	AlipayOrderNo string `json:"alipay_order_no"`
	ResultCode bool `json:"result_code"`
}

type AlipayPointOrderAddResponseResult struct {
	Response *AlipayPointOrderAddResponse `json:"alipay_point_order_add_response"`
}





/* 查询集分宝发放详情，发放API是alipay.point.order.add。请先熟悉发放API的流程。 */
type AlipayPointOrderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝用户给应用的授权。 */
func (r *AlipayPointOrderGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* isv提供的发放号订单号，由数字和组成，最大长度为32为，需要保证每笔发放的唯一性，支付宝会对该参数做唯一性控制。如果使用同样的订单号，支付宝将返回订单号已经存在的错误 */
func (r *AlipayPointOrderGetRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 用户标识符，用于指定集分宝发放的用户，和user_symbol_type一起使用，确定一个唯一的支付宝用户 */
func (r *AlipayPointOrderGetRequest) SetUserSymbol(value string) {
	r.SetValue("user_symbol", value)
}

/* 用户标识符类型，现在支持ALIPAY_USER_ID:表示支付宝用户ID,ALIPAY_LOGON_ID:表示支付宝登陆号 */
func (r *AlipayPointOrderGetRequest) SetUserSymbolType(value string) {
	r.SetValue("user_symbol_type", value)
}


func (r *AlipayPointOrderGetRequest) GetResponse(accessToken string) (*AlipayPointOrderGetResponse, []byte, error) {
	var resp AlipayPointOrderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.point.order.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayPointOrderGetResponse struct {
	AlipayOrderNo string `json:"alipay_order_no"`
	CreateTime string `json:"create_time"`
	DispatchUserId string `json:"dispatch_user_id"`
	Memo string `json:"memo"`
	MerchantOrderNo string `json:"merchant_order_no"`
	OrderStatus string `json:"order_status"`
	PointCount int `json:"point_count"`
	ReceiveUserId string `json:"receive_user_id"`
}

type AlipayPointOrderGetResponseResult struct {
	Response *AlipayPointOrderGetResponse `json:"alipay_point_order_get_response"`
}





/* OAuth2.0 授权的第二步，换取访问令牌。 */
type AlipaySystemOauthTokenRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 授权码，用户对应用授权后得到。 */
func (r *AlipaySystemOauthTokenRequest) SetCode(value string) {
	r.SetValue("code", value)
}

/* 获取访问令牌的类型，authorization_code表示用授权码换，refresh_token表示用刷新令牌来换。 */
func (r *AlipaySystemOauthTokenRequest) SetGrantType(value string) {
	r.SetValue("grant_type", value)
}

/* 刷新令牌，上次换取访问令牌是得到。 */
func (r *AlipaySystemOauthTokenRequest) SetRefreshToken(value string) {
	r.SetValue("refresh_token", value)
}


func (r *AlipaySystemOauthTokenRequest) GetResponse(accessToken string) (*AlipaySystemOauthTokenResponse, []byte, error) {
	var resp AlipaySystemOauthTokenResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.system.oauth.token", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipaySystemOauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	AlipayUserId string `json:"alipay_user_id"`
	ExpiresIn string `json:"expires_in"`
	ReExpiresIn string `json:"re_expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type AlipaySystemOauthTokenResponseResult struct {
	Response *AlipaySystemOauthTokenResponse `json:"alipay_system_oauth_token_response"`
}





/* 查询用户支付宝账务明细接口。<br/> 
异步API使用方法，请查看：<a href="http://open.taobao.com/doc/detail.htm?id=30">异步API使用说明</a><br/> 
1. 使用之前必须首先申请支付宝API，参考http://open.taobao.com//doc/detail.htm?id=101181#s2<br/> 
2. 只对通过商家认证的、先行赔付卖家以及B2C商家开放<br/> 
3. 提交任务后，通过taobao.topats.result.get来查看任务执行状态，如果任务已完成，则返回下载URL<br/> 
4. 如果订阅了主动通知服务，任务完成后TOP会通过HTTP长连接推送消息，通知的消息格式请参考异步API使用文档<br/> 
5. 下载到的结果是CSV格式的文本文件，默认采用UTF-8编码<br/> */
type AlipayTopatsUserAccountreportGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 返回下载结果文件的数据格式，只支持utf-8和gbk编码，默认是utf-8 */
func (r *AlipayTopatsUserAccountreportGetRequest) SetCharset(value string) {
	r.SetValue("charset", value)
}

/* 对账单结束时间。end_time - start_time <= 1个自然月 */
func (r *AlipayTopatsUserAccountreportGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 需要返回的字段列表。create_time:创建时间,type：账务类型,business_type:子业务类型,balance:当时支付宝账户余额,in_amount:收入金额,out_amount:支出金额,alipay_order_no:支付宝订单号,merchant_order_no:商户订单号,self_user_id:自己的支付宝ID,opt_user_id:对方的支付宝ID,memo:账号备注 */
func (r *AlipayTopatsUserAccountreportGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 对账单开始时间。最近一个月内的日期。 */
func (r *AlipayTopatsUserAccountreportGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 账务类型。多个类型是，用逗号分隔，不传则查询所有类型的。PAYMENT:在线支付，TRANSFER:转账，DEPOSIT:充值，WITHDRAW:提现，CHARGE:收费，PREAUTH:预授权，OTHER：其它。 */
func (r *AlipayTopatsUserAccountreportGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *AlipayTopatsUserAccountreportGetRequest) GetResponse(accessToken string) (*AlipayTopatsUserAccountreportGetResponse, []byte, error) {
	var resp AlipayTopatsUserAccountreportGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.topats.user.accountreport.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayTopatsUserAccountreportGetResponse struct {
	Task *Task `json:"task"`
}

type AlipayTopatsUserAccountreportGetResponseResult struct {
	Response *AlipayTopatsUserAccountreportGetResponse `json:"alipay_topats_user_accountreport_get_response"`
}





/* 查询支付宝账户冻结类型的冻结金额。 */
type AlipayUserAccountFreezeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 冻结类型，多个用,分隔。不传返回所有类型的冻结金额。 
DEPOSIT_FREEZE,充值冻结 
WITHDRAW_FREEZE,提现冻结 
PAYMENT_FREEZE,支付冻结 
BAIL_FREEZE,保证金冻结 
CHARGE_FREEZE,收费冻结 
PRE_DEPOSIT_FREEZE,预存款冻结 
LOAN_FREEZE,贷款冻结 
OTHER_FREEZE,其它 */
func (r *AlipayUserAccountFreezeGetRequest) SetFreezeType(value string) {
	r.SetValue("freeze_type", value)
}


func (r *AlipayUserAccountFreezeGetRequest) GetResponse(accessToken string) (*AlipayUserAccountFreezeGetResponse, []byte, error) {
	var resp AlipayUserAccountFreezeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.user.account.freeze.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayUserAccountFreezeGetResponse struct {
	FreezeItems []*AccountFreeze `json:"freeze_items"`
	TotalResults string `json:"total_results"`
}

type AlipayUserAccountFreezeGetResponseResult struct {
	Response *AlipayUserAccountFreezeGetResponse `json:"alipay_user_account_freeze_get_response"`
}





/* 查询支付宝账户余额 */
type AlipayUserAccountGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *AlipayUserAccountGetRequest) GetResponse(accessToken string) (*AlipayUserAccountGetResponse, []byte, error) {
	var resp AlipayUserAccountGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.user.account.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayUserAccountGetResponse struct {
	AlipayAccount *AlipayAccount `json:"alipay_account"`
}

type AlipayUserAccountGetResponseResult struct {
	Response *AlipayUserAccountGetResponse `json:"alipay_user_account_get_response"`
}





/* 获取支付宝用户订购信息。在不确认用户对应用是否订购的时候，可以调用此API查询。 */
type AlipayUserContractGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 订购者支付宝ID。session与subscriber_user_id二选一即可。 */
func (r *AlipayUserContractGetRequest) SetSubscriberUserId(value string) {
	r.SetValue("subscriber_user_id", value)
}


func (r *AlipayUserContractGetRequest) GetResponse(accessToken string) (*AlipayUserContractGetResponse, []byte, error) {
	var resp AlipayUserContractGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.user.contract.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayUserContractGetResponse struct {
	AlipayContract *AlipayContract `json:"alipay_contract"`
}

type AlipayUserContractGetResponseResult struct {
	Response *AlipayUserContractGetResponse `json:"alipay_user_contract_get_response"`
}





/* 查询支付宝用户信息 */
type AlipayUserGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝给用户的授权。如果没有top的授权，这个字段是必填项 */
func (r *AlipayUserGetRequest) SetAuthToken(value string) {
	r.SetValue("auth_token", value)
}

/* 需要返回的字段列表。alipay_user_id：支付宝用户userId,user_status：支付宝用户状态,user_type：支付宝用户类型,certified：是否通过实名认证,real_name：真实姓名,logon_id：支付宝登录号,sex：用户性别 */
func (r *AlipayUserGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *AlipayUserGetRequest) GetResponse(accessToken string) (*AlipayUserGetResponse, []byte, error) {
	var resp AlipayUserGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.user.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayUserGetResponse struct {
	AlipayUserDetail *AlipayUserDetail `json:"alipay_user_detail"`
}

type AlipayUserGetResponseResult struct {
	Response *AlipayUserGetResponse `json:"alipay_user_get_response"`
}





/* 查询支付宝账户交易记录 */
type AlipayUserTradeSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝订单号，为空查询所有记录 */
func (r *AlipayUserTradeSearchRequest) SetAlipayOrderNo(value string) {
	r.SetValue("alipay_order_no", value)
}

/* 结束时间。与开始时间间隔在七天之内 */
func (r *AlipayUserTradeSearchRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 商户订单号，为空查询所有记录 */
func (r *AlipayUserTradeSearchRequest) SetMerchantOrderNo(value string) {
	r.SetValue("merchant_order_no", value)
}

/* 订单来源，为空查询所有来源。淘宝(TAOBAO)，支付宝(ALIPAY)，其它(OTHER) */
func (r *AlipayUserTradeSearchRequest) SetOrderFrom(value string) {
	r.SetValue("order_from", value)
}

/* 订单状态，为空查询所有状态订单 */
func (r *AlipayUserTradeSearchRequest) SetOrderStatus(value string) {
	r.SetValue("order_status", value)
}

/* 订单类型，为空查询所有类型订单。 */
func (r *AlipayUserTradeSearchRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}

/* 页码。取值范围:大于零的整数; 默认值1 */
func (r *AlipayUserTradeSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页获取条数。最大值500。 */
func (r *AlipayUserTradeSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间，时间必须是今天范围之内。格式为yyyy-MM-dd HH:mm:ss，精确到秒 */
func (r *AlipayUserTradeSearchRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *AlipayUserTradeSearchRequest) GetResponse(accessToken string) (*AlipayUserTradeSearchResponse, []byte, error) {
	var resp AlipayUserTradeSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "alipay.user.trade.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AlipayUserTradeSearchResponse struct {
	TotalPages string `json:"total_pages"`
	TotalResults string `json:"total_results"`
	TradeRecords []*TradeRecord `json:"trade_records"`
}

type AlipayUserTradeSearchResponseResult struct {
	Response *AlipayUserTradeSearchResponse `json:"alipay_user_trade_search_response"`
}



