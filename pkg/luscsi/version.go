package luscsi

import (
	"fmt"
	"runtime"
	"strings"

	"sigs.k8s.io/yaml"
)

// These are set during build time via -ldflags
var (
	driverVersion = "N/A"
	gitCommit     = "N/A"
	buildDate     = "N/A"
)

// VersionInfo holds the version information of the driver
type VersionInfo struct {
	DriverName    string `json:"Driver Name"`
	DriverVersion string `json:"Driver Version"`
	GitCommit     string `json:"Git Commit"`
	BuildDate     string `json:"Build Date"`
	GoVersion     string `json:"Go Version"`
	Compiler      string `json:"Compiler"`
	Platform      string `json:"Platform"`
}

// GetVersion returns the version information of the driver
func GetVersion(driverName string) VersionInfo {
	return VersionInfo{
		DriverName:    driverName,
		DriverVersion: driverVersion,
		GitCommit:     gitCommit,
		BuildDate:     buildDate,
		GoVersion:     runtime.Version(),
		Compiler:      runtime.Compiler,
		Platform:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// GetVersionYAML returns the version information of the driver
// in YAML format
func GetVersionYAML(driverName string) (string, error) {
	info := GetVersion(driverName)
	marshalled, err := yaml.Marshal(&info)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(marshalled)), nil
}

// GetUserAgent returns user agent of the driver
func GetUserAgent(driverName, customUserAgent, userAgentSuffix string) string {
	customUserAgent = strings.TrimSpace(customUserAgent)
	userAgent := customUserAgent
	if customUserAgent == "" {
		userAgent = fmt.Sprintf("%s/%s %s/%s (%s-%s)", driverName, driverVersion, runtime.Compiler, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	}

	userAgentSuffix = strings.TrimSpace(userAgentSuffix)
	if userAgentSuffix != "" {
		userAgent = userAgent + " " + userAgentSuffix
	}
	return userAgent
}
