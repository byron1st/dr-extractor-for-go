package cmd

import (
	"fmt"
	"os"

	"github.com/byron1st/dr-extractor-for-go/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFlag = "config"

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract dependency relations of a target Go project",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := lib.SetConfig(viper.GetString(configFlag))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		for _, targetSourceCode := range config.TargetSourceCodes {
			if err := lib.ExtractCallgraph(targetSourceCode.MainPkgName, targetSourceCode.SourceCodePkgNames); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.Flags().StringP(configFlag, "c", "config.json", "Configuration")
	if err := viper.BindPFlag(configFlag, extractCmd.Flags().Lookup(configFlag)); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
