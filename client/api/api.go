package api

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/eduardobcolombo/learning-grpc/portpb"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

func (e *Environment) GetGRPC() error {
	address := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	fmt.Println("Client GRPC ", address)
	cc, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	defer cc.Close()
	e.psc = portpb.NewPortServiceClient(cc)

	return nil
}

func Initialize() {
	var wait time.Duration
	e := &Environment{}
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "")
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// e.readjsonFile("../../ports.json")

	router := mux.NewRouter()
	e.GetRoutes(router)
	e.GetMiddlewares(router)

	port := os.Getenv("API_PORT")
	srv := &http.Server{
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	fmt.Println("Running client api at port:", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error trying to shutting down: %s", err)
	}
	log.Println("shutting down")
	os.Exit(0)

}
