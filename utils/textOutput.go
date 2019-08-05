package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

// TerminalSize gets the terminal size
func TerminalSize() (width, height int) {
	width, height, err := terminal.GetSize(int(os.Stdin.Fd()))

	if err != nil {
		log.Fatal(err)
	}
	return width, height
}

// CenterHor centers text horizontally
func CenterHor(input string) string {
	var ret string
	w, _ := TerminalSize()
	ret = strings.Repeat(" ", w) + input
	return ret
}

// CenterText centers text
func CenterText(input string) string {
	// w, _ := TerminalSize()
	// return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(input))/2, input))
	return fmt.Sprintf("%-55s", fmt.Sprintf("%55s", input))
}

// WaitForEnter blocks execution until the user presses enter
func WaitForEnter() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// FilterToColor formats text to colored output
func FilterToColor(input interface{}, format string) interface{} {
	var chosen func(a ...interface{}) string
	switch format {
	case "yellow":
		chosen = color.New(color.FgYellow).SprintFunc()
	case "red":
		chosen = color.New(color.FgRed).SprintFunc()
	case "green":
		chosen = color.New(color.FgGreen).SprintFunc()
	default:
		chosen = color.New(color.FgBlack).SprintFunc()
	}

	switch v := input.(type) {
	case []string:
		ret := make([]string, len(v))
		for idx, el := range v {
			ret[idx] = fmt.Sprintf("%v", chosen(el))
		}
		return ret

	case string:
		ret := fmt.Sprintf("%v", chosen(v))
		return ret

	default:
		return nil

	}

}
