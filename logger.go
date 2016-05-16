package pointsofinterestlauncher

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

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
