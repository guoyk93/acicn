package gg

import (
	"os"
	"strings"
)

const (
	DebugWildcard  = "*"
	DebugSeparator = ":"
)

type DebugSet interface {
	Debug(v string) bool
}

type debugSet struct {
	full  bool
	none  bool
	items map[string]struct{}
}

// Debug check weather a debug flag is enabled
func (d *debugSet) Debug(v string) bool {
	if d.full {
		return true
	}

	if d.none {
		return false
	}

	if _, ok := d.items[v]; ok {
		return true
	}

	if _, ok := d.items[DebugWildcard]; ok {
		return true
	}

	for {
		if _, ok := d.items[v+DebugSeparator+DebugWildcard]; ok {
			return true
		}

		if idx := strings.LastIndex(v, DebugSeparator); idx < 0 {
			break
		} else {
			v = v[:idx]
		}
	}

	return false
}

// NewDebugSet create a new DebugSet
// use comma ',' separate multiple items
// use '*' for wildcard, for example '*', 'echo:*,view:*'
func NewDebugSet(debug string) DebugSet {
	debug = strings.ToLower(strings.TrimSpace(debug))
	if debug == "" {
		return &debugSet{none: true}
	}
	if debug == DebugWildcard {
		return &debugSet{full: true}
	}
	items := map[string]struct{}{}
	for _, item := range strings.Split(debug, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		items[item] = struct{}{}
	}
	return &debugSet{items: items}
}

var (
	debug = NewDebugSet(os.Getenv("DEBUG"))
)

// Debug check weather a debug flag is enabled from environment variable $DEBUG
func Debug(v string) bool {
	return debug.Debug(v)
}
