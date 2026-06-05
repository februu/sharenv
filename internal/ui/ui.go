package ui

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	dimStyle  = color.New(color.Faint)
	boldStyle = color.New(color.Bold)
	keyStyle  = color.New(color.FgCyan, color.Bold)
)

func Success(format string, a ...any) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgRed).Sprintf(format, a...))
}

func Warn(format string, a ...any) {
	fmt.Printf("%s\n", color.New(color.FgYellow).Sprintf(format, a...))
}

func Info(format string, a ...any) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}

func Header(text string) {
	fmt.Println()
	fmt.Println(text)
}

func ListItem(name string) {
	fmt.Printf("- %s\n", boldStyle.Sprint(name))
}

func KeyDisplay(label, value string) {
	fmt.Println()
	fmt.Printf("  %s\n", dimStyle.Sprint(label))
	fmt.Printf("  %s\n", keyStyle.Sprint(value))
	fmt.Println()
}
