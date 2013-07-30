// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/changkong/open_taobao"
	"github.com/changkong/open_taobao/api/user"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Controller struct {
	OpenTaobaoVersionNo string
	PkgUserVersionNo    string
	ConfMain            *ConfMainT
	ConfPackage         *ConfPackageT
	Data                *DataT
	PageData            *PageDataT
}

func NewController() (*Controller, error) {
	confMain, err := NewConfMain("./conf/main.json")
	if err != nil {
		return nil, err
	}

	confPackage, err := NewConfPackage("./conf/package.json")
	if err != nil {
		return nil, err
	}

	data, err := NewMetadata("./conf/ApiMetadata.xml", confPackage)
	if err != nil {
		return nil, err
	}

	pageData := NewPageData()

	var ctrl Controller
	ctrl.OpenTaobaoVersionNo = open_taobao.VersionNo
	ctrl.PkgUserVersionNo = user.VersionNo
	ctrl.ConfMain = confMain
	ctrl.ConfPackage = confPackage
	ctrl.Data = data
	ctrl.PageData = pageData

	return &ctrl, nil
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	if r.RequestURI == "/" {
	} else if r.RequestURI == "/favicon.ico" {

	} else {
		c.SetErr("不能识别请求:" + r.RequestURI)
	}
	c.showPage(w)
}

func (c *Controller) Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	c.SetInfo("")
	redirect(w, "/")
}

func (c *Controller) Save(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
	} else {
		c.ConfMain.AppKey = r.Form.Get("AppKey")
		c.ConfMain.AppSecret = r.Form.Get("AppSecret")
		c.ConfMain.RedirectUri = r.Form.Get("RedirectUri")
		err = c.ConfMain.Save()
		if err != nil {
			c.SetErr(err.Error())
		} else {
			c.SetInfo("save ok!")
		}
	}
	redirect(w, "/")
}

func (c *Controller) MakeApi(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())

	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
	} else {
		choose := r.Form["PkgChoose"]
		for _, v := range c.ConfPackage.Mx {
			v.PkgChoose = false
			for _, vv := range choose {
				if vv == v.Name {
					v.PkgChoose = true
					continue
				}
			}
		}

		errArray := MakeApis(c.ConfPackage, c.Data)
		if errArray == nil || len(*errArray) == 0 {
			c.SetInfo("make ok!")
		} else {
			c.SetInfo("")
			for _, e := range *errArray {
				c.AddErr(e)
			}
		}
	}
	redirect(w, "/")
}

func (c *Controller) Auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	url_, err := open_taobao.GetUrlForAuth(c.ConfMain.AppKey, c.ConfMain.RedirectUri, "")
	if err != nil {
		c.SetErr(err.Error())
		redirect(w, "/")
	} else {
		redirect(w, url_.String())
	}
}

func (c *Controller) UserBuyerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	var req user.UserBuyerGetRequest
	req.SetFields("user_id, nick")
	resp, data, err := req.GetResponse(c.ConfMain.AccessToken)
	if err != nil {
		c.SetErr(err.Error())
	} else {
		c.PageData.UserBuyerName = resp.User.Nick
		c.PageData.UserBuyerJson = string(data)
		c.SetInfo("UserBuyerGet ok！")
	}
	redirect(w, "/")
}

func (c *Controller) UserSellerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	var req user.UserSellerGetRequest
	req.SetFields("user_id, nick")
	resp, data, err := req.GetResponse(c.ConfMain.AccessToken)
	if err != nil {
		c.SetErr(err.Error())
	} else {
		c.PageData.UserSellerName = resp.User.Nick
		c.PageData.UserSellerJson = string(data)
		c.SetInfo("UserSellerGet ok！")
	}
	redirect(w, "/")
}

func (c *Controller) Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
	} else {
		TBErr := r.Form.Get("error")
		TBErrDesc := r.Form.Get("error_description")
		if TBErr != "" {
			c.SetErr(TBErrDesc)
		} else {
			code := r.Form.Get("code")
			fmt.Println("code:", code)
			var req open_taobao.TokenGetRequest
			req.SetAppKey(c.ConfMain.AppKey)
			req.SetAppSecret(c.ConfMain.AppSecret)
			req.SetRedirectUri(c.ConfMain.RedirectUri)
			req.SetCode(code)
			req.SetState("")
			resp, _, err := req.GetResponse()
			if err != nil {
				c.SetErr(err.Error())
			} else {
				fmt.Println("accessToken:", resp.AccessToken)
				c.ConfMain.AccessToken = resp.AccessToken
				err = c.ConfMain.Save()
				if err != nil {
					c.SetErr("授权成功，保存配置失败！" + err.Error())
				} else {
					c.SetInfo("授权成功！")
				}
			}

		}
	}
	redirect(w, "/")
}

func (c *Controller) showPage(w http.ResponseWriter) {
	t1html, err := ioutil.ReadFile("./views/home.tpl")
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	t1 := template.New("t1")
	t1, err = t1.Parse(string(t1html))
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}

	t1.Execute(w, c)
}

func redirect(w http.ResponseWriter, url string) {
	w.Header().Set("Location", url)
	w.WriteHeader(302)
}

func (c *Controller) SetErr(msg string) {
	c.PageData.setErr(msg)
}

func (c *Controller) AddErr(msg string) {
	c.PageData.addErr(msg)
}

func (c *Controller) SetInfo(msg string) {
	c.PageData.setInfo(msg)
}
