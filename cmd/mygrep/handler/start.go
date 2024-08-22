package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type Start struct{}

func (s Start) Matches(app *app.App) bool {
	cur := app.Pattern[app.Ptr.PatternL]
	letter := "1234567890qwertyiuopasdfghjklzxcvbnmQWERTYIUOPASDFGHJKLZXCVBNM_"

	return strings.ContainsAny(letter, string(cur))
}

func (s Start) Run(app *app.App) bool {
	l := app.Ptr.LineL
	app.Ptr.LineL++
	return string(app.Pattern[app.Ptr.PatternL]) == string(app.Line[l])
}
