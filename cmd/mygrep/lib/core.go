package lib

import (
	"fmt"

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

	for matcher.Ptr.PatternL <= matcher.Ptr.PatternR && matcher.Ptr.LineL <= matcher.Ptr.LineR {
		cur := matcher.Pattern[matcher.Ptr.PatternL]
		curLine := matcher.Line[matcher.Ptr.LineL]

		if cur == '\\' {
			matcher.Ptr.PatternL++
			continue
		}

		if cur == ' ' {
			if curLine != ' ' {
				return false, fmt.Errorf("no pattern found: with space")
			}

			matcher.Ptr.LineL++
			matcher.Ptr.PatternL++
			continue
		}

		if digit.Matches(&matcher) {
			if !digit.Run(&matcher) {
				return false, fmt.Errorf("no pattern found: %q", "\\d")
			}

		} else if word.Matches(&matcher) {
			if !word.Run(&matcher) {
				return false, fmt.Errorf("no pattern found: %q", "\\w")
			}

		} else if negativeCharGrp.Matches(&matcher) {
			if !negativeCharGrp.Run(&matcher) {
				return false, fmt.Errorf("no pattern found: %q", pattern)
			}

		} else if positiveCharGrp.Matches(&matcher) {
			if !positiveCharGrp.Run(&matcher) {
				return false, fmt.Errorf("no pattern found: %q", pattern)
			}

		} else if start.Matches(&matcher) {
			if !start.Run(&matcher) {
				return false, fmt.Errorf("no pattern found: %q", pattern)
			}

		} else {
			return false, fmt.Errorf("no pattern found at last: %q", pattern)
		}

		matcher.Ptr.PatternL++
	}

	return true, nil
}
