package lib

import (
	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/app"
	"github.com/codecrafters-io/grep-starter-go/cmd/mygrep/handler"
)

func MatchLine(line []byte, pattern string) (bool, error) {
	matcher := app.App{
		Line:    line,
		Pattern: pattern,
		Ptr: &app.Ptr{
			LineL:    0,
			LineR:    len(line) - 1,
			PatternL: 0,
			PatternR: len(pattern) - 1,
		},
	}

	digit := handler.Digit{}
	word := handler.Word{}
	start := handler.Start{}
	positiveCharGrp := handler.PositiveCharGroup{}
	negativeCharGrp := handler.NegativeCharactorGroup{}
	startAnchor := handler.StartAnchor{}

	for matcher.Ptr.PatternL <= matcher.Ptr.PatternR && matcher.Ptr.LineL <= matcher.Ptr.LineR {
		cur := matcher.Pattern[matcher.Ptr.PatternL]
		curLine := matcher.Line[matcher.Ptr.LineL]

		if cur == '\\' {
			if matcher.Ptr.PatternL < matcher.Ptr.PatternR && matcher.Pattern[matcher.Ptr.PatternL+1] == '\\' {
				if curLine != '\\' {
					return false, nil
				}
			}

			matcher.Ptr.PatternL++
			continue
		}

		if cur == ' ' {
			if curLine != ' ' {
				return false, nil
			}

			matcher.Ptr.LineL++
			matcher.Ptr.PatternL++
			continue
		}

		if startAnchor.Matches(&matcher) {
			if !startAnchor.Run(&matcher) {
				return false, nil
			}

		} else if digit.Matches(&matcher) {
			if !digit.Run(&matcher) {
				matcher.Ptr.LineL++
				continue
			}

		} else if word.Matches(&matcher) {
			if !word.Run(&matcher) {
				matcher.Ptr.LineL++
				continue
			}

		} else if negativeCharGrp.Matches(&matcher) {
			if !negativeCharGrp.Run(&matcher) {
				return false, nil
			}

		} else if positiveCharGrp.Matches(&matcher) {
			if !positiveCharGrp.Run(&matcher) {
				return false, nil
			}

		} else if start.Matches(&matcher) {
			if !start.Run(&matcher) {
				return false, nil
				// matcher.Ptr.LineL++
				// continue
			}

		} else {
			return false, nil
		}

		matcher.Ptr.PatternL++
	}

	if matcher.Ptr.PatternL <= matcher.Ptr.PatternR {
		return false, nil
	}

	return true, nil
}
