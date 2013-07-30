// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package refund

import (
	"github.com/changkong/open_taobao"
)

/* 获取单笔退款详情 */
type RefundGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段。目前支持有：refund_id, alipay_no, tid, oid, buyer_nick, seller_nick, total_fee, status, created, refund_fee, good_status, has_good_return, payment, reason, desc, num_iid, title, price, num, good_return_time, company_name, sid, address, shipping_type, refund_remind_timeout */
func (r *RefundGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 退款单号 */
func (r *RefundGetRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

func (r *RefundGetRequest) GetResponse(accessToken string) (*RefundGetResponse, []byte, error) {
	var resp RefundGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refund.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundGetResponse struct {
	Refund *Refund `json:"refund"`
}

type RefundGetResponseResult struct {
	Response *RefundGetResponse `json:"refund_get_response"`
}

/* 创建退款留言/凭证 */
type RefundMessageAddRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 留言内容。最大长度: 400个字节 */
func (r *RefundMessageAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K */
func (r *RefundMessageAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 退款编号。 */
func (r *RefundMessageAddRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

func (r *RefundMessageAddRequest) GetResponse(accessToken string) (*RefundMessageAddResponse, []byte, error) {
	var resp RefundMessageAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refund.message.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundMessageAddResponse struct {
	RefundMessage *RefundMessage `json:"refund_message"`
}

type RefundMessageAddResponseResult struct {
	Response *RefundMessageAddResponse `json:"refund_message_add_response"`
}

/* 单笔退款详情 */
type RefundMessagesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表。可选值：RefundMessage结构体中的所有字段，以半角逗号(,)分隔。 */
func (r *RefundMessagesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *RefundMessagesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
func (r *RefundMessagesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 退款单号 */
func (r *RefundMessagesGetRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

func (r *RefundMessagesGetRequest) GetResponse(accessToken string) (*RefundMessagesGetResponse, []byte, error) {
	var resp RefundMessagesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refund.messages.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundMessagesGetResponse struct {
	RefundMessages []*RefundMessage `json:"refund_messages"`
	TotalResults   int              `json:"total_results"`
}

type RefundMessagesGetResponseResult struct {
	Response *RefundMessagesGetResponse `json:"refund_messages_get_response"`
}

/* 卖家拒绝单笔退款交易，要求如下：
1. 传入的refund_id和相应的tid, oid必须匹配
2. 如果一笔订单只有一笔子订单，则tid必须与oid相同
3. 只有卖家才能执行拒绝退款操作
4. 以下三种情况不能退款：卖家未发货；7天无理由退换货；网游订单 */
type RefundRefuseRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 退款记录对应的交易子订单号 */
func (r *RefundRefuseRequest) SetOid(value string) {
	r.SetValue("oid", value)
}

/* 退款单号 */
func (r *RefundRefuseRequest) SetRefundId(value string) {
	r.SetValue("refund_id", value)
}

/* 拒绝退款时的说明信息，长度2-200 */
func (r *RefundRefuseRequest) SetRefuseMessage(value string) {
	r.SetValue("refuse_message", value)
}

/* 拒绝退款时的退款凭证，一般是卖家拒绝退款时使用的发货凭证，最大长度130000字节，支持的图片格式：GIF, JPG, PNG */
func (r *RefundRefuseRequest) SetRefuseProof(value string) {
	r.SetValue("refuse_proof", value)
}

/* 退款记录对应的交易订单号 */
func (r *RefundRefuseRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

func (r *RefundRefuseRequest) GetResponse(accessToken string) (*RefundRefuseResponse, []byte, error) {
	var resp RefundRefuseResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refund.refuse", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundRefuseResponse struct {
	Refund *Refund `json:"refund"`
}

type RefundRefuseResponseResult struct {
	Response *RefundRefuseResponse `json:"refund_refuse_response"`
}

/* 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型 */
type RefundsApplyGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee */
func (r *RefundsApplyGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。传入值为 1 代表第一页，传入值为 2 代表第二页，依此类推。默认返回的数据是从第一页开始 */
func (r *RefundsApplyGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
func (r *RefundsApplyGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 卖家昵称 */
func (r *RefundsApplyGetRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}

/* 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意)
WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货)
WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货)
SELLER_REFUSE_BUYER(卖家拒绝退款)
CLOSED(退款关闭)
SUCCESS(退款成功) */
func (r *RefundsApplyGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
fixed(一口价)
auction(拍卖)
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲)
cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易) */
func (r *RefundsApplyGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

func (r *RefundsApplyGetRequest) GetResponse(accessToken string) (*RefundsApplyGetResponse, []byte, error) {
	var resp RefundsApplyGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refunds.apply.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundsApplyGetResponse struct {
	Refunds      []*Refund `json:"refunds"`
	TotalResults int       `json:"total_results"`
}

type RefundsApplyGetResponseResult struct {
	Response *RefundsApplyGetResponse `json:"refunds_apply_get_response"`
}

/* 查询卖家收到的退款列表，查询外店的退款列表时需要指定交易类型 */
type RefundsReceiveGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家昵称 */
func (r *RefundsReceiveGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 查询修改时间结束。格式: yyyy-MM-dd HH:mm:ss */
func (r *RefundsReceiveGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 需要返回的字段。目前支持有：refund_id, tid, title, buyer_nick, seller_nick, total_fee, status, created, refund_fee, oid, good_status, company_name, sid, payment, reason, desc, has_good_return, modified, order_status */
func (r *RefundsReceiveGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 页码。取值范围:大于零的整数; 默认值:1 */
func (r *RefundsReceiveGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数; 默认值:40;最大值:100 */
func (r *RefundsReceiveGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询修改时间开始。格式: yyyy-MM-dd HH:mm:ss */
func (r *RefundsReceiveGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}

/* 退款状态，默认查询所有退款状态的数据，除了默认值外每次只能查询一种状态。
WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意)
WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货)
WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货)
SELLER_REFUSE_BUYER(卖家拒绝退款)
CLOSED(退款关闭)
SUCCESS(退款成功) */
func (r *RefundsReceiveGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易类型列表，一次查询多种类型可用半角逗号分隔，默认同时查询guarantee_trade, auto_delivery的2种类型的数据。
fixed(一口价)
auction(拍卖)
guarantee_trade(一口价、拍卖)
independent_simple_trade(旺店入门版交易)
independent_shop_trade(旺店标准版交易)
auto_delivery(自动发货)
ec(直冲)
cod(货到付款)
fenxiao(分销)
game_equipment(游戏装备)
shopex_trade(ShopEX交易)
netcn_trade(万网交易)
external_trade(统一外部交易) */
func (r *RefundsReceiveGetRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 是否启用has_next的分页方式，如果指定true,则返回的结果中不包含总记录数，但是会新增一个是否存在下一页的的字段，通过此种方式获取增量退款，接口调用成功率在原有的基础上有所提升。 */
func (r *RefundsReceiveGetRequest) SetUseHasNext(value string) {
	r.SetValue("use_has_next", value)
}

func (r *RefundsReceiveGetRequest) GetResponse(accessToken string) (*RefundsReceiveGetResponse, []byte, error) {
	var resp RefundsReceiveGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.refunds.receive.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type RefundsReceiveGetResponse struct {
	HasNext      bool      `json:"has_next"`
	Refunds      []*Refund `json:"refunds"`
	TotalResults int       `json:"total_results"`
}

type RefundsReceiveGetResponseResult struct {
	Response *RefundsReceiveGetResponse `json:"refunds_receive_get_response"`
}

/* 提供异步获取三个月内买家收到的退款详情信息接口。 */
type TopatsRefundsReceiveGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 退款修改结束时间，格式yyyyMMdd，取值范围：前90天内~昨天，其中start_time<=end_time，如20120531相当于取退款修改时间到2012-05-31 23:59:59为止的退款。注：如果start_time和end_time相同，表示取一天的退款数据。<span style="color:red">强烈建议直充类卖家把三个月退款拆分成3次来获取，否则单个任务会消耗很长时间。<span> */
func (r *TopatsRefundsReceiveGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* Refund结构体中的所有字段。<span style="color:red">请尽量按需获取，如果只取refund_id和modified字段，获取退款数据速度会超快。</span> */
func (r *TopatsRefundsReceiveGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 退款修改开始时间，格式yyyyMMdd，取值范围：前90天内~昨天。如：20120501相当于取退款修改时间从2012-05-01 00:00:00开始的退款。 */
func (r *TopatsRefundsReceiveGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TopatsRefundsReceiveGetRequest) GetResponse(accessToken string) (*TopatsRefundsReceiveGetResponse, []byte, error) {
	var resp TopatsRefundsReceiveGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.refunds.receive.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsRefundsReceiveGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsRefundsReceiveGetResponseResult struct {
	Response *TopatsRefundsReceiveGetResponse `json:"topats_refunds_receive_get_response"`
}
