// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package item

const VersionNo = "20131207"


/* 授权 */
type SellerAuthorize struct {
	Brands *BrandListObject `json:"brands"`
	ItemCats *ItemCatListObject `json:"item_cats"`
	XinpinItemCats *ItemCatListObject `json:"xinpin_item_cats"`

}

/*  */
type BrandListObject struct {
	Brand []*Brand `json:"brand"`

}

/* 品牌 */
type Brand struct {
	Name string `json:"name"`
	Pid int `json:"pid"`
	PropName string `json:"prop_name"`
	Vid int `json:"vid"`

}

/*  */
type ItemCatListObject struct {
	ItemCat []*ItemCat `json:"item_cat"`

}

/* 商品类目结构 */
type ItemCat struct {
	Cid int `json:"cid"`
	Features *FeatureListObject `json:"features"`
	IsParent bool `json:"is_parent"`
	ModifiedTime string `json:"modified_time"`
	ModifiedType string `json:"modified_type"`
	Name string `json:"name"`
	ParentCid int `json:"parent_cid"`
	SortOrder int `json:"sort_order"`
	Status string `json:"status"`

}

/*  */
type FeatureListObject struct {
	Feature []*Feature `json:"feature"`

}

/* 类目属性 */
type Feature struct {
	AttrKey string `json:"attr_key"`
	AttrValue string `json:"attr_value"`

}

/*  */
type ItemPropListObject struct {
	ItemProp []*ItemProp `json:"item_prop"`

}

/* 商品属性 */
type ItemProp struct {
	ChildTemplate string `json:"child_template"`
	Cid int `json:"cid"`
	Features *FeatureListObject `json:"features"`
	IsAllowAlias bool `json:"is_allow_alias"`
	IsColorProp bool `json:"is_color_prop"`
	IsEnumProp bool `json:"is_enum_prop"`
	IsInputProp bool `json:"is_input_prop"`
	IsItemProp bool `json:"is_item_prop"`
	IsKeyProp bool `json:"is_key_prop"`
	IsSaleProp bool `json:"is_sale_prop"`
	ModifiedTime string `json:"modified_time"`
	ModifiedType string `json:"modified_type"`
	Multi bool `json:"multi"`
	Must bool `json:"must"`
	Name string `json:"name"`
	ParentPid int `json:"parent_pid"`
	ParentVid int `json:"parent_vid"`
	Pid int `json:"pid"`
	PropValues *PropValueListObject `json:"prop_values"`
	Required bool `json:"required"`
	SortOrder int `json:"sort_order"`
	Status string `json:"status"`
	Type string `json:"type"`

}

/*  */
type PropValueListObject struct {
	PropValue []*PropValue `json:"prop_value"`

}

/* 属性值 */
type PropValue struct {
	Cid int `json:"cid"`
	Features *FeatureListObject `json:"features"`
	IsParent bool `json:"is_parent"`
	ModifiedTime string `json:"modified_time"`
	ModifiedType string `json:"modified_type"`
	Name string `json:"name"`
	NameAlias string `json:"name_alias"`
	Pid int `json:"pid"`
	PropName string `json:"prop_name"`
	SortOrder int `json:"sort_order"`
	Status string `json:"status"`
	Vid int `json:"vid"`

}

