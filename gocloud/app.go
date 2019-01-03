package gocloud

import (
	"context"
	"database/sql"
	"github.com/google/wire"
	"go.opencensus.io/trace"
	"gocloud.dev/blob"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
	"log"
	"sync"
)

// ApplicationSet is the Wire provider set for the setting up an Application that
// does not depend on the underlying platform.
var ApplicationSet = wire.NewSet(
	NewApplication,
	AppHealthChecks,
	trace.AlwaysSample,
)

// Application is the main server struct for Guestbook. It contains the state of
// the most recently read message of the day.
type Application struct {
	srv    *server.Server
	db     *sql.DB
	bucket *blob.Bucket

	// The following fields are protected by mu:
	mu   sync.RWMutex
	motd string // message of the day
}

// newApplication creates a new Application struct based on the backends and the message
// of the day variable.
func NewApplication(srv *server.Server, db *sql.DB, bucket *blob.Bucket, motdVar *runtimevar.Variable) *Application {
	app := &Application{
		srv:    srv,
		db:     db,
		bucket: bucket,
	}
	go app.WatchVariable(motdVar)
	return app
}

// WatchVariable listens for changes in v and updates the app's message of the
// day. It is run in a separate goroutine.
func (app *Application) WatchVariable(v *runtimevar.Variable) {
	ctx := context.Background()
	for {
		snap, err := v.Watch(ctx)
		if err != nil {
			log.Printf("watch runtime variable: %v", err)
			continue
		}
		log.Println("updated runtime to", snap.Value)
		app.mu.Lock()
		app.motd = snap.Value.(string)
		app.mu.Unlock()
	}
}
