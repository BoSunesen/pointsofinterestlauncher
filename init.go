package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/factories"
	"github.com/BoSunesen/pointsofinterest/webapi/initialization"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

func init() {
	initialization.Initialize(AppEngineLog{}, AppEngineContextFactory{}, AppEngineClientFactory{}, AppengineWorkerFactory{})
}

//TODO Move stuff out of init.go
type AppengineWorkerFactory struct{}

func (f AppengineWorkerFactory) CreateBackgroundWorker(key string, workFunction func(context.Context)) factories.BackgroundWorker {
	delayedFunction := delay.Func(key, workFunction)
	return AppEngineBackgroundWorker{delayedFunction: delayedFunction}
}

type AppEngineBackgroundWorker struct {
	delayedFunction *delay.Function
}

func (worker AppEngineBackgroundWorker) DoWork(r *http.Request) error {
	return worker.delayedFunction.Call(appengine.NewContext(r))
}

type AppEngineContextFactory struct{}

func (f AppEngineContextFactory) CreateContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

type AppEngineClientFactory struct{}

func (f AppEngineClientFactory) CreateClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}

type AppEngineLog struct{}

func (l AppEngineLog) Debugf(ctx context.Context, format string, v ...interface{}) {
	log.Debugf(ctx, format, v...)
}

func (l AppEngineLog) Infof(ctx context.Context, format string, v ...interface{}) {
	log.Infof(ctx, format, v...)
}

func (l AppEngineLog) Warningf(ctx context.Context, format string, v ...interface{}) {
	log.Warningf(ctx, format, v...)
}

func (l AppEngineLog) Errorf(ctx context.Context, format string, v ...interface{}) {
	log.Errorf(ctx, format, v...)
}

func (l AppEngineLog) Criticalf(ctx context.Context, format string, v ...interface{}) {
	log.Criticalf(ctx, format, v...)
}
