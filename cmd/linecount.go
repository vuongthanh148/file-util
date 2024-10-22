package cmd

import (
	"fmt"
	"kkcompany/internal/linecount"

	"github.com/spf13/cobra"
)

var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Printf("error: no file provided\n")
		}
		linecount.CountLines(file)
	},
}

func init() {
	linecountCmd.Flags().StringP("file", "f", "", "the input file")
}
