package main

import (
	"fmt"
	"github.com/yaofangou/open_taobao/api/user"
	"net/http"
	"strings"
)

type CtrlTest struct {
	UtilPage

	UserBuyerName  string
	UserBuyerJson  string
	UserSellerName string
	UserSellerJson string
}

func (c *CtrlTest) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())

	if r.RequestURI == "/test/" {
		c.ShowPage(w, c, "test.tpl")
		return
	}

	if strings.HasPrefix(r.RequestURI, "/test/user/buyer/get") {
		c.userBuyerGet(w, r)
	} else if strings.HasPrefix(r.RequestURI, "/test/user/seller/get") {
		c.userSellerGet(w, r)
	} else {
		c.SetErr("不能识别的请求:" + r.RequestURI)
	}

	c.Redirect(w, "/test/")
}

func (c *CtrlTest) userBuyerGet(w http.ResponseWriter, r *http.Request) {
	var req user.UserBuyerGetRequest
	req.SetFields("user_id, nick")
	resp, data, err := req.GetResponse(c.GetConfMain().AccessToken)
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	c.UserBuyerName = resp.User.Nick
	c.UserBuyerJson = string(data)
	c.SetInfo("UserBuyerGet ok！")
}

func (c *CtrlTest) userSellerGet(w http.ResponseWriter, r *http.Request) {
	var req user.UserSellerGetRequest
	req.SetFields("user_id, nick")
	resp, data, err := req.GetResponse(c.GetConfMain().AccessToken)
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	c.UserSellerName = resp.User.Nick
	c.UserSellerJson = string(data)
	c.SetInfo("UserSellerGet ok！")
}
