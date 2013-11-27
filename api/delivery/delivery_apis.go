// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package delivery

import (
	"github.com/changkong/open_taobao"
)





/* 查询标准地址区域代码信息 参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20100623_402652267.htm */
type AreasGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip. */
func (r *AreasGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *AreasGetRequest) GetResponse(accessToken string) (*AreasGetResponse, []byte, error) {
	var resp AreasGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.areas.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type AreasGetResponse struct {
	Areas []*Area `json:"areas"`
}

type AreasGetResponseResult struct {
	Response *AreasGetResponse `json:"areas_get_response"`
}





/* 新增运费模板 */
type DeliveryTemplateAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
func (r *DeliveryTemplateAddRequest) SetAssumer(value string) {
	r.SetValue("assumer", value)
}

/* 卖家发货地址区域ID 
<br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm  
</font> 
 
<br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font> */
func (r *DeliveryTemplateAddRequest) SetConsignAreaId(value string) {
	r.SetValue("consign_area_id", value)
}

/* 运费模板的名称，长度不能超过50个字节 */
func (r *DeliveryTemplateAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 增费：输入0.00-999.99（最多包含两位小数） 
 
<br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font> 
 
 
<br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateAddFees(value string) {
	r.SetValue("template_add_fees", value)
}

/* 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米） 
<br/> 
<br/> 
<br/> 
 
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateAddStandards(value string) {
	r.SetValue("template_add_standards", value)
}

/* 邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
<font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateDests(value string) {
	r.SetValue("template_dests", value)
}

/* 首费：输入0.00-999.99（最多包含两位小数） 
<br/><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateStartFees(value string) {
	r.SetValue("template_start_fees", value)
}

/* 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米） 
<br/> 
<br/> 
<br/> 
 
 
 
 
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateStartStandards(value string) {
	r.SetValue("template_start_standards", value)
}

/* 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 
如: post;express;ems;cod 
<br/><br/>
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分
。template_dests(指定地区)
template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font>
<br><br/><br/>

<font color=blue>
注意：<br/>
1、post,ems,express三种运费方式必须填写一个<br/>
2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture
只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>
3、如果是货到付款用户当没有填写cod运送方式的时候运费模板会默认继承express的费用为cod的费用<br/>
4、如果是保障速递户当没有填写bzsd运送方式的时候运费模板会默认继承express的费用为bzsd的费用<br/>
5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的
</font> */
func (r *DeliveryTemplateAddRequest) SetTemplateTypes(value string) {
	r.SetValue("template_types", value)
}

/* 可选值：0<br>说明：<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费 
 <br>3:表示按宝贝体积计算运费 */
func (r *DeliveryTemplateAddRequest) SetValuation(value string) {
	r.SetValue("valuation", value)
}


func (r *DeliveryTemplateAddRequest) GetResponse(accessToken string) (*DeliveryTemplateAddResponse, []byte, error) {
	var resp DeliveryTemplateAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.delivery.template.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type DeliveryTemplateAddResponse struct {
	DeliveryTemplate *DeliveryTemplate `json:"delivery_template"`
}

type DeliveryTemplateAddResponseResult struct {
	Response *DeliveryTemplateAddResponse `json:"delivery_template_add_response"`
}





/* 根据用户指定的模板ID删除指定的模板 */
type DeliveryTemplateDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 运费模板ID */
func (r *DeliveryTemplateDeleteRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}


func (r *DeliveryTemplateDeleteRequest) GetResponse(accessToken string) (*DeliveryTemplateDeleteResponse, []byte, error) {
	var resp DeliveryTemplateDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.delivery.template.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type DeliveryTemplateDeleteResponse struct {
	Complete bool `json:"complete"`
}

type DeliveryTemplateDeleteResponseResult struct {
	Response *DeliveryTemplateDeleteResponse `json:"delivery_template_delete_response"`
}





/* 获取用户指定运费模板信息 */
type DeliveryTemplateGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。 <br/>  
<B> 
可选值:示例中的值;各字段之间用","隔开 
</B> 
<br/><br/>  
<font color=blue> 
template_id：查询模板ID <br/>  
template_name:查询模板名称 <br/>  
assumer：查询服务承担放 <br/>  
valuation:查询计价规则 <br/>  
supports:查询增值服务列表 <br/>  
created:查询模板创建时间 <br/>  
modified:查询修改时间<br/>  
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/>  
address:卖家地址信息 
</font> 
<br/><br/>  
<font color=bule> 
query_cod:查询货到付款运费信息<br/>  
query_post:查询平邮运费信息 <br/>  
query_express:查询快递公司运费信息 <br/>  
query_ems:查询EMS运费信息<br/>  
query_bzsd:查询普通保障速递运费信息<br/>  
query_wlb:查询物流宝保障速递运费信息<br/>  
query_furniture:查询家装物流运费信息 
<font color=blue> 
<br/><br/> 
<font color=red>不能有空格</font> */
func (r *DeliveryTemplateGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 需要查询的运费模板ID列表 */
func (r *DeliveryTemplateGetRequest) SetTemplateIds(value string) {
	r.SetValue("template_ids", value)
}

/* 在没有登录的情况下根据淘宝用户昵称查询指定的模板 */
func (r *DeliveryTemplateGetRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *DeliveryTemplateGetRequest) GetResponse(accessToken string) (*DeliveryTemplateGetResponse, []byte, error) {
	var resp DeliveryTemplateGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.delivery.template.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type DeliveryTemplateGetResponse struct {
	DeliveryTemplates []*DeliveryTemplate `json:"delivery_templates"`
	TotalResults int `json:"total_results"`
}

type DeliveryTemplateGetResponseResult struct {
	Response *DeliveryTemplateGetResponse `json:"delivery_template_get_response"`
}





/* 修改运费模板 */
type DeliveryTemplateUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费 */
func (r *DeliveryTemplateUpdateRequest) SetAssumer(value string) {
	r.SetValue("assumer", value)
}

/* 模板名称，长度不能大于50个字节 */
func (r *DeliveryTemplateUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateAddFees(value string) {
	r.SetValue("template_add_fees", value)
}

/* 增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateAddStandards(value string) {
	r.SetValue("template_add_standards", value)
}

/* 邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateDests(value string) {
	r.SetValue("template_dests", value)
}

/* 需要修改的模板对应的模板ID */
func (r *DeliveryTemplateUpdateRequest) SetTemplateId(value string) {
	r.SetValue("template_id", value)
}

/* 首费：输入0.01-999.99（最多包含两位小数） 
<br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateStartFees(value string) {
	r.SetValue("template_start_fees", value)
}

/* 首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateStartStandards(value string) {
	r.SetValue("template_start_standards", value)
}

/* 运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4  
如: post;express;ems;cod  
<br/><font color=red> 
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同.  
 </font> 
<br/> 
<font color=blue> 
普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。 
货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font> */
func (r *DeliveryTemplateUpdateRequest) SetTemplateTypes(value string) {
	r.SetValue("template_types", value)
}


func (r *DeliveryTemplateUpdateRequest) GetResponse(accessToken string) (*DeliveryTemplateUpdateResponse, []byte, error) {
	var resp DeliveryTemplateUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.delivery.template.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type DeliveryTemplateUpdateResponse struct {
	Complete bool `json:"complete"`
}

type DeliveryTemplateUpdateResponseResult struct {
	Response *DeliveryTemplateUpdateResponse `json:"delivery_template_update_response"`
}





/* 根据用户ID获取用户下所有模板 */
type DeliveryTemplatesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。 <br/>  
<B> 
可选值:示例中的值;各字段之间用","隔开 
</B> 
<br/><br/>  
<font color=blue> 
template_id：查询模板ID <br/>  
template_name:查询模板名称 <br/>  
assumer：查询服务承担放 <br/>  
valuation:查询计价规则 <br/>  
supports:查询增值服务列表 <br/>  
created:查询模板创建时间 <br/>  
modified:查询修改时间<br/>  
consign_area_id:运费模板上设置的卖家发货地址最小级别ID<br/>  
address:卖家地址信息 
</font> 
<br/><br/>  
<font color=bule> 
query_cod:查询货到付款运费信息<br/>  
query_post:查询平邮运费信息 <br/>  
query_express:查询快递公司运费信息 <br/>  
query_ems:查询EMS运费信息<br/>  
query_bzsd:查询普通保障速递运费信息<br/>  
query_wlb:查询物流宝保障速递运费信息<br/>  
query_furniture:查询家装物流运费信息 
<font color=blue> 
<br/><br/> 
<font color=red>不能有空格</font> */
func (r *DeliveryTemplatesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *DeliveryTemplatesGetRequest) GetResponse(accessToken string) (*DeliveryTemplatesGetResponse, []byte, error) {
	var resp DeliveryTemplatesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.delivery.templates.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type DeliveryTemplatesGetResponse struct {
	DeliveryTemplates []*DeliveryTemplate `json:"delivery_templates"`
	TotalResults int `json:"total_results"`
}

type DeliveryTemplatesGetResponseResult struct {
	Response *DeliveryTemplatesGetResponse `json:"delivery_templates_get_response"`
}





/* 通过此接口新增卖家地址库,卖家最多可添加5条地址库,新增第一条卖家地址，将会自动设为默认地址库 */
type LogisticsAddressAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 详细街道地址，不需要重复填写省/市/区 */
func (r *LogisticsAddressAddRequest) SetAddr(value string) {
	r.SetValue("addr", value)
}

/* 默认退货地址。<br> 
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
func (r *LogisticsAddressAddRequest) SetCancelDef(value string) {
	r.SetValue("cancel_def", value)
}

/* 所在市 */
func (r *LogisticsAddressAddRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 联系人姓名 <font color='red'>长度不可超过20个字节</font> */
func (r *LogisticsAddressAddRequest) SetContactName(value string) {
	r.SetValue("contact_name", value)
}

/* 区、县 
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (r *LogisticsAddressAddRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 默认取货地址。<br> 
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
func (r *LogisticsAddressAddRequest) SetGetDef(value string) {
	r.SetValue("get_def", value)
}

/* 备注,<br><font color='red'>备注不能超过256字节</font> */
func (r *LogisticsAddressAddRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 手机号码，手机与电话必需有一个 
<br><font color='red'>手机号码不能超过20位</font> */
func (r *LogisticsAddressAddRequest) SetMobilePhone(value string) {
	r.SetValue("mobile_phone", value)
}

/* 电话号码,手机与电话必需有一个 */
func (r *LogisticsAddressAddRequest) SetPhone(value string) {
	r.SetValue("phone", value)
}

/* 所在省 */
func (r *LogisticsAddressAddRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 公司名称,<br><font color="red">公司名称长能不能超过20字节</font> */
func (r *LogisticsAddressAddRequest) SetSellerCompany(value string) {
	r.SetValue("seller_company", value)
}

/* 地区邮政编码 
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (r *LogisticsAddressAddRequest) SetZipCode(value string) {
	r.SetValue("zip_code", value)
}


func (r *LogisticsAddressAddRequest) GetResponse(accessToken string) (*LogisticsAddressAddResponse, []byte, error) {
	var resp LogisticsAddressAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.address.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsAddressAddResponse struct {
	AddressResult *AddressResult `json:"address_result"`
}

type LogisticsAddressAddResponseResult struct {
	Response *LogisticsAddressAddResponse `json:"logistics_address_add_response"`
}





/* 卖家地址库修改 */
type LogisticsAddressModifyRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 详细街道地址，不需要重复填写省/市/区 */
func (r *LogisticsAddressModifyRequest) SetAddr(value string) {
	r.SetValue("addr", value)
}

/* 默认退货地址。<br> 
<font color='red'>选择此项(true)，将当前地址设为默认退货地址，撤消原默认退货地址</font> */
func (r *LogisticsAddressModifyRequest) SetCancelDef(value string) {
	r.SetValue("cancel_def", value)
}

/* 所在市 */
func (r *LogisticsAddressModifyRequest) SetCity(value string) {
	r.SetValue("city", value)
}

/* 地址库ID */
func (r *LogisticsAddressModifyRequest) SetContactId(value string) {
	r.SetValue("contact_id", value)
}

/* 联系人姓名 
<font color='red'>长度不可超过20个字节</font> */
func (r *LogisticsAddressModifyRequest) SetContactName(value string) {
	r.SetValue("contact_name", value)
}

/* 区、县 
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (r *LogisticsAddressModifyRequest) SetCountry(value string) {
	r.SetValue("country", value)
}

/* 默认取货地址。<br> 
<font color='red'>选择此项(true)，将当前地址设为默认取货地址，撤消原默认取货地址</font> */
func (r *LogisticsAddressModifyRequest) SetGetDef(value string) {
	r.SetValue("get_def", value)
}

/* 备注,<br><font color='red'>备注不能超过256字节</font> */
func (r *LogisticsAddressModifyRequest) SetMemo(value string) {
	r.SetValue("memo", value)
}

/* 手机号码，手机与电话必需有一个 <br><font color='red'>手机号码不能超过20位</font> */
func (r *LogisticsAddressModifyRequest) SetMobilePhone(value string) {
	r.SetValue("mobile_phone", value)
}

/* 电话号码,手机与电话必需有一个 */
func (r *LogisticsAddressModifyRequest) SetPhone(value string) {
	r.SetValue("phone", value)
}

/* 所在省 */
func (r *LogisticsAddressModifyRequest) SetProvince(value string) {
	r.SetValue("province", value)
}

/* 公司名称, 
<br><font color='red'>公司名称长能不能超过20字节</font> */
func (r *LogisticsAddressModifyRequest) SetSellerCompany(value string) {
	r.SetValue("seller_company", value)
}

/* 地区邮政编码 
<br><font color='red'>如果所在地区是海外的可以为空，否则为必参</font> */
func (r *LogisticsAddressModifyRequest) SetZipCode(value string) {
	r.SetValue("zip_code", value)
}


func (r *LogisticsAddressModifyRequest) GetResponse(accessToken string) (*LogisticsAddressModifyResponse, []byte, error) {
	var resp LogisticsAddressModifyResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.address.modify", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsAddressModifyResponse struct {
	AddressResult *AddressResult `json:"address_result"`
}

type LogisticsAddressModifyResponseResult struct {
	Response *LogisticsAddressModifyResponse `json:"logistics_address_modify_response"`
}





/* 用此接口删除卖家地址库 */
type LogisticsAddressRemoveRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 地址库ID */
func (r *LogisticsAddressRemoveRequest) SetContactId(value string) {
	r.SetValue("contact_id", value)
}


func (r *LogisticsAddressRemoveRequest) GetResponse(accessToken string) (*LogisticsAddressRemoveResponse, []byte, error) {
	var resp LogisticsAddressRemoveResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.address.remove", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsAddressRemoveResponse struct {
	AddressResult *AddressResult `json:"address_result"`
}

type LogisticsAddressRemoveResponseResult struct {
	Response *LogisticsAddressRemoveResponse `json:"logistics_address_remove_response"`
}





/* 通过此接口查询卖家地址库， */
type LogisticsAddressSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 可选，参数列表如下：<br><font color='red'>
no_def:查询非默认地址<br>
get_def:查询默认取货地址<br>
cancel_def:查询默认退货地址<br>
缺省此参数时，查询所有当前用户的地址库
</font> */
func (r *LogisticsAddressSearchRequest) SetRdef(value string) {
	r.SetValue("rdef", value)
}


func (r *LogisticsAddressSearchRequest) GetResponse(accessToken string) (*LogisticsAddressSearchResponse, []byte, error) {
	var resp LogisticsAddressSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.address.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsAddressSearchResponse struct {
	Addresses []*AddressResult `json:"addresses"`
}

type LogisticsAddressSearchResponseResult struct {
	Response *LogisticsAddressSearchResponse `json:"logistics_address_search_response"`
}





/* 查询淘宝网合作的物流公司信息，用于发货接口。 */
type LogisticsCompaniesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开. 
如:id,code,name,reg_mail_no 
<br><font color='red'>说明：</font> 
<br>id：物流公司ID 
<br>code：物流公司code 
<br>name：物流公司名称 
<br>reg_mail_no：物流公司对应的运单规则 */
func (r *LogisticsCompaniesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司. */
func (r *LogisticsCompaniesGetRequest) SetIsRecommended(value string) {
	r.SetValue("is_recommended", value)
}

/* 推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司. */
func (r *LogisticsCompaniesGetRequest) SetOrderMode(value string) {
	r.SetValue("order_mode", value)
}


func (r *LogisticsCompaniesGetRequest) GetResponse(accessToken string) (*LogisticsCompaniesGetResponse, []byte, error) {
	var resp LogisticsCompaniesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.companies.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsCompaniesGetResponse struct {
	LogisticsCompanies []*LogisticsCompany `json:"logistics_companies"`
}

type LogisticsCompaniesGetResponseResult struct {
	Response *LogisticsCompaniesGetResponse `json:"logistics_companies_get_response"`
}





/* 创建物流订单，并发货。 */
type LogisticsConsignOrderCreateandsendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流公司ID */
func (r *LogisticsConsignOrderCreateandsendRequest) SetCompanyId(value string) {
	r.SetValue("company_id", value)
}

/* 物品的json数据。 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetItemJsonString(value string) {
	r.SetValue("item_json_string", value)
}

/* 物流订单物流类型，值固定选择：2 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetLogisType(value string) {
	r.SetValue("logis_type", value)
}

/* 运单号 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetMailNo(value string) {
	r.SetValue("mail_no", value)
}

/* 订单来源，值选择：30 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetOrderSource(value string) {
	r.SetValue("order_source", value)
}

/* 订单类型，值固定选择：30 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}

/* 收件人街道地址 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRAddress(value string) {
	r.SetValue("r_address", value)
}

/* 收件人区域ID */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRAreaId(value string) {
	r.SetValue("r_area_id", value)
}

/* 市 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRCityName(value string) {
	r.SetValue("r_city_name", value)
}

/* 区 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRDistName(value string) {
	r.SetValue("r_dist_name", value)
}

/* 手机号码 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRMobilePhone(value string) {
	r.SetValue("r_mobile_phone", value)
}

/* 收件人名称 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRName(value string) {
	r.SetValue("r_name", value)
}

/* 省 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRProvName(value string) {
	r.SetValue("r_prov_name", value)
}

/* 电话号码 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRTelephone(value string) {
	r.SetValue("r_telephone", value)
}

/* 收件人邮编 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetRZipCode(value string) {
	r.SetValue("r_zip_code", value)
}

/* 发件人街道地址 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSAddress(value string) {
	r.SetValue("s_address", value)
}

/* 发件人区域ID */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSAreaId(value string) {
	r.SetValue("s_area_id", value)
}

/* 市 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSCityName(value string) {
	r.SetValue("s_city_name", value)
}

/* 区 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSDistName(value string) {
	r.SetValue("s_dist_name", value)
}

/* 手机号码 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSMobilePhone(value string) {
	r.SetValue("s_mobile_phone", value)
}

/* 发件人名称 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSName(value string) {
	r.SetValue("s_name", value)
}

/* 省 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSProvName(value string) {
	r.SetValue("s_prov_name", value)
}

/* 电话号码 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSTelephone(value string) {
	r.SetValue("s_telephone", value)
}

/* 发件人出编 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetSZipCode(value string) {
	r.SetValue("s_zip_code", value)
}

/* 费用承担方式 1买家承担运费 2卖家承担运费 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetShipping(value string) {
	r.SetValue("shipping", value)
}

/* 交易流水号，淘外订单号或者商家内部交易流水号 */
func (r *LogisticsConsignOrderCreateandsendRequest) SetTradeId(value string) {
	r.SetValue("trade_id", value)
}

/* 用户ID */
func (r *LogisticsConsignOrderCreateandsendRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}


func (r *LogisticsConsignOrderCreateandsendRequest) GetResponse(accessToken string) (*LogisticsConsignOrderCreateandsendResponse, []byte, error) {
	var resp LogisticsConsignOrderCreateandsendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.consign.order.createandsend", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsConsignOrderCreateandsendResponse struct {
	IsSuccess bool `json:"is_success"`
	OrderId int `json:"order_id"`
	ResultDesc string `json:"result_desc"`
}

type LogisticsConsignOrderCreateandsendResponseResult struct {
	Response *LogisticsConsignOrderCreateandsendResponse `json:"logistics_consign_order_createandsend_response"`
}





/* 支持卖家发货后修改物流公司和运单号。支持订单类型支持在线下单和自己联系。 
自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。 
 
调用时订单状态是卖家已发货，自己联系在发货后24小时内在线下单未揽收成功才可使用 */
type LogisticsConsignResendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。 
<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font> */
func (r *LogisticsConsignResendRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* feature参数格式<br> 
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br> 
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br> 
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 
"|"不同商品间的分隔符。<br> 
例1商品和2商品，之间就用"|"分开。<br> 
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br> 
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br> 
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br> 
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br> 
参数格式：identCode=TIDA:12345,67890。 
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br> 
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (r *LogisticsConsignResendRequest) SetFeature(value string) {
	r.SetValue("feature", value)
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (r *LogisticsConsignResendRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
func (r *LogisticsConsignResendRequest) SetOutSid(value string) {
	r.SetValue("out_sid", value)
}

/* 商家的IP地址 */
func (r *LogisticsConsignResendRequest) SetSellerIp(value string) {
	r.SetValue("seller_ip", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (r *LogisticsConsignResendRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝交易ID */
func (r *LogisticsConsignResendRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsConsignResendRequest) GetResponse(accessToken string) (*LogisticsConsignResendResponse, []byte, error) {
	var resp LogisticsConsignResendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.consign.resend", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsConsignResendResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsConsignResendResponseResult struct {
	Response *LogisticsConsignResendResponse `json:"logistics_consign_resend_response"`
}





/* 用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货 */
type LogisticsDummySendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* feature参数格式<br> 
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br> 
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br> 
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 
"|"不同商品间的分隔符。<br> 
例1商品和2商品，之间就用"|"分开。<br> 
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br> 
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br> 
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br> 
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br> 
参数格式：identCode=TIDA:12345,67890。 
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br> 
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (r *LogisticsDummySendRequest) SetFeature(value string) {
	r.SetValue("feature", value)
}

/* 商家的IP地址 */
func (r *LogisticsDummySendRequest) SetSellerIp(value string) {
	r.SetValue("seller_ip", value)
}

/* 淘宝交易ID */
func (r *LogisticsDummySendRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsDummySendRequest) GetResponse(accessToken string) (*LogisticsDummySendResponse, []byte, error) {
	var resp LogisticsDummySendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.dummy.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsDummySendResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsDummySendResponseResult struct {
	Response *LogisticsDummySendResponse `json:"logistics_dummy_send_response"`
}





/* 用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。 */
type LogisticsOfflineSendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br> */
func (r *LogisticsOfflineSendRequest) SetCancelId(value string) {
	r.SetValue("cancel_id", value)
}

/* 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。非淘宝官方物流合作公司，填写具体的物流公司名称，如“顺丰”。 */
func (r *LogisticsOfflineSendRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* feature参数格式<br> 
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br> 
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br> 
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 
"|"不同商品间的分隔符。<br> 
例1商品和2商品，之间就用"|"分开。<br> 
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br> 
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br> 
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br> 
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br> 
参数格式：identCode=TIDA:12345,67890。 
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br> 
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (r *LogisticsOfflineSendRequest) SetFeature(value string) {
	r.SetValue("feature", value)
}

/* 表明是否是拆单
1表示拆单
0表示不拆单，默认值0 */
func (r *LogisticsOfflineSendRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；若company_code中传入的代码非淘宝官方物流合作公司，此处运单号不校验。 */
func (r *LogisticsOfflineSendRequest) SetOutSid(value string) {
	r.SetValue("out_sid", value)
}

/* 商家的IP地址 */
func (r *LogisticsOfflineSendRequest) SetSellerIp(value string) {
	r.SetValue("seller_ip", value)
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font> */
func (r *LogisticsOfflineSendRequest) SetSenderId(value string) {
	r.SetValue("sender_id", value)
}

/* 需要拆单发货的子订单集合，为空表示不做拆单发货。 */
func (r *LogisticsOfflineSendRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝交易ID */
func (r *LogisticsOfflineSendRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsOfflineSendRequest) GetResponse(accessToken string) (*LogisticsOfflineSendResponse, []byte, error) {
	var resp LogisticsOfflineSendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.offline.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOfflineSendResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsOfflineSendResponseResult struct {
	Response *LogisticsOfflineSendResponse `json:"logistics_offline_send_response"`
}





/* 调此接口取消发货的订单，重新选择物流公司发货。前提是物流公司未揽收货物。对未发货和已经被物流公司揽收的物流订单，是不能取消的。 */
type LogisticsOnlineCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 淘宝交易ID */
func (r *LogisticsOnlineCancelRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsOnlineCancelRequest) GetResponse(accessToken string) (*LogisticsOnlineCancelResponse, []byte, error) {
	var resp LogisticsOnlineCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.online.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOnlineCancelResponse struct {
	IsSuccess bool `json:"is_success"`
	ModifyTime string `json:"modify_time"`
	RecreatedOrderId int `json:"recreated_order_id"`
}

type LogisticsOnlineCancelResponseResult struct {
	Response *LogisticsOnlineCancelResponse `json:"logistics_online_cancel_response"`
}





/* 确认发货的目的是让交易流程继承走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】，然后买家才可以确认收货，货款打入卖家账号。货到付款的订单除外 */
type LogisticsOnlineConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 表明是否是拆单，默认值0，1表示拆单 */
func (r *LogisticsOnlineConfirmRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；若company_code中传入的代码非淘宝官方物流合作公司，此处运单号不校验。<br> */
func (r *LogisticsOnlineConfirmRequest) SetOutSid(value string) {
	r.SetValue("out_sid", value)
}

/* 商家的IP地址 */
func (r *LogisticsOnlineConfirmRequest) SetSellerIp(value string) {
	r.SetValue("seller_ip", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (r *LogisticsOnlineConfirmRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝交易ID */
func (r *LogisticsOnlineConfirmRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsOnlineConfirmRequest) GetResponse(accessToken string) (*LogisticsOnlineConfirmResponse, []byte, error) {
	var resp LogisticsOnlineConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.online.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOnlineConfirmResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsOnlineConfirmResponseResult struct {
	Response *LogisticsOnlineConfirmResponse `json:"logistics_online_confirm_response"`
}





/* 用户调用该接口可实现在线订单发货（支持货到付款） 
调用该接口实现在线下单发货，有两种情况：<br> 
<font color='red'>如果不输入运单号的情况：交易状态不会改变，需要调用taobao.logistics.online.confirm确认发货后交易状态才会变成卖家已发货。<br> 
如果输入运单号的情况发货：交易订单状态会直接变成卖家已发货 。</font> */
type LogisticsOnlineSendRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<br><font color='red'>如果为空，取的卖家的默认退货地址</font><br> */
func (r *LogisticsOnlineSendRequest) SetCancelId(value string) {
	r.SetValue("cancel_id", value)
}

/* 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。 
<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font> */
func (r *LogisticsOnlineSendRequest) SetCompanyCode(value string) {
	r.SetValue("company_code", value)
}

/* feature参数格式<br> 
范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br> 
identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br> 
“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 
"|"不同商品间的分隔符。<br> 
例1商品和2商品，之间就用"|"分开。<br> 
TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br> 
":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br> 
"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br> 
具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br> 
参数格式：identCode=TIDA:12345,67890。 
TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br> 
当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。 */
func (r *LogisticsOnlineSendRequest) SetFeature(value string) {
	r.SetValue("feature", value)
}

/* 表明是否是拆单，默认值0，1表示拆单 */
func (r *LogisticsOnlineSendRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入； */
func (r *LogisticsOnlineSendRequest) SetOutSid(value string) {
	r.SetValue("out_sid", value)
}

/* 商家的IP地址 */
func (r *LogisticsOnlineSendRequest) SetSellerIp(value string) {
	r.SetValue("seller_ip", value)
}

/* 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。<font color='red'>如果为空，取的卖家的默认取货地址</font> */
func (r *LogisticsOnlineSendRequest) SetSenderId(value string) {
	r.SetValue("sender_id", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (r *LogisticsOnlineSendRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝交易ID */
func (r *LogisticsOnlineSendRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsOnlineSendRequest) GetResponse(accessToken string) (*LogisticsOnlineSendResponse, []byte, error) {
	var resp LogisticsOnlineSendResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.online.send", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOnlineSendResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsOnlineSendResponseResult struct {
	Response *LogisticsOnlineSendResponse `json:"logistics_online_send_response"`
}





/* 查询物流订单的详细信息，涉及用户隐私字段。（注：该API主要是提供给卖家查询物流订单使用，买家查询物流订单，建议使用taobao.logistics.trace.search） */
type LogisticsOrdersDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *LogisticsOrdersDetailGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 创建时间结束.格式:yyyy-MM-dd HH:mm:ss */
func (r *LogisticsOrdersDetailGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 需返回的字段列表.可选值:Shipping 物流数据结构中所有字段.fileds中可以指定返回以上任意一个或者多个字段,以","分隔. */
func (r *LogisticsOrdersDetailGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
func (r *LogisticsOrdersDetailGetRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 页码.该字段没传 或 值<1 ,则默认page_no为1 */
func (r *LogisticsOrdersDetailGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.该字段没传 或 值<1 ，则默认page_size为40 */
func (r *LogisticsOrdersDetailGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 收货人姓名 */
func (r *LogisticsOrdersDetailGetRequest) SetReceiverName(value string) {
	r.SetValue("receiver_name", value)
}

/* 卖家是否发货.可选值:yes(是),no(否).如:yes. */
func (r *LogisticsOrdersDetailGetRequest) SetSellerConfirm(value string) {
	r.SetValue("seller_confirm", value)
}

/* 创建时间开始.格式:yyyy-MM-dd HH:mm:ss */
func (r *LogisticsOrdersDetailGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}

/* 物流状态.可查看数据结构 Shipping 中的status字段. */
func (r *LogisticsOrdersDetailGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易ID.如果加入tid参数的话,不用传其他的参数,但是仅会返回一条物流订单信息. */
func (r *LogisticsOrdersDetailGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

/* 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
func (r *LogisticsOrdersDetailGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *LogisticsOrdersDetailGetRequest) GetResponse(accessToken string) (*LogisticsOrdersDetailGetResponse, []byte, error) {
	var resp LogisticsOrdersDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.orders.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOrdersDetailGetResponse struct {
	Shippings []*Shipping `json:"shippings"`
	TotalResults int `json:"total_results"`
}

type LogisticsOrdersDetailGetResponseResult struct {
	Response *LogisticsOrdersDetailGetResponse `json:"logistics_orders_detail_get_response"`
}





/* 批量查询物流订单。（注：该API主要是提供给卖家查询物流订单使用，买家查询物流订单，建议使用taobao.logistics.trace.search） */
type LogisticsOrdersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 买家昵称 */
func (r *LogisticsOrdersGetRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 创建时间结束 */
func (r *LogisticsOrdersGetRequest) SetEndCreated(value string) {
	r.SetValue("end_created", value)
}

/* 需返回的字段列表.可选值:Shipping 物流数据结构中的以下字段: <br> 
tid,order_code,seller_nick,buyer_nick,delivery_start, delivery_end,out_sid,item_title,receiver_name, created,modified,status,type,freight_payer,seller_confirm,company_name,sub_tids,is_spilt；<br>多个字段之间用","分隔。如tid,seller_nick,buyer_nick,delivery_start。 */
func (r *LogisticsOrdersGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 谁承担运费.可选值:buyer(买家),seller(卖家).如:buyer */
func (r *LogisticsOrdersGetRequest) SetFreightPayer(value string) {
	r.SetValue("freight_payer", value)
}

/* 页码.该字段没传 或 值<1 ,则默认page_no为1 */
func (r *LogisticsOrdersGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.该字段没传 或 值<1 ,则默认page_size为40 */
func (r *LogisticsOrdersGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 收货人姓名 */
func (r *LogisticsOrdersGetRequest) SetReceiverName(value string) {
	r.SetValue("receiver_name", value)
}

/* 卖家是否发货.可选值:yes(是),no(否).如:yes */
func (r *LogisticsOrdersGetRequest) SetSellerConfirm(value string) {
	r.SetValue("seller_confirm", value)
}

/* 创建时间开始 */
func (r *LogisticsOrdersGetRequest) SetStartCreated(value string) {
	r.SetValue("start_created", value)
}

/* 物流状态.查看数据结构 Shipping 中的status字段. */
func (r *LogisticsOrdersGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 交易ID.如果加入tid参数的话,不用传其他的参数,若传入tid：非拆单场景，仅会返回一条物流订单信息；拆单场景，会返回多条物流订单信息 */
func (r *LogisticsOrdersGetRequest) SetTid(value string) {
	r.SetValue("tid", value)
}

/* 物流方式.可选值:post(平邮),express(快递),ems(EMS).如:post */
func (r *LogisticsOrdersGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *LogisticsOrdersGetRequest) GetResponse(accessToken string) (*LogisticsOrdersGetResponse, []byte, error) {
	var resp LogisticsOrdersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.orders.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOrdersGetResponse struct {
	Shippings []*Shipping `json:"shippings"`
	TotalResults int `json:"total_results"`
}

type LogisticsOrdersGetResponseResult struct {
	Response *LogisticsOrdersGetResponse `json:"logistics_orders_get_response"`
}





/* 卖家使用自己的物流公司发货，可以通过本接口将订单的仓内信息推送到淘宝，淘宝保存这些仓内信息，并可在页面展示这些仓内信息。 */
type LogisticsOrderstorePushRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 表明是否是拆单，默认值0，1表示拆单 */
func (r *LogisticsOrderstorePushRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 流转状态发生时间 */
func (r *LogisticsOrderstorePushRequest) SetOccureTime(value string) {
	r.SetValue("occure_time", value)
}

/* 仓内操作描述 */
func (r *LogisticsOrderstorePushRequest) SetOperateDetail(value string) {
	r.SetValue("operate_detail", value)
}

/* 快递业务员联系方式 */
func (r *LogisticsOrderstorePushRequest) SetOperatorContact(value string) {
	r.SetValue("operator_contact", value)
}

/* 快递业务员名称 */
func (r *LogisticsOrderstorePushRequest) SetOperatorName(value string) {
	r.SetValue("operator_name", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！ */
func (r *LogisticsOrderstorePushRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝订单交易号 */
func (r *LogisticsOrderstorePushRequest) SetTradeId(value string) {
	r.SetValue("trade_id", value)
}


func (r *LogisticsOrderstorePushRequest) GetResponse(accessToken string) (*LogisticsOrderstorePushResponse, []byte, error) {
	var resp LogisticsOrderstorePushResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.orderstore.push", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOrderstorePushResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsOrderstorePushResponseResult struct {
	Response *LogisticsOrderstorePushResponse `json:"logistics_orderstore_push_response"`
}





/* 卖家使用自己的物流公司发货，可以通过本接口将订单的流转信息推送到淘宝，淘宝保存这些流转信息，并可在页面展示这些流转信息。 */
type LogisticsOrdertracePushRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流公司名称 */
func (r *LogisticsOrdertracePushRequest) SetCompanyName(value string) {
	r.SetValue("company_name", value)
}

/* 流转节点的当前城市 */
func (r *LogisticsOrdertracePushRequest) SetCurrentCity(value string) {
	r.SetValue("current_city", value)
}

/* 网点名称 */
func (r *LogisticsOrdertracePushRequest) SetFacilityName(value string) {
	r.SetValue("facility_name", value)
}

/* 快递单号。各个快递公司的运单号格式不同。 */
func (r *LogisticsOrdertracePushRequest) SetMailNo(value string) {
	r.SetValue("mail_no", value)
}

/* 流转节点的下一个城市 */
func (r *LogisticsOrdertracePushRequest) SetNextCity(value string) {
	r.SetValue("next_city", value)
}

/* 描述当前节点的操作，操作是“揽收”、“派送”或“签收”。 */
func (r *LogisticsOrdertracePushRequest) SetNodeDescription(value string) {
	r.SetValue("node_description", value)
}

/* 流转节点发生时间 */
func (r *LogisticsOrdertracePushRequest) SetOccureTime(value string) {
	r.SetValue("occure_time", value)
}

/* 流转节点的详细地址及操作描述 */
func (r *LogisticsOrdertracePushRequest) SetOperateDetail(value string) {
	r.SetValue("operate_detail", value)
}

/* 快递业务员联系方式，手机号码或电话。 */
func (r *LogisticsOrdertracePushRequest) SetOperatorContact(value string) {
	r.SetValue("operator_contact", value)
}

/* 快递业务员名称 */
func (r *LogisticsOrdertracePushRequest) SetOperatorName(value string) {
	r.SetValue("operator_name", value)
}


func (r *LogisticsOrdertracePushRequest) GetResponse(accessToken string) (*LogisticsOrdertracePushResponse, []byte, error) {
	var resp LogisticsOrdertracePushResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.ordertrace.push", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsOrdertracePushResponse struct {
	Shipping *Shipping `json:"shipping"`
}

type LogisticsOrdertracePushResponseResult struct {
	Response *LogisticsOrdertracePushResponse `json:"logistics_ordertrace_push_response"`
}





/* 查询物流公司信息（可以查询目的地可不可达情况） */
type LogisticsPartnersGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 货物价格.只有当选择货到付款此参数才会有效 */
func (r *LogisticsPartnersGetRequest) SetGoodsValue(value string) {
	r.SetValue("goods_value", value)
}

/* 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。 */
func (r *LogisticsPartnersGetRequest) SetIsNeedCarriage(value string) {
	r.SetValue("is_need_carriage", value)
}

/* 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。 */
func (r *LogisticsPartnersGetRequest) SetServiceType(value string) {
	r.SetValue("service_type", value)
}

/* 物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20100623_402652267.htm  或者调用 taobao.areas.get 获取 */
func (r *LogisticsPartnersGetRequest) SetSourceId(value string) {
	r.SetValue("source_id", value)
}

/* 物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjbz/xzqhdm/t20100623_402652267.htm  或者调用 taobao.areas.get 获取 */
func (r *LogisticsPartnersGetRequest) SetTargetId(value string) {
	r.SetValue("target_id", value)
}


func (r *LogisticsPartnersGetRequest) GetResponse(accessToken string) (*LogisticsPartnersGetResponse, []byte, error) {
	var resp LogisticsPartnersGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.partners.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsPartnersGetResponse struct {
	LogisticsPartners []*LogisticsPartner `json:"logistics_partners"`
}

type LogisticsPartnersGetResponseResult struct {
	Response *LogisticsPartnersGetResponse `json:"logistics_partners_get_response"`
}





/* 用户根据淘宝交易号查询物流流转信息，如2010-8-10 15：23：00到达杭州集散地。 
此接口的返回信息都由物流公司提供。（备注：使用线下发货（offline.send）的运单，不支持运单状态的实时跟踪，只要一发货，状态就会变为<status>对方已签收</status>，该字段仅对线上发货（online.send）的运单有效。） */
type LogisticsTraceSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 表明是否是拆单，默认值0，1表示拆单 */
func (r *LogisticsTraceSearchRequest) SetIsSplit(value string) {
	r.SetValue("is_split", value)
}

/* 卖家昵称 */
func (r *LogisticsTraceSearchRequest) SetSellerNick(value string) {
	r.SetValue("seller_nick", value)
}

/* 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集 */
func (r *LogisticsTraceSearchRequest) SetSubTid(value string) {
	r.SetValue("sub_tid", value)
}

/* 淘宝交易号，请勿传非淘宝交易号 */
func (r *LogisticsTraceSearchRequest) SetTid(value string) {
	r.SetValue("tid", value)
}


func (r *LogisticsTraceSearchRequest) GetResponse(accessToken string) (*LogisticsTraceSearchResponse, []byte, error) {
	var resp LogisticsTraceSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.logistics.trace.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type LogisticsTraceSearchResponse struct {
	CompanyName string `json:"company_name"`
	OutSid string `json:"out_sid"`
	Status string `json:"status"`
	Tid int `json:"tid"`
	TraceList []*TransitStepInfo `json:"trace_list"`
}

type LogisticsTraceSearchResponseResult struct {
	Response *LogisticsTraceSearchResponse `json:"logistics_trace_search_response"`
}



