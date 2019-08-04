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

var top10Command = &cobra.Command{
	Use:   "t10",
	Short: "Print the top 10",
	Long:  `Quick way to check the current top 10`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.CallClear()
		c := make(chan bool)
		go utils.StartLoading(c)
		res, err := http.Get(constants.APIHead + constants.Top10)
		c <- true
		defer close(c)

		if err != nil {
			log.Fatal(err)
		}

		jres := top10{}
		data, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(data, &jres)

		prompt := promptui.Select{
			Label: "Select Player",
			Items: append(jres.Top10, "Exit"),
		}

		prompt.IsVimMode = true

		_, result, err := prompt.Run()

		if err != nil {
			log.Fatal(err)
		}

		if result == "Exit" {
			fmt.Println("Exiting...")
			os.Exit(1)
		}

		utils.CallClear()
		fmt.Printf("You choose %q\n", result)
		go utils.StartLoading(c)
		parsed_name := utils.ParsePlayerName(result)
		fmt.Println(parsed_name)
		po := SearchPlayer(parsed_name)
		c <- true
		utils.CallClear()

		pres := po.PrintPlayer()

		fmt.Println("\n" + pres)

	},
}
