package bridge

import (
	"github.com/webability-go/xamboo/assets"

	bridgebase "github.com/webability-go/xmodules/base/bridge"
)

func GetDatasource(ctx *assets.Context, datasourcename string) assets.Datasource {

	return bridgebase.ModuleBase.TryDatasource(ctx, datasourcename)
}
