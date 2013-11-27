// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package simba

const VersionNo = "20131127"


/* 推广组 */
type ADGroup struct {
	AdgroupId int `json:"adgroup_id"`
	CampaignId int `json:"campaign_id"`
	CategoryIds string `json:"category_ids"`
	CreateTime string `json:"create_time"`
	DefaultPrice int `json:"default_price"`
	IsNonsearchDefaultPrice bool `json:"is_nonsearch_default_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	NonsearchMaxPrice int `json:"nonsearch_max_price"`
	NonsearchStatus int `json:"nonsearch_status"`
	NumIid int `json:"num_iid"`
	OfflineType string `json:"offline_type"`
	OnlineStatus string `json:"online_status"`
	Reason string `json:"reason"`

}

/* 推广组类目出价 */
type ADGroupCatmatch struct {
	AdgroupId int `json:"adgroup_id"`
	CampaignId int `json:"campaign_id"`
	CatmatchId int `json:"catmatch_id"`
	CreateTime string `json:"create_time"`
	IsDefaultPrice bool `json:"is_default_price"`
	MaxPrice int `json:"max_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	OnlineStatus string `json:"online_status"`
	Qscore string `json:"qscore"`

}

/* 一页ADGroupCatMatch列表 */
type ADGroupCatMatchPage struct {
	AdgroupCatmatchList []*ADGroupCatmatch `json:"adgroup_catmatch_list"`
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	TotalItem int `json:"total_item"`

}

/* 类目出价预估信息 */
type ADGroupCatMatchForecast struct {
	AdgroupId int `json:"adgroup_id"`
	CatmatchId int `json:"catmatch_id"`
	Nick string `json:"nick"`
	PriceClick string `json:"price_click"`
	PriceCust string `json:"price_cust"`
	PriceRank string `json:"price_rank"`

}

/* 直通车商品分页对象 */
type SubwayItemPartition struct {
	ItemList []*SubwayItem `json:"item_list"`
	OrderBy bool `json:"order_by"`
	OrderField string `json:"order_field"`
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	TotalItem int `json:"total_item"`

}

/* 主站商品对象 */
type SubwayItem struct {
	ExtraAttributes *ExtraAttributes `json:"extra_attributes"`
	NumId int `json:"num_id"`
	Price float64 `json:"price"`
	Title string `json:"title"`

}

/* 直通车商品对象属性（Map） */
type ExtraAttributes struct {
	PublishTime string `json:"publish_time"`
	Quantity float64 `json:"quantity"`
	SalesCount float64 `json:"sales_count"`

}

/* 一页ADGroup列表 */
type ADGroupPage struct {
	AdgroupList []*ADGroup `json:"adgroup_list"`
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	TotalItem int `json:"total_item"`

}

/* 推广计划 */
type Campaign struct {
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	OnlineStatus string `json:"online_status"`
	SettleReason string `json:"settle_reason"`
	SettleStatus string `json:"settle_status"`
	Title string `json:"title"`

}

/* 推广计划的投放地域 */
type CampaignArea struct {
	Area string `json:"area"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`

}

/* 直通车可推广的地域 */
type AreaOption struct {
	AreaId int `json:"area_id"`
	Level int `json:"level"`
	Name string `json:"name"`
	ParentId int `json:"parent_id"`

}

/* 推广计划的日限额 */
type CampaignBudget struct {
	Budget int `json:"budget"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	IsSmooth bool `json:"is_smooth"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`

}

/* Campaign投放频道 */
type ChannelOption struct {
	ChannelId int `json:"channel_id"`
	IsNonsearch bool `json:"is_nonsearch"`
	IsSearch bool `json:"is_search"`
	Name string `json:"name"`
	TrafficName string `json:"traffic_name"`
	TrafficType string `json:"traffic_type"`

}

/* 推广计划的投放平台 */
type CampaignPlatform struct {
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	NonsearchChannels []int `json:"nonsearch_channels"`
	OutsideDiscount int `json:"outside_discount"`
	SearchChannels []int `json:"search_channels"`

}

/* 推广计划的分时折扣设置 */
type CampaignSchedule struct {
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	Schedule string `json:"schedule"`

}

/* 创意 */
type Creative struct {
	AdgroupId int `json:"adgroup_id"`
	AuditDesc string `json:"audit_desc"`
	AuditStatus string `json:"audit_status"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	CreativeId int `json:"creative_id"`
	ImgUrl string `json:"img_url"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	Title string `json:"title"`

}

/* 创意修改记录，只记录最后一次修改 */
type CreativeRecord struct {
	AuditDesc string `json:"audit_desc"`
	AuditStatus string `json:"audit_status"`
	CreateTime string `json:"create_time"`
	CreativeId int `json:"creative_id"`
	ImgUrl string `json:"img_url"`
	ModifiedTime string `json:"modified_time"`
	ModifyTime string `json:"modify_time"`
	Nick string `json:"nick"`
	OldImgUrl string `json:"old_img_url"`
	OldTitle string `json:"old_title"`
	Title string `json:"title"`

}

/* 广告创意分页对象 */
type CreativePage struct {
	CreativeList []*Creative `json:"creative_list"`
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	TotalItem int `json:"total_item"`

}

/* 类目对象 */
type INCategoryTop struct {
	CategoryChildTopList []*INCategoryChildTop `json:"category_child_top_list"`
	CategoryDesc string `json:"category_desc"`
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryPropertiesList []*INCategoryProperties `json:"category_properties_list"`
	CategroyWord string `json:"categroy_word"`

}

/* 类目对象 */
type INCategoryChildTop struct {
	CategoryDesc string `json:"category_desc"`
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryPropertiesList []*INCategoryProperties `json:"category_properties_list"`

}

/* 类目属性对象 */
type INCategoryProperties struct {
	PropertiesDesc string `json:"properties_desc"`
	PropertiesId int `json:"properties_id"`
	PropertiesName string `json:"properties_name"`

}

/* 类目数据分析对象 */
type INCategoryAnalysis struct {
	CategoryAreaPer string `json:"category_area_per"`
	CategoryHpPrice string `json:"category_hp_price"`
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategorySourcePer string `json:"category_source_per"`

}

/* 类目基础数据对象 */
type INCategoryBase struct {
	CategoryId int `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryPv int `json:"category_pv"`
	InRecordBaseList []*INRecordBase `json:"in_record_base_list"`

}

/* 词基础数据对象 */
type INRecordBase struct {
	AvgPrice int `json:"avg_price"`
	Click int `json:"click"`
	Competition int `json:"competition"`
	Ctr float64 `json:"ctr"`
	Date string `json:"date"`
	Pv int `json:"pv"`

}

/* 词数据分析对象 */
type INWordAnalysis struct {
	Word string `json:"word"`
	WordAreaPer string `json:"word_area_per"`
	WordHpPrice string `json:"word_hp_price"`
	WordSourcePer string `json:"word_source_per"`

}

/* 词基础数据对象 */
type INWordBase struct {
	InRecordBaseList []*INRecordBase `json:"in_record_base_list"`
	Word string `json:"word"`

}

/* 词和类目数据对象 */
type INWordCategory struct {
	AvgPrice int `json:"avg_price"`
	CategoryId int `json:"category_id"`
	Click int `json:"click"`
	Competition int `json:"competition"`
	Ctr float64 `json:"ctr"`
	Date string `json:"date"`
	Pv int `json:"pv"`
	Word string `json:"word"`

}

/* 词预估信息 */
type KeywordForecast struct {
	KeywordId int `json:"keyword_id"`
	Nick string `json:"nick"`
	PriceClick string `json:"price_click"`
	PriceCust string `json:"price_cust"`
	PriceRank string `json:"price_rank"`
	Word string `json:"word"`

}

/* 关键词分页对象 */
type KeywordPage struct {
	KeywordList []*Keyword `json:"keyword_list"`
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	TotalItem int `json:"total_item"`

}

/* 关键词 */
type Keyword struct {
	AdgroupId int `json:"adgroup_id"`
	AuditDesc string `json:"audit_desc"`
	AuditStatus string `json:"audit_status"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	IsDefaultPrice bool `json:"is_default_price"`
	IsGarbage bool `json:"is_garbage"`
	KeywordId int `json:"keyword_id"`
	MatchScope string `json:"match_scope"`
	MaxPrice int `json:"max_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	Qscore string `json:"qscore"`
	Word string `json:"word"`

}

/* 关键词质量得分 */
type KeywordQscore struct {
	AdgroupId int `json:"adgroup_id"`
	CampaignId int `json:"campaign_id"`
	KeywordId int `json:"keyword_id"`
	Nick string `json:"nick"`
	Qscore string `json:"qscore"`
	Word string `json:"word"`

}

/* 一页推荐词列表 */
type RecommendWordPage struct {
	PageNo int `json:"page_no"`
	PageSize int `json:"page_size"`
	RecommendWordList []*RecommendWord `json:"recommend_word_list"`
	TotalItem int `json:"total_item"`

}

/* 推荐词 */
type RecommendWord struct {
	AveragePrice string `json:"average_price"`
	Pertinence string `json:"pertinence"`
	Pv string `json:"pv"`
	Word string `json:"word"`

}

/* 质量得分类 */
type Qscore struct {
	CatmatchQscore string `json:"catmatch_qscore"`
	KeywordQscoreList []*KeywordQscore `json:"keyword_qscore_list"`

}

/* 推广组与定向推广位置关系 */
type ADGroupPlace struct {
	AdgroupId int `json:"adgroup_id"`
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	IsDefaultPrice int `json:"is_default_price"`
	MaxPrice int `json:"max_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`
	PlaceId int `json:"place_id"`
	PlaceName string `json:"place_name"`

}

/* 投放人群 */
type Demographic struct {
	CreateTime string `json:"create_time"`
	GroupId int `json:"group_id"`
	GroupName string `json:"group_name"`
	Id int `json:"id"`
	LastUpdateTime string `json:"last_update_time"`
	Name string `json:"name"`

}

/* 定向推广位置 */
type Place struct {
	CreateTime string `json:"create_time"`
	LastUpdateTime string `json:"last_update_time"`
	Name string `json:"name"`
	PlaceId int `json:"place_id"`

}

/* 投放人群设置 */
type DemographicSetting struct {
	CampaignId int `json:"campaign_id"`
	CreateTime string `json:"create_time"`
	DemographicId int `json:"demographic_id"`
	IncrementalPrice int `json:"incremental_price"`
	ModifiedTime string `json:"modified_time"`
	Nick string `json:"nick"`

}

/* 关键词排名推广商品信息 */
type RankedItem struct {
	LinkUrl string `json:"link_url"`
	MaxPrice float64 `json:"max_price"`
	Nick string `json:"nick"`
	Order int `json:"order"`
	RankScore int `json:"rank_score"`
	Title string `json:"title"`

}

/* 批量异步任务结果 */
type Task struct {
	CheckCode string `json:"check_code"`
	Created string `json:"created"`
	DownloadUrl string `json:"download_url"`
	Method string `json:"method"`
	Schedule string `json:"schedule"`
	Status string `json:"status"`
	Subtasks []*Subtask `json:"subtasks"`
	TaskId int `json:"task_id"`

}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess bool `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult string `json:"sub_task_result"`

}

