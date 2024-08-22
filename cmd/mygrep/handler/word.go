package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type Word struct{}

func (d Word) Matches(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]

	if cur != 'w' {
		return false
	}

	if app.Ptr.PatternL-1 < 0 {
		return false
	}

	prev := app.Pattern[app.Ptr.PatternL-1]
	return prev == '\\'
}

func (d Word) Run(app *app.App) bool {
	letter := "1234567890qwertyiuopasdfghjklzxcvbnmQWERTYIUOPASDFGHJKLZXCVBNM_"

	if strings.ContainsAny(letter, string(app.Line[app.Ptr.LineL])) {
		app.Ptr.LineL++
		return true
	}

	return false
}
