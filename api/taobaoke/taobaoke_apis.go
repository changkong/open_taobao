// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package taobaoke

import (
	"github.com/changkong/open_taobao"
)

/* 淘宝客类目推广URL.
<p>由于受到<a href="http://club.alimama.com/read-htm-tid-3133847.html">淘宝联盟pid升级规则</a>影响，该接口后续可能会下线。建议直接使用拼链接的方式（塞入完整的pid）来获取搜索结果，而不是调用该接口来获取结果</p> */
type TaobaokeCaturlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 标准商品后台类目id。该ID可以通过taobao.itemcats.get接口获取到。 */
func (r *TaobaokeCaturlGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 推广者的淘宝会员昵称.注：这里指的是淘宝的登录会员名 */
func (r *TaobaokeCaturlGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeCaturlGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20 */
func (r *TaobaokeCaturlGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 关键词 */
func (r *TaobaokeCaturlGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

func (r *TaobaokeCaturlGetRequest) GetResponse(accessToken string) (*TaobaokeCaturlGetResponse, []byte, error) {
	var resp TaobaokeCaturlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.caturl.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeCaturlGetResponse struct {
	TaobaokeItem *TaobaokeItem `json:"taobaoke_item"`
}

type TaobaokeCaturlGetResponseResult struct {
	Response *TaobaokeCaturlGetResponse `json:"taobaoke_caturl_get_response"`
}

/* 查询淘客折扣商品。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeItemsCouponGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 商品所在地 */
func (r *TaobaokeItemsCouponGetRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 标准商品后台类目id。该ID可以通过taobao.itemcats.get接口获取到。 */
func (r *TaobaokeItemsCouponGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 优惠商品类型.1:打折商品,默认值为1 */
func (r *TaobaokeItemsCouponGetRequest) SetCouponType(value string) {
	r.SetValue("coupon_type", value)
}

/* 设置30天累计推广量（与返回数据中的commission_num字段对应）上限.注：该字段要与start_commissionNum一起使用才生效 */
func (r *TaobaokeItemsCouponGetRequest) SetEndCommissionNum(value string) {
	r.SetValue("end_commission_num", value)
}

/* 最高佣金比率选项，如：2345表示23.45%。注：要起始佣金比率和最高佣金比率一起设置才有效。 */
func (r *TaobaokeItemsCouponGetRequest) SetEndCommissionRate(value string) {
	r.SetValue("end_commission_rate", value)
}

/* 最高累计推广佣金选项 */
func (r *TaobaokeItemsCouponGetRequest) SetEndCommissionVolume(value string) {
	r.SetValue("end_commission_volume", value)
}

/* 设置折扣比例范围上限,如：8000表示80.00%.注：要起始折扣比率和最高折扣比率一起设置才有效 */
func (r *TaobaokeItemsCouponGetRequest) SetEndCouponRate(value string) {
	r.SetValue("end_coupon_rate", value)
}

/* 可选值和start_credit一样.start_credit的值一定要小于或等于end_credit的值。注：end_credit与start_credit一起使用才生效 */
func (r *TaobaokeItemsCouponGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 设置商品总成交量（与返回字段volume对应）上限。 */
func (r *TaobaokeItemsCouponGetRequest) SetEndVolume(value string) {
	r.SetValue("end_volume", value)
}

/* 需返回的字段列表.可选值:num_iid,title,nick,pic_url,price,click_url,commission,commission_rate,commission_num,commission_volume,shop_click_url,seller_credit_score,item_location,volume,coupon_price,coupon_rate,coupon_start_time,coupon_end_time,shop_type;字段之间用","分隔 */
func (r *TaobaokeItemsCouponGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false */
func (r *TaobaokeItemsCouponGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 商品标题中包含的关键字. 注意:查询时keyword,cid至少选择其中一个参数 */
func (r *TaobaokeItemsCouponGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 推广者的淘宝会员昵称.注:指的是淘宝的会员登录名 */
func (r *TaobaokeItemsCouponGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道 */
func (r *TaobaokeItemsCouponGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 结果页数.1~99 */
func (r *TaobaokeItemsCouponGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数.最大每页100 */
func (r *TaobaokeItemsCouponGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeItemsCouponGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 点击串跳转类型，1：单品，2：单品中间页（无线暂无） */
func (r *TaobaokeItemsCouponGetRequest) SetReferType(value string) {
	r.SetValue("refer_type", value)
}

/* 店铺类型.默认all,商城:b,集市:c */
func (r *TaobaokeItemsCouponGetRequest) SetShopType(value string) {
	r.SetValue("shop_type", value)
}

/* default(默认排序),
price_desc(折扣价格从高到低),
price_asc(折扣价格从低到高),
credit_desc(信用等级从高到低),
credit_asc(信用等级从低到高),
commissionRate_desc(佣金比率从高到低),
commissionRate_asc(佣金比率从低到高),
volume_desc(成交量成高到低), volume_asc(成交量从低到高) */
func (r *TaobaokeItemsCouponGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 设置30天累计推广量（与返回数据中的commission_num字段对应）下限.注：该字段要与end_commissionNum一起使用才生效 */
func (r *TaobaokeItemsCouponGetRequest) SetStartCommissionNum(value string) {
	r.SetValue("start_commission_num", value)
}

/* 起始佣金比率选项，如：1234表示12.34% */
func (r *TaobaokeItemsCouponGetRequest) SetStartCommissionRate(value string) {
	r.SetValue("start_commission_rate", value)
}

/* 起始累计推广量佣金.注：返回的数据是30天内累计推广佣金，该字段要与最高累计推广佣金一起使用才生效 */
func (r *TaobaokeItemsCouponGetRequest) SetStartCommissionVolume(value string) {
	r.SetValue("start_commission_volume", value)
}

/* 设置折扣比例范围下限,如：7000表示70.00% */
func (r *TaobaokeItemsCouponGetRequest) SetStartCouponRate(value string) {
	r.SetValue("start_coupon_rate", value)
}

/* 卖家信用: 1heart(一心) 2heart (两心) 3heart(三心) 4heart(四心) 5heart(五心) 1diamond(一钻) 2diamond(两钻) 3diamond(三钻) 4diamond(四钻) 5diamond(五钻) 1crown(一冠) 2crown(两冠) 3crown(三冠) 4crown(四冠) 5crown(五冠) 1goldencrown(一黄冠) 2goldencrown(二黄冠) 3goldencrown(三黄冠) 4goldencrown(四黄冠) 5goldencrown(五黄冠) */
func (r *TaobaokeItemsCouponGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 设置商品总成交量（与返回字段volume对应）下限。 */
func (r *TaobaokeItemsCouponGetRequest) SetStartVolume(value string) {
	r.SetValue("start_volume", value)
}

func (r *TaobaokeItemsCouponGetRequest) GetResponse(accessToken string) (*TaobaokeItemsCouponGetResponse, []byte, error) {
	var resp TaobaokeItemsCouponGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.items.coupon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeItemsCouponGetResponse struct {
	TaobaokeItems []*TaobaokeItem `json:"taobaoke_items"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeItemsCouponGetResponseResult struct {
	Response *TaobaokeItemsCouponGetResponse `json:"taobaoke_items_coupon_get_response"`
}

/* 查询淘宝客推广商品详细信息。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeItemsDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:TaobaokeItemDetail淘宝客商品结构体中的所有字段;字段之间用","分隔。item_detail需要设置到Item模型下的字段,如设置:num_iid,detail_url等; 只设置item_detail,则不返回的Item下的所有信息.注：item结构中的skus、videos、props_name不返回 */
func (r *TaobaokeItemsDetailGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TaobaokeItemsDetailGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 淘宝用户昵称，注：指的是淘宝的会员登录名.如果昵称错误,那么客户就收不到佣金.每个淘宝昵称都对应于一个pid，在这里输入要结算佣金的淘宝昵称，当推广的商品成功后，佣金会打入此输入的淘宝昵称的账户。具体的信息可以登入阿里妈妈的网站查看. */
func (r *TaobaokeItemsDetailGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 淘宝客商品数字id串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔商品id. */
func (r *TaobaokeItemsDetailGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeItemsDetailGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeItemsDetailGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 点击串跳转类型，1：单品，2：单品中间页（无线暂无） */
func (r *TaobaokeItemsDetailGetRequest) SetReferType(value string) {
	r.SetValue("refer_type", value)
}

/* 商品track_iid串（带有追踪效果的商品id),最大输入10个,与num_iids必填其一 */
func (r *TaobaokeItemsDetailGetRequest) SetTrackIids(value string) {
	r.SetValue("track_iids", value)
}

func (r *TaobaokeItemsDetailGetRequest) GetResponse(accessToken string) (*TaobaokeItemsDetailGetResponse, []byte, error) {
	var resp TaobaokeItemsDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.items.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeItemsDetailGetResponse struct {
	TaobaokeItemDetails []*TaobaokeItemDetail `json:"taobaoke_item_details"`
	TotalResults        int                   `json:"total_results"`
}

type TaobaokeItemsDetailGetResponseResult struct {
	Response *TaobaokeItemsDetailGetResponse `json:"taobaoke_items_detail_get_response"`
}

/* 查询淘宝客推广商品,不能通过设置cid=0来查询。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeItemsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 商品所在地 */
func (r *TaobaokeItemsGetRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 是否自动发货 */
func (r *TaobaokeItemsGetRequest) SetAutoSend(value string) {
	r.SetValue("auto_send", value)
}

/* 是否支持抵价券，设置为true表示该商品支持抵价券，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetCashCoupon(value string) {
	r.SetValue("cash_coupon", value)
}

/* 是否支持货到付款，设置为true表示该商品是支持货到付款，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetCashOndelivery(value string) {
	r.SetValue("cash_ondelivery", value)
}

/* 标准商品后台类目id。该ID可以通过taobao.itemcats.get接口获取到。 */
func (r *TaobaokeItemsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 30天累计推广量（与返回数据中的commission_num字段对应）上限. */
func (r *TaobaokeItemsGetRequest) SetEndCommissionNum(value string) {
	r.SetValue("end_commissionNum", value)
}

/* 佣金比率上限，如：2345表示23.45%。注：start_commissionRate和end_commissionRate一起设置才有效。 */
func (r *TaobaokeItemsGetRequest) SetEndCommissionRate(value string) {
	r.SetValue("end_commissionRate", value)
}

/* 可选值和start_credit一样.start_credit的值一定要小于或等于end_credit的值。注：end_credit与start_credit一起使用才生效 */
func (r *TaobaokeItemsGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 最高价格 */
func (r *TaobaokeItemsGetRequest) SetEndPrice(value string) {
	r.SetValue("end_price", value)
}

/* 商品总成交量（与返回字段volume对应）上限。 */
func (r *TaobaokeItemsGetRequest) SetEndTotalnum(value string) {
	r.SetValue("end_totalnum", value)
}

/* 需返回的字段列表.可选值:num_iid,title,nick,pic_url,price,click_url,commission,commission_rate,commission_num,commission_volume,shop_click_url,seller_credit_score,item_location,volume
;字段之间用","分隔 */
func (r *TaobaokeItemsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否查询消保卖家 */
func (r *TaobaokeItemsGetRequest) SetGuarantee(value string) {
	r.SetValue("guarantee", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TaobaokeItemsGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 商品标题中包含的关键字. 注意:查询时keyword,cid至少选择其中一个参数 */
func (r *TaobaokeItemsGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 是否商城的商品，设置为true表示该商品是属于淘宝商城的商品，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetMallItem(value string) {
	r.SetValue("mall_item", value)
}

/* 淘宝用户昵称，注：指的是淘宝的会员登录名.如果昵称错误,那么客户就收不到佣金.每个淘宝昵称都对应于一个pid，在这里输入要结算佣金的淘宝昵称，当推广的商品成功后，佣金会打入此输入的淘宝昵称的账户。具体的信息可以登入阿里妈妈的网站查看.
<font color="red">注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准</font> */
func (r *TaobaokeItemsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 是否30天维修，设置为true表示该商品是支持30天维修，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetOnemonthRepair(value string) {
	r.SetValue("onemonth_repair", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeItemsGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 是否海外商品，设置为true表示该商品是属于海外商品，默认为false */
func (r *TaobaokeItemsGetRequest) SetOverseasItem(value string) {
	r.SetValue("overseas_item", value)
}

/* 结果页数.1~10 */
func (r *TaobaokeItemsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数.最大每页40 */
func (r *TaobaokeItemsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx".
<font color="red">注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20</font>。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeItemsGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 是否如实描述(即:先行赔付)商品，设置为true表示该商品是如实描述的商品，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetRealDescribe(value string) {
	r.SetValue("real_describe", value)
}

/* 点击串跳转类型，1：单品，2：单品中间页（无线暂无） */
func (r *TaobaokeItemsGetRequest) SetReferType(value string) {
	r.SetValue("refer_type", value)
}

/* 是否支持7天退换，设置为true表示该商品支持7天退换，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetSevendaysReturn(value string) {
	r.SetValue("sevendays_return", value)
}

/* 默认排序:default

price_desc(价格从高到低)
price_asc(价格从低到高)
credit_desc(信用等级从高到低)
commissionRate_desc(佣金比率从高到低)
commissionRate_asc(佣金比率从低到高)
commissionNum_desc(成交量成高到低)
commissionNum_asc(成交量从低到高)
commissionVolume_desc(总支出佣金从高到低)
commissionVolume_asc(总支出佣金从低到高)
delistTime_desc(商品下架时间从高到低)
delistTime_asc(商品下架时间从低到高) */
func (r *TaobaokeItemsGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 30天累计推广量（与返回数据中的commission_num字段对应）下限.注：该字段要与end_commissionNum一起使用才生效 */
func (r *TaobaokeItemsGetRequest) SetStartCommissionNum(value string) {
	r.SetValue("start_commissionNum", value)
}

/* 佣金比率下限，如：1234表示12.34% */
func (r *TaobaokeItemsGetRequest) SetStartCommissionRate(value string) {
	r.SetValue("start_commissionRate", value)
}

/* 卖家信用:

1heart(一心)
2heart (两心)
3heart(三心)
4heart(四心)
5heart(五心)
1diamond(一钻)
2diamond(两钻)
3diamond(三钻)
4diamond(四钻)
5diamond(五钻)
1crown(一冠)
2crown(两冠)
3crown(三冠)
4crown(四冠)
5crown(五冠)
1goldencrown(一黄冠)
2goldencrown(二黄冠)
3goldencrown(三黄冠)
4goldencrown(四黄冠)
5goldencrown(五黄冠) */
func (r *TaobaokeItemsGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 起始价格.传入价格参数时,需注意起始价格和最高价格必须一起传入,并且 start_price <= end_price */
func (r *TaobaokeItemsGetRequest) SetStartPrice(value string) {
	r.SetValue("start_price", value)
}

/* 商品总成交量（与返回字段volume对应）下限。 */
func (r *TaobaokeItemsGetRequest) SetStartTotalnum(value string) {
	r.SetValue("start_totalnum", value)
}

/* 是否支持VIP卡，设置为true表示该商品支持VIP卡，设置为false或不设置表示不判断这个属性 */
func (r *TaobaokeItemsGetRequest) SetVipCard(value string) {
	r.SetValue("vip_card", value)
}

func (r *TaobaokeItemsGetRequest) GetResponse(accessToken string) (*TaobaokeItemsGetResponse, []byte, error) {
	var resp TaobaokeItemsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.items.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeItemsGetResponse struct {
	TaobaokeItems []*TaobaokeItem `json:"taobaoke_items"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeItemsGetResponseResult struct {
	Response *TaobaokeItemsGetResponse `json:"taobaoke_items_get_response"`
}

/* 商品关联推荐。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeItemsRelateGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 分类id.推荐类型为5时cid不能为空。仅支持叶子类目ID，即通过taobao.itemcats.get获取到is_parent=false的cid。 */
func (r *TaobaokeItemsRelateGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 需返回的字段列表.可选值:num_iid,title,nick,pic_url,price,click_url,commission,ommission_rate,commission_num,commission_volume,shop_click_url,seller_credit_score,item_location,volume;字段之间用","分隔 */
func (r *TaobaokeItemsRelateGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false */
func (r *TaobaokeItemsRelateGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 指定返回结果的最大条数.实际返回结果个数根据算法来确定,所以该值会小于或者等于该值 */
func (r *TaobaokeItemsRelateGetRequest) SetMaxCount(value string) {
	r.SetValue("max_count", value)
}

/* 推广者的淘宝会员昵称.注:指的是淘宝的会员登录名 */
func (r *TaobaokeItemsRelateGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 淘宝客商品数字id.推荐类型为1,2,3时num_iid不能为空 */
func (r *TaobaokeItemsRelateGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道 */
func (r *TaobaokeItemsRelateGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeItemsRelateGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 点击串跳转类型，1：单品，2：单品中间页（无线暂无） */
func (r *TaobaokeItemsRelateGetRequest) SetReferType(value string) {
	r.SetValue("refer_type", value)
}

/* <p>推荐类型.</p>
<p>1:同类商品推荐;此时必须得输入num_iid</p>
<p>2:异类商品推荐;此时必须得输入num_iid</p>
<p>3:同店商品推荐;此时必须得输入num_iid</p>
<p>4:店铺热门推荐;此时必须得输入seller_id，这里的seller_id得通过<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:10449">taobao.taobaoke.shops.get</a>
跟<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:21419">taobao.taobaoke.widget.shops.convert</a>这两个接口去获取user_id字段</p>
<p>5:类目热门推荐;此时必须得输入cid</p> */
func (r *TaobaokeItemsRelateGetRequest) SetRelateType(value string) {
	r.SetValue("relate_type", value)
}

/* 卖家的用户id，这里的seller_id得通过<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:10449">taobao.taobaoke.shops.get</a>
跟<a href="http://api.taobao.com/apidoc/api.htm?path=cid:38-apiId:21419">taobao.taobaoke.widget.shops.convert</a>这两个接口去获取user_id字段。
注：推荐类型为4时seller_id不能为空 */
func (r *TaobaokeItemsRelateGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}

/* 店铺类型.默认all,商城:b,集市:c */
func (r *TaobaokeItemsRelateGetRequest) SetShopType(value string) {
	r.SetValue("shop_type", value)
}

/* default(默认排序,关联推荐相关度),price_desc(价格从高到低), price_asc(价格从低到高),commissionRate_desc(佣金比率从高到低), commissionRate_asc(佣金比率从低到高), commissionNum_desc(成交量成高到低), commissionNum_asc(成交量从低到高) */
func (r *TaobaokeItemsRelateGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

/* 商品数字ID(带有跟踪效果) */
func (r *TaobaokeItemsRelateGetRequest) SetTrackIid(value string) {
	r.SetValue("track_iid", value)
}

func (r *TaobaokeItemsRelateGetRequest) GetResponse(accessToken string) (*TaobaokeItemsRelateGetResponse, []byte, error) {
	var resp TaobaokeItemsRelateGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.items.relate.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeItemsRelateGetResponse struct {
	TaobaokeItems []*TaobaokeItem `json:"taobaoke_items"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeItemsRelateGetResponseResult struct {
	Response *TaobaokeItemsRelateGetResponse `json:"taobaoke_items_relate_get_response"`
}

/* 淘宝客关键词搜索URL.
<p>由于受到<a href="http://club.alimama.com/read-htm-tid-3133847.html">淘宝联盟pid升级规则</a>影响，该接口后续可能会下线。建议直接使用拼链接的方式（塞入完整的pid）来获取搜索结果，而不是调用该接口来获取结果</p>
<p>拼接的格式为：http://s8.taobao.com/browse/search_auction.htm?q=$url_encode($keywords)&pid=mm_xxx_xxx_xxx&search_type=auction&commend=all&at_topsearch=1&unid=$outer_code&spm=2014.$appkey.$apmc.$spmd</p>
<p>q指关键字，url_encode是编码方法（具体根据不同编程语言而定）</p>
<p>pid是你自己淘宝账号的完整的pid，登录alimama.com去获取</p>
<p>unid是推广渠道。跟调用api设置outer_code效果相同</p>
<p>spm可以到文档搜索“spm”了解下</p>
<p></p>
<p></p>
<p></p> */
type TaobaokeListurlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 推广者的淘宝会员昵称.注：这里指的是淘宝的登录会员名 */
func (r *TaobaokeListurlGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeListurlGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20 */
func (r *TaobaokeListurlGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 关键词 */
func (r *TaobaokeListurlGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

func (r *TaobaokeListurlGetRequest) GetResponse(accessToken string) (*TaobaokeListurlGetResponse, []byte, error) {
	var resp TaobaokeListurlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.listurl.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeListurlGetResponse struct {
	TaobaokeItem *TaobaokeItem `json:"taobaoke_item"`
}

type TaobaokeListurlGetResponseResult struct {
	Response *TaobaokeListurlGetResponse `json:"taobaoke_listurl_get_response"`
}

/* 查询卖家是否支持返利 */
type TaobaokeRebateAuthorizeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 卖家淘宝会员昵称.注：指的是淘宝的会员登录名 */
func (r *TaobaokeRebateAuthorizeGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 商品数字ID */
func (r *TaobaokeRebateAuthorizeGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 卖家淘宝会员数字ID. */
func (r *TaobaokeRebateAuthorizeGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}

func (r *TaobaokeRebateAuthorizeGetRequest) GetResponse(accessToken string) (*TaobaokeRebateAuthorizeGetResponse, []byte, error) {
	var resp TaobaokeRebateAuthorizeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.rebate.authorize.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeRebateAuthorizeGetResponse struct {
	Rebate bool `json:"rebate"`
}

type TaobaokeRebateAuthorizeGetResponseResult struct {
	Response *TaobaokeRebateAuthorizeGetResponse `json:"taobaoke_rebate_authorize_get_response"`
}

/* 根据输入开始时间，时间跨度，查询90天以内订单成功且支持返利淘宝客报表 */
type TaobaokeRebateReportGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:TaobaokePayment淘宝客订单构体中的所有字段;字段之间用","分隔. */
func (r *TaobaokeRebateReportGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 当前页数 */
func (r *TaobaokeRebateReportGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数，最小每页40条，默认每页40条，最大每页100条 */
func (r *TaobaokeRebateReportGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询报表的时间跨度，单位秒。
以用户输入的start_time时间为起始时间，start_time+span为结束时间，查询该时间段内的订单。span最小值为60秒，最大值为600秒，默认值为60秒 */
func (r *TaobaokeRebateReportGetRequest) SetSpan(value string) {
	r.SetValue("span", value)
}

/* 需要查询报表的开始日期，有效的日期为从当前日期开始起90天以内的订单 */
func (r *TaobaokeRebateReportGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

func (r *TaobaokeRebateReportGetRequest) GetResponse(accessToken string) (*TaobaokeRebateReportGetResponse, []byte, error) {
	var resp TaobaokeRebateReportGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.rebate.report.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeRebateReportGetResponse struct {
	TaobaokePayments []*TaobaokePayment `json:"taobaoke_payments"`
}

type TaobaokeRebateReportGetResponseResult struct {
	Response *TaobaokeRebateReportGetResponse `json:"taobaoke_rebate_report_get_response"`
}

/* 淘客报表查询，报表接口只能获取到该appkey对应的创建者（淘宝账号）的推广数据。目前只支持获取订单成功的报表。订单创建、订单失效、退款订单请直接在www.alimama.com后台进行查询。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http:http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeReportGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 买家确认收货时间，有效的日期为最近3个月内的某一天，格式为:yyyyMMdd,如20090520. */
func (r *TaobaokeReportGetRequest) SetDate(value string) {
	r.SetValue("date", value)
}

/* 需返回的字段列表.可选值:TaobaokeReportMember淘宝客报表成员结构体中的所有字段;字段之间用","分隔. */
func (r *TaobaokeReportGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 当前页数.只能获取1-499页数据. */
func (r *TaobaokeReportGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页返回结果数,默认是40条.最大每页100 */
func (r *TaobaokeReportGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

func (r *TaobaokeReportGetRequest) GetResponse(accessToken string) (*TaobaokeReportGetResponse, []byte, error) {
	var resp TaobaokeReportGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.report.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeReportGetResponse struct {
	TaobaokeReport *TaobaokeReport `json:"taobaoke_report"`
}

type TaobaokeReportGetResponseResult struct {
	Response *TaobaokeReportGetResponse `json:"taobaoke_report_get_response"`
}

/* 提供对参加了淘客推广的店铺的搜索。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeShopsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 店铺前台展示类目id，可以通过taobao.shopcats.list.get获取。 */
func (r *TaobaokeShopsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 店铺商品数查询结束值。需要跟start_auctioncount同时设置才生效，只设置该值不生效。 */
func (r *TaobaokeShopsGetRequest) SetEndAuctioncount(value string) {
	r.SetValue("end_auctioncount", value)
}

/* 店铺佣金比例查询结束值 */
func (r *TaobaokeShopsGetRequest) SetEndCommissionrate(value string) {
	r.SetValue("end_commissionrate", value)
}

/* 店铺掌柜信用等级查询结束
店铺的信用等级总共为20级 1-5:1heart-5heart;6-10:1diamond-5diamond;11-15:1crown-5crown;16-20:1goldencrown-5goldencrown */
func (r *TaobaokeShopsGetRequest) SetEndCredit(value string) {
	r.SetValue("end_credit", value)
}

/* 店铺累计推广数查询结束值 */
func (r *TaobaokeShopsGetRequest) SetEndTotalaction(value string) {
	r.SetValue("end_totalaction", value)
}

/* 需要返回的字段，目前包括如下字段 user_id click_url shop_title commission_rate seller_credit shop_type auction_count total_auction */
func (r *TaobaokeShopsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TaobaokeShopsGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 店铺主题关键字查询 */
func (r *TaobaokeShopsGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 淘宝用户昵称，注：指的是淘宝的会员登录名.如果昵称错误,那么客户就收不到佣金.每个淘宝昵称都对应于一个pid，在这里输入要结算佣金的淘宝昵称，当推广的商品成功后，佣金会打入此输入的淘宝昵称的账户。具体的信息可以登入阿里妈妈的网站查看 */
func (r *TaobaokeShopsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 是否只显示商城店铺 */
func (r *TaobaokeShopsGetRequest) SetOnlyMall(value string) {
	r.SetValue("only_mall", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeShopsGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 页码.结果页1~99 */
func (r *TaobaokeShopsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.最大每页40 */
func (r *TaobaokeShopsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeShopsGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 排序字段。目前支持的排序字段有：
commission_rate，auction_count，total_auction。必须输入这三个任意值，否则排序无效 */
func (r *TaobaokeShopsGetRequest) SetSortField(value string) {
	r.SetValue("sort_field", value)
}

/* 排序类型.必须输入desc,asc任一值，否则无效
desc-降序,asc-升序 */
func (r *TaobaokeShopsGetRequest) SetSortType(value string) {
	r.SetValue("sort_type", value)
}

/* 店铺宝贝数查询开始值。需要跟end_auctioncount同时设置才生效，只设置该值不生效。 */
func (r *TaobaokeShopsGetRequest) SetStartAuctioncount(value string) {
	r.SetValue("start_auctioncount", value)
}

/* 店铺佣金比例查询开始值，注意佣金比例是x10000的整数.50表示0.5% */
func (r *TaobaokeShopsGetRequest) SetStartCommissionrate(value string) {
	r.SetValue("start_commissionrate", value)
}

/* 店铺掌柜信用等级起始
店铺的信用等级总共为20级 1-5:1heart-5heart;6-10:1diamond-5diamond;11-15:1crown-5crown;16-20:1goldencrown-5goldencrown */
func (r *TaobaokeShopsGetRequest) SetStartCredit(value string) {
	r.SetValue("start_credit", value)
}

/* 店铺累计推广量开始值 */
func (r *TaobaokeShopsGetRequest) SetStartTotalaction(value string) {
	r.SetValue("start_totalaction", value)
}

func (r *TaobaokeShopsGetRequest) GetResponse(accessToken string) (*TaobaokeShopsGetResponse, []byte, error) {
	var resp TaobaokeShopsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.shops.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeShopsGetResponse struct {
	TaobaokeShops []*TaobaokeShop `json:"taobaoke_shops"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeShopsGetResponseResult struct {
	Response *TaobaokeShopsGetResponse `json:"taobaoke_shops_get_response"`
}

/* 淘宝客店铺关联推荐。
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeShopsRelateGetRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:TaobaokeShop淘宝客商品结构体中的user_id,seller_nick,shop_id,shop_title,seller_credit,shop_type,commission_rate,click_url,total_auction,auction_count,字段之间用","分隔 */
func (r *TaobaokeShopsRelateGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串,如果不传值,则默认是false */
func (r *TaobaokeShopsRelateGetRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 指定返回结果的最大条数,实际返回结果个数根据算法来确定 */
func (r *TaobaokeShopsRelateGetRequest) SetMaxCount(value string) {
	r.SetValue("max_count", value)
}

/* 淘宝用户昵称.注:指的是淘宝的会员登录名.如果昵称错误,那么客户就收不到佣金.每个淘宝昵称都对应于一个pid,在这里输入要结算佣金的淘宝昵称,当推广的商品成功后,佣金会打入此输入的淘宝昵称的账户.具体的信息可以登入阿里妈妈的网站查看 */
func (r *TaobaokeShopsRelateGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道 */
func (r *TaobaokeShopsRelateGetRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 用户的pid,必须是mm_xxxx_0_0这种格式中间的"xxxx". 注意nick和pid至少需要传递一个,如果2个都传了,将以pid为准,且pid的最大长度是20。第一次调用接口的用户，推荐该入参不要填写，使用nick=（淘宝账号）的方式去获取，以免出错。 */
func (r *TaobaokeShopsRelateGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 卖家id.seller_id和seller_nick不能同时为空,如果都有值,则以seller_id为主 */
func (r *TaobaokeShopsRelateGetRequest) SetSellerId(value string) {
	r.SetValue("seller_id", value)
}

/* 卖家昵称 */
func (r *TaobaokeShopsRelateGetRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}

/* 店铺类型.所有:all,商城:b,集市:c */
func (r *TaobaokeShopsRelateGetRequest) SetShopType(value string) {
	r.SetValue("shop_type", value)
}

/* default(默认排序,关联推荐相关度),commissionRate_desc(佣金比率从高到低), commissionRate_asc(佣金比率从低到高),credit_desc(信用等级从高到低), credit_asc(信用等级从低到高) */
func (r *TaobaokeShopsRelateGetRequest) SetSort(value string) {
	r.SetValue("sort", value)
}

func (r *TaobaokeShopsRelateGetRequest) GetResponse(accessToken string) (*TaobaokeShopsRelateGetResponse, []byte, error) {
	var resp TaobaokeShopsRelateGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.shops.relate.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeShopsRelateGetResponse struct {
	TaobaokeShops []*TaobaokeShop `json:"taobaoke_shops"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeShopsRelateGetResponseResult struct {
	Response *TaobaokeShopsRelateGetResponse `json:"taobaoke_shops_relate_get_response"`
}

/* 淘客商品转换，该接口只支持js方式去调用，具体的调用方式参考文档【http://open.taobao.com/doc/detail.htm?id=101372】

<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeWidgetItemsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:num_iid,title,nick,pic_url,price,click_url,commission,commission_rate,commission_num,commission_volume,shop_click_url,seller_credit_score,item_location,volume
;字段之间用","分隔. */
func (r *TaobaokeWidgetItemsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来自无线或者手机应用,如果是true，则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TaobaokeWidgetItemsConvertRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 淘宝客商品数字id串.最大输入40个.格式如:"value1,value2,value3" 用" , "号分隔商品数字id */
func (r *TaobaokeWidgetItemsConvertRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeWidgetItemsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 商品track_iid串（带有追踪效果的商品id),最大输入40个,与num_iids必填其一 */
func (r *TaobaokeWidgetItemsConvertRequest) SetTrackIids(value string) {
	r.SetValue("track_iids", value)
}

func (r *TaobaokeWidgetItemsConvertRequest) GetResponse(accessToken string) (*TaobaokeWidgetItemsConvertResponse, []byte, error) {
	var resp TaobaokeWidgetItemsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.widget.items.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeWidgetItemsConvertResponse struct {
	TaobaokeItems []*TaobaokeItem `json:"taobaoke_items"`
	TotalResults  int             `json:"total_results"`
}

type TaobaokeWidgetItemsConvertResponseResult struct {
	Response *TaobaokeWidgetItemsConvertResponse `json:"taobaoke_widget_items_convert_response"`
}

/* 淘客店铺转换，该接口只支持js方式去调用，具体的调用方式参考文档【http://open.taobao.com/doc/detail.htm?id=101372】
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeWidgetShopsConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 需返回的字段列表.可选值:TaobaokeShop淘宝客商品结构体中的user_id,shop_title,click_url,commission_rate;字段之间用","分隔. */
func (r *TaobaokeWidgetShopsConvertRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 标识一个应用是否来在无线或者手机应用,如果是true则会使用其他规则加密点击串.如果不传值,则默认是false. */
func (r *TaobaokeWidgetShopsConvertRequest) SetIsMobile(value string) {
	r.SetValue("is_mobile", value)
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道. */
func (r *TaobaokeWidgetShopsConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 卖家昵称串.最大输入10个.格式如:"value1,value2,value3" 用" , "号分隔。 */
func (r *TaobaokeWidgetShopsConvertRequest) SetSellerNicks(value string) {
	r.SetValue("seller_nicks", value)
}

func (r *TaobaokeWidgetShopsConvertRequest) GetResponse(accessToken string) (*TaobaokeWidgetShopsConvertResponse, []byte, error) {
	var resp TaobaokeWidgetShopsConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.widget.shops.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeWidgetShopsConvertResponse struct {
	TaobaokeShops []*TaobaokeShop `json:"taobaoke_shops"`
}

type TaobaokeWidgetShopsConvertResponseResult struct {
	Response *TaobaokeWidgetShopsConvertResponse `json:"taobaoke_widget_shops_convert_response"`
}

/* 活动推广API可以把天猫、聚划算、淘宝旅行、淘宝游戏等平台的活动链接转换为淘宝客推广链接，该接口只支持js方式去调用，具体的调用方式参考文档【http://open.taobao.com/doc/detail.htm?id=101372】
<p>淘宝客应用、网站接入的过程中，难免会遇到问题，这里对从接入到线上运营的各个环节最常碰到的问题，做了汇总，帮助开发者提高接入的效率。 </p> <p>一、淘宝客网站应用创建流程：<a href="http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028">http://open.taobao.com/doc/detail.htm?spm=0.0.0.34.b1f9de&id=1028</a></p> <p>二、淘宝客API结合实际使用场景的介绍：<a href="http://open.taobao.com/doc/detail.htm?id=1014">http://open.taobao.com/doc/detail.htm?id=1014</a></p> <p>三、淘宝客网站官方推荐的架构：<a href="http://open.taobao.com/doc/detail.htm?id=1011">http://open.taobao.com/doc/detail.htm?id=1011</a></p> <p>四、淘宝客最常见的几个问题以及解决方案汇总：<a href="http://open.taobao.com/doc/detail.htm?id=1005">http://open.taobao.com/doc/detail.htm?id=1005</a></p> */
type TaobaokeWidgetUrlConvertRequest struct {
	open_taobao.TaobaoMethodRequest
}

/* 自定义输入串.格式:英文和数字组成;长度不能大于12个字符,区分不同的推广渠道,如:bbs,表示bbs为推广渠道;blog,表示blog为推广渠道 */
func (r *TaobaokeWidgetUrlConvertRequest) SetOuterCode(value string) {
	r.SetValue("outer_code", value)
}

/* 需要转化为淘客链接的url，可以把天猫、聚划算、淘宝旅行、淘宝游戏等平台的活动链接转换为淘宝客推广链接 */
func (r *TaobaokeWidgetUrlConvertRequest) SetUrl(value string) {
	r.SetValue("url", value)
}

func (r *TaobaokeWidgetUrlConvertRequest) GetResponse(accessToken string) (*TaobaokeWidgetUrlConvertResponse, []byte, error) {
	var resp TaobaokeWidgetUrlConvertResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.taobaoke.widget.url.convert", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TaobaokeWidgetUrlConvertResponse struct {
	TaobaokeItem *TaobaokeItem `json:"taobaoke_item"`
}

type TaobaokeWidgetUrlConvertResponseResult struct {
	Response *TaobaokeWidgetUrlConvertResponse `json:"taobaoke_widget_url_convert_response"`
}
