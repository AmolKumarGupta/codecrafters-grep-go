package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type Digit struct {
}

func (d Digit) Matches(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]

	if cur != 'd' {
		return false
	}

	if app.Ptr.PatternL-1 < 0 {
		return false
	}

	prev := app.Pattern[app.Ptr.PatternL-1]
	return prev == '\\'
}

func (d Digit) Run(app *app.App) bool {
	// return strings.ContainsAny("0123456789", string(app.Line[app.Ptr.LineL]))
	if strings.ContainsAny("0123456789", string(app.Line[app.Ptr.LineL])) {
		app.Ptr.LineL++
		return true
	}

	return false
}
