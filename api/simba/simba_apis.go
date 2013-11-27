// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simba

import (
	"github.com/changkong/open_taobao"
)





/* 获取实时余额，”元”为单位 */
type SimbaAccountBalanceGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAccountBalanceGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAccountBalanceGetRequest) GetResponse(accessToken string) (*SimbaAccountBalanceGetResponse, []byte, error) {
	var resp SimbaAccountBalanceGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.account.balance.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAccountBalanceGetResponse struct {
	Balance float64 `json:"balance"`
}

type SimbaAccountBalanceGetResponseResult struct {
	Response *SimbaAccountBalanceGetResponse `json:"simba_account_balance_get_response"`
}





/* 创建一个推广组 */
type SimbaAdgroupAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaAdgroupAddRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 推广组默认出价，单位为分，不能小于5 不能大于日最高限额 */
func (r *SimbaAdgroupAddRequest) SetDefaultPrice(value string) {
	r.SetValue("default_price", value)
}

/* 创意图片地址，必须是商品的图片之一 */
func (r *SimbaAdgroupAddRequest) SetImgUrl(value string) {
	r.SetValue("img_url", value)
}

/* 商品Id */
func (r *SimbaAdgroupAddRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 创意标题，最多20个汉字 */
func (r *SimbaAdgroupAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *SimbaAdgroupAddRequest) GetResponse(accessToken string) (*SimbaAdgroupAddResponse, []byte, error) {
	var resp SimbaAdgroupAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupAddResponse struct {
	Adgroup *ADGroup `json:"adgroup"`
}

type SimbaAdgroupAddResponseResult struct {
	Response *SimbaAdgroupAddResponse `json:"simba_adgroup_add_response"`
}





/* 根据一组推广组id获取推广组类目出价列表, */
type SimbaAdgroupAdgroupcatmatchsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id列表 */
func (r *SimbaAdgroupAdgroupcatmatchsGetRequest) SetAdgroupIds(value string) {
	r.SetValue("adgroup_ids", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupAdgroupcatmatchsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupAdgroupcatmatchsGetRequest) GetResponse(accessToken string) (*SimbaAdgroupAdgroupcatmatchsGetResponse, []byte, error) {
	var resp SimbaAdgroupAdgroupcatmatchsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.adgroupcatmatchs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupAdgroupcatmatchsGetResponse struct {
	AdgroupCatmatchList []*ADGroupCatmatch `json:"adgroup_catmatch_list"`
}

type SimbaAdgroupAdgroupcatmatchsGetResponseResult struct {
	Response *SimbaAdgroupAdgroupcatmatchsGetResponse `json:"simba_adgroup_adgroupcatmatchs_get_response"`
}





/* 根据一个推广计划的id获取一页推广组类目出价列表 */
type SimbaAdgroupCampcatmatchsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaAdgroupCampcatmatchsGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupCampcatmatchsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码，从1开始 */
func (r *SimbaAdgroupCampcatmatchsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页尺寸，最大200 */
func (r *SimbaAdgroupCampcatmatchsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *SimbaAdgroupCampcatmatchsGetRequest) GetResponse(accessToken string) (*SimbaAdgroupCampcatmatchsGetResponse, []byte, error) {
	var resp SimbaAdgroupCampcatmatchsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.campcatmatchs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupCampcatmatchsGetResponse struct {
	Adgroupcatmatchs *ADGroupCatMatchPage `json:"adgroupcatmatchs"`
}

type SimbaAdgroupCampcatmatchsGetResponseResult struct {
	Response *SimbaAdgroupCampcatmatchsGetResponse `json:"simba_adgroup_campcatmatchs_get_response"`
}





/* 取得一个推广组的类目出价 */
type SimbaAdgroupCatmatchGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaAdgroupCatmatchGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupCatmatchGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupCatmatchGetRequest) GetResponse(accessToken string) (*SimbaAdgroupCatmatchGetResponse, []byte, error) {
	var resp SimbaAdgroupCatmatchGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.catmatch.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupCatmatchGetResponse struct {
	Adgroupcatmatch *ADGroupCatmatch `json:"adgroupcatmatch"`
}

type SimbaAdgroupCatmatchGetResponseResult struct {
	Response *SimbaAdgroupCatmatchGetResponse `json:"simba_adgroup_catmatch_get_response"`
}





/* 更新一个推广组的类目出价，可以设置类目出价、是否使用默认出价、是否打开类目出价 */
type SimbaAdgroupCatmatchUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 类目出价Id */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetCatmatchId(value string) {
	r.SetValue("catmatch_id", value)
}

/* 类目出价，单位为分，不能小于5。如果use_default_price字段为使用默认出价，则此max_price字段所传入的值不起作用。商品将会使用默认出价。 */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetMaxPrice(value string) {
	r.SetValue("max_price", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 是否启用类目出价； offline-不启用 online-启用,如果此字段状态为offline时，则max_price,use_default_price字段值不起作用 */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetOnlineStatus(value string) {
	r.SetValue("online_status", value)
}

/* 是否使用推广组默认出价false为不使用，true为使用 */
func (r *SimbaAdgroupCatmatchUpdateRequest) SetUseDefaultPrice(value string) {
	r.SetValue("use_default_price", value)
}


func (r *SimbaAdgroupCatmatchUpdateRequest) GetResponse(accessToken string) (*SimbaAdgroupCatmatchUpdateResponse, []byte, error) {
	var resp SimbaAdgroupCatmatchUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.catmatch.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupCatmatchUpdateResponse struct {
	Adgroupcatmatch *ADGroupCatmatch `json:"adgroupcatmatch"`
}

type SimbaAdgroupCatmatchUpdateResponseResult struct {
	Response *SimbaAdgroupCatmatchUpdateResponse `json:"simba_adgroup_catmatch_update_response"`
}





/* 获取指定推广组下给定出价的类目出价预估信息; */
type SimbaAdgroupCatmatchforecastGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID */
func (r *SimbaAdgroupCatmatchforecastGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 类目出价,出价范围在5-9999之间,单位分 */
func (r *SimbaAdgroupCatmatchforecastGetRequest) SetCatmatchPrice(value string) {
	r.SetValue("catmatch_price", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupCatmatchforecastGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupCatmatchforecastGetRequest) GetResponse(accessToken string) (*SimbaAdgroupCatmatchforecastGetResponse, []byte, error) {
	var resp SimbaAdgroupCatmatchforecastGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.catmatchforecast.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupCatmatchforecastGetResponse struct {
	AdgroupCatmatchForecast *ADGroupCatMatchForecast `json:"adgroup_catmatch_forecast"`
}

type SimbaAdgroupCatmatchforecastGetResponseResult struct {
	Response *SimbaAdgroupCatmatchforecastGetResponse `json:"simba_adgroup_catmatchforecast_get_response"`
}





/* 分页获取修改过的推广组类目出价ID , 推广组ID，修改时间 */
type SimbaAdgroupChangedcatmatchsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupChangedcatmatchsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaAdgroupChangedcatmatchsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaAdgroupChangedcatmatchsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaAdgroupChangedcatmatchsGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaAdgroupChangedcatmatchsGetRequest) GetResponse(accessToken string) (*SimbaAdgroupChangedcatmatchsGetResponse, []byte, error) {
	var resp SimbaAdgroupChangedcatmatchsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.changedcatmatchs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupChangedcatmatchsGetResponse struct {
	ChangedCatmatchs *ADGroupCatMatchPage `json:"changed_catmatchs"`
}

type SimbaAdgroupChangedcatmatchsGetResponseResult struct {
	Response *SimbaAdgroupChangedcatmatchsGetResponse `json:"simba_adgroup_changedcatmatchs_get_response"`
}





/* 删除一个推广组 */
type SimbaAdgroupDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaAdgroupDeleteRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupDeleteRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupDeleteRequest) GetResponse(accessToken string) (*SimbaAdgroupDeleteResponse, []byte, error) {
	var resp SimbaAdgroupDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupDeleteResponse struct {
	Adgroup *ADGroup `json:"adgroup"`
}

type SimbaAdgroupDeleteResponseResult struct {
	Response *SimbaAdgroupDeleteResponse `json:"simba_adgroup_delete_response"`
}





/* 获取删除的类目出价列表（只存类目出价ID和推广组ID） */
type SimbaAdgroupDeletedcatmatchsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupDeletedcatmatchsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaAdgroupDeletedcatmatchsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaAdgroupDeletedcatmatchsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaAdgroupDeletedcatmatchsGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaAdgroupDeletedcatmatchsGetRequest) GetResponse(accessToken string) (*SimbaAdgroupDeletedcatmatchsGetResponse, []byte, error) {
	var resp SimbaAdgroupDeletedcatmatchsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.deletedcatmatchs.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupDeletedcatmatchsGetResponse struct {
	DeletedCatmatchs *ADGroupCatMatchPage `json:"deleted_catmatchs"`
}

type SimbaAdgroupDeletedcatmatchsGetResponseResult struct {
	Response *SimbaAdgroupDeletedcatmatchsGetResponse `json:"simba_adgroup_deletedcatmatchs_get_response"`
}





/* 修改通投出价 */
type SimbaAdgroupNonsearchpricesUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id，通投位置价格，是否使用默认出价json 数组字符串，数组个数最多200个. 
json数组中的key必须和对应实体ADGroup中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值， 
adgroupId,isNonsearchDefaultPrice不能为空。nonsearchMaxPrice是整数，以“分”为单位，不能小于5，不能大于日限额,不能大于9999分。 启用非搜索默认出价时nonsearchMaxPrice为0 */
func (r *SimbaAdgroupNonsearchpricesUpdateRequest) SetAdgroupidPriceJson(value string) {
	r.SetValue("adgroupid_price_json", value)
}

/* 推广计划ID */
func (r *SimbaAdgroupNonsearchpricesUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupNonsearchpricesUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupNonsearchpricesUpdateRequest) GetResponse(accessToken string) (*SimbaAdgroupNonsearchpricesUpdateResponse, []byte, error) {
	var resp SimbaAdgroupNonsearchpricesUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.nonsearchprices.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupNonsearchpricesUpdateResponse struct {
	AdgroupList []*ADGroup `json:"adgroup_list"`
}

type SimbaAdgroupNonsearchpricesUpdateResponseResult struct {
	Response *SimbaAdgroupNonsearchpricesUpdateResponse `json:"simba_adgroup_nonsearchprices_update_response"`
}





/* 更改通投状态（暂停或启动） */
type SimbaAdgroupNonsearchstatesUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID通投状态json数组字符串，数组个数最多200个。json数组中的key必须和对应实体ADGroup 中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值推广组ID,通投状态不传默认为1 */
func (r *SimbaAdgroupNonsearchstatesUpdateRequest) SetAdgroupidNonsearchstateJson(value string) {
	r.SetValue("adgroupid_nonsearchstate_json", value)
}

/* 推广计划ID */
func (r *SimbaAdgroupNonsearchstatesUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupNonsearchstatesUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupNonsearchstatesUpdateRequest) GetResponse(accessToken string) (*SimbaAdgroupNonsearchstatesUpdateResponse, []byte, error) {
	var resp SimbaAdgroupNonsearchstatesUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.nonsearchstates.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupNonsearchstatesUpdateResponse struct {
	AdgroupList []*ADGroup `json:"adgroup_list"`
}

type SimbaAdgroupNonsearchstatesUpdateResponseResult struct {
	Response *SimbaAdgroupNonsearchstatesUpdateResponse `json:"simba_adgroup_nonsearchstates_update_response"`
}





/* 获取用户上架在线销售的全部宝贝 */
type SimbaAdgroupOnlineitemsvonGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupOnlineitemsvonGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 排序，true:降序， false:升序 */
func (r *SimbaAdgroupOnlineitemsvonGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 排序字段，starts：按开始时间排序bidCount:按销量排序 */
func (r *SimbaAdgroupOnlineitemsvonGetRequest) SetOrderField(value string) {
	r.SetValue("order_field", value)
}

/* 页码，从1开始,最大50。最大只能获取1W个宝贝 */
func (r *SimbaAdgroupOnlineitemsvonGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页尺寸，最大200 */
func (r *SimbaAdgroupOnlineitemsvonGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *SimbaAdgroupOnlineitemsvonGetRequest) GetResponse(accessToken string) (*SimbaAdgroupOnlineitemsvonGetResponse, []byte, error) {
	var resp SimbaAdgroupOnlineitemsvonGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.onlineitemsvon.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupOnlineitemsvonGetResponse struct {
	PageItem *SubwayItemPartition `json:"page_item"`
}

type SimbaAdgroupOnlineitemsvonGetResponseResult struct {
	Response *SimbaAdgroupOnlineitemsvonGetResponse `json:"simba_adgroup_onlineitemsvon_get_response"`
}





/* 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价 */
type SimbaAdgroupUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaAdgroupUpdateRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 默认出价，单位是分，不能小于5 */
func (r *SimbaAdgroupUpdateRequest) SetDefaultPrice(value string) {
	r.SetValue("default_price", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价 */
func (r *SimbaAdgroupUpdateRequest) SetNonsearchMaxPrice(value string) {
	r.SetValue("nonsearch_max_price", value)
}

/* 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online */
func (r *SimbaAdgroupUpdateRequest) SetOnlineStatus(value string) {
	r.SetValue("online_status", value)
}

/* 非搜索是否使用默认出价，false-不用；true-使用；默认为true; */
func (r *SimbaAdgroupUpdateRequest) SetUseNonsearchDefaultPrice(value string) {
	r.SetValue("use_nonsearch_default_price", value)
}


func (r *SimbaAdgroupUpdateRequest) GetResponse(accessToken string) (*SimbaAdgroupUpdateResponse, []byte, error) {
	var resp SimbaAdgroupUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroup.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupUpdateResponse struct {
	Adgroup *ADGroup `json:"adgroup"`
}

type SimbaAdgroupUpdateResponseResult struct {
	Response *SimbaAdgroupUpdateResponse `json:"simba_adgroup_update_response"`
}





/* 获取修改的推广组ID */
type SimbaAdgroupidsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupidsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaAdgroupidsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaAdgroupidsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaAdgroupidsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaAdgroupidsChangedGetRequest) GetResponse(accessToken string) (*SimbaAdgroupidsChangedGetResponse, []byte, error) {
	var resp SimbaAdgroupidsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroupids.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupidsChangedGetResponse struct {
	ChangedAdgroupids []int `json:"changed_adgroupids"`
}

type SimbaAdgroupidsChangedGetResponseResult struct {
	Response *SimbaAdgroupidsChangedGetResponse `json:"simba_adgroupids_changed_get_response"`
}





/* 获取删除的推广组ID */
type SimbaAdgroupidsDeletedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupidsDeletedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaAdgroupidsDeletedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaAdgroupidsDeletedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaAdgroupidsDeletedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaAdgroupidsDeletedGetRequest) GetResponse(accessToken string) (*SimbaAdgroupidsDeletedGetResponse, []byte, error) {
	var resp SimbaAdgroupidsDeletedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroupids.deleted.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupidsDeletedGetResponse struct {
	DeletedAdgroupIds []int `json:"deleted_adgroup_ids"`
}

type SimbaAdgroupidsDeletedGetResponseResult struct {
	Response *SimbaAdgroupidsDeletedGetResponse `json:"simba_adgroupids_deleted_get_response"`
}





/* 分页获取修改的推广组ID和修改时间 */
type SimbaAdgroupsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaAdgroupsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaAdgroupsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaAdgroupsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaAdgroupsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaAdgroupsChangedGetRequest) GetResponse(accessToken string) (*SimbaAdgroupsChangedGetResponse, []byte, error) {
	var resp SimbaAdgroupsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroups.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupsChangedGetResponse struct {
	Adgroups *ADGroupPage `json:"adgroups"`
}

type SimbaAdgroupsChangedGetResponseResult struct {
	Response *SimbaAdgroupsChangedGetResponse `json:"simba_adgroups_changed_get_response"`
}





/* 判断在一个推广计划中是否已经推广了一个商品 */
type SimbaAdgroupsItemExistRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaAdgroupsItemExistRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 商品Id */
func (r *SimbaAdgroupsItemExistRequest) SetItemId(value string) {
	r.SetValue("item_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupsItemExistRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaAdgroupsItemExistRequest) GetResponse(accessToken string) (*SimbaAdgroupsItemExistResponse, []byte, error) {
	var resp SimbaAdgroupsItemExistResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroups.item.exist", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupsItemExistResponse struct {
	Exist bool `json:"exist"`
}

type SimbaAdgroupsItemExistResponseResult struct {
	Response *SimbaAdgroupsItemExistResponse `json:"simba_adgroups_item_exist_response"`
}





/* 批量得到推广组 */
type SimbaAdgroupsbyadgroupidsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id列表 */
func (r *SimbaAdgroupsbyadgroupidsGetRequest) SetAdgroupIds(value string) {
	r.SetValue("adgroup_ids", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupsbyadgroupidsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码，从1开始 */
func (r *SimbaAdgroupsbyadgroupidsGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码 */
func (r *SimbaAdgroupsbyadgroupidsGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *SimbaAdgroupsbyadgroupidsGetRequest) GetResponse(accessToken string) (*SimbaAdgroupsbyadgroupidsGetResponse, []byte, error) {
	var resp SimbaAdgroupsbyadgroupidsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroupsbyadgroupids.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupsbyadgroupidsGetResponse struct {
	Adgroups *ADGroupPage `json:"adgroups"`
}

type SimbaAdgroupsbyadgroupidsGetResponseResult struct {
	Response *SimbaAdgroupsbyadgroupidsGetResponse `json:"simba_adgroupsbyadgroupids_get_response"`
}





/* 批量得到推广计划下的推广组 */
type SimbaAdgroupsbycampaignidGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaAdgroupsbycampaignidGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaAdgroupsbycampaignidGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码，从1开始 */
func (r *SimbaAdgroupsbycampaignidGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码 */
func (r *SimbaAdgroupsbycampaignidGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}


func (r *SimbaAdgroupsbycampaignidGetRequest) GetResponse(accessToken string) (*SimbaAdgroupsbycampaignidGetResponse, []byte, error) {
	var resp SimbaAdgroupsbycampaignidGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.adgroupsbycampaignid.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaAdgroupsbycampaignidGetResponse struct {
	Adgroups *ADGroupPage `json:"adgroups"`
}

type SimbaAdgroupsbycampaignidGetResponseResult struct {
	Response *SimbaAdgroupsbycampaignidGetResponse `json:"simba_adgroupsbycampaignid_get_response"`
}





/* 创建一个推广计划 */
type SimbaCampaignAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCampaignAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。 */
func (r *SimbaCampaignAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *SimbaCampaignAddRequest) GetResponse(accessToken string) (*SimbaCampaignAddResponse, []byte, error) {
	var resp SimbaCampaignAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignAddResponse struct {
	Campaign *Campaign `json:"campaign"`
}

type SimbaCampaignAddResponseResult struct {
	Response *SimbaCampaignAddResponse `json:"simba_campaign_add_response"`
}





/* 取得一个推广计划的投放地域设置 */
type SimbaCampaignAreaGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignAreaGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignAreaGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignAreaGetRequest) GetResponse(accessToken string) (*SimbaCampaignAreaGetResponse, []byte, error) {
	var resp SimbaCampaignAreaGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.area.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignAreaGetResponse struct {
	CampaignArea *CampaignArea `json:"campaign_area"`
}

type SimbaCampaignAreaGetResponseResult struct {
	Response *SimbaCampaignAreaGetResponse `json:"simba_campaign_area_get_response"`
}





/* 更新一个推广计划的投放地域 */
type SimbaCampaignAreaUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID； */
func (r *SimbaCampaignAreaUpdateRequest) SetArea(value string) {
	r.SetValue("area", value)
}

/* 推广计划Id */
func (r *SimbaCampaignAreaUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignAreaUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignAreaUpdateRequest) GetResponse(accessToken string) (*SimbaCampaignAreaUpdateResponse, []byte, error) {
	var resp SimbaCampaignAreaUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.area.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignAreaUpdateResponse struct {
	CampaignArea *CampaignArea `json:"campaign_area"`
}

type SimbaCampaignAreaUpdateResponseResult struct {
	Response *SimbaCampaignAreaUpdateResponse `json:"simba_campaign_area_update_response"`
}





/* 取得推广计划的可设置投放地域列表 */
type SimbaCampaignAreaoptionsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *SimbaCampaignAreaoptionsGetRequest) GetResponse(accessToken string) (*SimbaCampaignAreaoptionsGetResponse, []byte, error) {
	var resp SimbaCampaignAreaoptionsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.areaoptions.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignAreaoptionsGetResponse struct {
	AreaOptions []*AreaOption `json:"area_options"`
}

type SimbaCampaignAreaoptionsGetResponseResult struct {
	Response *SimbaCampaignAreaoptionsGetResponse `json:"simba_campaign_areaoptions_get_response"`
}





/* 取得一个推广计划的日限额 */
type SimbaCampaignBudgetGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignBudgetGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignBudgetGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignBudgetGetRequest) GetResponse(accessToken string) (*SimbaCampaignBudgetGetResponse, []byte, error) {
	var resp SimbaCampaignBudgetGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.budget.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignBudgetGetResponse struct {
	CampaignBudget *CampaignBudget `json:"campaign_budget"`
}

type SimbaCampaignBudgetGetResponseResult struct {
	Response *SimbaCampaignBudgetGetResponse `json:"simba_campaign_budget_get_response"`
}





/* 更新一个推广计划的日限额 */
type SimbaCampaignBudgetUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 如果为空则取消限额；否则必须为整数，单位是元，不得小于30； */
func (r *SimbaCampaignBudgetUpdateRequest) SetBudget(value string) {
	r.SetValue("budget", value)
}

/* 推广计划Id */
func (r *SimbaCampaignBudgetUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignBudgetUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 是否平滑消耗：false-否，true-是 */
func (r *SimbaCampaignBudgetUpdateRequest) SetUseSmooth(value string) {
	r.SetValue("use_smooth", value)
}


func (r *SimbaCampaignBudgetUpdateRequest) GetResponse(accessToken string) (*SimbaCampaignBudgetUpdateResponse, []byte, error) {
	var resp SimbaCampaignBudgetUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.budget.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignBudgetUpdateResponse struct {
	CampaignBudget *CampaignBudget `json:"campaign_budget"`
}

type SimbaCampaignBudgetUpdateResponseResult struct {
	Response *SimbaCampaignBudgetUpdateResponse `json:"simba_campaign_budget_update_response"`
}





/* 取得推广计划的可设置投放频道列表 */
type SimbaCampaignChanneloptionsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *SimbaCampaignChanneloptionsGetRequest) GetResponse(accessToken string) (*SimbaCampaignChanneloptionsGetResponse, []byte, error) {
	var resp SimbaCampaignChanneloptionsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.channeloptions.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignChanneloptionsGetResponse struct {
	ChannelOptions []*ChannelOption `json:"channel_options"`
}

type SimbaCampaignChanneloptionsGetResponseResult struct {
	Response *SimbaCampaignChanneloptionsGetResponse `json:"simba_campaign_channeloptions_get_response"`
}





/* 取得一个推广计划的投放平台设置 */
type SimbaCampaignPlatformGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignPlatformGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignPlatformGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignPlatformGetRequest) GetResponse(accessToken string) (*SimbaCampaignPlatformGetResponse, []byte, error) {
	var resp SimbaCampaignPlatformGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.platform.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignPlatformGetResponse struct {
	CampaignPlatform *CampaignPlatform `json:"campaign_platform"`
}

type SimbaCampaignPlatformGetResponseResult struct {
	Response *SimbaCampaignPlatformGetResponse `json:"simba_campaign_platform_get_response"`
}





/* 更新一个推广计划的平台设置 */
type SimbaCampaignPlatformUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignPlatformUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignPlatformUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。 */
func (r *SimbaCampaignPlatformUpdateRequest) SetNonsearchChannels(value string) {
	r.SetValue("nonsearch_channels", value)
}

/* 溢价的百分比，必须是大于等于 1小于等于200的整数 */
func (r *SimbaCampaignPlatformUpdateRequest) SetOutsideDiscount(value string) {
	r.SetValue("outside_discount", value)
}

/* 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值，必须包含淘宝内网。 */
func (r *SimbaCampaignPlatformUpdateRequest) SetSearchChannels(value string) {
	r.SetValue("search_channels", value)
}


func (r *SimbaCampaignPlatformUpdateRequest) GetResponse(accessToken string) (*SimbaCampaignPlatformUpdateResponse, []byte, error) {
	var resp SimbaCampaignPlatformUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.platform.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignPlatformUpdateResponse struct {
	CampaignPlatform *CampaignPlatform `json:"campaign_platform"`
}

type SimbaCampaignPlatformUpdateResponseResult struct {
	Response *SimbaCampaignPlatformUpdateResponse `json:"simba_campaign_platform_update_response"`
}





/* 取得一个推广计划的分时折扣设置 */
type SimbaCampaignScheduleGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignScheduleGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignScheduleGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignScheduleGetRequest) GetResponse(accessToken string) (*SimbaCampaignScheduleGetResponse, []byte, error) {
	var resp SimbaCampaignScheduleGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.schedule.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignScheduleGetResponse struct {
	CampaignSchedule *CampaignSchedule `json:"campaign_schedule"`
}

type SimbaCampaignScheduleGetResponseResult struct {
	Response *SimbaCampaignScheduleGetResponse `json:"simba_campaign_schedule_get_response"`
}





/* 更新一个推广计划的分时折扣设置 */
type SimbaCampaignScheduleUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignScheduleUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignScheduleUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比； */
func (r *SimbaCampaignScheduleUpdateRequest) SetSchedule(value string) {
	r.SetValue("schedule", value)
}


func (r *SimbaCampaignScheduleUpdateRequest) GetResponse(accessToken string) (*SimbaCampaignScheduleUpdateResponse, []byte, error) {
	var resp SimbaCampaignScheduleUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.schedule.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignScheduleUpdateResponse struct {
	CampaignSchedule *CampaignSchedule `json:"campaign_schedule"`
}

type SimbaCampaignScheduleUpdateResponseResult struct {
	Response *SimbaCampaignScheduleUpdateResponse `json:"simba_campaign_schedule_update_response"`
}





/* 更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。 */
type SimbaCampaignUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaCampaignUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaCampaignUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 用户设置的上下限状态；offline-下线；online-上线； */
func (r *SimbaCampaignUpdateRequest) SetOnlineStatus(value string) {
	r.SetValue("online_status", value)
}

/* 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。 */
func (r *SimbaCampaignUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *SimbaCampaignUpdateRequest) GetResponse(accessToken string) (*SimbaCampaignUpdateResponse, []byte, error) {
	var resp SimbaCampaignUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaign.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignUpdateResponse struct {
	Campaign *Campaign `json:"campaign"`
}

type SimbaCampaignUpdateResponseResult struct {
	Response *SimbaCampaignUpdateResponse `json:"simba_campaign_update_response"`
}





/* 取得一个客户的推广计划； */
type SimbaCampaignsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCampaignsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCampaignsGetRequest) GetResponse(accessToken string) (*SimbaCampaignsGetResponse, []byte, error) {
	var resp SimbaCampaignsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.campaigns.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCampaignsGetResponse struct {
	Campaigns []*Campaign `json:"campaigns"`
}

type SimbaCampaignsGetResponseResult struct {
	Response *SimbaCampaignsGetResponse `json:"simba_campaigns_get_response"`
}





/* 获取更改过的类目出价ID */
type SimbaCatmatchidsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCatmatchidsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaCatmatchidsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaCatmatchidsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaCatmatchidsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaCatmatchidsChangedGetRequest) GetResponse(accessToken string) (*SimbaCatmatchidsChangedGetResponse, []byte, error) {
	var resp SimbaCatmatchidsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.catmatchids.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCatmatchidsChangedGetResponse struct {
	ChangedCatmatchIds []int `json:"changed_catmatch_ids"`
}

type SimbaCatmatchidsChangedGetResponseResult struct {
	Response *SimbaCatmatchidsChangedGetResponse `json:"simba_catmatchids_changed_get_response"`
}





/* 获取删除的类目出价ID */
type SimbaCatmatchidsDeletedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCatmatchidsDeletedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaCatmatchidsDeletedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaCatmatchidsDeletedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaCatmatchidsDeletedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaCatmatchidsDeletedGetRequest) GetResponse(accessToken string) (*SimbaCatmatchidsDeletedGetResponse, []byte, error) {
	var resp SimbaCatmatchidsDeletedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.catmatchids.deleted.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCatmatchidsDeletedGetResponse struct {
	DeletedCatmatchIds []int `json:"deleted_catmatch_ids"`
}

type SimbaCatmatchidsDeletedGetResponseResult struct {
	Response *SimbaCatmatchidsDeletedGetResponse `json:"simba_catmatchids_deleted_get_response"`
}





/* 创建一个创意 */
type SimbaCreativeAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaCreativeAddRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 创意图片地址，必须是推广组对应商品的图片之一 */
func (r *SimbaCreativeAddRequest) SetImgUrl(value string) {
	r.SetValue("img_url", value)
}

/* 主人昵称 */
func (r *SimbaCreativeAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 创意标题，最多20个汉字 */
func (r *SimbaCreativeAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *SimbaCreativeAddRequest) GetResponse(accessToken string) (*SimbaCreativeAddResponse, []byte, error) {
	var resp SimbaCreativeAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creative.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativeAddResponse struct {
	Creative *Creative `json:"creative"`
}

type SimbaCreativeAddResponseResult struct {
	Response *SimbaCreativeAddResponse `json:"simba_creative_add_response"`
}





/* 删除一个创意 */
type SimbaCreativeDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 创意Id */
func (r *SimbaCreativeDeleteRequest) SetCreativeId(value string) {
	r.SetValue("creative_id", value)
}

/* 主人昵称 */
func (r *SimbaCreativeDeleteRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCreativeDeleteRequest) GetResponse(accessToken string) (*SimbaCreativeDeleteResponse, []byte, error) {
	var resp SimbaCreativeDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creative.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativeDeleteResponse struct {
	Creative *Creative `json:"creative"`
}

type SimbaCreativeDeleteResponseResult struct {
	Response *SimbaCreativeDeleteResponse `json:"simba_creative_delete_response"`
}





/* 更新一个创意的信息，可以设置创意标题、创意图片 */
type SimbaCreativeUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaCreativeUpdateRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 创意Id */
func (r *SimbaCreativeUpdateRequest) SetCreativeId(value string) {
	r.SetValue("creative_id", value)
}

/* 创意图片地址，必须是推广组对应商品的图片之一 */
func (r *SimbaCreativeUpdateRequest) SetImgUrl(value string) {
	r.SetValue("img_url", value)
}

/* 主人昵称 */
func (r *SimbaCreativeUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 创意标题，最多20个汉字 */
func (r *SimbaCreativeUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *SimbaCreativeUpdateRequest) GetResponse(accessToken string) (*SimbaCreativeUpdateResponse, []byte, error) {
	var resp SimbaCreativeUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creative.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativeUpdateResponse struct {
	Creativerecord *CreativeRecord `json:"creativerecord"`
}

type SimbaCreativeUpdateResponseResult struct {
	Response *SimbaCreativeUpdateResponse `json:"simba_creative_update_response"`
}





/* 获取修改的创意ID */
type SimbaCreativeidsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCreativeidsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaCreativeidsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaCreativeidsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaCreativeidsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaCreativeidsChangedGetRequest) GetResponse(accessToken string) (*SimbaCreativeidsChangedGetResponse, []byte, error) {
	var resp SimbaCreativeidsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creativeids.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativeidsChangedGetResponse struct {
	ChangedCreativeIds []int `json:"changed_creative_ids"`
}

type SimbaCreativeidsChangedGetResponseResult struct {
	Response *SimbaCreativeidsChangedGetResponse `json:"simba_creativeids_changed_get_response"`
}





/* 获取删除的创意ID */
type SimbaCreativeidsDeletedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCreativeidsDeletedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaCreativeidsDeletedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaCreativeidsDeletedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到这个时间点之后的数据，不能大于一个月 */
func (r *SimbaCreativeidsDeletedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaCreativeidsDeletedGetRequest) GetResponse(accessToken string) (*SimbaCreativeidsDeletedGetResponse, []byte, error) {
	var resp SimbaCreativeidsDeletedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creativeids.deleted.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativeidsDeletedGetResponse struct {
	DeletedCreativeIds []int `json:"deleted_creative_ids"`
}

type SimbaCreativeidsDeletedGetResponseResult struct {
	Response *SimbaCreativeidsDeletedGetResponse `json:"simba_creativeids_deleted_get_response"`
}





/* 分页获取修改过的广告创意ID和修改时间 */
type SimbaCreativesChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaCreativesChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaCreativesChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaCreativesChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaCreativesChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaCreativesChangedGetRequest) GetResponse(accessToken string) (*SimbaCreativesChangedGetResponse, []byte, error) {
	var resp SimbaCreativesChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creatives.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativesChangedGetResponse struct {
	Creatives *CreativePage `json:"creatives"`
}

type SimbaCreativesChangedGetResponseResult struct {
	Response *SimbaCreativesChangedGetResponse `json:"simba_creatives_changed_get_response"`
}





/* 取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意； 
如果同时提供了推广组Id和创意id列表，则优先使用推广组Id； */
type SimbaCreativesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaCreativesGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 创意Id数组，最多200个 */
func (r *SimbaCreativesGetRequest) SetCreativeIds(value string) {
	r.SetValue("creative_ids", value)
}

/* 主人昵称 */
func (r *SimbaCreativesGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCreativesGetRequest) GetResponse(accessToken string) (*SimbaCreativesGetResponse, []byte, error) {
	var resp SimbaCreativesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creatives.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativesGetResponse struct {
	Creatives []*Creative `json:"creatives"`
}

type SimbaCreativesGetResponseResult struct {
	Response *SimbaCreativesGetResponse `json:"simba_creatives_get_response"`
}





/* 根据一个创意Id列表取得创意对应的修改记录 */
type SimbaCreativesRecordGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 创意Id数组，最多200个 */
func (r *SimbaCreativesRecordGetRequest) SetCreativeIds(value string) {
	r.SetValue("creative_ids", value)
}

/* 主人昵称 */
func (r *SimbaCreativesRecordGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaCreativesRecordGetRequest) GetResponse(accessToken string) (*SimbaCreativesRecordGetResponse, []byte, error) {
	var resp SimbaCreativesRecordGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.creatives.record.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCreativesRecordGetResponse struct {
	Creativerecords []*CreativeRecord `json:"creativerecords"`
}

type SimbaCreativesRecordGetResponseResult struct {
	Response *SimbaCreativesRecordGetResponse `json:"simba_creatives_record_get_response"`
}





/* 取得当前登录用户的授权账户列表 */
type SimbaCustomersAuthorizedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *SimbaCustomersAuthorizedGetRequest) GetResponse(accessToken string) (*SimbaCustomersAuthorizedGetResponse, []byte, error) {
	var resp SimbaCustomersAuthorizedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.customers.authorized.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaCustomersAuthorizedGetResponse struct {
	Nicks []string `json:"nicks"`
}

type SimbaCustomersAuthorizedGetResponseResult struct {
	Response *SimbaCustomersAuthorizedGetResponse `json:"simba_customers_authorized_get_response"`
}





/* 获取类目信息 */
type SimbaInsightCatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询类目id数组，最大长度200 */
func (r *SimbaInsightCatsGetRequest) SetCategoryIds(value string) {
	r.SetValue("category_ids", value)
}

/* 主人昵称 */
func (r *SimbaInsightCatsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaInsightCatsGetRequest) GetResponse(accessToken string) (*SimbaInsightCatsGetResponse, []byte, error) {
	var resp SimbaInsightCatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.cats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatsGetResponse struct {
	InCategoryTops []*INCategoryTop `json:"in_category_tops"`
}

type SimbaInsightCatsGetResponseResult struct {
	Response *SimbaInsightCatsGetResponse `json:"simba_insight_cats_get_response"`
}





/* 类目分析数据查询 */
type SimbaInsightCatsanalysisGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询类目id数组，最大长度200 */
func (r *SimbaInsightCatsanalysisGetRequest) SetCategoryIds(value string) {
	r.SetValue("category_ids", value)
}

/* 主人昵称 */
func (r *SimbaInsightCatsanalysisGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 结果过滤。 
AREA：返回地域占比； 
SOURCE：返回来源占比； 
HPRICE：返回竞价分布。 
stu只能是AREA、SOURCE或HPRICE中的一个 */
func (r *SimbaInsightCatsanalysisGetRequest) SetStu(value string) {
	r.SetValue("stu", value)
}


func (r *SimbaInsightCatsanalysisGetRequest) GetResponse(accessToken string) (*SimbaInsightCatsanalysisGetResponse, []byte, error) {
	var resp SimbaInsightCatsanalysisGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.catsanalysis.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatsanalysisGetResponse struct {
	InCategoryAnalyses []*INCategoryAnalysis `json:"in_category_analyses"`
}

type SimbaInsightCatsanalysisGetResponseResult struct {
	Response *SimbaInsightCatsanalysisGetResponse `json:"simba_insight_catsanalysis_get_response"`
}





/* 类目基础数据查询 */
type SimbaInsightCatsbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询类目id数组，最大长度200 */
func (r *SimbaInsightCatsbaseGetRequest) SetCategoryIds(value string) {
	r.SetValue("category_ids", value)
}

/* 结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合 */
func (r *SimbaInsightCatsbaseGetRequest) SetFilter(value string) {
	r.SetValue("filter", value)
}

/* 主人昵称 */
func (r *SimbaInsightCatsbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 时间格式(DAY: 最近一天； WEEK：最近一周。MONTH：最近一个月。3MONTH：最近三个月) */
func (r *SimbaInsightCatsbaseGetRequest) SetTime(value string) {
	r.SetValue("time", value)
}


func (r *SimbaInsightCatsbaseGetRequest) GetResponse(accessToken string) (*SimbaInsightCatsbaseGetResponse, []byte, error) {
	var resp SimbaInsightCatsbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.catsbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatsbaseGetResponse struct {
	InCategoryBases []*INCategoryBase `json:"in_category_bases"`
}

type SimbaInsightCatsbaseGetResponseResult struct {
	Response *SimbaInsightCatsbaseGetResponse `json:"simba_insight_catsbase_get_response"`
}





/* 类目属性预测 */
type SimbaInsightCatsforecastGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaInsightCatsforecastGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 查询词数组，最大长度200 */
func (r *SimbaInsightCatsforecastGetRequest) SetWords(value string) {
	r.SetValue("words", value)
}


func (r *SimbaInsightCatsforecastGetRequest) GetResponse(accessToken string) (*SimbaInsightCatsforecastGetResponse, []byte, error) {
	var resp SimbaInsightCatsforecastGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.catsforecast.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatsforecastGetResponse struct {
	InCategoryTops []*INCategoryTop `json:"in_category_tops"`
}

type SimbaInsightCatsforecastGetResponseResult struct {
	Response *SimbaInsightCatsforecastGetResponse `json:"simba_insight_catsforecast_get_response"`
}





/* 类目相关词查询 */
type SimbaInsightCatsrelatedwordGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaInsightCatsrelatedwordGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 最大返回数量(1-10) */
func (r *SimbaInsightCatsrelatedwordGetRequest) SetResultNum(value string) {
	r.SetValue("result_num", value)
}

/* 查询词数组，最大长度200 */
func (r *SimbaInsightCatsrelatedwordGetRequest) SetWords(value string) {
	r.SetValue("words", value)
}


func (r *SimbaInsightCatsrelatedwordGetRequest) GetResponse(accessToken string) (*SimbaInsightCatsrelatedwordGetResponse, []byte, error) {
	var resp SimbaInsightCatsrelatedwordGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.catsrelatedword.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatsrelatedwordGetResponse struct {
	RelatedWords []string `json:"related_words"`
}

type SimbaInsightCatsrelatedwordGetResponseResult struct {
	Response *SimbaInsightCatsrelatedwordGetResponse `json:"simba_insight_catsrelatedword_get_response"`
}





/* 类目TOP词查询 */
type SimbaInsightCatstopwordGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 类目id数组，最大长度200 */
func (r *SimbaInsightCatstopwordGetRequest) SetCategoryIds(value string) {
	r.SetValue("category_ids", value)
}

/* 主人昵称 */
func (r *SimbaInsightCatstopwordGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 最大返回数量(1-100) */
func (r *SimbaInsightCatstopwordGetRequest) SetResultNum(value string) {
	r.SetValue("result_num", value)
}


func (r *SimbaInsightCatstopwordGetRequest) GetResponse(accessToken string) (*SimbaInsightCatstopwordGetResponse, []byte, error) {
	var resp SimbaInsightCatstopwordGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.catstopword.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightCatstopwordGetResponse struct {
	TopWords []string `json:"top_words"`
}

type SimbaInsightCatstopwordGetResponseResult struct {
	Response *SimbaInsightCatstopwordGetResponse `json:"simba_insight_catstopword_get_response"`
}





/* 获取一级类目 */
type SimbaInsightToplevelcatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaInsightToplevelcatsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaInsightToplevelcatsGetRequest) GetResponse(accessToken string) (*SimbaInsightToplevelcatsGetResponse, []byte, error) {
	var resp SimbaInsightToplevelcatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.toplevelcats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightToplevelcatsGetResponse struct {
	InCategoryTops []*INCategoryTop `json:"in_category_tops"`
}

type SimbaInsightToplevelcatsGetResponseResult struct {
	Response *SimbaInsightToplevelcatsGetResponse `json:"simba_insight_toplevelcats_get_response"`
}





/* 词分析数据查询 */
type SimbaInsightWordsanalysisGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaInsightWordsanalysisGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 结果过滤。 
AREA：  返回地域占比； 
SOURCE：返回来源占比； 
HPRICE：返回竞价分布。 
stu只能是AREA、SOURCE或HPRICE中的一个 */
func (r *SimbaInsightWordsanalysisGetRequest) SetStu(value string) {
	r.SetValue("stu", value)
}

/* 查询词组，最大长度200 */
func (r *SimbaInsightWordsanalysisGetRequest) SetWords(value string) {
	r.SetValue("words", value)
}


func (r *SimbaInsightWordsanalysisGetRequest) GetResponse(accessToken string) (*SimbaInsightWordsanalysisGetResponse, []byte, error) {
	var resp SimbaInsightWordsanalysisGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.wordsanalysis.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightWordsanalysisGetResponse struct {
	InWordAnalyses []*INWordAnalysis `json:"in_word_analyses"`
}

type SimbaInsightWordsanalysisGetResponseResult struct {
	Response *SimbaInsightWordsanalysisGetResponse `json:"simba_insight_wordsanalysis_get_response"`
}





/* 词基础数据查询 */
type SimbaInsightWordsbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合 */
func (r *SimbaInsightWordsbaseGetRequest) SetFilter(value string) {
	r.SetValue("filter", value)
}

/* 主人昵称 */
func (r *SimbaInsightWordsbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 时间格式(DAY: 最近一天； WEEK：最近一周。MONTH：最近一个月。3MONTH：最近三个月) */
func (r *SimbaInsightWordsbaseGetRequest) SetTime(value string) {
	r.SetValue("time", value)
}

/* 查询词组，最大长度170 */
func (r *SimbaInsightWordsbaseGetRequest) SetWords(value string) {
	r.SetValue("words", value)
}


func (r *SimbaInsightWordsbaseGetRequest) GetResponse(accessToken string) (*SimbaInsightWordsbaseGetResponse, []byte, error) {
	var resp SimbaInsightWordsbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.wordsbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightWordsbaseGetResponse struct {
	InWordBases []*INWordBase `json:"in_word_bases"`
}

type SimbaInsightWordsbaseGetResponseResult struct {
	Response *SimbaInsightWordsbaseGetResponse `json:"simba_insight_wordsbase_get_response"`
}





/* 词和类目查询 */
type SimbaInsightWordscatsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合 */
func (r *SimbaInsightWordscatsGetRequest) SetFilter(value string) {
	r.SetValue("filter", value)
}

/* 主人昵称 */
func (r *SimbaInsightWordscatsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 查询词和类目数组，最大长度200，每一个数组元素都是词+类目，以”^^”分割如下： 
词^^类目 */
func (r *SimbaInsightWordscatsGetRequest) SetWordCategories(value string) {
	r.SetValue("word_categories", value)
}


func (r *SimbaInsightWordscatsGetRequest) GetResponse(accessToken string) (*SimbaInsightWordscatsGetResponse, []byte, error) {
	var resp SimbaInsightWordscatsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.insight.wordscats.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaInsightWordscatsGetResponse struct {
	InWordCategories []*INWordCategory `json:"in_word_categories"`
}

type SimbaInsightWordscatsGetResponseResult struct {
	Response *SimbaInsightWordscatsGetResponse `json:"simba_insight_wordscats_get_response"`
}





/* 根据词ID和给定的出价获取词的预估信息 */
type SimbaKeywordKeywordforecastGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 词的出价,范围在5-9999之间,单位分 */
func (r *SimbaKeywordKeywordforecastGetRequest) SetBidwordPrice(value string) {
	r.SetValue("bidword_price", value)
}

/* 词ID */
func (r *SimbaKeywordKeywordforecastGetRequest) SetKeywordId(value string) {
	r.SetValue("keyword_id", value)
}

/* 经典名表行 */
func (r *SimbaKeywordKeywordforecastGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordKeywordforecastGetRequest) GetResponse(accessToken string) (*SimbaKeywordKeywordforecastGetResponse, []byte, error) {
	var resp SimbaKeywordKeywordforecastGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keyword.keywordforecast.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordKeywordforecastGetResponse struct {
	KeywordForecast *KeywordForecast `json:"keyword_forecast"`
}

type SimbaKeywordKeywordforecastGetResponseResult struct {
	Response *SimbaKeywordKeywordforecastGetResponse `json:"simba_keyword_keywordforecast_get_response"`
}





/* 获取修改的词ID */
type SimbaKeywordidsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaKeywordidsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaKeywordidsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaKeywordidsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaKeywordidsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaKeywordidsChangedGetRequest) GetResponse(accessToken string) (*SimbaKeywordidsChangedGetResponse, []byte, error) {
	var resp SimbaKeywordidsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordids.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordidsChangedGetResponse struct {
	ChangedKeywordIds []int `json:"changed_keyword_ids"`
}

type SimbaKeywordidsChangedGetResponseResult struct {
	Response *SimbaKeywordidsChangedGetResponse `json:"simba_keywordids_changed_get_response"`
}





/* 获取删除的词ID */
type SimbaKeywordidsDeletedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaKeywordidsDeletedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaKeywordidsDeletedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaKeywordidsDeletedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaKeywordidsDeletedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaKeywordidsDeletedGetRequest) GetResponse(accessToken string) (*SimbaKeywordidsDeletedGetResponse, []byte, error) {
	var resp SimbaKeywordidsDeletedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordids.deleted.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordidsDeletedGetResponse struct {
	DeletedKeywordIds []int `json:"deleted_keyword_ids"`
}

type SimbaKeywordidsDeletedGetResponseResult struct {
	Response *SimbaKeywordidsDeletedGetResponse `json:"simba_keywordids_deleted_get_response"`
}





/* 分页获取修改过的关键词ID、宝贝id、修改时间 */
type SimbaKeywordsChangedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaKeywordsChangedGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaKeywordsChangedGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,默认200最大1000 */
func (r *SimbaKeywordsChangedGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 得到此时间点之后的数据，不能大于一个月 */
func (r *SimbaKeywordsChangedGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}


func (r *SimbaKeywordsChangedGetRequest) GetResponse(accessToken string) (*SimbaKeywordsChangedGetResponse, []byte, error) {
	var resp SimbaKeywordsChangedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.changed.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsChangedGetResponse struct {
	Keywords *KeywordPage `json:"keywords"`
}

type SimbaKeywordsChangedGetResponseResult struct {
	Response *SimbaKeywordsChangedGetResponse `json:"simba_keywords_changed_get_response"`
}





/* 删除一批关键词 */
type SimbaKeywordsDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划Id */
func (r *SimbaKeywordsDeleteRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 关键词Id数组，最多100个 */
func (r *SimbaKeywordsDeleteRequest) SetKeywordIds(value string) {
	r.SetValue("keyword_ids", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsDeleteRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsDeleteRequest) GetResponse(accessToken string) (*SimbaKeywordsDeleteResponse, []byte, error) {
	var resp SimbaKeywordsDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsDeleteResponse struct {
	Keywords []*Keyword `json:"keywords"`
}

type SimbaKeywordsDeleteResponseResult struct {
	Response *SimbaKeywordsDeleteResponse `json:"simba_keywords_delete_response"`
}





/* 设置一批关键词的出价 */
type SimbaKeywordsPricevonSetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0，1(0代表不使用，1代表使用)。matchscope只能是1,2,4（1代表精确匹配，2代表子串匹配，4代表广泛匹配） */
func (r *SimbaKeywordsPricevonSetRequest) SetKeywordidPrices(value string) {
	r.SetValue("keywordid_prices", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsPricevonSetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsPricevonSetRequest) GetResponse(accessToken string) (*SimbaKeywordsPricevonSetResponse, []byte, error) {
	var resp SimbaKeywordsPricevonSetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.pricevon.set", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsPricevonSetResponse struct {
	Keywords []*Keyword `json:"keywords"`
}

type SimbaKeywordsPricevonSetResponseResult struct {
	Response *SimbaKeywordsPricevonSetResponse `json:"simba_keywords_pricevon_set_response"`
}





/* 取得一个推广组的所有关键词的质量得分列表 */
type SimbaKeywordsQscoreGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaKeywordsQscoreGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsQscoreGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsQscoreGetRequest) GetResponse(accessToken string) (*SimbaKeywordsQscoreGetResponse, []byte, error) {
	var resp SimbaKeywordsQscoreGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.qscore.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsQscoreGetResponse struct {
	KeywordQscoreList []*KeywordQscore `json:"keyword_qscore_list"`
}

type SimbaKeywordsQscoreGetResponseResult struct {
	Response *SimbaKeywordsQscoreGetResponse `json:"simba_keywords_qscore_get_response"`
}





/* 取得一个推广组的推荐关键词列表 */
type SimbaKeywordsRecommendGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID */
func (r *SimbaKeywordsRecommendGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsRecommendGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 排序方式:  
搜索量 search_volume  
市场平均价格 average_price  
相关度 relevance  
不排序 non  
默认为 non */
func (r *SimbaKeywordsRecommendGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 返回的第几页数据，默认为1 */
func (r *SimbaKeywordsRecommendGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 返回的每页数据量大小,最大200 */
func (r *SimbaKeywordsRecommendGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 相关度 */
func (r *SimbaKeywordsRecommendGetRequest) SetPertinence(value string) {
	r.SetValue("pertinence", value)
}

/* 搜索量,设置此值后返回的就是大于此搜索量的词列表 */
func (r *SimbaKeywordsRecommendGetRequest) SetSearch(value string) {
	r.SetValue("search", value)
}


func (r *SimbaKeywordsRecommendGetRequest) GetResponse(accessToken string) (*SimbaKeywordsRecommendGetResponse, []byte, error) {
	var resp SimbaKeywordsRecommendGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywords.recommend.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsRecommendGetResponse struct {
	RecommendWords *RecommendWordPage `json:"recommend_words"`
}

type SimbaKeywordsRecommendGetResponseResult struct {
	Response *SimbaKeywordsRecommendGetResponse `json:"simba_keywords_recommend_get_response"`
}





/* 取得一个推广组的所有关键词 */
type SimbaKeywordsbyadgroupidGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaKeywordsbyadgroupidGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsbyadgroupidGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsbyadgroupidGetRequest) GetResponse(accessToken string) (*SimbaKeywordsbyadgroupidGetResponse, []byte, error) {
	var resp SimbaKeywordsbyadgroupidGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordsbyadgroupid.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsbyadgroupidGetResponse struct {
	Keywords []*Keyword `json:"keywords"`
}

type SimbaKeywordsbyadgroupidGetResponseResult struct {
	Response *SimbaKeywordsbyadgroupidGetResponse `json:"simba_keywordsbyadgroupid_get_response"`
}





/* 根据一个关键词Id列表取得一组关键词 */
type SimbaKeywordsbykeywordidsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 关键词Id数组，最多200个； */
func (r *SimbaKeywordsbykeywordidsGetRequest) SetKeywordIds(value string) {
	r.SetValue("keyword_ids", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsbykeywordidsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsbykeywordidsGetRequest) GetResponse(accessToken string) (*SimbaKeywordsbykeywordidsGetResponse, []byte, error) {
	var resp SimbaKeywordsbykeywordidsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordsbykeywordids.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsbykeywordidsGetResponse struct {
	Keywords []*Keyword `json:"keywords"`
}

type SimbaKeywordsbykeywordidsGetResponseResult struct {
	Response *SimbaKeywordsbykeywordidsGetResponse `json:"simba_keywordsbykeywordids_get_response"`
}





/* 取得一个推广组的所有关键词和类目出价的质量得分列表 */
type SimbaKeywordscatQscoreGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组Id */
func (r *SimbaKeywordscatQscoreGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 主人昵称 */
func (r *SimbaKeywordscatQscoreGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordscatQscoreGetRequest) GetResponse(accessToken string) (*SimbaKeywordscatQscoreGetResponse, []byte, error) {
	var resp SimbaKeywordscatQscoreGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordscat.qscore.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordscatQscoreGetResponse struct {
	Qscore *Qscore `json:"qscore"`
}

type SimbaKeywordscatQscoreGetResponseResult struct {
	Response *SimbaKeywordscatQscoreGetResponse `json:"simba_keywordscat_qscore_get_response"`
}





/* 创建一批关键词 */
type SimbaKeywordsvonAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaKeywordsvonAddRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 关键词，出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0，1(0代表不使用，1代表使用)。matchscope只能是1,2,4（1代表精确匹配，2代表子串匹配，4代表广泛匹配）。 */
func (r *SimbaKeywordsvonAddRequest) SetKeywordPrices(value string) {
	r.SetValue("keyword_prices", value)
}

/* 主人昵称 */
func (r *SimbaKeywordsvonAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaKeywordsvonAddRequest) GetResponse(accessToken string) (*SimbaKeywordsvonAddResponse, []byte, error) {
	var resp SimbaKeywordsvonAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.keywordsvon.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaKeywordsvonAddResponse struct {
	Keywords []*Keyword `json:"keywords"`
}

type SimbaKeywordsvonAddResponseResult struct {
	Response *SimbaKeywordsvonAddResponse `json:"simba_keywordsvon_add_response"`
}





/* 获取登陆权限签名 */
type SimbaLoginAuthsignGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 主人昵称 */
func (r *SimbaLoginAuthsignGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaLoginAuthsignGetRequest) GetResponse(accessToken string) (*SimbaLoginAuthsignGetResponse, []byte, error) {
	var resp SimbaLoginAuthsignGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.login.authsign.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaLoginAuthsignGetResponse struct {
	SubwayToken string `json:"subway_token"`
}

type SimbaLoginAuthsignGetResponseResult struct {
	Response *SimbaLoginAuthsignGetResponse `json:"simba_login_authsign_get_response"`
}





/* 批量推广组添加定向推广投放位置，出价使用默认出价 */
type SimbaNonsearchAdgroupplacesAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 宝贝id投放位置id json数组字符串，数组个数最多200个。其中json数组中的key必须和对应实体AdGroupPlace中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值 */
func (r *SimbaNonsearchAdgroupplacesAddRequest) SetAdgroupPlacesJson(value string) {
	r.SetValue("adgroup_places_json", value)
}

/* 12345 */
func (r *SimbaNonsearchAdgroupplacesAddRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchAdgroupplacesAddRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchAdgroupplacesAddRequest) GetResponse(accessToken string) (*SimbaNonsearchAdgroupplacesAddResponse, []byte, error) {
	var resp SimbaNonsearchAdgroupplacesAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.adgroupplaces.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAdgroupplacesAddResponse struct {
	AdgroupPlaceList []*ADGroupPlace `json:"adgroup_place_list"`
}

type SimbaNonsearchAdgroupplacesAddResponseResult struct {
	Response *SimbaNonsearchAdgroupplacesAddResponse `json:"simba_nonsearch_adgroupplaces_add_response"`
}





/* 批量删除推广组定向推广投放位置 */
type SimbaNonsearchAdgroupplacesDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id，投放位置id  json数组字符串，数组个数最多200个。其中json数组中的key必须和对应实体AdGroupPlace中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值 */
func (r *SimbaNonsearchAdgroupplacesDeleteRequest) SetAdgroupPlacesJson(value string) {
	r.SetValue("adgroup_places_json", value)
}

/* 推广计划ID */
func (r *SimbaNonsearchAdgroupplacesDeleteRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchAdgroupplacesDeleteRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchAdgroupplacesDeleteRequest) GetResponse(accessToken string) (*SimbaNonsearchAdgroupplacesDeleteResponse, []byte, error) {
	var resp SimbaNonsearchAdgroupplacesDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.adgroupplaces.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAdgroupplacesDeleteResponse struct {
	AdgroupPlaceList []*ADGroupPlace `json:"adgroup_place_list"`
}

type SimbaNonsearchAdgroupplacesDeleteResponseResult struct {
	Response *SimbaNonsearchAdgroupplacesDeleteResponse `json:"simba_nonsearch_adgroupplaces_delete_response"`
}





/* 根据指定推广计划下推广组列表获取相应推广组投放位置包括通投位置和单独出价位置 */
type SimbaNonsearchAdgroupplacesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID数组 */
func (r *SimbaNonsearchAdgroupplacesGetRequest) SetAdgroupIds(value string) {
	r.SetValue("adgroup_ids", value)
}

/* 推广计划ID */
func (r *SimbaNonsearchAdgroupplacesGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchAdgroupplacesGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchAdgroupplacesGetRequest) GetResponse(accessToken string) (*SimbaNonsearchAdgroupplacesGetResponse, []byte, error) {
	var resp SimbaNonsearchAdgroupplacesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.adgroupplaces.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAdgroupplacesGetResponse struct {
	AdgroupPlaceList []*ADGroupPlace `json:"adgroup_place_list"`
}

type SimbaNonsearchAdgroupplacesGetResponseResult struct {
	Response *SimbaNonsearchAdgroupplacesGetResponse `json:"simba_nonsearch_adgroupplaces_get_response"`
}





/* 批量修改推广组定向推广投放位置价格 */
type SimbaNonsearchAdgroupplacesUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id，投放位置id，出价，是否默认出价 json数组字符串，数组个数最多200个。是否默认出价不能为空, 出价以“分”为单位，不能小于5分，不能大于最高日限额，不能大于9999分。 
json数组中的key必须和对应实体AdGroupPlace中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值 */
func (r *SimbaNonsearchAdgroupplacesUpdateRequest) SetAdgroupPlacesJson(value string) {
	r.SetValue("adgroup_places_json", value)
}

/* 推广计划ID */
func (r *SimbaNonsearchAdgroupplacesUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchAdgroupplacesUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchAdgroupplacesUpdateRequest) GetResponse(accessToken string) (*SimbaNonsearchAdgroupplacesUpdateResponse, []byte, error) {
	var resp SimbaNonsearchAdgroupplacesUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.adgroupplaces.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAdgroupplacesUpdateResponse struct {
	AdgroupPlaceList []*ADGroupPlace `json:"adgroup_place_list"`
}

type SimbaNonsearchAdgroupplacesUpdateResponseResult struct {
	Response *SimbaNonsearchAdgroupplacesUpdateResponse `json:"simba_nonsearch_adgroupplaces_update_response"`
}





/* 获取定向投放人群维度列表 */
type SimbaNonsearchAlldemographicsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *SimbaNonsearchAlldemographicsGetRequest) GetResponse(accessToken string) (*SimbaNonsearchAlldemographicsGetResponse, []byte, error) {
	var resp SimbaNonsearchAlldemographicsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.alldemographics.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAlldemographicsGetResponse struct {
	DemographicList []*Demographic `json:"demographic_list"`
}

type SimbaNonsearchAlldemographicsGetResponseResult struct {
	Response *SimbaNonsearchAlldemographicsGetResponse `json:"simba_nonsearch_alldemographics_get_response"`
}





/* 获取单独出价投放位置列表 */
type SimbaNonsearchAllplacesGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *SimbaNonsearchAllplacesGetRequest) GetResponse(accessToken string) (*SimbaNonsearchAllplacesGetResponse, []byte, error) {
	var resp SimbaNonsearchAllplacesGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.allplaces.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchAllplacesGetResponse struct {
	PlaceList []*Place `json:"place_list"`
}

type SimbaNonsearchAllplacesGetResponseResult struct {
	Response *SimbaNonsearchAllplacesGetResponse `json:"simba_nonsearch_allplaces_get_response"`
}





/* 获取给定campaign设置的投放人群维度列表 */
type SimbaNonsearchDemographicsGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划ID */
func (r *SimbaNonsearchDemographicsGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchDemographicsGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchDemographicsGetRequest) GetResponse(accessToken string) (*SimbaNonsearchDemographicsGetResponse, []byte, error) {
	var resp SimbaNonsearchDemographicsGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.demographics.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchDemographicsGetResponse struct {
	DemographicSettingList []*DemographicSetting `json:"demographic_setting_list"`
}

type SimbaNonsearchDemographicsGetResponseResult struct {
	Response *SimbaNonsearchDemographicsGetResponse `json:"simba_nonsearch_demographics_get_response"`
}





/* 设置投放人群维度加价，如果给定的campagin没有设置给定的人群维度则添加给定的人群维度，如果给定的campaign已存在给定的人群维度则修改加价 */
type SimbaNonsearchDemographicsUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划ID */
func (r *SimbaNonsearchDemographicsUpdateRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 投放人群维度Id，加价json数组字符串。数组长度不能超过15，json数组中的key必须和相应实体DemographicSetting中的属性字段保持一致，否则对应的实体对象属性获取不到相应的值 
incrementalPrice是整数，以“分”为单位，不能小于1，不能大于日限额,不能大于9999分; 可以为0表示不加价；投放人群维度ID必须有效 */
func (r *SimbaNonsearchDemographicsUpdateRequest) SetDemographicIdPriceJson(value string) {
	r.SetValue("demographic_id_price_json", value)
}

/* 主人昵称 */
func (r *SimbaNonsearchDemographicsUpdateRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaNonsearchDemographicsUpdateRequest) GetResponse(accessToken string) (*SimbaNonsearchDemographicsUpdateResponse, []byte, error) {
	var resp SimbaNonsearchDemographicsUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.nonsearch.demographics.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaNonsearchDemographicsUpdateResponse struct {
	DemographicSettingList []*DemographicSetting `json:"demographic_setting_list"`
}

type SimbaNonsearchDemographicsUpdateResponseResult struct {
	Response *SimbaNonsearchDemographicsUpdateResponse `json:"simba_nonsearch_demographics_update_response"`
}





/* 推广组基础报表数据对象 */
type SimbaRptAdgroupbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupbaseGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划id */
func (r *SimbaRptAdgroupbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 
定向投放：NOSEARCH）可以一次取多个例如：SEARCH,CAT */
func (r *SimbaRptAdgroupbaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2 */
func (r *SimbaRptAdgroupbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupbaseGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupbaseGetResponse, []byte, error) {
	var resp SimbaRptAdgroupbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupbaseGetResponse struct {
	RptAdgroupBaseList string `json:"rpt_adgroup_base_list"`
}

type SimbaRptAdgroupbaseGetResponseResult struct {
	Response *SimbaRptAdgroupbaseGetResponse `json:"simba_rpt_adgroupbase_get_response"`
}





/* 推广组下创意报表基础数据查询(汇总数据，不分类型) */
type SimbaRptAdgroupcreativebaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 查询推广计划id */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束日期，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始日期，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限验证信息 */
func (r *SimbaRptAdgroupcreativebaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupcreativebaseGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupcreativebaseGetResponse, []byte, error) {
	var resp SimbaRptAdgroupcreativebaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupcreativebase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupcreativebaseGetResponse struct {
	RptAdgroupcreativeBaseList string `json:"rpt_adgroupcreative_base_list"`
}

type SimbaRptAdgroupcreativebaseGetResponseResult struct {
	Response *SimbaRptAdgroupcreativebaseGetResponse `json:"simba_rpt_adgroupcreativebase_get_response"`
}





/* 推广组下的创意报表效果数据查询(汇总数据，不分类型) */
type SimbaRptAdgroupcreativeeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 查询推广计划id */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束日期，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始日期，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限验证信息 */
func (r *SimbaRptAdgroupcreativeeffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupcreativeeffectGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupcreativeeffectGetResponse, []byte, error) {
	var resp SimbaRptAdgroupcreativeeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupcreativeeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupcreativeeffectGetResponse struct {
	RptAdgroupcreativeEffectList string `json:"rpt_adgroupcreative_effect_list"`
}

type SimbaRptAdgroupcreativeeffectGetResponseResult struct {
	Response *SimbaRptAdgroupcreativeeffectGetResponse `json:"simba_rpt_adgroupcreativeeffect_get_response"`
}





/* 推广组效果报表数据对象 */
type SimbaRptAdgroupeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupeffectGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划id */
func (r *SimbaRptAdgroupeffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupeffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupeffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupeffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupeffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 
定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT */
func (r *SimbaRptAdgroupeffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2 */
func (r *SimbaRptAdgroupeffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupeffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupeffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupeffectGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupeffectGetResponse, []byte, error) {
	var resp SimbaRptAdgroupeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupeffectGetResponse struct {
	RptAdgroupEffectList string `json:"rpt_adgroup_effect_list"`
}

type SimbaRptAdgroupeffectGetResponseResult struct {
	Response *SimbaRptAdgroupeffectGetResponse `json:"simba_rpt_adgroupeffect_get_response"`
}





/* 推广组下的词基础报表数据查询(明细数据不分类型查询) */
type SimbaRptAdgroupkeywordbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划ID */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 主人昵称 */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupkeywordbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupkeywordbaseGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupkeywordbaseGetResponse, []byte, error) {
	var resp SimbaRptAdgroupkeywordbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupkeywordbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupkeywordbaseGetResponse struct {
	RptAdgroupkeywordBaseList string `json:"rpt_adgroupkeyword_base_list"`
}

type SimbaRptAdgroupkeywordbaseGetResponseResult struct {
	Response *SimbaRptAdgroupkeywordbaseGetResponse `json:"simba_rpt_adgroupkeywordbase_get_response"`
}





/* 推广组下的词效果报表数据查询(明细数据不分类型查询) */
type SimbaRptAdgroupkeywordeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组ID */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划ID */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 主人昵称 */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupkeywordeffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupkeywordeffectGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupkeywordeffectGetResponse, []byte, error) {
	var resp SimbaRptAdgroupkeywordeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupkeywordeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupkeywordeffectGetResponse struct {
	RptAdgroupkeywordEffectList string `json:"rpt_adgroupkeyword_effect_list"`
}

type SimbaRptAdgroupkeywordeffectGetResponseResult struct {
	Response *SimbaRptAdgroupkeywordeffectGetResponse `json:"simba_rpt_adgroupkeywordeffect_get_response"`
}





/* 推广组下的定向推广基础数据查询 */
type SimbaRptAdgroupnonsearchbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划id */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间,格式为：yyyy-mm-dd */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间,格式为：yyyy-mm-dd */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupnonsearchbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupnonsearchbaseGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupnonsearchbaseGetResponse, []byte, error) {
	var resp SimbaRptAdgroupnonsearchbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupnonsearchbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupnonsearchbaseGetResponse struct {
	RptAdgroupNonsearchBase string `json:"rpt_adgroup_nonsearch_base"`
}

type SimbaRptAdgroupnonsearchbaseGetResponseResult struct {
	Response *SimbaRptAdgroupnonsearchbaseGetResponse `json:"simba_rpt_adgroupnonsearchbase_get_response"`
}





/* 推广组下的定向推广效果数据查询 */
type SimbaRptAdgroupnonsearcheffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广组id */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetAdgroupId(value string) {
	r.SetValue("adgroup_id", value)
}

/* 推广计划id */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间,格式为：yyyy-mm-dd */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间,格式为：yyyy-mm-dd */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptAdgroupnonsearcheffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptAdgroupnonsearcheffectGetRequest) GetResponse(accessToken string) (*SimbaRptAdgroupnonsearcheffectGetResponse, []byte, error) {
	var resp SimbaRptAdgroupnonsearcheffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.adgroupnonsearcheffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptAdgroupnonsearcheffectGetResponse struct {
	RptAdgroupNonsearchEffect string `json:"rpt_adgroup_nonsearch_effect"`
}

type SimbaRptAdgroupnonsearcheffectGetResponseResult struct {
	Response *SimbaRptAdgroupnonsearcheffectGetResponse `json:"simba_rpt_adgroupnonsearcheffect_get_response"`
}





/* 推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) */
type SimbaRptCampadgroupbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询推广计划id */
func (r *SimbaRptCampadgroupbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束日期,格式yyyy-mm-dd */
func (r *SimbaRptCampadgroupbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptCampadgroupbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCampadgroupbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCampadgroupbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT */
func (r *SimbaRptCampadgroupbaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptCampadgroupbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始日期,格式yyyy-mm-dd */
func (r *SimbaRptCampadgroupbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限验证信息 */
func (r *SimbaRptCampadgroupbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCampadgroupbaseGetRequest) GetResponse(accessToken string) (*SimbaRptCampadgroupbaseGetResponse, []byte, error) {
	var resp SimbaRptCampadgroupbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.campadgroupbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCampadgroupbaseGetResponse struct {
	RptCampadgroupBaseList string `json:"rpt_campadgroup_base_list"`
}

type SimbaRptCampadgroupbaseGetResponseResult struct {
	Response *SimbaRptCampadgroupbaseGetResponse `json:"simba_rpt_campadgroupbase_get_response"`
}





/* 推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型) */
type SimbaRptCampadgroupeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询推广计划id */
func (r *SimbaRptCampadgroupeffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束日期，格式yyyy-mm-dd */
func (r *SimbaRptCampadgroupeffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptCampadgroupeffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCampadgroupeffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCampadgroupeffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT */
func (r *SimbaRptCampadgroupeffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptCampadgroupeffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始日期，格式yyyy-mm-dd */
func (r *SimbaRptCampadgroupeffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限验证信息 */
func (r *SimbaRptCampadgroupeffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCampadgroupeffectGetRequest) GetResponse(accessToken string) (*SimbaRptCampadgroupeffectGetResponse, []byte, error) {
	var resp SimbaRptCampadgroupeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.campadgroupeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCampadgroupeffectGetResponse struct {
	RptCampadgroupEffectList string `json:"rpt_campadgroup_effect_list"`
}

type SimbaRptCampadgroupeffectGetResponseResult struct {
	Response *SimbaRptCampadgroupeffectGetResponse `json:"simba_rpt_campadgroupeffect_get_response"`
}





/* 推广计划报表基础数据对象 */
type SimbaRptCampaignbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划id */
func (r *SimbaRptCampaignbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptCampaignbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptCampaignbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCampaignbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCampaignbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT */
func (r *SimbaRptCampaignbaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2 */
func (r *SimbaRptCampaignbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptCampaignbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptCampaignbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCampaignbaseGetRequest) GetResponse(accessToken string) (*SimbaRptCampaignbaseGetResponse, []byte, error) {
	var resp SimbaRptCampaignbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.campaignbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCampaignbaseGetResponse struct {
	RptCampaignBaseList string `json:"rpt_campaign_base_list"`
}

type SimbaRptCampaignbaseGetResponseResult struct {
	Response *SimbaRptCampaignbaseGetResponse `json:"simba_rpt_campaignbase_get_response"`
}





/* 推广计划效果报表数据对象 */
type SimbaRptCampaigneffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划id */
func (r *SimbaRptCampaigneffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptCampaigneffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptCampaigneffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCampaigneffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCampaigneffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 报表类型（搜索：SEARCH,类目出价：CAT, 
定向投放：NOSEARCH 全部：ALL）可以一次取多个例如：SEARCH,CAT */
func (r *SimbaRptCampaigneffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源（站内：1，站外：2）可多选以逗号分隔，默认值为：1,2 */
func (r *SimbaRptCampaigneffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptCampaigneffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptCampaigneffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCampaigneffectGetRequest) GetResponse(accessToken string) (*SimbaRptCampaigneffectGetResponse, []byte, error) {
	var resp SimbaRptCampaigneffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.campaigneffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCampaigneffectGetResponse struct {
	RptCampaignEffectList string `json:"rpt_campaign_effect_list"`
}

type SimbaRptCampaigneffectGetResponseResult struct {
	Response *SimbaRptCampaigneffectGetResponse `json:"simba_rpt_campaigneffect_get_response"`
}





/* 客户账户报表基础数据对象 */
type SimbaRptCustbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束日期，格式yyyy-mm-dd */
func (r *SimbaRptCustbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptCustbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCustbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCustbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptCustbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始日期，格式yyyy-mm-dd */
func (r *SimbaRptCustbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限验证信息 */
func (r *SimbaRptCustbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCustbaseGetRequest) GetResponse(accessToken string) (*SimbaRptCustbaseGetResponse, []byte, error) {
	var resp SimbaRptCustbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.custbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCustbaseGetResponse struct {
	RptCustBaseList string `json:"rpt_cust_base_list"`
}

type SimbaRptCustbaseGetResponseResult struct {
	Response *SimbaRptCustbaseGetResponse `json:"simba_rpt_custbase_get_response"`
}





/* 用户账户报表效果数据查询（只有汇总数据，无分类数据） */
type SimbaRptCusteffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 结束时间，格式yyyy-mm-dd */
func (r *SimbaRptCusteffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 主人昵称 */
func (r *SimbaRptCusteffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptCusteffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptCusteffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2 */
func (r *SimbaRptCusteffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 开始时间，格式yyyy-mm-dd */
func (r *SimbaRptCusteffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptCusteffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptCusteffectGetRequest) GetResponse(accessToken string) (*SimbaRptCusteffectGetResponse, []byte, error) {
	var resp SimbaRptCusteffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.custeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptCusteffectGetResponse struct {
	RptCustEffectList string `json:"rpt_cust_effect_list"`
}

type SimbaRptCusteffectGetResponseResult struct {
	Response *SimbaRptCusteffectGetResponse `json:"simba_rpt_custeffect_get_response"`
}





/* 推广计划下的人群基础数据查询 */
type SimbaRptDemographicbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划id */
func (r *SimbaRptDemographicbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间,格式为：yyyy-mm-dd */
func (r *SimbaRptDemographicbaseGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptDemographicbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptDemographicbaseGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptDemographicbaseGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间,格式为：yyyy-mm-dd */
func (r *SimbaRptDemographicbaseGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptDemographicbaseGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptDemographicbaseGetRequest) GetResponse(accessToken string) (*SimbaRptDemographicbaseGetResponse, []byte, error) {
	var resp SimbaRptDemographicbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.demographicbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptDemographicbaseGetResponse struct {
	RptDemographicBase string `json:"rpt_demographic_base"`
}

type SimbaRptDemographicbaseGetResponseResult struct {
	Response *SimbaRptDemographicbaseGetResponse `json:"simba_rpt_demographicbase_get_response"`
}





/* 推广计划下的人群维度效果数据查询 */
type SimbaRptDemographiceffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 推广计划id */
func (r *SimbaRptDemographiceffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 结束时间,格式为：yyyy-mm-dd */
func (r *SimbaRptDemographiceffectGetRequest) SetEndTime(value string) {
	r.SetValue("end_time", value)
}

/* 昵称 */
func (r *SimbaRptDemographiceffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 页码 */
func (r *SimbaRptDemographiceffectGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页大小 */
func (r *SimbaRptDemographiceffectGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 开始时间,格式为：yyyy-mm-dd */
func (r *SimbaRptDemographiceffectGetRequest) SetStartTime(value string) {
	r.SetValue("start_time", value)
}

/* 权限校验参数 */
func (r *SimbaRptDemographiceffectGetRequest) SetSubwayToken(value string) {
	r.SetValue("subway_token", value)
}


func (r *SimbaRptDemographiceffectGetRequest) GetResponse(accessToken string) (*SimbaRptDemographiceffectGetResponse, []byte, error) {
	var resp SimbaRptDemographiceffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.rpt.demographiceffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaRptDemographiceffectGetResponse struct {
	RptDemographicEffect string `json:"rpt_demographic_effect"`
}

type SimbaRptDemographiceffectGetResponseResult struct {
	Response *SimbaRptDemographiceffectGetResponse `json:"simba_rpt_demographiceffect_get_response"`
}





/* 取得一个关键词的推广组排名列表 */
type SimbaToolsItemsTopGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 输入的必须是一个符合ipv4或者ipv6格式的IP地址 */
func (r *SimbaToolsItemsTopGetRequest) SetIp(value string) {
	r.SetValue("ip", value)
}

/* 关键词 */
func (r *SimbaToolsItemsTopGetRequest) SetKeyword(value string) {
	r.SetValue("keyword", value)
}

/* 主人昵称 */
func (r *SimbaToolsItemsTopGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}


func (r *SimbaToolsItemsTopGetRequest) GetResponse(accessToken string) (*SimbaToolsItemsTopGetResponse, []byte, error) {
	var resp SimbaToolsItemsTopGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.simba.tools.items.top.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type SimbaToolsItemsTopGetResponse struct {
	Rankeditems []*RankedItem `json:"rankeditems"`
}

type SimbaToolsItemsTopGetResponseResult struct {
	Response *SimbaToolsItemsTopGetResponse `json:"simba_tools_items_top_get_response"`
}





/* 直通车推广计划下的词报表基础数据查询<br/> 
异步API使用方法，请查看：<a href="http://open.taobao.com/doc/detail.htm?id=30">异步API使用说明</a><br/> */
type TopatsSimbaCampkeywordbaseGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询推广计划ID */
func (r *TopatsSimbaCampkeywordbaseGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 查询的昵称 */
func (r *TopatsSimbaCampkeywordbaseGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 报表类型。可选值：SEARCH（搜索）、CAT（类目出价）、 NOSEARCH（定向投放），可多选，如：SEARCH,CAT */
func (r *TopatsSimbaCampkeywordbaseGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源。可选值：1（站内）、2（站外）、SUMMARY（汇总），其中SUMMARY必须单选，其它值可多选，如：1,2 */
func (r *TopatsSimbaCampkeywordbaseGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 时间参数（昨天：DAY、 前一周：7DAY 、前15天：15DAY 、前30天：30DAY 、前90天：90DAY）单选 */
func (r *TopatsSimbaCampkeywordbaseGetRequest) SetTimeSlot(value string) {
	r.SetValue("time_slot", value)
}


func (r *TopatsSimbaCampkeywordbaseGetRequest) GetResponse(accessToken string) (*TopatsSimbaCampkeywordbaseGetResponse, []byte, error) {
	var resp TopatsSimbaCampkeywordbaseGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.simba.campkeywordbase.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsSimbaCampkeywordbaseGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsSimbaCampkeywordbaseGetResponseResult struct {
	Response *TopatsSimbaCampkeywordbaseGetResponse `json:"topats_simba_campkeywordbase_get_response"`
}





/* 推广计划下的词报表效果数据查询<br/> 
异步API使用方法，请查看：<a href="http://open.taobao.com/doc/detail.htm?id=30">异步API使用说明</a><br/> */
type TopatsSimbaCampkeywordeffectGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 查询推广计划ID */
func (r *TopatsSimbaCampkeywordeffectGetRequest) SetCampaignId(value string) {
	r.SetValue("campaign_id", value)
}

/* 查询的昵称 */
func (r *TopatsSimbaCampkeywordeffectGetRequest) SetNick(value string) {
	r.SetValue("nick", value)
}

/* 报表类型。可选值：SEARCH（搜索）、CAT（类目出价）、 NOSEARCH（定向投放），可多选，如：SEARCH,CAT */
func (r *TopatsSimbaCampkeywordeffectGetRequest) SetSearchType(value string) {
	r.SetValue("search_type", value)
}

/* 数据来源。可选值：1（站内）、2（站外）、SUMMARY（汇总），其中SUMMARY必须单选，其它值可多选，如：1,2 */
func (r *TopatsSimbaCampkeywordeffectGetRequest) SetSource(value string) {
	r.SetValue("source", value)
}

/* 时间参数（昨天：DAY、 前一周：7DAY、 前15天：15DAY、 前30天：30DAY 、前90天：90DAY）单选 */
func (r *TopatsSimbaCampkeywordeffectGetRequest) SetTimeSlot(value string) {
	r.SetValue("time_slot", value)
}


func (r *TopatsSimbaCampkeywordeffectGetRequest) GetResponse(accessToken string) (*TopatsSimbaCampkeywordeffectGetResponse, []byte, error) {
	var resp TopatsSimbaCampkeywordeffectGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.topats.simba.campkeywordeffect.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type TopatsSimbaCampkeywordeffectGetResponse struct {
	Task *Task `json:"task"`
}

type TopatsSimbaCampkeywordeffectGetResponseResult struct {
	Response *TopatsSimbaCampkeywordeffectGetResponse `json:"topats_simba_campkeywordeffect_get_response"`
}



