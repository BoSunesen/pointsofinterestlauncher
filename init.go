package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/initialization"
)

func init() {
	webApiInitializer := initialization.NewWebApiInitializer(AppEngineLog{}, AppEngineContextFactory{}, AppEngineClientFactory{}, AppengineWorkerFactory{})
	webApiInitializer.Initialize()
}
