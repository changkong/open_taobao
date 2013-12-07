// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package marketing

import (
	"github.com/yaofangou/open_taobao"
)





/* 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。 */
type UmpPromotionGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品id */
func (r *UmpPromotionGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *UmpPromotionGetRequest) GetResponse(accessToken string) (*UmpPromotionGetResponse, []byte, error) {
	var resp UmpPromotionGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.promotion.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpPromotionGetResponse struct {
	Promotions *PromotionDisplayTop `json:"promotions"`
}

type UmpPromotionGetResponseResult struct {
	Response *UmpPromotionGetResponse `json:"ump_promotion_get_response"`
}





/* 活动名称与描述违禁词检查 */
type MarketingPromotionKfcRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动描述 */
func (r *MarketingPromotionKfcRequest) SetPromotionDesc(value string) {
	r.SetValue("promotion_desc", value)
}

/* 活动名称 */
func (r *MarketingPromotionKfcRequest) SetPromotionTitle(value string) {
	r.SetValue("promotion_title", value)
}


func (r *MarketingPromotionKfcRequest) GetResponse(accessToken string) (*MarketingPromotionKfcResponse, []byte, error) {
	var resp MarketingPromotionKfcResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.marketing.promotion.kfc", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type MarketingPromotionKfcResponse struct {
	IsSuccess bool `json:"is_success"`
}

type MarketingPromotionKfcResponseResult struct {
	Response *MarketingPromotionKfcResponse `json:"marketing_promotion_kfc_response"`
}





/* 根据商品ID查询卖家使用该第三方工具对商品设置的所有优惠策略 */
type MarketingPromotionsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的优惠策略结构字段列表。可选值为Promotion中所有字段，如：promotion_id, promotion_title, item_id, status, tag_id等等 */
func (r *MarketingPromotionsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否新标签标识 */
func (r *MarketingPromotionsGetRequest) SetIsNewTag(value string) {
	r.SetValue("is_new_tag", value)
}

/* 商品数字ID。根据该ID查询商品下通过第三方工具设置的所有优惠策略 */
func (r *MarketingPromotionsGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 优惠策略状态。可选值：ACTIVE(有效)，UNACTIVE(无效)，若不传或者传入其他值，则默认查询全部 */
func (r *MarketingPromotionsGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 标签ID */
func (r *MarketingPromotionsGetRequest) SetTagId(value string) {
	r.SetValue("tag_id", value)
}


func (r *MarketingPromotionsGetRequest) GetResponse(accessToken string) (*MarketingPromotionsGetResponse, []byte, error) {
	var resp MarketingPromotionsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.marketing.promotions.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type MarketingPromotionsGetResponse struct {
	Promotions *PromotionListObject `json:"promotions"`
	TotalResults int `json:"total_results"`
}

type MarketingPromotionsGetResponseResult struct {
	Response *MarketingPromotionsGetResponse `json:"marketing_promotions_get_response"`
}





/* 创建某个卖家的店铺优惠券领取活动。返回，优惠券领取活动ID，优惠券领取链接。 */
type PromotionActivityAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 优惠券总领用数量 */
func (r *PromotionActivityAddRequest) SetCouponCount(value string) {
	r.SetValue("coupon_count", value)
}

/* 优惠券的id，唯一标识这个优惠券 */
func (r *PromotionActivityAddRequest) SetCouponId(value string) {
	r.SetValue("coupon_id", value)
}

/* 每个人最多领用数量，0代表不限 */
func (r *PromotionActivityAddRequest) SetPersonLimitCount(value string) {
	r.SetValue("person_limit_count", value)
}

/* 是否将该活动优惠券同步到淘券市场 
1（不同步） 
2（同步） */
func (r *PromotionActivityAddRequest) SetUploadToTaoquan(value string) {
	r.SetValue("upload_to_taoquan", value)
}


func (r *PromotionActivityAddRequest) GetResponse(accessToken string) (*PromotionActivityAddResponse, []byte, error) {
	var resp PromotionActivityAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.activity.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionActivityAddResponse struct {
	Activity *Activity `json:"activity"`
}

type PromotionActivityAddResponseResult struct {
	Response *PromotionActivityAddResponse `json:"promotion_activity_add_response"`
}





/* 取消正在开始的店铺优惠券活动 */
type PromotionActivityCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *PromotionActivityCancelRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}


func (r *PromotionActivityCancelRequest) GetResponse(accessToken string) (*PromotionActivityCancelResponse, []byte, error) {
	var resp PromotionActivityCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.activity.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionActivityCancelResponse struct {
	IsSuccess bool `json:"is_success"`
}

type PromotionActivityCancelResponseResult struct {
	Response *PromotionActivityCancelResponse `json:"promotion_activity_cancel_response"`
}





/* 删除卖家创建的店铺优惠券活动 */
type PromotionActivityDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 优惠券的id */
func (r *PromotionActivityDeleteRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}


func (r *PromotionActivityDeleteRequest) GetResponse(accessToken string) (*PromotionActivityDeleteResponse, []byte, error) {
	var resp PromotionActivityDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.activity.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionActivityDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type PromotionActivityDeleteResponseResult struct {
	Response *PromotionActivityDeleteResponse `json:"promotion_activity_delete_response"`
}





/* 查询某个卖家的店铺优惠券领取活动 
返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量 
领取活动状态，优惠券领取链接 
最多50个优惠券 */
type PromotionActivityGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动的id */
func (r *PromotionActivityGetRequest) SetActivityId(value string) {
	r.SetValue("activity_id", value)
}


func (r *PromotionActivityGetRequest) GetResponse(accessToken string) (*PromotionActivityGetResponse, []byte, error) {
	var resp PromotionActivityGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.activity.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionActivityGetResponse struct {
	Activitys *ActivityListObject `json:"activitys"`
}

type PromotionActivityGetResponseResult struct {
	Response *PromotionActivityGetResponse `json:"promotion_activity_get_response"`
}





/* 创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张 */
type PromotionCouponAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 订单满多少元才能用这个优惠券，500就是满500元才能使用 */
func (r *PromotionCouponAddRequest) SetCondition(value string) {
	r.SetValue("condition", value)
}

/* 优惠券的面额，必须是3，5，10，20，50，100 */
func (r *PromotionCouponAddRequest) SetDenominations(value string) {
	r.SetValue("denominations", value)
}

/* 优惠券的截止日期 */
func (r *PromotionCouponAddRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 优惠券的生效时间 */
func (r *PromotionCouponAddRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *PromotionCouponAddRequest) GetResponse(accessToken string) (*PromotionCouponAddResponse, []byte, error) {
	var resp PromotionCouponAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.coupon.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionCouponAddResponse struct {
	CouponId int `json:"coupon_id"`
}

type PromotionCouponAddResponseResult struct {
	Response *PromotionCouponAddResponse `json:"promotion_coupon_add_response"`
}





/* 通过接口批量发放店铺优惠券（每次只能发送100张，只能发给当前授权卖家店铺的会员），发送成功则返回为空，发送失败则返回失败的买家列表和发送成功的买家和优惠券的number。注：如果所有买家都发放失败的话，is_success也为true，建议调用者根据返回的集合判断是否送入的买家都发放成功了 */
type PromotionCouponSendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称用半角','号分割 */
func (r *PromotionCouponSendRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 优惠券的id */
func (r *PromotionCouponSendRequest) SetCouponId(value string) {
	r.SetValue("coupon_id", value)
}


func (r *PromotionCouponSendRequest) GetResponse(accessToken string) (*PromotionCouponSendResponse, []byte, error) {
	var resp PromotionCouponSendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.coupon.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionCouponSendResponse struct {
	CouponResults *CouponResultListObject `json:"coupon_results"`
	FailureBuyers *ErrorMessageListObject `json:"failure_buyers"`
	IsSuccess bool `json:"is_success"`
}

type PromotionCouponSendResponseResult struct {
	Response *PromotionCouponSendResponse `json:"promotion_coupon_send_response"`
}





/* 买家可以将自己“未使用”状态下的优惠券转发给任意淘宝用户。 
只能转发“未使用”状态下的优惠券 */
type PromotionCouponTransferRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 优惠券编号 */
func (r *PromotionCouponTransferRequest) SetCouponNumber(value string) {
	r.SetValue("coupon_number", value)
}

/* 要赠送的淘宝昵称 */
func (r *PromotionCouponTransferRequest) SetReceiveingBuyerName(value string) {
	r.SetValue("receiveing_buyer_name", value)
}


func (r *PromotionCouponTransferRequest) GetResponse(accessToken string) (*PromotionCouponTransferResponse, []byte, error) {
	var resp PromotionCouponTransferResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.coupon.transfer", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionCouponTransferResponse struct {
	IsSuccess bool `json:"is_success"`
}

type PromotionCouponTransferResponseResult struct {
	Response *PromotionCouponTransferResponse `json:"promotion_coupon_transfer_response"`
}





/* 通过接口可以查询某个店铺优惠券的买家详细信息返回的信息，买家昵称， 使用渠道，使用状态，总数量 */
type PromotionCoupondetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *PromotionCoupondetailGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 优惠券的id */
func (r *PromotionCoupondetailGetRequest) SetCouponId(value string) {
	r.SetValue("coupon_id", value)
}

/* 查N天内的数据，N<=15 */
func (r *PromotionCoupondetailGetRequest) SetDays(value string) {
	r.SetValue("days", value)
}

/* 传入优惠券截止时间，即失效时间。查询输入日期向前1天的数据；不传则查询当前日期向前1天的数据。比如查询明天才失效的优惠卷，要传入明天之后1天内的日期，才能查询到该优惠卷。 */
func (r *PromotionCoupondetailGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 这是一个扩展字段 供版本升级用 
当前如果新版本的话 可以传入new字符串 */
func (r *PromotionCoupondetailGetRequest) SetExtendParams(value string) {
	r.SetValue("extend_params", value)
}

/* 查询的页号，结果集是分页返回的，每页20-100条 */
func (r *PromotionCoupondetailGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页行数 */
func (r *PromotionCoupondetailGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 优惠券使用情况unused：代表未使用using：代表使用中used：代表已使用。必须是unused，using，used */
func (r *PromotionCoupondetailGetRequest) SetState(value string) {
	r.SetValue("state", value)
}


func (r *PromotionCoupondetailGetRequest) GetResponse(accessToken string) (*PromotionCoupondetailGetResponse, []byte, error) {
	var resp PromotionCoupondetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.coupondetail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionCoupondetailGetResponse struct {
	CouponDetails *CouponDetailListObject `json:"coupon_details"`
	IsHaveNextPage bool `json:"is_have_next_page"`
	TotalResults int `json:"total_results"`
}

type PromotionCoupondetailGetResponseResult struct {
	Response *PromotionCoupondetailGetResponse `json:"promotion_coupondetail_get_response"`
}





/* 查询卖家已经创建的优惠券，接口返回信息：优惠券ID，面值，创建时间，有效期，使用条件，使用渠道，创建渠道，优惠券总数量 */
type PromotionCouponsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 优惠券的id，唯一标识这个优惠券 */
func (r *PromotionCouponsGetRequest) SetCouponId(value string) {
	r.SetValue("coupon_id", value)
}

/* 优惠券的面额，必须是3，5，10，20，50,100 */
func (r *PromotionCouponsGetRequest) SetDenominations(value string) {
	r.SetValue("denominations", value)
}

/* 优惠券的截止日期 */
func (r *PromotionCouponsGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 查询的页号，结果集是分页返回的，每页20条 */
func (r *PromotionCouponsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数 */
func (r *PromotionCouponsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *PromotionCouponsGetRequest) GetResponse(accessToken string) (*PromotionCouponsGetResponse, []byte, error) {
	var resp PromotionCouponsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.coupons.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionCouponsGetResponse struct {
	Coupons *CouponListObject `json:"coupons"`
	TotalResults int `json:"total_results"`
}

type PromotionCouponsGetResponseResult struct {
	Response *PromotionCouponsGetResponse `json:"promotion_coupons_get_response"`
}





/* 限时打折详情查询。查询出指定限时打折的对应商品记录信息。 */
type PromotionLimitdiscountDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 限时打折ID。这个针对查询唯一限时打折情况。 */
func (r *PromotionLimitdiscountDetailGetRequest) SetLimitDiscountId(value string) {
	r.SetValue("limit_discount_id", value)
}


func (r *PromotionLimitdiscountDetailGetRequest) GetResponse(accessToken string) (*PromotionLimitdiscountDetailGetResponse, []byte, error) {
	var resp PromotionLimitdiscountDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.limitdiscount.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionLimitdiscountDetailGetResponse struct {
	ItemDiscountDetailList *LimitDiscountDetailListObject `json:"item_discount_detail_list"`
}

type PromotionLimitdiscountDetailGetResponseResult struct {
	Response *PromotionLimitdiscountDetailGetResponse `json:"promotion_limitdiscount_detail_get_response"`
}





/* 分页查询某个卖家的限时打折信息。每页20条数据，按照结束时间降序排列。也可指定某一个限时打折id查询唯一的限时打折信息。 */
type PromotionLimitdiscountGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 限时打折结束时间。输入的时间会被截取，年月日有效，时分秒忽略。 */
func (r *PromotionLimitdiscountGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 限时打折ID。这个针对查询唯一限时打折情况。若此字段不为空，则说明操作为单条限时打折记录查询，其他字段忽略。若想分页按条件查询，这个字段置为空。 */
func (r *PromotionLimitdiscountGetRequest) SetLimitDiscountId(value string) {
	r.SetValue("limit_discount_id", value)
}

/* 分页页号。默认1。当页数大于最大页数时，结果为最大页数的数据。 */
func (r *PromotionLimitdiscountGetRequest) SetPageNumber(value string) {
	r.SetValue("page_number", value)
}

/* 限时打折开始时间。输入的时间会被截取，年月日有效，时分秒忽略。 */
func (r *PromotionLimitdiscountGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 限时打折活动状态。ALL:全部状态;OVER:已结束;DOING:进行中;PROPARE:未开始(只支持大写)。当limit_discount_id为空时，为空时，默认为全部的状态。 */
func (r *PromotionLimitdiscountGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *PromotionLimitdiscountGetRequest) GetResponse(accessToken string) (*PromotionLimitdiscountGetResponse, []byte, error) {
	var resp PromotionLimitdiscountGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.limitdiscount.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionLimitdiscountGetResponse struct {
	LimitDiscountList *LimitDiscountListObject `json:"limit_discount_list"`
	TotalCount int `json:"total_count"`
}

type PromotionLimitdiscountGetResponseResult struct {
	Response *PromotionLimitdiscountGetResponse `json:"promotion_limitdiscount_get_response"`
}





/* 搭配套餐查询。每个卖家最多创建50个搭配套餐，所以查询不会分页，会将所有的满足状态的搭配套餐全部查出。该接口不会校验商品的下架或库存为0，查询结果的状态表明搭配套餐在数据库中的状态，商品的状态请isv自己验证。在卖家后台页面点击查看会触发数据库状态的修改。 */
type PromotionMealGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 搭配套餐id */
func (r *PromotionMealGetRequest) SetMealId(value string) {
	r.SetValue("meal_id", value)
}

/* 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。 */
func (r *PromotionMealGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *PromotionMealGetRequest) GetResponse(accessToken string) (*PromotionMealGetResponse, []byte, error) {
	var resp PromotionMealGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.promotion.meal.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PromotionMealGetResponse struct {
	MealList *MealListObject `json:"meal_list"`
}

type PromotionMealGetResponseResult struct {
	Response *PromotionMealGetResponse `json:"promotion_meal_get_response"`
}





/* 查询活动列表 */
type UmpActivitiesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分页的页码 */
func (r *UmpActivitiesGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页的最大条数 */
func (r *UmpActivitiesGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 工具id */
func (r *UmpActivitiesGetRequest) SetToolId(value string) {
	r.SetValue("tool_id", value)
}


func (r *UmpActivitiesGetRequest) GetResponse(accessToken string) (*UmpActivitiesGetResponse, []byte, error) {
	var resp UmpActivitiesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activities.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivitiesGetResponse struct {
	Contents []string `json:"contents"`
	TotalCount int `json:"total_count"`
}

type UmpActivitiesGetResponseResult struct {
	Response *UmpActivitiesGetResponse `json:"ump_activities_get_response"`
}





/* 按照营销活动id的列表ids，查询对应的营销活动列表。 */
type UmpActivitiesListGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销活动id列表。 */
func (r *UmpActivitiesListGetRequest) SetIds(value string) {
	r.SetValue("ids", value)
}


func (r *UmpActivitiesListGetRequest) GetResponse(accessToken string) (*UmpActivitiesListGetResponse, []byte, error) {
	var resp UmpActivitiesListGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activities.list.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivitiesListGetResponse struct {
	Activities []string `json:"activities"`
}

type UmpActivitiesListGetResponseResult struct {
	Response *UmpActivitiesListGetResponse `json:"ump_activities_list_get_response"`
}





/* 新增优惠活动。设置优惠活动的基本信息，比如活动时间，活动针对的对象（可以是满足某些条件的买家） */
type UmpActivityAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动内容，通过ump sdk里面的MarkeitngTool来生成 */
func (r *UmpActivityAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 工具id */
func (r *UmpActivityAddRequest) SetToolId(value string) {
	r.SetValue("tool_id", value)
}


func (r *UmpActivityAddRequest) GetResponse(accessToken string) (*UmpActivityAddResponse, []byte, error) {
	var resp UmpActivityAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activity.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivityAddResponse struct {
	ActId int `json:"act_id"`
}

type UmpActivityAddResponseResult struct {
	Response *UmpActivityAddResponse `json:"ump_activity_add_response"`
}





/* 删除营销活动。对应的活动详情等将会被全部删除。 */
type UmpActivityDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpActivityDeleteRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}


func (r *UmpActivityDeleteRequest) GetResponse(accessToken string) (*UmpActivityDeleteResponse, []byte, error) {
	var resp UmpActivityDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activity.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivityDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpActivityDeleteResponseResult struct {
	Response *UmpActivityDeleteResponse `json:"ump_activity_delete_response"`
}





/* 查询营销活动 */
type UmpActivityGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpActivityGetRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}


func (r *UmpActivityGetRequest) GetResponse(accessToken string) (*UmpActivityGetResponse, []byte, error) {
	var resp UmpActivityGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activity.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivityGetResponse struct {
	Content string `json:"content"`
}

type UmpActivityGetResponseResult struct {
	Response *UmpActivityGetResponse `json:"ump_activity_get_response"`
}





/* 修改营销活动 */
type UmpActivityUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpActivityUpdateRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* 营销活动内容，json格式，通过ump sdk 的marketingTool来生成 */
func (r *UmpActivityUpdateRequest) SetContent(value string) {
	r.SetValue("content", value)
}


func (r *UmpActivityUpdateRequest) GetResponse(accessToken string) (*UmpActivityUpdateResponse, []byte, error) {
	var resp UmpActivityUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.activity.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpActivityUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpActivityUpdateResponseResult struct {
	Response *UmpActivityUpdateResponse `json:"ump_activity_update_response"`
}





/* 增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。 */
type UmpDetailAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 增加工具详情 */
func (r *UmpDetailAddRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理 */
func (r *UmpDetailAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}


func (r *UmpDetailAddRequest) GetResponse(accessToken string) (*UmpDetailAddResponse, []byte, error) {
	var resp UmpDetailAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.detail.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailAddResponse struct {
	DetailId int `json:"detail_id"`
}

type UmpDetailAddResponseResult struct {
	Response *UmpDetailAddResponse `json:"ump_detail_add_response"`
}





/* 删除活动详情 */
type UmpDetailDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动详情id */
func (r *UmpDetailDeleteRequest) SetDetailId(value string) {
	r.SetValue("detail_id", value)
}


func (r *UmpDetailDeleteRequest) GetResponse(accessToken string) (*UmpDetailDeleteResponse, []byte, error) {
	var resp UmpDetailDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.detail.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpDetailDeleteResponseResult struct {
	Response *UmpDetailDeleteResponse `json:"ump_detail_delete_response"`
}





/* 查询活动详情 */
type UmpDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动详情的id */
func (r *UmpDetailGetRequest) SetDetailId(value string) {
	r.SetValue("detail_id", value)
}


func (r *UmpDetailGetRequest) GetResponse(accessToken string) (*UmpDetailGetResponse, []byte, error) {
	var resp UmpDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailGetResponse struct {
	Content string `json:"content"`
}

type UmpDetailGetResponseResult struct {
	Response *UmpDetailGetResponse `json:"ump_detail_get_response"`
}





/* 批量添加营销活动。替代单条添加营销详情的的API。此接口适用针对某个营销活动的多档设置，会按顺序插入detail。若在整个事务过程中出现断点，会将已插入完成的detail_id返回，注意记录这些id，并将其删除，会对交易过程中的优惠产生影响。 */
type UmpDetailListAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销活动id。 */
func (r *UmpDetailListAddRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* 营销详情的列表。此列表由detail的json字符串组成。最多插入为10个。 */
func (r *UmpDetailListAddRequest) SetDetails(value string) {
	r.SetValue("details", value)
}


func (r *UmpDetailListAddRequest) GetResponse(accessToken string) (*UmpDetailListAddResponse, []byte, error) {
	var resp UmpDetailListAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.detail.list.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailListAddResponse struct {
	DetailIdList []int `json:"detail_id_list"`
}

type UmpDetailListAddResponseResult struct {
	Response *UmpDetailListAddResponse `json:"ump_detail_list_add_response"`
}





/* 更新活动详情 */
type UmpDetailUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容 */
func (r *UmpDetailUpdateRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 活动详情id */
func (r *UmpDetailUpdateRequest) SetDetailId(value string) {
	r.SetValue("detail_id", value)
}


func (r *UmpDetailUpdateRequest) GetResponse(accessToken string) (*UmpDetailUpdateResponse, []byte, error) {
	var resp UmpDetailUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.detail.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailUpdateResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpDetailUpdateResponseResult struct {
	Response *UmpDetailUpdateResponse `json:"ump_detail_update_response"`
}





/* 分页查询优惠详情列表 */
type UmpDetailsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销活动id */
func (r *UmpDetailsGetRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* 分页的页码 */
func (r *UmpDetailsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页的最大条数 */
func (r *UmpDetailsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *UmpDetailsGetRequest) GetResponse(accessToken string) (*UmpDetailsGetResponse, []byte, error) {
	var resp UmpDetailsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.details.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpDetailsGetResponse struct {
	Contents []string `json:"contents"`
	TotalCount int `json:"total_count"`
}

type UmpDetailsGetResponseResult struct {
	Response *UmpDetailsGetResponse `json:"ump_details_get_response"`
}





/* 根据营销积木块代码获取积木块。接口返回该代码最新版本的积木块。如果要查询某个非最新版本的积木块，可以使用积木块id查询接口。 */
type UmpMbbGetbycodeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销积木块code */
func (r *UmpMbbGetbycodeRequest) SetCode(value string) {
	r.SetValue("code", value)
}


func (r *UmpMbbGetbycodeRequest) GetResponse(accessToken string) (*UmpMbbGetbycodeResponse, []byte, error) {
	var resp UmpMbbGetbycodeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.mbb.getbycode", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpMbbGetbycodeResponse struct {
	Mbb string `json:"mbb"`
}

type UmpMbbGetbycodeResponseResult struct {
	Response *UmpMbbGetbycodeResponse `json:"ump_mbb_getbycode_response"`
}





/* 根据积木块id获取营销积木块。 */
type UmpMbbGetbyidRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 积木块的id */
func (r *UmpMbbGetbyidRequest) SetId(value string) {
	r.SetValue("id", value)
}


func (r *UmpMbbGetbyidRequest) GetResponse(accessToken string) (*UmpMbbGetbyidResponse, []byte, error) {
	var resp UmpMbbGetbyidResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.mbb.getbyid", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpMbbGetbyidResponse struct {
	Mbb string `json:"mbb"`
}

type UmpMbbGetbyidResponseResult struct {
	Response *UmpMbbGetbyidResponse `json:"ump_mbb_getbyid_response"`
}





/* 获取营销积木块列表，可以根据类型获取，也可以将该字段设为空，获取所有的 */
type UmpMbbsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 积木块类型。如果该字段为空表示查出所有类型的 
现在有且仅有如下几种：resource,condition,action,target */
func (r *UmpMbbsGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *UmpMbbsGetRequest) GetResponse(accessToken string) (*UmpMbbsGetResponse, []byte, error) {
	var resp UmpMbbsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.mbbs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpMbbsGetResponse struct {
	Mbbs []string `json:"mbbs"`
}

type UmpMbbsGetResponseResult struct {
	Response *UmpMbbsGetResponse `json:"ump_mbbs_get_response"`
}





/* 通过营销积木id列表来获取营销积木块列表。 */
type UmpMbbsListGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销积木块id组成的字符串。 */
func (r *UmpMbbsListGetRequest) SetIds(value string) {
	r.SetValue("ids", value)
}


func (r *UmpMbbsListGetRequest) GetResponse(accessToken string) (*UmpMbbsListGetResponse, []byte, error) {
	var resp UmpMbbsListGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.mbbs.list.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpMbbsListGetResponse struct {
	Mbbs []string `json:"mbbs"`
}

type UmpMbbsListGetResponseResult struct {
	Response *UmpMbbsListGetResponse `json:"ump_mbbs_list_get_response"`
}





/* 指定某项活动中，某个商家的某些类型物品（指定商品，指定类目或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。 */
type UmpRangeAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpRangeAddRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id.多个id用逗号隔开，一次不超过50个 */
func (r *UmpRangeAddRequest) SetIds(value string) {
	r.SetValue("ids", value)
}

/* 范围的类型，比如是全店，商品，类目 
见：MarketingConstants.PARTICIPATE_TYPE_* */
func (r *UmpRangeAddRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *UmpRangeAddRequest) GetResponse(accessToken string) (*UmpRangeAddResponse, []byte, error) {
	var resp UmpRangeAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.range.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpRangeAddResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpRangeAddResponseResult struct {
	Response *UmpRangeAddResponse `json:"ump_range_add_response"`
}





/* 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置 */
type UmpRangeDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpRangeDeleteRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}

/* id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id */
func (r *UmpRangeDeleteRequest) SetIds(value string) {
	r.SetValue("ids", value)
}

/* 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_* */
func (r *UmpRangeDeleteRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *UmpRangeDeleteRequest) GetResponse(accessToken string) (*UmpRangeDeleteResponse, []byte, error) {
	var resp UmpRangeDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.range.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpRangeDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpRangeDeleteResponseResult struct {
	Response *UmpRangeDeleteResponse `json:"ump_range_delete_response"`
}





/* 查询某个卖家所有参加或者不参加某项活动的物品 */
type UmpRangeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 活动id */
func (r *UmpRangeGetRequest) SetActId(value string) {
	r.SetValue("act_id", value)
}


func (r *UmpRangeGetRequest) GetResponse(accessToken string) (*UmpRangeGetResponse, []byte, error) {
	var resp UmpRangeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.range.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpRangeGetResponse struct {
	Ranges *RangeListObject `json:"ranges"`
}

type UmpRangeGetResponseResult struct {
	Response *UmpRangeGetResponse `json:"ump_range_get_response"`
}





/* 新增优惠工具。通过ump sdk来完成将积木块拼装成为工具的操作，再使用本接口上传优惠工具。 */
type UmpToolAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 工具内容，由sdk里面的MarketingTool生成的json字符串 */
func (r *UmpToolAddRequest) SetContent(value string) {
	r.SetValue("content", value)
}


func (r *UmpToolAddRequest) GetResponse(accessToken string) (*UmpToolAddResponse, []byte, error) {
	var resp UmpToolAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.tool.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpToolAddResponse struct {
	ToolId int `json:"tool_id"`
}

type UmpToolAddResponseResult struct {
	Response *UmpToolAddResponse `json:"ump_tool_add_response"`
}





/* 删除营销工具。当工具正在被使用的时候，是不能被删除的。 */
type UmpToolDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 营销工具id */
func (r *UmpToolDeleteRequest) SetToolId(value string) {
	r.SetValue("tool_id", value)
}


func (r *UmpToolDeleteRequest) GetResponse(accessToken string) (*UmpToolDeleteResponse, []byte, error) {
	var resp UmpToolDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.tool.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpToolDeleteResponse struct {
	IsSuccess bool `json:"is_success"`
}

type UmpToolDeleteResponseResult struct {
	Response *UmpToolDeleteResponse `json:"ump_tool_delete_response"`
}





/* 根据工具id获取一个工具对象 */
type UmpToolGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 工具的id */
func (r *UmpToolGetRequest) SetToolId(value string) {
	r.SetValue("tool_id", value)
}


func (r *UmpToolGetRequest) GetResponse(accessToken string) (*UmpToolGetResponse, []byte, error) {
	var resp UmpToolGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.tool.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpToolGetResponse struct {
	Content string `json:"content"`
}

type UmpToolGetResponseResult struct {
	Response *UmpToolGetResponse `json:"ump_tool_get_response"`
}





/* 修改工具 */
type UmpToolUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 工具的内容，由sdk的marketingBuilder生成 */
func (r *UmpToolUpdateRequest) SetContent(value string) {
	r.SetValue("content", value)
}

/* 工具id */
func (r *UmpToolUpdateRequest) SetToolId(value string) {
	r.SetValue("tool_id", value)
}


func (r *UmpToolUpdateRequest) GetResponse(accessToken string) (*UmpToolUpdateResponse, []byte, error) {
	var resp UmpToolUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.tool.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpToolUpdateResponse struct {
	ToolId int `json:"tool_id"`
}

type UmpToolUpdateResponseResult struct {
	Response *UmpToolUpdateResponse `json:"ump_tool_update_response"`
}





/* 查询工具列表 */
type UmpToolsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 工具编码 */
func (r *UmpToolsGetRequest) SetToolCode(value string) {
	r.SetValue("tool_code", value)
}


func (r *UmpToolsGetRequest) GetResponse(accessToken string) (*UmpToolsGetResponse, []byte, error) {
	var resp UmpToolsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.ump.tools.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type UmpToolsGetResponse struct {
	Tools []string `json:"tools"`
}

type UmpToolsGetResponseResult struct {
	Response *UmpToolsGetResponse `json:"ump_tools_get_response"`
}



