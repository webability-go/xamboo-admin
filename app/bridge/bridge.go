package bridge

import (
	"net/http"
	"plugin"

	"github.com/webability-go/xamboo/assets"

	adminmenubridge "github.com/webability-go/xmodules/adminmenu/bridge"
	basebridge "github.com/webability-go/xmodules/base/bridge"
	userbridge "github.com/webability-go/xmodules/user/bridge"
	useradminbridge "github.com/webability-go/xmodules/useradmin/bridge"
)

/* This package declare all the available functions of the app to be able to call them. */
/* Include this package when you want to call the app */

const (
	// DOES NOT MATTER IF THE USER OR CLIENT IS OR NOT CONNECTED
	ANY = 1
	// THE USER MUST BE CONNECTED TO USE THE BRIDGE
	USER = 2
)

var (
	linked bool = false
)

// Setup fabrica el enlace bridge-modulo SO listo para usar. Verifica luego enlace de funciones, verifica login clientes y usuarios, verifica idionas y deviles.
// Es la primera funcion que hay que llamar cuando se usa el bridge
func Setup(ctx *assets.Context, connection int) bool {

	// is appname is empty: search for "app" entry in ctx
	appname, _ := ctx.Sysparams.GetString("adminapp")
	if appname == "" {
		http.Error(ctx.Writer, "Admin Library name is not available in config file (parameter adminapp missing)", http.StatusInternalServerError)
		return false
	}

	// Ask for the plugins we need
	app, ok := ctx.Plugins[appname]
	if !ok {
		// 500 internal error
		http.Error(ctx.Writer, "Library admin/app is not available", http.StatusInternalServerError)
		return false
	}

	// Initialize the plugin (just in case)
	err := Start(ctx, app)
	if err != nil {
		// 500 internal error
		http.Error(ctx.Writer, "Library admin/app could not link with system", http.StatusInternalServerError)
		return false
	}

	// modules installed minimum must be base, user, useradmin

	// Verification of session is done during StartContext

	switch connection {
	case ANY:
	case USER:
		sessionid, _ := ctx.Sessionparams.GetString("usersessionid")
		if sessionid == "" {
			http.Error(ctx.Writer, "Error, user not connected", http.StatusUnauthorized)
			return false
		}
	}
	return true
}

func Start(ctx *assets.Context, lib *plugin.Plugin) error {
	if !linked {
		err := link(lib)
		if err != nil {
			return err
		}
	}

	return nil
}

func link(lib *plugin.Plugin) error {

	err := basebridge.Link(lib)
	if err != nil {
		return err
	}
	err = userbridge.Link(lib)
	if err != nil {
		return err
	}
	err = adminmenubridge.Link(lib)
	if err != nil {
		return err
	}
	err = useradminbridge.Link(lib)
	if err != nil {
		return err
	}

	linked = true
	return nil
}
