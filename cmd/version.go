package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the current version of igmpqd",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:\t%s (Commit: %s)\n", GitDescribe, GitCommit)
		fmt.Printf("Built:\t\t%s\n", time.Unix(BuildTime, 0))
		fmt.Printf("Fingerprint:\t%s/%s/%s/%s\n", runtime.Compiler, runtime.GOOS, runtime.GOARCH, runtime.Version())
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
