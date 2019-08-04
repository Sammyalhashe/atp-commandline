package utils

import "strings"

// ParsePlayerName parses the player name so it fits the api sepcification
func ParsePlayerName(name string) (ret string) {
	ret = strings.TrimSpace(name)
	ret = strings.ToLower(ret)
	ret = strings.ReplaceAll(ret, " ", "-")

	return ret
}
