package main

import (
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"

	"github.com/webability-go/xmodules/user/security"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}

	application := "true"
	userkey, _ := ctx.Sessionparams.GetString("userkey")
	if userkey == "" {
		userkey = "0"
		application = "false"
	}
	username, _ := ctx.Sessionparams.GetString("username")

	loginlogo, _ := ctx.Sysparams.GetString("loginlogo")
	welcome, _ := ctx.Sysparams.GetString("loginwelcome")
	if welcome != "" {
		language.Set("welcome", welcome)
	}
	version := config.Config.Version
	sversion, _ := ctx.Sysparams.GetString("VERSION")
	if sversion != "" {
		version = sversion
	}
	headerlogo, _ := ctx.Sysparams.GetString("headerlogo")
	headertitle, _ := ctx.Sysparams.GetString("headertitle")
	if headertitle != "" {
		language.Set("headertitle", headertitle)
	}
	footertitle, _ := ctx.Sysparams.GetString("footertitle")
	if footertitle != "" {
		language.Set("footertitle", footertitle)
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"APPLICATION": application,
		"SOUND":       "1",
		"HELP":        "1",
		"USER":        username,
		"USERKEY":     userkey,
		"VERSION":     version,
		"LOGINLOGO":   loginlogo,
		"HEADERLOGO":  headerlogo,

		"#": language,
	}

	return template.Execute(params)
}

func Formlogin(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}

	sessionid, _ := ctx.Sessionparams.GetString("usersessionid")
	userkey, _ := ctx.Sessionparams.GetString("userkey")
	username, _ := ctx.Sessionparams.GetString("username")
	var data map[string]interface{}
	if sessionid != "" {
		data = map[string]interface{}{
			"success":  true,
			"userkey":  userkey,
			"username": username,
			"help":     1,
			"sound":    1,
		}
	} else {
		data = map[string]interface{}{
			"success": false,
			"messages": map[string]string{
				"username": language.Get("login.wrong"),
			},
			"popup": false,
		}
	}
	return data
}

func Formpassword(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}

	// send random id
	// ask for rid
	// if match, connect the user
	return "OK"
}

func Disconnect(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, e interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}
	return "OK"
}
