package cmd

import (
	"fmt"
	"kkcompany/internal/version"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version info",
	Run: func(cmd *cobra.Command, args []string) {
		cmdVersion := exec.Command("./update_version.sh")
		cmdVersion.Stdout = os.Stdout
		cmdVersion.Stderr = os.Stderr
		err := cmdVersion.Run()
		if err != nil {
			fmt.Println("Error running update_version.sh:", err)
			os.Exit(1)
		}
		version.ShowVersion()
	},
}
