package api_domain

import (
	"github.com/batazor/shortlink/internal/pkg/notify"
)

var (
	// Link CRUD methods
	METHOD_ADD    = notify.NewEventID()
	METHOD_GET    = notify.NewEventID()
	METHOD_LIST   = notify.NewEventID()
	METHOD_UPDATE = notify.NewEventID()
	METHOD_DELETE = notify.NewEventID()

	// Link CQRS methods
	METHOD_CQRS_GET = notify.NewEventID()
)
