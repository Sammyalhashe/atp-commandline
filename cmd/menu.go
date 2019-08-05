package cmd

import (
	"fmt"
	"os"

	"github.com/Sammyalhashe/gomod/utils"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCommand)
}

// Menu creates a menu to access various commands
func Menu() {
	values := [3]string{"Top 10", "Player Search", "Exit"}
	fmt.Print(utils.CenterHor("asdf"))
	// templates := &promptui.SelectTemplates{
	// 	Help:     utils.CenterText("j(up) k(down)"),
	// 	Label:    utils.CenterText("{{ . }}"),
	// 	Active:   utils.CenterText("-> {{. | cyan }}"),
	// 	Inactive: utils.CenterText("{{ . | red }}"),
	// 	Selected: utils.CenterText("-> {{. | cyan }}"),
	// }
	prompt := promptui.Select{
		Label:     "Enter Option",
		Items:     values[:],
		// Templates: templates,
	}

	utils.CallClear()
	idx, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	if idx == 0 {
		Top10()
	} else if idx == 1 {
		PlayerOverviewFunc()
	} else {
		fmt.Println("Exiting...")
		os.Exit(1)
	}
}

var menuCommand = &cobra.Command{
	Use:   "menu",
	Short: "Opens application menu",
	Long:  "Opens application menu",
	Run: func(cmd *cobra.Command, args []string) {
		Menu()
	},
}
