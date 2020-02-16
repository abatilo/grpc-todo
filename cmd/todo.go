package main

import (
	"github.com/abatilo/grpc-todo/cmd/fooclient"
	"github.com/abatilo/grpc-todo/cmd/fooservice"
	"github.com/spf13/cobra"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use:   "todo",
			Short: "A gRPC based todo list sample",
		}
	)

	rootCmd.AddCommand(fooservice.Cmd)
	rootCmd.AddCommand(fooclient.Cmd)
	rootCmd.Execute()
}
