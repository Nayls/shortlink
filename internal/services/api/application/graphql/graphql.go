package graphql

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/markbates/pkger"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/text/message"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/services/api/application/graphql/resolver"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
)

// API ...
type API struct { // nolint unused
	store db.DB
	ctx   context.Context
}

// GetHandler ...
func (api *API) GetHandler() *relay.Handler {
	buf := bytes.Buffer{}

	err := pkger.Walk("internal/services/api/application/graphql/schema", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := pkger.Open(path)
			if err != nil {
				return err
			}

			f := make([]byte, info.Size())
			_, err = file.Read(f)

			// Add a newline if the file does not end in a newline.
			if len(f) > 0 && f[len(f)-1] != '\n' {
				if errWriteByte := buf.WriteByte('\n'); err != nil {
					panic(errWriteByte)
				}
			}

			if err != nil {
				panic(err)
			}

			if _, err := buf.Write(f); err != nil {
				panic(err)
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	s := graphql.MustParseSchema(buf.String(), &resolver.Resolver{Store: api.store})
	handler := relay.Handler{Schema: s}

	return &handler
}

// Run ...
func (api *API) Run(
	ctx context.Context,
	i18n *message.Printer,
	config api_type.Config,
	log logger.Logger,
	tracer *opentracing.Tracer,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error { // nolint unparam
	api.ctx = ctx

	log.Info("Run GraphQL API")

	handler := api.GetHandler()

	http.Handle("/api/query", http.TimeoutHandler(handler, config.Timeout, `{"error":"context deadline exceeded"}`))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

	return err
}
