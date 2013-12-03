// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dmt

import (
	"github.com/yaofangou/open_taobao"
)





/* 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符 */
type PictureCategoryAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片分类的父分类,一级分类的parent_id为0,二级分类的则为其父分类的picture_category_id */
func (r *PictureCategoryAddRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}

/* 图片分类名称，最大长度20字符，中英文都算一字符，不能为空 */
func (r *PictureCategoryAddRequest) SetPictureCategoryName(value string) {
	r.SetValue("picture_category_name", value)
}


func (r *PictureCategoryAddRequest) GetResponse(accessToken string) (*PictureCategoryAddResponse, []byte, error) {
	var resp PictureCategoryAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.category.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureCategoryAddResponse struct {
	PictureCategory *PictureCategory `json:"picture_category"`
}

type PictureCategoryAddResponseResult struct {
	Response *PictureCategoryAddResponse `json:"picture_category_add_response"`
}





/* 获取图片分类信息 */
type PictureCategoryGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片分类的修改时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片分类。 */
func (r *PictureCategoryGetRequest) SetModifiedTime(value string) {
	r.SetValue("modified_time", value)
}

/* 取二级分类时设置为对应父分类id 
取一级分类时父分类id设为0 
取全部分类的时候不设或设为-1 */
func (r *PictureCategoryGetRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}

/* 图片分类ID */
func (r *PictureCategoryGetRequest) SetPictureCategoryId(value string) {
	r.SetValue("picture_category_id", value)
}

/* 图片分类名，不支持模糊查询 */
func (r *PictureCategoryGetRequest) SetPictureCategoryName(value string) {
	r.SetValue("picture_category_name", value)
}

/* 分类类型,fixed代表店铺装修分类类别，auction代表宝贝分类类别，user-define代表用户自定义分类类别 */
func (r *PictureCategoryGetRequest) SetType(value string) {
	r.SetValue("type", value)
}


func (r *PictureCategoryGetRequest) GetResponse(accessToken string) (*PictureCategoryGetResponse, []byte, error) {
	var resp PictureCategoryGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.category.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureCategoryGetResponse struct {
	PictureCategories []*PictureCategory `json:"picture_categories"`
}

type PictureCategoryGetResponseResult struct {
	Response *PictureCategoryGetResponse `json:"picture_category_get_response"`
}





/* 更新图片分类的名字，或者更新图片分类的父分类（即分类移动）。只能移动2级分类到非2级分类，默认分类和1级分类不可移动。 */
type PictureCategoryUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 要更新的图片分类的id */
func (r *PictureCategoryUpdateRequest) SetCategoryId(value string) {
	r.SetValue("category_id", value)
}

/* 图片分类的新名字，最大长度20字符，不能为空 */
func (r *PictureCategoryUpdateRequest) SetCategoryName(value string) {
	r.SetValue("category_name", value)
}

/* 图片分类的新父分类id */
func (r *PictureCategoryUpdateRequest) SetParentId(value string) {
	r.SetValue("parent_id", value)
}


func (r *PictureCategoryUpdateRequest) GetResponse(accessToken string) (*PictureCategoryUpdateResponse, []byte, error) {
	var resp PictureCategoryUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.category.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureCategoryUpdateResponse struct {
	Done bool `json:"done"`
}

type PictureCategoryUpdateResponseResult struct {
	Response *PictureCategoryUpdateResponse `json:"picture_category_update_response"`
}





/* 删除图片空间图片 */
type PictureDeleteRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片ID字符串,可以一个也可以一组,用英文逗号间隔,如450,120,155 */
func (r *PictureDeleteRequest) SetPictureIds(value string) {
	r.SetValue("picture_ids", value)
}


func (r *PictureDeleteRequest) GetResponse(accessToken string) (*PictureDeleteResponse, []byte, error) {
	var resp PictureDeleteResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.delete", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureDeleteResponse struct {
	Success bool `json:"success"`
}

type PictureDeleteResponseResult struct {
	Response *PictureDeleteResponse `json:"picture_delete_response"`
}





/* 获取图片信息 */
type PictureGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片使用，如果是pc宝贝detail使用，设置为client:computer，查询出来的图片是符合pc的宝贝detail显示的 
如果是手机宝贝detail使用，设置为client:phone，查询出来的图片是符合手机的宝贝detail显示的 */
func (r *PictureGetRequest) SetClientType(value string) {
	r.SetValue("client_type", value)
}

/* 是否删除,unfroze代表没有删除 */
func (r *PictureGetRequest) SetDeleted(value string) {
	r.SetValue("deleted", value)
}

/* 查询上传结束时间点,格式:yyyy-MM-dd HH:mm:ss */
func (r *PictureGetRequest) SetEndDate(value string) {
	r.SetValue("end_date", value)
}

/* 图片被修改的时间点，格式:yyyy-MM-dd HH:mm:ss。查询此修改时间点之后到目前的图片。 */
func (r *PictureGetRequest) SetModifiedTime(value string) {
	r.SetValue("modified_time", value)
}

/* 图片查询结果排序,time:desc按上传时间从晚到早(默认), time:asc按上传时间从早到晚,sizes:desc按图片从大到小，sizes:asc按图片从小到大,默认time:desc */
func (r *PictureGetRequest) SetOrderBy(value string) {
	r.SetValue("order_by", value)
}

/* 页码.传入值为1代表第一页,传入值为2代表第二页,依此类推,默认值为1 */
func (r *PictureGetRequest) SetPageNo(value string) {
	r.SetValue("page_no", value)
}

/* 每页条数.每页返回最多返回100条,默认值40 */
func (r *PictureGetRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 图片分类ID */
func (r *PictureGetRequest) SetPictureCategoryId(value string) {
	r.SetValue("picture_category_id", value)
}

/* 图片ID */
func (r *PictureGetRequest) SetPictureId(value string) {
	r.SetValue("picture_id", value)
}

/* 查询上传开始时间点,格式:yyyy-MM-dd HH:mm:ss */
func (r *PictureGetRequest) SetStartDate(value string) {
	r.SetValue("start_date", value)
}

/* 图片标题,最大长度50字符,中英文都算一字符 */
func (r *PictureGetRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 图片url查询接口 */
func (r *PictureGetRequest) SetUrls(value string) {
	r.SetValue("urls", value)
}


func (r *PictureGetRequest) GetResponse(accessToken string) (*PictureGetResponse, []byte, error) {
	var resp PictureGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureGetResponse struct {
	Pictures []*Picture `json:"pictures"`
	TotalResults int `json:"totalResults"`
}

type PictureGetResponseResult struct {
	Response *PictureGetResponse `json:"picture_get_response"`
}





/* 查询图片是否被引用，被引用返回true，未被引用返回false */
type PictureIsreferencedGetRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片id */
func (r *PictureIsreferencedGetRequest) SetPictureId(value string) {
	r.SetValue("picture_id", value)
}


func (r *PictureIsreferencedGetRequest) GetResponse(accessToken string) (*PictureIsreferencedGetResponse, []byte, error) {
	var resp PictureIsreferencedGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.isreferenced.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureIsreferencedGetResponse struct {
	IsReferenced bool `json:"is_referenced"`
}

type PictureIsreferencedGetResponseResult struct {
	Response *PictureIsreferencedGetResponse `json:"picture_isreferenced_get_response"`
}





/* 替换一张图片，只替换图片数据，图片名称，图片分类等保持不变。 */
type PictureReplaceRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片二进制文件流,不能为空,允许png、jpg、gif图片格式 */
func (r *PictureReplaceRequest) SetImageData(value string) {
	r.SetValue("image_data", value)
}

/* 要替换的图片的id，必须大于0 */
func (r *PictureReplaceRequest) SetPictureId(value string) {
	r.SetValue("picture_id", value)
}


func (r *PictureReplaceRequest) GetResponse(accessToken string) (*PictureReplaceResponse, []byte, error) {
	var resp PictureReplaceResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.replace", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureReplaceResponse struct {
	Done bool `json:"done"`
}

type PictureReplaceResponseResult struct {
	Response *PictureReplaceResponse `json:"picture_replace_response"`
}





/* 修改指定图片的图片名 */
type PictureUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 新的图片名，最大长度50字符，不能为空 */
func (r *PictureUpdateRequest) SetNewName(value string) {
	r.SetValue("new_name", value)
}

/* 要更改名字的图片的id */
func (r *PictureUpdateRequest) SetPictureId(value string) {
	r.SetValue("picture_id", value)
}


func (r *PictureUpdateRequest) GetResponse(accessToken string) (*PictureUpdateResponse, []byte, error) {
	var resp PictureUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureUpdateResponse struct {
	Done bool `json:"done"`
}

type PictureUpdateResponseResult struct {
	Response *PictureUpdateResponse `json:"picture_update_response"`
}





/* 上传单张图片 */
type PictureUploadRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 图片上传的来源，有电脑版本宝贝发布，手机版本宝贝发布 
client:computer电脑版本宝贝使用 
client:phone手机版本宝贝使用 */
func (r *PictureUploadRequest) SetClientType(value string) {
	r.SetValue("client_type", value)
}

/* 包括后缀名的图片标题,不能为空，如Bule.jpg,有些卖家希望图片上传后取图片文件的默认名 */
func (r *PictureUploadRequest) SetImageInputTitle(value string) {
	r.SetValue("image_input_title", value)
}

/* 图片二进制文件流,不能为空,允许png、jpg、gif图片格式,2M以内。 */
func (r *PictureUploadRequest) SetImg(value string) {
	r.SetValue("img", value)
}

/* 图片分类ID，设置具体某个分类ID或设置0上传到默认分类，只能传入一个分类 */
func (r *PictureUploadRequest) SetPictureCategoryId(value string) {
	r.SetValue("picture_category_id", value)
}

/* 图片标题,如果为空,传的图片标题就取去掉后缀名的image_input_title,超过50字符长度会截取50字符,重名会在标题末尾加"(1)";标题末尾已经有"(数字)"了，则数字加1 */
func (r *PictureUploadRequest) SetTitle(value string) {
	r.SetValue("title", value)
}


func (r *PictureUploadRequest) GetResponse(accessToken string) (*PictureUploadResponse, []byte, error) {
	var resp PictureUploadResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.upload", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureUploadResponse struct {
	Picture *Picture `json:"picture"`
}

type PictureUploadResponseResult struct {
	Response *PictureUploadResponse `json:"picture_upload_response"`
}





/* 查询用户的图片空间使用信息，包括：订购量，已使用容量，免费容量，总的可使用容量，订购有效期，剩余容量 */
type PictureUserinfoGetRequest struct {
	open_taobao.TaobaoMethodRequest
}



func (r *PictureUserinfoGetRequest) GetResponse(accessToken string) (*PictureUserinfoGetResponse, []byte, error) {
	var resp PictureUserinfoGetResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.picture.userinfo.get", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type PictureUserinfoGetResponse struct {
	UserInfo *UserInfo `json:"user_info"`
}

type PictureUserinfoGetResponseResult struct {
	Response *PictureUserinfoGetResponse `json:"picture_userinfo_get_response"`
}





/* 此api添加一个视频，视频发布后需要转码，审核等操作步骤，表示添加完成之后并不是立马能播放该视频，需要等后台转码成功，审核通过之后。线上视频转码审核过程比较快 */
type VideoAddRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 视频封面url,不能为空且不能超过512个英文字母 */
func (r *VideoAddRequest) SetCoverUrl(value string) {
	r.SetValue("cover_url", value)
}

/* 视频描述信息，不能为空且不能超过256个汉字 */
func (r *VideoAddRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 视频标签，以','隔开，不能为空且总长度不超过128个汉字 */
func (r *VideoAddRequest) SetTags(value string) {
	r.SetValue("tags", value)
}

/* 视频标题，不能为空且不超过128个汉字 */
func (r *VideoAddRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 上传唯一识别符,上传id */
func (r *VideoAddRequest) SetUploadId(value string) {
	r.SetValue("upload_id", value)
}

/* 视频上传者数字id */
func (r *VideoAddRequest) SetUploaderId(value string) {
	r.SetValue("uploader_id", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (r *VideoAddRequest) SetVideoAppKey(value string) {
	r.SetValue("video_app_key", value)
}


func (r *VideoAddRequest) GetResponse(accessToken string) (*VideoAddResponse, []byte, error) {
	var resp VideoAddResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.video.add", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VideoAddResponse struct {
	Model int `json:"model"`
}

type VideoAddResponseResult struct {
	Response *VideoAddResponse `json:"video_add_response"`
}





/* 此api用于更新一个视频的基本信息，包括视频标题，标签，描述，封面等，不包括是否允许他人观看，更新是否允许他人观看详见taobao.videos.open */
type VideoUpdateRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 视频封面url,不能超过512个英文字母 */
func (r *VideoUpdateRequest) SetCoverUrl(value string) {
	r.SetValue("cover_url", value)
}

/* 视频描述信息，不能超过256个汉字 */
func (r *VideoUpdateRequest) SetDescription(value string) {
	r.SetValue("description", value)
}

/* 视频标签，以','隔开，且总长度不超过128个汉字 */
func (r *VideoUpdateRequest) SetTags(value string) {
	r.SetValue("tags", value)
}

/* 视频标题，不超过128个汉字。title, tags,cover_url和description至少必须传一个 */
func (r *VideoUpdateRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (r *VideoUpdateRequest) SetVideoAppKey(value string) {
	r.SetValue("video_app_key", value)
}

/* 视频id */
func (r *VideoUpdateRequest) SetVideoId(value string) {
	r.SetValue("video_id", value)
}


func (r *VideoUpdateRequest) GetResponse(accessToken string) (*VideoUpdateResponse, []byte, error) {
	var resp VideoUpdateResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.video.update", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VideoUpdateResponse struct {
	Model bool `json:"model"`
}

type VideoUpdateResponseResult struct {
	Response *VideoUpdateResponse `json:"video_update_response"`
}





/* 此api用于搜索应用上传的所有视频，应用由appKey 唯一识别，搜索条件包含视频标题，标签，状态，关键字（标题or标签，不能同时设置），上传者数字id */
type VideosSearchRequest struct {
	open_taobao.TaobaoMethodRequest
}


/* 页码。默认返回的数据是从第一页开始 */
func (r *VideosSearchRequest) SetCurrentPage(value string) {
	r.SetValue("current_page", value)
}

/* 需要返回的视频对象字段。VideoItem结构体中所有字段均可返回；多个字段用“,”分隔；其中video_play_info中的播放url可选择性返回，其余属性全部返回；如果想返回整个子对象中所有url，那字段为video_play_info，如果是想返回子对象里面的字段，那字段为video_play_info.web_url。 */
func (r *VideosSearchRequest) SetFields(value string) {
	r.SetValue("fields", value)
}

/* 关键字(标题or标签，不能同时设置title,tag，否则冲突) */
func (r *VideosSearchRequest) SetKeywords(value string) {
	r.SetValue("keywords", value)
}

/* 每页条数，默认值是12 */
func (r *VideosSearchRequest) SetPageSize(value string) {
	r.SetValue("page_size", value)
}

/* 视频状态列表；视频状态：等待转码（1），转码中（2），转码失败（3），等待审核（4），未通过审核（5），通过审核（6） */
func (r *VideosSearchRequest) SetStates(value string) {
	r.SetValue("states", value)
}

/* 视频标签 */
func (r *VideosSearchRequest) SetTag(value string) {
	r.SetValue("tag", value)
}

/* 视频标题 */
func (r *VideosSearchRequest) SetTitle(value string) {
	r.SetValue("title", value)
}

/* 视频上传者数字id */
func (r *VideosSearchRequest) SetUploaderId(value string) {
	r.SetValue("uploader_id", value)
}

/* 在淘宝视频中的应用key，该值向淘宝视频申请产生 */
func (r *VideosSearchRequest) SetVideoAppKey(value string) {
	r.SetValue("video_app_key", value)
}


func (r *VideosSearchRequest) GetResponse(accessToken string) (*VideosSearchResponse, []byte, error) {
	var resp VideosSearchResponseResult
	data, err := r.TaobaoMethodRequest.GetResponse(accessToken, "taobao.videos.search", &resp)
	if err != nil {
		return nil, data, err
	}
	return resp.Response, data, err
}

type VideosSearchResponse struct {
	SearchResult *TOPSearchResult `json:"search_result"`
}

type VideosSearchResponseResult struct {
	Response *VideosSearchResponse `json:"videos_search_response"`
}



