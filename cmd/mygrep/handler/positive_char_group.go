package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type PositiveCharGroup struct{}

func (p PositiveCharGroup) Matches(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]
	return cur == '['
}

func (p PositiveCharGroup) Run(app *app.App) bool {
	start := app.Ptr.PatternL + 1

	for app.Pattern[app.Ptr.PatternL] != ']' {
		app.Ptr.PatternL++
	}

	searchable := app.Pattern[start:app.Ptr.PatternL]
	line := app.Line[app.Ptr.LineL : app.Ptr.LineR+1]

	return strings.ContainsAny(string(line), searchable)
}
