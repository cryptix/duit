package duit

import (
	draw "github.com/ktye/duitdraw"
)

// Icon is a single codepoint in the given font. Typically for an "icon font" like fontawesome.
type Icon struct {
	Rune rune       // Codepoint to draw.
	Font *draw.Font `json:"-"` // Font to draw in. If nil, nothing is typically drawn.
}
