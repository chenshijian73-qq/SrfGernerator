package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var bannerBase64 = "ICAgX19fX18gX19fXyBfX19fX18KICAvIF9fXy8vIF9fIFxfICBfXy8KICBcX18gXC8gL18vIC8vIC8gICAKIF9fXy8gLyBfLCBfLy8gLyAgICAKL19fX18vXy8gfF98L18vICAgIAo="

var versionTpl = `%c[%d;%d;%dm%s%c[0m
Name: srf
Version: %s
BuildDate: %s
Arch: %s
CommitID: %s
`

var (
	Version   string
	BuildDate string
	CommitID  string
)

var (
	showVersion bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
		fmt.Printf(versionTpl, 0x1B, 0, 0, 34, banner, 0x1B, Version, BuildDate, runtime.GOOS+"/"+runtime.GOARCH, CommitID)
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	rootCmd.AddCommand(versionCmd)
}
