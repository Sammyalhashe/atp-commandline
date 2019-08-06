package cmd

import (
	"encoding/json"
	"errors"
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
	rootCmd.AddCommand(titlesCommand)
}

// ItemStruct struct that details a year
type ItemStruct struct {
	Year   int `json:"year"`
	Titles struct {
		Number      int `json:"number"`
		Tournaments []struct {
			Tournament string `json:"tournament"`
			URL        string `json:"url"`
		} `json:"tournaments"`
	} `json:"titles"`
	Finals struct {
		Number      int `json:"number"`
		Tournaments []struct {
			Tournament string `json:"tournament"`
			URL        string `json:"url"`
		} `json:"tournaments"`
	} `json:"finals"`
}

// TitlesJSON struct modelling return from titles request
type TitlesJSON struct {
	Singles struct {
		Type  string       `json:"type"`
		Items []ItemStruct `json:"items"`
	} `json:"singles"`
}

func grabYears(titleStruct TitlesJSON) []string {
	items := titleStruct.Singles.Items
	n := len(items)
	ret := make([]string, n)
	for idx, el := range items {
		ret[idx] = fmt.Sprintf("%d", el.Year)
	}

	return ret
}

func printItem(item ItemStruct) string {
	var ret string
	ret += fmt.Sprintf("Year: %d\n", item.Year)
	ret += "Titles:\n"
	ret += fmt.Sprintf("\tNumber: %d\n", item.Titles.Number)
	tournaments := item.Titles.Tournaments
	for _, tourny := range tournaments {
		ret += fmt.Sprintf("\t\tTournament: %v\n", tourny.Tournament)
		ret += fmt.Sprintf("\t\tURL: %v\n", constants.ATPUrl + tourny.URL)
	}
	ret += "Finals:\n"
	ret += fmt.Sprintf("\tNumber: %d\n", item.Finals.Number)
	finals := item.Finals.Tournaments
	for _, tourny := range finals {
		ret += fmt.Sprintf("\t\tTournament: %v\n", tourny.Tournament)
		ret += fmt.Sprintf("\t\t\tURL: %v\n", tourny.URL)
	}
	return ret
}

// Titles function performs the operation to get players titles
func Titles() {
	validate := func(input string) error {
		if !ValidName(input) {
			return errors.New("Must be valid name")
		}
		return nil
	}

	// promptTemaplates := &promptui.PromptTemplates{
	// 	Prompt:  "{{ . }}",
	// 	Valid:   "{{ . | bold  }}",
	// 	Invalid: "{{ . | red }}",
	// 	Success: "{{ . | bold }}",
	// }

	prompt1 := promptui.Prompt{
		Label: utils.FilterToColor("Enter Player Name", "green"),
		// Templates: promptTemaplates,
		Validate: validate,
	}
	result1, err1 := prompt1.Run()
	if err1 != nil {
		log.Fatal(err1)
	}

	name := utils.ParsePlayerName(result1)
	utils.CallClear()
	c := make(chan bool)
	go utils.StartLoading(c)
	url := constants.APIHead + "/" + name + constants.PlayerTitlesURL
	res, err := http.Get(url)
	c <- true
	defer close(c)

	if err != nil {
		log.Fatal(err)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . | green }}",
		Active:   "🎾  {{. | cyan }}",
		Inactive: "{{ . | red }}",
		Selected: "🎾  {{. | cyan }}",
	}

	jRes := TitlesJSON{}
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(data, &jRes)

	years := append(grabYears(jRes), "Exit")

	prompt := promptui.Select{
		Label:     "Select year you want to break down",
		Templates: templates,
		Items:     years,
	}

	prompt.IsVimMode = true

	idx, _, err := prompt.Run()

	if err != nil {
		log.Fatal(err)
	}

	result := years[idx]

	if result == "Exit" {
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	fRes := jRes.Singles.Items[idx]
	fString := printItem(fRes)
	fmt.Println(fString)

	utils.WaitForEnter()
	Menu()

}

var titlesCommand = &cobra.Command{
	Use:   "titles",
	Short: "Get the titles and finals for a player",
	Long:  "Get the titles and finals for a player",
	Run: func(cmd *cobra.Command, args []string) {
		Titles()
	},
}
