// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package trvael

import (
	"github.com/changkong/open_taobao"
)





/* 增加一个旅游商品，目前支持的类目包括：国内跟团游、国内自由行、国内一日游、出境自由行、出境跟团游、境外跟团游、境外自由行、境外一日游。 */
type TravelItemsAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale。 */
func (r *TravelItemsAddRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数.B商家在发布非虚拟商品时，这个字段必须设置，返点必须是 5的倍数，即0.5%的倍数。注意该字段值最高值不超过500，即50%。 */
func (r *TravelItemsAddRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄。 */
func (r *TravelItemsAddRequest) SetChooseLogis(value string) {
	r.SetValue("choose_logis", value)
}

/* 商品所属叶子类目ID。 */
func (r *TravelItemsAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 宝贝所在地市。 */
func (r *TravelItemsAddRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* Json串：{"combos":[{"combo_name":"套餐一","price_calendar":[{"child_num":100,"child_price":100,"date":"2012-06-08","diff_price":1000,"man_num":100,"man_price":1000},{"child_num":100,"child_price":100,"date":"2012-06-09","diff_price":1000,"man_num":100,"man_price":1000}]}]} */
func (r *TravelItemsAddRequest) SetComboPriceCalendar(value string) {
	r.SetValue("combo_price_calendar", value)
}

/* 商品描述，不超过50000个字符。 */
func (r *TravelItemsAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 最晚成团提前天数，最小0天，最大为300天。 */
func (r *TravelItemsAddRequest) SetDuration(value string) {
	r.SetValue("duration", value)
}

/* 电子交易凭证有效期，目前此字段只涉及到的信息为有效期; 如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 如果有效期为【购买成功日 至】类型则格式为2012-08-16 如果有效期为天数类型则格式为15。 */
func (r *TravelItemsAddRequest) SetExpirydate(value string) {
	r.SetValue("expirydate", value)
}

/* 费用包含，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetFeeInclude(value string) {
	r.SetValue("fee_include", value)
}

/* 费用不包，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetFeeNotInclude(value string) {
	r.SetValue("fee_not_include", value)
}

/* 机票信息，不超过1500字符 */
func (r *TravelItemsAddRequest) SetFlightInfo(value string) {
	r.SetValue("flight_info", value)
}

/* 集合地，不超过30个字符。 */
func (r *TravelItemsAddRequest) SetGatheringPlace(value string) {
	r.SetValue("gathering_place", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折)。 */
func (r *TravelItemsAddRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)。 */
func (r *TravelItemsAddRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐)，B商家不用设置该字段，均为true。 */
func (r *TravelItemsAddRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 酒店信息，不超过1500字符 */
func (r *TravelItemsAddRequest) SetHotelInfo(value string) {
	r.SetValue("hotel_info", value)
}

/* 商品主图。类型:JPG,GIF;最大长度:500k，支持的文件类型：gif,jpg,jpeg,png。 */
func (r *TravelItemsAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："2000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。在度假线路类目中，该属性ID为“自由行包含元素”或“一日游包含元素”属性ID。 */
func (r *TravelItemsAddRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，如“自定义自由行包含元素”。 */
func (r *TravelItemsAddRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否为铁定出游商品的参数设置为true，那么该商品为铁定出游商品；设置为false，那么该商品为非铁定出游商品。默认为false */
func (r *TravelItemsAddRequest) SetIsTdcy(value string) {
	r.SetValue("is_tdcy", value)
}

/* 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
func (r *TravelItemsAddRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 码商信息，格式为 码商id:nick。 */
func (r *TravelItemsAddRequest) SetMerchant(value string) {
	r.SetValue("merchant", value)
}

/* 网点ID。 */
func (r *TravelItemsAddRequest) SetNetworkId(value string) {
	r.SetValue("network_id", value)
}

/* 商品库存。如果发布旅游度假线路宝贝，该字段可以忽略。 */
func (r *TravelItemsAddRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数。(暂未使用) */
func (r *TravelItemsAddRequest) SetOnsaleAutoRefundRatio(value string) {
	r.SetValue("onsale_auto_refund_ratio", value)
}

/* 预定须知，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetOrderInfo(value string) {
	r.SetValue("order_info", value)
}

/* 商家的外部编码，最大为512字节。 */
func (r *TravelItemsAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 自费项目，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetOwnExpense(value string) {
	r.SetValue("own_expense", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path。 */
func (r *TravelItemsAddRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 商品一口价。如果发布旅游度假线路的宝贝，该字段可以忽略。 */
func (r *TravelItemsAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品属性列表。格式为：pid:vid;pid:vid。例如发布度假线路商品，那么这里就需要填写：出发地属性id:城市id;目的地市属性id:目的地市id;……等等。 */
func (r *TravelItemsAddRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 宝贝所在地省份。 */
func (r *TravelItemsAddRequest) SetProv(value string) {
	r.SetValue("prov", value)
}

/* 退款比例，百分比%前的数字,1-100的正整数值。 */
func (r *TravelItemsAddRequest) SetRefundRatio(value string) {
	r.SetValue("refund_ratio", value)
}

/* 退改规定，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetRefundRegulation(value string) {
	r.SetValue("refund_regulation", value)
}

/* 商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)。 */
func (r *TravelItemsAddRequest) SetSecondKill(value string) {
	r.SetValue("second_kill", value)
}

/* 关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *TravelItemsAddRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* 购物店信息，不超过1500个字符。 */
func (r *TravelItemsAddRequest) SetShopingInfo(value string) {
	r.SetValue("shoping_info", value)
}

/* sku销售属性串对应的价格，精确到分，每一个属性串都会对应一个价格，单位为分。sku_prices的数组大小应该和sku_properties的数组大小一致。如果发布线路的商品，该字段可以忽略。 */
func (r *TravelItemsAddRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* Sku销售属性串，调用taobao.travel.itemsprops.get接口获取商品销售属性code，然后拼接成pid:vid;pid:vid格式。如果发布线路的商品，该字段可以忽略。 */
func (r *TravelItemsAddRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* Sku销售属性串对应的库存，每一个属性串都会对应一个库存，显然sku_quantities的数组大小应该和sku_properties的数组大小一致。如果发布线路的商品，该字段可以忽略。 */
func (r *TravelItemsAddRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改，集市卖家默认拍下减库存;商城卖家默认付款减库存。 */
func (r *TravelItemsAddRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 门票信息，不超过1500字符 */
func (r *TravelItemsAddRequest) SetTicketInfo(value string) {
	r.SetValue("ticket_info", value)
}

/* 商品标题。 */
func (r *TravelItemsAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 核销打款 1代表核销打款 0代表非核销打款。 */
func (r *TravelItemsAddRequest) SetVerification(value string) {
	r.SetValue("verification", value)
}


func (r *TravelItemsAddRequest) GetResponse(accessToken string) (*TravelItemsAddResponse, []byte, error) {
	var resp TravelItemsAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.travel.items.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TravelItemsAddResponse struct {
	TravelItems *TravelItems `json:"travel_items"`
}

type TravelItemsAddResponseResult struct {
	Response *TravelItemsAddResponse `json:"travel_items_add_response"`
}





/* 此接口用于查询一个旅游度假线路商品信息，根据传入的商品数字ID查询商品信息，目前仅支持8个类目：国内跟团游、国内自由行、国内一日游、出境自由行、出境跟团游、境外跟团游、境外自由行、境外一日游商品的查询。 */
type TravelItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID。 */
func (r *TravelItemsGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *TravelItemsGetRequest) GetResponse(accessToken string) (*TravelItemsGetResponse, []byte, error) {
	var resp TravelItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.travel.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TravelItemsGetResponse struct {
	TravelItems *TravelItems `json:"travel_items"`
}

type TravelItemsGetResponseResult struct {
	Response *TravelItemsGetResponse `json:"travel_items_get_response"`
}





/* 更新一个旅游度假线路商品，目前仅支持8个类目：国内跟团游、国内自由行、国内一日游、出境自由行、出境跟团游、境外跟团游、境外自由行、境外一日游商品的更新。 */
type TravelItemsUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 套餐价格日历增量更新字段，添加若干新套餐。（说明：如果使用增量更新字段，则全量更新字段combo_price_calendar不用设置，如果设置了则会优先使用全量更新），套餐价格日历json格式。如：{"combos":[{"combo_name":"套餐一","price_calendar":[{"child_num":100,"child_price":100,"date":"2012-06-08","diff_price":1000,"man_num":100,"man_price":1000},{"child_num":100,"child_price":100,"date":"2012-06-09","diff_price":1000,"man_num":100,"man_price":1000}]}]} */
func (r *TravelItemsUpdateRequest) SetAddComboPriceCalendar(value string) {
	r.SetValue("add_combo_price_calendar", value)
}

/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale。 */
func (r *TravelItemsUpdateRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数. */
func (r *TravelItemsUpdateRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄。 */
func (r *TravelItemsUpdateRequest) SetChooseLogis(value string) {
	r.SetValue("choose_logis", value)
}

/* 商品所属类目ID。发布旅游线路商品有两个值：1(国内线路类目ID)，2(国际线路类目ID)。 */
func (r *TravelItemsUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 宝贝所在地市。如果发布旅游度假线路的宝贝，该字段可以忽略。 */
func (r *TravelItemsUpdateRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* Json串，全量更新套餐价格日历信息（针对旅游度假线路的销售属性），定义了两个套餐日历价格，套餐明分别为：套餐一和套餐二。如：{"combos":[{"combo_name":"套餐一","price_calendar":[{"child_num":100,"child_price":100,"date":"2012-06-08","diff_price":1000,"man_num":100,"man_price":1000},{"child_num":100,"child_price":100,"date":"2012-06-09","diff_price":1000,"man_num":100,"man_price":1000}]}]} */
func (r *TravelItemsUpdateRequest) SetComboPriceCalendar(value string) {
	r.SetValue("combo_price_calendar", value)
}

/* 商品描述，不超过50000个字符。 */
func (r *TravelItemsUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 最晚成团提前天数，最小0天，最大为300天。 */
func (r *TravelItemsUpdateRequest) SetDuration(value string) {
	r.SetValue("duration", value)
}

/* 支持宝贝信息的删除，各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：locality_life表示删除电子凭证，merchant表示删除码商信息，refund_ratio表示删除电子凭证退款比例，network_id表示删除电子凭证网点信息,seller_cids表示删除关联店铺类目，outer_id表示删除商家外部编码，second_kill表示删除秒杀信息，input_pids表示删除用户自定义属性 */
func (r *TravelItemsUpdateRequest) SetEmptyFields(value string) {
	r.SetValue("empty_fields", value)
}

/* 电子交易凭证有效期，目前此字段只涉及到的信息为有效期; 如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 如果有效期为【购买成功日 至】类型则格式为2012-08-16 如果有效期为天数类型则格式为15。 */
func (r *TravelItemsUpdateRequest) SetExpirydate(value string) {
	r.SetValue("expirydate", value)
}

/* 费用包含，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetFeeInclude(value string) {
	r.SetValue("fee_include", value)
}

/* 费用不包，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetFeeNotInclude(value string) {
	r.SetValue("fee_not_include", value)
}

/* 机票信息，不超过1500字符 */
func (r *TravelItemsUpdateRequest) SetFlightInfo(value string) {
	r.SetValue("flight_info", value)
}

/* 集合地，不超过30个字符。 */
func (r *TravelItemsUpdateRequest) SetGatheringPlace(value string) {
	r.SetValue("gathering_place", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折)。 */
func (r *TravelItemsUpdateRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票)。 */
func (r *TravelItemsUpdateRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐)，B商家不用设置该字段，均为true。 */
func (r *TravelItemsUpdateRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 酒店信息，不超过1500字符 */
func (r *TravelItemsUpdateRequest) SetHotelInfo(value string) {
	r.SetValue("hotel_info", value)
}

/* 商品主图。类型:JPG,GIF;最大长度:500k，支持的文件类型：gif,jpg,jpeg,png。 */
func (r *TravelItemsUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："2000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。在度假线路类目中，该属性ID为“自由行包含元素”或“一日游包含元素”属性ID。 */
func (r *TravelItemsUpdateRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，如“自定义自由行包含元素”。 */
func (r *TravelItemsUpdateRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是铁定出游商品 */
func (r *TravelItemsUpdateRequest) SetIsTdcy(value string) {
	r.SetValue("is_tdcy", value)
}

/* 商品数字ID。 */
func (r *TravelItemsUpdateRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 码商信息，格式为 码商id:nick。 */
func (r *TravelItemsUpdateRequest) SetMerchant(value string) {
	r.SetValue("merchant", value)
}

/* 网点ID。 */
func (r *TravelItemsUpdateRequest) SetNetworkId(value string) {
	r.SetValue("network_id", value)
}

/* 商品库存。如果发布旅游度假线路宝贝，该字段可以忽略，参考后面：add_combo_price_calendar,update_combo_price_calendar,remove_combo_price_calendar 这些字段去使用商品销售属性 */
func (r *TravelItemsUpdateRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数。 */
func (r *TravelItemsUpdateRequest) SetOnsaleAutoRefundRatio(value string) {
	r.SetValue("onsale_auto_refund_ratio", value)
}

/* 预定须知，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetOrderInfo(value string) {
	r.SetValue("order_info", value)
}

/* 商家的外部编码，最大为512字节。 */
func (r *TravelItemsUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 自费项目，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetOwnExpense(value string) {
	r.SetValue("own_expense", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path。 */
func (r *TravelItemsUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 商品一口价。如果发布旅游度假线路宝贝，该字段可以忽略，参考后面：add_combo_price_calendar,update_combo_price_calendar,remove_combo_price_calendar 这些字段去使用商品销售属性 */
func (r *TravelItemsUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品属性列表。格式为：pid:vid;pid:vid。例如发布旅游度假线路商品，那么这里就需要填写：出发地属性id:城市id;目的地市属性id:目的地市id;……等等。 */
func (r *TravelItemsUpdateRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 宝贝所在地省份。如果发布旅游线路宝贝，该字段可以忽略。 */
func (r *TravelItemsUpdateRequest) SetProv(value string) {
	r.SetValue("prov", value)
}

/* 退款比例，百分比%前的数字,1-100的正整数值。 */
func (r *TravelItemsUpdateRequest) SetRefundRatio(value string) {
	r.SetValue("refund_ratio", value)
}

/* 退改规定，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetRefundRegulation(value string) {
	r.SetValue("refund_regulation", value)
}

/* 套餐价格日历增量更新字段，删除若干套餐。（说明：如果使用增量更新字段，则全量更新字段combo_price_calendar不用设置，如果设置了则会优先使用全量更新）。删除时，需要设置套餐属性id（pid），套餐属性值id（vid）。格式为：pid:vid1;pid:vid2;pid:vid3。 */
func (r *TravelItemsUpdateRequest) SetRemoveComboPriceCalendar(value string) {
	r.SetValue("remove_combo_price_calendar", value)
}

/* 商品属性（不包含销售属性）增量更新字段，删除商品属性。（说明：如果使用增量更新字段，则全量更新字段props不用设置，如果设置了则会优先使用全量更新）。格式：pid1:vid1;pid2:vid2;pid3:vid3。 */
func (r *TravelItemsUpdateRequest) SetRemoveProps(value string) {
	r.SetValue("remove_props", value)
}

/* 商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)。 */
func (r *TravelItemsUpdateRequest) SetSecondKill(value string) {
	r.SetValue("second_kill", value)
}

/* 关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *TravelItemsUpdateRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* 购物店信息，不超过1500个字符。 */
func (r *TravelItemsUpdateRequest) SetShopingInfo(value string) {
	r.SetValue("shoping_info", value)
}

/* Sku销售属性串对应的价格，每一个属性串都会对应一个价格，精确到两位小数，单位为元。sku_prices的数组大小应该和sku_properties的数组大小一致。如果发布线路的商品，参考后面：add_combo_price_calendar,update_combo_price_calendar,remove_combo_price_calendar 这些字段去使用商品销售属性 */
func (r *TravelItemsUpdateRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* Sku销售属性串，调用taobao.travel.itemsprops.get接口获取商品销售属性code，然后拼接成pid:vid;pid:vid格式。如果发布线路的商品，参考后面：add_combo_price_calendar,update_combo_price_calendar,remove_combo_price_calendar 这些字段去使用商品销售属性 */
func (r *TravelItemsUpdateRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* Sku销售属性串对应的库存，每一个属性串都会对应一个库存，显然sku_quantities的数组大小应该和sku_properties的数组大小一致。如果发布线路的商品，参考后面：add_combo_price_calendar,update_combo_price_calendar,remove_combo_price_calendar 这些字段去使用商品销售属性 */
func (r *TravelItemsUpdateRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改，集市卖家默认拍下减库存;商城卖家默认付款减库存。 */
func (r *TravelItemsUpdateRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 门票信息，不超过1500字符 */
func (r *TravelItemsUpdateRequest) SetTicketInfo(value string) {
	r.SetValue("ticket_info", value)
}

/* 商品标题。注意：在商品更新时，如果不设置该属性，默认不进行更新，下同。 */
func (r *TravelItemsUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 否	套餐价格日历增量更新字段，更新若干套餐。（说明：如果使用增量更新字段，则全量更新字段combo_price_calendar不用设置，如果设置了则会优先使用全量更新），套餐价格日历json格式。更新时，需要设置套餐属性id（pid），套餐属性值id（vid），套餐名称可以不设置，设置了也会忽略，会以传入的pid和vid为准。如：{"combos":[{"combo_name":"套餐一","pid":102020,"vid":289129,"price_calendar":[{"child_num":100,"child_price":100,"date":"2012-06-08","diff_price":1000,"man_num":100,"man_price":1000},{"child_num":100,"child_price":100,"date":"2012-06-09","diff_price":1000,"man_num":100,"man_price":1000}]}]} */
func (r *TravelItemsUpdateRequest) SetUpdateComboPriceCalendar(value string) {
	r.SetValue("update_combo_price_calendar", value)
}

/* 商品属性（不包含销售属性）增量更新字段，更新或者添加商品属性。（说明：如果使用增量更新字段，则全量更新字段props不用设置，如果设置了则会优先使用全量更新）。格式：pid1:vid1;pid2:vid2;pid3:vid3。对于重复设置的同一个属性的多个值，对于单选属性，则会以最后一个为准；对于多选，则会对该属性新增属性值。 */
func (r *TravelItemsUpdateRequest) SetUpdateOrAddProps(value string) {
	r.SetValue("update_or_add_props", value)
}

/* 核销打款 1代表核销打款 0代表非核销打款。 */
func (r *TravelItemsUpdateRequest) SetVerification(value string) {
	r.SetValue("verification", value)
}


func (r *TravelItemsUpdateRequest) GetResponse(accessToken string) (*TravelItemsUpdateResponse, []byte, error) {
	var resp TravelItemsUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.travel.items.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TravelItemsUpdateResponse struct {
	TravelItems *TravelItems `json:"travel_items"`
}

type TravelItemsUpdateResponseResult struct {
	Response *TravelItemsUpdateResponse `json:"travel_items_update_response"`
}





/* 此接口用于获取卖家发布旅游度假线路商品时目的地地区级联信息。 */
type TravelItemsareaGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品所属叶子类目ID，支持旅游度假线路8个类目。 */
func (r *TravelItemsareaGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}


func (r *TravelItemsareaGetRequest) GetResponse(accessToken string) (*TravelItemsareaGetResponse, []byte, error) {
	var resp TravelItemsareaGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.travel.itemsarea.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TravelItemsareaGetResponse struct {
	TravelItemsAreaNodes []*TravelItemsAreaNode `json:"travel_items_area_nodes"`
}

type TravelItemsareaGetResponseResult struct {
	Response *TravelItemsareaGetResponse `json:"travel_itemsarea_get_response"`
}





/* 此接口用于获取旅游商品类目属性信息。 */
type TravelItemspropsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品所属叶子类目ID，支持旅游度假线路8个类目。 */
func (r *TravelItemspropsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 属性id (取某个类目属性时，传pid；若不传该值，返回该类目下所有属性) */
func (r *TravelItemspropsGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}


func (r *TravelItemspropsGetRequest) GetResponse(accessToken string) (*TravelItemspropsGetResponse, []byte, error) {
	var resp TravelItemspropsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.travel.itemsprops.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TravelItemspropsGetResponse struct {
	TravelItemsProps []*TravelItemsProp `json:"travel_items_props"`
}

type TravelItemspropsGetResponseResult struct {
	Response *TravelItemspropsGetResponse `json:"travel_itemsprops_get_response"`
}



