package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/Sammyalhashe/gomod/constants"
	"github.com/Sammyalhashe/gomod/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(top10Command)
}

type top10 struct {
	Top10 []string `json:"top_10"`
}

// Top10 gets the top 10 and executes a prompt
func Top10() {

	utils.CallClear()
	c := make(chan bool)
	go utils.StartLoading(c)
	res, err := http.Get(constants.APIHead + constants.Top10)
	c <- true
	defer close(c)

	if err != nil {
		log.Fatal(err)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "🎾  {{. | cyan }}",
		Inactive: "{{ . | red }}",
		Selected: "🎾  {{. | cyan }}",
	}

	jRes := top10{}
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(data, &jRes)

	if err != nil {
		log.Fatal(err)
	}
	final := append(jRes.Top10, "Exit")
	prompt := promptui.Select{
		Label:     utils.FilterToColor("Select Player", "green"),
		Templates: templates,
		Items:     final,
	}

	prompt.IsVimMode = true

	idx, result, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	result = final[idx]

	if result == "Exit" {
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	utils.CallClear()
	// fmt.Printf("You choose %q\n", result)
	go utils.StartLoading(c)
	parsedName := utils.ParsePlayerName(result)
	po := SearchPlayer(parsedName)
	c <- true
	utils.CallClear()

	pres := po.PrintPlayer()

	fmt.Println("\n" + pres)

    utils.WaitForEnter()

    Menu()
}

var top10Command = &cobra.Command{
	Use:   "t10",
	Short: "Print the top 10",
	Long:  `Quick way to check the current top 10`,
	Run: func(cmd *cobra.Command, args []string) {
		Top10()
	},
}
