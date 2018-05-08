package cmd

import (
	"fmt"

	"github.com/plombardi89/klusternaut/pkg/version"
	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Short: "Verify the Klusternaut version",
		Long: "Use this command to check the version of Klusternaut.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", version.GetVersionJSON())
		},
	}
}
