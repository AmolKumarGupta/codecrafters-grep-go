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

	if l == 0 {
		s := string(app.Pattern[app.Ptr.PatternL])
		for app.Ptr.LineL <= app.Ptr.LineR {
			if s == string(app.Line[app.Ptr.LineL]) {
				app.Ptr.LineL++
				return true
			}
			app.Ptr.LineL++
		}

		return false
	}

	app.Ptr.LineL++
	return string(app.Pattern[app.Ptr.PatternL]) == string(app.Line[l])
}
