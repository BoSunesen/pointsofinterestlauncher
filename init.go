package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/initialization"
)

func init() {
	initialization.Initialize(AppEngineLog{}, AppEngineContextFactory{}, AppEngineClientFactory{}, AppengineWorkerFactory{})
}
