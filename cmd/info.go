/*
Copyright © 2025 Baygeldi Cholukov <baygeldicolukov@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Info about richxcame",
	Run: func(cmd *cobra.Command, args []string) {
		Box := box.New(box.Config{
			Type:  "Classic",
			Px:    2,
			Py:    1,
			Color: "Cyan",
		})
		Box.Print("", "richxcame")

		fmt.Println(`
Hi 👋, My name is Baygeldi. Software engineer who ❤️ VIM

About me:
  🥷  Coding ninja
  ☕️  Coffee geek
  🎧  Music enthusiast
  🧠  Math lover


Contacts:
  Email:  baygeldicholukov@gmail.com
  Phone:  +993 62726535
  Github: richxcame
			`)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
