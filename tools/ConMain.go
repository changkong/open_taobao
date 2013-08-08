// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
)

var RedirectUriRegexp = regexp.MustCompile("^http://[\\w.:]+/auth/callback$")

type ConfMainT struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	RedirectUri string

	ConfFile string `json:"-"` // 不需要输出
}

func (c *ConfMainT) IsNoAccessToken() bool {
	return c.AccessToken == ""
}

func (c *ConfMainT) GetAuthMsg() string {
	if c.AccessToken == "" {
		return "用户未授权"
	}
	return ""
}

func NewConfMain(confFile string) (*ConfMainT, error) {
	fmt.Println("read", confFile)
	var confMain ConfMainT
	confMain.ConfFile = confFile

	body, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &confMain)
	if err != nil {
		return nil, err
	}

	return &confMain, nil
}

func (c *ConfMainT) CheckPara() error {
	if c.AppKey == "" {
		return errors.New("App Key 不能为空!")
	}
	if c.AppSecret == "" {
		return errors.New("App Secret 不能为空!")
	}
	if c.RedirectUri == "" {
		return errors.New("RedirectUri 不能为空!")
	}
	if !RedirectUriRegexp.MatchString(c.RedirectUri) {
		return errors.New("测试程序, 只支持 RedirectUri 为 http://域名或IP/callback 的格式")
	}

	return nil
}

func (c *ConfMainT) SavePara(appKey, appSecret, redirectUri string) error {
	c.AppKey = appKey
	c.AppSecret = appSecret
	c.RedirectUri = redirectUri

	body, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.ConfFile, body, 0600)
}

func (c *ConfMainT) SaveAccessToken(accessToken string) error {
	c.AccessToken = accessToken

	if accessToken == "" {
		return errors.New("accessToken 为空, 保存错误!")
	}

	body, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.ConfFile, body, 0600)
}
