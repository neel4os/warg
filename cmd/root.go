package cmd

import (
	"fmt"
	"os"

	"github.com/neel4os/warg/cmd/initialize"
	"github.com/neel4os/warg/cmd/start"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	rootCmd.AddCommand(initialize.New())
	rootCmd.AddCommand(start.New())
	cobra.OnInitialize(initconfig)
	return rootCmd
}

var rootCmd = &cobra.Command{
	Use:   "warg",
	Short: "warg is a CLI tool to manage your Warg instance",
	Long:  `warg is a CLI tool to manage your Warg instance. It builds on top of the Warg API to provide a more user-friendly way to interact with your Warg instance.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func initconfig() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
