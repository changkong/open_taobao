// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dmt

const VersionNo = "20131207"


/* 图片分类 */
type PictureCategory struct {
	Created string `json:"created"`
	Modified string `json:"modified"`
	ParentId int `json:"parent_id"`
	PictureCategoryId int `json:"picture_category_id"`
	PictureCategoryName string `json:"picture_category_name"`
	Position int `json:"position"`
	Type string `json:"type"`

}

/*  */
type PictureCategoryListObject struct {
	PictureCategory []*PictureCategory `json:"picture_category"`

}

/*  */
type PictureListObject struct {
	Picture []*Picture `json:"picture"`

}

/* 图片 */
type Picture struct {
	ClientType string `json:"client_type"`
	Created string `json:"created"`
	Deleted string `json:"deleted"`
	Md5 string `json:"md5"`
	Modified string `json:"modified"`
	PictureCategoryId int `json:"picture_category_id"`
	PictureId int `json:"picture_id"`
	PicturePath string `json:"picture_path"`
	Pixel string `json:"pixel"`
	Referenced bool `json:"referenced"`
	Sizes int `json:"sizes"`
	Status string `json:"status"`
	Title string `json:"title"`
	Uid int `json:"uid"`

}

/* 图片空间的用户信息获取，包括订购容量等 */
type UserInfo struct {
	AvailableSpace string `json:"available_space"`
	FreeSpace string `json:"free_space"`
	OrderExpiryDate string `json:"order_expiry_date"`
	OrderSpace string `json:"order_space"`
	RemainingSpace string `json:"remaining_space"`
	UsedSpace string `json:"used_space"`
	WaterMark string `json:"water_mark"`

}

/* 搜索返回的结果类 */
type TOPSearchResult struct {
	Paginator *TOPPaginator `json:"paginator"`
	VideoItems *VideoItemListObject `json:"video_items"`

}

/* 分页信息 */
type TOPPaginator struct {
	CurrentPage int `json:"current_page"`
	IsLastPage bool `json:"is_last_page"`
	PageSize int `json:"page_size"`
	TotalResults int `json:"total_results"`

}

/*  */
type VideoItemListObject struct {
	VideoItem []*VideoItem `json:"video_item"`

}

/* 视频 */
type VideoItem struct {
	CoverUrl string `json:"cover_url"`
	Description string `json:"description"`
	Duration int `json:"duration"`
	IsOpenToOther bool `json:"is_open_to_other"`
	State int `json:"state"`
	Tags []string `json:"tags"`
	Title string `json:"title"`
	UploadTime string `json:"upload_time"`
	UploaderId int `json:"uploader_id"`
	VideoId int `json:"video_id"`
	VideoPlayInfo *VideoPlayInfo `json:"video_play_info"`

}

/* 视频播放信息 */
type VideoPlayInfo struct {
	AndroidpadUrl string `json:"androidpad_url"`
	AndroidpadV23Url *AndroidVlowUrl `json:"androidpad_v23_url"`
	AndroidphoneUrl string `json:"androidphone_url"`
	AndroidphoneV23Url *AndroidVlowUrl `json:"androidphone_v23_url"`
	FlashUrl string `json:"flash_url"`
	IpadUrl string `json:"ipad_url"`
	IphoneUrl string `json:"iphone_url"`
	WebUrl string `json:"web_url"`

}

/* android  phone和pad播放的mp4文件类。适用2.3版本的Android。 */
type AndroidVlowUrl struct {
	Hd string `json:"hd"`
	Ld string `json:"ld"`
	Sd string `json:"sd"`
	Ud string `json:"ud"`

}

