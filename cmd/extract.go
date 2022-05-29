/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/byron1st/dr-extractor-golang/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := lib.ExtractCallgraph(viper.GetString("pkg"), viper.GetString("base")); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.Flags().StringP("pkg", "p", "", "Target package")
	viper.BindPFlag("pkg", extractCmd.Flags().Lookup("pkg"))

	extractCmd.Flags().StringP("base", "b", "", "Base package")
	viper.BindPFlag("base", extractCmd.Flags().Lookup("base"))
}
