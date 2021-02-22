package main

import (
	"fmt"

	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms/context"
	bridgeuser "github.com/webability-go/xmodules/user/bridge"

	"admin/app/bridge"
)

const (
	ACCESS          = "_adminmenu"
	LOCALDATASOURCE = "adminmenudatasource"
	PREFIX          = "adminmenu"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := bridge.Setup(ctx, bridge.USER)
	if !ok {
		return ""
	}

	// Check security
	ds := bridge.GetDatasource(ctx, LOCALDATASOURCE)
	if !ok {
		return "Error: no datasource"
	}

	userkey := bridge.GetUserKey(ctx)
	acc := bridgeuser.ModuleUser.HasAccess(ds, userkey, ACCESS, "")
	if !acc {
		return "Error: no access"
	}

	bridgeuser.ModuleUser.SetUserParam(ds, userkey, PREFIX+"_group", "VALUE1")

	// get default group
	grp := bridgeuser.ModuleUser.GetUserParam(ds, userkey, PREFIX+"_group")
	fmt.Println(grp)

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{

		"#": language,
	}

	return template.Execute(params)
}
