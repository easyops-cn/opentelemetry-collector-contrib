//go:build darwin
// +build darwin

package infoscraper

import (
	"os/exec"
	"runtime"
	"strings"
)

func _getSystemInfo() (System, error) {
	// 获取操作系统版本
	osVersionCmd := exec.Command("sw_vers", "-productVersion")
	osVersionOutput, err := osVersionCmd.Output()
	if err != nil {
		return System{}, err
	}
	version := strings.TrimSpace(string(osVersionOutput))

	return System{
		Distribution: "MacOS",
		Arch:         runtime.GOARCH,
		Version:      version,
	}, nil
}
