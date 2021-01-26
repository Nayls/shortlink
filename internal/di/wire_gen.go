// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/context"
	"github.com/batazor/shortlink/internal/di/internal/flags"
	"github.com/batazor/shortlink/internal/di/internal/logger"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	"github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/di/internal/traicing"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	store2 "github.com/batazor/shortlink/internal/services/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"net/http"
)

// Injectors from default.go:

func InitializeFullService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler)
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup6, err := store.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := profiling.New()
	autoMaxProAutoMaxPro, cleanup7, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewFullService(context, configConfig, logger, mq, handler, serveMux, tracer, dbStore, pprofEndpoint, autoMaxProAutoMaxPro, rpcServer, clientConn)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_api.go:

func InitializeAPIService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler)
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup6, err := store.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkStore, err := InitLinkStore(context, logger, dbStore)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := profiling.New()
	autoMaxProAutoMaxPro, cleanup7, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup8, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup9, err := rpc.InitClient(logger, tracer)
	if err != nil {
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewAPIService(context, configConfig, logger, mq, handler, serveMux, tracer, dbStore, linkStore, pprofEndpoint, autoMaxProAutoMaxPro, rpcServer, clientConn)
	if err != nil {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup9()
		cleanup8()
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_bot.go:

func InitializeBotService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	mq, cleanup3, err := mq_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup4, err := sentry.New()
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler)
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewBotService(logger, mq, serveMux, autoMaxProAutoMaxPro)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_logger.go:

func InitializeLoggerService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, cleanup3, err := sentry.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler)
	mq, cleanup4, err := mq_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerService(logger, serveMux, mq, autoMaxProAutoMaxPro)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// Injectors from service_metadata.go:

func InitializeMetadataService() (*Service, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup3, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	dbStore, cleanup4, err := store.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracer, cleanup5, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	rpcServer, cleanup6, err := rpc.InitServer(logger, tracer)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	metaStore, err := InitMetaStore(context, logger, dbStore)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	handler, cleanup7, err := sentry.New()
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := monitoring.New(handler)
	service, err := NewMetadataService(logger, autoMaxProAutoMaxPro, dbStore, rpcServer, metaStore, serveMux, handler)
	if err != nil {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup7()
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// default.go:

// Service - heplers
type Service struct {
	Ctx           context.Context
	Cfg           *config.Config
	Log           logger.Logger
	Tracer        *opentracing.Tracer
	Sentry        *sentryhttp.Handler
	DB            *db.Store
	LinkStore     *store2.LinkStore
	MetaStore     *meta_store.MetaStore
	MQ            mq.MQ
	ServerRPC     *rpc.RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(ctx.New, autoMaxPro.New, flags.New, config.New, logger_di.New, traicing_di.New)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService, store.New, sentry.New, monitoring.New, profiling.New, mq_di.New, rpc.InitServer, rpc.InitClient,
)

func NewFullService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,

	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:           ctx2,
		Cfg:           cfg,
		Log:           log,
		MQ:            mq2,
		Tracer:        tracer,
		Monitoring:    monitoring2,
		Sentry:        sentryHandler,
		DB:            db2,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

// service_api.go:

// InitLinkStore =======================================================================================================
func InitLinkStore(ctx2 context.Context, log logger.Logger, conn *db.Store) (*store2.LinkStore, error) {
	st := store2.LinkStore{}
	linkStore, err := st.Use(ctx2, log, conn)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	DefaultSet, store.New, InitLinkStore, sentry.New, monitoring.New, profiling.New, mq_di.New, rpc.InitServer, rpc.InitClient, NewAPIService,
)

func NewAPIService(ctx2 context.Context,

	cfg *config.Config,
	log logger.Logger, mq2 mq.MQ,

	sentryHandler *sentryhttp.Handler, monitoring2 *http.ServeMux,
	tracer *opentracing.Tracer, db2 *db.Store,
	linkStore *store2.LinkStore,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:           ctx2,
		Log:           log,
		MQ:            mq2,
		Tracer:        tracer,
		Monitoring:    monitoring2,
		Sentry:        sentryHandler,
		DB:            db2,
		LinkStore:     linkStore,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

// service_bot.go:

// BotService ==========================================================================================================
var BotSet = wire.NewSet(
	DefaultSet, sentry.New, monitoring.New, mq_di.New, NewBotService,
)

func NewBotService(
	log logger.Logger, mq2 mq.MQ, monitoring2 *http.ServeMux,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Log:        log,
		MQ:         mq2,
		Monitoring: monitoring2,
	}, nil
}

// service_logger.go:

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(
	DefaultSet, mq_di.New, sentry.New, monitoring.New, NewLoggerService,
)

func NewLoggerService(
	log logger.Logger, monitoring2 *http.ServeMux, mq2 mq.MQ,

	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		Log:        log,
		MQ:         mq2,
		Monitoring: monitoring2,
	}, nil
}

// service_metadata.go:

// InitMetaStore =======================================================================================================
func InitMetaStore(ctx2 context.Context, log logger.Logger, conn *db.Store) (*meta_store.MetaStore, error) {
	st := meta_store.MetaStore{}
	metaStore, err := st.Use(ctx2, log, conn)
	if err != nil {
		return nil, err
	}

	return metaStore, nil
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(
	DefaultSet, store.New, rpc.InitServer, InitMetaStore, sentry.New, monitoring.New, NewMetadataService,
)

func NewMetadataService(
	log logger.Logger,
	autoMaxProcsOption autoMaxPro.AutoMaxPro, db2 *db.Store,
	serverRPC *rpc.RPCServer,
	metaStore *meta_store.MetaStore, monitoring2 *http.ServeMux,
	sentryHandler *sentryhttp.Handler,
) (*Service, error) {
	return &Service{
		Log:        log,
		ServerRPC:  serverRPC,
		DB:         db2,
		MetaStore:  metaStore,
		Monitoring: monitoring2,
		Sentry:     sentryHandler,
	}, nil
}
