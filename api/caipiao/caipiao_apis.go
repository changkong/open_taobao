// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caipiao

import (
	"github.com/yaofangou/open_taobao"
)





/* 根据卖家id与appkey获取商品信息。 */
type CaipiaoGoodsInfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *CaipiaoGoodsInfoGetRequest) GetResponse(accessToken string) (*CaipiaoGoodsInfoGetResponse, []byte, error) {
	var resp CaipiaoGoodsInfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.goods.info.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoGoodsInfoGetResponse struct {
	Results []*LotteryWangcaiSellerGoodsInfo `json:"results"`
	TotalResults int `json:"total_results"`
}

type CaipiaoGoodsInfoGetResponseResult struct {
	Response *CaipiaoGoodsInfoGetResponse `json:"caipiao_goods_info_get_response"`
}





/* 录入参加送彩票商品信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据商品id与赠送类型（goodsid_presenttype_uk）来决定是新增数据还是修改数据。 */
type CaipiaoGoodsInfoInputRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetActEndDate(value string) {
	r.SetValue("act_end_date", value)
}

/* 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetActStartDate(value string) {
	r.SetValue("act_start_date", value)
}

/* 店铺相关商品参加的送彩票活动描述 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsDesc(value string) {
	r.SetValue("goods_desc", value)
}

/* 商品在淘宝的唯一id，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsId(value string) {
	r.SetValue("goods_id", value)
}

/* 商品主图地址 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsImage(value string) {
	r.SetValue("goods_image", value)
}

/* 商品价格,保留两位小数，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsPrice(value string) {
	r.SetValue("goods_price", value)
}

/* 商品标题 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsTitle(value string) {
	r.SetValue("goods_title", value)
}

/* 商品类目编号，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetGoodsType(value string) {
	r.SetValue("goods_type", value)
}

/* 彩种id,不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetLotteryTypeId(value string) {
	r.SetValue("lottery_type_id", value)
}

/* 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空 */
func (r *CaipiaoGoodsInfoInputRequest) SetPresentType(value string) {
	r.SetValue("present_type", value)
}


func (r *CaipiaoGoodsInfoInputRequest) GetResponse(accessToken string) (*CaipiaoGoodsInfoInputResponse, []byte, error) {
	var resp CaipiaoGoodsInfoInputResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.goods.info.input", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoGoodsInfoInputResponse struct {
	InputResult bool `json:"input_result"`
}

type CaipiaoGoodsInfoInputResponseResult struct {
	Response *CaipiaoGoodsInfoInputResponse `json:"caipiao_goods_info_input_response"`
}





/* 卖家给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。 */
type CaipiaoLotterySendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 彩票接收方淘宝数字ID，不可为空、0和负数。 */
func (r *CaipiaoLotterySendRequest) SetBuyerNumId(value string) {
	r.SetValue("buyer_num_id", value)
}

/* 彩种ID，此彩种ID为彩票系统中的序号。 */
func (r *CaipiaoLotterySendRequest) SetLotteryTypeId(value string) {
	r.SetValue("lottery_type_id", value)
}

/* 送彩方淘宝数字ID， 不可为空、0和负数。 */
func (r *CaipiaoLotterySendRequest) SetSellerNumId(value string) {
	r.SetValue("seller_num_id", value)
}

/* 彩票注数，不可为空、0和负数，最大值为100。 */
func (r *CaipiaoLotterySendRequest) SetStakeCount(value string) {
	r.SetValue("stake_count", value)
}

/* 送彩票给接收方的赠言。 不能超过50个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。 */
func (r *CaipiaoLotterySendRequest) SetSweetyWords(value string) {
	r.SetValue("sweety_words", value)
}


func (r *CaipiaoLotterySendRequest) GetResponse(accessToken string) (*CaipiaoLotterySendResponse, []byte, error) {
	var resp CaipiaoLotterySendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.lottery.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoLotterySendResponse struct {
	SendResult bool `json:"send_result"`
}

type CaipiaoLotterySendResponseResult struct {
	Response *CaipiaoLotterySendResponse `json:"caipiao_lottery_send_response"`
}





/* 卖家使用nick给买家送彩票，可以指定彩种和注数。赠送成功，返回true; 以下几种情况情况， 返回false: 注数超过100注、卖家未签署支付宝代扣协议、卖家或者买家信息不存在等。 */
type CaipiaoLotterySendbynickRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 彩票接收方nick， 不可为空、""。 */
func (r *CaipiaoLotterySendbynickRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 彩种ID，此彩种ID为彩票系统中的序号。 */
func (r *CaipiaoLotterySendbynickRequest) SetLotteryTypeId(value string) {
	r.SetValue("lottery_type_id", value)
}

/* 彩票注数，不可为空、0和负数，最大值为100。 */
func (r *CaipiaoLotterySendbynickRequest) SetStakeCount(value string) {
	r.SetValue("stake_count", value)
}

/* 送彩票给接收方的赠言。 不能超过20个字符，1个中文字符、1个英文字母及1个数字等均当作一个字符，如果超过，则会截取。 */
func (r *CaipiaoLotterySendbynickRequest) SetSweetyWords(value string) {
	r.SetValue("sweety_words", value)
}


func (r *CaipiaoLotterySendbynickRequest) GetResponse(accessToken string) (*CaipiaoLotterySendbynickResponse, []byte, error) {
	var resp CaipiaoLotterySendbynickResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.lottery.sendbynick", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoLotterySendbynickResponse struct {
	SendResult bool `json:"send_result"`
}

type CaipiaoLotterySendbynickResponseResult struct {
	Response *CaipiaoLotterySendbynickResponse `json:"caipiao_lottery_sendbynick_response"`
}





/* 获取彩票系统支持的可用于赠送的彩种列表 */
type CaipiaoLotterytypesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *CaipiaoLotterytypesGetRequest) GetResponse(accessToken string) (*CaipiaoLotterytypesGetResponse, []byte, error) {
	var resp CaipiaoLotterytypesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.lotterytypes.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoLotterytypesGetResponse struct {
	Results []*LotteryType `json:"results"`
	TotalResults int `json:"total_results"`
}

type CaipiaoLotterytypesGetResponseResult struct {
	Response *CaipiaoLotterytypesGetResponse `json:"caipiao_lotterytypes_get_response"`
}





/* 查询卖家指定时间范围内的赠送订单列表, 只查询3个月以内的数据 */
type CaipiaoPresentItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 赠送订单的截止时间，格式为yyyyMMdd, 距当前最长时间间隔是3个月，最近可以取当天的时间。不可为空。要求endDate-startDate必须<=3个月， */
func (r *CaipiaoPresentItemsGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 当前页码， 不可为空、0和负数。 */
func (r *CaipiaoPresentItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页的大小，不可为空、0和负数。最大为500，如果超过500，则取默认的500. */
func (r *CaipiaoPresentItemsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 赠送订单的开始时间， 格式为yyyyMMdd, 距当前最长时间间隔是3个月， 最近可以取当天的时间. 不可为空。 */
func (r *CaipiaoPresentItemsGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}


func (r *CaipiaoPresentItemsGetRequest) GetResponse(accessToken string) (*CaipiaoPresentItemsGetResponse, []byte, error) {
	var resp CaipiaoPresentItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.present.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoPresentItemsGetResponse struct {
	Results []*LotteryWangcaiPresent `json:"results"`
	TotalResults int `json:"total_results"`
}

type CaipiaoPresentItemsGetResponseResult struct {
	Response *CaipiaoPresentItemsGetResponse `json:"caipiao_present_items_get_response"`
}





/* 查询卖家一段时间内按天统计的彩票赠送数据，只支持查询90天以内的数据. */
type CaipiaoPresentStatGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 指定查询的天数，从当前日期（不包括当前日期）向前推算的天数，可为空。如果为空、0、负数或者大于90天，则设置为默认的90天。举例：当天是20120703， days=2， 则统计数据的日期为：20120702，20120701. */
func (r *CaipiaoPresentStatGetRequest) SetDays(value string) {
	r.SetValue("days", value)
}


func (r *CaipiaoPresentStatGetRequest) GetResponse(accessToken string) (*CaipiaoPresentStatGetResponse, []byte, error) {
	var resp CaipiaoPresentStatGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.present.stat.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoPresentStatGetResponse struct {
	Results []*LotteryWangcaiPresentStat `json:"results"`
	TotalResults int `json:"total_results"`
}

type CaipiaoPresentStatGetResponseResult struct {
	Response *CaipiaoPresentStatGetResponse `json:"caipiao_present_stat_get_response"`
}





/* 获取卖家赠送的中奖订单列表(90天以内)，支持两种方式查询： 
1.按num查询中奖订单，不支持分页。 
2.按date查询中奖订单或所有订单(以searchType区分)，支持分页。 
注意num与date只要传入一个即可，如果两个参数都传会以date方式查询。 */
type CaipiaoPresentWinItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询日期，格式请严格遵守yyyy-MM-dd */
func (r *CaipiaoPresentWinItemsGetRequest) SetDate(value string) {
	r.SetValue("date", value)
}

/* 查询个数，最大值为500.如果为空、0和负数，则取默认值500 */
func (r *CaipiaoPresentWinItemsGetRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 查询页码，空，零，负的情况默认为1（注意每页数据量为50） */
func (r *CaipiaoPresentWinItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 0：查询中奖订单，1：查询所有订单，默认为0，注意按列表数量查询只会查询中奖订单。 */
func (r *CaipiaoPresentWinItemsGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}


func (r *CaipiaoPresentWinItemsGetRequest) GetResponse(accessToken string) (*CaipiaoPresentWinItemsGetResponse, []byte, error) {
	var resp CaipiaoPresentWinItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.present.win.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoPresentWinItemsGetResponse struct {
	Results []*LotteryWangcaiPresent `json:"results"`
	TotalResults int `json:"total_results"`
}

type CaipiaoPresentWinItemsGetResponseResult struct {
	Response *CaipiaoPresentWinItemsGetResponse `json:"caipiao_present_win_items_get_response"`
}





/* 录入参加送彩票店铺信息，如果录入成功，返回true，如果录入失败，返回false，后端会根据卖家id与赠送类型（sellerid_presenttype_uk）来决定是新增数据还是修改数据。 */
type CaipiaoShopInfoInputRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动结束时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空 */
func (r *CaipiaoShopInfoInputRequest) SetActEndDate(value string) {
	r.SetValue("act_end_date", value)
}

/* 活动开始时间，格式需严格遵守yyyy-MM-dd HH:mm:ss，不可为空 */
func (r *CaipiaoShopInfoInputRequest) SetActStartDate(value string) {
	r.SetValue("act_start_date", value)
}

/* 赠送类型：0-满就送；1-好评送；2-分享送；3-游戏送；4-收藏送，不可为空 */
func (r *CaipiaoShopInfoInputRequest) SetPresentType(value string) {
	r.SetValue("present_type", value)
}

/* 店铺参加的送彩票活动描述 */
func (r *CaipiaoShopInfoInputRequest) SetShopDesc(value string) {
	r.SetValue("shop_desc", value)
}

/* 店铺名称 */
func (r *CaipiaoShopInfoInputRequest) SetShopName(value string) {
	r.SetValue("shop_name", value)
}

/* 店铺类目编号，不可为空 */
func (r *CaipiaoShopInfoInputRequest) SetShopType(value string) {
	r.SetValue("shop_type", value)
}


func (r *CaipiaoShopInfoInputRequest) GetResponse(accessToken string) (*CaipiaoShopInfoInputResponse, []byte, error) {
	var resp CaipiaoShopInfoInputResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.shop.info.input", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoShopInfoInputResponse struct {
	InputResult bool `json:"input_result"`
}

type CaipiaoShopInfoInputResponseResult struct {
	Response *CaipiaoShopInfoInputResponse `json:"caipiao_shop_info_input_response"`
}





/* 检查用户是否签署了支付宝代扣协议。如果签署了，返回true; 如果没签署，返回false, 同时返回签署代扣协议的Url。 */
type CaipiaoSignstatusCheckRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *CaipiaoSignstatusCheckRequest) GetResponse(accessToken string) (*CaipiaoSignstatusCheckResponse, []byte, error) {
	var resp CaipiaoSignstatusCheckResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.caipiao.signstatus.check", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type CaipiaoSignstatusCheckResponse struct {
	Sign bool `json:"sign"`
	SignUrl string `json:"sign_url"`
}

type CaipiaoSignstatusCheckResponseResult struct {
	Response *CaipiaoSignstatusCheckResponse `json:"caipiao_signstatus_check_response"`
}



