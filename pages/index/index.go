package main

import (
	"github.com/webability-go/xcore/v2"

	"github.com/webability-go/xamboo/cms"
	"github.com/webability-go/xamboo/cms/context"

	"github.com/webability-go/xmodules/adminmenu/assets"
	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/user/security"
)

func Run(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}

	//	bridge.EntityLog_LogStat(ctx)
	params := &xcore.XDataset{
		"#": language,
	}

	return template.Execute(params)
}

func Menu(ctx *context.Context, template *xcore.XTemplate, language *xcore.XLanguage, s interface{}) interface{} {

	ok := security.Verify(ctx, security.ANY)
	if !ok {
		return ""
	}

	Order := ctx.Request.Form.Get("Order")

	if Order == "get" {
		return getMenu(ctx, s.(*cms.CMS), language, "")
	}
	if Order == "getchildren" {
		father := ctx.Request.Form.Get("father")
		return getMenu(ctx, s.(*cms.CMS), language, father)
	}
	if Order == "openclose" {

		//    $id = $this->base->HTTPRequest->getParameter('id');
		//    $status = $this->base->HTTPRequest->getParameter('status');
		//    $this->base->security->setParameter('mastertree', $id, $status=='true'?1:0);
	}

	return map[string]string{
		"status": "OK",
	}
}

func getMenu(ctx *context.Context, s *cms.CMS, language *xcore.XLanguage, father string) map[string]interface{} {

	rows := []interface{}{}

	if father == "" {
		// Main Menu Title and reload button
		optr := map[string]interface{}{
			"id":        "maintitle",
			"template":  "maintitle",
			"loadable":  false,
			"closeable": false,
			"closed":    false,
		}
		rows = append(rows, optr)
	}

	adminmenu := assets.GetEntries(ctx)
	ds := base.TryDatasource(ctx, assets.DATASOURCE)
	options, _ := adminmenu.GetMenu(ds, "_admin", father)

	if options != nil {
		for _, option := range *options {
			template := "option_0"
			loadable := false
			call1, _ := option.GetString("call1")
			call2, _ := option.GetString("call2")
			call3, _ := option.GetString("call3")
			if call1 == "openclose" {
				if call2 == "" {
					template = "folder_0"
				} else {
					if call3 == "" {
						template = "folder_1"
					} else {
						template = "folder_2"
					}
				}
				loadable = true
			} else {
				if call2 == "" {
					template = "option_0"
				} else {
					if call3 == "" {
						template = "option_1"
					} else {
						template = "option_2"
					}
				}
			}
			id, _ := option.GetString("key")
			name, _ := option.GetString("name")
			icon1, _ := option.GetString("icon1")

			optr := map[string]interface{}{
				"id":        id,
				"template":  template,
				"name":      name,
				"icon1":     icon1,
				"loadable":  loadable,
				"closeable": loadable,
				"closed":    loadable,
			}
			if loadable {
				optr["loaded"] = false
			}
			if father != "" {
				optr["father"] = father
			}
			if call1 != "" {
				optr["call1"] = call1
			}
			if call2 != "" {
				optr["call2"] = call2
				icon2, _ := option.GetString("icon2")
				optr["icon2"] = icon2
			}
			if call3 != "" {
				optr["call3"] = call3
				icon3, _ := option.GetString("icon3")
				optr["icon3"] = icon3
			}

			rows = append(rows, optr)
		}
	}

	// All the options are in tables.
	//	options, err := bridge.

	return map[string]interface{}{
		"row": rows,
	}

}
