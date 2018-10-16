package cmd

import (
	"fmt"
	"os"

	"github.com/CodersGarage/black-marlin-web/log"
	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of black-marlin-web service
	RootCmd = &cobra.Command{
		Use:   "black-marlin-web",
		Short: "black-marlin-web is a grpc/http service",
		Long:  `An gRPC/HTTP JSON API backend service of black-marlin-web`,
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
}

// Execute executes the root command
func Execute() {
	log.Init()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
