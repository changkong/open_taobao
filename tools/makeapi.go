// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func createGoFile(data TemplateData, tplfile, gofile string) error {
	targetFile := filepath.Join(data.PkgPath, gofile)
	fmt.Println("    create Go file:", targetFile)

	tpl, err := ioutil.ReadFile(filepath.Join(data.ViewPath, tplfile))
	if err != nil {
		return err
	}
	t1 := template.New("t1")
	t1, err = t1.Parse(string(tpl))
	if err != nil {
		return err
	}

	err = os.MkdirAll(data.PkgPath, 0600)
	if err != nil {
		return err
	}

	f, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer f.Close()
	t1.Execute(f, data)

	return nil
}

func MakePkg(tData TemplateData) error {
	fmt.Println("    [api count] =", len(tData.Apis), "    [struct count] =", len(tData.Structs))

	err := createGoFile(tData, "api.tpl", tData.PkgName+"_apis.go")
	if err != nil {
		return err
	}

	err = createGoFile(tData, "struct.tpl", tData.PkgName+"_structs.go")
	if err != nil {
		return err
	}

	err = createGoFile(tData, "doc.tpl", tData.PkgName+"_doc.go")
	if err != nil {
		return err
	}

	fmt.Println("    format:", tData.PkgFullName)
	cmd := exec.Command("go", "fmt", tData.PkgFullName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("    install:", tData.PkgFullName)
	cmd = exec.Command("go", "install", tData.PkgFullName)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil

}

type TemplateData struct {
	MetaVersionNo string
	PkgName       string
	PkgFullName   string
	PkgPath       string
	PkgDesc       string
	ViewPath      string
	Apis          []*ApiT
	Structs       []*StructT
}

func MakeApis(confPackage *ConfPackageT, data *DataT) *[]string {

	fmt.Println("\n----------------------------------------------")
	fmt.Println("make apis")
	fmt.Println("----------------------------------------------")

	errArray := make([]string, 0)

	var tdata TemplateData
	tdata.MetaVersionNo = data.MetaVersionNo

	for _, v := range confPackage.Mx {
		if !v.PkgChoose {
			continue
		}

		tdata.PkgFullName = "github.com/yaofangou/open_taobao/api/" + v.Name
		fmt.Println("  make:", tdata.PkgFullName)

		tdata.PkgDesc = v.Desc
		tdata.PkgName = v.Name
		tdata.PkgPath = "../api/" + v.Name
		tdata.ViewPath = "./template/"
		tdata.Apis = data.MapPkgApi[v.Name]
		tdata.Structs = data.MapPkgStruct[v.Name]

		err := MakePkg(tdata)
		if err != nil {
			fmt.Println("      [error]", err)
			errArray = append(errArray, "package ["+tdata.PkgFullName+"] make err!")
		}
	}

	return &errArray
}
