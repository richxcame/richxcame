/*
Copyright © 2025 Baygeldi Cholukov <baygeldicolukov@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "richxcame",
	Short: "Personal CLI tool",
	Long: `Info about me and other stuff that I use a lot. For example:
info    - Info about me.
passgen - Password generator and save it to clipboard.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
