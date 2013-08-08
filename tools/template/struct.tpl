// Copyright 2013 The Changkong Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package {{.PkgName}}

const VersionNo = "{{.MetaVersionNo}}"

{{range .Structs }}
/* {{.Desc}} */
type {{.Name}} struct {
{{range .Props.Prop }}	{{.GoName}} {{.GoType}} `json:"{{.Name}}"`
{{end}}
}
{{end}}
