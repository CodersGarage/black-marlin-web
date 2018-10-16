package cmd

import (
	"context"
	"github.com/CodersGarage/black-marlin-web/api"
	"github.com/CodersGarage/black-marlin-web/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve starts http and gRPC server",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	addr := ":8080"

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)

	hServer := http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	go func() {
		log.Logger().Infoln("Http server has been started on ", addr)
		if err := hServer.ListenAndServe(); err != nil {
			log.Logger().Errorln("Failed to start http server on ", err)
			os.Exit(-1)
		}
	}()

	<-stop

	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	hServer.Shutdown(ctx)

	log.Logger().Infoln("Http server has been shutdown gracefully")
}
