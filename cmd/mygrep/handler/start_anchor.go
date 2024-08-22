package handler

import (
	"strings"

	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
)

type StartAnchor struct{}

func (s StartAnchor) Matches(app *app.App) bool {
	return app.Ptr.PatternL == 0 && app.Pattern[app.Ptr.PatternL] == '^'
}

func (s StartAnchor) Run(app *app.App) bool {
	letter := "1234567890qwertyiuopasdfghjklzxcvbnmQWERTYIUOPASDFGHJKLZXCVBNM_"

	app.Ptr.PatternL++

	for strings.ContainsAny(letter, string(app.Pattern[app.Ptr.PatternL])) {
		if app.Pattern[app.Ptr.PatternL] != app.Line[app.Ptr.LineL] {
			return false
		}

		if app.Ptr.PatternL+1 > app.Ptr.PatternR {
			break
		}

		if app.Ptr.LineL+1 > app.Ptr.LineR {
			break
		}

		app.Ptr.PatternL++
		app.Ptr.LineL++
	}

	return true
}
