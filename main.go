package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var attributeMap = map[string]color.Attribute{
	"red":         color.FgRed,
	"green":       color.FgGreen,
	"blue":        color.FgBlue,
	"yellow":      color.FgYellow,
	"magenta":     color.FgMagenta,
	"cyan":        color.FgCyan,
	"white":       color.FgWhite,
	"black":       color.FgBlack,
	"hired":       color.FgHiRed,
	"higreen":     color.FgHiGreen,
	"hiblue":      color.FgHiBlue,
	"hiyellow":    color.FgHiYellow,
	"himagenta":   color.FgHiMagenta,
	"hicyan":      color.FgHiCyan,
	"hiwhite":     color.FgHiWhite,
	"bgred":       color.BgRed,
	"bggreen":     color.BgGreen,
	"bgblue":      color.BgBlue,
	"bgyellow":    color.BgYellow,
	"bgmagenta":   color.BgMagenta,
	"bgcyan":      color.BgCyan,
	"bgwhite":     color.BgWhite,
	"bghired":     color.BgHiRed,
	"bghigreen":   color.BgHiGreen,
	"bghiblue":    color.BgHiBlue,
	"bghiyellow":  color.BgHiYellow,
	"bghimagenta": color.BgHiMagenta,
	"bghicyan":    color.BgHiCyan,
	"bghiwhite":   color.BgHiWhite,
	"bold":        color.Bold,
	"underline":   color.Underline,
	"italic":      color.Italic,
	"faint":       color.Faint,
	"crossedout":  color.CrossedOut,
	"blinkslow":   color.BlinkSlow,
	"blinkrapid":  color.BlinkRapid,
	"reverse":     color.ReverseVideo,
	"concealed":   color.Concealed,
}

func mustParseOpts(args []string) *color.Color {
	c := color.New()
	for _, arg := range args {
		if attr, ok := attributeMap[strings.ToLower(arg)]; ok {
			c.Add(attr)
		} else {
			fatal("error: unknown attribute/color:", arg)
		}
	}
	return c
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: color [options]")
		fmt.Println(`Example: echo "hello" | color blue bgyellow bold italic`)
		return
	}

	c := mustParseOpts(os.Args[1:])

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

func fatal(v ...any) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}
