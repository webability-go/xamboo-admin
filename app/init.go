package main

import (
	"golang.org/x/text/language"

	"github.com/webability-go/xmodules/adminmenu"
	"github.com/webability-go/xmodules/base"
	"github.com/webability-go/xmodules/tools"
	"github.com/webability-go/xmodules/user"
	"github.com/webability-go/xmodules/useradmin"

	"github.com/webability-go/xamboo/assets"
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

func (la *LocalApp) StartHost(h assets.Host) {
	// Any host that use THIS app should have the same datasource container
	if Container == nil {
		DatasourceContainerName, _ = h.Config.GetString("datasourcecontainername")
		DatasourceContainerFile, _ = h.Config.GetString("datasourcecontainer")
		Container = base.Create(DatasourceContainerFile)
		base.Containers.AddContainer(DatasourceContainerName, Container)
	}
}

func (la *LocalApp) StartContext(ctx *assets.Context) {

	ds := Container.TryDatasource(ctx, "userdatasource")
	if ds == nil {
		return
	}
	user.StartContext(ds, ctx)
}

func (la *LocalApp) GetDatasourcesConfigFile() string {
	return DatasourceContainerFile
}

func (la *LocalApp) GetDatasourceSet() assets.DatasourceSet {
	return Container
}

func (la *LocalApp) GetCompiledModules() assets.ModuleSet {
	return base.ModulesList
}

// Modules linkers
var ModuleBase = base.ModuleBase
var ModuleUser = user.ModuleUser
var ModuleAdminMenu = adminmenu.ModuleAdminMenu
var ModuleUserAdmin = useradmin.ModuleUserAdmin
