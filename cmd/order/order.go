package order

import (
	"fmt"

	"goapi/internal/config"
	"goapi/internal/handler"
	"goapi/internal/svc"

	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	configFile string
)

var Cmd = &cobra.Command{
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
	Cmd.Flags().StringVarP(&configFile, "config", "f", "etc/order-api.yaml", "config file path")
}
