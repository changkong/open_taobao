// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package wlb

import (
	"github.com/yaofangou/open_taobao"
)





/* 查询库存明细 */
type InventoryIpcInventorydetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主订单号，可选 */
func (r *InventoryIpcInventorydetailGetRequest) SetBizOrderId(value string) {
	r.SetValue("biz_order_id", value)
}

/* 子订单号，可选 */
func (r *InventoryIpcInventorydetailGetRequest) SetBizSubOrderId(value string) {
	r.SetValue("biz_sub_order_id", value)
}

/* 当前页数 */
func (r *InventoryIpcInventorydetailGetRequest) SetPageIndex(value string) {
	r.SetValue("page_index", value)
}

/* 一页显示多少条 */
func (r *InventoryIpcInventorydetailGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓储商品id */
func (r *InventoryIpcInventorydetailGetRequest) SetScItemId(value string) {
	r.SetValue("sc_item_id", value)
}

/* 1:查询预扣  4：查询占用 */
func (r *InventoryIpcInventorydetailGetRequest) SetStatusQuery(value string) {
	r.SetValue("status_query", value)
}


func (r *InventoryIpcInventorydetailGetRequest) GetResponse(accessToken string) (*InventoryIpcInventorydetailGetResponse, []byte, error) {
	var resp InventoryIpcInventorydetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.inventory.ipc.inventorydetail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type InventoryIpcInventorydetailGetResponse struct {
	InventoryDetails []*IpcInventoryDetailDo `json:"inventory_details"`
}

type InventoryIpcInventorydetailGetResponseResult struct {
	Response *InventoryIpcInventorydetailGetResponse `json:"inventory_ipc_inventorydetail_get_response"`
}





/* 查询库存详情，通过商品ID获取发送请求的卖家的库存详情 */
type WlbInventoryDetailGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 库存类型列表，值包括： 
VENDIBLE--可销售库存 
FREEZE--冻结库存 
ONWAY--在途库存 
DEFECT--残次品库存 
ENGINE_DAMAGE--机损 
BOX_DAMAGE--箱损 
EXPIRATION--过保 */
func (r *WlbInventoryDetailGetRequest) SetInventoryTypeList(value string) {
	r.SetValue("inventory_type_list", value)
}

/* 商品ID */
func (r *WlbInventoryDetailGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 仓库编码 */
func (r *WlbInventoryDetailGetRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *WlbInventoryDetailGetRequest) GetResponse(accessToken string) (*WlbInventoryDetailGetResponse, []byte, error) {
	var resp WlbInventoryDetailGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.inventory.detail.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbInventoryDetailGetResponse struct {
	InventoryList []*WlbInventory `json:"inventory_list"`
	ItemId int `json:"item_id"`
}

type WlbInventoryDetailGetResponseResult struct {
	Response *WlbInventoryDetailGetResponse `json:"wlb_inventory_detail_get_response"`
}





/* 将库存同步至IC */
type WlbInventorySyncRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品ID */
func (r *WlbInventorySyncRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 外部实体类型.存如下值  
IC_ITEM --表示IC商品;  
IC_SKU --表示IC最小单位商品; 
WLB_ITEM  --表示WLB商品. 
若值不在范围内，则按WLB_ITEM处理 */
func (r *WlbInventorySyncRequest) SetItemType(value string) {
	r.SetValue("item_type", value)
}

/* 库存数量 */
func (r *WlbInventorySyncRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}


func (r *WlbInventorySyncRequest) GetResponse(accessToken string) (*WlbInventorySyncResponse, []byte, error) {
	var resp WlbInventorySyncResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.inventory.sync", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbInventorySyncResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbInventorySyncResponseResult struct {
	Response *WlbInventorySyncResponse `json:"wlb_inventory_sync_response"`
}





/* 通过商品ID等几个条件来分页查询库存变更记录 */
type WlbInventorylogQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束修改时间,小于等于该时间 */
func (r *WlbInventorylogQueryRequest) SetGmtEnd(value string) {
	r.SetValue("gmt_end", value)
}

/* 起始修改时间,大于等于该时间 */
func (r *WlbInventorylogQueryRequest) SetGmtStart(value string) {
	r.SetValue("gmt_start", value)
}

/* 商品ID */
func (r *WlbInventorylogQueryRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 库存操作作类型(可以为空) 
CHU_KU 1-出库 
RU_KU 2-入库 
FREEZE 3-冻结 
THAW 4-解冻 
CHECK_FREEZE 5-冻结确认 
CHANGE_KU 6-库存类型变更 
若值不在范围内，则按CHU_KU处理 */
func (r *WlbInventorylogQueryRequest) SetOpType(value string) {
	r.SetValue("op_type", value)
}

/* 可指定授权的用户来查询 */
func (r *WlbInventorylogQueryRequest) SetOpUserId(value string) {
	r.SetValue("op_user_id", value)
}

/* 单号 */
func (r *WlbInventorylogQueryRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}

/* 当前页 */
func (r *WlbInventorylogQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数 */
func (r *WlbInventorylogQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓库编码 */
func (r *WlbInventorylogQueryRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *WlbInventorylogQueryRequest) GetResponse(accessToken string) (*WlbInventorylogQueryResponse, []byte, error) {
	var resp WlbInventorylogQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.inventorylog.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbInventorylogQueryResponse struct {
	InventoryLogList []*WlbItemInventoryLog `json:"inventory_log_list"`
	TotalCount int `json:"total_count"`
}

type WlbInventorylogQueryResponseResult struct {
	Response *WlbInventorylogQueryResponse `json:"wlb_inventorylog_query_response"`
}





/* 添加物流宝商品，支持物流宝子商品和属性添加 */
type WlbItemAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品颜色 */
func (r *WlbItemAddRequest) SetColor(value string) {
	r.SetValue("color", value)
}

/* 货类 */
func (r *WlbItemAddRequest) SetGoodsCat(value string) {
	r.SetValue("goods_cat", value)
}

/* 商品高度，单位毫米 */
func (r *WlbItemAddRequest) SetHeight(value string) {
	r.SetValue("height", value)
}

/* 是否危险品 */
func (r *WlbItemAddRequest) SetIsDangerous(value string) {
	r.SetValue("is_dangerous", value)
}

/* 是否易碎品 */
func (r *WlbItemAddRequest) SetIsFriable(value string) {
	r.SetValue("is_friable", value)
}

/* 是否sku */
func (r *WlbItemAddRequest) SetIsSku(value string) {
	r.SetValue("is_sku", value)
}

/* 商品编码 */
func (r *WlbItemAddRequest) SetItemCode(value string) {
	r.SetValue("item_code", value)
}

/* 商品长度，单位毫米 */
func (r *WlbItemAddRequest) SetLength(value string) {
	r.SetValue("length", value)
}

/* 商品名称 */
func (r *WlbItemAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商品包装材料类型 */
func (r *WlbItemAddRequest) SetPackageMaterial(value string) {
	r.SetValue("package_material", value)
}

/* 商品价格，单位：分 */
func (r *WlbItemAddRequest) SetPrice(value string) {
	r.SetValue("price", value)
}

/* 计价货类 */
func (r *WlbItemAddRequest) SetPricingCat(value string) {
	r.SetValue("pricing_cat", value)
}

/* 属性名列表,目前支持的属性： 
毛重:GWeight	 
净重:Nweight 
皮重: Tweight 
自定义属性： 
paramkey1 
paramkey2 
paramkey3 
paramkey4 */
func (r *WlbItemAddRequest) SetProNameList(value string) {
	r.SetValue("pro_name_list", value)
}

/* 属性值列表： 
10,8 */
func (r *WlbItemAddRequest) SetProValueList(value string) {
	r.SetValue("pro_value_list", value)
}

/* 商品备注 */
func (r *WlbItemAddRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* 是否支持批次 */
func (r *WlbItemAddRequest) SetSupportBatch(value string) {
	r.SetValue("support_batch", value)
}

/* 商品标题 */
func (r *WlbItemAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* NORMAL--普通商品 
COMBINE--组合商品 
DISTRIBUTION--分销 */
func (r *WlbItemAddRequest) SetType(value string) {
	r.SetValue("type", value)
}

/* 商品体积，单位立方厘米 */
func (r *WlbItemAddRequest) SetVolume(value string) {
	r.SetValue("volume", value)
}

/* 商品重量，单位G */
func (r *WlbItemAddRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}

/* 商品宽度，单位毫米 */
func (r *WlbItemAddRequest) SetWidth(value string) {
	r.SetValue("width", value)
}


func (r *WlbItemAddRequest) GetResponse(accessToken string) (*WlbItemAddResponse, []byte, error) {
	var resp WlbItemAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemAddResponse struct {
	ItemId int `json:"item_id"`
}

type WlbItemAddResponseResult struct {
	Response *WlbItemAddResponse `json:"wlb_item_add_response"`
}





/* 添加商品的授权规则：添加规则之后代销商可以增加商品代销关系 */
type WlbItemAuthorizationAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 授权类型：1=全量授权，0=部分授权 
当部分授权时，需要指定授权数量quantity */
func (r *WlbItemAuthorizationAddRequest) SetAuthType(value string) {
	r.SetValue("auth_type", value)
}

/* 授权结束时间 */
func (r *WlbItemAuthorizationAddRequest) SetAuthorizeEndTime(value string) {
	r.SetValue("authorize_end_time", value)
}

/* 授权开始时间 */
func (r *WlbItemAuthorizationAddRequest) SetAuthorizeStartTime(value string) {
	r.SetValue("authorize_start_time", value)
}

/* 被授权人用户id */
func (r *WlbItemAuthorizationAddRequest) SetConsignUserNick(value string) {
	r.SetValue("consign_user_nick", value)
}

/* 商品id列表，以英文逗号,分隔，最多可放入50个商品ID。 */
func (r *WlbItemAuthorizationAddRequest) SetItemIdList(value string) {
	r.SetValue("item_id_list", value)
}

/* 规则名称 */
func (r *WlbItemAuthorizationAddRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 授权数量 */
func (r *WlbItemAuthorizationAddRequest) SetQuantity(value string) {
	r.SetValue("quantity", value)
}

/* 授权规则：目前只有GRANT_FIX，按照数量授权 */
func (r *WlbItemAuthorizationAddRequest) SetRuleCode(value string) {
	r.SetValue("rule_code", value)
}


func (r *WlbItemAuthorizationAddRequest) GetResponse(accessToken string) (*WlbItemAuthorizationAddResponse, []byte, error) {
	var resp WlbItemAuthorizationAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.authorization.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemAuthorizationAddResponse struct {
	RuleIdList []int `json:"rule_id_list"`
}

type WlbItemAuthorizationAddResponseResult struct {
	Response *WlbItemAuthorizationAddResponse `json:"wlb_item_authorization_add_response"`
}





/* 删除授权关系.若有建立代销关系，会将其代销关系冻结即将状态置为失效(代销关系)。 */
type WlbItemAuthorizationDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 授权关系ID */
func (r *WlbItemAuthorizationDeleteRequest) SetAuthorizeId(value string) {
	r.SetValue("authorize_id", value)
}


func (r *WlbItemAuthorizationDeleteRequest) GetResponse(accessToken string) (*WlbItemAuthorizationDeleteResponse, []byte, error) {
	var resp WlbItemAuthorizationDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.authorization.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemAuthorizationDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemAuthorizationDeleteResponseResult struct {
	Response *WlbItemAuthorizationDeleteResponse `json:"wlb_item_authorization_delete_response"`
}





/* 查询授权关系,可由货主或被授权人查询。 */
type WlbItemAuthorizationQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 授权商品ID */
func (r *WlbItemAuthorizationQueryRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 授权名称 */
func (r *WlbItemAuthorizationQueryRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 当前页 */
func (r *WlbItemAuthorizationQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *WlbItemAuthorizationQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 授权编码 */
func (r *WlbItemAuthorizationQueryRequest) SetRuleCode(value string) {
	r.SetValue("rule_code", value)
}

/* 状态： 只能输入如下值,范围外的默认按VALID处理;不选则查询所有;  
VALID -- 1 有效； INVALIDATION -- 2 失效 */
func (r *WlbItemAuthorizationQueryRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 类型：可由不同角色来查询，默认值OWNER, 
OWNER -- 授权人, 
ON_COMMISSION -- 被授权人 */
func (r *WlbItemAuthorizationQueryRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *WlbItemAuthorizationQueryRequest) GetResponse(accessToken string) (*WlbItemAuthorizationQueryResponse, []byte, error) {
	var resp WlbItemAuthorizationQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.authorization.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemAuthorizationQueryResponse struct {
	AuthorizationList []*WlbAuthorization `json:"authorization_list"`
	TotalCount int `json:"total_count"`
}

type WlbItemAuthorizationQueryResponseResult struct {
	Response *WlbItemAuthorizationQueryResponse `json:"wlb_item_authorization_query_response"`
}





/* 根据用户id，item id list和store code来查询商品库存信息和批次信息 */
type WlbItemBatchQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 需要查询的商品ID列表，以字符串表示，ID间以;隔开 */
func (r *WlbItemBatchQueryRequest) SetItemIds(value string) {
	r.SetValue("item_ids", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (r *WlbItemBatchQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (r *WlbItemBatchQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓库编号 */
func (r *WlbItemBatchQueryRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *WlbItemBatchQueryRequest) GetResponse(accessToken string) (*WlbItemBatchQueryResponse, []byte, error) {
	var resp WlbItemBatchQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.batch.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemBatchQueryResponse struct {
	ItemInventoryBatchList []*WlbItemBatchInventory `json:"item_inventory_batch_list"`
	TotalCount int `json:"total_count"`
}

type WlbItemBatchQueryResponseResult struct {
	Response *WlbItemBatchQueryResponse `json:"wlb_item_batch_query_response"`
}





/* 创建商品组合关系 */
type WlbItemCombinationCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 组合商品的id列表 */
func (r *WlbItemCombinationCreateRequest) SetDestItemList(value string) {
	r.SetValue("dest_item_list", value)
}

/* 要建立组合关系的商品id */
func (r *WlbItemCombinationCreateRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 组成组合商品的比例列表，描述组合商品的组合比例，默认为1,1,1 */
func (r *WlbItemCombinationCreateRequest) SetProportionList(value string) {
	r.SetValue("proportion_list", value)
}


func (r *WlbItemCombinationCreateRequest) GetResponse(accessToken string) (*WlbItemCombinationCreateResponse, []byte, error) {
	var resp WlbItemCombinationCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.combination.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemCombinationCreateResponse struct {
	GmtCreate string `json:"gmt_create"`
}

type WlbItemCombinationCreateResponseResult struct {
	Response *WlbItemCombinationCreateResponse `json:"wlb_item_combination_create_response"`
}





/* 删除商品组合关系 */
type WlbItemCombinationDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 组合商品的id列表 */
func (r *WlbItemCombinationDeleteRequest) SetDestItemList(value string) {
	r.SetValue("dest_item_list", value)
}

/* 组合关系的商品id */
func (r *WlbItemCombinationDeleteRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *WlbItemCombinationDeleteRequest) GetResponse(accessToken string) (*WlbItemCombinationDeleteResponse, []byte, error) {
	var resp WlbItemCombinationDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.combination.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemCombinationDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemCombinationDeleteResponseResult struct {
	Response *WlbItemCombinationDeleteResponse `json:"wlb_item_combination_delete_response"`
}





/* 根据商品id查询商品组合关系 */
type WlbItemCombinationGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要查询的组合商品id */
func (r *WlbItemCombinationGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *WlbItemCombinationGetRequest) GetResponse(accessToken string) (*WlbItemCombinationGetResponse, []byte, error) {
	var resp WlbItemCombinationGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.combination.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemCombinationGetResponse struct {
	ItemIdList []int `json:"item_id_list"`
}

type WlbItemCombinationGetResponseResult struct {
	Response *WlbItemCombinationGetResponse `json:"wlb_item_combination_get_response"`
}





/* 代销商创建商品代销关系：货主创建授权规则，获得授权规则后方可创建商品代销关系。 */
type WlbItemConsignmentCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品id */
func (r *WlbItemConsignmentCreateRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 代销数量 */
func (r *WlbItemConsignmentCreateRequest) SetNumber(value string) {
	r.SetValue("number", value)
}

/* 货主商品id */
func (r *WlbItemConsignmentCreateRequest) SetOwnerItemId(value string) {
	r.SetValue("owner_item_id", value)
}

/* 货主id */
func (r *WlbItemConsignmentCreateRequest) SetOwnerUserId(value string) {
	r.SetValue("owner_user_id", value)
}

/* 通过taobao.wlb.item.authorization.add接口创建后得到的rule_id，规则中设定了代销商可以代销的商品数量 */
func (r *WlbItemConsignmentCreateRequest) SetRuleId(value string) {
	r.SetValue("rule_id", value)
}


func (r *WlbItemConsignmentCreateRequest) GetResponse(accessToken string) (*WlbItemConsignmentCreateResponse, []byte, error) {
	var resp WlbItemConsignmentCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.consignment.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemConsignmentCreateResponse struct {
	ConsignmentId int `json:"consignment_id"`
}

type WlbItemConsignmentCreateResponseResult struct {
	Response *WlbItemConsignmentCreateResponse `json:"wlb_item_consignment_create_response"`
}





/* 删除商品的代销关系 */
type WlbItemConsignmentDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销商前台宝贝ID */
func (r *WlbItemConsignmentDeleteRequest) SetIcItemId(value string) {
	r.SetValue("ic_item_id", value)
}

/* 货主的物流宝商品ID */
func (r *WlbItemConsignmentDeleteRequest) SetOwnerItemId(value string) {
	r.SetValue("owner_item_id", value)
}

/* 授权关系id */
func (r *WlbItemConsignmentDeleteRequest) SetRuleId(value string) {
	r.SetValue("rule_id", value)
}


func (r *WlbItemConsignmentDeleteRequest) GetResponse(accessToken string) (*WlbItemConsignmentDeleteResponse, []byte, error) {
	var resp WlbItemConsignmentDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.consignment.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemConsignmentDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemConsignmentDeleteResponseResult struct {
	Response *WlbItemConsignmentDeleteResponse `json:"wlb_item_consignment_delete_response"`
}





/* 代销商角度分页查询物流宝商品的代销关系 */
type WlbItemConsignmentPageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 代销商宝贝id */
func (r *WlbItemConsignmentPageGetRequest) SetIcItemId(value string) {
	r.SetValue("ic_item_id", value)
}

/* 供应商商品id */
func (r *WlbItemConsignmentPageGetRequest) SetOwnerItemId(value string) {
	r.SetValue("owner_item_id", value)
}

/* 供应商用户昵称 */
func (r *WlbItemConsignmentPageGetRequest) SetOwnerUserNick(value string) {
	r.SetValue("owner_user_nick", value)
}


func (r *WlbItemConsignmentPageGetRequest) GetResponse(accessToken string) (*WlbItemConsignmentPageGetResponse, []byte, error) {
	var resp WlbItemConsignmentPageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.consignment.page.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemConsignmentPageGetResponse struct {
	TotalCount int `json:"total_count"`
	WlbConsignMents []*WlbConsignMent `json:"wlb_consign_ments"`
}

type WlbItemConsignmentPageGetResponseResult struct {
	Response *WlbItemConsignmentPageGetResponse `json:"wlb_item_consignment_page_get_response"`
}





/* 通过ItemId,UserId来删除单个商品 */
type WlbItemDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品ID */
func (r *WlbItemDeleteRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 商品所有人淘宝nick */
func (r *WlbItemDeleteRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *WlbItemDeleteRequest) GetResponse(accessToken string) (*WlbItemDeleteResponse, []byte, error) {
	var resp WlbItemDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemDeleteResponseResult struct {
	Response *WlbItemDeleteResponse `json:"wlb_item_delete_response"`
}





/* 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。 */
type WlbItemGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品ID */
func (r *WlbItemGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *WlbItemGetRequest) GetResponse(accessToken string) (*WlbItemGetResponse, []byte, error) {
	var resp WlbItemGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemGetResponse struct {
	Item *WlbItem `json:"item"`
}

type WlbItemGetResponseResult struct {
	Response *WlbItemGetResponse `json:"wlb_item_get_response"`
}





/* 根据物流宝商品ID查询商品映射关系 */
type WlbItemMapGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要查询映射关系的物流宝商品id */
func (r *WlbItemMapGetRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *WlbItemMapGetRequest) GetResponse(accessToken string) (*WlbItemMapGetResponse, []byte, error) {
	var resp WlbItemMapGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.map.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemMapGetResponse struct {
	OutEntityItemList []*OutEntityItem `json:"out_entity_item_list"`
}

type WlbItemMapGetResponseResult struct {
	Response *WlbItemMapGetResponse `json:"wlb_item_map_get_response"`
}





/* 根据外部实体类型和实体id查询映射的物流宝商品id */
type WlbItemMapGetByExtentityRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 外部实体类型对应的商品id */
func (r *WlbItemMapGetByExtentityRequest) SetExtEntityId(value string) {
	r.SetValue("ext_entity_id", value)
}

/* 外部实体类型： IC_ITEM--ic商品 IC_SKU--ic销售单元 */
func (r *WlbItemMapGetByExtentityRequest) SetExtEntityType(value string) {
	r.SetValue("ext_entity_type", value)
}


func (r *WlbItemMapGetByExtentityRequest) GetResponse(accessToken string) (*WlbItemMapGetByExtentityResponse, []byte, error) {
	var resp WlbItemMapGetByExtentityResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.map.get.by.extentity", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemMapGetByExtentityResponse struct {
	ItemId int `json:"item_id"`
}

type WlbItemMapGetByExtentityResponseResult struct {
	Response *WlbItemMapGetByExtentityResponse `json:"wlb_item_map_get_by_extentity_response"`
}





/* 根据状态、卖家、SKU等信息查询商品列表 */
type WlbItemQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示; 
若值不在范围内，则按true处理 */
func (r *WlbItemQueryRequest) SetIsSku(value string) {
	r.SetValue("is_sku", value)
}

/* 商家编码 */
func (r *WlbItemQueryRequest) SetItemCode(value string) {
	r.SetValue("item_code", value)
}

/* ITEM类型(只允许输入以下英文或空) 
NORMAL  0:普通商品;  
COMBINE  1:是否是组合商品  
DISTRIBUTION  2:是否是分销商品(货主是别人) 
若值不在范围内，则按NORMAL处理 */
func (r *WlbItemQueryRequest) SetItemType(value string) {
	r.SetValue("item_type", value)
}

/* 商品名称 */
func (r *WlbItemQueryRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 当前页 */
func (r *WlbItemQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *WlbItemQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品 */
func (r *WlbItemQueryRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}

/* 只能输入以下值或空： 
ITEM_STATUS_VALID -- 1 表示 有效； 
ITEM_STATUS_LOCK  -- 2 表示锁住。 
若值不在范围内，按ITEM_STATUS_VALID处理 */
func (r *WlbItemQueryRequest) SetStatus(value string) {
	r.SetValue("status", value)
}

/* 商品前台销售名字 */
func (r *WlbItemQueryRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *WlbItemQueryRequest) GetResponse(accessToken string) (*WlbItemQueryResponse, []byte, error) {
	var resp WlbItemQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemQueryResponse struct {
	ItemList []*WlbItem `json:"item_list"`
	TotalCount int `json:"total_count"`
}

type WlbItemQueryResponseResult struct {
	Response *WlbItemQueryResponse `json:"wlb_item_query_response"`
}





/* 同步仓储商品与前台商品，建立仓储商品与前台商品的关系,并更新IC中的信息到仓储商品Item中 */
type WlbItemSynchronizeRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 外部实体ID */
func (r *WlbItemSynchronizeRequest) SetExtEntityId(value string) {
	r.SetValue("ext_entity_id", value)
}

/* 外部实体类型.存如下值 
IC_ITEM   --表示IC商品 
IC_SKU    --表示IC最小单位商品 
若输入其他值，则按IC_ITEM处理 */
func (r *WlbItemSynchronizeRequest) SetExtEntityType(value string) {
	r.SetValue("ext_entity_type", value)
}

/* 商品ID */
func (r *WlbItemSynchronizeRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 商品所有人淘宝nick */
func (r *WlbItemSynchronizeRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *WlbItemSynchronizeRequest) GetResponse(accessToken string) (*WlbItemSynchronizeResponse, []byte, error) {
	var resp WlbItemSynchronizeResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.synchronize", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemSynchronizeResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemSynchronizeResponseResult struct {
	Response *WlbItemSynchronizeResponse `json:"wlb_item_synchronize_response"`
}





/* 通过物流宝商品ID和IC商品ID删除映射关系。 */
type WlbItemSynchronizeDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 外部实体ID */
func (r *WlbItemSynchronizeDeleteRequest) SetExtEntityId(value string) {
	r.SetValue("ext_entity_id", value)
}

/* 外部实体类型.存如下值 IC_ITEM --表示IC商品; IC_SKU --表示IC最小单位商品;若输入其他值，则按IC_ITEM处理 */
func (r *WlbItemSynchronizeDeleteRequest) SetExtEntityType(value string) {
	r.SetValue("ext_entity_type", value)
}

/* 物流宝商品ID */
func (r *WlbItemSynchronizeDeleteRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}


func (r *WlbItemSynchronizeDeleteRequest) GetResponse(accessToken string) (*WlbItemSynchronizeDeleteResponse, []byte, error) {
	var resp WlbItemSynchronizeDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.synchronize.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemSynchronizeDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemSynchronizeDeleteResponseResult struct {
	Response *WlbItemSynchronizeDeleteResponse `json:"wlb_item_synchronize_delete_response"`
}





/* 修改物流宝商品信息 */
type WlbItemUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品颜色 */
func (r *WlbItemUpdateRequest) SetColor(value string) {
	r.SetValue("color", value)
}

/* 需要删除的商品属性key列表 */
func (r *WlbItemUpdateRequest) SetDeletePropertyKeyList(value string) {
	r.SetValue("delete_property_key_list", value)
}

/* 商品货类 */
func (r *WlbItemUpdateRequest) SetGoodsCat(value string) {
	r.SetValue("goods_cat", value)
}

/* 商品高度，单位厘米 */
func (r *WlbItemUpdateRequest) SetHeight(value string) {
	r.SetValue("height", value)
}

/* 要修改的商品id */
func (r *WlbItemUpdateRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 是否危险品 */
func (r *WlbItemUpdateRequest) SetIsDangerous(value string) {
	r.SetValue("is_dangerous", value)
}

/* 是否易碎品 */
func (r *WlbItemUpdateRequest) SetIsFriable(value string) {
	r.SetValue("is_friable", value)
}

/* 商品长度，单位厘米 */
func (r *WlbItemUpdateRequest) SetLength(value string) {
	r.SetValue("length", value)
}

/* 要修改的商品名称 */
func (r *WlbItemUpdateRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 商品包装材料类型 */
func (r *WlbItemUpdateRequest) SetPackageMaterial(value string) {
	r.SetValue("package_material", value)
}

/* 商品计价货类 */
func (r *WlbItemUpdateRequest) SetPricingCat(value string) {
	r.SetValue("pricing_cat", value)
}

/* 要修改的商品备注 */
func (r *WlbItemUpdateRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* 要修改的商品标题 */
func (r *WlbItemUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 需要修改的商品属性值的列表，如果属性不存在，则新增属性 */
func (r *WlbItemUpdateRequest) SetUpdatePropertyKeyList(value string) {
	r.SetValue("update_property_key_list", value)
}

/* 需要修改的属性值的列表 */
func (r *WlbItemUpdateRequest) SetUpdatePropertyValueList(value string) {
	r.SetValue("update_property_value_list", value)
}

/* 商品体积，单位立方厘米 */
func (r *WlbItemUpdateRequest) SetVolume(value string) {
	r.SetValue("volume", value)
}

/* 商品重量，单位G */
func (r *WlbItemUpdateRequest) SetWeight(value string) {
	r.SetValue("weight", value)
}

/* 商品宽度，单位厘米 */
func (r *WlbItemUpdateRequest) SetWidth(value string) {
	r.SetValue("width", value)
}


func (r *WlbItemUpdateRequest) GetResponse(accessToken string) (*WlbItemUpdateResponse, []byte, error) {
	var resp WlbItemUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.item.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbItemUpdateResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbItemUpdateResponseResult struct {
	Response *WlbItemUpdateResponse `json:"wlb_item_update_response"`
}





/* 确认物流宝可执行消息 */
type WlbNotifyMessageConfirmRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝通知消息的id，通过taobao.wlb.notify.message.page.get接口得到的WlbMessage数据结构中的id字段 */
func (r *WlbNotifyMessageConfirmRequest) SetMessageId(value string) {
	r.SetValue("message_id", value)
}


func (r *WlbNotifyMessageConfirmRequest) GetResponse(accessToken string) (*WlbNotifyMessageConfirmResponse, []byte, error) {
	var resp WlbNotifyMessageConfirmResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.notify.message.confirm", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbNotifyMessageConfirmResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbNotifyMessageConfirmResponseResult struct {
	Response *WlbNotifyMessageConfirmResponse `json:"wlb_notify_message_confirm_response"`
}





/* 物流宝提供的消息通知查询接口，消息内容包括;出入库单不一致消息，取消订单成功消息，盘点单消息 */
type WlbNotifyMessagePageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询条件：记录截至时间 */
func (r *WlbNotifyMessagePageGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 通知消息编码： 
STOCK_IN_NOT_CONSISTENT---入库单不一致 
CANCEL_ORDER_SUCCESS---取消订单成功 
INVENTORY_CHECK---盘点 
CANCEL_ORDER_FAILURE---取消订单失败 
ORDER_REJECT--wms拒单 
ORDER_CONFIRMED--订单处理成功 */
func (r *WlbNotifyMessagePageGetRequest) SetMsgCode(value string) {
	r.SetValue("msg_code", value)
}

/* 分页查询页数 */
func (r *WlbNotifyMessagePageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页查询的每页页数 */
func (r *WlbNotifyMessagePageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询条件：记录开始时间 */
func (r *WlbNotifyMessagePageGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}

/* 消息状态： 
不需要确认：NO_NEED_CONFIRM 
已确认：CONFIRMED 
待确认：TO_BE_CONFIRM */
func (r *WlbNotifyMessagePageGetRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *WlbNotifyMessagePageGetRequest) GetResponse(accessToken string) (*WlbNotifyMessagePageGetResponse, []byte, error) {
	var resp WlbNotifyMessagePageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.notify.message.page.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbNotifyMessagePageGetResponse struct {
	TotalCount int `json:"total_count"`
	WlbMessages []*WlbMessage `json:"wlb_messages"`
}

type WlbNotifyMessagePageGetResponseResult struct {
	Response *WlbNotifyMessagePageGetResponse `json:"wlb_notify_message_page_get_response"`
}





/* 取消物流宝订单 */
type WlbOrderCancelRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝订单编号 */
func (r *WlbOrderCancelRequest) SetWlbOrderCode(value string) {
	r.SetValue("wlb_order_code", value)
}


func (r *WlbOrderCancelRequest) GetResponse(accessToken string) (*WlbOrderCancelResponse, []byte, error) {
	var resp WlbOrderCancelResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.cancel", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderCancelResponse struct {
	ErrorCodeList string `json:"error_code_list"`
	ModifyTime string `json:"modify_time"`
}

type WlbOrderCancelResponseResult struct {
	Response *WlbOrderCancelResponse `json:"wlb_order_cancel_response"`
}





/* 如果erp导入淘宝交易订单到物流宝，当物流宝订单已发货的时候，erp需要调用该接口来通知物流订单和淘宝交易订单已发货 */
type WlbOrderConsignRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝订单编号 */
func (r *WlbOrderConsignRequest) SetWlbOrderCode(value string) {
	r.SetValue("wlb_order_code", value)
}


func (r *WlbOrderConsignRequest) GetResponse(accessToken string) (*WlbOrderConsignResponse, []byte, error) {
	var resp WlbOrderConsignResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.consign", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderConsignResponse struct {
	ModifyTime string `json:"modify_time"`
}

type WlbOrderConsignResponseResult struct {
	Response *WlbOrderConsignResponse `json:"wlb_order_consign_response"`
}





/* 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生 */
type WlbOrderCreateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 支付宝交易号 */
func (r *WlbOrderCreateRequest) SetAlipayNo(value string) {
	r.SetValue("alipay_no", value)
}

/* 该字段暂时保留 */
func (r *WlbOrderCreateRequest) SetAttributes(value string) {
	r.SetValue("attributes", value)
}

/* 买家呢称 */
func (r *WlbOrderCreateRequest) SetBuyerNick(value string) {
	r.SetValue("buyer_nick", value)
}

/* 期望结束时间，在入库单会使用到 */
func (r *WlbOrderCreateRequest) SetExpectEndTime(value string) {
	r.SetValue("expect_end_time", value)
}

/* 计划开始送达时间  在入库单中可能会使用 */
func (r *WlbOrderCreateRequest) SetExpectStartTime(value string) {
	r.SetValue("expect_start_time", value)
}

/* {"invoince_info": [{"bill_type":"发票类型，必选", "bill_title":"发票抬头，必选", "bill_amount":"发票金额(单位是分)，必选","bill_content":"发票内容，可选"}]} */
func (r *WlbOrderCreateRequest) SetInvoinceInfo(value string) {
	r.SetValue("invoince_info", value)
}

/* 该物流宝订单是否已完成，如果完成则设置为true，如果为false，则需要等待继续创建订单商品信息。 */
func (r *WlbOrderCreateRequest) SetIsFinished(value string) {
	r.SetValue("is_finished", value)
}

/* 物流宝订单编号，该接口约定每次最多只能传50条order_item_list，如果一个物流宝订单超过50条商品的时候，需要批量来调用该接口，第一次调用的时候，wlb_order_code为空，如果第一次创建成功，该接口返回wlb_order_code，其后继续再该订单上添加商品条目，需要带上wlb_order_code，out_biz_code，order_item_list,is_finished四个字段。 */
func (r *WlbOrderCreateRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}

/* 用字符串格式来表示订单标记列表：比如COD^PRESELL^SPLIT^LIMIT 等，中间用“^”来隔开 ---------------------------------------- 订单标记list（所有字母全部大写）： 1: COD –货到付款 2: LIMIT-限时配送 3: PRESELL-预售 5:COMPLAIN-已投诉 7:SPLIT-拆单， 8:EXCHANGE-换货， 9:VISIT-上门 ， 10: MODIFYTRANSPORT-是否可改配送方式， 
: 是否可改配送方式  默认可更改 
11 CONSIGN 物流宝代理发货,自动更改发货状态 
12: SELLER_AFFORD 是否卖家承担运费 默认是，即没 13: SYNC_RETURN_BILL，同时退回发票 */
func (r *WlbOrderCreateRequest) SetOrderFlag(value string) {
	r.SetValue("order_flag", value)
}

/* 订单商品列表： {"order_item_list":[{"trade_code":"可选,淘宝交易订单，并且不是赠品，必须要传订单来源编号"," sub_trade_code ":"可选,淘宝子交易号","item_id":"必须,商品Id","item_code":"必须,商家编码","item_name":"可选,物流宝商品名称","item_quantity":"必选,计划数量","item_price":"必选,物品价格,单位为分","owner_user_nick
":"可选,货主nick 代销模式下会存在","flag":"判断是否为赠品0 不是1是","remarks":"可选,备注","batch_remark":"可选，批次描述信息会把这个信息带给WMS，但不会跟物流宝库存相关联"，"inventory_type":"库存类型1 可销售库存 101 类型用来定义残次品 201 冻结类型库存 301 在途库存","picture_url":"图片Url","distributor_user_nick": "分销商NICK",必选"ext_order_item_code":"可选，外部商品的商家编码"]} ======================================== 如果订单中的商品条目数大于50条的时候，我们会校验，不能创建成功，需要你按照50个一批的数量传，需要分批调用该接口，第二次传的时候，需要带上wlb_order_code和is_finished和order_item_list三个字段是必传的，is_finished为true表示传输完毕，为false表示还没完全传完。 */
func (r *WlbOrderCreateRequest) SetOrderItemList(value string) {
	r.SetValue("order_item_list", value)
}

/* 订单子类型： 
（1）OTHER： 其他 
（2）TAOBAO_TRADE： 淘宝交易 
（3）OTHER_TRADE：其他交易 
（4）ALLCOATE： 调拨 
（5）PURCHASE:采购 */
func (r *WlbOrderCreateRequest) SetOrderSubType(value string) {
	r.SetValue("order_sub_type", value)
}

/* 订单类型: 
（1）NORMAL_OUT ：正常出库 
（2）NORMAL_IN：正常入库 
（3）RETURN_IN：退货入库 
（4）EXCHANGE_OUT：换货出库 */
func (r *WlbOrderCreateRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}

/* 外部订单业务ID，该编号在isv中是唯一编号， 用来控制并发，去重用 */
func (r *WlbOrderCreateRequest) SetOutBizCode(value string) {
	r.SetValue("out_biz_code", value)
}

/* 包裹件数，入库单和出库单中会用到 */
func (r *WlbOrderCreateRequest) SetPackageCount(value string) {
	r.SetValue("package_count", value)
}

/* 应收金额，cod订单必选 */
func (r *WlbOrderCreateRequest) SetPayableAmount(value string) {
	r.SetValue("payable_amount", value)
}

/* 源订单编号 */
func (r *WlbOrderCreateRequest) SetPrevOrderCode(value string) {
	r.SetValue("prev_order_code", value)
}

/* 收货方信息，必须传， 手机和电话必选其一。
收货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话

如果某一个字段的数据为空时，必须传NA */
func (r *WlbOrderCreateRequest) SetReceiverInfo(value string) {
	r.SetValue("receiver_info", value)
}

/* 备注 */
func (r *WlbOrderCreateRequest) SetRemark(value string) {
	r.SetValue("remark", value)
}

/* 投递时间范围要求,格式'15:20'用分号隔开 */
func (r *WlbOrderCreateRequest) SetScheduleEnd(value string) {
	r.SetValue("schedule_end", value)
}

/* 投递时间范围要求,格式'13:20'用分号隔开 */
func (r *WlbOrderCreateRequest) SetScheduleStart(value string) {
	r.SetValue("schedule_start", value)
}

/* 投递时延要求:  
（1）INSTANT_ARRIVED： 当日达  
（2）TOMMORROY_MORNING_ARRIVED：次晨达 
（3）TOMMORROY_ARRIVED：次日达 
（4）工作日：WORK_DAY  
（5）节假日：WEEKED_DAY */
func (r *WlbOrderCreateRequest) SetScheduleType(value string) {
	r.SetValue("schedule_type", value)
}

/* 发货方信息，发货方信息必须传， 手机和电话必选其一。 发货方信息：
邮编^^^省^^^市^^^区^^^具体地址^^^收件方名称^^^手机^^^电话
如果某一个字段的数据为空时，必须传NA */
func (r *WlbOrderCreateRequest) SetSenderInfo(value string) {
	r.SetValue("sender_info", value)
}

/* cod服务费，只有cod订单的时候，才需要这个字段 */
func (r *WlbOrderCreateRequest) SetServiceFee(value string) {
	r.SetValue("service_fee", value)
}

/* 仓库编码 */
func (r *WlbOrderCreateRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}

/* 出库单中可能会用到 
运输公司名称^^^运输公司联系人^^^运输公司运单号^^^运输公司电话^^^运输公司联系人身份证号 
 
======================================== 
如果某一个字段的数据为空时，必须传NA */
func (r *WlbOrderCreateRequest) SetTmsInfo(value string) {
	r.SetValue("tms_info", value)
}

/* 运单编号，退货单时可能会使用 */
func (r *WlbOrderCreateRequest) SetTmsOrderCode(value string) {
	r.SetValue("tms_order_code", value)
}

/* 物流公司编码 */
func (r *WlbOrderCreateRequest) SetTmsServiceCode(value string) {
	r.SetValue("tms_service_code", value)
}

/* 总金额 */
func (r *WlbOrderCreateRequest) SetTotalAmount(value string) {
	r.SetValue("total_amount", value)
}


func (r *WlbOrderCreateRequest) GetResponse(accessToken string) (*WlbOrderCreateResponse, []byte, error) {
	var resp WlbOrderCreateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.create", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderCreateResponse struct {
	CreateTime string `json:"create_time"`
	OrderCode string `json:"order_code"`
}

type WlbOrderCreateResponseResult struct {
	Response *WlbOrderCreateResponse `json:"wlb_order_create_response"`
}





/* 分页查询物流宝订单 */
type WlbOrderPageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询截止时间 */
func (r *WlbOrderPageGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 物流订单编号 */
func (r *WlbOrderPageGetRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}

/* 订单状态 */
func (r *WlbOrderPageGetRequest) SetOrderStatus(value string) {
	r.SetValue("order_status", value)
}

/* 订单子类型：  
（1）OTHER： 其他  
（2）TAOBAO_TRADE： 淘宝交易  
（3）OTHER_TRADE：其他交易 
（4）ALLCOATE： 调拨 
（5）CHECK:  盘点单 
（6）PURCHASE: 采购单 */
func (r *WlbOrderPageGetRequest) SetOrderSubType(value string) {
	r.SetValue("order_sub_type", value)
}

/* 订单类型:  
（1）NORMAL_OUT ：正常出库  
（2）NORMAL_IN：正常入库  
（3）RETURN_IN：退货入库  
（4）EXCHANGE_OUT：换货出库 */
func (r *WlbOrderPageGetRequest) SetOrderType(value string) {
	r.SetValue("order_type", value)
}

/* 分页的第几页 */
func (r *WlbOrderPageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页多少条 */
func (r *WlbOrderPageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询开始时间 */
func (r *WlbOrderPageGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *WlbOrderPageGetRequest) GetResponse(accessToken string) (*WlbOrderPageGetResponse, []byte, error) {
	var resp WlbOrderPageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.page.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderPageGetResponse struct {
	OrderList []*WlbOrder `json:"order_list"`
	TotalCount int `json:"total_count"`
}

type WlbOrderPageGetResponseResult struct {
	Response *WlbOrderPageGetResponse `json:"wlb_order_page_get_response"`
}





/* 为订单的自动流转设置订单调度规则。规则设定之后，可以根据发货地区，精确路由订单至指定仓库处理。 */
type WlbOrderScheduleRuleAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 备用发货仓库服务id（通过taobao.wlb.subscription.query接口获得service_id） */
func (r *WlbOrderScheduleRuleAddRequest) SetBackupStoreId(value string) {
	r.SetValue("backup_store_id", value)
}

/* 发货仓库服务id（通过taobao.wlb.subscription.query接口获得service_id） */
func (r *WlbOrderScheduleRuleAddRequest) SetDefaultStoreId(value string) {
	r.SetValue("default_store_id", value)
}

/* 发货规则的额外规则设置： 
REMARK_STOP--有订单留言不自动下发 
COD_STOP--货到付款订单不自动下发 
CHECK_WAREHOUSE_DELIVER--检查仓库的配送范围 */
func (r *WlbOrderScheduleRuleAddRequest) SetOption(value string) {
	r.SetValue("option", value)
}

/* 国家地区标准编码列表 */
func (r *WlbOrderScheduleRuleAddRequest) SetProvAreaIds(value string) {
	r.SetValue("prov_area_ids", value)
}


func (r *WlbOrderScheduleRuleAddRequest) GetResponse(accessToken string) (*WlbOrderScheduleRuleAddResponse, []byte, error) {
	var resp WlbOrderScheduleRuleAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.schedule.rule.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderScheduleRuleAddResponse struct {
	ScheduleRuleId int `json:"schedule_rule_id"`
}

type WlbOrderScheduleRuleAddResponseResult struct {
	Response *WlbOrderScheduleRuleAddResponse `json:"wlb_order_schedule_rule_add_response"`
}





/* 修改物流宝订单调度规则 */
type WlbOrderScheduleRuleUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 备用发货仓库id */
func (r *WlbOrderScheduleRuleUpdateRequest) SetBackupStoreId(value string) {
	r.SetValue("backup_store_id", value)
}

/* 默认发货仓库id */
func (r *WlbOrderScheduleRuleUpdateRequest) SetDefaultStoreId(value string) {
	r.SetValue("default_store_id", value)
}

/* 订单调度规则的额外规则设置： REMARK_STOP--有订单留言不自动下发 COD_STOP--货到付款订单不自动下发 CHECK_WAREHOUSE_DELIVER--检查仓库的配送范围 */
func (r *WlbOrderScheduleRuleUpdateRequest) SetOption(value string) {
	r.SetValue("option", value)
}

/* 国家地区标准编码列表 */
func (r *WlbOrderScheduleRuleUpdateRequest) SetProvAreaIds(value string) {
	r.SetValue("prov_area_ids", value)
}

/* 要修改的订单调度规则明细id */
func (r *WlbOrderScheduleRuleUpdateRequest) SetScheduleRuleId(value string) {
	r.SetValue("schedule_rule_id", value)
}


func (r *WlbOrderScheduleRuleUpdateRequest) GetResponse(accessToken string) (*WlbOrderScheduleRuleUpdateResponse, []byte, error) {
	var resp WlbOrderScheduleRuleUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.order.schedule.rule.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderScheduleRuleUpdateResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbOrderScheduleRuleUpdateResponseResult struct {
	Response *WlbOrderScheduleRuleUpdateResponse `json:"wlb_order_schedule_rule_update_response"`
}





/* 外部ERP可通过该接口查询一段时间内的物流宝订单，以及订单详情 */
type WlbOrderdetailDateGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询条件截止日期 */
func (r *WlbOrderdetailDateGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (r *WlbOrderdetailDateGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (r *WlbOrderdetailDateGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 查询起始日期 */
func (r *WlbOrderdetailDateGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *WlbOrderdetailDateGetRequest) GetResponse(accessToken string) (*WlbOrderdetailDateGetResponse, []byte, error) {
	var resp WlbOrderdetailDateGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.orderdetail.date.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderdetailDateGetResponse struct {
	OrderDetailList []*WlbOrderDetail `json:"order_detail_list"`
	TotalCount int `json:"total_count"`
}

type WlbOrderdetailDateGetResponseResult struct {
	Response *WlbOrderdetailDateGetResponse `json:"wlb_orderdetail_date_get_response"`
}





/* 分页查询物流宝订单商品详情 */
type WlbOrderitemPageGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝订单编码 */
func (r *WlbOrderitemPageGetRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}

/* 分页查询参数，指定查询页数，默认为1 */
func (r *WlbOrderitemPageGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页查询参数，每页查询数量，默认20，最大值50,大于50时按照50条查询 */
func (r *WlbOrderitemPageGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *WlbOrderitemPageGetRequest) GetResponse(accessToken string) (*WlbOrderitemPageGetResponse, []byte, error) {
	var resp WlbOrderitemPageGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.orderitem.page.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderitemPageGetResponse struct {
	OrderItemList []*WlbOrderItem `json:"order_item_list"`
	TotalCount int `json:"total_count"`
}

type WlbOrderitemPageGetResponseResult struct {
	Response *WlbOrderitemPageGetResponse `json:"wlb_orderitem_page_get_response"`
}





/* 删除单个订单调度规则 */
type WlbOrderscheduleruleDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 订单调度规则ID */
func (r *WlbOrderscheduleruleDeleteRequest) SetId(value string) {
	r.SetValue("id", value)
}

/* 商品userNick */
func (r *WlbOrderscheduleruleDeleteRequest) SetUserNick(value string) {
	r.SetValue("user_nick", value)
}


func (r *WlbOrderscheduleruleDeleteRequest) GetResponse(accessToken string) (*WlbOrderscheduleruleDeleteResponse, []byte, error) {
	var resp WlbOrderscheduleruleDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.orderschedulerule.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderscheduleruleDeleteResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbOrderscheduleruleDeleteResponseResult struct {
	Response *WlbOrderscheduleruleDeleteResponse `json:"wlb_orderschedulerule_delete_response"`
}





/* 查询某个卖家的所有订单调度规则，提供分页查询。 */
type WlbOrderscheduleruleQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 当前页 */
func (r *WlbOrderscheduleruleQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *WlbOrderscheduleruleQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *WlbOrderscheduleruleQueryRequest) GetResponse(accessToken string) (*WlbOrderscheduleruleQueryResponse, []byte, error) {
	var resp WlbOrderscheduleruleQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.orderschedulerule.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderscheduleruleQueryResponse struct {
	OrderScheduleRuleList []*WlbOrderScheduleRule `json:"order_schedule_rule_list"`
	TotalCount int `json:"total_count"`
}

type WlbOrderscheduleruleQueryResponseResult struct {
	Response *WlbOrderscheduleruleQueryResponse `json:"wlb_orderschedulerule_query_response"`
}





/* 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表 */
type WlbOrderstatusGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝订单编码 */
func (r *WlbOrderstatusGetRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}


func (r *WlbOrderstatusGetRequest) GetResponse(accessToken string) (*WlbOrderstatusGetResponse, []byte, error) {
	var resp WlbOrderstatusGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.orderstatus.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOrderstatusGetResponse struct {
	WlbOrderStatus []*WlbProcessStatus `json:"wlb_order_status"`
}

type WlbOrderstatusGetResponseResult struct {
	Response *WlbOrderstatusGetResponse `json:"wlb_orderstatus_get_response"`
}





/* 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。 */
type WlbOutInventoryChangeNotifyRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 库存变化数量 */
func (r *WlbOutInventoryChangeNotifyRequest) SetChangeCount(value string) {
	r.SetValue("change_count", value)
}

/* 物流宝商品id或前台宝贝id（由type类型决定） */
func (r *WlbOutInventoryChangeNotifyRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* OUT--出库 
IN--入库 */
func (r *WlbOutInventoryChangeNotifyRequest) SetOpType(value string) {
	r.SetValue("op_type", value)
}

/* 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号 */
func (r *WlbOutInventoryChangeNotifyRequest) SetOrderSourceCode(value string) {
	r.SetValue("order_source_code", value)
}

/* 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。 */
func (r *WlbOutInventoryChangeNotifyRequest) SetOutBizCode(value string) {
	r.SetValue("out_biz_code", value)
}

/* 本次库存变化后库存余额 */
func (r *WlbOutInventoryChangeNotifyRequest) SetResultCount(value string) {
	r.SetValue("result_count", value)
}

/* （1）OTHER： 其他  
（2）TAOBAO_TRADE： 淘宝交易  
（3）OTHER_TRADE：其他交易  
（4）ALLCOATE： 调拨  
（5）CHECK:盘点 
（6）PURCHASE:采购 */
func (r *WlbOutInventoryChangeNotifyRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 目前非必须，系统会选择默认值 */
func (r *WlbOutInventoryChangeNotifyRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}

/* WLB_ITEM--物流宝商品 
IC_ITEM--淘宝商品 
IC_SKU--淘宝sku */
func (r *WlbOutInventoryChangeNotifyRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *WlbOutInventoryChangeNotifyRequest) GetResponse(accessToken string) (*WlbOutInventoryChangeNotifyResponse, []byte, error) {
	var resp WlbOutInventoryChangeNotifyResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.out.inventory.change.notify", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbOutInventoryChangeNotifyResponse struct {
	GmtModified string `json:"gmt_modified"`
}

type WlbOutInventoryChangeNotifyResponseResult struct {
	Response *WlbOutInventoryChangeNotifyResponse `json:"wlb_out_inventory_change_notify_response"`
}





/* 查询BI的统计的补货统计数据 */
type WlbReplenishStatisticsRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 商品编码 */
func (r *WlbReplenishStatisticsRequest) SetItemCode(value string) {
	r.SetValue("item_code", value)
}

/* 商品名称 */
func (r *WlbReplenishStatisticsRequest) SetName(value string) {
	r.SetValue("name", value)
}

/* 分页参数，默认第一页 */
func (r *WlbReplenishStatisticsRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页每页页数，默认20，最大50 */
func (r *WlbReplenishStatisticsRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 仓库编码 */
func (r *WlbReplenishStatisticsRequest) SetStoreCode(value string) {
	r.SetValue("store_code", value)
}


func (r *WlbReplenishStatisticsRequest) GetResponse(accessToken string) (*WlbReplenishStatisticsResponse, []byte, error) {
	var resp WlbReplenishStatisticsResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.replenish.statistics", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbReplenishStatisticsResponse struct {
	ReplenishList []*WlbReplenish `json:"replenish_list"`
	TotalCount int `json:"total_count"`
}

type WlbReplenishStatisticsResponseResult struct {
	Response *WlbReplenishStatisticsResponse `json:"wlb_replenish_statistics_response"`
}





/* 查询商家定购的所有服务,可通过入参状态来筛选 */
type WlbSubscriptionQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 当前页 */
func (r *WlbSubscriptionQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *WlbSubscriptionQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 状态  
AUDITING 1-待审核;  
CANCEL 2-撤销 ; 
CHECKED 3-审核通过 ; 
FAILED 4-审核未通过 ; 
SYNCHRONIZING 5-同步中; 
只允许输入上面指定的值，且可以为空，为空时查询所有状态。若输错了，则按AUDITING处理。 */
func (r *WlbSubscriptionQueryRequest) SetStatus(value string) {
	r.SetValue("status", value)
}


func (r *WlbSubscriptionQueryRequest) GetResponse(accessToken string) (*WlbSubscriptionQueryResponse, []byte, error) {
	var resp WlbSubscriptionQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.subscription.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbSubscriptionQueryResponse struct {
	SellerSubscriptionList []*WlbSellerSubscription `json:"seller_subscription_list"`
	TotalCount int `json:"total_count"`
}

type WlbSubscriptionQueryResponseResult struct {
	Response *WlbSubscriptionQueryResponse `json:"wlb_subscription_query_response"`
}





/* 通过物流订单编号分页查询物流信息 */
type WlbTmsorderQueryRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流订单编号 */
func (r *WlbTmsorderQueryRequest) SetOrderCode(value string) {
	r.SetValue("order_code", value)
}

/* 当前页 */
func (r *WlbTmsorderQueryRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录 */
func (r *WlbTmsorderQueryRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *WlbTmsorderQueryRequest) GetResponse(accessToken string) (*WlbTmsorderQueryResponse, []byte, error) {
	var resp WlbTmsorderQueryResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.tmsorder.query", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbTmsorderQueryResponse struct {
	TmsOrderList []*WlbTmsOrder `json:"tms_order_list"`
	TotalCount int `json:"total_count"`
}

type WlbTmsorderQueryResponseResult struct {
	Response *WlbTmsorderQueryResponse `json:"wlb_tmsorder_query_response"`
}





/* 根据交易类型和交易id查询物流宝订单详情 */
type WlbTradeorderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 指定交易类型的交易号 */
func (r *WlbTradeorderGetRequest) SetTradeId(value string) {
	r.SetValue("trade_id", value)
}

/* 交易类型: 
TAOBAO--淘宝交易 
PAIPAI--拍拍交易 
YOUA--有啊交易 */
func (r *WlbTradeorderGetRequest) SetTradeType(value string) {
	r.SetValue("trade_type", value)
}


func (r *WlbTradeorderGetRequest) GetResponse(accessToken string) (*WlbTradeorderGetResponse, []byte, error) {
	var resp WlbTradeorderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.tradeorder.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbTradeorderGetResponse struct {
	WlbOrderList []*WlbOrder `json:"wlb_order_list"`
}

type WlbTradeorderGetResponseResult struct {
	Response *WlbTradeorderGetResponse `json:"wlb_tradeorder_get_response"`
}





/* 商家申请电子面单 */
type WlbWaybillallocationRequestwaybillnumRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 申请号码数量 */
func (r *WlbWaybillallocationRequestwaybillnumRequest) SetNum(value string) {
	r.SetValue("num", value)
}

/* 外部去重号 */
func (r *WlbWaybillallocationRequestwaybillnumRequest) SetOutBizCode(value string) {
	r.SetValue("out_biz_code", value)
}

/* 面单类型 */
func (r *WlbWaybillallocationRequestwaybillnumRequest) SetPoolType(value string) {
	r.SetValue("pool_type", value)
}

/* 服务编码 */
func (r *WlbWaybillallocationRequestwaybillnumRequest) SetServiceCode(value string) {
	r.SetValue("service_code", value)
}

/* 用户ID */
func (r *WlbWaybillallocationRequestwaybillnumRequest) SetUserId(value string) {
	r.SetValue("user_id", value)
}


func (r *WlbWaybillallocationRequestwaybillnumRequest) GetResponse(accessToken string) (*WlbWaybillallocationRequestwaybillnumResponse, []byte, error) {
	var resp WlbWaybillallocationRequestwaybillnumResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.waybillallocation.requestwaybillnum", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbWaybillallocationRequestwaybillnumResponse struct {
	ServiceCode string `json:"service_code"`
	TmsWaybillList []string `json:"tms_waybill_list"`
}

type WlbWaybillallocationRequestwaybillnumResponseResult struct {
	Response *WlbWaybillallocationRequestwaybillnumResponse `json:"wlb_waybillallocation_requestwaybillnum_response"`
}





/* 根据物流宝订单编号查询物流宝订单概要信息 */
type WlbWlborderGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 物流宝订单编码 */
func (r *WlbWlborderGetRequest) SetWlbOrderCode(value string) {
	r.SetValue("wlb_order_code", value)
}


func (r *WlbWlborderGetRequest) GetResponse(accessToken string) (*WlbWlborderGetResponse, []byte, error) {
	var resp WlbWlborderGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.wlb.wlborder.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type WlbWlborderGetResponse struct {
	WlbOrder *WlbOrder `json:"wlb_order"`
}

type WlbWlborderGetResponseResult struct {
	Response *WlbWlborderGetResponse `json:"wlb_wlborder_get_response"`
}



