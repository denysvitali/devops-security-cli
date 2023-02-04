package main

import (
	"github.com/alexeyco/simpletable"
	"strings"
)

var DefaultStyle = simpletable.StyleCompactClassic

func GetTableStyle() *simpletable.Style {
	if config == nil {
		return DefaultStyle
	}

	switch strings.ToLower(config.TableFormat) {
	case "markdown", "md":
		return simpletable.StyleMarkdown
	case "rounded":
		return simpletable.StyleRounded
	case "compact":
		return simpletable.StyleCompact
	case "compact-lite":
		return simpletable.StyleCompactLite
	case "compact-classic":
		return simpletable.StyleCompactClassic
	default:
		return DefaultStyle
	}
}
