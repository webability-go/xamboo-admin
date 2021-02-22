package main

import (
	"net/http"

	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/utils"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	// read file if syntax .../WxH/...webp
	original := ctx.Request.URL.Path

	if original == "" || original == "/" {
		return s.(*cms.CMS).Run("home/home", true, nil, "", "", "")
	}

	alternatepublic, _ := ctx.Sysparams.GetString("alternatepublic")
	alternatepages, _ := ctx.Sysparams.GetString("alternatepages")

	if alternatepublic != "" && utils.FileExists(alternatepublic+original) {
		ctx.Request.RequestURI = original
		ctx.Request.URL.Path = original
		http.FileServer(http.Dir(alternatepublic)).ServeHTTP(ctx.Writer, ctx.Request)
		return nil
	}

	if alternatepages != "" {
		thispagedir := s.(*cms.CMS).PagesDir
		s.(*cms.CMS).PagesDir = alternatepages
		data := s.(*cms.CMS).Run(original, false, nil, "", "", "")
		s.(*cms.CMS).PagesDir = thispagedir
		return data
	}

	return "Error 404, there is no alternatepublic or alternatepage: " + alternatepublic + original + " " + alternatepages + original
}
