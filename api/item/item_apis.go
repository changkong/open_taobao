// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package item

import (
	"github.com/changkong/open_taobao"
)





/* 查询B商家被授权品牌列表、类目列表和 c 商家新品类目列表 */
type ItemcatsAuthorizeGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要返回的字段。目前支持有： 
brand.vid, brand.name,  
item_cat.cid, item_cat.name, item_cat.status,item_cat.sort_order,item_cat.parent_cid,item_cat.is_parent, 
xinpin_item_cat.cid,  
xinpin_item_cat.name,  
xinpin_item_cat.status, 
xinpin_item_cat.sort_order, 
xinpin_item_cat.parent_cid, 
xinpin_item_cat.is_parent */
func (r *ItemcatsAuthorizeGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}


func (r *ItemcatsAuthorizeGetRequest) GetResponse(accessToken string) (*ItemcatsAuthorizeGetResponse, []byte, error) {
	var resp ItemcatsAuthorizeGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.itemcats.authorize.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemcatsAuthorizeGetResponse struct {
	SellerAuthorize *SellerAuthorize `json:"seller_authorize"`
}

type ItemcatsAuthorizeGetResponseResult struct {
	Response *ItemcatsAuthorizeGetResponse `json:"itemcats_authorize_get_response"`
}





/* 获取后台供卖家发布商品的标准商品类目。 
<br/>1. 增量获取类目数据可用<a href="http://api.taobao.com/apidoc/api.htm?path=cid:3-apiId:11072">taobao.itemcats.increment.get</a>接口。 
<br/>2. 全量获取类目数据可用<a href="http://api.taobao.com/apidoc/api.htm?path=cid:3-apiId:11078">taobao.topats.itemcats.get</a>接口。 */
type ItemcatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品所属类目ID列表，用半角逗号(,)分隔 例如:(18957,19562,) (cids、parent_cid至少传一个) */
func (r *ItemcatsGetRequest) SetCids(value string) {
	r.SetValue("cids", value)
}

/* 需要返回的字段列表，见ItemCat，默认返回：cid,parent_cid,name,is_parent */
func (r *ItemcatsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 父商品类目 id，0表示根节点, 传输该参数返回所有子类目。 (cids、parent_cid至少传一个) */
func (r *ItemcatsGetRequest) SetParentCid(value string) {
	r.SetValue("parent_cid", value)
}


func (r *ItemcatsGetRequest) GetResponse(accessToken string) (*ItemcatsGetResponse, []byte, error) {
	var resp ItemcatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.itemcats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItemcatsGetResponse struct {
	ItemCats []*ItemCat `json:"item_cats"`
	LastModified string `json:"last_modified"`
}

type ItemcatsGetResponseResult struct {
	Response *ItemcatsGetResponse `json:"itemcats_get_response"`
}





/* 通过设置必要的参数，来获取商品后台标准类目属性，以及这些属性里面详细的属性值prop_values。 */
type ItempropsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性的Key，支持多条，以“,”分隔 */
func (r *ItempropsGetRequest) SetAttrKeys(value string) {
	r.SetValue("attr_keys", value)
}

/* 类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid */
func (r *ItempropsGetRequest) SetChildPath(value string) {
	r.SetValue("child_path", value)
}

/* 叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID */
func (r *ItempropsGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values */
func (r *ItempropsGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
func (r *ItempropsGetRequest) SetIsColorProp(value string) {
	r.SetValue("is_color_prop", value)
}

/* 是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。 */
func (r *ItempropsGetRequest) SetIsEnumProp(value string) {
	r.SetValue("is_enum_prop", value)
}

/* 在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件) */
func (r *ItempropsGetRequest) SetIsInputProp(value string) {
	r.SetValue("is_input_prop", value)
}

/* 是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否) */
func (r *ItempropsGetRequest) SetIsItemProp(value string) {
	r.SetValue("is_item_prop", value)
}

/* 是否关键属性。可选值:true(是),false(否) */
func (r *ItempropsGetRequest) SetIsKeyProp(value string) {
	r.SetValue("is_key_prop", value)
}

/* 是否销售属性。可选值:true(是),false(否) */
func (r *ItempropsGetRequest) SetIsSaleProp(value string) {
	r.SetValue("is_sale_prop", value)
}

/* 父属性ID */
func (r *ItempropsGetRequest) SetParentPid(value string) {
	r.SetValue("parent_pid", value)
}

/* 属性id (取类目属性时，传pid，不用同时传PID和parent_pid) */
func (r *ItempropsGetRequest) SetPid(value string) {
	r.SetValue("pid", value)
}

/* 获取类目的类型：1代表集市、2代表天猫 */
func (r *ItempropsGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *ItempropsGetRequest) GetResponse(accessToken string) (*ItempropsGetResponse, []byte, error) {
	var resp ItempropsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.itemprops.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItempropsGetResponse struct {
	ItemProps []*ItemProp `json:"item_props"`
	LastModified string `json:"last_modified"`
}

type ItempropsGetResponseResult struct {
	Response *ItempropsGetResponse `json:"itemprops_get_response"`
}





/* 传入类目ID,必需是叶子类目，通过taobao.itemcats.get获取类目ID 
返回字段目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order 
作用:获取标准类目属性值 */
type ItempropvaluesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 属性的Key，支持多条，以“,”分隔 */
func (r *ItempropvaluesGetRequest) SetAttrKeys(value string) {
	r.SetValue("attr_keys", value)
}

/* 叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID */
func (r *ItempropvaluesGetRequest) SetCid(value string) {
	r.SetValue("cid", value)
}

/* 需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order */
func (r *ItempropvaluesGetRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2) */
func (r *ItempropvaluesGetRequest) SetPvs(value string) {
	r.SetValue("pvs", value)
}

/* 获取类目的类型：1代表集市、2代表天猫 */
func (r *ItempropvaluesGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *ItempropvaluesGetRequest) GetResponse(accessToken string) (*ItempropvaluesGetResponse, []byte, error) {
	var resp ItempropvaluesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.itempropvalues.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type ItempropvaluesGetResponse struct {
	LastModified string `json:"last_modified"`
	PropValues []*PropValue `json:"prop_values"`
}

type ItempropvaluesGetResponseResult struct {
	Response *ItempropvaluesGetResponse `json:"itempropvalues_get_response"`
}





/* 全量获取后台类目、类目属性、类目属性值数据 
<br/>异步API使用方法，请查看：<a href="http://open.taobao.com/doc/detail.htm?id=30">异步API使用说明</a> 
<br/>1. 每天8点左右会产生今天的全量数据，在8点之前提交获取全量类目数据任务需要等到8点之后才完成，在8点之后提交获取全量类目数据任务可以接近实时返回。 
<br/>2. 提交任务后，通过taobao.topats.result.get来查看任务执行状态，如果任务已完成，则返回下载URL 
<br/>3. 如果订阅了主动通知服务，任务完成后TOP会通过HTTP长连接推送消息，通知的消息格式请参考异步API使用文档 
<br/>4. 下载到的结果是zip压缩包，解压后得到一个标准的csv/json格式的文本文件，文件内容的默认编码格式是UTF-8 */
type TopatsItemcatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 一级类目ID列表（非一级类目将会被忽略），用半角分号(;)分隔，例如:"16;19562"，一次最多可以获取10个类目的增量数据。<span style="color:red">注：传入0代表获取所有类目的数据,默认获取所有类目数据</span> */
func (r *TopatsItemcatsGetRequest) SetCids(value string) {
	r.SetValue("cids", value)
}

/* 类目数据输出格式，可选值为：csv, json（默认csv格式返回） */
func (r *TopatsItemcatsGetRequest) SetOutputFormat(value string) {
	r.SetValue("output_format", value)
}

/* 获取类目的类型：1代表集市、2代表天猫 */
func (r *TopatsItemcatsGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *TopatsItemcatsGetRequest) GetResponse(accessToken string) (*TopatsItemcatsGetResponse, []byte, error) {
	var resp TopatsItemcatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.itemcats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsItemcatsGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsItemcatsGetResponseResult struct {
	Response *TopatsItemcatsGetResponse `json:"topats_itemcats_get_response"`
}



