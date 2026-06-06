package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgGreen).Sprintf(format, a...))
}

func Error(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgRed).Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgYellow).Sprintf(format, a...))
}

func Info(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgCyan).Sprintf(format, a...))
}

func Dim(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.Faint).Sprintf(format, a...))
}

func Highlight(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.BgHiRed, color.FgWhite).Sprintf(format, a...))
}

func ListItems(strings []string) {

	if len(strings) == 0 {
		Dim("None")
		return
	}

	for i, s := range strings {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(s)
	}
	fmt.Println()
}

func Prompt(format string, a ...any) string {
	fmt.Printf("%s", color.New(color.FgBlue).Sprintf(format, a...))
	var input string
	fmt.Scanln(&input)
	return input
}
