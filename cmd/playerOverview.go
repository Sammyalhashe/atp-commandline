package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sammyalhashe/gomod/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand()
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
						PrizeMoney int   `json:"prize_money"`
					} `json:"data-singles"`
					DataDoubles struct {
						Rank       int   `json:"rank"`
						RankMove   int   `json:"rank_move"`
						WL         []int `json:"w-l"`
						Titles     int   `json:"titles"`
						PrizeMoney int   `json:"prize_money"`
					} `json:"data-doubles"`
					Year string `json:"year"`
				} `json:"current_year_stats"`
				CareerStats struct {
					DataSingles struct {
						Rank        int    `json:"rank"`
						DateHighest string `json:"date_highest"`
						WL          []int  `json:"w-l"`
						Titles      int    `json:"titles"`
						PrizeMoney  int    `json:"prize_money"`
					} `json:"data-singles"`
					DataDoubles struct {
						Rank        int    `json:"rank"`
						DateHighest string `json:"date_highest"`
						WL          []int  `json:"w-l"`
						Titles      int    `json:"titles"`
						PrizeMoney  int    `json:"prize_money"`
					} `json:"data-doubles"`
				} `json:"career_stats"`
				Fundamentals struct {
					Age struct {
						Age       int    `json:"age"`
						Birthdate string `json:"birthdate"`
					} `json:"age"`
					TurnedPro int `json:"turned_pro"`
					Weight    struct {
						Lbs int `json:"lbs"`
						Kg  int `json:"kg"`
					} `json:"weight"`
					Height struct {
						Ft string `json:"ft"`
						Cm int    `json:"cm"`
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

	var res string

	res += "Player: " + player.Name + "\n"
	res += "Bio " + player.Bio + "\n"

	return res

}

// SearchPlayer creates an http request to atp-scraper for the player overview
func SearchPlayer(name string) PlayerOverview {
	res, err := http.Get(constants.APIHead + constants.PlayerOverview + name)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	jres := PlayerOverview{}

	json.Unmarshal(data, &jres)

	return jres
}

var playerOverviewCommand = &cobra.Command{
	Use:   "player-overview",
	Short: "Gets a player overview",
	Long:  `Gets a player overview, which includes info such as player rank, earnings, coach, etc...`,
	Run:   func(cmd *cobra.Command, args []string) {},
}
