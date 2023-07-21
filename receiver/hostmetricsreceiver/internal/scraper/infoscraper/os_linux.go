//go:build linux
// +build linux

package infoscraper

import (
	"os"
	"runtime"
	"strings"
)

// Ubuntu 20.04 LTS
/**
NAME="Ubuntu"
VERSION="20.04.2 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.04.2 LTS"
VERSION_ID="20.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
*/

// CentOS 8
/**
NAME="CentOS Linux"
VERSION="8 (Core)"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8 (Core)"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
*/

// Debian 10
/**
PRETTY_NAME="Debian GNU/Linux 10 (buster)"
NAME="Debian GNU/Linux"
VERSION_ID="10"
VERSION="10 (buster)"
VERSION_CODENAME=buster
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
*/

// Fedora 34
/**
NAME=Fedora
VERSION="34 (Workstation Edition)"
ID=fedora
VERSION_ID=34
VERSION_CODENAME=""
PLATFORM_ID="platform:f34"
PRETTY_NAME="Fedora 34 (Workstation Edition)"
ANSI_COLOR="0;34"
LOGO=fedora-logo-icon
CPE_NAME="cpe:/o:fedoraproject:fedora:34"
HOME_URL="https://fedoraproject.org/"
DOCUMENTATION_URL="https://docs.fedoraproject.org/en-US/fedora/34/"
SUPPORT_URL="https://fedoraproject.org/wiki/Communicating_and_getting_help"
BUG_REPORT_URL="https://bugzilla.redhat.com/"
REDHAT_BUGZILLA_PRODUCT="Fedora"
REDHAT_BUGZILLA_PRODUCT_VERSION=34
REDHAT_SUPPORT_PRODUCT="Fedora"
REDHAT_SUPPORT_PRODUCT_VERSION=34
PRIVACY_POLICY_URL="https://fedoraproject.org/wiki/Legal:PrivacyPolicy"
*/

// Oracle Linux
/**
NAME="Oracle Linux Server"
VERSION="9.2"
ID="ol"
ID_LIKE="fedora"
VARIANT="Server"
VARIANT_ID="server"
VERSION_ID="9.2"
PLATFORM_ID="platform:el9"
PRETTY_NAME="Oracle Linux Server 9.2"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:oracle:linux:9:2:server"
HOME_URL="https://linux.oracle.com/"
BUG_REPORT_URL="https://github.com/oracle/oracle-linux"

ORACLE_BUGZILLA_PRODUCT="Oracle Linux 9"
ORACLE_BUGZILLA_PRODUCT_VERSION=9.2
ORACLE_SUPPORT_PRODUCT="Oracle Linux"
ORACLE_SUPPORT_PRODUCT_VERSION=9.2
*/
func _getSystemInfo() (System, error) {
	osReleaseContent, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return System{}, err
	}
	lines := strings.Split(string(osReleaseContent), "\n")
	distribution := "unknown"
	version := ""

	for _, line := range lines {
		if strings.HasPrefix(line, "NAME=") {
			distribution = strings.TrimPrefix(line, "NAME=")
		}
		if strings.HasPrefix(line, "VERSION_ID=") {
			version = strings.TrimPrefix(line, "VERSION_ID=")
		}
	}

	return System{
		Distribution: distribution,
		Arch:         runtime.GOARCH,
		Version:      version,
	}, nil
}
