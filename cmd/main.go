package main

import (
	"fmt"
	"os"

	"goapi/internal/config"
	"goapi/internal/handler"
	"goapi/internal/svc"

	"github.com/spf13/cobra"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile string
	version    = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:     "goapi",
	Short:   "A RESTful API service built with go-zero framework",
	Version: version,
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Start the user service",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(configFile, &c)

		server := rest.MustNewServer(c.RestConf)
		defer server.Stop()

		ctx := svc.NewServiceContext(c)
		handler.RegisterHandlers(server, ctx)

		fmt.Printf("Starting user service at %s:%d...\n", c.Host, c.Port)
		server.Start()
	},
}

var productCmd = &cobra.Command{
	Use:   "product",
	Short: "Start the product service",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(configFile, &c)

		server := rest.MustNewServer(c.RestConf)
		defer server.Stop()

		ctx := svc.NewServiceContext(c)
		handler.RegisterHandlers(server, ctx)

		fmt.Printf("Starting product service at %s:%d...\n", c.Host, c.Port)
		server.Start()
	},
}

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Start the order service",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(configFile, &c)

		server := rest.MustNewServer(c.RestConf)
		defer server.Stop()

		ctx := svc.NewServiceContext(c)
		handler.RegisterHandlers(server, ctx)

		fmt.Printf("Starting order service at %s:%d...\n", c.Host, c.Port)
		server.Start()
	},
}

func init() {
	userCmd.Flags().StringVarP(&configFile, "config", "f", "etc/user-api.yaml", "config file path")
	productCmd.Flags().StringVarP(&configFile, "config", "f", "etc/product-api.yaml", "config file path")
	orderCmd.Flags().StringVarP(&configFile, "config", "f", "etc/order-api.yaml", "config file path")

	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(productCmd)
	rootCmd.AddCommand(orderCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
