package main

import (
	"github.com/njayp/replik/pkg/grpc/client"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd().Execute()
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.AddCommand(getCmd())
	return cmd
}

func getCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "get",
		Run:  Get,
		Args: cobra.ExactArgs(1),
	}
}

func Get(cmd *cobra.Command, args []string) {
	client.Get(cmd.Context(), args[0])
}
