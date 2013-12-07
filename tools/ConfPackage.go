// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ConfPackageMxT struct {
	Name      string
	FullName  string
	Caption   string
	Desc      string
	Prefix    []string
	PkgChoose bool
}

type ConfPackageT struct {
	Mx       []*ConfPackageMxT
	ConfFile string
}

func NewConfPackage(confFile string) (*ConfPackageT, error) {
	fmt.Println("read", confFile)

	var conf ConfPackageT
	body, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	conf.ConfFile = confFile

	var mx []*ConfPackageMxT
	err = json.Unmarshal(body, &mx)
	if err != nil {
		return nil, err
	}

	for _, v := range mx {
		v.FullName = "github.com/yaofangou/open_taobao/api/" + v.Name
		v.PkgChoose = true
	}

	conf.Mx = mx

	return &conf, nil
}
