package bridge

import (
	"github.com/webability-go/xamboo/applications"
	"github.com/webability-go/xamboo/cms/context"

	bridgebase "github.com/webability-go/xmodules/base/bridge"
)

func GetDatasource(ctx *context.Context, datasourcename string) applications.Datasource {

	return bridgebase.ModuleBase.TryDatasource(ctx, datasourcename)
}
