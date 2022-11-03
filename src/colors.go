package main

import "github.com/fatih/color"

var (
	boldGreen      = color.New(color.Bold, color.FgGreen).PrintfFunc()
	boldRed        = color.New(color.Bold, color.FgRed).PrintfFunc()
	cyan           = color.New(color.FgCyan).PrintfFunc()
	underlineWhite = color.New(color.Underline, color.FgWhite).PrintfFunc()
	green          = color.New(color.FgGreen).PrintfFunc()
	red            = color.New(color.FgRed).PrintfFunc()
	yellow         = color.New(color.FgYellow).PrintfFunc()
	magenta        = color.New(color.FgMagenta).PrintfFunc()
)
