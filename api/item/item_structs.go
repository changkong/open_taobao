// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package item

const VersionNo = "20130729"

/* 授权 */
type SellerAuthorize struct {
	Brands         []*Brand   `json:"brands"`
	ItemCats       []*ItemCat `json:"item_cats"`
	XinpinItemCats []*ItemCat `json:"xinpin_item_cats"`
}

/* 品牌 */
type Brand struct {
	Name     string `json:"name"`
	Pid      int    `json:"pid"`
	PropName string `json:"prop_name"`
	Vid      int    `json:"vid"`
}

/* 商品类目结构 */
type ItemCat struct {
	Cid          int        `json:"cid"`
	Features     []*Feature `json:"features"`
	IsParent     bool       `json:"is_parent"`
	ModifiedTime string     `json:"modified_time"`
	ModifiedType string     `json:"modified_type"`
	Name         string     `json:"name"`
	ParentCid    int        `json:"parent_cid"`
	SortOrder    int        `json:"sort_order"`
	Status       string     `json:"status"`
}

/* 类目属性 */
type Feature struct {
	AttrKey   string `json:"attr_key"`
	AttrValue string `json:"attr_value"`
}

/* 商品属性 */
type ItemProp struct {
	ChildTemplate string       `json:"child_template"`
	Cid           int          `json:"cid"`
	Features      []*Feature   `json:"features"`
	IsAllowAlias  bool         `json:"is_allow_alias"`
	IsColorProp   bool         `json:"is_color_prop"`
	IsEnumProp    bool         `json:"is_enum_prop"`
	IsInputProp   bool         `json:"is_input_prop"`
	IsItemProp    bool         `json:"is_item_prop"`
	IsKeyProp     bool         `json:"is_key_prop"`
	IsSaleProp    bool         `json:"is_sale_prop"`
	ModifiedTime  string       `json:"modified_time"`
	ModifiedType  string       `json:"modified_type"`
	Multi         bool         `json:"multi"`
	Must          bool         `json:"must"`
	Name          string       `json:"name"`
	ParentPid     int          `json:"parent_pid"`
	ParentVid     int          `json:"parent_vid"`
	Pid           int          `json:"pid"`
	PropValues    []*PropValue `json:"prop_values"`
	Required      bool         `json:"required"`
	SortOrder     int          `json:"sort_order"`
	Status        string       `json:"status"`
	Type          string       `json:"type"`
}

/* 属性值 */
type PropValue struct {
	Cid          int        `json:"cid"`
	Features     []*Feature `json:"features"`
	IsParent     bool       `json:"is_parent"`
	ModifiedTime string     `json:"modified_time"`
	ModifiedType string     `json:"modified_type"`
	Name         string     `json:"name"`
	NameAlias    string     `json:"name_alias"`
	Pid          int        `json:"pid"`
	PropName     string     `json:"prop_name"`
	SortOrder    int        `json:"sort_order"`
	Status       string     `json:"status"`
	Vid          int        `json:"vid"`
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
