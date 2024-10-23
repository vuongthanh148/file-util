package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "futil",
	Short: "File Utility",
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help about commands",
	Long:  `Help about commands`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			subCmd, _, err := rootCmd.Find(args)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			subCmd.Help()
		} else {
			rootCmd.Help()
		}
	},
}

func Execute() {
	// Handle SIGINT to stop reading from stdin
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		os.Exit(0)
	}()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(linecountCmd)
	rootCmd.AddCommand(checksumCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
