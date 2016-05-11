package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/handlers"
	"net/http"
)

func init() {
	http.Handle("/ping/", handlers.LoggingDecorator{handlers.PingHandler{}, "Ping"})
	http.Handle("/poi/", handlers.LoggingDecorator{handlers.PoiHandler{}, "POI"})
}