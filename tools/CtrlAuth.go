// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/yaofangou/open_taobao"
	"net/http"
	"strings"
)

type CtrlAuth struct {
	UtilPage
}

func (c *CtrlAuth) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	if r.RequestURI == "/auth/" {
		c.ShowPage(w, c, "auth.tpl")
		return
	}
	needRedirect := true

	if strings.HasPrefix(r.RequestURI, "/auth/save") {
		c.save(w, r)
	} else if strings.HasPrefix(r.RequestURI, "/auth/do") {
		needRedirect = c.do(w, r)
	} else if strings.HasPrefix(r.RequestURI, "/auth/callback") {
		c.callback(w, r)
	} else {
		c.SetErr("不能识别的请求:" + r.RequestURI)
	}

	if needRedirect {
		c.Redirect(w, "/auth/")
	}
}

func (c *CtrlAuth) save(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	err = c.GetConfMain().SavePara(
		r.Form.Get("AppKey"),
		r.Form.Get("AppSecret"),
		r.Form.Get("RedirectUri"))
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	c.SetInfo("save ok!")
}

// 返回 true 代表执行完后，需要恢复转向到 /auth
func (c *CtrlAuth) do(w http.ResponseWriter, r *http.Request) bool {
	err := c.GetConfMain().CheckPara()
	if err != nil {
		c.SetErr(err.Error())
		return true
	}
	url_, err := open_taobao.GetUrlForAuth(
		c.GetConfMain().AppKey, c.GetConfMain().RedirectUri, "")
	if err != nil {
		c.SetErr(err.Error())
		return true
	}

	c.Redirect(w, url_.String())
	return false
}

func (c *CtrlAuth) callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("...1")
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	fmt.Println("...2")
	TBErr := r.Form.Get("error")
	TBErrDesc := r.Form.Get("error_description")
	if TBErr != "" {
		c.SetErr(TBErrDesc)
		return
	}

	fmt.Println("...3")
	code := r.Form.Get("code")
	fmt.Println("code:", code)
	var req open_taobao.TokenGetRequest
	req.SetAppKey(c.GetConfMain().AppKey)
	req.SetAppSecret(c.GetConfMain().AppSecret)
	req.SetRedirectUri(c.GetConfMain().RedirectUri)
	req.SetCode(code)
	req.SetState("")
	resp, _, err := req.GetResponse()
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	fmt.Println("...4")
	fmt.Println("accessToken:", resp.AccessToken)
	err = c.GetConfMain().SaveAccessToken(resp.AccessToken)
	if err != nil {
		c.SetErr("授权成功，保存配置失败！" + err.Error())
		return
	}

	fmt.Println("...5")
	c.SetInfo("授权成功！")
}
