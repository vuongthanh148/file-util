package cmd

import (
	"kkcompany/internal/checksum"

	"github.com/spf13/cobra"
)

var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Compute file checksum",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		md5Flag, _ := cmd.Flags().GetBool("md5")
		sha1Flag, _ := cmd.Flags().GetBool("sha1")
		sha256Flag, _ := cmd.Flags().GetBool("sha256")
		checksum.ComputeChecksum(file, md5Flag, sha1Flag, sha256Flag)
	},
}

func init() {
	checksumCmd.Flags().StringP("file", "f", "", "the input file")
	checksumCmd.Flags().Bool("md5", false, "compute md5 checksum")
	checksumCmd.Flags().Bool("sha1", false, "compute sha1 checksum")
	checksumCmd.Flags().Bool("sha256", false, "compute sha256 checksum")
}
