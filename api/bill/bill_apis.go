// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package bill

import (
	"github.com/changkong/open_taobao"
)

/* 查询费用账户信息 */
type BillAccountsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需要获取的科目ID */
func (r *BillAccountsGetRequest) SetAids(value string) {
	r.SetValue("aids", value)
}

/* 需要返回的字段 */
func (r *BillAccountsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

func (r *BillAccountsGetRequest) GetResponse(accessToken string) (*BillAccountsGetResponse, []byte, error) {
	var resp BillAccountsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.bill.accounts.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type BillAccountsGetResponse struct {
	Accounts     []*Account `json:"accounts"`
	TotalResults int        `json:"total_results"`
}

type BillAccountsGetResponseResult struct {
	Response *BillAccountsGetResponse `json:"bill_accounts_get_response"`
}

/* 查询单笔账单明细 */
type BillBillGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 账单编号 */
func (r *BillBillGetRequest) SetBid(value string) {
	r.SetValue("bid", value)
}

/* 传入需要返回的字段 */
func (r *BillBillGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

func (r *BillBillGetRequest) GetResponse(accessToken string) (*BillBillGetResponse, []byte, error) {
	var resp BillBillGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.bill.bill.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type BillBillGetResponse struct {
	Bill *Bill `json:"bill"`
}

type BillBillGetResponseResult struct {
	Response *BillBillGetResponse `json:"bill_bill_get_response"`
}

/* 查询账单明细数据 */
type BillBillsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 科目编号 */
func (r *BillBillsGetRequest) SetAccountId(value string) {
	r.SetValue("account_id", value)
}

/* 结束时间,限制:结束时间-开始时间不能大于1天(根据order_id或者trade_id查询除外) */
func (r *BillBillsGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 传入需要返回的字段,参见Bill结构体 */
func (r *BillBillsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 子订单编号 */
func (r *BillBillsGetRequest) SetOrderId(value string) {
	r.SetValue("order_id", value)
}

/* 页数,建议不要超过100页,越大性能越低,有可能会超时 */
func (r *BillBillsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小,默认40条,可选范围 ：40~100 */
func (r *BillBillsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间 */
func (r *BillBillsGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 查询条件中的时间类型:1-交易订单完成时间biz_time 2-支付宝扣款时间pay_time 如果不填默认为2即根据支付时间查询,查询的结果会根据该时间倒排序 */
func (r *BillBillsGetRequest) SetTimeType(value string) {
	r.SetValue("time_type", value)
}

/* 交易编号 */
func (r *BillBillsGetRequest) SetTradeId(value string) {
	r.SetValue("trade_id", value)
}

func (r *BillBillsGetRequest) GetResponse(accessToken string) (*BillBillsGetResponse, []byte, error) {
	var resp BillBillsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.bill.bills.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type BillBillsGetResponse struct {
	Bills        []*Bill `json:"bills"`
	HasNext      bool    `json:"has_next"`
	TotalResults int     `json:"total_results"`
}

type BillBillsGetResponseResult struct {
	Response *BillBillsGetResponse `json:"bill_bills_get_response"`
}

/* 查询单笔虚拟账户明细 */
type BillBookBillGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 虚拟账户流水编号 */
func (r *BillBookBillGetRequest) SetBid(value string) {
	r.SetValue("bid", value)
}

/* 需要返回的字段:参见BookBill结构体 */
func (r *BillBookBillGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

func (r *BillBookBillGetRequest) GetResponse(accessToken string) (*BillBookBillGetResponse, []byte, error) {
	var resp BillBookBillGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.bill.book.bill.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type BillBookBillGetResponse struct {
	Bookbill *BookBill `json:"bookbill"`
}

type BillBookBillGetResponseResult struct {
	Response *BillBookBillGetResponse `json:"bill_book_bill_get_response"`
}

/* 查询虚拟账户明细数据 */
type BillBookBillsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 虚拟账户科目编号 */
func (r *BillBookBillsGetRequest) SetAccountId(value string) {
	r.SetValue("account_id", value)
}

/* 记账结束时间,end_time与start_time相差不能超过30天 */
func (r *BillBookBillsGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 需要返回的字段:bid,account_id,journal_type,amount,book_time,description,gmt_create,gmt_modified ,如果不是以上字段将自动忽略 */
func (r *BillBookBillsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 明细流水类型:流水类型:101、可用金充值；102、可用金扣除；103、冻结；104、解冻；105、冻结金充值；106、冻结金扣除 */
func (r *BillBookBillsGetRequest) SetJournalTypes(value string) {
	r.SetValue("journal_types", value)
}

/* 页码,传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始 */
func (r *BillBookBillsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小,建议40~100,不能超过100 */
func (r *BillBookBillsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 记账开始时间 */
func (r *BillBookBillsGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *BillBookBillsGetRequest) GetResponse(accessToken string) (*BillBookBillsGetResponse, []byte, error) {
	var resp BillBookBillsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.bill.book.bills.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type BillBookBillsGetResponse struct {
	Bills        []*BookBill `json:"bills"`
	HasNext      bool        `json:"has_next"`
	TotalResults int         `json:"total_results"`
}

type BillBookBillsGetResponseResult struct {
	Response *BillBookBillsGetResponse `json:"bill_book_bills_get_response"`
}
