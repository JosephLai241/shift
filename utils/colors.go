// Defining terminal text color and style.

package utils

import "github.com/fatih/color"

var Blue = color.New(color.FgBlue).Add(color.Bold)
var Green = color.New(color.FgGreen).Add(color.Bold)
var Red = color.New(color.FgRed).Add((color.Bold))
var White = color.New(color.FgWhite).Add(color.Bold)
var Yellow = color.New(color.FgYellow).Add(color.Bold)

var BlueSprint = color.New(color.FgBlue).Add(color.Bold).SprintFunc()
var WhiteSprint = color.New(color.FgWhite).Add(color.Bold).SprintFunc()
var YellowSprint = color.New(color.FgYellow).Add(color.Bold).SprintFunc()
