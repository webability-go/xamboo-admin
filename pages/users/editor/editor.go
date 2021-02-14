package main

import (
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms/context"

	"kiwilimon/bridge"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	if !bridge.Setup(ctx, bridge.USER) {
		return "" // error already managed
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{}

	return template.Execute(params)
}
