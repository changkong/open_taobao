// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type PropsT struct {
	Prop []*PropT `xml:"prop"`
}

type PropT struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`  // Level 为 Basic 的类型有 Number String Date Boolean Price
	Level string `xml:"level"` // 四种 Basic/Object/Basic Array/Object Array
	Desc  string `xml:"desc"`
}

func (p *PropT) GoName() string {
	return GetGoName(p.Name)
}

func (p *PropT) GoType() string {
	if p.Level == "Object Array" {
		return GetGoType(p.Type, p.Level) + "ListObject"
	}
	return GetGoType(p.Type, p.Level)
}

type StructT struct {
	Name  string `xml:"name"`
	Desc  string `xml:"desc"`
	Props PropsT `xml:"props"`
}

type ApiT struct {
	Name     string    `xml:"name"`
	Desc     string    `xml:"desc"`
	Request  RequestT  `xml:"request"`
	Response ResponseT `xml:"response"`
}

func (a *ApiT) GoName() string {
	return GetApiKey(a.Name)
}

func (a *ApiT) JsonName() string {
	return GetApiJsonName(a.Name)
}

type RequestT struct {
	Param []ParamReqT `xml:"param"`
}

type ParamReqT struct {
	Name        string `xml:"name"`
	Type        string `xml:"type"`     // Field List/String/Number/Boolean/Price/byte[]/Date
	Required    string `xml:"required"` // required/optional/special 其中 special 多半参照 desc 的说明
	MaxValue    string `xml:"max_value"`
	MaxLength   string `xml:"max_length"`
	MaxListSize string `xml:"max_list_size"`
	MinValue    string `xml:"min_value"`
	Desc        string `xml:"desc"`
}

func (p *ParamReqT) GoName() string {
	return GetGoName(p.Name)
}

func (p *ParamReqT) GoType() string {
	return GetGoType(p.Type, "Basic")
}

type ParamRespT struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`  // 参见 PropT
	Level string `xml:"level"` // 参见 PropT
	Desc  string `xml:"desc"`
}

func (p *ParamRespT) GoName() string {
	return GetGoName(p.Name)
}

func (p *ParamRespT) GoType() string {
	if p.Level == "Object Array" {
		return GetGoType(p.Type, p.Level) + "ListObject"
	}
	return GetGoType(p.Type, p.Level)
}

type ResponseT struct {
	Param []*ParamRespT `xml:"param"`
}

type StructsT struct {
	Struct []*StructT `xml:"struct"`
}

type ApisT struct {
	Api []*ApiT `xml:"api"`
}

type MetadataT struct {
	VersionNo string   `xml:"versionNo,attr"`
	Structs   StructsT `xml:"structs"`
	Apis      ApisT    `xml:"apis"`
}

type DataT struct {
	MetaVersionNo string
	ConfFile      string
	Structs       []*StructT              // 所有 struct 的数组
	Apis          []*ApiT                 // 所有的 api 的数组
	MapStruct     map[string]*StructT     // 所有 struct 的 map, key为 struct name
	MapApi        map[string]*ApiT        // 所有 api 的 map, key为 api name
	MapPkgStruct  map[string]([]*StructT) // key 为包名 value 为包下对应的 struct
	MapPkgApi     map[string]([]*ApiT)    // key 为包名 value 为包下对应的 api
}

func NewMetadata(filename string, confPackage *ConfPackageT) (*DataT, error) {
	fmt.Println("read", filename)

	var metadata MetadataT

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &metadata)
	if err != nil {
		return nil, err
	}

	var d DataT
	d.MetaVersionNo = metadata.VersionNo
	d.ConfFile = filename
	d.Structs = metadata.Structs.Struct
	d.Apis = metadata.Apis.Api
	d.MapStruct = make(map[string]*StructT)
	d.MapApi = make(map[string]*ApiT)
	d.MapPkgStruct = make(map[string][]*StructT)
	d.MapPkgApi = make(map[string][]*ApiT)

	for _, v := range d.Structs {
		d.MapStruct[v.Name] = v
	}

	for _, v := range d.Apis {
		d.MapApi[v.Name] = v
	}

	for _, v := range confPackage.Mx {
		d.MapPkgStruct[v.Name] = make([]*StructT, 0)
		d.MapPkgApi[v.Name] = make([]*ApiT, 0)
	}

	// 按包将 api 分开
	for _, pkg := range confPackage.Mx {
		apis := d.MapPkgApi[pkg.Name]
		for _, api := range d.Apis {
			for _, prefix := range pkg.Prefix {
				if strings.HasPrefix(api.Name, prefix) {
					apis = append(apis, api)
				}
			}
		}
		d.MapPkgApi[pkg.Name] = apis // 注意设置回去
	}

	// 按包获取 api 需要的 struct
	for _, pkg := range confPackage.Mx {
		apis := d.MapPkgApi[pkg.Name]
		structs := make([]*StructT, 0)
		okStruct := make(map[string]bool)
		for _, api := range apis {
			for _, para := range api.Response.Param {
				if para.Level == "Object" || para.Level == "Object Array" {
					addApiNeedStructs(para.Type, para.Level, &d, &structs, okStruct)
				}
			}
		}
		d.MapPkgStruct[pkg.Name] = structs // 注意设置回去
	}

	return &d, nil
}

func addApiNeedStructs(curStructName string, level string, data *DataT, structArray *[]*StructT, okStruct map[string]bool) {

	key := curStructName + level
	_, ok := okStruct[key]
	if ok {
		return
	}

	var curStruct *StructT

	if level == "Object Array" {
		structName := curStructName + "ListObject"

		props := PropsT{Prop: []*PropT{&PropT{Name: GetTaobaoName(curStructName), Type: curStructName, Level: "List Object"}}}
		curStruct = &StructT{Name: structName, Desc: "", Props: props}
		data.MapStruct[curStruct.Name] = curStruct

		okStruct[key] = true
		*structArray = append(*structArray, curStruct)
	}

	if !okStruct[curStructName+"Object"] {
		okStruct[curStructName+"Object"] = true
		curStruct = data.MapStruct[curStructName]
		*structArray = append(*structArray, curStruct)
	}

	for _, v := range curStruct.Props.Prop {
		if v.Level == "Object" || v.Level == "Object Array" {
			addApiNeedStructs(v.Type, v.Level, data, structArray, okStruct)
		}
	}
}

func GetApiJsonName(apiName string) string {
	if apiName[0:7] == "taobao." {
		apiName = apiName[7:]
	}
	return strings.Replace(apiName, ".", "_", -1)
}

func GetApiKey(apiName string) string {
	result := ""
	if apiName[0:7] == "taobao." {
		apiName = apiName[7:]
	}

	needUpper := true
	for i, c := range apiName {
		if needUpper {
			result += strings.ToUpper(apiName[i : i+1])
			needUpper = false
		} else if c == '.' {
			needUpper = true
		} else {
			result += apiName[i : i+1]
		}
	}

	return result
}

func GetGoName(jsonName string) string {
	result := ""
	needUpper := true
	for i, c := range jsonName {
		if needUpper {
			result += strings.ToUpper(jsonName[i : i+1])
			needUpper = false
		} else if c == '_' || c == '.' {
			needUpper = true
		} else {
			result += jsonName[i : i+1]
		}
	}
	return result
}

func GetTaobaoName(goName string) string {
	result := ""
	for i, c := range goName {
		if unicode.IsUpper(c) {
			if i == 0 {
				result += strings.ToLower(goName[i : i+1])
			} else {
				result += "_" + strings.ToLower(goName[i:i+1])
			}
		} else {
			result += goName[i : i+1]
		}
	}
	return result
}

func GetGoType(TaobaoType, TaobaoLevel string) string {
	result := ""
	if TaobaoLevel == "Basic Array" || TaobaoLevel == "List Object" {
		result = "[]"
	}
	switch TaobaoLevel {
	case "Basic", "Basic Array":
		switch TaobaoType {
		case "Number":
			result += "int"
		case "String", "Date", "Field List": // 日期转换为字符串
			result += "string"
		case "Boolean":
			result += "bool"
		case "Price":
			result += "float64"
		}
	case "Object", "Object Array", "List Object":
		result += "*"
		result += TaobaoType
	}
	return result
}
