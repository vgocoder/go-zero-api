package main

import (
	"fmt"
	"os"

	"goapi/cmd/order"
	"goapi/cmd/product"
	"goapi/cmd/user"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "goapi",
	Short:   "A RESTful API service built with go-zero framework",
	Version: version,
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage different services",
}

func init() {
	serviceCmd.AddCommand(user.Cmd)
	serviceCmd.AddCommand(product.Cmd)
	serviceCmd.AddCommand(order.Cmd)

	rootCmd.AddCommand(serviceCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
