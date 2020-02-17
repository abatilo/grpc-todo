package main

import (
	"github.com/abatilo/grpc-todo/cmd/todoclient"
	"github.com/abatilo/grpc-todo/cmd/todoservice"
	"github.com/spf13/cobra"
)

func main() {
	var (
		rootCmd = &cobra.Command{
			Use:   "todo",
			Short: "A gRPC based todo list sample",
		}
	)

	rootCmd.AddCommand(todoservice.Cmd)
	rootCmd.AddCommand(todoclient.Cmd)
	rootCmd.Execute()
}
