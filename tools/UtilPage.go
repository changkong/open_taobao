// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/yaofangou/open_taobao"
	"github.com/yaofangou/open_taobao/api/user"
	"html/template"
	"net/http"
	"strings"
)

var (
	CurTpl    string
	Result    string
	ResultMsg []string

	confMain    *ConfMainT
	confPackage *ConfPackageT
	data        *DataT
)

type UtilPage struct {
}

func (u *UtilPage) IsShowAuth() bool {
	return CurTpl == "auth.tpl"
}

func (u *UtilPage) IsShowTest() bool {
	return CurTpl == "test.tpl"
}

func (u *UtilPage) IsShowMake() bool {
	return CurTpl == "make.tpl"
}

func (u *UtilPage) GetConfMain() *ConfMainT {
	return confMain
}

func (u *UtilPage) GetConfPackage() *ConfPackageT {
	return confPackage
}

func (u *UtilPage) GetData() *DataT {
	return data
}

func (u *UtilPage) GetVersionNo() string {
	return VersionNo
}

func (u *UtilPage) GetPkgTaobaoVersionNo() string {
	return open_taobao.VersionNo
}

func (u *UtilPage) GetPkgUserVersionNo() string {
	return user.VersionNo
}

func (u *UtilPage) Init() error {
	var err error
	confMain, err = NewConfMain("./conf/main.json")
	if err != nil {
		return err
	}

	confPackage, err = NewConfPackage("./conf/package.json")
	if err != nil {
		return err
	}

	data, err = NewMetadata("./conf/ApiMetadata.xml", confPackage)
	if err != nil {
		return err
	}

	u.SetInfo("")

	return nil
}

func (u *UtilPage) MsgNeedBr(index int) bool {
	return len(ResultMsg) > 1 &&
		(index+1) < len(ResultMsg)
}

func (u *UtilPage) SetErr(msg string) {
	Result = "error"
	if msg == "" {
		ResultMsg = []string{}
	} else {
		ResultMsg = []string{msg}
	}
	fmt.Println("[error]", msg)
}

func (u *UtilPage) AddErr(msg string) {
	Result = "error"
	ResultMsg = append(ResultMsg, msg)
}

func (u *UtilPage) SetInfo(msg string) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		Result = ""
		ResultMsg = []string{}
	} else {
		Result = "info"
		ResultMsg = []string{msg}
	}
	fmt.Println("[info]", msg)
}

func (u *UtilPage) IsMsgNo() bool {
	return Result == ""
}

func (u *UtilPage) IsMsgInfo() bool {
	return Result == "info"
}

func (u *UtilPage) IsMsgErr() bool {
	return Result == "error"
}

func (u *UtilPage) GetResultMsg() []string {
	return ResultMsg
}

func (u *UtilPage) Redirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

func (u *UtilPage) Static(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	http.ServeFile(w, r, "."+r.URL.String())
}

func (u *UtilPage) ShowPage(w http.ResponseWriter, d interface{}, tpl string) {
	t1, err := template.ParseFiles("./pages/layout/main.tpl", "./pages/"+tpl)
	if err != nil {
		fmt.Println("[error]", err)
		fmt.Fprintf(w, err.Error())
		return
	}
	CurTpl = tpl // 一定放在前面执行，模板中会引用

	err = t1.ExecuteTemplate(w, "main", d)
	if err != nil {
		fmt.Println("[error]", err)
		fmt.Fprintf(w, err.Error())
		return
	}
	u.SetInfo("")
}

func (u *UtilPage) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	if r.RequestURI == "/" {
		u.Redirect(w, "/auth/")
	} else if r.RequestURI == "/favicon.ico" {

	} else {
		u.SetErr("不能识别请求:" + r.RequestURI)
		u.ShowPage(w, u, CurTpl)
	}
}
