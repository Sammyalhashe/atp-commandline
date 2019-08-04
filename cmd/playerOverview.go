package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Sammyalhashe/gomod/utils"
	"io/ioutil"
	"net/http"
	//"reflect"

	"github.com/Sammyalhashe/gomod/constants"
	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playerOverviewCommand)
}

// PlayerOverview struct defining player overview return json
type PlayerOverview struct {
	Player []struct {
		Name  string `json:"name"`
		Stats struct {
			Player struct {
				Name             string `json:"name"`
				Bio              string `json:"bio"`
				CurrentYearStats struct {
					DataSingles struct {
						Rank       int   `json:"rank"`
						RankMove   int   `json:"rank_move"`
						WL         []int `json:"w-l"`
						Titles     int   `json:"titles"`
						PrizeMoney float64   `json:"prize_money"`
					} `json:"data-singles"`
					DataDoubles struct {
						Rank       int   `json:"rank"`
						RankMove   int   `json:"rank_move"`
						WL         []int `json:"w-l"`
						Titles     int   `json:"titles"`
						PrizeMoney float64   `json:"prize_money"`
					} `json:"data-doubles"`
					Year string `json:"year"`
				} `json:"current_year_stats"`
				CareerStats struct {
					DataSingles struct {
						Rank        int    `json:"rank"`
						DateHighest string `json:"date_highest"`
						WL          []int  `json:"w-l"`
						Titles      int    `json:"titles"`
						PrizeMoney  float64    `json:"prize_money"`
					} `json:"data-singles"`
					DataDoubles struct {
						Rank        int    `json:"rank"`
						DateHighest string `json:"date_highest"`
						WL          []int  `json:"w-l"`
						Titles      int    `json:"titles"`
						PrizeMoney  float64    `json:"prize_money"`
					} `json:"data-doubles"`
				} `json:"career_stats"`
				Fundamentals struct {
					Age struct {
						Age       int    `json:"age"`
						Birthdate string `json:"birthdate"`
					} `json:"age"`
					TurnedPro int `json:"turned_pro"`
					Weight    struct {
						Lbs float64 `json:"lbs"`
						Kg  float64 `json:"kg"`
					} `json:"weight"`
					Height struct {
						Ft string `json:"ft"`
						Cm float64    `json:"cm"`
					} `json:"height"`
					Birthplace string   `json:"birthplace"`
					Residence  string   `json:"residence"`
					Coach      []string `json:"coach"`
					Plays      struct {
						Hand     string `json:"hand"`
						Backhand string `json:"backhand"`
					} `json:"plays"`
				} `json:"fundamentals"`
			} `json:"player"`
		} `json:"stats"`
	} `json:"player"`
}

// PrintPlayer Prints the player overview with a specific format
func (p PlayerOverview) PrintPlayer() string {
	player := p.Player[0].Stats.Player

	//v := reflect.ValueOf(player)
	//
	//for i := 0; i < v.NumField(); i++ {
	//	fmt.Println(v.Field(i))
	//}

	var res string

	res += "Player: " + player.Name + "\n"
	res += "Bio " + player.Bio + "\n"

	currYear := player.CurrentYearStats
	res += "This Year: " + currYear.Year + "\n"

	res += "\tSingles:\n"
	singles := currYear.DataSingles
	fmt.Println(singles.Rank)
	res += "\t\tRank: " + fmt.Sprintf("%d\n", singles.Rank)
	res += "\t\tRankMove: " + fmt.Sprintf("%d\n", singles.RankMove)
	res += "\t\tWin-Loss: " + fmt.Sprintf("%d - ", singles.WL[0]) + fmt.Sprintf("%d\n", singles.WL[1])
	res += "\t\tTitles: " + fmt.Sprintf("%d\n", singles.Titles)
	res += "\t\tPrize Money: " + fmt.Sprintf("%f\n", singles.PrizeMoney)

	res += "\tDoubles:\n"
	doubles := currYear.DataDoubles
	res += "\t\tRank: " + fmt.Sprintf("%d\n", doubles.Rank)
	res += "\t\tRankMove: " + fmt.Sprintf("%d\n", doubles.RankMove)
	res += "\t\tWin-Loss: " + fmt.Sprintf("%d - ", doubles.WL[0]) + fmt.Sprintf("%d\n", doubles.WL[1])
	res += "\t\tTitles: " + fmt.Sprintf("%d\n", doubles.Titles)
	res += "\t\tPrize Money: " + fmt.Sprintf("%f\n", doubles.PrizeMoney)

	res += "Career:\n"
	career := player.CareerStats

	res += "\tSingles:\n"
	singlesCarr := career.DataSingles
	res += "\t\tRank: " + fmt.Sprintf("%d\n", singlesCarr.Rank)
	res += "\t\tDate Highest: " + fmt.Sprintf("%v\n", singlesCarr.DateHighest)
	res += "\t\tWin-Loss: " + fmt.Sprintf("%d - ", singlesCarr.WL[0]) + fmt.Sprintf("%d\n", singlesCarr.WL[1])
	res += "\t\tTitles: " + fmt.Sprintf("%d\n", singlesCarr.Titles)
	res += "\t\tPrize Money: " + fmt.Sprintf("%f\n", singlesCarr.PrizeMoney)

	res += "\tDoubles:\n"
	doublesCarr := career.DataDoubles
	res += "\t\tRank: " + fmt.Sprintf("%d\n", doublesCarr.Rank)
	res += "\t\tDate Highest: " + fmt.Sprintf("%v\n", doublesCarr.DateHighest)
	res += "\t\tWin-Loss: " + fmt.Sprintf("%d - ", doublesCarr.WL[0]) + fmt.Sprintf("%d\n", doublesCarr.WL[1])
	res += "\t\tTitles: " + fmt.Sprintf("%d\n", doublesCarr.Titles)
	res += "\t\tPrize Money: " + fmt.Sprintf("%f\n", doublesCarr.PrizeMoney)

	return res

}

// SearchPlayer creates an http request to atp-scraper for the player overview
func SearchPlayer(name string) PlayerOverview {
	res, err := http.Get(constants.APIHead + constants.PlayerOverview + name)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	jRes := PlayerOverview{}

	err = json.Unmarshal(data, &jRes)

	if err != nil {
		log.Fatal(err)
	}

	return jRes
}

var playerOverviewCommand = &cobra.Command{
	Use:   "po",
	Short: "Gets a player overview",
	Long:  `Gets a player overview, which includes info such as player rank, earnings, coach, etc...`,
	Run:   func(cmd *cobra.Command, args []string) {
		validate := func(input string) error {
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Enter Player Name",
			Validate: validate,
		}

		utils.CallClear()
		result, err := prompt.Run()

		if err != nil {
			fmt.Println("Prompt failed...")
			log.Fatal(err)
		}

		utils.CallClear()
		fmt.Printf("You choose %q\n", result)

		c := make(chan bool)
		defer close(c)
		go utils.StartLoading(c)
		playerOverview := SearchPlayer(utils.ParsePlayerName(result))
		c <- true
		utils.CallClear()

		pres := playerOverview.PrintPlayer()
		fmt.Println("\n" + pres)


	},
}
