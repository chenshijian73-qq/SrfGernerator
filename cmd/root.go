package cmd

import (
	"fmt"
	"github.com/chenshijian73-qq/cobra_scaffold/internal/logic"
	"github.com/spf13/cobra"
	"os"
)

var (
	interval int
	file     *string
	outSrt   *string
)

var rootCmd = &cobra.Command{
	Use:   "srf -f [Name] -t [Time] -o [Output]",
	Short: "Generate srt for a given document and interval",
	Run: func(cmd *cobra.Command, args []string) {
		logic.SrtGernerate(interval, *file, *outSrt)
	},
}

func init() {
	// 非持久化参数
	rootCmd.Flags().IntVarP(&interval, "time", "t", 0, "input interval time（Unit: second）")
	file = rootCmd.Flags().StringP("file", "f", "", "input filename")
	outSrt = rootCmd.Flags().StringP("out", "o", "output.srt", "output filename")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
