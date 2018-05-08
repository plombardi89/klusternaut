package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "klusternaut",
	Short: "Automagical Kubernetes clusters for developers",
}

func Execute()  {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(VersionCmd())
}