package main

import (
	"bufio"
	"github.com/fatih/color"
	"os"
)

var questionColor = color.New(color.FgCyan)

func prompt(s string) string {
	_, _ = questionColor.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func promptCheckbox(s string) bool {
	_, _ = questionColor.Print(s)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	switch scanner.Text() {
	case "y", "Y":
		return true
	default:
		return false
	}
}
