// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package product

import (
	"github.com/yaofangou/open_taobao"
)





/* 查询用户设置的售后服务模板，仅返回标题和id */
type AftersaleGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *AftersaleGetRequest) GetResponse(accessToken string) (*AftersaleGetResponse, []byte, error) {
	var resp AftersaleGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.aftersale.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AftersaleGetResponse struct {
	AfterSales *AfterSaleListObject `json:"after_sales"`
}

type AftersaleGetResponseResult struct {
	Response *AftersaleGetResponse `json:"aftersale_get_response"`
}





/* 此接口用于新增一个商品  
商品所属的卖家是当前会话的用户  
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则新增商品会失败）  
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得）  
商品的运费承担方式和邮费设置有相关性，卖家承担运费不用设置邮费，买家承担运费需要设置邮费  
当关键属性值选择了“其他”的时候，需要输入input_pids和input_str商品才能添加成功。 */
type ItemAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 售后说明模板id */
func (r *ItemAddRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale */
func (r *ItemAddRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (r *ItemAddRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： 
no_mark(不做类型标记) 
time_card(点卡软件代充) 
fee_card(话费软件代充) */
func (r *ItemAddRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 商品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1;
基础色只支持以下14种颜色：28335//绿色
28338//蓝色
90554//桔色
28324//黄色
28341//黑色
28320//白色
28326//红色
28329//紫色
3232480//粉红色
107121//透明
132069//褐色
28332//浅灰色
3232478//深灰色
130164//花色 */
func (r *ItemAddRequest) SetChangeProp(value string) {
	r.SetValue("change_prop", value)
}

/* 叶子类目id */
func (r *ItemAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。 */
func (r *ItemAddRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
func (r *ItemAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 商品描述模块化，模块列表，具体数据结构参见Item_Desc_Module。详细的使用方法可参考下面FAQ中说明。 */
func (r *ItemAddRequest) SetDescModules(value string) {
	r.SetValue("desc_modules", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemAddRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemAddRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 厂家联系方式 */
func (r *ItemAddRequest) SetFoodSecurityContact(value string) {
	r.SetValue("food_security.contact", value)
}

/* 产品标准号 */
func (r *ItemAddRequest) SetFoodSecurityDesignCode(value string) {
	r.SetValue("food_security.design_code", value)
}

/* 厂名 */
func (r *ItemAddRequest) SetFoodSecurityFactory(value string) {
	r.SetValue("food_security.factory", value)
}

/* 厂址 */
func (r *ItemAddRequest) SetFoodSecurityFactorySite(value string) {
	r.SetValue("food_security.factory_site", value)
}

/* 食品添加剂 */
func (r *ItemAddRequest) SetFoodSecurityFoodAdditive(value string) {
	r.SetValue("food_security.food_additive", value)
}

/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
func (r *ItemAddRequest) SetFoodSecurityHealthProductNo(value string) {
	r.SetValue("food_security.health_product_no", value)
}

/* 配料表 */
func (r *ItemAddRequest) SetFoodSecurityMix(value string) {
	r.SetValue("food_security.mix", value)
}

/* 保质期 */
func (r *ItemAddRequest) SetFoodSecurityPeriod(value string) {
	r.SetValue("food_security.period", value)
}

/* 储藏方法 */
func (r *ItemAddRequest) SetFoodSecurityPlanStorage(value string) {
	r.SetValue("food_security.plan_storage", value)
}

/* 生产许可证号 */
func (r *ItemAddRequest) SetFoodSecurityPrdLicenseNo(value string) {
	r.SetValue("food_security.prd_license_no", value)
}

/* 生产结束日期,格式必须为yyyy-MM-dd */
func (r *ItemAddRequest) SetFoodSecurityProductDateEnd(value string) {
	r.SetValue("food_security.product_date_end", value)
}

/* 生产开始日期，格式必须为yyyy-MM-dd */
func (r *ItemAddRequest) SetFoodSecurityProductDateStart(value string) {
	r.SetValue("food_security.product_date_start", value)
}

/* 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemAddRequest) SetFoodSecurityStockDateEnd(value string) {
	r.SetValue("food_security.stock_date_end", value)
}

/* 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemAddRequest) SetFoodSecurityStockDateStart(value string) {
	r.SetValue("food_security.stock_date_start", value)
}

/* 供货商 */
func (r *ItemAddRequest) SetFoodSecuritySupplier(value string) {
	r.SetValue("food_security.supplier", value)
}

/* 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id  
如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee） */
func (r *ItemAddRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他） */
func (r *ItemAddRequest) SetGlobalStockCountry(value string) {
	r.SetValue("global_stock_country", value)
}

/* 全球购商品采购地（库存类型）， 
有两种库存类型：现货和代购 
参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
func (r *ItemAddRequest) SetGlobalStockType(value string) {
	r.SetValue("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折) */
func (r *ItemAddRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) */
func (r *ItemAddRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐) */
func (r *ItemAddRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false;默认值:false(不保修) */
func (r *ItemAddRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 商品主图片。类型:JPG,GIF;最大长度:500K */
func (r *ItemAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
func (r *ItemAddRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemAddRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节 */
func (r *ItemAddRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D */
func (r *ItemAddRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外店显示 */
func (r *ItemAddRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货 */
func (r *ItemAddRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
func (r *ItemAddRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 */
func (r *ItemAddRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。 
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。 
在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m) 
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m） */
func (r *ItemAddRequest) SetItemSize(value string) {
	r.SetValue("item_size", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 
只能传入数值类型（包含小数），不能带单位，单位默认为kg。 */
func (r *ItemAddRequest) SetItemWeight(value string) {
	r.SetValue("item_weight", value)
}

/* 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 */
func (r *ItemAddRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
func (r *ItemAddRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
func (r *ItemAddRequest) SetLocalityLifeChooseLogis(value string) {
	r.SetValue("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期; 
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 
如果有效期为【购买成功日 至】类型则格式为2012-08-16 
如果有效期为天数类型则格式为15 */
func (r *ItemAddRequest) SetLocalityLifeExpirydate(value string) {
	r.SetValue("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (r *ItemAddRequest) SetLocalityLifeMerchant(value string) {
	r.SetValue("locality_life.merchant", value)
}

/* 网点ID */
func (r *ItemAddRequest) SetLocalityLifeNetworkId(value string) {
	r.SetValue("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (r *ItemAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	r.SetValue("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例， 
百分比%前的数字,1-100的正整数值 */
func (r *ItemAddRequest) SetLocalityLifeRefundRatio(value string) {
	r.SetValue("locality_life.refund_ratio", value)
}

/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
func (r *ItemAddRequest) SetLocalityLifeRefundmafee(value string) {
	r.SetValue("locality_life.refundmafee", value)
}

/* 核销打款  
1代表核销打款 0代表非核销打款 */
func (r *ItemAddRequest) SetLocalityLifeVerification(value string) {
	r.SetValue("locality_life.verification", value)
}

/* 所在地城市。如杭州 。可以通过http://dl.open.taobao.com/sdk/商品城市列表.rar查询 */
func (r *ItemAddRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江，具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar  取到 */
func (r *ItemAddRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和。 
拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。 */
func (r *ItemAddRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商品外部编码，该字段的最大长度是512个字节 */
func (r *ItemAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。 
对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。 */
func (r *ItemAddRequest) SetPaimaiInfoDeposit(value string) {
	r.SetValue("paimai_info.deposit", value)
}

/* 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。 */
func (r *ItemAddRequest) SetPaimaiInfoInterval(value string) {
	r.SetValue("paimai_info.interval", value)
}

/* 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。 */
func (r *ItemAddRequest) SetPaimaiInfoMode(value string) {
	r.SetValue("paimai_info.mode", value)
}

/* 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
func (r *ItemAddRequest) SetPaimaiInfoReserve(value string) {
	r.SetValue("paimai_info.reserve", value)
}

/* 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
func (r *ItemAddRequest) SetPaimaiInfoValidHour(value string) {
	r.SetValue("paimai_info.valid_hour", value)
}

/* 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
func (r *ItemAddRequest) SetPaimaiInfoValidMinute(value string) {
	r.SetValue("paimai_info.valid_minute", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *ItemAddRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写 */
func (r *ItemAddRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板） */
func (r *ItemAddRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 
拍卖商品对应的起拍价。 */
func (r *ItemAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (r *ItemAddRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过511字节 */
func (r *ItemAddRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值 */
func (r *ItemAddRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
func (r *ItemAddRequest) SetScenicTicketBookCost(value string) {
	r.SetValue("scenic_ticket_book_cost", value)
}

/* 景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
func (r *ItemAddRequest) SetScenicTicketPayWay(value string) {
	r.SetValue("scenic_ticket_pay_way", value)
}

/* 商品卖点信息，最长150个字符。仅天猫商家可用。 */
func (r *ItemAddRequest) SetSellPoint(value string) {
	r.SetValue("sell_point", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemAddRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemAddRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* Sku的外部id串，结构如：1234,1342,… 
sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (r *ItemAddRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemAddRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性（自定义的除外）也需要在props字段填写。sku的销售属性需要一同选取，如:颜色，尺寸。如果新增商品包含了sku，则此字段一定要传入。这个字段的长度要控制在512个字节以内。 
如果有自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在冒号:和分号;以及逗号， */
func (r *ItemAddRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* Sku的数量串，结构如：num1,num2,num3 如：2,3 */
func (r *ItemAddRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 此参数暂时不起作用 */
func (r *ItemAddRequest) SetSkuSpecIds(value string) {
	r.SetValue("sku_spec_ids", value)
}

/* 新旧程度。可选值：new(新)，second(二手)，unused(闲置)。B商家不能发布二手商品。 
如果是二手商品，特定类目下属性里面必填新旧成色属性 */
func (r *ItemAddRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改
集市卖家默认拍下减库存;
商城卖家默认付款减库存 */
func (r *ItemAddRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符； */
func (r *ItemAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。 
拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。 
注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。 */
func (r *ItemAddRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 有效期。可选值:7,14;单位:天;默认值:14 */
func (r *ItemAddRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemAddRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}


func (r *ItemAddRequest) GetResponse(accessToken string) (*ItemAddResponse, []byte, error) {
	var resp ItemAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemAddResponse struct {
	Item *Item `json:"item"`
}

type ItemAddResponseResult struct {
	Response *ItemAddResponse `json:"item_add_response"`
}





/* 在新的发布模式下，isv需要先获取一份发布规则，然后根据规则传入数据。该接口提供规则查询功能 */
type ItemAddRulesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 发布宝贝的叶子类目id */
func (r *ItemAddRulesGetRequest) SetCategoryId(value string) {
	r.SetValue("category_id", value)
}


func (r *ItemAddRulesGetRequest) GetResponse(accessToken string) (*ItemAddRulesGetResponse, []byte, error) {
	var resp ItemAddRulesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.add.rules.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemAddRulesGetResponse struct {
	AddRules string `json:"add_rules"`
}

type ItemAddRulesGetResponseResult struct {
	Response *ItemAddRulesGetResponse `json:"item_add_rules_get_response"`
}





/* 根据类目id和宝贝描述规范化打标类型获取该类目可用的宝贝描述模块中的锚点 */
type ItemAnchorGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 对应类目编号 */
func (r *ItemAnchorGetRequest) SetCatId(value string) {
	r.SetValue("cat_id", value)
}

/* 宝贝模板类型是人工打标还是自动打标：人工打标为1，自动打标为0.人工和自动打标为-1. */
func (r *ItemAnchorGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *ItemAnchorGetRequest) GetResponse(accessToken string) (*ItemAnchorGetResponse, []byte, error) {
	var resp ItemAnchorGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.anchor.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemAnchorGetResponse struct {
	AnchorModules *IdsModuleListObject `json:"anchor_modules"`
	TotalResults int `json:"total_results"`
}

type ItemAnchorGetResponseResult struct {
	Response *ItemAnchorGetResponse `json:"item_anchor_get_response"`
}





/* 淘宝助理提供的发布商城商品接口，在发布时 先查询是否有这个产品，有则将商品绑定到该产品上发布；如果没有这个产品，自动帮用户新建产品，再将商品绑定到该产品上发布。错误码参考taobao.product.add、taobao.product.img.upload、taobao.product.propimg.upload、taobao.item.add */
type ItemBsellerAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 售后服务说明模板id */
func (r *ItemBsellerAddRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale */
func (r *ItemBsellerAddRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数 */
func (r *ItemBsellerAddRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：  
*no_mark(不做类型标记)  
*time_card(点卡软件代充)  
*fee_card(话费软件代充) */
func (r *ItemBsellerAddRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 自动重发。可选值:true,false;默认值:false(不重发) */
func (r *ItemBsellerAddRequest) SetAutoRepost(value string) {
	r.SetValue("auto_repost", value)
}

/* 叶子类目id */
func (r *ItemBsellerAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 货到付款运费模板ID */
func (r *ItemBsellerAddRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 颜色销售属性和对应销售属性图片url列表，格式为"颜色属性pid,颜色属性值vid,颜色图片url"，多个颜色用分号分隔，例如：125,567,http://img05.daily.taobao.net/bao/uploaded/i5/T16.lbXl02XXXmYQgY_030226.jpg;566,578,http://img05.daily.taobao.net/bao/uploaded/i5/T1zphcXcVuXXbDdSva_122254.jpg; */
func (r *ItemBsellerAddRequest) SetColorPropAndPicUrls(value string) {
	r.SetValue("color_prop_and_pic_urls", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
func (r *ItemBsellerAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* ems费用。取值范围:0-100000000;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemBsellerAddRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-10000.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemBsellerAddRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 宝贝特征值，格式为： 
【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp */
func (r *ItemBsellerAddRequest) SetFeatures(value string) {
	r.SetValue("features", value)
}

/* 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id  
如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee */
func (r *ItemBsellerAddRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 针对全球购卖家的库存类型业务， 
有两种库存类型：现货和代购 
参数值为1时代表现货，值为2时代表代购 
如果传值为这两个值之外的值，会报错; 
如果不是全球购卖家，这两个值即使设置也不会处理 */
func (r *ItemBsellerAddRequest) SetGlobalStockType(value string) {
	r.SetValue("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折) */
func (r *ItemBsellerAddRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) */
func (r *ItemBsellerAddRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐) */
func (r *ItemBsellerAddRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false;默认值:false(不保修) */
func (r *ItemBsellerAddRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 加价幅度。如果为0，代表系统代理幅度 */
func (r *ItemBsellerAddRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户的内店装修模板id。 */
func (r *ItemBsellerAddRequest) SetInnerShopAuctionTemplateId(value string) {
	r.SetValue("inner_shop_auction_template_id", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemBsellerAddRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节 */
func (r *ItemBsellerAddRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D商品 */
func (r *ItemBsellerAddRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外部网店显示 */
func (r *ItemBsellerAddRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货 */
func (r *ItemBsellerAddRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否在淘宝显示 */
func (r *ItemBsellerAddRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 */
func (r *ItemBsellerAddRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。 该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。 在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m) */
func (r *ItemBsellerAddRequest) SetItemSize(value string) {
	r.SetValue("item_size", value)
}

/* 特定种类的商品约束，以key：value的形式传入，以分号间隔。如食品安全需要输入：food_security.prd_license_no：12345;food_security.design_code：444;...。这些key：value可参考taobao.item.add的对应输入参数。 */
func (r *ItemBsellerAddRequest) SetItemSpecProp(value string) {
	r.SetValue("item_spec_prop", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 */
func (r *ItemBsellerAddRequest) SetItemWeight(value string) {
	r.SetValue("item_weight", value)
}

/* 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 */
func (r *ItemBsellerAddRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
func (r *ItemBsellerAddRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
func (r *ItemBsellerAddRequest) SetLocalityLifeChooseLogis(value string) {
	r.SetValue("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期; 
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 
如果有效期为【购买成功日 至】类型则格式为2012-08-16 
如果有效期为天数类型则格式为15 */
func (r *ItemBsellerAddRequest) SetLocalityLifeExpirydate(value string) {
	r.SetValue("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (r *ItemBsellerAddRequest) SetLocalityLifeMerchant(value string) {
	r.SetValue("locality_life.merchant", value)
}

/* 网点ID */
func (r *ItemBsellerAddRequest) SetLocalityLifeNetworkId(value string) {
	r.SetValue("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (r *ItemBsellerAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	r.SetValue("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例， 
百分比%前的数字,1-100的正整数值 */
func (r *ItemBsellerAddRequest) SetLocalityLifeRefundRatio(value string) {
	r.SetValue("locality_life.refund_ratio", value)
}

/* 核销打款  
1代表核销打款 0代表非核销打款 */
func (r *ItemBsellerAddRequest) SetLocalityLifeVerification(value string) {
	r.SetValue("locality_life.verification", value)
}

/* 所在地城市。如杭州 。具体可以用taobao.areas.get取到 */
func (r *ItemBsellerAddRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江，具体可以用taobao.areas.get取到 */
func (r *ItemBsellerAddRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和 */
func (r *ItemBsellerAddRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商家编码 */
func (r *ItemBsellerAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 用户的外店装修模板id */
func (r *ItemBsellerAddRequest) SetOuterShopAuctionTemplateId(value string) {
	r.SetValue("outer_shop_auction_template_id", value)
}

/* 商品主图在图片空间的完整url。该图片必须属于当前用户。 */
func (r *ItemBsellerAddRequest) SetPicUrl(value string) {
	r.SetValue("pic_url", value)
}

/* 平邮费用。取值范围:0.01-10000.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写 */
func (r *ItemBsellerAddRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板） */
func (r *ItemBsellerAddRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 */
func (r *ItemBsellerAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (r *ItemBsellerAddRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过511字节 */
func (r *ItemBsellerAddRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值 */
func (r *ItemBsellerAddRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 秒杀商品类型。暂时不能使用。打上秒杀标记的商品，用户只能下架并不能再上架，其他任何编辑或删除操作都不能进行。如果用户想取消秒杀标记，需要联系小二进行操作。如果秒杀结束需要自由编辑请联系活动负责人（小二）去掉秒杀标记。可选类型 
web_only(只能通过web网络秒杀) 
wap_only(只能通过wap网络秒杀) 
web_and_wap(既能通过web秒杀也能通过wap秒杀) */
func (r *ItemBsellerAddRequest) SetSecondKill(value string) {
	r.SetValue("second_kill", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemBsellerAddRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemBsellerAddRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids可写成：“, ,” */
func (r *ItemBsellerAddRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemBsellerAddRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性(自定义的除外)也需要在props字段填写 . 规则：如果该SKU存在旧商品，则修改；否则新增Sku。如果更新时有对Sku进行操作，则Sku的properties一定要传入。如果存在自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText，其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号 */
func (r *ItemBsellerAddRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* Sku的数量串，结构如：num1,num2,num3 如：2,3 */
func (r *ItemBsellerAddRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 产品的规格信息。 */
func (r *ItemBsellerAddRequest) SetSkuSpecIds(value string) {
	r.SetValue("sku_spec_ids", value)
}

/* 新旧程度。可选值：new(新)，second(二手)，unused(闲置)。B商家不能发布二手商品。 
如果是二手商品，特定类目下属性里面必填新旧成色属性 */
func (r *ItemBsellerAddRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 1~4个淘宝图片空间的图片url列表，“,”分隔，做为商品的副图 */
func (r *ItemBsellerAddRequest) SetSubPicUrls(value string) {
	r.SetValue("sub_pic_urls", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持;0(默认)不更改 */
func (r *ItemBsellerAddRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题。不能超过60字符，受违禁词控制 */
func (r *ItemBsellerAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的 */
func (r *ItemBsellerAddRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 有效期。可选值:7,14;单位:天;默认值:14 */
func (r *ItemBsellerAddRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemBsellerAddRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}


func (r *ItemBsellerAddRequest) GetResponse(accessToken string) (*ItemBsellerAddResponse, []byte, error) {
	var resp ItemBsellerAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.bseller.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemBsellerAddResponse struct {
	Item *Item `json:"item"`
}

type ItemBsellerAddResponseResult struct {
	Response *ItemBsellerAddResponse `json:"item_bseller_add_response"`
}





/* 提供的发布集市商品接口，接口参数除了包括taobao.item.vip.add的参数外，新增一个sub_pic_paths参数，为图片空间的url，本方法会在发布商品同时，将这些图片作为副图关联到新商品。 */
type ItemCsellerAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 售后服务说明模板id */
func (r *ItemCsellerAddRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale */
func (r *ItemCsellerAddRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：  
*no_mark(不做类型标记)  
*time_card(点卡软件代充)  
*fee_card(话费软件代充) */
func (r *ItemCsellerAddRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 自动重发。可选值:true,false;默认值:false(不重发) */
func (r *ItemCsellerAddRequest) SetAutoRepost(value string) {
	r.SetValue("auto_repost", value)
}

/* 叶子类目id */
func (r *ItemCsellerAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 货到付款运费模板ID */
func (r *ItemCsellerAddRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
func (r *ItemCsellerAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* ems费用。取值范围:0-100000000;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemCsellerAddRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-10000.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemCsellerAddRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 宝贝特征值，格式为： 
【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp */
func (r *ItemCsellerAddRequest) SetFeatures(value string) {
	r.SetValue("features", value)
}

/* 运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id  
如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee */
func (r *ItemCsellerAddRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 针对全球购卖家的库存类型业务， 
有两种库存类型：现货和代购 
参数值为1时代表现货，值为2时代表代购 
如果传值为这两个值之外的值，会报错; 
如果不是全球购卖家，这两个值即使设置也不会处理 */
func (r *ItemCsellerAddRequest) SetGlobalStockType(value string) {
	r.SetValue("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false;默认值:false(不打折) */
func (r *ItemCsellerAddRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) */
func (r *ItemCsellerAddRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false;默认值:false(不推荐) */
func (r *ItemCsellerAddRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false;默认值:false(不保修) */
func (r *ItemCsellerAddRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 加价幅度。如果为0，代表系统代理幅度 */
func (r *ItemCsellerAddRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户的内店装修模板id。 */
func (r *ItemCsellerAddRequest) SetInnerShopAuctionTemplateId(value string) {
	r.SetValue("inner_shop_auction_template_id", value)
}

/* 用户自行输入的类目属性ID串。结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemCsellerAddRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节 */
func (r *ItemCsellerAddRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D商品 */
func (r *ItemCsellerAddRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外部网店显示 */
func (r *ItemCsellerAddRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货 */
func (r *ItemCsellerAddRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否在淘宝显示 */
func (r *ItemCsellerAddRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 */
func (r *ItemCsellerAddRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。 该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。 在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m) 该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m） */
func (r *ItemCsellerAddRequest) SetItemSize(value string) {
	r.SetValue("item_size", value)
}

/* 特定种类的商品约束，以key：value的形式传入，以分号间隔。如食品安全需要输入：food_security.prd_license_no：12345;food_security.design_code：444;...。这些key：value可参考taobao.item.add的对应输入参数。 */
func (r *ItemCsellerAddRequest) SetItemSpecProp(value string) {
	r.SetValue("item_spec_prop", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 */
func (r *ItemCsellerAddRequest) SetItemWeight(value string) {
	r.SetValue("item_weight", value)
}

/* 商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 */
func (r *ItemCsellerAddRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss) */
func (r *ItemCsellerAddRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄 */
func (r *ItemCsellerAddRequest) SetLocalityLifeChooseLogis(value string) {
	r.SetValue("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期; 
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 
如果有效期为【购买成功日 至】类型则格式为2012-08-16 
如果有效期为天数类型则格式为15 */
func (r *ItemCsellerAddRequest) SetLocalityLifeExpirydate(value string) {
	r.SetValue("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (r *ItemCsellerAddRequest) SetLocalityLifeMerchant(value string) {
	r.SetValue("locality_life.merchant", value)
}

/* 网点ID */
func (r *ItemCsellerAddRequest) SetLocalityLifeNetworkId(value string) {
	r.SetValue("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (r *ItemCsellerAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	r.SetValue("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例， 
百分比%前的数字,1-100的正整数值 */
func (r *ItemCsellerAddRequest) SetLocalityLifeRefundRatio(value string) {
	r.SetValue("locality_life.refund_ratio", value)
}

/* 核销打款  
1代表核销打款 0代表非核销打款 */
func (r *ItemCsellerAddRequest) SetLocalityLifeVerification(value string) {
	r.SetValue("locality_life.verification", value)
}

/* 所在地城市。如杭州 。具体可以用taobao.areas.get取到 */
func (r *ItemCsellerAddRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江，具体可以用taobao.areas.get取到 */
func (r *ItemCsellerAddRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和 */
func (r *ItemCsellerAddRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商家编码 */
func (r *ItemCsellerAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 用户的外店装修模板id */
func (r *ItemCsellerAddRequest) SetOuterShopAuctionTemplateId(value string) {
	r.SetValue("outer_shop_auction_template_id", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *ItemCsellerAddRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 平邮费用。取值范围:0.01-10000.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写 */
func (r *ItemCsellerAddRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板） */
func (r *ItemCsellerAddRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 */
func (r *ItemCsellerAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过511字节 */
func (r *ItemCsellerAddRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值 */
func (r *ItemCsellerAddRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 秒杀商品类型。暂时不能使用。打上秒杀标记的商品，用户只能下架并不能再上架，其他任何编辑或删除操作都不能进行。如果用户想取消秒杀标记，需要联系小二进行操作。如果秒杀结束需要自由编辑请联系活动负责人（小二）去掉秒杀标记。可选类型 
web_only(只能通过web网络秒杀) 
wap_only(只能通过wap网络秒杀) 
web_and_wap(既能通过web秒杀也能通过wap秒杀) */
func (r *ItemCsellerAddRequest) SetSecondKill(value string) {
	r.SetValue("second_kill", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemCsellerAddRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemCsellerAddRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids可写成：“, ,” */
func (r *ItemCsellerAddRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemCsellerAddRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性(自定义的除外)也需要在props字段填写 . 规则：如果该SKU存在旧商品，则修改；否则新增Sku。如果更新时有对Sku进行操作，则Sku的properties一定要传入。如果存在自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText，其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号 */
func (r *ItemCsellerAddRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* Sku的数量串，结构如：num1,num2,num3 如：2,3 */
func (r *ItemCsellerAddRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 产品的规格信息。 */
func (r *ItemCsellerAddRequest) SetSkuSpecIds(value string) {
	r.SetValue("sku_spec_ids", value)
}

/* 新旧程度。可选值：new(新)，second(二手)，unused(闲置)。B商家不能发布二手商品。 
如果是二手商品，特定类目下属性里面必填新旧成色属性 */
func (r *ItemCsellerAddRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 1~4个淘宝图片空间的图片相对url列表，“,”分隔，做为商品的副图 */
func (r *ItemCsellerAddRequest) SetSubPicPaths(value string) {
	r.SetValue("sub_pic_paths", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持;0(默认)不更改 */
func (r *ItemCsellerAddRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题。不能超过60字符，受违禁词控制 */
func (r *ItemCsellerAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的 */
func (r *ItemCsellerAddRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 有效期。可选值:7,14;单位:天;默认值:14 */
func (r *ItemCsellerAddRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemCsellerAddRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}


func (r *ItemCsellerAddRequest) GetResponse(accessToken string) (*ItemCsellerAddResponse, []byte, error) {
	var resp ItemCsellerAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.cseller.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemCsellerAddResponse struct {
	Item *Item `json:"item"`
}

type ItemCsellerAddResponseResult struct {
	Response *ItemCsellerAddResponse `json:"item_cseller_add_response"`
}





/* 删除单条商品 */
type ItemDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，该参数必须 */
func (r *ItemDeleteRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemDeleteRequest) GetResponse(accessToken string) (*ItemDeleteResponse, []byte, error) {
	var resp ItemDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemDeleteResponse struct {
	Item *Item `json:"item"`
}

type ItemDeleteResponseResult struct {
	Response *ItemDeleteResponse `json:"item_delete_response"`
}





/* 调用该接口 会创建一个二级类目为网络文学原创电子书商品 */
type ItemEbookSerialAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 作者。长度不能超过60个字符 */
func (r *ItemEbookSerialAddRequest) SetAuthor(value string) {
	r.SetValue("author", value)
}

/* 叶子类目id */
func (r *ItemEbookSerialAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 版权到期时间，如2013-08-06 */
func (r *ItemEbookSerialAddRequest) SetCopyrightEnd(value string) {
	r.SetValue("copyright_end", value)
}

/* 版权文件。不得小于350*500；类型:jpg,png；大小不能超过2M */
func (r *ItemEbookSerialAddRequest) SetCopyrightFiles(value string) {
	r.SetValue("copyright_files", value)
}

/* 商品主图片。类型:JPG,PNG;最大:2M */
func (r *ItemEbookSerialAddRequest) SetCover(value string) {
	r.SetValue("cover", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
func (r *ItemEbookSerialAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 书名。长度不能超过60个字符 */
func (r *ItemEbookSerialAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商品外部编码，该字段的最大长度是512个字节 */
func (r *ItemEbookSerialAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 不能为0；如：0.50元/章 或者 0.50元/千字；取值范围:0.01-9999.99;精确到2位小数;单位:元。如:5.07，表示:5元7分. */
func (r *ItemEbookSerialAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 相关链接。不超过128个字符 */
func (r *ItemEbookSerialAddRequest) SetRelationLink(value string) {
	r.SetValue("relation_link", value)
}

/* 售卖方式。目前取值范围0、1； 
0：按章节售卖 1：按千字售卖 */
func (r *ItemEbookSerialAddRequest) SetSellWay(value string) {
	r.SetValue("sell_way", value)
}

/* 宝贝标题。不能超过60字符，受违禁词控制 */
func (r *ItemEbookSerialAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *ItemEbookSerialAddRequest) GetResponse(accessToken string) (*ItemEbookSerialAddResponse, []byte, error) {
	var resp ItemEbookSerialAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.ebook.serial.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemEbookSerialAddResponse struct {
	Item *Item `json:"item"`
}

type ItemEbookSerialAddResponseResult struct {
	Response *ItemEbookSerialAddResponse `json:"item_ebook_serial_add_response"`
}





/* 更新连载电子书信息 */
type ItemEbookSerialUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 作者。长度不能超过60个字符 */
func (r *ItemEbookSerialUpdateRequest) SetAuthor(value string) {
	r.SetValue("author", value)
}

/* 叶子类目id */
func (r *ItemEbookSerialUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 商品主图片。类型:JPG,PNG;最大:2M */
func (r *ItemEbookSerialUpdateRequest) SetCover(value string) {
	r.SetValue("cover", value)
}

/* 宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制 */
func (r *ItemEbookSerialUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 宝贝数字id */
func (r *ItemEbookSerialUpdateRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 书名。长度不能超过60个字符 */
func (r *ItemEbookSerialUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商品外部编码，该字段的最大长度是512个字节 */
func (r *ItemEbookSerialUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 不能为0；如：0.50元/章 或者 0.50元/千字；取值范围:0.01-9999.99;精确到2位小数;单位:元。如:5.07，表示:5元7分. */
func (r *ItemEbookSerialUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 相关链接。不超过128个字符 */
func (r *ItemEbookSerialUpdateRequest) SetRelationLink(value string) {
	r.SetValue("relation_link", value)
}

/* 宝贝标题。不能超过60字符，受违禁词控制 */
func (r *ItemEbookSerialUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *ItemEbookSerialUpdateRequest) GetResponse(accessToken string) (*ItemEbookSerialUpdateResponse, []byte, error) {
	var resp ItemEbookSerialUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.ebook.serial.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemEbookSerialUpdateResponse struct {
	Item *Item `json:"item"`
}

type ItemEbookSerialUpdateResponseResult struct {
	Response *ItemEbookSerialUpdateResponse `json:"item_ebook_serial_update_response"`
}





/* 获取单个商品的详细信息  
卖家未登录时只能获得这个商品的公开数据，卖家登录后可以获取商品的所有数据 */
type ItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要返回的商品对象字段，如title,price,desc_modules等。可选值：Item商品结构体中所有字段均可返回；多个字段用“,”分隔。<br>新增返回字段：item_weight(商品的重量，格式为数字，包含小数)、item_size(商品的体积，格式为数字，包含小数)、change_prop（商品基础色数据） */
func (r *ItemGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品数字ID */
func (r *ItemGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 商品数字ID(带有跟踪效果) */
func (r *ItemGetRequest) SetTrackIid(value string) {
	r.SetValue("track_iid", value)
}


func (r *ItemGetRequest) GetResponse(accessToken string) (*ItemGetResponse, []byte, error) {
	var resp ItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemGetResponse struct {
	Item *Item `json:"item"`
}

type ItemGetResponseResult struct {
	Response *ItemGetResponse `json:"item_get_response"`
}





/* 删除itemimg_id 所指定的商品图片  
传入的num_iid所对应的商品必须属于当前会话的用户  
itemimg_id对应的图片需要属于num_iid对应的商品 */
type ItemImgDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品图片ID */
func (r *ItemImgDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 商品数字ID，必选 */
func (r *ItemImgDeleteRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemImgDeleteRequest) GetResponse(accessToken string) (*ItemImgDeleteResponse, []byte, error) {
	var resp ItemImgDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.img.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemImgDeleteResponse struct {
	ItemImg *ItemImg `json:"item_img"`
}

type ItemImgDeleteResponseResult struct {
	Response *ItemImgDeleteResponse `json:"item_img_delete_response"`
}





/* 添加一张商品图片到num_iid指定的商品中  
传入的num_iid所对应的商品必须属于当前会话的用户  
如果更新图片需要设置itemimg_id，且该itemimg_id的图片记录需要属于传入的num_iid对应的商品。如果新增图片则不用设置  
商品图片有数量和大小上的限制，根据卖家享有的服务（如：卖家订购了多图服务等），商品图片数量限制不同。 */
type ItemImgUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品图片id(如果是更新图片，则需要传该参数) */
func (r *ItemImgUploadRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 商品图片内容类型:JPG,GIF;最大:500KB 。支持的文件类型：gif,jpg,jpeg,png */
func (r *ItemImgUploadRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 是否将该图片设为主图,可选值:true,false;默认值:false(非主图) */
func (r *ItemImgUploadRequest) SetIsMajor(value string) {
	r.SetValue("is_major", value)
}

/* 商品数字ID，该参数必须 */
func (r *ItemImgUploadRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 图片序号 */
func (r *ItemImgUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}


func (r *ItemImgUploadRequest) GetResponse(accessToken string) (*ItemImgUploadResponse, []byte, error) {
	var resp ItemImgUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.img.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemImgUploadResponse struct {
	ItemImg *ItemImg `json:"item_img"`
}

type ItemImgUploadResponseResult struct {
	Response *ItemImgUploadResponse `json:"item_img_upload_response"`
}





/* * 关联一张商品图片到num_iid指定的商品中 
    * 传入的num_iid所对应的商品必须属于当前会话的用户 
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行 
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额 */
type ItemJointImgRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品图片id(如果是更新图片，则需要传该参数) */
func (r *ItemJointImgRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 上传的图片是否关联为商品主图 */
func (r *ItemJointImgRequest) SetIsMajor(value string) {
	r.SetValue("is_major", value)
}

/* 商品数字ID，必选 */
func (r *ItemJointImgRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 图片URL,图片空间图片的相对地址 */
func (r *ItemJointImgRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 图片序号 */
func (r *ItemJointImgRequest) SetPosition(value string) {
	r.SetValue("position", value)
}


func (r *ItemJointImgRequest) GetResponse(accessToken string) (*ItemJointImgResponse, []byte, error) {
	var resp ItemJointImgResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.joint.img", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemJointImgResponse struct {
	ItemImg *ItemImg `json:"item_img"`
}

type ItemJointImgResponseResult struct {
	Response *ItemJointImgResponse `json:"item_joint_img_response"`
}





/* * 关联一张商品属性图片到num_iid指定的商品中 
    * 传入的num_iid所对应的商品必须属于当前会话的用户 
    * 图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的 
    * 商品图片关联在卖家身份和图片来源上的限制，卖家要是B卖家或订购了多图服务才能关联图片，并且图片要来自于卖家自己的图片空间才行 
    * 商品图片数量有限制。不管是上传的图片还是关联的图片，他们的总数不能超过一定限额，最多不能超过24张（每个颜色属性都有一张） */
type ItemJointPropimgRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性图片ID。如果是新增不需要填写 */
func (r *ItemJointPropimgRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 商品数字ID，必选 */
func (r *ItemJointPropimgRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded ) */
func (r *ItemJointPropimgRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 图片序号 */
func (r *ItemJointPropimgRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。 */
func (r *ItemJointPropimgRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *ItemJointPropimgRequest) GetResponse(accessToken string) (*ItemJointPropimgResponse, []byte, error) {
	var resp ItemJointPropimgResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.joint.propimg", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemJointPropimgResponse struct {
	PropImg *PropImg `json:"prop_img"`
}

type ItemJointPropimgResponseResult struct {
	Response *ItemJointPropimgResponse `json:"item_joint_propimg_response"`
}





/* 更新商品价格 */
type ItemPriceUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 售后服务说明模板id */
func (r *ItemPriceUpdateRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
func (r *ItemPriceUpdateRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (r *ItemPriceUpdateRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：  
no_mark(不做类型标记)  
time_card(点卡软件代充)  
fee_card(话费软件代充) */
func (r *ItemPriceUpdateRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 叶子类目id */
func (r *ItemPriceUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 货到付款运费模板ID */
func (r *ItemPriceUpdateRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制 */
func (r *ItemPriceUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemPriceUpdateRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemPriceUpdateRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
func (r *ItemPriceUpdateRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 支持会员打折。可选值:true,false; */
func (r *ItemPriceUpdateRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true) */
func (r *ItemPriceUpdateRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false; */
func (r *ItemPriceUpdateRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false; */
func (r *ItemPriceUpdateRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 商品图片。类型:JPG,GIF;最大长度:500k */
func (r *ItemPriceUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 加价幅度 如果为0，代表系统代理幅度 */
func (r *ItemPriceUpdateRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemPriceUpdateRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
func (r *ItemPriceUpdateRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D */
func (r *ItemPriceUpdateRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外店显示 */
func (r *ItemPriceUpdateRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
func (r *ItemPriceUpdateRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否替换sku */
func (r *ItemPriceUpdateRequest) SetIsReplaceSku(value string) {
	r.SetValue("is_replace_sku", value)
}

/* 是否在淘宝上显示 */
func (r *ItemPriceUpdateRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
func (r *ItemPriceUpdateRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
func (r *ItemPriceUpdateRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
func (r *ItemPriceUpdateRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (r *ItemPriceUpdateRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (r *ItemPriceUpdateRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和 */
func (r *ItemPriceUpdateRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商品数字ID，该参数必须 */
func (r *ItemPriceUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 商家编码 */
func (r *ItemPriceUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *ItemPriceUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
func (r *ItemPriceUpdateRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板） */
func (r *ItemPriceUpdateRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 */
func (r *ItemPriceUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (r *ItemPriceUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节 */
func (r *ItemPriceUpdateRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
func (r *ItemPriceUpdateRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemPriceUpdateRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemPriceUpdateRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (r *ItemPriceUpdateRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemPriceUpdateRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。 */
func (r *ItemPriceUpdateRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
func (r *ItemPriceUpdateRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
func (r *ItemPriceUpdateRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 */
func (r *ItemPriceUpdateRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题. 不能超过60字符,受违禁词控制 */
func (r *ItemPriceUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 有效期。可选值:7,14;单位:天; */
func (r *ItemPriceUpdateRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemPriceUpdateRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}


func (r *ItemPriceUpdateRequest) GetResponse(accessToken string) (*ItemPriceUpdateResponse, []byte, error) {
	var resp ItemPriceUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.price.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemPriceUpdateResponse struct {
	Item *Item `json:"item"`
}

type ItemPriceUpdateResponseResult struct {
	Response *ItemPriceUpdateResponse `json:"item_price_update_response"`
}





/* 删除propimg_id 所指定的商品属性图片  
传入的num_iid所对应的商品必须属于当前会话的用户  
propimg_id对应的属性图片需要属于num_iid对应的商品 */
type ItemPropimgDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品属性图片ID */
func (r *ItemPropimgDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 商品数字ID，必选 */
func (r *ItemPropimgDeleteRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemPropimgDeleteRequest) GetResponse(accessToken string) (*ItemPropimgDeleteResponse, []byte, error) {
	var resp ItemPropimgDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.propimg.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemPropimgDeleteResponse struct {
	PropImg *PropImg `json:"prop_img"`
}

type ItemPropimgDeleteResponseResult struct {
	Response *ItemPropimgDeleteResponse `json:"item_propimg_delete_response"`
}





/* 添加一张商品属性图片到num_iid指定的商品中  
传入的num_iid所对应的商品必须属于当前会话的用户  
图片的属性必须要是颜色的属性，这个在前台显示的时候需要和sku进行关联的  
商品属性图片只有享有服务的卖家（如：淘宝大卖家、订购了淘宝多图服务的卖家）才能上传  
商品属性图片有数量和大小上的限制，最多不能超过24张（每个颜色属性都有一张）。 */
type ItemPropimgUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性图片ID。如果是新增不需要填写 */
func (r *ItemPropimgUploadRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 属性图片内容。类型:JPG,GIF;最大长度:500K;图片大小不超过:1M */
func (r *ItemPropimgUploadRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 商品数字ID，必选 */
func (r *ItemPropimgUploadRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 图片位置 */
func (r *ItemPropimgUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。 */
func (r *ItemPropimgUploadRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *ItemPropimgUploadRequest) GetResponse(accessToken string) (*ItemPropimgUploadResponse, []byte, error) {
	var resp ItemPropimgUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.propimg.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemPropimgUploadResponse struct {
	PropImg *PropImg `json:"prop_img"`
}

type ItemPropimgUploadResponseResult struct {
	Response *ItemPropimgUploadResponse `json:"item_propimg_upload_response"`
}





/* 提供按照全量或增量形式修改宝贝/SKU库存的功能 */
type ItemQuantityUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，必填参数 */
func (r *ItemQuantityUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* SKU的商家编码，可选参数。如果不填则默认修改宝贝的库存，如果填了则按照商家编码搜索出对应的SKU并修改库存。当sku_id和本字段都填写时以sku_id为准搜索对应SKU */
func (r *ItemQuantityUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 库存修改值，必选。当全量更新库存时，quantity必须为大于等于0的正整数；当增量更新库存时，quantity为整数，可小于等于0。若增量更新时传入的库存为负数，则负数与实际库存之和不能小于0。比如当前实际库存为1，传入增量更新quantity=-1，库存改为0 */
func (r *ItemQuantityUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 要操作的SKU的数字ID，可选。如果不填默认修改宝贝的库存，如果填上则修改该SKU的库存 */
func (r *ItemQuantityUpdateRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}

/* 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新 */
func (r *ItemQuantityUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *ItemQuantityUpdateRequest) GetResponse(accessToken string) (*ItemQuantityUpdateResponse, []byte, error) {
	var resp ItemQuantityUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.quantity.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemQuantityUpdateResponse struct {
	Item *Item `json:"item"`
}

type ItemQuantityUpdateResponseResult struct {
	Response *ItemQuantityUpdateResponse `json:"item_quantity_update_response"`
}





/* 将当前用户指定商品设置为橱窗推荐状态 
橱窗推荐需要用户有剩余橱窗位才可以顺利执行 
这个Item所属卖家从传入的session中获取，需要session绑定 
需要判断橱窗推荐是否已满，橱窗推荐已满停止调用橱窗推荐接口，2010年1月底开放查询剩余橱窗推荐数后可以按数量橱窗推荐商品 */
type ItemRecommendAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，该参数必须 */
func (r *ItemRecommendAddRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemRecommendAddRequest) GetResponse(accessToken string) (*ItemRecommendAddResponse, []byte, error) {
	var resp ItemRecommendAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.recommend.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemRecommendAddResponse struct {
	Item *Item `json:"item"`
}

type ItemRecommendAddResponseResult struct {
	Response *ItemRecommendAddResponse `json:"item_recommend_add_response"`
}





/* 取消当前用户指定商品的橱窗推荐状态 
这个Item所属卖家从传入的session中获取，需要session绑定 */
type ItemRecommendDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，该参数必须 */
func (r *ItemRecommendDeleteRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemRecommendDeleteRequest) GetResponse(accessToken string) (*ItemRecommendDeleteResponse, []byte, error) {
	var resp ItemRecommendDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.recommend.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemRecommendDeleteResponse struct {
	Item *Item `json:"item"`
}

type ItemRecommendDeleteResponseResult struct {
	Response *ItemRecommendDeleteResponse `json:"item_recommend_delete_response"`
}





/* 新增一个sku到num_iid指定的商品中  
传入的iid所对应的商品必须属于当前会话的用户 */
type ItemSkuAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功 */
func (r *ItemSkuAddRequest) SetItemPrice(value string) {
	r.SetValue("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (r *ItemSkuAddRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取。必选 */
func (r *ItemSkuAddRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* Sku的商家外部id */
func (r *ItemSkuAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemSkuAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。 
如果包含自定义属性则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的‘$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号， */
func (r *ItemSkuAddRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数 */
func (r *ItemSkuAddRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 产品的规格信息 */
func (r *ItemSkuAddRequest) SetSpecId(value string) {
	r.SetValue("spec_id", value)
}


func (r *ItemSkuAddRequest) GetResponse(accessToken string) (*ItemSkuAddResponse, []byte, error) {
	var resp ItemSkuAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.sku.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkuAddResponse struct {
	Sku *Sku `json:"sku"`
}

type ItemSkuAddResponseResult struct {
	Response *ItemSkuAddResponse `json:"item_sku_add_response"`
}





/* 删除一个sku的数据 
需要删除的sku通过属性properties进行匹配查找 */
type ItemSkuDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败 */
func (r *ItemSkuDeleteRequest) SetItemNum(value string) {
	r.SetValue("item_num", value)
}

/* sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功 */
func (r *ItemSkuDeleteRequest) SetItemPrice(value string) {
	r.SetValue("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (r *ItemSkuDeleteRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取。必选 */
func (r *ItemSkuDeleteRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
func (r *ItemSkuDeleteRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *ItemSkuDeleteRequest) GetResponse(accessToken string) (*ItemSkuDeleteResponse, []byte, error) {
	var resp ItemSkuDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.sku.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkuDeleteResponse struct {
	Sku *Sku `json:"sku"`
}

type ItemSkuDeleteResponseResult struct {
	Response *ItemSkuDeleteResponse `json:"item_sku_delete_response"`
}





/* 获取sku_id所对应的sku数据  
sku_id对应的sku要属于传入的nick对应的卖家 */
type ItemSkuGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
func (r *ItemSkuGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 卖家nick(num_iid和nick必传一个)，只传卖家nick时候，该api返回的结果不包含cspu（SKu上的产品规格信息）。 */
func (r *ItemSkuGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 商品的数字IID（num_iid和nick必传一个，推荐用num_iid），传商品的数字id返回的结果里包含cspu（SKu上的产品规格信息）。 */
func (r *ItemSkuGetRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* Sku的id。可以通过taobao.item.get得到 */
func (r *ItemSkuGetRequest) SetSkuId(value string) {
	r.SetValue("sku_id", value)
}


func (r *ItemSkuGetRequest) GetResponse(accessToken string) (*ItemSkuGetResponse, []byte, error) {
	var resp ItemSkuGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.sku.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkuGetResponse struct {
	Sku *Sku `json:"sku"`
}

type ItemSkuGetResponseResult struct {
	Response *ItemSkuGetResponse `json:"item_sku_get_response"`
}





/* 更新商品SKU的价格 */
type ItemSkuPriceUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
func (r *ItemSkuPriceUpdateRequest) SetItemPrice(value string) {
	r.SetValue("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (r *ItemSkuPriceUpdateRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取 */
func (r *ItemSkuPriceUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* Sku的商家外部id */
func (r *ItemSkuPriceUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
func (r *ItemSkuPriceUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充 */
func (r *ItemSkuPriceUpdateRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数 */
func (r *ItemSkuPriceUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}


func (r *ItemSkuPriceUpdateRequest) GetResponse(accessToken string) (*ItemSkuPriceUpdateResponse, []byte, error) {
	var resp ItemSkuPriceUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.sku.price.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkuPriceUpdateResponse struct {
	Sku *Sku `json:"sku"`
}

type ItemSkuPriceUpdateResponseResult struct {
	Response *ItemSkuPriceUpdateResponse `json:"item_sku_price_update_response"`
}





/* *更新一个sku的数据  
*需要更新的sku通过属性properties进行匹配查找  
*商品的数量和价格必须大于等于0  
*sku记录会更新到指定的num_iid对应的商品中  
*num_iid对应的商品必须属于当前的会话用户 */
type ItemSkuUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功 */
func (r *ItemSkuUpdateRequest) SetItemPrice(value string) {
	r.SetValue("item_price", value)
}

/* Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN */
func (r *ItemSkuUpdateRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* Sku所属商品数字id，可通过 taobao.item.get 获取 */
func (r *ItemSkuUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* Sku的商家外部id */
func (r *ItemSkuUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中） */
func (r *ItemSkuUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。 
如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号， */
func (r *ItemSkuUpdateRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}

/* Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数 */
func (r *ItemSkuUpdateRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 产品的规格信息。 */
func (r *ItemSkuUpdateRequest) SetSpecId(value string) {
	r.SetValue("spec_id", value)
}


func (r *ItemSkuUpdateRequest) GetResponse(accessToken string) (*ItemSkuUpdateResponse, []byte, error) {
	var resp ItemSkuUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.sku.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkuUpdateResponse struct {
	Sku *Sku `json:"sku"`
}

type ItemSkuUpdateResponseResult struct {
	Response *ItemSkuUpdateResponse `json:"item_sku_update_response"`
}





/* * 获取多个商品下的所有sku */
type ItemSkusGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”分隔。 */
func (r *ItemSkusGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* sku所属商品数字id，必选。num_iid个数不能超过40个 */
func (r *ItemSkusGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}


func (r *ItemSkusGetRequest) GetResponse(accessToken string) (*ItemSkusGetResponse, []byte, error) {
	var resp ItemSkusGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.skus.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemSkusGetResponse struct {
	Skus *SkuListObject `json:"skus"`
}

type ItemSkusGetResponseResult struct {
	Response *ItemSkusGetResponse `json:"item_skus_get_response"`
}





/* 查询当前登录用户的店铺的宝贝详情页的模板名称 */
type ItemTemplatesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *ItemTemplatesGetRequest) GetResponse(accessToken string) (*ItemTemplatesGetResponse, []byte, error) {
	var resp ItemTemplatesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.templates.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemTemplatesGetResponse struct {
	ItemTemplateList *ItemTemplateListObject `json:"item_template_list"`
}

type ItemTemplatesGetResponseResult struct {
	Response *ItemTemplatesGetResponse `json:"item_templates_get_response"`
}





/* 根据传入的num_iid更新对应的商品的数据  
传入的num_iid所对应的商品必须属于当前会话的用户  
商品的属性和sku的属性有包含的关系，商品的价格要位于sku的价格区间之中（例如，sku价格有5元、10元两种，那么商品的价格就需要大于等于5元，小于等于10元，否则更新商品会失败）  
商品的类目和商品的价格、sku的价格都有一定的相关性（具体的关系要通过类目属性查询接口获得）  
当关键属性值更新为“其他”的时候，需要输入input_pids和input_str商品才能更新成功。该接口不支持产品属性修改。 */
type ItemUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 售后服务说明模板id */
func (r *ItemUpdateRequest) SetAfterSaleId(value string) {
	r.SetValue("after_sale_id", value)
}

/* 商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true */
func (r *ItemUpdateRequest) SetApproveStatus(value string) {
	r.SetValue("approve_status", value)
}

/* 商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50% */
func (r *ItemUpdateRequest) SetAuctionPoint(value string) {
	r.SetValue("auction_point", value)
}

/* 代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型：  
no_mark(不做类型标记)  
time_card(点卡软件代充)  
fee_card(话费软件代充) */
func (r *ItemUpdateRequest) SetAutoFill(value string) {
	r.SetValue("auto_fill", value)
}

/* 商品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1 */
func (r *ItemUpdateRequest) SetChangeProp(value string) {
	r.SetValue("change_prop", value)
}

/* 叶子类目id */
func (r *ItemUpdateRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 货到付款运费模板ID */
func (r *ItemUpdateRequest) SetCodPostageId(value string) {
	r.SetValue("cod_postage_id", value)
}

/* 商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制 */
func (r *ItemUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 商品描述模块化，模块列表；数据结构可参考Item_Desc_Module 。详细使用说明：http://open.taobao.com/support/question_detail.htm?spm=0.0.0.0.cRcj3S&id=147498 ； */
func (r *ItemUpdateRequest) SetDescModules(value string) {
	r.SetValue("desc_modules", value)
}

/* 支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock） */
func (r *ItemUpdateRequest) SetEmptyFields(value string) {
	r.SetValue("empty_fields", value)
}

/* ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分 */
func (r *ItemUpdateRequest) SetEmsFee(value string) {
	r.SetValue("ems_fee", value)
}

/* 快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分 */
func (r *ItemUpdateRequest) SetExpressFee(value string) {
	r.SetValue("express_fee", value)
}

/* 厂家联系方式 */
func (r *ItemUpdateRequest) SetFoodSecurityContact(value string) {
	r.SetValue("food_security.contact", value)
}

/* 产品标准号 */
func (r *ItemUpdateRequest) SetFoodSecurityDesignCode(value string) {
	r.SetValue("food_security.design_code", value)
}

/* 厂名 */
func (r *ItemUpdateRequest) SetFoodSecurityFactory(value string) {
	r.SetValue("food_security.factory", value)
}

/* 厂址 */
func (r *ItemUpdateRequest) SetFoodSecurityFactorySite(value string) {
	r.SetValue("food_security.factory_site", value)
}

/* 食品添加剂 */
func (r *ItemUpdateRequest) SetFoodSecurityFoodAdditive(value string) {
	r.SetValue("food_security.food_additive", value)
}

/* 健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题； */
func (r *ItemUpdateRequest) SetFoodSecurityHealthProductNo(value string) {
	r.SetValue("food_security.health_product_no", value)
}

/* 配料表 */
func (r *ItemUpdateRequest) SetFoodSecurityMix(value string) {
	r.SetValue("food_security.mix", value)
}

/* 保质期 */
func (r *ItemUpdateRequest) SetFoodSecurityPeriod(value string) {
	r.SetValue("food_security.period", value)
}

/* 储藏方法 */
func (r *ItemUpdateRequest) SetFoodSecurityPlanStorage(value string) {
	r.SetValue("food_security.plan_storage", value)
}

/* 生产许可证号 */
func (r *ItemUpdateRequest) SetFoodSecurityPrdLicenseNo(value string) {
	r.SetValue("food_security.prd_license_no", value)
}

/* 生产结束日期,格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityProductDateEnd(value string) {
	r.SetValue("food_security.product_date_end", value)
}

/* 生产开始日期，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityProductDateStart(value string) {
	r.SetValue("food_security.product_date_start", value)
}

/* 进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityStockDateEnd(value string) {
	r.SetValue("food_security.stock_date_end", value)
}

/* 进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd */
func (r *ItemUpdateRequest) SetFoodSecurityStockDateStart(value string) {
	r.SetValue("food_security.stock_date_start", value)
}

/* 供货商 */
func (r *ItemUpdateRequest) SetFoodSecuritySupplier(value string) {
	r.SetValue("food_security.supplier", value)
}

/* 运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担); */
func (r *ItemUpdateRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他） */
func (r *ItemUpdateRequest) SetGlobalStockCountry(value string) {
	r.SetValue("global_stock_country", value)
}

/* 全球购商品采购地（库存类型） 
全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用 */
func (r *ItemUpdateRequest) SetGlobalStockType(value string) {
	r.SetValue("global_stock_type", value)
}

/* 支持会员打折。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否有发票。可选值:true,false (商城卖家此字段必须为true) */
func (r *ItemUpdateRequest) SetHasInvoice(value string) {
	r.SetValue("has_invoice", value)
}

/* 橱窗推荐。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 是否有保修。可选值:true,false; */
func (r *ItemUpdateRequest) SetHasWarranty(value string) {
	r.SetValue("has_warranty", value)
}

/* 商品图片。类型:JPG,GIF;最大长度:500k */
func (r *ItemUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。 */
func (r *ItemUpdateRequest) SetIncrement(value string) {
	r.SetValue("increment", value)
}

/* 用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。 */
func (r *ItemUpdateRequest) SetInputPids(value string) {
	r.SetValue("input_pids", value)
}

/* 用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。 */
func (r *ItemUpdateRequest) SetInputStr(value string) {
	r.SetValue("input_str", value)
}

/* 是否是3D */
func (r *ItemUpdateRequest) SetIs3D(value string) {
	r.SetValue("is_3D", value)
}

/* 是否在外店显示 */
func (r *ItemUpdateRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记 */
func (r *ItemUpdateRequest) SetIsLightningConsignment(value string) {
	r.SetValue("is_lightning_consignment", value)
}

/* 是否替换sku */
func (r *ItemUpdateRequest) SetIsReplaceSku(value string) {
	r.SetValue("is_replace_sku", value)
}

/* 是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品） */
func (r *ItemUpdateRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。 */
func (r *ItemUpdateRequest) SetIsXinpin(value string) {
	r.SetValue("is_xinpin", value)
}

/* 表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。 
该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。 
在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m) 
在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0 */
func (r *ItemUpdateRequest) SetItemSize(value string) {
	r.SetValue("item_size", value)
}

/* 商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0 */
func (r *ItemUpdateRequest) SetItemWeight(value string) {
	r.SetValue("item_weight", value)
}

/* 商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN” */
func (r *ItemUpdateRequest) SetLang(value string) {
	r.SetValue("lang", value)
}

/* 上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。 */
func (r *ItemUpdateRequest) SetListTime(value string) {
	r.SetValue("list_time", value)
}

/* 编辑电子凭证宝贝时候表示是否使用邮寄 
0: 代表不使用邮寄； 
1：代表使用邮寄； 
如果不设置这个值，代表不使用邮寄 */
func (r *ItemUpdateRequest) SetLocalityLifeChooseLogis(value string) {
	r.SetValue("locality_life.choose_logis", value)
}

/* 本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期; 
如果有效期为起止日期类型，此值为2012-08-06,2012-08-16 
如果有效期为【购买成功日 至】类型则格式为2012-08-16 
如果有效期为天数类型则格式为15 */
func (r *ItemUpdateRequest) SetLocalityLifeExpirydate(value string) {
	r.SetValue("locality_life.expirydate", value)
}

/* 码商信息，格式为 码商id:nick */
func (r *ItemUpdateRequest) SetLocalityLifeMerchant(value string) {
	r.SetValue("locality_life.merchant", value)
}

/* 网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID */
func (r *ItemUpdateRequest) SetLocalityLifeNetworkId(value string) {
	r.SetValue("locality_life.network_id", value)
}

/* 电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数 */
func (r *ItemUpdateRequest) SetLocalityLifeOnsaleAutoRefundRatio(value string) {
	r.SetValue("locality_life.onsale_auto_refund_ratio", value)
}

/* 退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例 */
func (r *ItemUpdateRequest) SetLocalityLifeRefundRatio(value string) {
	r.SetValue("locality_life.refund_ratio", value)
}

/* 退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担) */
func (r *ItemUpdateRequest) SetLocalityLifeRefundmafee(value string) {
	r.SetValue("locality_life.refundmafee", value)
}

/* 核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款 */
func (r *ItemUpdateRequest) SetLocalityLifeVerification(value string) {
	r.SetValue("locality_life.verification", value)
}

/* 所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (r *ItemUpdateRequest) SetLocationCity(value string) {
	r.SetValue("location.city", value)
}

/* 所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到 */
func (r *ItemUpdateRequest) SetLocationState(value string) {
	r.SetValue("location.state", value)
}

/* 商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。 */
func (r *ItemUpdateRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商品数字ID，该参数必须 */
func (r *ItemUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 商家编码 */
func (r *ItemUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。 
对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。 */
func (r *ItemUpdateRequest) SetPaimaiInfoDeposit(value string) {
	r.SetValue("paimai_info.deposit", value)
}

/* 降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。 */
func (r *ItemUpdateRequest) SetPaimaiInfoInterval(value string) {
	r.SetValue("paimai_info.interval", value)
}

/* 拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。 */
func (r *ItemUpdateRequest) SetPaimaiInfoMode(value string) {
	r.SetValue("paimai_info.mode", value)
}

/* 降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数 */
func (r *ItemUpdateRequest) SetPaimaiInfoReserve(value string) {
	r.SetValue("paimai_info.reserve", value)
}

/* 自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
func (r *ItemUpdateRequest) SetPaimaiInfoValidHour(value string) {
	r.SetValue("paimai_info.valid_hour", value)
}

/* 自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。 */
func (r *ItemUpdateRequest) SetPaimaiInfoValidMinute(value string) {
	r.SetValue("paimai_info.valid_minute", value)
}

/* 商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path */
func (r *ItemUpdateRequest) SetPicPath(value string) {
	r.SetValue("pic_path", value)
}

/* 平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写 */
func (r *ItemUpdateRequest) SetPostFee(value string) {
	r.SetValue("post_fee", value)
}

/* 宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板） */
func (r *ItemUpdateRequest) SetPostageId(value string) {
	r.SetValue("postage_id", value)
}

/* 商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。 */
func (r *ItemUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 商品所属的产品ID(B商家发布商品需要用) */
func (r *ItemUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节 */
func (r *ItemUpdateRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。 */
func (r *ItemUpdateRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100 */
func (r *ItemUpdateRequest) SetScenicTicketBookCost(value string) {
	r.SetValue("scenic_ticket_book_cost", value)
}

/* 景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付 */
func (r *ItemUpdateRequest) SetScenicTicketPayWay(value string) {
	r.SetValue("scenic_ticket_pay_way", value)
}

/* 商品卖点信息，最长150个字符。仅天猫商家可用。 */
func (r *ItemUpdateRequest) SetSellPoint(value string) {
	r.SetValue("sell_point", value)
}

/* 是否承诺退换货服务!虚拟商品无须设置此项! */
func (r *ItemUpdateRequest) SetSellPromise(value string) {
	r.SetValue("sell_promise", value)
}

/* 重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。 */
func (r *ItemUpdateRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节 */
func (r *ItemUpdateRequest) SetSkuOuterIds(value string) {
	r.SetValue("sku_outer_ids", value)
}

/* 更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分 */
func (r *ItemUpdateRequest) SetSkuPrices(value string) {
	r.SetValue("sku_prices", value)
}

/* 更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid,多个sku之间用逗号分隔。该字段内的销售属性(自定义的除外)也需要在props字段填写 . 规则：如果该SKU存在旧商品，则修改；否则新增Sku。如果更新时有对Sku进行操作，则Sku的properties一定要传入。如果存在自定义销售属性，则格式为pid:vid;pid2:vid2;$pText:vText，其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号 */
func (r *ItemUpdateRequest) SetSkuProperties(value string) {
	r.SetValue("sku_properties", value)
}

/* 更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4 */
func (r *ItemUpdateRequest) SetSkuQuantities(value string) {
	r.SetValue("sku_quantities", value)
}

/* 暂时不可用 */
func (r *ItemUpdateRequest) SetSkuSpecIds(value string) {
	r.SetValue("sku_spec_ids", value)
}

/* 商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。 */
func (r *ItemUpdateRequest) SetStuffStatus(value string) {
	r.SetValue("stuff_status", value)
}

/* 商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改
集市卖家默认拍下减库存;
商城卖家默认付款减库存 */
func (r *ItemUpdateRequest) SetSubStock(value string) {
	r.SetValue("sub_stock", value)
}

/* 宝贝标题. 不能超过30字符,受违禁词控制 */
func (r *ItemUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 有效期。可选值:7,14;单位:天; */
func (r *ItemUpdateRequest) SetValidThru(value string) {
	r.SetValue("valid_thru", value)
}

/* 商品的重量(商超卖家专用字段) */
func (r *ItemUpdateRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}


func (r *ItemUpdateRequest) GetResponse(accessToken string) (*ItemUpdateResponse, []byte, error) {
	var resp ItemUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemUpdateResponse struct {
	Item *Item `json:"item"`
}

type ItemUpdateResponseResult struct {
	Response *ItemUpdateResponse `json:"item_update_response"`
}





/* * 单个商品下架 
    * 输入的num_iid必须属于当前会话用户 */
type ItemUpdateDelistingRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，该参数必须 */
func (r *ItemUpdateDelistingRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemUpdateDelistingRequest) GetResponse(accessToken string) (*ItemUpdateDelistingResponse, []byte, error) {
	var resp ItemUpdateDelistingResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.update.delisting", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemUpdateDelistingResponse struct {
	Item *Item `json:"item"`
}

type ItemUpdateDelistingResponseResult struct {
	Response *ItemUpdateDelistingResponse `json:"item_update_delisting_response"`
}





/* * 单个商品上架 
* 输入的num_iid必须属于当前会话用户 */
type ItemUpdateListingRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num */
func (r *ItemUpdateListingRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 商品数字ID，该参数必须 */
func (r *ItemUpdateListingRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}


func (r *ItemUpdateListingRequest) GetResponse(accessToken string) (*ItemUpdateListingResponse, []byte, error) {
	var resp ItemUpdateListingResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.item.update.listing", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemUpdateListingResponse struct {
	Item *Item `json:"item"`
}

type ItemUpdateListingResponseResult struct {
	Response *ItemUpdateListingResponse `json:"item_update_listing_response"`
}





/* 跟据卖家设定的商品外部id获取商品  
这个商品对应卖家从传入的session中获取，需要session绑定 */
type ItemsCustomGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值：Item商品结构体中的所有字段；多个字段之间用“,”分隔。如果想返回整个子对象，那字段为item_img，如果是想返回子对象里面的字段，那字段为item_img.url。新增返回字段：one_station标记商品是否淘1站商品 */
func (r *ItemsCustomGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品的外部商品ID，支持批量，最多不超过40个。 */
func (r *ItemsCustomGetRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}


func (r *ItemsCustomGetRequest) GetResponse(accessToken string) (*ItemsCustomGetResponse, []byte, error) {
	var resp ItemsCustomGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.items.custom.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemsCustomGetResponse struct {
	Items *ItemListObject `json:"items"`
}

type ItemsCustomGetResponseResult struct {
	Response *ItemsCustomGetResponse `json:"items_custom_get_response"`
}





/* 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤  
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取 */
type ItemsInventoryGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 分类字段。可选值:<br>
regular_shelved(定时上架)<br>
never_on_shelf(从未上架)<br>
off_shelf(我下架的)<br>
for_shelved(等待所有上架)<br>

sold_out(全部卖完)<br>
violation_off_shelf(违规下架的)<br>
默认查询for_shelved(等待所有上架)这个状态的商品<br>
<font color='red'>注：for_shelved(等待所有上架)=regular_shelved(定时上架)+never_on_shelf(从未上架)+off_shelf(我下架的)</font> */
func (r *ItemsInventoryGetRequest) SetBanner(value string) {
	r.SetValue("banner", value)
}

/* 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
func (r *ItemsInventoryGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 商品结束修改时间 */
func (r *ItemsInventoryGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 需返回的字段列表。可选值为 Item 商品结构体中的以下字段： 
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru, list_time,price,has_discount,has_invoice,has_warranty,has_showcase, modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。<br> 
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.get。 */
func (r *ItemsInventoryGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
func (r *ItemsInventoryGetRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 商品是否在外部网店显示 */
func (r *ItemsInventoryGetRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 商品是否在淘宝显示 */
func (r *ItemsInventoryGetRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
func (r *ItemsInventoryGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 页码。取值范围:大于零小于等于101的整数;默认值为1，即返回第一页数据。当页码超过101页时系统就会报错，故请大家在用此接口获取数据时尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品。 */
func (r *ItemsInventoryGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。 */
func (r *ItemsInventoryGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 搜索字段。搜索商品的title。 */
func (r *ItemsInventoryGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
func (r *ItemsInventoryGetRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* 商品起始修改时间 */
func (r *ItemsInventoryGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}


func (r *ItemsInventoryGetRequest) GetResponse(accessToken string) (*ItemsInventoryGetResponse, []byte, error) {
	var resp ItemsInventoryGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.items.inventory.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemsInventoryGetResponse struct {
	Items *ItemListObject `json:"items"`
	TotalResults int `json:"total_results"`
}

type ItemsInventoryGetResponseResult struct {
	Response *ItemsInventoryGetResponse `json:"items_inventory_get_response"`
}





/* 查看非公开属性时需要用户登录 */
type ItemsListGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要返回的商品对象字段。可选值：Item商品结构体中字段均可返回(除item_weight,item_size)；多个字段用“,”分隔。如果想返回整个子对象，那字段为itemimg，如果是想返回子对象里面的字段，那字段为itemimg.url。 */
func (r *ItemsListGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品数字id列表，多个num_iid用逗号隔开，一次不超过20个。 */
func (r *ItemsListGetRequest) SetNumIids(value string) {
	r.SetValue("num_iids", value)
}

/* 商品数字id列表，多个track_iid用逗号隔开，一次不超过20个。 */
func (r *ItemsListGetRequest) SetTrackIids(value string) {
	r.SetValue("track_iids", value)
}


func (r *ItemsListGetRequest) GetResponse(accessToken string) (*ItemsListGetResponse, []byte, error) {
	var resp ItemsListGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.items.list.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemsListGetResponse struct {
	Items *ItemListObject `json:"items"`
}

type ItemsListGetResponseResult struct {
	Response *ItemsListGetResponse `json:"items_list_get_response"`
}





/* 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤  
只能获得商品的部分信息，商品的详细信息请通过taobao.item.get获取 */
type ItemsOnsaleGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到 */
func (r *ItemsOnsaleGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 结束的修改时间 */
func (r *ItemsOnsaleGetRequest) SetEndModified(value string) {
	r.SetValue("end_modified", value)
}

/* 需返回的字段列表。可选值：Item商品结构体中的以下字段： 
approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。
不支持其他字段，如果需要获取其他字段数据，调用taobao.item.get。 */
func (r *ItemsOnsaleGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否参与会员折扣。可选值：true，false。默认不过滤该条件 */
func (r *ItemsOnsaleGetRequest) SetHasDiscount(value string) {
	r.SetValue("has_discount", value)
}

/* 是否橱窗推荐。 可选值：true，false。默认不过滤该条件 */
func (r *ItemsOnsaleGetRequest) SetHasShowcase(value string) {
	r.SetValue("has_showcase", value)
}

/* 商品是否在外部网店显示 */
func (r *ItemsOnsaleGetRequest) SetIsEx(value string) {
	r.SetValue("is_ex", value)
}

/* 商品是否在淘宝显示 */
func (r *ItemsOnsaleGetRequest) SetIsTaobao(value string) {
	r.SetValue("is_taobao", value)
}

/* 排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc */
func (r *ItemsOnsaleGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
func (r *ItemsOnsaleGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过2万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品 */
func (r *ItemsOnsaleGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 搜索字段。搜索商品的title。 */
func (r *ItemsOnsaleGetRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 卖家店铺内自定义类目ID。多个之间用“,”分隔。可以根据taobao.sellercats.list.get获得.(<font color="red">注：目前最多支持32个ID号传入</font>) */
func (r *ItemsOnsaleGetRequest) SetSellerCids(value string) {
	r.SetValue("seller_cids", value)
}

/* 起始的修改时间 */
func (r *ItemsOnsaleGetRequest) SetStartModified(value string) {
	r.SetValue("start_modified", value)
}


func (r *ItemsOnsaleGetRequest) GetResponse(accessToken string) (*ItemsOnsaleGetResponse, []byte, error) {
	var resp ItemsOnsaleGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.items.onsale.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemsOnsaleGetResponse struct {
	Items *ItemListObject `json:"items"`
	TotalResults int `json:"total_results"`
}

type ItemsOnsaleGetResponseResult struct {
	Response *ItemsOnsaleGetResponse `json:"items_onsale_get_response"`
}





/* 获取类目ID，必需是叶子类目ID；调用taobao.itemcats.get.v2获取  
传入关键属性,结构:pid:vid;pid:vid.调用taobao.itemprops.get.v2获取pid, 
调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props. */
type ProductAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 非关键属性结构:pid:vid;pid:vid.<br>
非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;
<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font> */
func (r *ProductAddRequest) SetBinds(value string) {
	r.SetValue("binds", value)
}

/* 商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id. */
func (r *ProductAddRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
<br><font color=red>注：包含所有自定义属性的传入</font> */
func (r *ProductAddRequest) SetCustomerProps(value string) {
	r.SetValue("customer_props", value)
}

/* 产品描述.最大不超过25000个字符 */
func (r *ProductAddRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 存放产品扩展信息，由List(ProductExtraInfo)转化成jsonArray存入. */
func (r *ProductAddRequest) SetExtraInfo(value string) {
	r.SetValue("extra_info", value)
}

/* 产品主图片.最大1M,目前仅支持GIF,JPG. */
func (r *ProductAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 是不是主图 */
func (r *ProductAddRequest) SetMajor(value string) {
	r.SetValue("major", value)
}

/* 市场ID，1为新增C2C市场的产品信息， 2为新增B2C市场的产品信息。 
不填写此值则C用户新增B2C市场的产品信息，B用户新增B2C市场的产品信息。 */
func (r *ProductAddRequest) SetMarketId(value string) {
	r.SetValue("market_id", value)
}

/* 上市时间。目前只支持鞋城类目传入此参数 */
func (r *ProductAddRequest) SetMarketTime(value string) {
	r.SetValue("market_time", value)
}

/* 产品名称,最大30个字符. */
func (r *ProductAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 外部产品ID */
func (r *ProductAddRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 包装清单。注意，在管控类目下，包装清单不能为空，同时保证清单的格式为： 
名称:数字;名称:数字; 
其中，名称不能违禁、不能超过60字符，数字不能超过999 */
func (r *ProductAddRequest) SetPackingList(value string) {
	r.SetValue("packing_list", value)
}

/* 产品市场价.精确到2位小数;单位为元.如：200.07 */
func (r *ProductAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸 */
func (r *ProductAddRequest) SetPropertyAlias(value string) {
	r.SetValue("property_alias", value)
}

/* 关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props. */
func (r *ProductAddRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid. */
func (r *ProductAddRequest) SetSaleProps(value string) {
	r.SetValue("sale_props", value)
}

/* 商品卖点描述，长度限制为20个汉字 */
func (r *ProductAddRequest) SetSellPt(value string) {
	r.SetValue("sell_pt", value)
}


func (r *ProductAddRequest) GetResponse(accessToken string) (*ProductAddResponse, []byte, error) {
	var resp ProductAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductAddResponse struct {
	Product *Product `json:"product"`
}

type ProductAddResponseResult struct {
	Response *ProductAddResponse `json:"product_add_response"`
}





/* 两种方式查看一个产品详细信息:  
传入product_id来查询  
传入cid和props来查询 */
type ProductGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传. */
func (r *ProductGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234” */
func (r *ProductGetRequest) SetCustomerProps(value string) {
	r.SetValue("customer_props", value)
}

/* 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔. */
func (r *ProductGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。 
不填写此值则默认取C2C的产品信息。 */
func (r *ProductGetRequest) SetMarketId(value string) {
	r.SetValue("market_id", value)
}

/* Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询 */
func (r *ProductGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729. */
func (r *ProductGetRequest) SetProps(value string) {
	r.SetValue("props", value)
}


func (r *ProductGetRequest) GetResponse(accessToken string) (*ProductGetResponse, []byte, error) {
	var resp ProductGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductGetResponse struct {
	Product *Product `json:"product"`
}

type ProductGetResponseResult struct {
	Response *ProductGetResponse `json:"product_get_response"`
}





/* 1.传入非主图ID  
2.传入产品ID  
删除产品非主图 */
type ProductImgDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 非主图ID */
func (r *ProductImgDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 产品ID.Product的id,通过taobao.product.add接口新增产品的时候会返回id. */
func (r *ProductImgDeleteRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}


func (r *ProductImgDeleteRequest) GetResponse(accessToken string) (*ProductImgDeleteResponse, []byte, error) {
	var resp ProductImgDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.img.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductImgDeleteResponse struct {
	ProductImg *ProductImg `json:"product_img"`
}

type ProductImgDeleteResponseResult struct {
	Response *ProductImgDeleteResponse `json:"product_img_delete_response"`
}





/* 1.传入产品ID  
2.传入图片内容  
注意：图片最大为500K,只支持JPG,GIF格式,如果需要传多张，可调多次 */
type ProductImgUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品图片ID.修改图片时需要传入 */
func (r *ProductImgUploadRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 图片内容.图片最大为500K,只支持JPG,GIF格式. */
func (r *ProductImgUploadRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 是否将该图片设为主图.可选值:true,false;默认值:false. */
func (r *ProductImgUploadRequest) SetIsMajor(value string) {
	r.SetValue("is_major", value)
}

/* 图片序号 */
func (r *ProductImgUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 产品ID.Product的id */
func (r *ProductImgUploadRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}


func (r *ProductImgUploadRequest) GetResponse(accessToken string) (*ProductImgUploadResponse, []byte, error) {
	var resp ProductImgUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.img.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductImgUploadResponse struct {
	ProductImg *ProductImg `json:"product_img"`
}

type ProductImgUploadResponseResult struct {
	Response *ProductImgUploadResponse `json:"product_img_upload_response"`
}





/* 1.传入属性图片ID  
2.传入产品ID  
删除一个产品的属性图片 */
type ProductPropimgDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性图片ID */
func (r *ProductPropimgDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 产品ID.Product的id. */
func (r *ProductPropimgDeleteRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}


func (r *ProductPropimgDeleteRequest) GetResponse(accessToken string) (*ProductPropimgDeleteResponse, []byte, error) {
	var resp ProductPropimgDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.propimg.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductPropimgDeleteResponse struct {
	ProductPropImg *ProductPropImg `json:"product_prop_img"`
}

type ProductPropimgDeleteResponseResult struct {
	Response *ProductPropimgDeleteResponse `json:"product_propimg_delete_response"`
}





/* 传入产品ID  
传入props,目前仅支持颜色属性.调用taobao.itemprops.get.v2取得颜色属性pid, 
再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;  
传入图片内容  
注意：图片最大为2M,只支持JPG,GIF,如果需要传多张，可调多次 */
type ProductPropimgUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品属性图片ID */
func (r *ProductPropimgUploadRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 图片内容.图片最大为2M,只支持JPG,GIF. */
func (r *ProductPropimgUploadRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 图片序号 */
func (r *ProductPropimgUploadRequest) SetPosition(value string) {
	r.SetValue("position", value)
}

/* 产品ID.Product的id */
func (r *ProductPropimgUploadRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串; */
func (r *ProductPropimgUploadRequest) SetProps(value string) {
	r.SetValue("props", value)
}


func (r *ProductPropimgUploadRequest) GetResponse(accessToken string) (*ProductPropimgUploadResponse, []byte, error) {
	var resp ProductPropimgUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.propimg.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductPropimgUploadResponse struct {
	ProductPropImg *ProductPropImg `json:"product_prop_img"`
}

type ProductPropimgUploadResponseResult struct {
	Response *ProductPropimgUploadResponse `json:"product_propimg_upload_response"`
}





/* 传入产品ID  
可修改字段：outer_id,binds,sale_props,name,price,desc,image  
注意：1.可以修改主图,不能修改子图片,主图最大500K,目前仅支持GIF,JPG 
      2.商城卖家产品发布24小时后不能作删除或修改操作 */
type ProductUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid */
func (r *ProductUpdateRequest) SetBinds(value string) {
	r.SetValue("binds", value)
}

/* 产品描述.最大不超过25000个字符 */
func (r *ProductUpdateRequest) SetDesc(value string) {
	r.SetValue("desc", value)
}

/* 存放产品扩展信息，由List(ProductExtraInfo)转化成jsonArray存入. */
func (r *ProductUpdateRequest) SetExtraInfo(value string) {
	r.SetValue("extra_info", value)
}

/* 产品主图.最大500K,目前仅支持GIF,JPG */
func (r *ProductUpdateRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 是否是主图 */
func (r *ProductUpdateRequest) SetMajor(value string) {
	r.SetValue("major", value)
}

/* 市场ID，1为更新C2C市场的产品信息， 2为更新B2C市场的产品信息。 
不填写此值则C用户更新B2C市场的产品信息，B用户更新B2C市场的产品信息。 */
func (r *ProductUpdateRequest) SetMarketId(value string) {
	r.SetValue("market_id", value)
}

/* 产品名称.最大不超过30个字符 */
func (r *ProductUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 自定义非关键属性 */
func (r *ProductUpdateRequest) SetNativeUnkeyprops(value string) {
	r.SetValue("native_unkeyprops", value)
}

/* 外部产品ID */
func (r *ProductUpdateRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}

/* 保证清单。 */
func (r *ProductUpdateRequest) SetPackingList(value string) {
	r.SetValue("packing_list", value)
}

/* 产品市场价.精确到2位小数;单位为元.如:200.07 */
func (r *ProductUpdateRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 产品ID */
func (r *ProductUpdateRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid */
func (r *ProductUpdateRequest) SetSaleProps(value string) {
	r.SetValue("sale_props", value)
}

/* 产品卖点描述，最长40个字节 */
func (r *ProductUpdateRequest) SetSellPt(value string) {
	r.SetValue("sell_pt", value)
}


func (r *ProductUpdateRequest) GetResponse(accessToken string) (*ProductUpdateResponse, []byte, error) {
	var resp ProductUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.product.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductUpdateResponse struct {
	Product *Product `json:"product"`
}

type ProductUpdateResponseResult struct {
	Response *ProductUpdateResponse `json:"product_update_response"`
}





/* 根据淘宝会员帐号搜索所有产品信息，只有天猫商家发布商品时才需要用到。  
注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页 */
type ProductsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔 */
func (r *ProductsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 用户昵称 */
func (r *ProductsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始. */
func (r *ProductsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值为40 */
func (r *ProductsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *ProductsGetRequest) GetResponse(accessToken string) (*ProductsGetResponse, []byte, error) {
	var resp ProductsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.products.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductsGetResponse struct {
	Products *ProductListObject `json:"products"`
}

type ProductsGetResponseResult struct {
	Response *ProductsGetResponse `json:"products_get_response"`
}





/* 只有天猫商家发布商品时才需要用到； 
两种方式搜索所有产品信息(二种至少传一种):  
传入关键字q搜索  
传入cid和props搜索  
返回值支持:product_id,name,pic_path,cid,props,price,tsc 
当用户指定了cid并且cid为垂直市场（3C电器城、鞋城）的类目id时，默认只返回小二确认的产品。如果用户没有指定cid，或cid为普通的类目，默认返回商家确认或小二确认的产品。如果用户自定了status字段，以指定的status类型为准 */
type ProductsSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品类目ID.调用taobao.itemcats.get获取. */
func (r *ProductsSearchRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234” */
func (r *ProductsSearchRequest) SetCustomerProps(value string) {
	r.SetValue("customer_props", value)
}

/* 需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用","分隔.新增字段status(product的当前状态) */
func (r *ProductsSearchRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。 
不填写此值则默认取C2C的产品信息。 */
func (r *ProductsSearchRequest) SetMarketId(value string) {
	r.SetValue("market_id", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始. */
func (r *ProductsSearchRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值为40. */
func (r *ProductsSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid  
,再用taobao.itempropvalues.get取得vid. */
func (r *ProductsSearchRequest) SetProps(value string) {
	r.SetValue("props", value)
}

/* 搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个 */
func (r *ProductsSearchRequest) SetQ(value string) {
	r.SetValue("q", value)
}

/* 想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用","分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入"3",只要商家确认的传入"0"，既要小二确认又要商家确认的传入"0,3"。目前只支持者两种类型的状态搜索，输入其他状态无效。 */
func (r *ProductsSearchRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型 */
func (r *ProductsSearchRequest) SetVerticalMarket(value string) {
	r.SetValue("vertical_market", value)
}


func (r *ProductsSearchRequest) GetResponse(accessToken string) (*ProductsSearchResponse, []byte, error) {
	var resp ProductsSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.products.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ProductsSearchResponse struct {
	Products *ProductListObject `json:"products"`
	TotalResults int `json:"total_results"`
}

type ProductsSearchResponseResult struct {
	Response *ProductsSearchResponse `json:"products_search_response"`
}





/* 跟据卖家设定的Sku的外部id获取商品，如果一个outer_id对应多个Sku会返回所有符合条件的sku  
这个Sku所属卖家从传入的session中获取，需要session绑定(注：iid标签里是num_iid的值，可以用作num_iid使用) */
type SkusCustomGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值：Sku结构体中的所有字段；字段之间用“,”隔开 */
func (r *SkusCustomGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* Sku的外部商家ID */
func (r *SkusCustomGetRequest) SetOuterId(value string) {
	r.SetValue("outer_id", value)
}


func (r *SkusCustomGetRequest) GetResponse(accessToken string) (*SkusCustomGetResponse, []byte, error) {
	var resp SkusCustomGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.skus.custom.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SkusCustomGetResponse struct {
	Skus *SkuListObject `json:"skus"`
}

type SkusCustomGetResponseResult struct {
	Response *SkusCustomGetResponse `json:"skus_custom_get_response"`
}





/* 提供按照全量/增量的方式批量修改SKU库存的功能 */
type SkusQuantityUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品数字ID，必填参数 */
func (r *SkusQuantityUpdateRequest) SetNumIid(value string) {
	r.SetValue("num_iid", value)
}

/* 特殊可选，skuIdQuantities为空的时候用该字段通过outerId来指定sku和其库存修改值。格式为outerId:库存修改值;outerId:库存修改值。当skuIdQuantities不为空的时候该字段失效。当一个outerId对应多个sku时，所有匹配到的sku都会被修改库存。最多支持20个SKU同时修改。 */
func (r *SkusQuantityUpdateRequest) SetOuteridQuantities(value string) {
	r.SetValue("outerid_quantities", value)
}

/* sku库存批量修改入参，用于指定一批sku和每个sku的库存修改值，特殊可填。格式为skuId:库存修改值;skuId:库存修改值。最多支持20个SKU同时修改。 */
func (r *SkusQuantityUpdateRequest) SetSkuidQuantities(value string) {
	r.SetValue("skuid_quantities", value)
}

/* 库存更新方式，可选。1为全量更新，2为增量更新。如果不填，默认为全量更新。当选择全量更新时，如果库存更新值传入的是负数，会出错并返回错误码；当选择增量更新时，如果库存更新值为负数且绝对值大于当前库存，则sku库存会设置为0. */
func (r *SkusQuantityUpdateRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *SkusQuantityUpdateRequest) GetResponse(accessToken string) (*SkusQuantityUpdateResponse, []byte, error) {
	var resp SkusQuantityUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.skus.quantity.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SkusQuantityUpdateResponse struct {
	Item *Item `json:"item"`
}

type SkusQuantityUpdateResponseResult struct {
	Response *SkusQuantityUpdateResponse `json:"skus_quantity_update_response"`
}





/* 提供异步获取三个月内的所有（出售中和仓库中）商品详情信息接口。 */
type TopatsItemsAllGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品修改结束时间，格式yyyyMMdd，取值范围：前90天内~昨天，其中start_time<=end_time，如20120531相当于取商品修改时间到2012-05-31 23:59:59为止的商品。注：如果start_time和end_time相同，表示取一天的商品数据。<span style="color:red">强烈建议图书类卖家把三个月商品拆分成3次来获取，否则单个任务会消耗很长时间。<span> */
func (r *TopatsItemsAllGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* Item结构体中的所有字段。<span style="color:red">请尽量按需获取，如果只取num_iid和modified字段，获取商品数据速度会超快。</span> */
func (r *TopatsItemsAllGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 商品修改开始时间，格式yyyyMMdd，取值范围：前90天内~昨天。如：20120501相当于取商品修改时间从2012-05-01 00:00:00开始的订单。 */
func (r *TopatsItemsAllGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 商品状态，可选值：onsale,instock。多个值用半角逗号分隔。默认查询两个状态的商品。 */
func (r *TopatsItemsAllGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *TopatsItemsAllGetRequest) GetResponse(accessToken string) (*TopatsItemsAllGetResponse, []byte, error) {
	var resp TopatsItemsAllGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.items.all.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsItemsAllGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsItemsAllGetResponseResult struct {
	Response *TopatsItemsAllGetResponse `json:"topats_items_all_get_response"`
}





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





/* 获取到所有的类目，品牌的控制信息，达尔文类目是会以品牌和类目作为控制的力度来管理商品，所有的控制信息可以通过这个接口获取到信息。 */
type TmallBrandcatControlGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *TmallBrandcatControlGetRequest) GetResponse(accessToken string) (*TmallBrandcatControlGetResponse, []byte, error) {
	var resp TmallBrandcatControlGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.brandcat.control.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallBrandcatControlGetResponse struct {
	BrandCatControlInfo *BrandCatControlInfo `json:"brand_cat_control_info"`
}

type TmallBrandcatControlGetResponseResult struct {
	Response *TmallBrandcatControlGetResponse `json:"tmall_brandcat_control_get_response"`
}





/* 针对监管类目品牌关键属性可输入判断逻辑比较复杂，提供简化属性可输入判断的接口，该接口兼容非关键属性和非监管类目品牌 */
type TmallBrandcatPropinputGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 品牌ID，如果类目没有品牌，指定null */
func (r *TmallBrandcatPropinputGetRequest) SetBrandId(value string) {
	r.SetValue("brand_id", value)
}

/* 类目ID */
func (r *TmallBrandcatPropinputGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 属性ID，如果属性有子属性，请指定最后一级子属性ID，tmall.brandcat.propinput.get返回的即为的该属性ID对应的输入特征，对于有子属性模板的情况指定顶级属性ID即可 */
func (r *TmallBrandcatPropinputGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}


func (r *TmallBrandcatPropinputGetRequest) GetResponse(accessToken string) (*TmallBrandcatPropinputGetResponse, []byte, error) {
	var resp TmallBrandcatPropinputGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.brandcat.propinput.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallBrandcatPropinputGetResponse struct {
	PropertyInput *PropertyInputDO `json:"property_input"`
}

type TmallBrandcatPropinputGetResponseResult struct {
	Response *TmallBrandcatPropinputGetResponse `json:"tmall_brandcat_propinput_get_response"`
}





/* 获取被管控的品牌类目下销售信息，销售信息又分成两中，一种是规格属性，要求针对这个属性创建规格信息，第二种是非规格属性（一般会被称为营销属性），这种信息主要是会影响到价格的变化，在做sku的时候会使用到。 */
type TmallBrandcatSalesproGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 被管控的品牌Id */
func (r *TmallBrandcatSalesproGetRequest) SetBrandId(value string) {
	r.SetValue("brand_id", value)
}

/* 被管控的类目Id */
func (r *TmallBrandcatSalesproGetRequest) SetCatId(value string) {
	r.SetValue("cat_id", value)
}


func (r *TmallBrandcatSalesproGetRequest) GetResponse(accessToken string) (*TmallBrandcatSalesproGetResponse, []byte, error) {
	var resp TmallBrandcatSalesproGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.brandcat.salespro.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallBrandcatSalesproGetResponse struct {
	CatBrandSaleProps *CatBrandSalePropListObject `json:"cat_brand_sale_props"`
}

type TmallBrandcatSalesproGetResponseResult struct {
	Response *TmallBrandcatSalesproGetResponse `json:"tmall_brandcat_salespro_get_response"`
}





/* 增加产品规格 */
type TmallProductSpecAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品二维码 */
func (r *TmallProductSpecAddRequest) SetBarcode(value string) {
	r.SetValue("barcode", value)
}

/* 存放产品规格认证类型-认证图片url映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为调用tmall.product.spec.pic.upload返回的认证图片url文本 */
func (r *TmallProductSpecAddRequest) SetCertifiedPicStr(value string) {
	r.SetValue("certified_pic_str", value)
}

/* 存放产品规格认证类型-认证文本映射信息，格式为k:v;k:v;，其中key为认证类型数字id，value为认证文本值 */
func (r *TmallProductSpecAddRequest) SetCertifiedTxtStr(value string) {
	r.SetValue("certified_txt_str", value)
}

/* 产品基础色，数据格式为：pid:vid:rvid1,rvid2,rvid3;pid:vid:rvid1 */
func (r *TmallProductSpecAddRequest) SetChangeProp(value string) {
	r.SetValue("change_prop", value)
}

/* 用户自定义销售属性，结构：pid1:value1;pid2:value2。在 */
func (r *TmallProductSpecAddRequest) SetCustomerSpecProps(value string) {
	r.SetValue("customer_spec_props", value)
}

/* 产品图片 */
func (r *TmallProductSpecAddRequest) SetImage(value string) {
	r.SetValue("image", value)
}

/* 产品规格吊牌价，以分为单位，无默认值，上限999999999 */
func (r *TmallProductSpecAddRequest) SetLabelPrice(value string) {
	r.SetValue("label_price", value)
}

/* 产品上市时间 */
func (r *TmallProductSpecAddRequest) SetMarketTime(value string) {
	r.SetValue("market_time", value)
}

/* 产品货号 */
func (r *TmallProductSpecAddRequest) SetProductCode(value string) {
	r.SetValue("product_code", value)
}

/* 产品ID */
func (r *TmallProductSpecAddRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 产品的规格属性 */
func (r *TmallProductSpecAddRequest) SetSpecProps(value string) {
	r.SetValue("spec_props", value)
}

/* 规格属性别名,只允许传颜色别名 */
func (r *TmallProductSpecAddRequest) SetSpecPropsAlias(value string) {
	r.SetValue("spec_props_alias", value)
}


func (r *TmallProductSpecAddRequest) GetResponse(accessToken string) (*TmallProductSpecAddResponse, []byte, error) {
	var resp TmallProductSpecAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.product.spec.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallProductSpecAddResponse struct {
	ProductSpec *ProductSpec `json:"product_spec"`
}

type TmallProductSpecAddResponseResult struct {
	Response *TmallProductSpecAddResponse `json:"tmall_product_spec_add_response"`
}





/* 通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核 
通过参看这个ProductSpec的status判断： 
1:表示审核通过 
3:表示等待审核。 
如果你的id找不到数据，那么就是审核被拒绝。 */
type TmallProductSpecGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要获取信息的产品规格信息。 */
func (r *TmallProductSpecGetRequest) SetSpecId(value string) {
	r.SetValue("spec_id", value)
}


func (r *TmallProductSpecGetRequest) GetResponse(accessToken string) (*TmallProductSpecGetResponse, []byte, error) {
	var resp TmallProductSpecGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.product.spec.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallProductSpecGetResponse struct {
	ProductSpec *ProductSpec `json:"product_spec"`
}

type TmallProductSpecGetResponseResult struct {
	Response *TmallProductSpecGetResponse `json:"tmall_product_spec_get_response"`
}





/* 上传指定类型的产品规格认证文件，并返回存有上传成功图片url的产品规格对象 */
type TmallProductSpecPicUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 上传的认证图片文件 */
func (r *TmallProductSpecPicUploadRequest) SetCertifyPic(value string) {
	r.SetValue("certify_pic", value)
}

/* 上传的认证图片的认证类型<br>
1：代表产品包装正面图<br>
2：代表完整产品资质<br>
3：代表产品包装反面图<br>
4：代表产品包装侧面图<br>
5：代表产品包装条形码特写<br>
6：代表特殊用途化妆品批准文号<br>
7：代表3C认证图标<br> */
func (r *TmallProductSpecPicUploadRequest) SetCertifyType(value string) {
	r.SetValue("certify_type", value)
}


func (r *TmallProductSpecPicUploadRequest) GetResponse(accessToken string) (*TmallProductSpecPicUploadResponse, []byte, error) {
	var resp TmallProductSpecPicUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.product.spec.pic.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallProductSpecPicUploadResponse struct {
	SpecPicUrl string `json:"spec_pic_url"`
}

type TmallProductSpecPicUploadResponseResult struct {
	Response *TmallProductSpecPicUploadResponse `json:"tmall_product_spec_pic_upload_response"`
}





/* 按productID下载或品牌下载产品规格，返回一组的产品规格信息。 */
type TmallProductSpecsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 类目的ID号，该id必须和properties同时传入。 
而且只有当product_id不传入的时候才起效果。 */
func (r *TmallProductSpecsGetRequest) SetCatId(value string) {
	r.SetValue("cat_id", value)
}

/* 产品的ID。这个不能和properties和cat_id同时起效果 */
func (r *TmallProductSpecsGetRequest) SetProductId(value string) {
	r.SetValue("product_id", value)
}

/* 关键属性的字符串，pid:vid;pid:vid 
该字段必须和cat_id同时传入才起效果。 而且只有当product_id不传入的时候才起效果。 */
func (r *TmallProductSpecsGetRequest) SetProperties(value string) {
	r.SetValue("properties", value)
}


func (r *TmallProductSpecsGetRequest) GetResponse(accessToken string) (*TmallProductSpecsGetResponse, []byte, error) {
	var resp TmallProductSpecsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.product.specs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallProductSpecsGetResponse struct {
	ProductSpecs *ProductSpecListObject `json:"product_specs"`
}

type TmallProductSpecsGetResponseResult struct {
	Response *TmallProductSpecsGetResponse `json:"tmall_product_specs_get_response"`
}





/* 批量根据specId查询产品规格审核信息包括产品规格状态，申请人，拒绝原因等 */
type TmallProductSpecsTicketGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 产品规格ID，多个用逗号分隔 */
func (r *TmallProductSpecsTicketGetRequest) SetSpecIds(value string) {
	r.SetValue("spec_ids", value)
}


func (r *TmallProductSpecsTicketGetRequest) GetResponse(accessToken string) (*TmallProductSpecsTicketGetResponse, []byte, error) {
	var resp TmallProductSpecsTicketGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "tmall.product.specs.ticket.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TmallProductSpecsTicketGetResponse struct {
	Tickets *TicketListObject `json:"tickets"`
}

type TmallProductSpecsTicketGetResponseResult struct {
	Response *TmallProductSpecsTicketGetResponse `json:"tmall_product_specs_ticket_get_response"`
}



