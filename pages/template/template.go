package main

import (
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms/context"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	// Va a buscar los datos de la página
	// JS: core mandatory load for every page
	jss := []string{
		// WAJAF
		"/js?js=core.js",
		"/js?js=corebrowser.js",
		"/js?js=coreext.js",
		"/js?js=coretemplate.js",
		"/js?js=eventManager.js",
		"/js?js=ajaxManager.js",
		"/js?js=ondemandManager.js",
		"/js?js=ddManager.js",
		"/js?js=soundManager.js",
		"/js?js=animManager.js",
		"/js?js=helpManager.js",
		"/js?js=wa4glManager.js",
	}
	jsdata := ""
	for _, file := range jss {
		jsdata += "<script src=\"" + file + "\" /></script>\n"
	}
	js, _ := ctx.Sysparams.GetString("mainjs")
	if js != "" {
		jsdata += "<script src=\"" + js + "\" /></script>\n"
	}

	title, _ := ctx.Sysparams.GetString("maintitle")
	if title != "" {
		language.Set("title", title)
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"LANGUAGE": ctx.Language,
		"JS":       jsdata,

		"#": language,
	}

	return template.Execute(params)
}
