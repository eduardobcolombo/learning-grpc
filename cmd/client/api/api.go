package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/mid"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/port"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/portpb"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type API struct {
	router *mux.Router
}

type CoreConfig struct {
	Port port.Core
}

func newAPI() *API {
	return &API{
		router: mux.NewRouter(),
	}
}

func (api *API) Handler() http.Handler {
	return api.router
}

func (api *API) routes(core CoreConfig, cfg *Config) {
	ver := "/v1"

	portHandler := Handler{
		core: core.Port,
	}

	api.router.HandleFunc(ver+"/ports", portHandler.RetrievePorts).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc(ver+"/ports", portHandler.UpdatePorts).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) mid(cfg *Config) {
	api.router.Use(mux.CORSMethodMiddleware(api.router))
	api.router.Use(mid.CORS())
	api.router.Use(mid.Authenticate)
}

func Initialize(log *zap.SugaredLogger) int {
	var wait time.Duration
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error(err)
		return 1
	}

	grpcClientConn, err := GRPCInit(cfg)
	if err != nil {
		log.Error(err)
		return 1
	}

	cfg.psc = portpb.NewPortServiceClient(grpcClientConn)

	core := CoreConfig{
		Port: port.NewCore(log, cfg.psc),
	}

	api := newAPI()
	api.routes(core, cfg)
	api.mid(cfg)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + cfg.APIPort,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.Handler(),
	}

	fmt.Println("Running client api at port:", cfg.APIPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	log.Error("Closing GRPC Connection")
	defer grpcClientConn.Close()

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Error trying to shutting down: %s", err)
	}
	log.Error("shutting down")

	return 0

}
