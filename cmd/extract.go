package cmd

import (
	"fmt"
	"os"

	"github.com/byron1st/dr-extractor-for-go/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract dependency relations of a target Go project",
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
