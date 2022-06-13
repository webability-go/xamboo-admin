package main

import (
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms/context"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	loadinglogo, _ := ctx.Sysparams.GetString("loadinglogo")
	loadingbgcolor, _ := ctx.Sysparams.GetString("loadingbgcolor")
	loadingbordercolor, _ := ctx.Sysparams.GetString("loadingbordercolor")
	loadingcolor, _ := ctx.Sysparams.GetString("loadingcolor")

	params := &xcore.XDataset{
		"LOADINGLOGO":        loadinglogo,
		"LOADINGBGCOLOR":     loadingbgcolor,
		"LOADINGBORDERCOLOR": loadingbordercolor,
		"LOADINGCOLOR":       loadingcolor,

		"#": language,
	}
	css, _ := ctx.Sysparams.GetString("maincss")
	if css != "" {
		params.Set("CSS", "WA.Managers.ondemand.loadCSS('"+css+"', cssloaded, false);")
		params.Set("NUMCSS", "4")
	} else {
		params.Set("NUMCSS", "3")
	}

	return template.Execute(params)
}
