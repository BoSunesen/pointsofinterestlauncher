package pointsofinterestlauncher

import (
	"github.com/BoSunesen/pointsofinterest/webapi/factories"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/delay"
	"google.golang.org/appengine/urlfetch"
	"net/http"
)

type AppengineWorkerFactory struct{}

func (f AppengineWorkerFactory) CreateBackgroundWorker(key string, workFunction func(context.Context)) factories.BackgroundWorker {
	delayedFunction := delay.Func(key, workFunction)
	return AppEngineBackgroundWorker{delayedFunction: delayedFunction}
}

type AppEngineBackgroundWorker struct {
	delayedFunction *delay.Function
}

func (worker AppEngineBackgroundWorker) DoWork(ctx context.Context) error {
	return worker.delayedFunction.Call(ctx)
}

type AppEngineContextFactory struct{}

func (f AppEngineContextFactory) CreateContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

type AppEngineClientFactory struct{}

func (f AppEngineClientFactory) CreateClient(ctx context.Context) *http.Client {
	return urlfetch.Client(ctx)
}
