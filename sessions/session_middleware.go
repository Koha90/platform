package sessions

import (
	"context"
	"time"

	gorilla "github.com/gorilla/sessions"
	"platform/config"
	"platform/pipeline"
)

type SessionComponent struct {
	store *gorilla.CookieStore
	config.Configuration
}

func (sc *SessionComponent) Init() {
	cookiekey, found := sc.Configuration.GetString("sessions:key")
	if !found {
		panic("session key not found in configuration")
	}
	if sc.GetBoolDefault("session:cyclekey", true) {
		cookiekey += time.Now().String()
	}
	sc.store = gorilla.NewCookieStore([]byte(cookiekey))
}

func (sc *SessionComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(ctx2 *pipeline.ComponentContext)) {
	session, _ := sc.store.Get(ctx.Request, SESSION__CONTEXT_KEY)
	c := context.WithValue(ctx.Request.Context(), SESSION__CONTEXT_KEY, session)
	ctx.Request = ctx.Request.WithContext(c)
	next(ctx)
	session.Save(ctx.Request, ctx.ResponseWriter)
}
