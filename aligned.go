package proto

import "strings"

type aligned struct {
	source string
	left   bool
}

var (
	alignedEquals      = leftAligned(" = ")
	alignedShortEquals = leftAligned("=")
	alignedSpace       = leftAligned(" ")
	alignedComma       = leftAligned(",")
)

func leftAligned(src string) aligned  { return aligned{src, true} }
func rightAligned(src string) aligned { return aligned{src, false} }

func (a aligned) preferredWidth() int { return len(a.source) }

func (a aligned) formatted(width int) string {
	if len(a.source) > width {
		return a.source[:width]
	}
	if a.left {
		return a.source + strings.Repeat(" ", width-len(a.source))
	}
	return strings.Repeat(" ", width-len(a.source)) + a.source
}
