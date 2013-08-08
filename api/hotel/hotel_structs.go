// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package hotel

const VersionNo = "20130808"

/* Hotel（酒店）结构。各字段详细说明可参考接口定义，如：酒店发布接口。 */
type Hotel struct {
	Address         string      `json:"address"`
	AliasName       string      `json:"alias_name"`
	AuditDenyReason string      `json:"audit_deny_reason"`
	City            int         `json:"city"`
	CityStr         string      `json:"city_str"`
	Country         string      `json:"country"`
	CountryStr      string      `json:"country_str"`
	Created         string      `json:"created"`
	DecorateTime    string      `json:"decorate_time"`
	Desc            string      `json:"desc"`
	District        int         `json:"district"`
	DistrictStr     string      `json:"district_str"`
	Hid             int         `json:"hid"`
	Level           string      `json:"level"`
	Modified        string      `json:"modified"`
	Name            string      `json:"name"`
	OpeningTime     string      `json:"opening_time"`
	Orientation     string      `json:"orientation"`
	PicUrl          string      `json:"pic_url"`
	Province        int         `json:"province"`
	ProvinceStr     string      `json:"province_str"`
	RoomTypes       []*RoomType `json:"room_types"`
	Rooms           int         `json:"rooms"`
	Service         string      `json:"service"`
	Status          int         `json:"status"`
	Storeys         int         `json:"storeys"`
	Tel             string      `json:"tel"`
}

/* RoomType（房型）结构。各字段详细说明可参考接口定义，如：房型发布接口。 */
type RoomType struct {
	AliasName       string `json:"alias_name"`
	AuditDenyReason string `json:"audit_deny_reason"`
	Created         string `json:"created"`
	Hid             int    `json:"hid"`
	Modified        string `json:"modified"`
	Name            string `json:"name"`
	Rid             int    `json:"rid"`
	Status          int    `json:"status"`
}

/* 酒店图片 */
type HotelImage struct {
	Hid int    `json:"hid"`
	Pic string `json:"pic"`
}

/* HotelOrder（酒店订单）结构。各字段详细说明可参考接口定义。注意：trade_status，refund_status，logistics_status不是严格准确的，请以交易API，物流API等得到的订单状态、物流状态为准确依据。 */
type HotelOrder struct {
	AlipayTradeNo   string        `json:"alipay_trade_no"`
	ArriveEarly     string        `json:"arrive_early"`
	ArriveLate      string        `json:"arrive_late"`
	BuyerNick       string        `json:"buyer_nick"`
	CheckinDate     string        `json:"checkin_date"`
	CheckoutDate    string        `json:"checkout_date"`
	ContactName     string        `json:"contact_name"`
	ContactPhone    string        `json:"contact_phone"`
	ContactPhoneBak string        `json:"contact_phone_bak"`
	Created         string        `json:"created"`
	EndTime         string        `json:"end_time"`
	Gid             int           `json:"gid"`
	Guests          []*OrderGuest `json:"guests"`
	Hid             int           `json:"hid"`
	LogisticsStatus string        `json:"logistics_status"`
	Message         string        `json:"message"`
	Modified        string        `json:"modified"`
	Nights          int           `json:"nights"`
	Oid             int           `json:"oid"`
	OutOid          string        `json:"out_oid"`
	PayTime         string        `json:"pay_time"`
	Payment         int           `json:"payment"`
	RefundStatus    string        `json:"refund_status"`
	Rid             int           `json:"rid"`
	RoomNumber      int           `json:"room_number"`
	SellerNick      string        `json:"seller_nick"`
	Tid             int           `json:"tid"`
	TotalRoomPrice  int           `json:"total_room_price"`
	TradeStatus     string        `json:"trade_status"`
	Type            string        `json:"type"`
}

/* 酒店订单入住人结构。 */
type OrderGuest struct {
	Name      string `json:"name"`
	Oid       int    `json:"oid"`
	PersonPos int    `json:"person_pos"`
	RoomPos   int    `json:"room_pos"`
	Tel       string `json:"tel"`
}

/* Room（酒店商品）结构。各字段详细说明可参考接口定义，如：商品发布接口。 */
type Room struct {
	Area            string    `json:"area"`
	Bbn             string    `json:"bbn"`
	BedType         string    `json:"bed_type"`
	Breakfast       string    `json:"breakfast"`
	Created         string    `json:"created"`
	Deposit         int       `json:"deposit"`
	Desc            string    `json:"desc"`
	Fee             int       `json:"fee"`
	Gid             int       `json:"gid"`
	Guide           string    `json:"guide"`
	Hid             int       `json:"hid"`
	Hotel           *Hotel    `json:"hotel"`
	Iid             int       `json:"iid"`
	Modified        string    `json:"modified"`
	MultiRoomQuotas string    `json:"multi_room_quotas"`
	PaymentType     string    `json:"payment_type"`
	PicUrl          string    `json:"pic_url"`
	PriceType       string    `json:"price_type"`
	Rid             int       `json:"rid"`
	RoomQuotas      string    `json:"room_quotas"`
	RoomType        *RoomType `json:"room_type"`
	Service         string    `json:"service"`
	Size            string    `json:"size"`
	Status          int       `json:"status"`
	Storey          string    `json:"storey"`
	Title           string    `json:"title"`
}

/* RoomImage（酒店图片）结构。各字段详细说明可参考接口定义，如：商品图片上传接口。 */
type RoomImage struct {
	AllImages string `json:"all_images"`
	Gid       int    `json:"gid"`
	Image     string `json:"image"`
	Position  int    `json:"position"`
}
