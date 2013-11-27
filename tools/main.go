// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	var utilPage UtilPage
	err := utilPage.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/auth/", new(CtrlAuth).Show)
	http.HandleFunc("/test/", new(CtrlTest).Show)
	http.HandleFunc("/make/", new(CtrlMake).Show)

	http.HandleFunc("/assets/", utilPage.Static)
	http.HandleFunc("/", utilPage.Home)

	fmt.Println("start 80 ......")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
