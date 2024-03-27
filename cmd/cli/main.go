package main

import (
	"github.com/njayp/replik/pkg/grpc/client"
	"github.com/spf13/cobra"
)

func main() {
	cmd().Execute()
}

func cmd() *cobra.Command {
	cmd := &cobra.Command{}
	cmd.AddCommand(sub())
	return cmd
}

func sub() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			client := client.NewClient()
			files, err := client.List(ctx, args[0])
			if err != nil {
				panic(err)
			}
			for _, file := range files.Files {
				// TODO create waitgroup and wait until all finish. maybe flush manager
				err := client.File(ctx, file.Path)
				if err != nil {
					println(err.Error())
				}
			}
		},
	}
}
