package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/initialization"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func init() {
	initialization.Initialize(AppEngineLog{})
}

type AppEngineLog struct{}

func (l AppEngineLog) Debugf(r *http.Request, format string, v ...interface{}) {
	log.Debugf(GetContext(r), format, v...)
}

func (l AppEngineLog) Infof(r *http.Request, format string, v ...interface{}) {
	log.Infof(GetContext(r), format, v...)
}

func (l AppEngineLog) Warningf(r *http.Request, format string, v ...interface{}) {
	log.Warningf(GetContext(r), format, v...)
}

func (l AppEngineLog) Errorf(r *http.Request, format string, v ...interface{}) {
	log.Errorf(GetContext(r), format, v...)
}

func (l AppEngineLog) Criticalf(r *http.Request, format string, v ...interface{}) {
	log.Criticalf(GetContext(r), format, v...)
}

func GetContext(r *http.Request) context.Context {
	if r == nil {
		return appengine.BackgroundContext()
	}
	return appengine.NewContext(r)
}
