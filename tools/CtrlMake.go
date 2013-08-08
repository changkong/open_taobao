package main

import (
	"strings"
	"fmt"
	"net/http"
)

type CtrlMake struct {
	UtilPage
}

func (c *CtrlMake) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" [url]", r.URL.String())
	if r.RequestURI == "/make/" {
		c.ShowPage(w, c, "make.tpl")
		return
	}
	if strings.HasPrefix(r.RequestURI, "/make/do") {
		c.makeDo(w, r)
	}
	c.Redirect(w, "/make/")
}

func (c *CtrlMake) makeDo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		c.SetErr(err.Error())
		return
	}

	choose := r.Form["PkgChoose"]
	for _, v := range c.GetConfPackage().Mx {
		v.PkgChoose = false
		for _, vv := range choose {
			if vv == v.Name {
				v.PkgChoose = true
				continue
			}
		}
	}

	errArray := MakeApis(c.GetConfPackage(), c.GetData())
	if errArray == nil || len(*errArray) == 0 {
		c.SetInfo("make ok!")
		return
	}

	c.SetInfo("")
	for _, e := range *errArray {
		c.AddErr(e)
	}
}
