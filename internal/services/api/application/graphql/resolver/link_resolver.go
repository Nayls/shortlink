package resolver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/batazor/shortlink/internal/pkg/notify"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// Link ...
func (r *Resolver) Link(ctx context.Context, args struct { //nolint unparam
	Hash *string
}) (*LinkResolver, error) {
	responseCh := make(chan interface{})

	go notify.Publish(ctx, v1.METHOD_GET, *args.Hash, &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_GET"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return &LinkResolver{
				Link: nil,
			}, err
		}
		response := r.Payload.(*v1.Link) // nolint errcheck
		return &LinkResolver{
			Link: response,
		}, err
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// Links ...
func (r *Resolver) Links(ctx context.Context, args struct {
	Filter *query.Filter
}) (*[]*LinkResolver, error) { // nolint unused
	responseCh := make(chan interface{})

	// Filter to string
	filterRaw, err := json.Marshal(args.Filter)
	if err != nil {
		err := fmt.Errorf("Error parse filter args")
		return nil, err
	}

	// Default value for filter; null -> nil
	if string(filterRaw) == "null" {
		filterRaw = nil
	}

	go notify.Publish(ctx, v1.METHOD_LIST, string(filterRaw), &notify.Callback{CB: responseCh, ResponseFilter: "RESPONSE_STORE_LIST"})

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		responses := r.Payload.([]*v1.Link) // nolint errcheck

		links := []*LinkResolver{}
		for _, item := range responses {
			links = append(links, &LinkResolver{
				Link: item,
			})
		}

		return &links, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}
