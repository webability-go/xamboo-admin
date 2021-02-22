package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xmodules/adminmenu"
	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"
	"github.com/webability-go/xmodules/user"
	"github.com/webability-go/xmodules/useradmin"

	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms/context"
	"github.com/webability-go/xamboo/config"
)

const (
	VERSION = "0.0.1"
)

var (
	Container               *base.Container
	Application             = LocalApp{}
	DatasourceContainerName = ""
	DatasourceContainerFile = ""
)

// Called ONCE, at the loading of the plugin
func init() {
	tools.Language = language.Spanish
}

// LocalApp is needed to link the Xamboo
type LocalApp struct{}

func (la *LocalApp) StartHost(h config.Host) {
	// Any host that use THIS app should have the same datasource container
	if Container == nil {
		DatasourceContainerName, _ = h.CMS.Config.GetString("datasourcecontainername")
		DatasourceContainerFile, _ = h.CMS.Config.GetString("datasourcecontainer")
		Container = base.Create(DatasourceContainerFile)
		base.Containers.AddContainer(DatasourceContainerName, Container)
	}
}

func (la *LocalApp) StartContext(ctx *context.Context) {

	ds := Container.TryDatasource(ctx, "userdatasource")
	if ds == nil {
		return
	}
	user.StartContext(ds, ctx)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return DatasourceContainerFile
}

func (la *LocalApp) GetDatasourceSet() applications.DatasourceSet {
	return Container
}

func (la *LocalApp) GetCompiledModules() applications.ModuleSet {
	return base.ModulesList
}

// Modules linkers
var ModuleBase = base.ModuleBase
var ModuleUser = user.ModuleUser
var ModuleAdminMenu = adminmenu.ModuleAdminMenu
var ModuleUserAdmin = useradmin.ModuleUserAdmin
