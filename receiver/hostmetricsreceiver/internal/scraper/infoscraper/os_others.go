//go:build windows
// +build windows

package infoscraper

import (
	"runtime"
)

func _getSystemInfo() (System, error) {
	return System{
		Distribution: "unimplemented",
		Arch:         runtime.GOARCH,
		Version:      "",
	}, nil
}
