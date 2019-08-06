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
	values := []string{"Top 10", "Player Search",  "Player Titles", "Exit"}
	fmt.Print(utils.CenterHor("asdf"))
	// templates := &promptui.SelectTemplates{
	// 	Help:     utils.CenterText("j(up) k(down)"),
	// 	Label:    utils.CenterText("{{ . }}"),
	// 	Active:   utils.CenterText("-> {{. | cyan }}"),
	// 	Inactive: utils.CenterText("{{ . | red }}"),
	// 	Selected: utils.CenterText("-> {{. | cyan }}"),
	// }
	prompt := promptui.Select{
		Label: "Enter Option",
		Items: values[:],
		// Templates: templates,
	}

	utils.CallClear()
	_, res, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	switch res {
	case "Top 10":
		Top10()
	case "Player Search":
		PlayerOverviewFunc()
	case "Player Titles":
		Titles()
	default:
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	// if idx == 0 {
	// 	Top10()
	// } else if idx == 1 {
	// 	PlayerOverviewFunc()
	// } else if idx == 2 {
	// 	Titles()
	// } else {
	// 	fmt.Println("Exiting...")
	// 	os.Exit(1)
	// }
}

var menuCommand = &cobra.Command{
	Use:   "menu",
	Short: "Opens application menu",
	Long:  "Opens application menu",
	Run: func(cmd *cobra.Command, args []string) {
		Menu()
	},
}
