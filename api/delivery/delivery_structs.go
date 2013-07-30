// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package delivery

const VersionNo = "20130729"

/* 撤销运单结果 */
type CancelOrderResult struct {
	RecreateOrderId int `json:"recreate_order_id"`
}

/* 物流费用信息 */
type OrderCharge struct {
	OriginalTotalCost float64             `json:"original_total_cost"`
	OtherCost         float64             `json:"other_cost"`
	TotalCost         float64             `json:"total_cost"`
	TotalSavedCost    float64             `json:"total_saved_cost"`
	TransportCharge   *TransportCharge    `json:"transport_charge"`
	VasCharge         *LogisticsVasCharge `json:"vas_charge"`
}

/* 运输费用 */
type TransportCharge struct {
	Cost         float64 `json:"cost"`
	CostBy       string  `json:"cost_by"`
	OriginalCost float64 `json:"original_cost"`
	SavedCost    float64 `json:"saved_cost"`
}

/* 增值服务的总的费用 */
type LogisticsVasCharge struct {
	OriginalTotalVasCost float64                   `json:"original_total_vas_cost"`
	TotalVasCost         float64                   `json:"total_vas_cost"`
	TotalVasSaveCost     float64                   `json:"total_vas_save_cost"`
	VasCostList          []*LogisticsVasItemCharge `json:"vas_cost_list"`
}

/* 物流增值服务每项费用 */
type LogisticsVasItemCharge struct {
	Cost         float64 `json:"cost"`
	OriginalCost float64 `json:"original_cost"`
	VasCode      string  `json:"vas_code"`
	VasId        string  `json:"vas_id"`
}

/* 发货返回结果 */
type ConsignResult struct {
	LogisticsId string `json:"logistics_id"`
	OrderId     int    `json:"order_id"`
}

/* 查询线路的结果 */
type QueryRouteResult struct {
	CompanyRouteSummarys []*CompanyRouteSummary `json:"company_route_summarys"`
	IsTurnLevel          bool                   `json:"is_turn_level"`
	PagesRouteDetails    *RouteAlpPage          `json:"pages_route_details"`
	RouteVas             []*RouteVasInfo        `json:"route_vas"`
	TotalCorps           int                    `json:"total_corps"`
	TotalRoutes          int                    `json:"total_routes"`
}

/* 公司的线路统计 */
type CompanyRouteSummary struct {
	CompanyCode string `json:"company_code"`
	CompanyId   string `json:"company_id"`
	CompanyName string `json:"company_name"`
	RouteCounts int    `json:"route_counts"`
}

/* 线路的分页信息 */
type RouteAlpPage struct {
	Datas       []*ComplexLogisticsRoute `json:"datas"`
	End         int                      `json:"end"`
	PageCount   int                      `json:"page_count"`
	PageIndex   int                      `json:"page_index"`
	PageSize    int                      `json:"page_size"`
	RecordCount int                      `json:"record_count"`
	Start       int                      `json:"start"`
}

/* 线路的整条完整信息 */
type ComplexLogisticsRoute struct {
	CarriageInfo     *RouteCarriageInfo   `json:"carriage_info"`
	Company          *FreightCompany      `json:"company"`
	ExtenalInfo      *RouteExtenalInfo    `json:"extenal_info"`
	FromAreaId       int                  `json:"from_area_id"`
	FromCityName     string               `json:"from_city_name"`
	FromCountyName   string               `json:"from_county_name"`
	FromProvinceName string               `json:"from_province_name"`
	PromotionInfo    *RoutePromotionInfo  `json:"promotion_info"`
	RouteCode        string               `json:"route_code"`
	StatisticsInfo   *RouteStatisticsInfo `json:"statistics_info"`
	ToAreaId         int                  `json:"to_area_id"`
	ToCityName       string               `json:"to_city_name"`
	ToCountyName     string               `json:"to_county_name"`
	ToProvinceName   string               `json:"to_province_name"`
}

/* 线路运输相关的基本信息 */
type RouteCarriageInfo struct {
	AddFee             float64 `json:"add_fee"`
	Comments           string  `json:"comments"`
	GiveTime           string  `json:"give_time"`
	InitialFee         float64 `json:"initial_fee"`
	LeastExpense       float64 `json:"least_expense"`
	OrigVolumeRate     float64 `json:"orig_volume_rate"`
	OrigWeightRate     float64 `json:"orig_weight_rate"`
	PriceDescription   string  `json:"price_description"`
	TakeTime           string  `json:"take_time"`
	TransportMode      string  `json:"transport_mode"`
	TransportName      string  `json:"transport_name"`
	TransportTime      string  `json:"transport_time"`
	TransportTimeHours int     `json:"transport_time_hours"`
	TransportTypeCode  string  `json:"transport_type_code"`
	TransportWay       string  `json:"transport_way"`
	VolumeRate         float64 `json:"volume_rate"`
	WeightRate         float64 `json:"weight_rate"`
}

/* 线路的相关公司信息 */
type FreightCompany struct {
	Comments           string          `json:"comments"`
	CompanyId          int             `json:"company_id"`
	CompanyName        string          `json:"company_name"`
	CompanyeCode       string          `json:"companye_code"`
	CorpLevel          string          `json:"corp_level"`
	CustomerServiceTel string          `json:"customer_service_tel"`
	LogoUrl            string          `json:"logo_url"`
	ShopUrl            string          `json:"shop_url"`
	Sort               int             `json:"sort"`
	VasFeeHelpUrl      string          `json:"vas_fee_help_url"`
	WangwangList       []*WangwangInfo `json:"wangwang_list"`
}

/* 旺旺信息 */
type WangwangInfo struct {
	Site       string `json:"site"`
	WangwangId string `json:"wangwang_id"`
}

/* 线路附加信息 */
type RouteExtenalInfo struct {
	Cod                bool     `json:"cod"`
	CreditOpened       bool     `json:"credit_opened"`
	CreditTotalBalance float64  `json:"credit_total_balance"`
	Recommend          bool     `json:"recommend"`
	SpecialCodes       []string `json:"special_codes"`
	Top                bool     `json:"top"`
}

/* 线路的促销信息 */
type RoutePromotionInfo struct {
	BasePromotionId      string `json:"base_promotion_id"`
	PromotionDescription string `json:"promotion_description"`
	PromotionUrl         string `json:"promotion_url"`
	UnbasePromotion      bool   `json:"unbase_promotion"`
}

/* 线路的统计信息 */
type RouteStatisticsInfo struct {
	EvaluationCount      int     `json:"evaluation_count"`
	EvaluationScore      float64 `json:"evaluation_score"`
	FromNetworkCount     int     `json:"from_network_count"`
	ToNetworkCount       int     `json:"to_network_count"`
	TrunkRouteOrderCount int     `json:"trunk_route_order_count"`
}

/* 线路的增值服务信息 */
type RouteVasInfo struct {
	RouteCode string          `json:"route_code"`
	VasList   []*LogisticsVas `json:"vas_list"`
}

/* 货运增值服务信息 */
type LogisticsVas struct {
	ChargeCalculateType string `json:"charge_calculate_type"`
	Comments            string `json:"comments"`
	DefaultSelected     bool   `json:"default_selected"`
	Displayable         bool   `json:"displayable"`
	Needed              bool   `json:"needed"`
	ValueDisplayable    bool   `json:"value_displayable"`
	VasCode             string `json:"vas_code"`
	VasName             string `json:"vas_name"`
}

/* 地址区域结构 */
type Area struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ParentId int    `json:"parent_id"`
	Type     int    `json:"type"`
	Zip      string `json:"zip"`
}

/* 运费模板对象 */
type DeliveryTemplate struct {
	Address       string    `json:"address"`
	Assumer       int       `json:"assumer"`
	ConsignAreaId int       `json:"consign_area_id"`
	Created       string    `json:"created"`
	FeeList       []*TopFee `json:"fee_list"`
	Modified      string    `json:"modified"`
	Name          string    `json:"name"`
	Supports      string    `json:"supports"`
	TemplateId    int       `json:"template_id"`
	Valuation     int       `json:"valuation"`
}

/* 运费模板中运费信息对象 */
type TopFee struct {
	AddFee        string `json:"add_fee"`
	AddStandard   string `json:"add_standard"`
	Destination   string `json:"destination"`
	ServiceType   string `json:"service_type"`
	StartFee      string `json:"start_fee"`
	StartStandard string `json:"start_standard"`
}

/* 地址库返回数据信息 */
type AddressResult struct {
	Addr          string `json:"addr"`
	AreaId        int    `json:"area_id"`
	CancelDef     bool   `json:"cancel_def"`
	City          string `json:"city"`
	ContactId     int    `json:"contact_id"`
	ContactName   string `json:"contact_name"`
	Country       string `json:"country"`
	GetDef        bool   `json:"get_def"`
	Memo          string `json:"memo"`
	MobilePhone   string `json:"mobile_phone"`
	ModifyDate    string `json:"modify_date"`
	Phone         string `json:"phone"`
	Province      string `json:"province"`
	SellerCompany string `json:"seller_company"`
	SendDef       bool   `json:"send_def"`
	ZipCode       string `json:"zip_code"`
}

/* 物流公司基础数据结构 */
type LogisticsCompany struct {
	Code      string `json:"code"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	RegMailNo string `json:"reg_mail_no"`
}

/* 获取的物流订单详情列表 返回的Shipping包含的具体信息为入参fields请求的字段信息 */
type Shipping struct {
	BuyerNick       string    `json:"buyer_nick"`
	CompanyName     string    `json:"company_name"`
	Created         string    `json:"created"`
	DeliveryEnd     string    `json:"delivery_end"`
	DeliveryStart   string    `json:"delivery_start"`
	FreightPayer    string    `json:"freight_payer"`
	IsQuickCodOrder bool      `json:"is_quick_cod_order"`
	IsSpilt         int       `json:"is_spilt"`
	IsSuccess       bool      `json:"is_success"`
	ItemTitle       string    `json:"item_title"`
	Location        *Location `json:"location"`
	Modified        string    `json:"modified"`
	OrderCode       string    `json:"order_code"`
	OutSid          string    `json:"out_sid"`
	ReceiverMobile  string    `json:"receiver_mobile"`
	ReceiverName    string    `json:"receiver_name"`
	ReceiverPhone   string    `json:"receiver_phone"`
	SellerConfirm   string    `json:"seller_confirm"`
	SellerNick      string    `json:"seller_nick"`
	Status          string    `json:"status"`
	SubTids         []int     `json:"sub_tids"`
	Tid             int       `json:"tid"`
	Type            string    `json:"type"`
}

/* 用户地址 */
type Location struct {
	Address  string `json:"address"`
	City     string `json:"city"`
	Country  string `json:"country"`
	District string `json:"district"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}

/* 查询揽送范围之内的物流公司信息 */
type LogisticsPartner struct {
	Carriage      *CarriageDetail `json:"carriage"`
	CoverRemark   string          `json:"cover_remark"`
	Partner       *PartnerDetail  `json:"partner"`
	UncoverRemark string          `json:"uncover_remark"`
}

/* 物流公司资费相关信息 */
type CarriageDetail struct {
	AddFee        int    `json:"add_fee"`
	AddWeight     int    `json:"add_weight"`
	DamagePayment string `json:"damage_payment"`
	GotTime       string `json:"got_time"`
	InitialFee    int    `json:"initial_fee"`
	InitialWeight int    `json:"initial_weight"`
	LostPayment   string `json:"lost_payment"`
	WayDay        string `json:"way_day"`
}

/* 物流公司详细信息 */
type PartnerDetail struct {
	AccountNo   string `json:"account_no"`
	CompanyCode string `json:"company_code"`
	CompanyId   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
	FullName    string `json:"full_name"`
	RegMailNo   string `json:"reg_mail_no"`
	WangwangId  string `json:"wangwang_id"`
}

/* 物流跟踪信息的一条 */
type TransitStepInfo struct {
	Action          string `json:"action"`
	Desc            string `json:"desc"`
	NodeDescription string `json:"node_description"`
	StatusDesc      string `json:"status_desc"`
	StatusTime      string `json:"status_time"`
	Time            string `json:"time"`
}

/* 批量异步任务结果 */
type Task struct {
	CheckCode   string     `json:"check_code"`
	Created     string     `json:"created"`
	DownloadUrl string     `json:"download_url"`
	Method      string     `json:"method"`
	Schedule    string     `json:"schedule"`
	Status      string     `json:"status"`
	Subtasks    []*Subtask `json:"subtasks"`
	TaskId      int        `json:"task_id"`
}

/* 批量异步任务的子任务结果 */
type Subtask struct {
	IsSuccess      bool   `json:"is_success"`
	SubTaskRequest string `json:"sub_task_request"`
	SubTaskResult  string `json:"sub_task_result"`
}
